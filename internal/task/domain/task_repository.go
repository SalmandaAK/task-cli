package domain

type TaskRepository interface {
	FindAllTasks() ([]*Task, error)
	CreateTask(*Task) error
	FindTaskById(id int) (*Task, error)
	UpdateTask(*Task) error
	DeleteTask(*Task) error
}
