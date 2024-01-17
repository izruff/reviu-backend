package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *APIHandlers) SearchUsers(c *gin.Context) {

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
		"userId":    userID,
		"username":  user.Username.String,
		"nickname":  user.Nickname.String,
		"about":     user.About.String,
		"createdAt": user.CreatedAt.Time,
	})
}

func (s *APIHandlers) GetUserFollowers(c *gin.Context) {

}

func (s *APIHandlers) GetUserFollowings(c *gin.Context) {

}

func (s *APIHandlers) SearchPosts(c *gin.Context) {

}

func (s *APIHandlers) GetPost(c *gin.Context) {

}

func (s *APIHandlers) SearchCommentsInPost(c *gin.Context) {

}

func (s *APIHandlers) GetComment(c *gin.Context) {

}

func (s *APIHandlers) SearchTopics(c *gin.Context) {

}

func (s *APIHandlers) GetTopic(c *gin.Context) {

}

func (s *APIHandlers) SearchTags(c *gin.Context) {

}
