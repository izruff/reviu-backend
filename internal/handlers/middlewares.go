package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/utils"
)

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
