package models

import (
	"time"
)

// Todo represents a single todo item
type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewTodo creates a new Todo with the given title and description
func NewTodo(title, description string) *Todo {
	now := time.Now()
	return &Todo{
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
