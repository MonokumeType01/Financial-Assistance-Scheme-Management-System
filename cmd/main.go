package main

import (
	"log"
	"os"

	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/config"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/handlers"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/routes"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	router := gin.Default()

	// Services & Handlers
	applicantService := services.NewApplicantService(config.DB)
	applicantHandler := handlers.NewApplicantHandler(applicantService)

	// Routes
	routes.SetupRoutes(router, applicantHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	router.Run(":" + port)
}
