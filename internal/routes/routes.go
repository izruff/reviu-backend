package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/api"
)

func SetupRoutes(r *gin.Engine, s *api.APIServer) {
	// For testing
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Routes for user authentication
	auth := r.Group("/auth")
	auth.POST("/login", s.LogIn)
	auth.POST("/signup", s.SignUp)

	// Routes for interacting with the user model
	users := r.Group("/users")
	user := users.Group("/id/:userID")
	user.GET("/", s.GetUserInfoByID)
	user.PATCH("/", s.UpdateUserInfo)
}
