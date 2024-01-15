package api

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine, s *APIServer) {
	// For testing
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Routes for user authentication
	auth := r.Group("/auth")
	auth.POST("/login", s.handlers.Login)
	auth.POST("/signup", s.handlers.Signup)

	// Routes for interacting with the user model
	users := r.Group("/users")
	user := users.Group("/id/:userID")
	{
		user.GET("/", s.handlers.GetUserProfile)
		withAuth := user.Group("/", s.handlers.JWTAuth)
		{
			withAuth.PATCH("/", s.handlers.UpdateUserProfile)
			withAuth.DELETE("/", s.handlers.DeleteUser)

			withAuth.GET("/post-list", s.handlers.GetUserPosts)
			withAuth.POST("/follow", s.handlers.FollowUser)
			withAuth.GET("/follow-list", s.handlers.GetFollowers)
			withAuth.GET("/bookmark-list", s.handlers.GetUserBookmarkedPosts)
		}
		modOnly := user.Group("/", s.handlers.JWTAuth) // TODO: replace with middleware that also checks mod role
		{
			modOnly.POST("/ban", s.handlers.BanUser)
		}
	}
	users.GET("/search", s.handlers.SearchUsernames)

	// Routes for interacting with the post and comment model
	posts := r.Group("/posts")
	post := posts.Group("/id/:postID")
	{
		post.GET("/", s.handlers.GetPostOnly)
		post.GET("/all", s.handlers.GetPostWithChildren)
		withAuth := post.Group("/", s.handlers.JWTAuth)
		{
			withAuth.POST("/comment", s.handlers.CommentOnPost)
			withAuth.PATCH("/edit", s.handlers.EditPost)
			withAuth.POST("/vote", s.handlers.VotePost)
			withAuth.POST("/bookmark", s.handlers.BookmarkPost)
		}
		modOnly := post.Group("/", s.handlers.JWTAuth) // TODO: replace with middleware that also checks mod role
		{
			modOnly.PATCH("/delete", s.handlers.DeletePost)
		}

		comment := post.Group("/comment/:commentID")
		{
			comment.GET("/", s.handlers.GetCommentOnly)
			comment.GET("/all", s.handlers.GetCommentWithChildren)
			withAuth := comment.Group("/", s.handlers.JWTAuth)
			{
				withAuth.POST("/reply", s.handlers.ReplyToComment)
				withAuth.PATCH("/edit", s.handlers.EditComment)
			}
			modOnly := comment.Group("/", s.handlers.JWTAuth) // TODO: replace with middleware that also checks mod role
			{
				modOnly.PATCH("/delete", s.handlers.DeleteComment)
			}
		}
	}
	posts.POST("/create", s.handlers.JWTAuth, s.handlers.CreatePost)
	posts.GET("/search", s.handlers.SearchPosts)

	// Routes for interacting with the topic model
	topics := r.Group("/topics")
	topic := topics.Group("/id/:topicID")
	{
		topic.GET("/", s.handlers.GetTopicContents)
	}
	topics.POST("/create", s.handlers.JWTAuth, s.handlers.CreateTopic)
	topics.GET("/search", s.handlers.SearchTopics)
}
