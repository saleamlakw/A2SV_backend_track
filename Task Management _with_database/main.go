package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/saleamlakw/TaskManagement/data"
	"github.com/saleamlakw/TaskManagement/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func main(){
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	
	service:=data.NewTaskService(client)

	fmt.Println("Connected to MongoDB!")
	ro :=gin.Default()
	router.Route(ro,service)
	err=godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error loading .env file")
	}
	mongo_url:=os.Getenv("MONGO_URL")
	ro.Run(mongo_url)
	}
