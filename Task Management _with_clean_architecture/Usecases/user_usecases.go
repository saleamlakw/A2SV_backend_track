package Usecases

import (
	"context"
	"fmt"
	"github.com/saleamlakw/TaskManager/Domain"
	"github.com/saleamlakw/TaskManager/Repositories"
	"github.com/saleamlakw/TaskManager/Infrastructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase interface{
	GetUserById(ctx context.Context,id string)(Domain.User,error)
	CreateUser(ctx context.Context,newUser Domain.User)(Domain.User,error)
	Login(ctx context.Context,user Domain.User)(string,error)
	GetUsers(ctx context.Context) (*[]Domain.User, error) 
	PromoteUserToAdmin(ctx context.Context, id string) error 
}

type userUsecase struct{
	UserRepository Repositories.UserRepository
}
func NewUserUsecase(repositary Repositories.UserRepository)*userUsecase{
	return &userUsecase{
		UserRepository: repositary,
	}
}

func (uu userUsecase)GetUserById(ctx context.Context,id string)(Domain.User,error){
	var user Domain.User
	user,err:=uu.UserRepository.GetUserById(ctx,id)
	return user,err
}

func (uu userUsecase)CreateUser(ctx context.Context,newUser Domain.User)(Domain.User,error){
	accountExists, err :=uu.UserRepository.AccountExists(ctx,newUser.UserName)
	if err != nil {
		return Domain.User{},err
	}
	if accountExists {
		return Domain.User{},fmt.Errorf("this username is already taken")
	}
	hashedPassword:=Infrastructure.HashPassword(newUser.Password)
	newUser.Password=hashedPassword
	newUser.ID=primitive.NewObjectID().Hex()
	newUser.Password=string(hashedPassword)
	if result:=uu.UserRepository.CountDocuments(ctx);result==0{
		newUser.Role="admin"
	}else{
		newUser.Role="user"
	}
	createdUser,err:=uu.UserRepository.CreateUser(ctx,newUser)
	return createdUser,err
	
}

func (uu userUsecase)Login(ctx context.Context,user Domain.User)(string,error){
	foundUser,err:=uu.UserRepository.GetUserByUsername(ctx,user.UserName)
	if err!=nil{
		return "",fmt.Errorf("username  not found:%v",err)
	}
	passwordIsValid,_:=Infrastructure.VerifyPassword(user.Password,foundUser.Password)
	if !passwordIsValid{
		return "",fmt.Errorf("invalid password : %v",err)
	}
	token,err:=Infrastructure.GenerateToken(foundUser)
	return token,err
}
func (uu userUsecase)GetUsers(ctx context.Context) (*[]Domain.User, error) {
	user,err:=uu.UserRepository.GetUsers(ctx)
	return user,err

}
func (uu userUsecase)PromoteUserToAdmin(ctx context.Context, id string) error {
	user,err:=uu.GetUserById(context.TODO(), id)
	if err != nil {
		return fmt.Errorf("user not found")
	}
	if user.Role=="admin"{
		return fmt.Errorf("the user is already an admin")
	}
	err=uu.UserRepository.PromoteUserToAdmin(ctx,id)
	return err
}