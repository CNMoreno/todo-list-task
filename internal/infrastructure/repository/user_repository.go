package repository

import "todo-list-task/internal/domain"

// UserRepository defines the interface for user-related persistence operations.
type UserRepository interface {
	Create(user *domain.User) (string, error)
	Login(user domain.User) (string, error)
}
