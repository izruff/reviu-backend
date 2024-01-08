package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/api"
)

func SetupRoutes(r *gin.Engine, s *api.APIServer) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
