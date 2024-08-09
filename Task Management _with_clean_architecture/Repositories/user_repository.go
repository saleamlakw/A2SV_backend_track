package Repositories

import (
	"context"
	"github.com/saleamlakw/TaskManager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface{
	PromoteUserToAdmin(ctx context.Context, id string) error 
	GetUsers(ctx context.Context) (*[]Domain.User, error)
	CreateUser(ctx context.Context,newUser Domain.User)(Domain.User,error)
	CountDocuments(ctx context.Context)int64
	GetUserByUsername(ctx context.Context,username string)(Domain.User,error)
	GetUserById(ctx context.Context,id string)(Domain.User,error)
	AccountExists(ctx context.Context,username string)(bool,error)
	UserExists(ctx context.Context,id string)(bool,error)

}

type userRepository struct {
	UserCollection *mongo.Collection
}
func NewUserRepository(client *mongo.Client)*userRepository{
	return &userRepository{
		UserCollection:client.Database("taskmanager").Collection("users"),
	}
}

func (ur *userRepository)UserExists(ctx context.Context,id string)(bool,error){
	count,err:=ur.UserCollection.CountDocuments(ctx,bson.M{"id":id})
	if err!=nil{
		return false,err
	}
	return count>0,nil
}

func (ur *userRepository) AccountExists(ctx context.Context,username string)(bool,error){
	count,err:=ur.UserCollection.CountDocuments(ctx,bson.M{"username":username})
	if err!=nil{
		return false,err
	}
	return count>0,nil
}

func (ur *userRepository)GetUserById(ctx context.Context,id string)(Domain.User,error){
	var user Domain.User
	err:=ur.UserCollection.FindOne(ctx,bson.M{"id":id}).Decode(&user) 
	return user,err
}

func (ur *userRepository)GetUserByUsername(ctx context.Context,username string)(Domain.User,error){
	var user Domain.User
	err:=ur.UserCollection.FindOne(ctx,bson.M{"username":username}).Decode(&user) 
	print("----",user.UserName)
	return user,err
}

 func (ur *userRepository)CountDocuments(ctx context.Context)int64{
	result,_:=ur.UserCollection.EstimatedDocumentCount(ctx)
	return result
 }
func (ur *userRepository) CreateUser(ctx context.Context,newUser Domain.User)(Domain.User,error){
	_,err:=ur.UserCollection.InsertOne(ctx,newUser)
	return newUser,err
	
}


func (ur *userRepository)GetUsers(ctx context.Context) (*[]Domain.User, error) {
	users := []Domain.User{}
	cursor, err := ur.UserCollection.Find(ctx,bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var user Domain.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	cursor.Close(ctx)
	return &users, nil

}
func (ur *userRepository)PromoteUserToAdmin(ctx context.Context, id string) error {
	query := bson.M{"id": id}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "role", Value: "admin"},
	}},
	}
	_,err := ur.UserCollection.UpdateOne(ctx, query, update)
	if err != nil {
		return err
	}
	return nil
}