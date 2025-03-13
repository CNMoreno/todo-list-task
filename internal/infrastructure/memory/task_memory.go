package memory

import (
	"errors"
	"sync"
	"todo-list-task/internal/domain"
)

type InMemoryTaskRepository struct {
	tasks map[string]*domain.Task
	mu    sync.RWMutex
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]*domain.Task),
	}
}

// CreateTask creates a new task in the in-memory repository.
func (r *InMemoryTaskRepository) CreateTask(task *domain.Task) (*domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[task.ID] = task
	return task, nil
}

// GetTask get a task in the in-memory repository
func (r *InMemoryTaskRepository) GetTask(id string) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return task, nil
}

// GetTasks get all task in the in-memory repository
func (r *InMemoryTaskRepository) GetTasks() ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var tasks []*domain.Task
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// UpdateTask update task by id in the in-memory repository
func (r *InMemoryTaskRepository) UpdateTask(id string, task *domain.Task) (*domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return nil, errors.New("task not found")
	}
	task.ID = id
	r.tasks[id] = task
	return task, nil
}

// DeleteTask delete task by id in the in-memory repository
func (r *InMemoryTaskRepository) DeleteTask(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}
