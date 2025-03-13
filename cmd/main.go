package main

import (
	"github.com/gin-gonic/gin"
	"todo-list-task/internal/app"
	"todo-list-task/internal/infrastructure/http"
	"todo-list-task/internal/infrastructure/memory"
	"todo-list-task/internal/middleware"
	"todo-list-task/internal/utils"
)

func main() {

	taskRepo := memory.NewInMemoryTaskRepository()
	taskService := app.NewTaskService(taskRepo)
	taskHandler := http.NewTaskHandler(taskService)

	bcryptCrypto := app.BcryptCrypto{}
	appCrypto := utils.NewHashPassword(bcryptCrypto)
	userRepo := memory.NewInMemoryUserRepository(appCrypto)

	userService := app.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)

	r := gin.Default()

	r.POST("/users", userHandler.RegisterUser)
	r.POST("/login", userHandler.LoginUser)

	r.POST("/tasks", middleware.AuthMiddleware(), taskHandler.RegisterTask)
	r.GET("/tasks/:id", middleware.AuthMiddleware(), taskHandler.GetTaskByID)
	r.GET("/tasks", middleware.AuthMiddleware(), taskHandler.GetAllTask)
	r.PUT("/tasks/:id", middleware.AuthMiddleware(), taskHandler.UpdateTask)
	r.DELETE("tasks/:id", middleware.AuthMiddleware(), taskHandler.DeleteTask)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
