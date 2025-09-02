package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/jmiryas/urlshortener/handlers"
	"github.com/jmiryas/urlshortener/middleware"
)

func SetupRoutes() *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	
	// Middleware
	app.Use(middleware.Logger())
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Hello!")
	})

	v1 := app.Group("/api/v1")

	v1.Post("/shorten", handlers.ShortenURL)
	v1.Get("/:token", handlers.RedirectURL)
	v1.Get("/stats/:token", handlers.GetStats)
	
	v1.Get("/analytics/:token", handlers.GetAnalytics)

	v1.Post("/auth/register", handlers.Register)
	v1.Post("/auth/login", handlers.Login)
	
	return app
}