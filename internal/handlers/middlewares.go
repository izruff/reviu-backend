package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/utils"
)

func (s *APIHandlers) JWTAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader[:6] != "Bearer" {
		c.AbortWithStatus(401)
	}

	tokenString := authHeader[7:]
	if err := utils.IsValidJWT(tokenString); err != nil {
		c.AbortWithStatus(401)
	}
}
