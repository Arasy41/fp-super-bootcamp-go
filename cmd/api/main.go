package main

import (
	"api-culinary-review/config"
	"api-culinary-review/internal/middlewares"
	"api-culinary-review/internal/routes"
	"api-culinary-review/pkg/database"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	db := database.ConnectDB(*cfg)

	r := routes.SetupRouter(db)

	r.Use(middlewares.CORSMiddleware())

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
