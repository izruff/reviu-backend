package http

import (
	"net/http"
	"strconv"

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
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

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

	postID, err := s.services.CreatePost(json.Title, json.Content, userID, json.Topic, json.Hub, json.Tags)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": postID,
	})
}

func (s *APIHandlers) GetPostInteractions(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

	postID, parseErr := strconv.ParseInt(c.Param("postID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	viewed, voted, err := s.services.GetPostInteractionsByUserID(postID, userID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	postInteractions := &postInteractionsResponse{
		Viewed: viewed,
		Voted:  voted,
	}
	c.JSON(http.StatusOK, postInteractions)
}

func (s *APIHandlers) ViewPost(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

	postID, parseErr := strconv.ParseInt(c.Param("postID"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr.Error(),
		})
		return
	}

	if err := s.services.ViewPost(postID, userID); err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message, // TODO: error handling when the post was already viewed
		})
		return
	}

	c.Status(http.StatusCreated)
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

	commentID, err := s.services.CreateComment(json.Content, userID, null.NewInt(json.PostID, true), null.NewInt(0, false))
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": commentID,
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

	commentID, err := s.services.CreateComment(json.Content, userID, null.NewInt(0, false), null.NewInt(json.ParentCommentID, true))
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": commentID,
	})
}

func (s *APIHandlers) VoteComment(c *gin.Context) {
	value, _ := c.Get("userID")
	userID := value.(int64)

	var json voteCommentJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := s.services.VoteComment(json.CommentID, userID, json.Up); err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.Status(http.StatusCreated)
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

	comment, err := s.services.GetCommentByID(json.CommentID)
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

	if err := s.services.UpdateCommentByID(json.CommentID, json.Content); err != nil {
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
		"id": topicID,
	})
}
