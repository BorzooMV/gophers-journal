package model

import (
	"fmt"
	"time"
)

type Post struct {
	Id          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Body        string     `json:"body"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (p Post) Validate() error {
	if p.Title == "" {
		return fmt.Errorf("title not provided")
	}

	if p.Description == "" {
		return fmt.Errorf("description not provided")
	}

	if p.Body == "" {
		return fmt.Errorf("body not provided")
	}

	return nil
}
