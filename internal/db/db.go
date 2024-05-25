package db

import (
	"log"
	"plog-backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "user=imaikosuke password=postgresql0202 dbname=plog sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	// マイグレーションを自動的に実行
	err = db.AutoMigrate(&models.User{}, &models.Photolog{}, &models.Image{}, &models.Comment{})
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	log.Println("Successfully connected to the database")
}
