package app

import (
	"Ctrl/internal/database"
	userService2 "Ctrl/internal/services"
	"Ctrl/internal/transport/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func Run() {
	data, err := database.InitDB()
	if err != nil {
		log.Fatalf("database dead %v", err)
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	taskRepo := database.NewTaskRepository(data)
	taskService := userService2.NewTaskService(taskRepo)

	userRepo := database.NewUserRepository(data)
	userServices := userService2.NewUserService(userRepo, taskService)

	userHandler := http.NewUserHandler(userServices)
	taskHandler := http.NewTaskHandler(taskService)

	combinedHandler := &http.CombinedHandler{
		UserHandlerService: userHandler,
		TaskHandlerService: taskHandler,
	}

	strictCombined := http.NewStrictHandler(combinedHandler, nil)
	http.RegisterHandlers(e, strictCombined)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("server dead %v", err)
	}
}
