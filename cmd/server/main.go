package main

import (
	"fmt"
	"os"

	"github.com/izruff/reviu-backend/internal/database"
	"github.com/izruff/reviu-backend/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	dsn := os.Getenv("DATABASE_URI")
	db, err := database.OpenDB(dsn)
	if err != nil {
		panic(err)
	}

	r := router.Setup(db)
	fmt.Print("Listening on port 8080 at http://localhost:8080!")

	r.Run()
}
