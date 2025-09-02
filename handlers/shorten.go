package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jmiryas/urlshortener/config"
	"github.com/jmiryas/urlshortener/models"
	"github.com/jmiryas/urlshortener/storage"
	"github.com/jmiryas/urlshortener/utils"
)

type ShortenRequest struct {
	URL string `json:"url" validate:"required,url"`
}

func ShortenURL(c *fiber.Ctx) error {
	var req ShortenRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	req.URL = strings.TrimSpace(req.URL)

	if req.URL == "" {
		return c.Status(400).JSON(fiber.Map{"error": "URL is required"})
	}

	// if _, err := url.ParseRequestURI(req.URL); err != nil {
	// 	return c.Status(422).JSON(fiber.Map{"error": "Invalid URL format"})
	// }

	if !utils.IsValidURL(req.URL) {
		return c.Status(422).JSON(fiber.Map{"error": "Invalid URL format"})
	}

	normalizedURL := utils.NormalizeURL(req.URL)

	// Generate token using CRC32 for shorter tokens
	token := utils.GenerateToken(normalizedURL)

	// Check if URL already exists
	var existingURL models.URL

	// result := storage.DB.Where("original_url = ?", req.URL).First(&existingURL)
	result := storage.DB.Where("original_url = ?", normalizedURL).First(&existingURL)

	base_url := config.Get("BASE_URL", "http://localhost:3000/")

	if result.Error != nil {
		// Create new URL
		newURL := models.URL{
			OriginalURL: normalizedURL,
			ShortToken: token,
		}

		if err := storage.DB.Create(&newURL).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create short URL"})
		}

		return c.Status(201).JSON(fiber.Map{
			"short_token":  token,
			"shorten_url": base_url + token,
			"original_url": normalizedURL,
		})
	}

	return c.JSON(fiber.Map{
		"short_token": existingURL.ShortToken,
		"shorten_url": base_url + existingURL.ShortToken,
		"original_url": existingURL.OriginalURL,
	})
}