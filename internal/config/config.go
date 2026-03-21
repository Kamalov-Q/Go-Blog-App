package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DatabaseURL string
}

func LoadConfig() *Config {

	// Load .env file (ignor error if not found)
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set properly")
	}

	return &Config{
		Port: port,
		DatabaseURL: databaseURL,
	}
}