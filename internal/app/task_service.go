package app

import (
	"github.com/google/uuid"
	"todo-list-task/internal/domain"
	"todo-list-task/internal/infrastructure/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (t TaskService) RegisterTask(task *domain.TaskRequest) (*domain.Task, error) {
	taskSave := &domain.Task{
		Title:       task.Title,
		Description: task.Description,
		ID:          uuid.NewString(),
	}
	return t.repo.CreateTask(taskSave)
}

func (t TaskService) GetTask(id string) (*domain.Task, error) {
	return t.repo.GetTask(id)
}
func (t TaskService) GetTasks() ([]*domain.Task, error) {
	return t.repo.GetTasks()
}

func (t TaskService) UpdateTaskByID(id string, task domain.TaskRequest) (*domain.Task, error) {
	taskSave := &domain.Task{
		Title:       task.Title,
		Description: task.Description,
		Completed:   true,
	}
	return t.repo.UpdateTask(id, taskSave)
}

func (t TaskService) DeleteTaskByID(id string) error {
	return t.repo.DeleteTask(id)
}
