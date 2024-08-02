package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManagement/models"
	"github.com/saleamlakw/TaskManagement/data"
	
)

func GetTask(c *gin.Context){
	tasks:=data.GetTask()
	c.IndentedJSON(http.StatusOK,tasks)
}

func PostTask(c *gin.Context){
	var newTask models.Task
	if err:= c.BindJSON(&newTask);err!=nil{
		return 
	}
	data.CreateTask(newTask)
	c.IndentedJSON(http.StatusCreated,newTask)
}

func GetTaskById(c *gin.Context){
	id:=c.Param("id")
	task,err:=data.GetTaskById(id)
	if err==nil{
		c.IndentedJSON(http.StatusOK,task)
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"task not found"})
}

func DeleteTask(c *gin.Context){
	id:=c.Param("id")
	if err:=data.DeleteTask(id);err==nil{
		c.IndentedJSON(http.StatusOK,gin.H{"message":"task deleted"})
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"task not found"})
}

func UpdateTask(c *gin.Context){
	id:=c.Param("id")
	var updatedTask models.Task
	if err:=c.BindJSON(&updatedTask);err!=nil{
		return
	}
	task,err:=data.UpdateTask(updatedTask,id)
	if err==nil{
		c.IndentedJSON(http.StatusOK,task)
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"task not found"})
}