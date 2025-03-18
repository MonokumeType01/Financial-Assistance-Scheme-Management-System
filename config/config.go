package config

import (
	"fmt"
	"log"
	"os"

	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	database.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	// AutoMigrate updated models
	database.AutoMigrate(&models.Applicant{}, &models.HouseholdMember{})

	DB = database
	log.Println("Database connected successfully!")
}
