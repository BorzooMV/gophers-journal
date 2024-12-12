package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/BorzooMV/gophers-journal/internal/model"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	qs := "SELECT * FROM posts;"
	rows, err := db.Query(qs)
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't fetch data from database:\n%v", err.Error()), http.StatusInternalServerError)
	}
	defer rows.Close()

	allPosts := []model.Post{}

	for rows.Next() {
		var p model.Post
		err := rows.Scan(&p.Id, &p.Title, &p.Description, &p.Body, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			http.Error(w, fmt.Sprintf("couldn't scan the row\n%v", err.Error()), http.StatusInternalServerError)
		}

		allPosts = append(allPosts, p)
	}

	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(allPosts)
	if err != nil {
		http.Error(w, fmt.Sprintf("couldn't encode posts\n%v", err.Error()), http.StatusInternalServerError)
	}
}

func GetPostWithId(w http.ResponseWriter, r *http.Request, db *sql.DB, id string) {
	qs := "SELECT * FROM posts WHERE id = $1 LIMIT 1;"
	var post model.Post
	row := db.QueryRow(qs, id)
	err := row.Scan(&post.Id, &post.Title, &post.Description, &post.Body, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		http.Error(w, "can't find requested data", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func CreateNewPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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

	qs := "INSERT INTO posts (title, description, body) VALUES($1,$2,$3);"

	// Query database
	_, err = db.Exec(qs, newPost.Title, newPost.Description, newPost.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("couldn't query database\n%v", err.Error()), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newPost)
}
