package api

import "github.com/gin-gonic/gin"

func SetupRouter(s *APIServer) *gin.Engine {
	// TODO: configure listening address and other stuff
	r := gin.Default()
	SetupRoutes(r, s)
	return r
}
