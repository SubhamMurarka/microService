package user_handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/SubhamMurarka/microService/Users/models"
	"github.com/SubhamMurarka/microService/Users/user_service"
	"github.com/SubhamMurarka/microService/Users/util"
	"github.com/gofiber/fiber/v2"
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

func Authenticate(c *fiber.Ctx) error {
	clientToken := c.Get("token")
	if clientToken == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "No Authorisation header provided"})
	}

	claims, err := util.ValidateToken(clientToken)
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Token is Expired"})
		} else {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	c.Locals("email", claims.Email)
	c.Locals("username", claims.Username)
	c.Locals("userid", claims.UserID)
	return c.Next()
}
