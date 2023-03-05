package repository

import (
	"context"
	"fmt"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	queryCreateTables = `	
		CREATE TYPE modes AS ENUM (
			'MAN',
			'AUT'
		);
		CREATE TABLE IF NOT EXISTS statuses (
			id INT PRIMARY KEY,
			name VARCHAR(64) NOT NULL UNIQUE
		);
		CREATE TABLE IF NOT EXISTS agents (
			login VARCHAR(64) PRIMARY KEY,
			password VARCHAR(128) NOT NULL,
			status_id INT REFERENCES statuses(id),
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
		);
		CREATE TABLE IF NOT EXISTS transitions (
			id SERIAL NOT NULL,
			status_id INT REFERENCES statuses(id),
			mode modes NOT NULL,
			permitted_ids int[] NOT NULL,
		  CONSTRAINT transitions_pk
        PRIMARY KEY (status_id, mode)
		);
		CREATE TABLE IF NOT EXISTS transitions_log (
			id SERIAL PRIMARY KEY,
			agent_login VARCHAR(64) REFERENCES statuses(id),
			old_status_id INT REFERENCES statuses(id),
			new_status_id INT REFERENCES statuses(id),
			mode modes NOT NULL,
			processed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
		);
	`

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
		user.Password,
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
	).Scan(&user.Login, &user.Password, &user.StatusID, &user.CreatedAt)
	if err != nil {
		return entity.Agent{}, fmt.Errorf("error to get user: %w, %s", err, login)
	}

	return user, nil
}
