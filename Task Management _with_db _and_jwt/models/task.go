package models

import (
	"time"
)

type Task struct {
	ID          string    `json:"_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
	User_id     string    `json:"user_id"`
}
