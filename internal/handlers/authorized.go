package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/models"
	"gopkg.in/guregu/null.v3"
)

func (s *APIHandlers) UpdateUserProfile(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

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

	c.Status(http.StatusOK)
}

func (s *APIHandlers) GetUserPrivates(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

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
		"email": user.Email.String,
	})
}

func (s *APIHandlers) GetUserSubscriptions(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

	subscriptions, err := s.services.GetUserSubscriptions(userID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	// We use the subscription model so it can store user preferences later

	response := []gin.H{}
	for _, subscription := range subscriptions {
		response = append(response, gin.H{
			"topicId": subscription.TopicID.Int64,
		})
	}
	c.JSON(http.StatusOK, response)
}

func (s *APIHandlers) GetUserBookmarks(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

	bookmarks, err := s.services.GetUserBookmarks(userID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	// TODO: we use the bookmark model so it can store user preferences later (?)

	response := []gin.H{}
	for _, bookmark := range bookmarks {
		response = append(response, gin.H{
			"postId": bookmark.PostID.Int64,
		})
	}
	c.JSON(http.StatusOK, response)
}

func (s *APIHandlers) FollowUser(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

	var json followOrUnfollowUserJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := s.services.FollowUserByID(userID, json.FollowingID); err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.Status(http.StatusOK)
}

func (s *APIHandlers) UnfollowUser(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

	var json followOrUnfollowUserJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := s.services.UnfollowUserByID(userID, json.FollowingID); err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.Status(http.StatusOK)
}

func (s *APIHandlers) CreatePost(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

	var json createPostJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postID, err := s.services.CreatePost(json.Title, json.Content, userID, json.TopicID, json.Tags)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"postId": postID,
	})
}

func (s *APIHandlers) EditPost(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

	var json editPostJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	post, err := s.services.GetPostByID(json.PostID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}
	if post.AuthorID.Int64 != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You do not have permission to edit this post.",
		})
		return
	}

	updatedPost := &models.Post{
		Title:   json.Title,
		Content: json.Content,
	}

	if err := s.services.UpdatePostByID(json.PostID, updatedPost); err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.Status(http.StatusOK)
}

func (s *APIHandlers) ReplyToPost(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

	var json replyToPostJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	commentID, err := s.services.CreateComment(json.Content, userID, json.PostID, null.NewInt(0, false))
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"commentId": commentID,
	})
}

func (s *APIHandlers) VotePost(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

	var json votePostJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := s.services.VotePost(json.PostID, userID, json.Up); err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (s *APIHandlers) BookmarkPost(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

	var json bookmarkPostJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := s.services.BookmarkPostWithID(json.PostID, userID); err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (s *APIHandlers) ReplyToComment(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

	var json replyToCommentJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	commentID, err := s.services.CreateComment(json.Content, userID, json.PostID, null.NewInt(json.ParentCommentID, true))
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"commentId": commentID,
	})
}

func (s *APIHandlers) EditComment(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

	var json editCommentJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	comment, err := s.services.GetCommentByID(json.CommentID, json.PostID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}
	if comment.AuthorID.Int64 != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You do not have permission to edit this comment.",
		})
		return
	}

	if err := s.services.UpdateCommentByID(json.CommentID, json.PostID, json.Content); err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.Status(http.StatusOK)
}

func (s *APIHandlers) CreateTopic(c *gin.Context) {
	var json createTopicJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	topicID, err := s.services.CreateTopic(json.Topic, json.Hub)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"topicID": topicID,
	})
}
