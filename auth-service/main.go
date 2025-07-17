package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/xavierspacelix/logistic-system/auth-service/config"
	"github.com/xavierspacelix/logistic-system/auth-service/model"
	"github.com/xavierspacelix/logistic-system/auth-service/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
}
func main() {
	config.InitDB()
	config.DB.AutoMigrate(&model.User{})

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8001") // default port
}
