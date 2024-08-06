package routes
import (
	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManager/controllers"
)

func AuthRoute(router *gin.Engine){
	router.POST("/user/signup",controllers.SignUp)
	// router.POST("/user/login",controllers.LoginIn)
} 
