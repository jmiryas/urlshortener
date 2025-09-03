package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"github.com/jmiryas/urlshortener/config"
)

func Protected(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	
	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Authorization header required"})
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.Get("JWT_SECRET", "SALINGJAGA"), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Locals("user_id", claims["user_id"])
	c.Locals("username", claims["username"])

	return c.Next()
}