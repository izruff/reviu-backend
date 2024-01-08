package main

import (
	"fmt"
	"os"

	"github.com/izruff/reviu-backend/internal/api"
	"github.com/izruff/reviu-backend/internal/database"
	"github.com/izruff/reviu-backend/internal/router"
	"github.com/izruff/reviu-backend/internal/services"
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
	services := services.NewPostgresServices(db)
	server := api.NewAPIServer(listenAddr, services)

	r := router.SetupRouter(server)
	r.Run()
	fmt.Println("Listening on address", listenAddr)
}
