package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
	LogPath  string
}

func Load() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		Port:    getEnv("PORT", "3000"),
		DBHost:  getEnv("DB_HOST", "localhost"),
		DBPort:  getEnv("DB_PORT", "5432"),
		DBUser:  getEnv("DB_USER", "postgres"),
		DBPass:  getEnv("DB_PASSWORD", "postgres"),
		DBName:  getEnv("DB_NAME", "urlshortener_db"),
		LogPath: getEnv("LOG_PATH", "logs"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}