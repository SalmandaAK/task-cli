package main

import (
	"os"
	"path/filepath"

	"github.com/SalmandaAK/task-cli/internal/console/controller"
	"github.com/SalmandaAK/task-cli/internal/console/router"
	"github.com/SalmandaAK/task-cli/internal/task/db"
	"github.com/SalmandaAK/task-cli/internal/task/service"
)

func main() {
	filePath := filepath.Join(filepath.Dir(""), "task.json")
	repo := db.NewTaskJSONRepository(filePath)
	service := service.NewTaskService(repo)
	controller := controller.NewTaskHandler(service)
	router := router.NewRouter(controller)

	router.Run(os.Args)
}
