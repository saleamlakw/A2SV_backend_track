package controllers

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManagement/data"
	"github.com/saleamlakw/TaskManagement/models"
	"go.mongodb.org/mongo-driver/mongo"
)
var db mongo.Collection
func Init(database *mongo.Database){
	db=*database.Collection("tasks")
}
func GetTask(c *gin.Context){
	tasks,err:=data.GetTask(context.TODO(),&db)
	if err!=nil{
		c.IndentedJSON(http.StatusInternalServerError,gin.H{"message":"failed to retrive tasks"})
		return
	}
	c.IndentedJSON(http.StatusOK,tasks)
}

func PostTask(c *gin.Context){
	var newTask models.Task
	if err:= c.BindJSON(&newTask);err!=nil{
		return 
	}
	data.CreateTask(context.TODO(),newTask,&db)
	c.IndentedJSON(http.StatusCreated,newTask)
}

func GetTaskById(c *gin.Context){
	id:=c.Param("id")
	task,err:=data.GetTaskById(context.TODO(),id,&db)
	if err!=nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK,task)
	
}

func DeleteTask(c *gin.Context){
	id:=c.Param("id")
	if err:=data.DeleteTask(context.TODO(),id,&db);err==nil{
		c.IndentedJSON(http.StatusOK,gin.H{"message":"task deleted"})
		return
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"task not found"})
}

func UpdateTask(c *gin.Context){
	id:=c.Param("id")
	var updatedTask models.Task
	if err:=c.BindJSON(&updatedTask);err!=nil{
		return
	}
	task,err:=data.UpdateTask(context.TODO(),updatedTask,id,&db)
	if err==nil{
		c.IndentedJSON(http.StatusOK,task)
		return 
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"task not found"})
}