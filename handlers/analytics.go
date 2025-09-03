package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmiryas/urlshortener/models"
	"github.com/jmiryas/urlshortener/storage"
)

func GetAnalytics(c *fiber.Ctx) error {
	token := c.Params("token")

	var url models.URL

	if err := storage.DB.Where("short_token = ?", token).First(&url).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "URL not found"})
	}

	var totalClicks int64

	storage.DB.Model(&models.Visit{}).Where("url_id = ?", url.ID).Count(&totalClicks)

	var uniqueVisitors int64

	storage.DB.Model(&models.Visit{}).
		Where("url_id = ?", url.ID).
		Distinct("ip_address").
		Count(&uniqueVisitors)

	var referrerStats []struct {
		Referrer string `json:"referrer"`
		Count    int64  `json:"count"`
	}
	
	storage.DB.Model(&models.Visit{}).
		Select("referrer, count(*) as count").
		Where("url_id = ? AND referrer != ''", url.ID).
		Group("referrer").
		Order("count DESC").
		Scan(&referrerStats)

	return c.JSON(fiber.Map{
		"url": fiber.Map{
			"original_url": url.OriginalURL,
			"short_token":  url.ShortToken,
			"click_count":  url.ClickCount,
		},
		"total_clicks":    totalClicks,
		"unique_visitors": uniqueVisitors,
		"referrers":       referrerStats,
	})
}