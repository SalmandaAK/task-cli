package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/SalmandaAK/task-cli/console/view"
	"github.com/SalmandaAK/task-cli/task/domain"
	"github.com/SalmandaAK/task-cli/task/service"
)

var (
	errEmptyTaskDescription      = errors.New("task description must not be empty")
	errInvalidTaskStatus         = errors.New("invalid task status")
	errEmptyIdAndTaskDescription = errors.New("ID and task description must not be empty")
	errInvalidId                 = errors.New("invalid ID")
	errEmptyId                   = errors.New("ID must not be empty")
)

type TaskController struct {
	ts *service.TaskService
}

func NewTaskHandler(ts *service.TaskService) *TaskController {
	return &TaskController{
		ts: ts,
	}
}

// args: list (<nil> || <status>)
func (c *TaskController) List(args []string) {
	status := args[1]
	if status != "" {
		c.ListByStatus(status)
		return
	}
	tasks, err := c.ts.ListTasks()
	if err != nil {
		view.DisplayError("listing task", err)
		return
	}
	view.DisplayTasks(tasks)
}

func (c *TaskController) ListByStatus(status string) {
	var tasks []*domain.Task
	var err error
	if status == "todo" || status == "in-progress" || status == "done" {
		tasks, err = c.ts.ListTasksByStatus(status)
	} else {
		err = errInvalidTaskStatus
	}
	if err != nil {
		view.DisplayError("listing tasks", err)
		return
	}
	view.DisplayTasks(tasks)
}

// args: add <task description>
func (c *TaskController) AddTask(args []string) {
	description := args[1]
	if description == "" {
		view.DisplayError("adding task", errEmptyTaskDescription)
		return
	}
	id, err := c.ts.AddTask(description)
	if err != nil {
		view.DisplayError("adding task", err)
		return
	}
	view.DisplayRespond(fmt.Sprintf("Task added successfully (ID: %v)", id))
}

// args: update <id> <task description>
func (c *TaskController) UpdateTask(args []string) {
	if args[1] == "" || args[2] == "" {
		view.DisplayError("updating task", errEmptyIdAndTaskDescription)
		return
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		view.DisplayError("updating task", errInvalidId)
		return
	}
	// at this point, description would not be empty
	description := args[2]
	err = c.ts.UpdateTask(id, description)
	if err != nil {
		view.DisplayError("updating task", err)
	}
	view.DisplayRespond(fmt.Sprintf("Task updated successfully (ID: %v)", id))
}

// args: delete <id>
func (c *TaskController) DeleteTask(args []string) {
	if args[1] == "" {
		view.DisplayError("deleting task", errEmptyId)
		return
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		view.DisplayError("deleting task", errInvalidId)
		return
	}
	if err = c.ts.DeleteTask(id); err != nil {
		view.DisplayError("deleting task", err)
		return
	}
	view.DisplayRespond(fmt.Sprintf("Task has been deleted (ID: %v)", id))
}

// args: (mark-in-progress || mark-done) <id>
func (c *TaskController) MarkTaskStatus(args []string) {
	if args[1] == "" {
		view.DisplayError("marking task", errEmptyId)
		return
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		view.DisplayError("marking task", errInvalidId)
		return
	}
	var status string
	if args[0] == "mark-in-progress" {
		status = "in-progress"
	}
	if args[0] == "mark-done" {
		status = "done"
	}
	err = c.ts.MarkTaskStatus(id, status)
	if err != nil {
		view.DisplayError("marking task", err)
		return
	}
	view.DisplayRespond(fmt.Sprintf("Task (ID: %v) marked as '%v' successfully", id, status))
}
