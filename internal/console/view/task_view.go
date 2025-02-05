package view

import (
	"fmt"

	"github.com/SalmandaAK/task-cli/internal/task/domain"
)

func DisplayError(actionPerformed string, err error) {
	fmt.Printf("Error %v: %v\n", actionPerformed, err)
}

func DisplayRespond(message string) {
	fmt.Printf("%v\n", message)
}

func DisplayTasks(tasks []*domain.Task) {
	if len(tasks) == 0 {
		fmt.Println("Task list is empty")
		return
	}
	PrintTasks(tasks)
}
