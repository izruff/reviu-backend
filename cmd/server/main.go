package main

import (
	"os"

	"github.com/izruff/reviu-backend/internal/api"
	"github.com/izruff/reviu-backend/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	dsn := os.Getenv("DATABASE_URI")
	db, err := database.OpenPostgresDB(dsn)
	if err != nil {
		panic(err)
	}

	listenAddr := os.Getenv("LISTEN_ADDR")
	server := api.NewAPIServer(listenAddr, db)

	server.Run()
}
