package router

import (
	"net/http"

	"github.com/BorzooMV/gophers-journal/internal/handlers"
)

func PostsRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.GetAllPosts(w, r)
	case http.MethodPost:
		handlers.CreateNewPost(w, r)
	default:
		http.Error(w, "wrong method provided", http.StatusBadRequest)
	}
}
