package tests

import (
	"testing"

	sqlite "github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/jmiryas/urlshortener/models"
	"github.com/jmiryas/urlshortener/routes"
	"github.com/jmiryas/urlshortener/storage"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := db.AutoMigrate(&models.URL{}, &models.Visit{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	storage.DB = db
	return db
}

func TeardownTestDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	_ = sqlDB.Close()
	storage.DB = nil
}

func SetupApp() *fiber.App {
	return routes.SetupRoutes()
}
