package http

import (
	"github.com/izruff/reviu-backend/internal/services"
	"github.com/jmoiron/sqlx"
)

type APIHandlers struct {
	services *services.APIServices
	origin   string
}

func NewAPIHandlers(db *sqlx.DB, origin string) *APIHandlers {
	services := services.NewAPIServices(db)

	return &APIHandlers{
		services: services,
		origin:   origin,
	}
}
