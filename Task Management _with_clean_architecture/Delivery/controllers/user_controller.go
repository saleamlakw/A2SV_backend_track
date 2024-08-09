package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/saleamlakw/TaskManager/Usecases"
	"github.com/saleamlakw/TaskManager/Domain"
)
type userController struct{
	UserUsecase Usecases.UserUsecase
}

func NewUserController(userusecase Usecases.UserUsecase)*userController{
	return &userController{
		UserUsecase: userusecase,
	}
}

func (uc *userController)GetuserById(c *gin.Context) {
	user_id := c.Param("id")
	user, err := uc.UserUsecase.GetUserById(context.TODO(),user_id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}
func (uc *userController) GetUsers(c *gin.Context) {
	users, err := uc.UserUsecase.GetUsers(context.TODO())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to retrive users"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}


func (uc *userController)SignUp(c *gin.Context) {
	var newUser Domain.User
	err := c.BindJSON(&newUser)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	validateErr := validate.Struct(newUser)
	if validateErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
		return
	}
	
	createdUser, err := uc.UserUsecase.CreateUser(context.TODO(), newUser)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, createdUser)
}
func (uc *userController)Login(c *gin.Context) {
	var user Domain.User
	err := c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := uc.UserUsecase.Login(context.TODO(), user)
	if err!=nil{
		c.IndentedJSON(http.StatusUnauthorized,gin.H{"error":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, token)
}

func (uc *userController)PromoteUserToAdmin(c *gin.Context) {
	id := c.Param("id")
	err := uc.UserUsecase.PromoteUserToAdmin(context.TODO(),id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "role updated successfully"})
}
