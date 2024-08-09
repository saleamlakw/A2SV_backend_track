package Repositories

import (
	"context"

	"github.com/saleamlakw/TaskManager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, filter interface{}, uid string) (*[]Domain.Task, error)
	CreateTask(ctx context.Context, newTask Domain.Task) (Domain.Task, error)
	IsUserAssignedToTask(ctx context.Context, id string, uid string) bool
	GetTaskById(ctx context.Context, id string) (Domain.Task, error)
	DeleteTask(ctx context.Context, id string) error
	UpdateTask(ctx context.Context, updatedTask Domain.Task, id string) (Domain.Task, error)
}

type taskRepository struct {
	TaskCollection *mongo.Collection
}

func NewTaskRepository(client *mongo.Client) *taskRepository {
	return &taskRepository{
		TaskCollection: client.Database("taskmanager").Collection("tasks"),
	}
}

func (tr *taskRepository) GetTasks(ctx context.Context, filter interface{}, uid string) (*[]Domain.Task, error) {
	tasks := []Domain.Task{}
	cursor, err := tr.TaskCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var task Domain.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)

	}
	defer cursor.Close(ctx)
	return &tasks, nil

}
func (tr *taskRepository) CreateTask(ctx context.Context, newTask Domain.Task) (Domain.Task, error) {
	_, err := tr.TaskCollection.InsertOne(ctx, newTask)
	return newTask, err
}
func (tr *taskRepository) IsUserAssignedToTask(ctx context.Context, id string, uid string) bool {
	var task Domain.Task
	query := bson.D{bson.E{Key: "id", Value: id}}
	err := tr.TaskCollection.FindOne(ctx, query).Decode(&task)
	if err != nil {
		return false
	}
	return task.User_id == uid
}
func (tr *taskRepository) GetTaskById(ctx context.Context, id string) (Domain.Task, error) {
	var task Domain.Task
	query := bson.D{bson.E{Key: "id", Value: id}}
	err := tr.TaskCollection.FindOne(ctx, query).Decode(&task)

	return task, err
}
func (tr *taskRepository) DeleteTask(ctx context.Context, id string) error {
	query := bson.D{bson.E{Key: "id", Value: id}}
	if _, err := tr.TaskCollection.DeleteOne(ctx, query); err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(ctx context.Context, updatedTask Domain.Task, id string) (Domain.Task, error) {
	query := bson.D{{Key: "id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "title", Value: updatedTask.Title},
		{Key: "description", Value: updatedTask.Description},
		{Key: "duedate", Value: updatedTask.DueDate},
		{Key: "status", Value: updatedTask.Status},
	}},
	}
	var updatedResult Domain.Task
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := tr.TaskCollection.FindOneAndUpdate(ctx, query, update, opts).Decode(&updatedResult)
	if err != nil {
		return Domain.Task{}, err
	}
	return updatedResult, nil
}
