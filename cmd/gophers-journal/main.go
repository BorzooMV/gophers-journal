package main

import (
	"fmt"
	"net/http"

	"github.com/BorzooMV/gophers-journal/internal/handlers"
)

func main() {
	// Define routes
	http.HandleFunc("/api/posts", handlers.GetAllPosts)

	// Start the server
	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
