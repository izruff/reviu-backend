package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/models"
)

type updateUserProfileByIDJSON struct {
	UpdatedUser *models.User `json:"updatedUser" binding:"required"`
}

type searchJSON struct {
}

func (s *APIHandlers) GetUserProfile(c *gin.Context) {
	userID, parseErr := strconv.ParseInt(c.Param("userID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	user, err := s.services.GetUserByID(userID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	// TODO: add more info such as follow count
	c.JSON(http.StatusOK, gin.H{
		"id":        userID,
		"username":  user.Username.String,
		"nickname":  user.Nickname.String,
		"about":     user.About.String,
		"createdAt": user.CreatedAt.Time,
	})
}

func (s *APIHandlers) UpdateUserProfile(c *gin.Context) {
	userID, parseErr := strconv.ParseInt(c.Param("userID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	var json updateUserProfileByIDJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := s.services.UpdateUserByID(userID, json.UpdatedUser)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (s *APIHandlers) DeleteUser(c *gin.Context) {

}

func (s *APIHandlers) GetUserPosts(c *gin.Context) {

}

func (s *APIHandlers) FollowUser(c *gin.Context) {

}

func (s *APIHandlers) GetFollowers(c *gin.Context) {

}

func (s *APIHandlers) GetUserBookmarkedPosts(c *gin.Context) {

}

func (s *APIHandlers) BanUser(c *gin.Context) {

}

func (s *APIHandlers) SearchUsernames(c *gin.Context) {

}
