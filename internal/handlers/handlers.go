package handlers

import (
	"github.com/izruff/reviu-backend/internal/services"
	"github.com/jmoiron/sqlx"
)

type APIHandlers struct {
	services *services.PostgresServices
}

func NewAPIHandlers(db *sqlx.DB) *APIHandlers {
	services := services.NewPostgresServices(db)

	return &APIHandlers{
		services: services,
	}
}
