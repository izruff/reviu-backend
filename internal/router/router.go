package router

import (
	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/routes"
	"github.com/jmoiron/sqlx"
)

func Setup(db *sqlx.DB) *gin.Engine {
	r := gin.Default()
	routes.SetupRoutes(r, db)
	return r
}
