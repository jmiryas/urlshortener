package storage

import (
	"fmt"
	"log"

	"github.com/jmiryas/urlshortener/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dsn := buildDSN()

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(getLogLevel()),
	})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connected successfully")
}

func buildDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Get("DB_HOST", "localhost"),
		config.Get("DB_USER", "postgres"),
		config.Get("DB_PASSWORD", "postgres"),
		config.Get("DB_NAME", "urlshortener_db"),
		config.Get("DB_PORT", "5432"),
		config.GetSSLMode(),
	)
}

func getLogLevel() logger.LogLevel {
	if config.IsProduction() {
		return logger.Warn
	}
	return logger.Info
}