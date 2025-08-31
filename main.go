package main

import (
	"log"

	"github.com/jmiryas/urlshortener/config"
	"github.com/jmiryas/urlshortener/routes"
	"github.com/jmiryas/urlshortener/storage"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	storage.InitDB(cfg)

	// Setup routes
	app := routes.SetupRoutes(cfg)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}