package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Global variables for configuration
var (
	BOTAPI             string
	DB_HOST            string
	DB_PORT            string
	DB_USER            string
	DB_PASSWORD        string
	DB_NAME            string
	LATITUDE           string
	LONGITUDE          string
	TIMEZONE           string
	TWILIGHTCONVENTION string
	ASRCONVENTION      string
	PRECISETOSECONDS   string
)

func LoadConfig() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve values from .env file
	BOTAPI = os.Getenv("BOT_API")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")

	// Additional configuration
	LATITUDE = os.Getenv("LATITUDE")
	LONGITUDE = os.Getenv("LONGITUDE") // Fixed typo
	TIMEZONE = os.Getenv("TIMEZONE")
	TWILIGHTCONVENTION = os.Getenv("TWILIGHTCONVENTION")
	ASRCONVENTION = os.Getenv("ASRCONVENTION")
	PRECISETOSECONDS = os.Getenv("PRECISETOSECONDS")
}
