package db

import (
	"blog-api/internal/models"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB(databaseURL string) *gorm.DB {

	database, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the models
	if err := database.AutoMigrate(&models.User{}, &models.Blog{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully")
	return database

}