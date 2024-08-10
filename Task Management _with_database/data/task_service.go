package data

import (
	"context"

	"github.com/saleamlakw/TaskManagement/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskService interface{
	GetTask(ctx context.Context)  (*[]models.Task,error)
	CreateTask(ctx context.Context,newTask models.Task) (models.Task,error)
	GetTaskById(ctx context.Context,id string)(models.Task,error)
	DeleteTask(ctx context.Context,id string)error
	UpdateTask(ctx context.Context,updatedTask models.Task,id string)(models.Task,error)

}

type taskService struct{
	db *mongo.Collection
}
func NewTaskService(client *mongo.Client)*taskService{
	return &taskService{
		db:client.Database("taskmanager").Collection("tasks"),
	}
}

func (ts *taskService)GetTask(ctx context.Context)  (*[]models.Task,error){
	tasks:=[]models.Task{} 
	cursor,err:=ts.db.Find(ctx,bson.D{{}});
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
func (ts *taskService)CreateTask(ctx context.Context,newTask models.Task) (models.Task,error){
	newTask.ID = primitive.NewObjectID().Hex()
	_,err:=ts.db.InsertOne(ctx,newTask)
	return newTask,err
}
func (ts *taskService)GetTaskById(ctx context.Context,id string)(models.Task,error){
	var task models.Task
	query:=bson.D{bson.E{Key:"id",Value:id}}
	err:=ts.db.FindOne(ctx,query).Decode(&task)
	
	return task,err
}
func (ts *taskService)DeleteTask(ctx context.Context,id string)error{
	query:=bson.D{bson.E{Key:"id",Value:id}}
	if _,err:=ts.db.DeleteOne(ctx,query);err!=nil{
		return err
	}
return nil
}

func (ts *taskService)UpdateTask(ctx context.Context,updatedTask models.Task,id string)(models.Task,error){
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
	err := ts.db.FindOneAndUpdate(ctx, query, update, opts).Decode(&updatedResult)
	if err != nil {
		return models.Task{}, err
	}
	return updatedResult, nil
	}
