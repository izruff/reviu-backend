package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/izruff/reviu-backend/internal/routes"
)

func Setup(db *sql.DB) *gin.Engine {
	r := gin.Default()
	routes.SetupRoutes(r, db)
	return r
}
