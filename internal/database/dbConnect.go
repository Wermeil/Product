package database

import (
	"Ctrl/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.GetDBDSN()
	var db *gorm.DB
	var err error

	// Пытаемся подключиться несколько раз
	for i := 0; i < 3; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Database connected successfully")
			return db, nil
		}

		log.Printf("Attempt %d: Database connection failed: %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	return nil, err
}
