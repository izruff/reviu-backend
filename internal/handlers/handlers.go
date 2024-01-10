package handlers

import (
	"github.com/izruff/reviu-backend/internal/services"
	"github.com/jmoiron/sqlx"
)

type APIHandlers struct {
	services *services.APIServices
}

func NewAPIHandlers(db *sqlx.DB) *APIHandlers {
	services := services.NewAPIServices(db)

	return &APIHandlers{
		services: services,
	}
}
