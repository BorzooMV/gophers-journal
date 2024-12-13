package main

import (
	"fmt"
	"log"

	"github.com/BorzooMV/gophers-journal/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := services.ConnectDb()
	defer db.Close()

	fmt.Println("Start cleaning posts table in database...")
	db.Exec("DELETE FROM posts;")
	fmt.Println("posts table is empty")
	db.Exec("ALTER SEQUENCE posts_id_seq RESTART WITH 1")
}
