package user_handler

import (
	"errors"
	"log"
	"net/http"

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
		return c.Status(http.StatusBadRequest).JSON(util.EncodeErrorResponse(err))
	}

	res, err := h.Service.CreateUser(c.Context(), &u)
	if err != nil {
		switch {
		case errors.Is(err, util.ErrEmailExists):
			return c.Status(http.StatusConflict).JSON(util.EncodeErrorResponse(err))
		case errors.Is(err, util.ErrEmptyFields):
			return c.Status(http.StatusBadRequest).JSON(util.EncodeErrorResponse(err))
		default:
			log.Printf("error creating user %v", err)
			return c.Status(http.StatusInternalServerError).JSON(util.EncodeErrorResponse(err))
		}
	}

	return c.Status(http.StatusOK).JSON(util.EncodeSuccessResponse(res))
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var user models.LoginUserReq

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(util.EncodeErrorResponse(err))
	}

	u, err := h.Service.Login(c.Context(), &user)
	if err != nil {
		switch {
		case errors.Is(err, util.ErrInvalidCredentials):
			return c.Status(http.StatusUnauthorized).JSON(util.EncodeErrorResponse(err))
		default:
			log.Printf("error logging user %v", err)
			return c.Status(http.StatusInternalServerError).JSON(util.EncodeErrorResponse(err))
		}
	}

	return c.Status(http.StatusOK).JSON(util.EncodeSuccessResponse(u))
}
