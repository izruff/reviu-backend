package api

import "github.com/izruff/reviu-backend/internal/services"

type APIServer struct {
	listenAddr string
	services   *services.PostgresServices
}

func NewAPIServer(listenAddr string, services *services.PostgresServices) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		services:   services,
	}
}
