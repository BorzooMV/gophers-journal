package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/BorzooMV/gophers-journal/internal/model"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	postsFile, err := os.ReadFile("assets/data/sample-posts.json")
	if err != nil {
		http.Error(w, "couldn't read the posts file!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(postsFile)

}

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var newPost model.Post

	// Decode the incoming request body
	bodyDecoder := json.NewDecoder(r.Body)
	err := bodyDecoder.Decode(&newPost)
	if err != nil {
		http.Error(w, "couldn't read body", http.StatusInternalServerError)
		return
	}

	// Validate the new post
	err = newPost.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add a timestamp to the new post
	newPost.CreatedAt = time.Now()

	// Read existing posts from the JSON file
	jsonFile, err := os.Open("assets/data/sample-posts.json")
	if err != nil {
		http.Error(w, "couldn't open the file", http.StatusInternalServerError)
		return
	}

	var prevPosts []model.Post
	err = json.NewDecoder(jsonFile).Decode(&prevPosts)
	if err != nil {
		http.Error(w, "Couldn't decode JSON file", http.StatusInternalServerError)
		return
	}

	jsonFile.Close()

	// Append the new post to the previous posts
	prevPosts = append(prevPosts, newPost)

	// Reopen the file in write mode to overwrite the content
	jsonFile, err = os.OpenFile("assets/data/sample-posts.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		http.Error(w, "Couldn't open the file for writing", http.StatusInternalServerError)
		return
	}
	defer jsonFile.Close()

	// Write the updated data back to the file
	err = json.NewEncoder(jsonFile).Encode(prevPosts)
	if err != nil {
		http.Error(w, "Couldn't encode JSON file", http.StatusInternalServerError)
		return
	}

	// Respond with the created post
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newPost)
	if err != nil {
		http.Error(w, "Couldn't encode response.", http.StatusInternalServerError)
	}
}
