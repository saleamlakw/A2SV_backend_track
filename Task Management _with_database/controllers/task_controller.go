package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/saleamlakw/TaskManagement/data"
	"github.com/saleamlakw/TaskManagement/models"
)

type taskController struct {
	service data.TaskService
}

func NewTaskController(service data.TaskService) *taskController {
	return &taskController{
		service: service,
	}
}
func (tc *taskController) GetTask(c *gin.Context) {
	tasks, err := tc.service.GetTask(context.TODO())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to retrive tasks"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func (tc *taskController) PostTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request body", "error": err.Error()})
		return
	}

	validate := validator.New()
	validateErr := validate.Struct(newTask)
	if validateErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": validateErr.Error()})
		return
	}

	createdTask, err := tc.service.CreateTask(context.TODO(), newTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, createdTask)
}

func (tc *taskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.service.GetTaskById(context.TODO(), id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)

}

func (tc *taskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := tc.service.DeleteTask(context.TODO(), id); err == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted successfuly"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func (tc *taskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		return
	}
	validate := validator.New()
	validateErr := validate.Struct(updatedTask)
	if validateErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": validateErr.Error()})
		return
	}
	task, err := tc.service.UpdateTask(context.TODO(), updatedTask, id)
	if err == nil {
		c.IndentedJSON(http.StatusOK, task)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
