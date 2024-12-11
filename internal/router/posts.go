package router

import (
	"database/sql"
	"net/http"

	"github.com/BorzooMV/gophers-journal/internal/handlers"
)

type Router struct {
	DB *sql.DB
}

func (ro Router) PostsRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.GetAllPosts(w, r, ro.DB)
	case http.MethodPost:
		handlers.CreateNewPost(w, r, ro.DB)
	default:
		http.Error(w, "wrong method provided", http.StatusBadRequest)
	}
}
