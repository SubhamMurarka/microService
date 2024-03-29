package auth

import (
	"encoding/json"
	"io"
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

	url := "http://localhost:" + port + "/auth?" + token

	res, err := http.Get(url)

	log.Fatal("error in authorisation : ", err)
	responsebyte, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("error reading the response body : ", err)
	}
	err = json.Unmarshal(responsebyte, &Authres)
	if err != nil {
		log.Fatal("error marshalling the response : ", err)
	}

	return nil
}
