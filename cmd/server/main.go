package main

import (
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"

	handler "github.com/izruff/reviu-backend/internal/adapters/handler/http"
	repository "github.com/izruff/reviu-backend/internal/adapters/repository/postgres"
	service "github.com/izruff/reviu-backend/internal/core/services"
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

	postgresRepo := repository.NewPostgresRepository(db)
	services := service.NewAPIServices(postgresRepo)
	httpHandler := handler.NewHTTPHandler(services, origin)

	// TODO: configure listening address and other stuff
	r := gin.Default()
	SetupRoutes(r, httpHandler)

	r.Run(listenAddr)
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

func SetupRoutes(r *gin.Engine, h *handler.HTTPHandler) {
	r.Use(h.CORSMiddleware)

	// Simple route for testing ping.
	r.GET("/ping", h.Ping)

	// These routes are classified into four categories:
	//   account: requests related to user authentication,
	//   public: requests which can be made by any visitor,
	//   authorized: requests only available for registered users, and
	//   moderator: requests for moderation purposes.

	account := r.Group("/account")
	{
		account.POST("/login", h.Login)
		account.POST("/signup", h.Signup)
		account.GET("/check-token", h.CheckToken)

		// For security purposes, the actions below require a full login
		// (email/username and password input), not a JWT token.
		// TODO
	}

	public := r.Group("/public")
	{
		users := public.Group("/users")
		users.GET("", h.SearchUsers)
		user := users.Group("/id/:userID")
		{
			user.GET("", h.GetUserProfile)
			user.GET("/relations", h.GetUserRelations)
			// user.GET("/activity", h.GetUserActivity)
		}
		userByName := users.Group("/name/:username")
		{
			userByName.GET("", h.GetUserProfileByUsername)
			userByName.GET("/relations", h.GetUserRelationsByUsername)
			// userByName.GET("/activity", h.GetUserActivityByUsername)
		}

		posts := public.Group("/posts")
		posts.GET("/search", h.SearchPosts)
		posts.GET("/id/:postID", h.GetPost)
		posts.GET("/id/:postID/replies", h.GetRepliesToPost)

		comments := public.Group("/comments")
		comments.GET("/search", h.SearchComments)
		comments.GET("/id/:commentID", h.GetComment)
		comments.GET("/id/:commentID/replies", h.GetRepliesToComment)

		topics := public.Group("/topics")
		topics.GET("", h.SearchTopics)
		topic := topics.Group("/id/:topicID")
		{
			topic.GET("", h.GetTopic)
		}
		// topics.GET("/id/:topicID", h.GetTopic)

		tags := public.Group("/tags")
		tags.GET("", h.SearchTags)
	}

	authorized := r.Group("/authorized", h.JWTAuth)
	{
		users := authorized.Group("/users")
		users.PATCH("/me", h.UpdateUserProfile)
		users.GET("/me/private", h.GetUserPrivates)
		users.GET("/me/subscriptions", h.GetUserSubscriptions)
		users.GET("/me/bookmarks", h.GetUserBookmarks)
		users.POST("/follow", h.FollowUser)
		users.DELETE("/unfollow", h.UnfollowUser)

		posts := authorized.Group("/posts")
		posts.POST("/create", h.CreatePost)
		post := posts.Group("/id/:postID")
		{
			post.GET("", h.GetPostInteractions)
			post.POST("/view", h.ViewPost)
			post.PATCH("/edit", h.EditPost)
			post.POST("/reply", h.ReplyToPost)
			post.POST("/vote", h.VotePost)
			post.POST("/bookmark", h.BookmarkPost)
		}

		comments := authorized.Group("/comments")
		comment := comments.Group("/id/:commentID")
		{
			comment.POST("/reply", h.ReplyToComment)
			comment.POST("/vote", h.VoteComment)
			comment.PATCH("/edit", h.EditComment)
		}

		topics := authorized.Group("/topics")
		topics.POST("/create", h.CreateTopic)
	}

	moderator := r.Group("/moderator")
	{
		moderator.POST("/users/ban", h.BanUser)
		moderator.PATCH("/posts/delete", h.MarkPostAsDeleted)
		moderator.PATCH("/comments/delete", h.MarkCommentAsDeleted)
	}
}
