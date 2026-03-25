package db

import (
	"log"

	"urlshortener/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dbURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the database schemas
	if err := db.AutoMigrate(&model.Shortlink{}); err != nil {
		log.Printf("Failed to auto migrate database schemas: %v", err)
		return nil, err
	}

	return db, nil
}
