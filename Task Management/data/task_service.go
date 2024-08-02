package data

import (
	"errors"
	"time"
	"github.com/saleamlakw/TaskManagement/models"
)
var tasks []models.Task
func init (){
	tasks=*NewTask()
}

func NewTask() *[]models.Task {
	tasks := []models.Task{
		{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
		{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
		{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
	}
	return &tasks
}
func GetTask()  *[]models.Task{
	return &tasks
}
func CreateTask(newTask models.Task){
	tasks=append(tasks,newTask)
}
func GetTaskById(id string)(models.Task,error){
	for _,task := range tasks{
		if task.ID == id{
			return task, nil
		}
	}
	return models.Task{},errors.New("task not found")
}
func DeleteTask(id string)error{
	for ind,task :=range tasks{
		if task.ID==id{
			tasks=append(tasks[:ind],tasks[ind+1:]...)
			
			return nil
		}
	}
	return errors.New("task not found")
}

func UpdateTask(updatedTask models.Task,id string)(models.Task,error){
	for i ,task:= range tasks{
		if task.ID==id{
			if updatedTask.Description!=""{
				tasks[i].Description=updatedTask.Description
			}
			if updatedTask.Status!=""{
				tasks[i].Status=updatedTask.Status
			}
			if updatedTask.Title!=""{
				tasks[i].Title=updatedTask.Title
			}
			if !updatedTask.DueDate.IsZero(){
				tasks[i].DueDate=updatedTask.DueDate
			}
			
			return updatedTask,nil



		}
	}
	return models.Task{},errors.New("Task not found")
}
