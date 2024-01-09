package services

import "github.com/jmoiron/sqlx"

type PostgresServices struct {
	db *sqlx.DB
}

func NewPostgresServices(db *sqlx.DB) *PostgresServices {
	return &PostgresServices{
		db: db,
	}
}
