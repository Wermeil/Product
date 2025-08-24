package database

import (
	"Ctrl/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.GetDBDSN()
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database error %v", err)
	}
	return db, nil
}
