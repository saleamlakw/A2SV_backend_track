package router

import (
	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManagement/controllers"
	"github.com/saleamlakw/TaskManagement/data"
)
func Route(router *gin.Engine,service data.TaskService){
	tc:=controllers.NewTaskController(service)
	router.GET("/tasks",tc.GetTask)
	router.GET("/tasks/:id",tc.GetTaskById)
	router.POST("/tasks/",tc.PostTask)
	router.DELETE("/tasks/:id",tc.DeleteTask)
	router.PUT("/tasks/:id",tc.UpdateTask)
}