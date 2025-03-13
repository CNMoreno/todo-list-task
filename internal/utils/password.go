package utils

import (
	"todo-list-task/internal/infrastructure/repository"

	"golang.org/x/crypto/bcrypt"
)

type DefaultAppCrypto struct {
	crypto repository.AppCrypto
}

func NewHashPassword(crypto repository.AppCrypto) *DefaultAppCrypto {
	return &DefaultAppCrypto{
		crypto: crypto,
	}
}

// HashPassword generate a hash for password with bcrypt.
func (a DefaultAppCrypto) HashPassword(password string) (string, error) {
	bytes, err := a.crypto.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compare plane password with his hash.
func (a DefaultAppCrypto) CheckPasswordHash(password, hash string) bool {
	err := a.crypto.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
