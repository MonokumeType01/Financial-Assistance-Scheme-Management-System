package config

import (
	"log"
	"os"

	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB_DSN")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	database.AutoMigrate(&models.Applicant{}, &models.HouseholdMember{})

	DB = database
	log.Println("✅ Database connected successfully!")
}
