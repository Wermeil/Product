package database

import (
	"Ctrl/internal/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// RedisClient обертка вокруг redis.Client для лучшего контроля
type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(cfg *config.RedisConfig) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}
	return &RedisClient{client: client}, nil
}
