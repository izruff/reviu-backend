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
	auth.POST("/login", s.handlers.Login)
	auth.POST("/signup", s.handlers.Signup)

	// Routes for interacting with the user model
	users := r.Group("/users")
	user := users.Group("/id/:userID")
	{
		user.GET("/", s.handlers.GetUserProfileByID)
		user.PATCH("/", s.handlers.UpdateUserProfileByID)
	}
	users.GET("/search", s.handlers.SearchUsername)
}
