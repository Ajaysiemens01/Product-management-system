package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

var APIKey string
var PORT string

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	APIKey = os.Getenv("API_KEY")
	PORT = os.Getenv("PORT")
}
