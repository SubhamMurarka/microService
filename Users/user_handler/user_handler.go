package user_handler

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/SubhamMurarka/microService/Users/config"
	"github.com/SubhamMurarka/microService/Users/models"
	"github.com/SubhamMurarka/microService/Users/user_service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Handler struct {
	user_service.Service
}

func NewHandler(s user_service.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var u models.CreateUserReq
	if err := c.BodyParser(&u); err != nil {
		log.Println("Error parsing request body for Signup:", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.Service.CreateUser(c.Context(), &u)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var user models.LoginUserReq

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	u, err := h.Service.Login(c.Context(), &user)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(u)
}

var SECRET_KEY string = config.Config.JwtSecret

func Authorise(c *fiber.Ctx) error {
	token := c.Query("token")
	if token == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Token is missing"})
	}
	claims, err := ValidateToken(token)
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Token is Expired"})
		} else {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	res := models.AuthRes{
		ID:       claims.UserID,
		Username: claims.Username,
		Email:    claims.Email,
	}

	return c.Status(http.StatusOK).JSON(res)
}

func ValidateToken(signedToken string) (*models.TokenCreateParams, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.TokenCreateParams{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.TokenCreateParams)

	if !ok {
		return nil, errors.New("the token is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, errors.New("token is expired")
	}

	return claims, nil
}
