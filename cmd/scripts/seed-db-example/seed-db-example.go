package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/BorzooMV/gophers-journal/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	var Posts []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Body        string `json:"body"`
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	file, err := os.ReadFile("assets/data/sample-posts-recipe.json")
	if err != nil {
		log.Fatalf("couldn't read the recipe file:\n%v\n", err.Error())
	}

	err = json.Unmarshal(file, &Posts)
	if err != nil {
		log.Fatalf("couldn't unmarshal the recipe file content:\n%v\n", err.Error())
	}

	db := services.ConnectDb()
	defer db.Close()

	fmt.Println("Start seeding database with fake data...")
	for _, item := range Posts {
		qs := "INSERT INTO posts (title, description, body) VALUES ($1,$2,$3);"
		_, err := db.Exec(qs, item.Title, item.Description, item.Body)
		if err != nil {
			log.Fatalf("couldn't insert into db:\n%v\n", err.Error())
		}
	}

	fmt.Println("Seeding completed")
}
