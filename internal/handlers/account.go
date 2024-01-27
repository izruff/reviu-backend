package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/utils"
)

func (h *APIHandlers) Login(c *gin.Context) {
	var json loginJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, token, err := h.services.Login(json.UsernameOrEmail, json.Password)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", token, utils.CookieExpiryMinutes*60, "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"userId": userID,
	})
}

func (h *APIHandlers) Signup(c *gin.Context) {
	var json signupJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, token, err := h.services.Signup(json.Email, json.Username, json.Password)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", token, utils.CookieExpiryMinutes*60, "/", "", true, true)

	c.JSON(http.StatusCreated, gin.H{
		"userId": userID,
	})
}

// TODO: recoverAccount, changePassword, changeEmail, and maybe refactoring
