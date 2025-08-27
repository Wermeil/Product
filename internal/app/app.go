package app

import (
	"Ctrl/internal/config"
	"Ctrl/internal/database"
	Service "Ctrl/internal/services"
	"Ctrl/internal/transport/http"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"time"
)

func Run() {
	time.Sleep(3 * time.Second) // Ждем 3 секунды
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}
	cfg := config.Load()
	data, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("database dead %v", err)
	}
	redisClient, err := database.NewRedisClient(&cfg.Redis)
	if err != nil {
		log.Printf("Redis warning: %v (continuing without cache)", err)
		// Можно продолжить без Redis, если он не критичен
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	taskRepo := database.NewTaskRepository(data)
	taskService := Service.NewTaskService(taskRepo)

	userRepo := database.NewUserRepository(data)
	userServices := Service.NewUserService(userRepo, taskService, redisClient)

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
