package database

import (
	"api-culinary-review/config"
	// "api-culinary-review/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate models
	// err = db.AutoMigrate(
	// 	&models.User{},
	// 	&models.Profile{},
	// 	&models.Recipe{},
	// 	&models.Review{},
	// 	&models.Image{},
	// 	&models.Tag{},
	// 	&models.Favorite{},
	// )

	// if err != nil {
	// 	log.Fatalf("Failed to migrate database: %v", err)
	// }

	return db
}
