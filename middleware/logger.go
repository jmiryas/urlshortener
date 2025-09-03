package middleware

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		
		os.MkdirAll("logs", 0755)
		
		today := time.Now().Format("2006-01-02")
		logFile := filepath.Join("logs", today+".log")
		
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Failed to open log file: %v\n", err)
			return c.Next()
		}
		defer file.Close()
		
		var requestBody interface{}
		if c.Request().Body() != nil && len(c.Request().Body()) > 0 {
			json.Unmarshal(c.Request().Body(), &requestBody)
		}
		
		requestLog := map[string]interface{}{
			"time":       time.Now().Format(time.RFC3339),
			"type":       "REQUEST",
			"method":     c.Method(),
			"path":       c.Path(),
			"ip":         c.IP(),
			"user_agent": c.Get("User-Agent"),
			"body":       requestBody,
		}
		requestJSON, _ := json.MarshalIndent(requestLog, "", "  ")
		file.WriteString(string(requestJSON) + "\n")
		
		err = c.Next()
		
		var responseBody interface{}
		if c.Response().Body() != nil && len(c.Response().Body()) > 0 {
			json.Unmarshal(c.Response().Body(), &responseBody)
		}
		
		responseLog := map[string]interface{}{
			"time":     time.Now().Format(time.RFC3339),
			"type":     "RESPONSE",
			"method":   c.Method(),
			"path":     c.Path(),
			"status":   c.Response().StatusCode(),
			"duration": time.Since(start).String(),
			"ip":       c.IP(),
			"body":     responseBody,
		}
		responseJSON, _ := json.MarshalIndent(responseLog, "", "  ")
		file.WriteString(string(responseJSON) + "\n")
		
		return err
	}
}