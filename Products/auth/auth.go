package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SubhamMurarka/microService/Products/config"
	"github.com/gofiber/fiber/v2"
)

type AuthRes struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var Authres AuthRes

func Authorise(c *fiber.Ctx) error {
	token := c.Get("token")
	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "No token present"})
	}

	port := config.Config.ServerPortUser
	url := "http://localhost:" + port + "/auth?token=" + token

	res, err := http.Get(url)
	if err != nil {
		log.Printf("Error while calling authentication service: %v", err)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to authenticate token"})
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Authentication service returned status: %s", res.Status)
		return c.Status(http.StatusUnauthorized).JSON("try to login again")
	}

	if err := json.NewDecoder(res.Body).Decode(&Authres); err != nil {
		log.Printf("Error decoding authentication response: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Next()
}
