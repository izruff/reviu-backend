package repository

import "github.com/jmoiron/sqlx"

type PostgresQueries struct {
	db *sqlx.DB
}

func NewPostgresQueries(db *sqlx.DB) *PostgresQueries {
	return &PostgresQueries{
		db: db,
	}
}
