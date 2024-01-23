package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/utils"
)

// https://stackoverflow.com/questions/29418478/go-gin-framework-cors
func (s *APIHandlers) CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func (s *APIHandlers) JWTAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader[:6] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "no valid authorization header",
		})
		return
	}

	tokenString := authHeader[7:]
	userID, err := utils.IsValidJWT(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Set("userID", userID)
}
