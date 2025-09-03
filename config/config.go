package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

func Get(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func IsProduction() bool {
	return Get("APP_ENV", "development") == "production"
}

func GetSSLMode() string {
	if IsProduction() {
		return "require"
	}
	
	return "disable"
}