package controller

import (
	"log"
	"task/service"
)

type TaskController struct {
	TaskService service.TaskService
}

func NewTaskController(taskService service.TaskService) TaskController {
	return TaskController{
		TaskService: taskService,
	}
}

func (tsk *TaskController) Add() {
	err := tsk.TaskService.Add()
	if err != nil {
		log.Fatal(err)
	}
}

func (tsk *TaskController) List() {
	err := tsk.TaskService.List()
	if err != nil {
		log.Fatal(err)
	}
}
