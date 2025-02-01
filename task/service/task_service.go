package service

import (
	"slices"
	"time"

	"github.com/SalmandaAK/task-cli/task/domain"
)

type TaskService struct {
	r domain.TaskRepository
}

func NewTaskService(r domain.TaskRepository) *TaskService {
	return &TaskService{
		r: r,
	}
}

func (s *TaskService) AddTask(description string) (int, error) {
	// New Task Id will be generated in repository layer
	t := &domain.Task{
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format(time.DateTime),
	}
	err := s.r.CreateTask(t)
	if err != nil {
		return 0, err
	}
	return t.Id, nil
}

func (s *TaskService) ListTasks() ([]*domain.Task, error) {
	tasks, err := s.r.FindAllTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) ListTasksByStatus(status string) ([]*domain.Task, error) {
	var filteredTask []*domain.Task
	tasks, err := s.r.FindAllTasks()
	if err != nil {
		return nil, err
	}
	if len(tasks) > 0 {
		filteredTask = slices.DeleteFunc(tasks, func(t *domain.Task) bool {
			return t.Status != status
		})
	}
	return filteredTask, nil
}

func (s *TaskService) UpdateTask(id int, description string) error {
	t, err := s.r.FindTaskById(id)
	if err != nil {
		return err
	}
	t.Description = description
	t.UpdatedAt = time.Now().Format(time.DateTime)
	return s.r.UpdateTask(t)
}

func (s *TaskService) DeleteTask(id int) error {
	t, err := s.r.FindTaskById(id)
	if err != nil {
		return err
	}
	return s.r.DeleteTask(t)
}

func (s *TaskService) MarkTaskStatus(id int, status string) error {
	t, err := s.r.FindTaskById(id)
	if err != nil {
		return err
	}
	t.Status = status
	t.UpdatedAt = time.Now().Format(time.DateTime)
	return s.r.UpdateTask(t)
}
