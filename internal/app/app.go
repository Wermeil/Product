package app

import (
	"Ctrl/internal/config"
	"Ctrl/internal/database"
	"Ctrl/internal/kafka"
	Service "Ctrl/internal/services"
	rout "Ctrl/internal/transport/http"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}
	cfg := config.Load()
	kafkaProducer := kafka.NewProducer(cfg.Kafka.Broker)
	data, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("database dead %v", err)
	}
	redisClient, err := database.NewRedisClient(&cfg.Redis)
	if err != nil {
		log.Fatal("Redis warning: %v", err)
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	taskRepo := database.NewTaskRepository(data)
	taskService := Service.NewTaskService(taskRepo)

	userRepo := database.NewUserRepository(data)
	userServices := Service.NewUserService(userRepo, taskService, redisClient, kafkaProducer)

	userHandler := rout.NewUserHandler(userServices)
	taskHandler := rout.NewTaskHandler(taskService)

	combinedHandler := &rout.CombinedHandler{
		UserHandlerService: userHandler,
		TaskHandlerService: taskHandler,
	}

	strictCombined := rout.NewStrictHandler(combinedHandler, nil)
	rout.RegisterHandlers(e, strictCombined)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("server dead %v", err)
	}
}
