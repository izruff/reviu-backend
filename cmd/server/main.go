package main

import (
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/izruff/reviu-backend/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	dsn := os.Getenv("DATABASE_URI")
	db, err := OpenPostgresDB(dsn)
	if err != nil {
		panic(err)
	}

	listenAddr := os.Getenv("LISTEN_ADDR")
	origin := os.Getenv("ORIGIN")

	// TODO: configure listening address and other stuff
	r := gin.Default()
	SetupRoutes(r, httpHandler)

	r.Run()
}

func OpenPostgresDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./internal/migrations/postgres",
		"postgres", driver)
	if err != nil {
		return nil, err
	}
	m.Up()

	return db, nil
}

func SetupRoutes(r *gin.Engine) {
	r.Use(s.handlers.CORSMiddleware)

	// Simple route for testing ping.
	r.GET("/ping", s.handlers.Ping)

	// These routes are classified into four categories:
	//   account: requests related to user authentication,
	//   public: requests which can be made by any visitor,
	//   authorized: requests only available for registered users, and
	//   moderator: requests for moderation purposes.

	account := r.Group("/account")
	{
		account.POST("/login", s.handlers.Login)
		account.POST("/signup", s.handlers.Signup)
		account.GET("/check-token", s.handlers.CheckToken)

		// For security purposes, the actions below require a full login
		// (email/username and password input), not a JWT token.
		// TODO
	}

	public := r.Group("/public")
	{
		users := public.Group("/users")
		users.GET("", s.handlers.SearchUsers)
		user := users.Group("/id/:userID")
		{
			user.GET("", s.handlers.GetUserProfile)
			user.GET("/relations", s.handlers.GetUserRelations)
			// user.GET("/activity", s.handlers.GetUserActivity)
		}
		userByName := users.Group("/name/:username")
		{
			userByName.GET("", s.handlers.GetUserProfileByUsername)
			userByName.GET("/relations", s.handlers.GetUserRelationsByUsername)
			// userByName.GET("/activity", s.handlers.GetUserActivityByUsername)
		}

		posts := public.Group("/posts")
		posts.GET("/search", s.handlers.SearchPosts)
		posts.GET("/id/:postID", s.handlers.GetPost)
		posts.GET("/id/:postID/replies", s.handlers.GetRepliesToPost)

		comments := public.Group("/comments")
		comments.GET("/search", s.handlers.SearchComments)
		comments.GET("/id/:commentID", s.handlers.GetComment)
		comments.GET("/id/:commentID/replies", s.handlers.GetRepliesToComment)

		topics := public.Group("/topics")
		topics.GET("", s.handlers.SearchTopics)
		topic := topics.Group("/id/:topicID")
		{
			topic.GET("", s.handlers.GetTopic)
		}
		// topics.GET("/id/:topicID", s.handlers.GetTopic)

		tags := public.Group("/tags")
		tags.GET("", s.handlers.SearchTags)
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
			post.GET("", s.handlers.GetPostInteractions)
			post.POST("/view", s.handlers.ViewPost)
			post.PATCH("/edit", s.handlers.EditPost)
			post.POST("/reply", s.handlers.ReplyToPost)
			post.POST("/vote", s.handlers.VotePost)
			post.POST("/bookmark", s.handlers.BookmarkPost)
		}

		comments := authorized.Group("/comments")
		comment := comments.Group("/id/:commentID")
		{
			comment.POST("/reply", s.handlers.ReplyToComment)
			comment.POST("/vote", s.handlers.VoteComment)
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
