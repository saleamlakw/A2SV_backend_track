package routes
import (
	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManager/controllers"
	// "github.com/saleamlakw/TaskManager/middleware"
)

func UserRoute(router *gin.Engine){
	// router.Use(middleware.Authenticate())
	router.GET("/users",controllers.GetUsers)
	router.GET("/users/:id",controllers.Getuser)
}