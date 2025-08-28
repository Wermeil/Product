package config

import (
	"os"
	"strconv"
)

type Config struct {
	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	Redis      RedisConfig
	Kafka      KafkaConfig
}
type RedisConfig struct {
	Addr     string // Адрес, например: "localhost:6379"
	Password string // Пароль (если есть)
	DB       int    // Номер базы данных (по умолчанию 0)
}
type KafkaConfig struct {
	Broker string // например: "localhost:9092"
	Topic  string // например: "user-events"
}

func Load() *Config {
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB")) // Преобразуем строку в int
	if err != nil {
		return nil
	}
	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
		Redis: RedisConfig{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       redisDB,
		},
		Kafka: KafkaConfig{
			Broker: os.Getenv("KAFKA_BROKER"),
			Topic:  os.Getenv("KAFKA_TOPIC"),
		},
	}
}
func (c *Config) GetDBDSN() string {
	return "host=" + c.DBHost +
		" user=" + c.DBUser +
		" password=" + c.DBPassword +
		" dbname=" + c.DBName +
		" port=" + c.DBPort +
		" sslmode=" + c.DBSSLMode
}
