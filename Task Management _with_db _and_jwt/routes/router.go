package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManager/controllers"
	"github.com/saleamlakw/TaskManager/middleware"
)

func Route(router *gin.Engine) {
	r:=router.Group("/",middleware.AuthMiddleware)
	r.GET("/tasks",controllers.GetTask)
	r.GET("/tasks/:id",controllers.GetTaskById)
	r.POST("/tasks",middleware.IsAdminMiddleware, controllers.PostTask)
	r.DELETE("/tasks/:id",middleware.IsAdminMiddleware,controllers.DeleteTask)
	r.PUT("/tasks/:id",middleware.IsAdminMiddleware,controllers.UpdateTask)

	router.POST("/user/signup",controllers.SignUp)
	router.POST("/user/login",controllers.Login)

	r.GET("/users",middleware.IsAdminMiddleware,controllers.GetUsers)
	r.GET("/users/:id",middleware.IsAdminMiddleware,controllers.Getuser)
	r.POST("/promote/:id",middleware.IsAdminMiddleware,controllers.PromoteUserToAdmin)
}
