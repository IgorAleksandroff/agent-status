package repository

import (
	"context"

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
			department_ids int[] NOT NULL,
		  CONSTRAINT transitions_pk
        PRIMARY KEY (status_id, mode)
		);
		CREATE TABLE IF NOT EXISTS transitions_log (
			agent_login VARCHAR(64) REFERENCES statuses(id),
			old_status_id INT REFERENCES statuses(id),
			new_status_id INT REFERENCES statuses(id),
			mode modes NOT NULL,
			processed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
		);
	`
)

type pgRepo struct {
	db *sqlx.DB
}

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

func (p *pgRepo) init(ctx context.Context) error {
	_, err := p.db.ExecContext(ctx, queryCreateTables)
	if err != nil {
		return err
	}
	return nil
}

func (p *pgRepo) Close() error {
	return p.db.Close()
}
