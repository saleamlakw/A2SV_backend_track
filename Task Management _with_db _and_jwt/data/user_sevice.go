package data

import (
	"context"
	"fmt"
	// "errors"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	// "github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)
func AccountExists(ctx context.Context,usersdb *mongo.Collection,username string)(bool,error){
	count,err:=usersdb.CountDocuments(ctx,bson.M{"username":username})
	if err!=nil{
		return false,err
	}
	return count>0,nil
}
func HashPassword(Password string)string{
	hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(Password),bcrypt.DefaultCost)
	print(hashedPassword)
	if err!=nil{
		log.Panic(err)
	}
	return string(hashedPassword)
}
func VerifyPassword(userPassword string ,providedPassword string)(bool,string){
	msg:=""
	check:=true
	err:=bcrypt.CompareHashAndPassword([]byte(providedPassword),[]byte(userPassword))
	if err!=nil{
		msg="passowrd is incorrect"
        check=false
	}
	return check,msg
}
func GetUserById(ctx context.Context,id string ,db *mongo.Collection)(models.User,error){
	var user models.User
	err:=db.FindOne(ctx,bson.M{"id":id}).Decode(&user) 
	return user,err
}

func CreateUser(ctx context.Context,newUser models.User,db *mongo.Collection)(models.User,error){
	hashedPassword:=HashPassword(newUser.Password)
	newUser.Password=hashedPassword
	newUser.ID=primitive.NewObjectID().Hex()
	newUser.Password=string(hashedPassword)
	if result,_:=db.EstimatedDocumentCount(ctx);result==0{
		newUser.Role="admin"
	}else{
		newUser.Role="user"
	}
	_,err:=db.InsertOne(ctx,newUser)
	return newUser,err
	
}

func Login(ctx context.Context,user models.User,db *mongo.Collection)(string,error){
	var foundUser models.User
	err:=db.FindOne(context.TODO(),bson.M{"username":user.UserName}).Decode(&foundUser)
	if err!=nil{
		return "",fmt.Errorf("user not found:%v",err)
	}
	passwordIsValid,_:=VerifyPassword(user.Password,foundUser.Password)
	if !passwordIsValid{
		return "",fmt.Errorf("invalid password : %v",err)
	}
	claims := jwt.MapClaims{
        "userid":   foundUser.ID,
        "username": foundUser.UserName,
        "role":     foundUser.Role,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
    }
	var SecretKey = os.Getenv("SECRET_KEY")
	token,err:= jwt.NewWithClaims(jwt.SigningMethodHS256,claims).SignedString([]byte(SecretKey))
	if err!=nil{
		return "",fmt.Errorf("error generating token : %v",err)
	}
	return token,nil
}
func GetUsers(ctx context.Context, db *mongo.Collection, role string, uid string) (*[]models.User, error) {
	var filter interface{}
	users := []models.User{}
	if role == "admin" {
		filter = bson.D{{}}
	} else {
		filter = bson.M{"user_id": uid}
	}
	cursor, err := db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	cursor.Close(ctx)
	return &users, nil

}
func PromoteUserToAdmin(ctx context.Context, db *mongo.Collection, id string) error {
	query := bson.M{"id": id}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "role", Value: "admin"},
	}},
	}
	_,err := db.UpdateOne(ctx, query, update)
	if err != nil {
		return err
	}
	return nil
}