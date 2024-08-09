package Usecases

import (
	"context"
	"fmt"

	"github.com/saleamlakw/TaskManager/Domain"
	"github.com/saleamlakw/TaskManager/Repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type TaskUsecase interface{
	GetTasks(ctx context.Context, role string, uid string) (*[]Domain.Task, error)
	CreateTask(ctx context.Context, newTask Domain.Task) (Domain.Task, error)
	IsUserAssignedToTask(ctx context.Context, id string, uid string) bool
	GetTaskById(ctx context.Context, id string,urole string,uid string) (Domain.Task, error)
	DeleteTask(ctx context.Context, id string) error
	UpdateTask(ctx context.Context, updatedTask Domain.Task, id string) (Domain.Task, error)
}

type taskUsecase struct{
	TaskRepository Repositories.TaskRepository
	UserRepository Repositories.UserRepository
}
func NewTaskUsecase(taskrepositary Repositories.TaskRepository,userrepositary Repositories.UserRepository)*taskUsecase{
	return &taskUsecase{
		TaskRepository: taskrepositary,
		UserRepository:userrepositary,

	}
}
func(tu *taskUsecase) GetTasks(ctx context.Context, role string, uid string) (*[]Domain.Task, error) {
	var filter interface{}
	if role == "admin" {
		filter = bson.D{{}}
	} else {
		filter = bson.M{"user_id": uid}
	}
	task,err:=tu.TaskRepository.GetTasks(ctx,filter,uid)
	return task,err

}
func (tu *taskUsecase)CreateTask(ctx context.Context, newTask Domain.Task) (Domain.Task, error) {
	exists, _ := tu.UserRepository.UserExists(ctx,newTask.User_id)
	if !exists {
		return Domain.Task{}, fmt.Errorf("user not found")
	}
	newTask.ID = primitive.NewObjectID().Hex()
	createdTask,err:=tu.TaskRepository.CreateTask(ctx,newTask)
	return createdTask, err
}
func (tu *taskUsecase)IsUserAssignedToTask(ctx context.Context, id string, uid string) bool {
	return tu.TaskRepository.IsUserAssignedToTask(ctx,id,uid)
}
func (tu *taskUsecase)GetTaskById(ctx context.Context, id string ,urole string,uid string) (Domain.Task, error) {
	if urole == "user" {
		IsUserAssignedToTask := tu.IsUserAssignedToTask(context.TODO(), id, uid)
		if !IsUserAssignedToTask {
			return Domain.Task{},fmt.Errorf("user is not assigned to this task")
		}
	}
	task,err:=tu.TaskRepository.GetTaskById(ctx,id)
	return task, err
}
func (tu *taskUsecase)DeleteTask(ctx context.Context, id string) error {
	err:=tu.TaskRepository.DeleteTask(ctx,id)
	return err
}

func (tu *taskUsecase)UpdateTask(ctx context.Context, updatedTask Domain.Task, id string) (Domain.Task, error) {
	updatedResult,err:=tu.TaskRepository.UpdateTask(ctx,updatedTask,id)
	return updatedResult, err
}
