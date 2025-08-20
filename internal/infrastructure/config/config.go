package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

// NewConfig loads environment variables and returns a Config struct.
func NewConfig() *Config {
	// In Node.js, this is like calling require('dotenv').config()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}