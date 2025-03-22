package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	applicantService, schemeService := initializeServices()
	applicantHandler := handlers.NewApplicantHandler(applicantService)
	schemeHandler := handlers.NewSchemeHandler(schemeService)

	// Routes
	routes.SetupRoutes(router, applicantHandler, schemeHandler)

	srv := &http.Server{
		Addr:    ":" + getPort(),
		Handler: router,
	}

	//run server
	go func() {
		log.Printf("Server is running on port %s", getPort())
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	shutdown(srv)
}

func initializeServices() (*services.ApplicantService, *services.SchemeService) {
	applicantService := services.NewApplicantService(config.DB)
	schemeService := services.NewSchemeService(config.DB)
	return applicantService, schemeService
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func shutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited.")
}
