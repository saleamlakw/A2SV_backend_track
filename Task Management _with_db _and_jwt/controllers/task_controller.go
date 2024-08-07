package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/saleamlakw/TaskManager/data"
	"github.com/saleamlakw/TaskManager/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var db mongo.Database

func Init(database *mongo.Database) {
	db = *database
}
func GetTask(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "role not found"})
		c.Abort()
		return
	}

	urole, ok := role.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role"})
		c.Abort()
		return
	}

	userid, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userid not found"})
		c.Abort()
		return
	}

	uid, ok := userid.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userid"})
		c.Abort()
		return
	}
	tasksdb := db.Collection("tasks")
	tasks, err := data.GetTask(context.TODO(), tasksdb, urole, uid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to retrive tasks"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func PostTask(c *gin.Context) {
	tasksdb := db.Collection("tasks")
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	createdTask, err := data.CreateTask(context.TODO(), newTask, tasksdb)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "task not created"})
		return
	}
	c.IndentedJSON(http.StatusCreated, createdTask)
}

func GetTaskById(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "role not found"})
		c.Abort()
		return
	}

	urole, ok := role.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role"})
		c.Abort()
		return
	}

	userid, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userid not found"})
		c.Abort()
		return
	}

	uid, ok := userid.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userid"})
		c.Abort()
		return
	}
	tasksdb := db.Collection("tasks")
	id := c.Param("id")
	if urole == "user" {
		IsUserAssignedToTask := data.IsUserAssignedToTask(context.TODO(), id, tasksdb, urole, uid)
		if !IsUserAssignedToTask {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "this task is not assigned to this user"})
			c.Abort()
			return
		}
	}

	task, err := data.GetTaskById(context.TODO(), id, tasksdb)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)

}

func DeleteTask(c *gin.Context) {
	tasksdb := db.Collection("tasks")
	id := c.Param("id")
	if err := data.DeleteTask(context.TODO(), id, tasksdb); err == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func UpdateTask(c *gin.Context) {
	tasksdb := db.Collection("tasks")
	id := c.Param("id")
	var updatedTask models.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		return
	}
	task, err := data.UpdateTask(context.TODO(), updatedTask, id, tasksdb)
	if err == nil {
		c.IndentedJSON(http.StatusOK, task)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
func Getuser(c *gin.Context) {
	user_id := c.Param("id")
	usersdb := db.Collection("users")
	user, err := data.GetUserById(context.TODO(), user_id, usersdb)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}
func GetUsers(c *gin.Context) {

}


func SignUp(c *gin.Context) {
	var newUser models.User
	err := c.BindJSON(&newUser)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	validate := validator.New()
	validateErr := validate.Struct(newUser)
	if validateErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": validateErr.Error()})
		return
	}
	usersdb := db.Collection("users")
	accountExists, err := data.AccountExists(context.TODO(), usersdb, newUser.UserName)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if accountExists {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "this username is already taken"})
		return
	}
	createdUser, err := data.CreateUser(context.TODO(), newUser, usersdb)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, createdUser)
}
func Login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userdb := db.Collection("users")
	token, _ := data.Login(context.TODO(), user, userdb)
	c.IndentedJSON(http.StatusOK, token)
}

func PromoteUserToAdmin(c *gin.Context) {
	id := c.Param("id")
	userdb := db.Collection("users")
	err := data.PromoteUserToAdmin(context.TODO(), userdb, id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "role updated sucessfully"})
}
