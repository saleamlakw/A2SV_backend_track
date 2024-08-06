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
func Init(database *mongo.Database){
	db=*database
}
func GetTask(c *gin.Context){
	tasksdb:=db.Collection("tasks")
	tasks,err:=data.GetTask(context.TODO(),tasksdb)
	if err!=nil{
		c.IndentedJSON(http.StatusInternalServerError,gin.H{"message":"failed to retrive tasks"})
		return
	}
	c.IndentedJSON(http.StatusOK,tasks)
}

func PostTask(c *gin.Context){
	tasksdb:=db.Collection("tasks")
	var newTask models.Task
	if err:= c.BindJSON(&newTask);err!=nil{
		return 
	}
	err:=data.CreateTask(context.TODO(),newTask,tasksdb)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"task not created"})
		return 
	}
	c.IndentedJSON(http.StatusCreated,newTask)
}

func GetTaskById(c *gin.Context){
	tasksdb:=db.Collection("tasks")
	id:=c.Param("id")
	task,err:=data.GetTaskById(context.TODO(),id,tasksdb)
	if err!=nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK,task)
	
}

func DeleteTask(c *gin.Context){
	tasksdb:=db.Collection("tasks")
	id:=c.Param("id")
	if err:=data.DeleteTask(context.TODO(),id,tasksdb);err==nil{
		c.IndentedJSON(http.StatusOK,gin.H{"message":"task deleted"})
		return
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"task not found"})
}

func UpdateTask(c *gin.Context){
	tasksdb:=db.Collection("tasks")
	id:=c.Param("id")
	var updatedTask models.Task
	if err:=c.BindJSON(&updatedTask);err!=nil{
		return
	}
	task,err:=data.UpdateTask(context.TODO(),updatedTask,id,tasksdb)
	if err==nil{
		c.IndentedJSON(http.StatusOK,task)
		return 
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"task not found"})
}
func Getuser(c *gin.Context){
	user_id:=c.Param("id")
	// if err:=data.MatchUserTypeToUid(c,user_id);err!=nil{
	// 	c.IndentedJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
	// 	return 
	// }
	usersdb:=db.Collection("users")
	user,err:=data.GetUserById(context.TODO(),user_id,usersdb)
	if err!=nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"user not found"})
		return
	}
	c.IndentedJSON(http.StatusOK,user)
}
func GetUsers(c *gin.Context){


}

// func LoginIn(c *gin.Context){
// 	var user models.User
// 	if err:=c.BindJSON(&user);err!= nil{
// 		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
// 	return
// 	}

// 	usersdb:=db.Collection("users")
// 	foundUser:=data.Login(context.TODO(),user,usersdb)
// 	c.IndentedJSON(http.StatusOK,foundUser)
// }

func SignUp(c *gin.Context){
	 var newUser  models.User
	 err:=c.BindJSON(&newUser)
	 if err!= nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	 }
	 validate:=validator.New()
	 validateErr:=validate.Struct(newUser)
	 if validateErr!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":validateErr.Error()})
		return
	 }
	usersdb:=db.Collection("users")
	accountExists,err:=data.AccountExists(context.TODO(),usersdb,newUser.UserName)
	if err!=nil{
		c.IndentedJSON(http.StatusInternalServerError,gin.H{"message":err.Error()})
		return
	}
	if accountExists{
		c.IndentedJSON(http.StatusInternalServerError,gin.H{"message":"this username is already taken"})
		return
	}
	createdUser,err:=data.CreateUser(context.TODO(),newUser,usersdb)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return 
	}
	c.IndentedJSON(http.StatusCreated,createdUser)
}
