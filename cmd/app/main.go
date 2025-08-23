package main

import (
	"Ctrl/internal/db"
	"Ctrl/internal/handlers"
	"Ctrl/internal/tasksService"
	"Ctrl/internal/userService"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	data, err := db.InitDB()
	if err != nil {
		log.Fatalf("db dead %v", err)
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	taskRepo := tasksService.NewTaskRepository(data)
	taskService := tasksService.NewTaskService(taskRepo)

	userRepo := userService.NewUserRepository(data)
	userServices := userService.NewUserService(userRepo, taskService)

	userHandler := handlers.NewUserHandler(userServices)
	taskHandler := handlers.NewTaskHandler(taskService)

	combinedHandler := &handlers.CombinedHandler{
		UserHandlerService: userHandler,
		TaskHandlerService: taskHandler,
	}

	strictCombined := handlers.NewStrictHandler(combinedHandler, nil)
	handlers.RegisterHandlers(e, strictCombined)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("server dead %v", err)
	}
}
