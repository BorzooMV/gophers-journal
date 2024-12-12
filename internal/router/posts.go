package router

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/BorzooMV/gophers-journal/internal/handlers"
)

type Router struct {
	DB *sql.DB
}

func (ro Router) PostsRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if hasValidSuffix, suffix := hasAcceptableSuffix(r.URL.Path, "/api/posts/"); hasValidSuffix {
			handlers.GetPostWithId(w, r, ro.DB, suffix)
		} else {
			handlers.GetAllPosts(w, r, ro.DB)
		}
	case http.MethodPost:
		handlers.CreateNewPost(w, r, ro.DB)
	case http.MethodDelete:
		if hasValidSuffix, suffix := hasAcceptableSuffix(r.URL.Path, "/api/posts/"); hasValidSuffix {
			handlers.DeletePostWithId(w, ro.DB, suffix)
		}
	default:
		http.Error(w, "wrong method provided", http.StatusBadRequest)
	}
}

func hasSuffix(s string, p string) bool {
	return s != strings.TrimPrefix(s, p)
}

func hasAcceptableSuffix(s string, p string) (hasValidSuffix bool, suffix string) {
	if hasSuffix(s, p) {
		suffix := strings.Split(strings.TrimPrefix(s, p), "/")[0]
		hasValidSuffix = suffix != ""
		return hasValidSuffix, suffix
	}

	return hasValidSuffix, suffix
}
