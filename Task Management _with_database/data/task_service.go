package data

import (
	"context"

	"github.com/saleamlakw/TaskManagement/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func GetTask(ctx context.Context,db *mongo.Collection)  (*[]models.Task,error){
	tasks:=[]models.Task{} 
	cursor,err:=db.Find(ctx,bson.D{{}});
	if err!=nil{
		return nil,err
	}
	for cursor.Next(ctx){
		var task models.Task
		err:=cursor.Decode(&task)
		if err!=nil{
			return nil,err
		}
		tasks=append(tasks, task)

	}
	cursor.Close(ctx)
	return &tasks,nil

}
func CreateTask(ctx context.Context,newTask models.Task,db *mongo.Collection) error{
	newTask.ID = primitive.NewObjectID().Hex()
	_,err:=db.InsertOne(ctx,newTask)
	return err
}
func GetTaskById(ctx context.Context,id string,db *mongo.Collection)(models.Task,error){
	var task models.Task
	query:=bson.D{bson.E{Key:"id",Value:id}}
	err:=db.FindOne(ctx,query).Decode(&task)
	
	return task,err
}
func DeleteTask(ctx context.Context,id string,db *mongo.Collection)error{
	query:=bson.D{bson.E{Key:"id",Value:id}}
	if _,err:=db.DeleteOne(ctx,query);err!=nil{
		return err
	}
return nil
}

func UpdateTask(ctx context.Context,updatedTask models.Task,id string,db *mongo.Collection)(models.Task,error){
	query:=bson.D{{Key:"id",Value:id}}
	update:=bson.D{{Key:"$set" ,Value:bson.D{
		{Key: "title", Value: updatedTask.Title},
		{Key: "description", Value: updatedTask.Description},
		{Key: "duedate", Value: updatedTask.DueDate},
		{Key: "status", Value: updatedTask.Status},
	}},
}
	var updatedResult models.Task
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := db.FindOneAndUpdate(ctx, query, update, opts).Decode(&updatedResult)
	if err != nil {
		return models.Task{}, err
	}
	return updatedResult, nil
	}
