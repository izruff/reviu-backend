package api

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine, s *APIServer) {
	// These routes are classified into four categories:
	//   account: requests related to user authentication,
	//   public: requests which can be made by any visitor,
	//   authorized: requests only available for registered users, and
	//   moderator: requests for moderation purposes.

	account := r.Group("/account")
	{
		account.POST("/login", s.handlers.Login)
		account.POST("/signup", s.handlers.Signup)

		// For security purposes, the actions below require a full login
		// (email/username and password input), not a JWT token.
		// TODO
	}

	public := r.Group("/public")
	{
		users := public.Group("/users")
		users.GET("/", s.handlers.SearchUsers)
		user := users.Group("/id/:userID")
		{
			user.GET("/", s.handlers.GetUserProfile)
			user.GET("/followers", s.handlers.GetUserFollowers)
			user.GET("/followings", s.handlers.GetUserFollowings)
		}

		posts := public.Group("/posts")
		posts.GET("/", s.handlers.SearchPosts)
		posts.GET("/id/:postID", s.handlers.GetPost)

		comments := public.Group("/posts/id/:postID/comments")
		comments.GET("/", s.handlers.SearchCommentsInPost)
		comments.GET("/id/:commentID", s.handlers.GetComment)

		topics := public.Group("/topics")
		topics.GET("/", s.handlers.SearchTopics)
		topics.GET("/id/:topicID", s.handlers.GetTopic)

		tags := public.Group("/tags")
		tags.GET("/", s.handlers.SearchTags)
	}

	authorized := r.Group("/authorized", s.handlers.JWTAuth)
	{
		users := authorized.Group("/users")
		users.PATCH("/me", s.handlers.UpdateUserProfile)
		users.GET("/me/private", s.handlers.GetUserPrivates)
		users.GET("/me/subscriptions", s.handlers.GetUserSubscriptions)
		users.GET("/me/bookmarks", s.handlers.GetUserBookmarks)
		users.POST("/follow", s.handlers.FollowUser)
		users.DELETE("/unfollow", s.handlers.UnfollowUser)

		posts := authorized.Group("/posts")
		posts.POST("/create", s.handlers.CreatePost)
		post := posts.Group("/id/:postID")
		{
			post.PATCH("/edit", s.handlers.EditPost)
			post.POST("/comment", s.handlers.CreateCommentOnPost)
			post.POST("/vote", s.handlers.VotePost)
			post.POST("/bookmark", s.handlers.BookmarkPost)

			comment := post.Group("/comments/id/:commentID")
			comment.POST("/reply", s.handlers.ReplyToComment)
			comment.PATCH("/edit", s.handlers.EditComment)
		}

		topics := authorized.Group("/topics")
		topics.POST("/create", s.handlers.CreateTopic)
	}

	moderator := r.Group("/moderator")
	{
		moderator.POST("/users/ban", s.handlers.BanUser)
		moderator.PATCH("/posts/delete", s.handlers.MarkPostAsDeleted)
		moderator.PATCH("/comments/delete", s.handlers.MarkCommentAsDeleted)
	}
}
