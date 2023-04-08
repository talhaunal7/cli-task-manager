package initialize

import (
	"github.com/boltdb/bolt"
	"log"
	"os"
	"task/controller"
	"task/service"
)

var (
	db             *bolt.DB
	taskService    service.TaskService
	taskController controller.TaskController
)

func Start() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	db, err = bolt.Open(dirname+"/my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	taskService = service.NewTaskService(db)
	taskController = controller.NewTaskController(taskService)

}

func GetController() *controller.TaskController {
	return &taskController
}
