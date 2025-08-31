package handlers

import (
	"github.com/jmiryas/urlshortener/models"
	"github.com/jmiryas/urlshortener/storage"

	"github.com/gofiber/fiber/v2"
)

func GetStats(c *fiber.Ctx) error {
	token := c.Params("token")

	var url models.URL
	if err := storage.DB.Where("short_token = ?", token).First(&url).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "URL not found"})
	}

	return c.JSON(fiber.Map{
		"short_token":  url.ShortToken,
		"original_url": url.OriginalURL,
		"click_count":  url.ClickCount,
		"created_at":   url.CreatedAt,
	})
}