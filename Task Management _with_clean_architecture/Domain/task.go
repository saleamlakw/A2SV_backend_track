package Domain

import (
	"time"
)

type Task struct {
	ID          string    `json:"_id"`          
	Title       string    `json:"title" validate:"required,min=1,max=100"` 
	Description string    `json:"description" validate:"max=500"`         
	DueDate     time.Time `json:"due_date"`            
	Status      string    `json:"status" validate:"required,oneof=pending in-progress completed"` 
	User_id     string    `json:"user_id" validate:"required"`
}