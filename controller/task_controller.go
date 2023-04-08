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

func (tsk *TaskController) Add(args []string) {
	err := tsk.TaskService.Add(args)
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

func (tsk *TaskController) Do(key string) {
	err := tsk.TaskService.Do(key)
	if err != nil {
		log.Fatal(err)
	}
}
