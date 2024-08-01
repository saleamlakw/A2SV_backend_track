package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManagement/models"
	"github.com/saleamlakw/TaskManagement/data"
	
)
var tasks []models.Task
func init (){
	tasks=*data.NewTask()
}
func GetTask(c *gin.Context){
	c.IndentedJSON(http.StatusOK,tasks)
}

func PostTask(c *gin.Context){
	var newTask models.Task
	if err:= c.BindJSON(&newTask);err!=nil{
		return 
	}
	tasks=append(tasks,newTask)
	c.IndentedJSON(http.StatusCreated,newTask)
}

func GetTaskById(c *gin.Context){
	id:=c.Param("id")
	for _,task := range tasks{
		if task.ID == id{
			c.IndentedJSON(http.StatusOK,task)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"task not found"})
}

func DeleteTask(c *gin.Context){
	id:=c.Param("id")
	for ind,task :=range tasks{
		if task.ID==id{
			tasks=append(tasks[:ind],tasks[ind+1:]...)
			c.IndentedJSON(http.StatusOK,gin.H{"message":"task deleted"})
			return 
		}
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"task not found"})
}

func UpdateTask(c *gin.Context){
	id:=c.Param("id")
	var updatedTask models.Task
	if err:=c.BindJSON(&updatedTask);err!=nil{
		return
	}
	for i ,task:= range tasks{
		if task.ID==id{
			if updatedTask.Description!=""{
				tasks[i].Description=updatedTask.Description
			}
			if updatedTask.Status!=""{
				tasks[i].Status=updatedTask.Status
			}
			if updatedTask.Title!=""{
				tasks[i].Title=updatedTask.Title
			}
			if !updatedTask.DueDate.IsZero(){
				tasks[i].DueDate=updatedTask.DueDate
			}
			c.IndentedJSON(http.StatusOK,updatedTask)
			return 



		}
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"task not found"})
}