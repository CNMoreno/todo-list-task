package app

import (
	"github.com/google/uuid"
	"todo-list-task/internal/domain"
	"todo-list-task/internal/infrastructure/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u UserService) Register(user *domain.UserRequest) (string, error) {
	saveUser := &domain.User{
		Username: user.Username,
		Password: user.Password,
		ID:       uuid.NewString(),
	}
	return u.repo.Create(saveUser)
}

func (u UserService) Login(user domain.User) (string, error) {

	return u.repo.Login(user)
}
