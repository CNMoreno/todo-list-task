package memory

import (
	"fmt"
	"sync"
	"todo-list-task/internal/domain"
	"todo-list-task/internal/utils"
)

type InMemoryUserRepository struct {
	users     map[string]*domain.User
	mu        sync.RWMutex
	appCrypto *utils.DefaultAppCrypto
}

func NewInMemoryUserRepository(appCrypto *utils.DefaultAppCrypto) *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:     make(map[string]*domain.User),
		appCrypto: appCrypto,
	}
}

func (r *InMemoryUserRepository) Create(user *domain.User) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	password, err := r.appCrypto.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = password
	r.users[user.ID] = user
	token, err := utils.GenerateJWT()

	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *InMemoryUserRepository) Login(user domain.User) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, u := range r.users {
		if u.Username == user.Username && r.appCrypto.CheckPasswordHash(user.Password, u.Password) {
			token, err := utils.GenerateJWT()

			if err != nil {
				return "", err
			}

			return token, nil

		}
	}
	return "", fmt.Errorf("username or password incorrect")
}
