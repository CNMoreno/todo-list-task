package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-list-task/internal/app"
	handlerHttp "todo-list-task/internal/infrastructure/http"
	"todo-list-task/internal/infrastructure/memory"
	"todo-list-task/internal/middleware"
	"todo-list-task/internal/utils"
)

func main() {

	taskRepo := memory.NewInMemoryTaskRepository()
	taskService := app.NewTaskService(taskRepo)
	taskHandler := handlerHttp.NewTaskHandler(taskService)

	bcryptCrypto := app.BcryptCrypto{}
	appCrypto := utils.NewHashPassword(bcryptCrypto)
	userRepo := memory.NewInMemoryUserRepository(appCrypto)

	userService := app.NewUserService(userRepo)
	userHandler := handlerHttp.NewUserHandler(userService)

	r := gin.Default()

	r.POST("/users", userHandler.RegisterUser)
	r.POST("/login", userHandler.LoginUser)

	r.POST("/tasks", middleware.AuthMiddleware(), taskHandler.RegisterTask)
	r.GET("/tasks/:id", middleware.AuthMiddleware(), taskHandler.GetTaskByID)
	r.GET("/tasks", middleware.AuthMiddleware(), taskHandler.GetAllTask)
	r.PUT("/tasks/:id", middleware.AuthMiddleware(), taskHandler.UpdateTask)
	r.DELETE("tasks/:id", middleware.AuthMiddleware(), taskHandler.DeleteTask)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Recibida señal de cierre, apagando servidor...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Error al apagar el servidor: %v", err)
		}
		log.Println("Servidor apagado correctamente")
	}()

	log.Println("Servidor iniciado en el puerto 8080")
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Error en el servidor: %v", err)
	}

	log.Println("Salida limpia del programa")
}
