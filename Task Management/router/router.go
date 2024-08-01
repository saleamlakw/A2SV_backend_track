package router

import (
	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManagement/controllers"
)
func Route(router *gin.Engine){
	router.GET("/tasks",controllers.GetTask)
	router.GET("/tasks/:id",controllers.GetTaskById)
	router.POST("/tasks/:id",controllers.PostTask)
	router.DELETE("/tasks/:id",controllers.DeleteTask)
	router.PUT("/tasks/:id",controllers.UpdateTask)
}