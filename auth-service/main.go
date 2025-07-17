package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xavierspacelix/logistic-system/auth-service/config"
	_ "github.com/xavierspacelix/logistic-system/auth-service/docs"
	"github.com/xavierspacelix/logistic-system/auth-service/model"
	"github.com/xavierspacelix/logistic-system/auth-service/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

// @title Auth Service API
// @version 1.0
// @description This is a service for user authentication.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8001
// @BasePath /api/auth

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.InitDB()
	config.DB.AutoMigrate(&model.User{})

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8001") // default port
}
