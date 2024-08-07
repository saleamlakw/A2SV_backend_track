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
	fmt.Println("count---",count)
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
// }
// func CheckUserType(c *gin.Context,role string)error{
// 	user_type:=c.GetString("user_type")
// 	if user_type!=role{
// 		return errors.New("unauthorized to access this resource")
// 	}
// 	return nil
// }
// func MatchUserTypeToUid(c *gin.Context,userid string)error{
// 	user_type:=c.GetString("user_type")
// 	uid:=c.GetString("uid")
// 	if user_type=="USER" && uid!=userid{
// 		return errors.New("unauthorized to access this resource")
// 	}
// 	err:=CheckUserType(c,user_type)
// 	return err 
// }
func GetUserById(ctx context.Context,id string ,db *mongo.Collection)(models.User,error){
	var user models.User
	err:=db.FindOne(ctx,bson.M{"id":id}).Decode(&user) 
	return user,err
}
// type SignedDetails struct{
// 	Email      string
// 	First_name string
// 	Last_name  string 
// 	Uid        string
// 	User_type  string 
// 	jwt.StandardClaims   
// }
// var SecretKey = os.Getenv("SECRET_KEY")

// func GenerateAllTokens(email string,fname string ,lname string ,user_type string,user_id string)(string,string,error){
// 	clamis:=&SignedDetails{
// 		Email: email,
// 		First_name: fname,
// 		Last_name: lname,
// 		Uid: user_id,
// 		User_type: user_type,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt : time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
// 		},
// 	}
// 	refreshClaims := &SignedDetails{
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
// 		},
// 	}
// 	token,err:= jwt.NewWithClaims(jwt.SigningMethodHS256,clamis).SignedString([]byte(SecretKey))
// 	if err != nil{
// 		log.Panic(err)
// 		return "","",err
// 	}
// 	refreshToken,err:= jwt.NewWithClaims(jwt.SigningMethodHS256,refreshClaims).SignedString([]byte(SecretKey))
// 	if err != nil{
// 		log.Panic(err)
// 		return "","",err
// 	}
// 	return token,refreshToken,err
// }

func CreateUser(ctx context.Context,newUser models.User,db *mongo.Collection)(models.User,error){
	hashedPassword:=HashPassword(newUser.Password)
	newUser.Password=hashedPassword
	newUser.ID=primitive.NewObjectID().Hex()
	newUser.Password=string(hashedPassword)
	if result,_:=db.EstimatedDocumentCount(ctx);result==0{
		// fmt.Println("+++",result)
		newUser.Role="admin"
	}else{
		// fmt.Println("+++",result)
		newUser.Role="user"
	}
	// newUser.User_id=newUser.ID.Hex()
	// token,refreshToken,_:= GenerateAllTokens(newUser.Email,newUser.FirstName,newUser.LastName,newUser.User_type,newUser.User_id) 
	// newUser.Token=token
	// newUser.Refresh_token=refreshToken
	_,err:=db.InsertOne(ctx,newUser)
	return newUser,err
	
}
// func UpdateAllTokens(ctx context.Context,db *mongo.Collection,signedToken string, signedRefreshToken string, userId string) {
//     var updateObj primitive.D

//     updateObj = append(updateObj, bson.E{"token", signedToken})
//     updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})


//     upsert := true
//     filter := bson.M{"user_id": userId}
//     opt := options.UpdateOptions{
//         Upsert: &upsert,
//     }

//     _, err := db.UpdateOne(
//         ctx,
//         filter,
//         bson.D{
//             {"$set", updateObj},
//         },
//         &opt,
//     )

//     if err != nil {
//         log.Panic(err)
//         return
//     }

//     return
// }
func Login(ctx context.Context,user models.User,db *mongo.Collection)(string,models.User){
	var foundUser models.User
	err:=db.FindOne(context.TODO(),bson.M{"username":user.UserName}).Decode(&foundUser)
	if err!=nil{
		return "",models.User{}
	}
	passwordIsValid,_:=VerifyPassword(user.Password,foundUser.Password)
	if !passwordIsValid{
		return "",models.User{}
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
		return "",models.User{}
	}
	return token,foundUser
}