package repository

import "todo-list-task/internal/domain"

// TaskRepository defines the interface for task persistence operations.
type TaskRepository interface {
	CreateTask(task *domain.Task) (*domain.Task, error)
	GetTask(id string) (*domain.Task, error)
	GetTasks() ([]*domain.Task, error)
	UpdateTask(id string, task *domain.Task) (*domain.Task, error)
	DeleteTask(id string) error
}
