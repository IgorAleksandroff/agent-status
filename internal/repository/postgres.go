package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hash/fnv"
	"log"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	querySaveUser = `INSERT INTO agents (login, password, status_id) VALUES ($1, $2, $3)
		ON CONFLICT (login) DO NOTHING`
	queryGetUser = `SELECT login, password, status, created_at FROM agents WHERE login = $1`

	queryLockForID     = `select pg_advisory_xact_lock($1)`
	queryGetAgent      = `SELECT login, status FROM agents WHERE login = $1`
	queryIsTransaction = `SELECT 1 FROM transitions WHERE source = $1 AND destination = $2 AND mode = $3`
	queryUpdateAgent   = `UPDATE agents SET status = $2, WHERE login = $1`
	queryLog           = `INSERT INTO transitions_log (agent_login, source, destination, mode) VALUES ($1, $2, $3, $4)
		returning id`
)

type pgRepo struct {
	db *sqlx.DB
}

var ErrUserRegister = errors.New("user already exist")

func NewPostgres(ctx context.Context, addressDB string) (*pgRepo, error) {
	db, err := sqlx.Connect("postgres", addressDB)
	if err != nil {
		return nil, errors.Wrap(err, "failed postgres connect")
	}

	repo := pgRepo{db: db}
	if err = repo.init(ctx); err != nil {
		return nil, errors.Wrap(err, "failed postgres init")
	}

	return &repo, nil
}

// todo сделать через миграции
func (p pgRepo) init(ctx context.Context) error {
	_, err := p.db.ExecContext(ctx, queryCreateTables)
	if err != nil {
		return err
	}
	return nil
}

func (p pgRepo) Close() error {
	return p.db.Close()
}

func (p pgRepo) SaveUser(ctx context.Context, user entity.Agent) error {
	res, err := p.db.ExecContext(ctx, querySaveUser,
		user.Login,
		*user.Password,
	)
	if err != nil {
		return fmt.Errorf("error to save user: %w, %+v", err, user)

	}
	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error get rows affected: %w", err)
	}
	if affected == 0 {
		return ErrUserRegister
	}

	return nil
}

func (p pgRepo) GetUser(ctx context.Context, login string) (entity.Agent, error) {
	var user entity.Agent

	err := p.db.QueryRowContext(
		ctx,
		queryGetUser,
		login,
	).Scan(&user.Login, user.Password, user.Status, &user.CreatedAt)
	if err != nil {
		return entity.Agent{}, fmt.Errorf("error to get user: %w, %s", err, login)
	}

	return user, nil
}

func (p pgRepo) AgentSetStatusTx(ctx context.Context, agent entity.Agent, mode entity.Mode) (*int64, error) {
	var logID *int64

	tx, err := p.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	transaction := entity.Transition{
		Dst:  *agent.Status,
		Mode: mode,
	}

	callbacks := []func(ctx context.Context, tx *sqlx.Tx) error{
		// защита от гонки при изменении статуса
		func(ctx context.Context, tx *sqlx.Tx) error {
			uid, err := hash64(agent.Login)
			if err != nil {
				return err
			}

			return p.lockForUID(ctx, tx, uid)
		},
		// получение актуального статуса
		func(ctx context.Context, tx *sqlx.Tx) error {
			a, err := p.getAgent(ctx, tx, agent.Login)
			if err != nil {
				return err
			}

			transaction.Src = *a.Status

			return err
		},
		// поверка валидности перехода
		func(ctx context.Context, tx *sqlx.Tx) error {
			return p.isTransactions(ctx, tx, transaction)
		},
		func(ctx context.Context, tx *sqlx.Tx) error {
			return p.setStatus(ctx, tx, agent.Login, *agent.Status)
		},
		func(ctx context.Context, tx *sqlx.Tx) error {
			logID, err = p.transactionLog(ctx, tx, entity.Logs{
				Agent:      agent.Login,
				Transition: transaction,
			})

			return err
		},
	}

	for _, cb := range callbacks {
		err = cb(ctx, tx)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Println("error rolling back a transaction:", rollbackErr)
			}

			return nil, errors.WithStack(err)
		}
	}

	return logID, errors.WithStack(tx.Commit())
}

func (p pgRepo) lockForUID(ctx context.Context, tx *sqlx.Tx, uid *int64) error {
	if _, err := tx.ExecContext(ctx, queryLockForID, uid); err != nil {
		return err
	}

	return nil
}

func (p pgRepo) getAgent(ctx context.Context, tx *sqlx.Tx, login string) (entity.Agent, error) {
	var a entity.Agent

	err := tx.QueryRowContext(
		ctx,
		queryGetAgent,
		login,
	).Scan(&a.Login, &a.Status)
	if err != nil {
		return entity.Agent{}, fmt.Errorf("error to get order: %w, %s", err, login)
	}

	return a, nil
}

func (p pgRepo) isTransactions(ctx context.Context, tx *sqlx.Tx, tr entity.Transition) error {
	row := tx.QueryRowContext(ctx, queryIsTransaction,
		tr.Src,
		tr.Dst,
		tr.Mode,
	)
	if errors.Is(row.Err(), sql.ErrNoRows) {
		return errors.Errorf("invalid transaction - %+v", tr)
	}

	return row.Err()
}

func (p pgRepo) setStatus(ctx context.Context, tx *sqlx.Tx, login string, status entity.Status) error {
	res, err := tx.ExecContext(ctx, queryUpdateAgent,
		login,
		status,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error to get rows after update agent: %w, %s", err, login)
	}
	if rows <= 0 {
		return fmt.Errorf("rows affected %v <= 0, after update user: %s", rows, login)
	}

	return nil
}

func (p pgRepo) transactionLog(ctx context.Context, tx *sqlx.Tx, log entity.Logs) (*int64, error) {
	var logID int64

	err := tx.GetContext(ctx, &logID, queryLog,
		log.Agent,
		log.Src,
		log.Dst,
		log.Mode,
	)
	if err != nil {
		return nil, err
	}

	return &logID, nil
}

func hash64(s string) (*int64, error) {
	h := fnv.New64()

	_, err := h.Write([]byte(s))
	if err != nil {
		return nil, err
	}

	v := int64(h.Sum64())

	return &v, nil
}
