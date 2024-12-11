package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/BorzooMV/gophers-journal/internal/router"
	"github.com/BorzooMV/gophers-journal/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the Database
	db := services.ConnectDb()
	defer db.Close()

	// Define routes
	appRouter := router.Router{DB: db}
	http.Handle("/api/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/posts") {
			appRouter.PostsRouter(w, r)
			return
		}
		http.NotFound(w, r)
	}))

	// Start the server
	fmt.Println("Listening on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
