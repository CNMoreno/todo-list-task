package domain

// Task represents a to-do item in the system.
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// TaskRequest represents the incoming data structure for creating or updating a task.
type TaskRequest struct {
	Title       string `json:"title" binding:"required" validate:"required"`
	Description string `json:"description" binding:"required" validate:"required"`
	Completed   bool   `json:"completed"`
}

