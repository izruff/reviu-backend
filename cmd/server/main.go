package main

import (
	"fmt"

	"github.com/izruff/reviu-backend/internal/router"
)

func main() {
	r := router.Setup()
	fmt.Print("Listening on port 8080 at http://localhost:8080!")

	r.Run()
}
