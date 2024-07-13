package http

import "github.com/izruff/reviu-backend/internal/core/ports"

type HTTPHandler struct {
	svc    ports.Service
	origin string
}

func NewHTTPHandler(svc ports.Service, origin string) *HTTPHandler {
	return &HTTPHandler{
		svc:    svc,
		origin: origin,
	}
}
