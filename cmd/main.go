package main

import (
	"log"
	"os"

	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/config"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	router := gin.Default()
	routes.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	router.Run(":" + port)
}
