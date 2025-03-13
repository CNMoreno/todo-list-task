package domain

// User represents a user entity in the system.
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserRequest represents the incoming data structure for user registration and login.
type UserRequest struct {
	Username string `json:"username" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required,min=8" validate:"required,min=8"`
}

// UserResponse represents the outgoing data structure after successful authentication.
type UserResponse struct {
	Token string `json:"token"`
}
