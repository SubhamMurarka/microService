package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Authorise(c *fiber.Ctx) error {
	token := c.Get("token")

	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "No token present"})
	}

	// TODO to implement authorisation function
}
