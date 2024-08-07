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
	usersdb:=db.Collection("users")
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	createdTask, err := data.CreateTask(context.TODO(), newTask, tasksdb,usersdb)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
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
func GetuserById(c *gin.Context) {
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
	user_id := c.Param("id")
	if urole=="user" &&user_id!=uid{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "An authorized access"})
		c.Abort()
		return
	}
	usersdb := db.Collection("users")
	user, err := data.GetUserById(context.TODO(), user_id, usersdb)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}
func GetUsers(c *gin.Context) {
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
	usersdb := db.Collection("users")
	users, err := data.GetUsers(context.TODO(), usersdb, urole, uid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to retrive users"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)

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
	token, err := data.Login(context.TODO(), user, userdb)
	if err!=nil{
		c.IndentedJSON(http.StatusUnauthorized,gin.H{"error":err.Error()})
	}
	c.IndentedJSON(http.StatusOK, token)
}

func PromoteUserToAdmin(c *gin.Context) {
	userdb := db.Collection("users")
	id := c.Param("id")
	user,err:=data.GetUserById(context.TODO(), id, userdb)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	if user.Role=="admin"{
		c.JSON(http.StatusBadRequest, gin.H{"error": "the user is alrady an admin"})
		c.Abort()
		return
	}

	
	
	err = data.PromoteUserToAdmin(context.TODO(), userdb, id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "role updated sucessfully"})
}
