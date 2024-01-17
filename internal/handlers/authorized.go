package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/models"
	"gopkg.in/guregu/null.v3"
)

type updateUserProfileJSON struct {
	Nickname null.String `json:"nickname" binding:"required"`
	About    null.String `json:"about" binding:"required"`
}

func (s *APIHandlers) UpdateUserProfile(c *gin.Context) {
	userID, parseErr := strconv.ParseInt(c.Param("userID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	var json updateUserProfileJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedUser := &models.User{
		Nickname: json.Nickname,
		About:    json.About,
	}

	err := s.services.UpdateUserByID(userID, updatedUser)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (s *APIHandlers) GetUserPrivates(c *gin.Context) {

}

func (s *APIHandlers) GetUserSubscriptions(c *gin.Context) {

}

func (s *APIHandlers) GetUserBookmarks(c *gin.Context) {

}

func (s *APIHandlers) FollowUser(c *gin.Context) {

}

func (s *APIHandlers) UnfollowUser(c *gin.Context) {

}

func (s *APIHandlers) CreatePost(c *gin.Context) {

}

func (s *APIHandlers) GetFollowers(c *gin.Context) {

}

func (s *APIHandlers) EditPost(c *gin.Context) {

}

func (s *APIHandlers) CreateCommentOnPost(c *gin.Context) {

}

func (s *APIHandlers) VotePost(c *gin.Context) {

}

func (s *APIHandlers) BookmarkPost(c *gin.Context) {

}

func (s *APIHandlers) ReplyToComment(c *gin.Context) {

}

func (s *APIHandlers) EditComment(c *gin.Context) {

}

func (s *APIHandlers) CreateTopic(c *gin.Context) {

}
