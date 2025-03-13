package app

import "golang.org/x/crypto/bcrypt"

// BcryptCrypto struct implements the AppCrypto interface using bcrypt.
type BcryptCrypto struct{}

// GenerateFromPassword wraps bcrypt GenerateFromPassword function.
func (BcryptCrypto) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

// CompareHashAndPassword wraps bcrypt CompareHashAndPassword function.
func (BcryptCrypto) CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
