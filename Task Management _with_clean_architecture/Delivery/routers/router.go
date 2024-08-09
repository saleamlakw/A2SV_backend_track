package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManager/Delivery/controllers"
	"github.com/saleamlakw/TaskManager/Infrastructure"
	"github.com/saleamlakw/TaskManager/Repositories"
	"github.com/saleamlakw/TaskManager/Usecases"
	"go.mongodb.org/mongo-driver/mongo"
)

func Route(router *gin.Engine,client *mongo.Client) {
	tr:=Repositories.NewTaskRepository(client)
	ur:=Repositories.NewUserRepository(client)

	tu:=Usecases.NewTaskUsecase(tr,ur)
	uu:=Usecases.NewUserUsecase(ur)

	tc:=controllers.NewTaskController(tu)
	uc:=controllers.NewUserController(uu)
	r:=router.Group("/",Infrastructure.AuthMiddleware)
	r.GET("/tasks",tc.GetTask)
	r.GET("/tasks/:id",tc.GetTaskById)
	r.POST("/tasks",Infrastructure.IsAdminMiddleware, tc.PostTask)
	r.DELETE("/tasks/:id",Infrastructure.IsAdminMiddleware,tc.DeleteTask)
	r.PUT("/tasks/:id",Infrastructure.IsAdminMiddleware,tc.UpdateTask)

	router.POST("/users/signup",uc.SignUp)
	router.POST("/users/login",uc.Login)

	r.GET("/users",Infrastructure.IsAdminMiddleware,uc.GetUsers)
	r.GET("/users/:id",Infrastructure.IsAuthorizedToViewUserDetails,uc.GetuserById)
	r.POST("/promote/:id",Infrastructure.IsAdminMiddleware,uc.PromoteUserToAdmin)
}
