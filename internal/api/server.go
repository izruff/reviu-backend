package api

import (
	"fmt"

	"github.com/izruff/reviu-backend/internal/handlers"
	"github.com/jmoiron/sqlx"
)

type APIServer struct {
	listenAddr string
	origin     string
	handlers   *handlers.APIHandlers
}

func NewAPIServer(listenAddr string, db *sqlx.DB, origin string) *APIServer {
	handlers := handlers.NewAPIHandlers(db, origin)

	return &APIServer{
		listenAddr: listenAddr,
		origin:     origin,
		handlers:   handlers,
	}
}

func (s *APIServer) Run() {
	r := SetupRouter(s)
	r.Run()
	fmt.Println("Listening on address", s.listenAddr)
}
