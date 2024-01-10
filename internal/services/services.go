package services

import (
	"github.com/izruff/reviu-backend/internal/repository"
	"github.com/jmoiron/sqlx"
)

type APIServices struct {
	queries *repository.PostgresQueries
}

func NewAPIServices(db *sqlx.DB) *APIServices {
	queries := repository.NewPostgresQueries(db)

	return &APIServices{
		queries: queries,
	}
}
