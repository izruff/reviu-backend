package api

import (
	"fmt"

	"github.com/izruff/reviu-backend/internal/handlers"
	"github.com/jmoiron/sqlx"
)

type APIServer struct {
	listenAddr string
	handlers   *handlers.APIHandlers
}

func NewAPIServer(listenAddr string, db *sqlx.DB) *APIServer {
	handlers := handlers.NewAPIHandlers(db)

	return &APIServer{
		listenAddr: listenAddr,
		handlers:   handlers,
	}
}

func (s *APIServer) Run() {
	r := SetupRouter(s)
	r.Run()
	fmt.Println("Listening on address", s.listenAddr)
}
