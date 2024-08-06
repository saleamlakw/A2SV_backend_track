package models

// import (
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

type User struct {
	ID       string 			`json:"_id"`
	UserName string             `json:"username" validate:"required,min=2,max=100"`
	Password string             `json:"password" validate:"required,min=5"`
	Role     string             `json:"role"`
}
