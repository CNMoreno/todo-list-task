package repository

// AppCrypto interface defines the methods for password hashing and comparison.
type AppCrypto interface {
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
	CompareHashAndPassword(hashedPassword []byte, password []byte) error
}
