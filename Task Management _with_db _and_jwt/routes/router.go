package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManager/controllers"
)

func Route(router *gin.Engine) {
	router.GET("/tasks", controllers.GetTask)
	router.GET("/tasks/:id", controllers.GetTaskById)
	router.POST("/tasks", controllers.PostTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
}
