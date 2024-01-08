package router

import (
	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/api"
	"github.com/izruff/reviu-backend/internal/routes"
)

func SetupRouter(s *api.APIServer) *gin.Engine {
	// TODO: configure listening address and other stuff
	r := gin.Default()
	routes.SetupRoutes(r, s)
	return r
}
