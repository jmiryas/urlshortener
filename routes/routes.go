package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmiryas/urlshortener/config"
	"github.com/jmiryas/urlshortener/handlers"
	"github.com/jmiryas/urlshortener/middleware"
)

func SetupRoutes(cfg *config.Config) *fiber.App {
	app := fiber.New()
	
	// Middleware
	app.Use(middleware.Logger())
	
	// Routes
	
	app.Get("/", func (c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello!"})
	})

	app.Post("/shorten", handlers.ShortenURL)

	app.Get("/:token", handlers.RedirectURL)
	
	app.Get("/stats/:token", handlers.GetStats)
	
	return app
}