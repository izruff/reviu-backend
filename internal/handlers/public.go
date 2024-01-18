package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/models"
)

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

	var response []gin.H
	for _, user := range users {
		response = append(response, gin.H{
			"userId":   user.ID.Int64,
			"username": user.Username.String,
			"nickname": user.Nickname.String,
		})
	}
	c.JSON(http.StatusOK, response)
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
	// TODO: have short or long option
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	// TODO: add more info such as follow count
	c.JSON(http.StatusOK, gin.H{
		"username":  user.Username.String,
		"nickname":  user.Nickname.String,
		"about":     user.About.String,
		"createdAt": user.CreatedAt.Time,
	})
}

func (s *APIHandlers) GetUserFollowers(c *gin.Context) {
	userID, parseErr := strconv.ParseInt(c.Param("userID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	users, err := s.services.GetFollowersList(userID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	var response []gin.H
	for _, user := range users {
		response = append(response, gin.H{
			"userId":   user.ID.Int64,
			"username": user.Username.String,
			"nickname": user.Nickname.String,
		})
	}
	c.JSON(http.StatusOK, response)
}

func (s *APIHandlers) GetUserFollowings(c *gin.Context) {
	userID, parseErr := strconv.ParseInt(c.Param("userID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	users, err := s.services.GetFollowingsList(userID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	var response []gin.H
	for _, user := range users {
		response = append(response, gin.H{
			"userId":   user.ID.Int64,
			"username": user.Username.String,
			"nickname": user.Nickname.String,
		})
	}
	c.JSON(http.StatusOK, response)
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

	var response []gin.H
	for _, post := range posts {
		response = append(response, gin.H{
			"postId":    post.ID.Int64,
			"title":     post.Title.String,
			"content":   post.Content.String,
			"authorId":  post.AuthorID.Int64,
			"topicId":   post.TopicID.Int64,
			"createdAt": post.CreatedAt.Time,
		})
	}
	c.JSON(http.StatusOK, response)
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

	c.JSON(http.StatusOK, gin.H{
		"title":     post.Title.String,
		"content":   post.Content.String,
		"authorId":  post.AuthorID.Int64,
		"topicId":   post.TopicID.Int64,
		"createdAt": post.CreatedAt.Time,
	})
}

func (s *APIHandlers) SearchCommentsInPost(c *gin.Context) {
	postID, parseErr := strconv.ParseInt(c.Param("postID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	var options models.SearchCommentsOptions
	if err := c.ShouldBindQuery(&options); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	options.PostID = postID
	comments, err := s.services.SearchCommentsInPost(&options)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	var response []gin.H
	for _, comment := range comments {
		response = append(response, gin.H{
			"commentId": comment.ID.Int64,
			"content":   comment.Content.String,
			"authorId":  comment.AuthorID.Int64,
			"createdAt": comment.CreatedAt.Time,
		})
	}
	c.JSON(http.StatusOK, response)
}

func (s *APIHandlers) GetComment(c *gin.Context) {
	postID, parseErr := strconv.ParseInt(c.Param("postID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	commentID, parseErr := strconv.ParseInt(c.Param("commentID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	comment, err := s.services.GetCommentByID(commentID, postID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content":   comment.Content.String,
		"authorId":  comment.AuthorID.Int64,
		"createdAt": comment.CreatedAt.Time,
	})
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

	var response []gin.H
	for _, topic := range topics {
		response = append(response, gin.H{
			"topicId": topic.ID.Int64,
			"topic":   topic.Topic.String,
			"hub":     topic.Hub.String,
		})
	}
	c.JSON(http.StatusOK, response)
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

	c.JSON(http.StatusOK, gin.H{
		"topic": topic.Topic.String,
		"hub":   topic.Hub.String,
	})
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

	var response []gin.H
	for _, tag := range tags {
		response = append(response, gin.H{
			"tagId": tag.ID.Int64,
			"tag":   tag.Tag.String,
			"hub":   tag.Hub.String,
		})
	}
	c.JSON(http.StatusOK, response)
}
