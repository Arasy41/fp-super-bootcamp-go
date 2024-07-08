package database

import (
	"api-culinary-review/config"
	"api-culinary-review/internal/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Jakarta", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.LogMode(true)

	// Auto Migrate models
	err = db.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Recipe{},
		&models.Review{},
		&models.Image{},
		&models.Tag{},
		&models.Favorite{},
	).Error

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Connected to database")
	return db
}
