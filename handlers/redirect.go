package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmiryas/urlshortener/models"
	"github.com/jmiryas/urlshortener/storage"
)

func RedirectURL(c *fiber.Ctx) error {
	token := c.Params("token")

	var url models.URL
	if err := storage.DB.Where("short_token = ?", token).First(&url).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "URL not found"})
	}

	// Update click count
	storage.DB.Model(&url).Update("click_count", url.ClickCount + 1)

	return c.Redirect(url.OriginalURL, 302)
}