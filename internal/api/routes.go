package api

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine, s *APIServer) {
	// For testing
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Routes for user authentication
	auth := r.Group("/auth")
	auth.POST("/login", s.handlers.LogIn)
	auth.POST("/signup", s.handlers.SignUp)

	// Routes for interacting with the user model
	users := r.Group("/users")
	user := users.Group("/id/:userID")
	user.GET("/", s.handlers.GetUserInfoByID)
	user.PATCH("/", s.handlers.UpdateUserInfo)
}