package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/TaskManagement/controllers"
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
	controllers.Init(client.Database("taskmanager"))
	fmt.Println("Connected to MongoDB!")
	ro :=gin.Default()
	router.Route(ro)
	ro.Run("localhost:8080")
	}
