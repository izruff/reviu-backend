package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/core/domain"
	"gopkg.in/guregu/null.v3"
)

func (h *HTTPHandler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (h *HTTPHandler) SearchUsers(c *gin.Context) {
	var options domain.SearchUsersOptions
	if err := c.ShouldBindQuery(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	users, err := h.svc.SearchUsers(&options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *HTTPHandler) GetUserProfile(c *gin.Context) {
	userID, parseErr := strconv.ParseInt(c.Param("userID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	user, err := h.svc.GetUserByID(userID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *HTTPHandler) GetUserRelations(c *gin.Context) {
	userID, parseErr := strconv.ParseInt(c.Param("userID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	followers, err := h.svc.GetUserFollowers(userID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	followings, err := h.svc.GetUserFollowings(userID)
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

func (h *HTTPHandler) GetUserProfileByUsername(c *gin.Context) {
	// TODO: refactor this and GetUserProfile
	username := c.Param("username")
	userID, err := h.svc.GetUserIDByUsername(username)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	user, err := h.svc.GetUserByID(userID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *HTTPHandler) GetUserRelationsByUsername(c *gin.Context) {
	// TODO: refactor this and GetUserRelations
	username := c.Param("username")
	userID, err := h.svc.GetUserIDByUsername(username)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	followers, err := h.svc.GetUserFollowers(userID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	followings, err := h.svc.GetUserFollowings(userID)
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

func (h *HTTPHandler) SearchPosts(c *gin.Context) {
	var options domain.SearchPostsOptions
	if err := c.ShouldBindQuery(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	posts, err := h.svc.SearchPosts(&options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *HTTPHandler) GetPost(c *gin.Context) {
	postID, parseErr := strconv.ParseInt(c.Param("postID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	post, err := h.svc.GetPostByID(postID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *HTTPHandler) GetRepliesToPost(c *gin.Context) {
	postID, parseErr := strconv.ParseInt(c.Param("postID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	options := &domain.SearchCommentsOptions{
		PostID: null.NewInt(postID, true),
	}
	comments, err := h.svc.SearchComments(options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *HTTPHandler) SearchComments(c *gin.Context) {
	var options domain.SearchCommentsOptions
	if err := c.ShouldBindQuery(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	comments, err := h.svc.SearchComments(&options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *HTTPHandler) GetComment(c *gin.Context) {
	commentID, parseErr := strconv.ParseInt(c.Param("commentID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	comment, err := h.svc.GetCommentByID(commentID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h *HTTPHandler) GetRepliesToComment(c *gin.Context) {
	commentID, parseErr := strconv.ParseInt(c.Param("commentID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	options := &domain.SearchCommentsOptions{
		ParentCommentID: null.NewInt(commentID, true),
	}
	comments, err := h.svc.SearchComments(options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *HTTPHandler) SearchTopics(c *gin.Context) {
	var options domain.SearchTopicsOptions
	if err := c.ShouldBindQuery(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	topics, err := h.svc.SearchTopics(&options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, topics)
}

func (h *HTTPHandler) GetTopic(c *gin.Context) {
	topicID, parseErr := strconv.ParseInt(c.Param("topicID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	topic, err := h.svc.GetTopicByID(topicID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, topic)
}

func (h *HTTPHandler) SearchTags(c *gin.Context) {
	var options domain.SearchTagsOptions
	if err := c.ShouldBindQuery(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tags, err := h.svc.SearchTags(&options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, tags)
}
