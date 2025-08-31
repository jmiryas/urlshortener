package main

import (
	"log"

	"github.com/jmiryas/urlshortener/config"
	"github.com/jmiryas/urlshortener/routes"
	"github.com/jmiryas/urlshortener/storage"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database
	storage.InitDB()

	// Setup routes
	app := routes.SetupRoutes()

	// Start server
	port := config.Get("PORT", "3000")
	
	storage.RunMigrations()

	log.Printf("Server starting on port %s", port)

	log.Fatal(app.Listen(":" + port))
}