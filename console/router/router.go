package router

import (
	"fmt"
	"os"

	"github.com/SalmandaAK/task-cli/console/controller"
)

var (
	errEmptyCommand   = "command must not be empty"
	errUnknownCommand = "unknown command"
)

type ControllerFunc func([]string)

func (hf ControllerFunc) Execute(args []string) {
	hf(args)
}

type Router struct {
	m map[string]ControllerFunc
}

func NewRouter(tc *controller.TaskController) *Router {
	return &Router{
		m: map[string]ControllerFunc{
			"add":              tc.AddTask,
			"update":           tc.UpdateTask,
			"delete":           tc.DeleteTask,
			"mark-in-progress": tc.MarkTaskStatus,
			"mark-done":        tc.MarkTaskStatus,
			"list":             tc.List,
		},
	}
}

func (r *Router) Run(osArgs []string) {
	if len(osArgs) < 2 {
		fmt.Fprintf(os.Stderr, "Error input command: %v\n", errEmptyCommand)
		return
	}
	commandArgs := make([]string, 3)
	copy(commandArgs, osArgs[1:]) // copy osArgs into a slice with fixed length to avoid panic out of index when the args are validated in task controller
	taskControllerFunc, ok := r.m[commandArgs[0]]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: %v: %v\n", errUnknownCommand, commandArgs[0])
		return
	}
	taskControllerFunc.Execute(commandArgs)
}
