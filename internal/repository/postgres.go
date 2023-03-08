package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	querySaveUser = `INSERT INTO agents (login, password, status_id) VALUES ($1, $2, $3)
		ON CONFLICT (login) DO NOTHING`
	queryGetUser = `SELECT login, password, status_id, created_at FROM agents WHERE login = $1`
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

func (p pgRepo) AgentSetStatusTx(ctx context.Context, agent entity.Agent, mode entity.Mode) error {
	tx, err := p.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return errors.WithStack(err)
	}

	transaction := entity.Transition{
		Dst:  *agent.Status,
		Mode: mode,
	}

	callbacks := []func(ctx context.Context, tx *sqlx.Tx) error{
		func(ctx context.Context, tx *sqlx.Tx) error {
			a, err := p.getAgent(ctx, tx, agent.Login)
			if err != nil {
				return err
			}

			transaction.Src = *a.Status

			return err
		},
		func(ctx context.Context, tx *sqlx.Tx) error {
			v, err := p.isValidTransaction(ctx, tx, transaction)
			if err != nil {
				return err
			}
			if !v {
				return errors.Errorf("invalid transaction - %+v", transaction)
			}

			return nil
		},
		func(ctx context.Context, tx *sqlx.Tx) error {
			return p.setStatus(ctx, tx, agent.Login, *agent.Status)
		},
		func(ctx context.Context, tx *sqlx.Tx) error {
			return p.transactionLog(ctx, tx, entity.Logs{
				Agent:      agent.Login,
				Transition: transaction,
			})
		},
	}

	for _, cb := range callbacks {
		err = cb(ctx, tx)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Println("error rolling back a transaction:", rollbackErr)
			}

			return errors.WithStack(err)
		}
	}

	return errors.WithStack(tx.Commit())
}

func (p pgRepo) getAgent(ctx context.Context, tx *sqlx.Tx, login string) (entity.Agent, error) {
	// todo
	return entity.Agent{}, nil
}

func (p pgRepo) isValidTransaction(ctx context.Context, tx *sqlx.Tx, transition entity.Transition) (bool, error) {
	// todo
	return false, nil
}

func (p pgRepo) setStatus(ctx context.Context, tx *sqlx.Tx, login string, status entity.Status) error {
	// todo
	return nil
}

func (p pgRepo) transactionLog(ctx context.Context, tx *sqlx.Tx, logs entity.Logs) error {
	// todo
	return nil
}
