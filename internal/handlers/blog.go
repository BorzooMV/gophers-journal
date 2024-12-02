package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/BorzooMV/gophers-journal/internal/model"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	postsFile, err := os.ReadFile("assets/data/sample-posts.json")
	if err != nil {
		http.Error(w, "Couldn't read the posts file!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(postsFile)

}

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var newPost model.Post
	bodyDecoder := json.NewDecoder(r.Body)
	err := bodyDecoder.Decode(&newPost)
	if err != nil {
		http.Error(w, "Couldn't read body", http.StatusInternalServerError)
	}
}
