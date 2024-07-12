package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/models"
	"gopkg.in/guregu/null.v3"
)

func (s *APIHandlers) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (s *APIHandlers) SearchUsers(c *gin.Context) {
	var options models.SearchUsersOptions
	if err := c.ShouldBindQuery(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	users, err := s.services.SearchUsers(&options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, users)
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

	c.JSON(http.StatusOK, user)
}

func (s *APIHandlers) GetUserRelations(c *gin.Context) {
	userID, parseErr := strconv.ParseInt(c.Param("userID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	followers, err := s.services.GetUserFollowers(userID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	followings, err := s.services.GetUserFollowings(userID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"followers":  followers,
		"followings": followings,
	})
}

func (s *APIHandlers) GetUserProfileByUsername(c *gin.Context) {
	// TODO: refactor this and GetUserProfile
	username := c.Param("username")
	userID, err := s.services.GetUserIDByUsername(username)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
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

	c.JSON(http.StatusOK, user)
}

func (s *APIHandlers) GetUserRelationsByUsername(c *gin.Context) {
	// TODO: refactor this and GetUserRelations
	username := c.Param("username")
	userID, err := s.services.GetUserIDByUsername(username)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	followers, err := s.services.GetUserFollowers(userID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	followings, err := s.services.GetUserFollowings(userID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"followers":  followers,
		"followings": followings,
	})
}

func (s *APIHandlers) SearchPosts(c *gin.Context) {
	var options models.SearchPostsOptions
	if err := c.ShouldBindQuery(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	posts, err := s.services.SearchPosts(&options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (s *APIHandlers) GetPost(c *gin.Context) {
	postID, parseErr := strconv.ParseInt(c.Param("postID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	post, err := s.services.GetPostByID(postID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (s *APIHandlers) GetRepliesToPost(c *gin.Context) {
	postID, parseErr := strconv.ParseInt(c.Param("postID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	options := &models.SearchCommentsOptions{
		PostID: null.NewInt(postID, true),
	}
	comments, err := s.services.SearchComments(options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (s *APIHandlers) SearchComments(c *gin.Context) {
	var options models.SearchCommentsOptions
	if err := c.ShouldBindQuery(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	comments, err := s.services.SearchComments(&options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (s *APIHandlers) GetComment(c *gin.Context) {
	commentID, parseErr := strconv.ParseInt(c.Param("commentID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	comment, err := s.services.GetCommentByID(commentID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (s *APIHandlers) GetRepliesToComment(c *gin.Context) {
	commentID, parseErr := strconv.ParseInt(c.Param("commentID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	options := &models.SearchCommentsOptions{
		ParentCommentID: null.NewInt(commentID, true),
	}
	comments, err := s.services.SearchComments(options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (s *APIHandlers) SearchTopics(c *gin.Context) {
	var options models.SearchTopicsOptions
	if err := c.ShouldBindQuery(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	topics, err := s.services.SearchTopics(&options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, topics)
}

func (s *APIHandlers) GetTopic(c *gin.Context) {
	topicID, parseErr := strconv.ParseInt(c.Param("topicID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	topic, err := s.services.GetTopicByID(topicID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, topic)
}

func (s *APIHandlers) SearchTags(c *gin.Context) {
	var options models.SearchTagsOptions
	if err := c.ShouldBindQuery(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tags, err := s.services.SearchTags(&options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, tags)
}
