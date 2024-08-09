package controllers

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManager/Usecases"
	"github.com/saleamlakw/TaskManager/Domain"
	"github.com/go-playground/validator/v10"
)
type taskController struct{
	UseCase Usecases.TaskUsecase
}
func NewTaskController(usecase Usecases.TaskUsecase )*taskController{
	return &taskController{
		UseCase: usecase,
	}
}
func (tc *taskController)GetTask(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "role not found"})
		return
	}

	urole, ok := role.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role"})
		return
	}

	userid, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userid not found"})
		return
	}

	uid, ok := userid.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userid"})
		return
	}
	tasks, err := tc.UseCase.GetTasks(context.TODO(), urole, uid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to retrive tasks"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func (tc *taskController)PostTask(c *gin.Context) {
	var newTask Domain.Task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	validate := validator.New()
	validateErr := validate.Struct(newTask)
	if validateErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": validateErr.Error()})
		return
	}

	createdTask, err := tc.UseCase.CreateTask(context.TODO(), newTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, createdTask)
}


func (tc *taskController)GetTaskById(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "role not found"})
		return
	}

	urole, ok := role.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role"})
		return
	}

	userid, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userid not found"})
		return
	}

	uid, ok := userid.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userid"})
		return
	}

	id := c.Param("id")
	task, err := tc.UseCase.GetTaskById(context.TODO(), id,urole,uid)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)

}

func (tc *taskController)DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := tc.UseCase.DeleteTask(context.TODO(), id); err == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func (tc *taskController)UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask Domain.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	task, err := tc.UseCase.UpdateTask(context.TODO(), updatedTask, id)
	if err == nil {
		c.IndentedJSON(http.StatusOK, task)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
