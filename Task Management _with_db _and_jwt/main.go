package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManager/controllers"
	"github.com/saleamlakw/TaskManager/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)
func main(){
	err:=godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error loading .env file")
	}
	url:=os.Getenv("MONGODB_URL")
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	controllers.Init(client.Database("taskmanager"))
	fmt.Println("Connected to MongoDB!")
	router :=gin.Default()
	routes.Route(router)
	routes.AuthRoute(router)
	routes.UserRoute(router)
	port:=os.Getenv("PORT")
	router.Run("localhost:"+port)
	}
