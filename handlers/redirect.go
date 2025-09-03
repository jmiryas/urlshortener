package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmiryas/urlshortener/models"
	"github.com/jmiryas/urlshortener/storage"
	"gorm.io/gorm"
)

func RedirectURL(c *fiber.Ctx) error {
	token := c.Params("token")

	var url models.URL
	if err := storage.DB.Where("short_token = ?", token).First(&url).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "URL not found"})
	}

	ipAddress := c.IP()
	userAgent := c.Get("User-Agent")
	referrer := c.Get("Referer")

	storage.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&url).Update("click_count", url.ClickCount+1).Error; err != nil {
			return err
		}

		var existingVisit models.Visit
		oneDayAgo := time.Now().Add(-24 * time.Hour)
		
		result := tx.Where("url_id = ? AND ip_address = ? AND created_at > ?", 
			url.ID, ipAddress, oneDayAgo).First(&existingVisit)

		if result.Error != nil {
			if err := tx.Model(&url).Update("unique_visits", url.UniqueVisits+1).Error; err != nil {
				return err
			}
		}

		visit := models.Visit{
			URLID:      url.ID,
			IPAddress:  ipAddress,
			UserAgent:  userAgent,
			Referrer:   referrer,
			VisitTime:  time.Now(),
		}

		if err := tx.Create(&visit).Error; err != nil {
			return err
		}

		return nil
	})

	return c.Redirect(url.OriginalURL, 302)
}