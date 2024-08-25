package main

import (
	"api-culinary-review/config"
	"api-culinary-review/docs"
	"api-culinary-review/internal/routes"
	"api-culinary-review/pkg/database"
	"api-culinary-review/pkg/helper"
	"log"
)

// @title API Culinary Review
// @version 1.0
// @description This is a sample server for culinary review API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host screeching-joanna-arasycorp-919c2cee.koyeb.app
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.LoadConfig()
	db := database.ConnectDB(*cfg)

	environment := helper.Getenv("ENVIRONMENT", "development")

	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Culinary Review REST API"
	docs.SwaggerInfo.Description = "This is REST API Culinary Review."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = helper.Getenv("HOST", "screeching-joanna-arasycorp-919c2cee.koyeb.app")
	if environment == "development" {
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	r := routes.SetupRouter(db)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
