package main
import (
	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManagement/router"
)
func main(){
	ro :=gin.Default()
	router.Route(ro)
	ro.Run("localhost:8080")
}