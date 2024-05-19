package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrEmptyFields    = errors.New("invalid product")
	ErrEmptyProductID = errors.New("no product id")
	ErrPurchaseReq    = errors.New("invalid purchase request")
	ErrEmptyToken     = errors.New("token is missing")
	ErrInvalidToken   = errors.New("token is invalid")
	ErrExpiredToken   = errors.New("token is Expired")
)

func EncodeSuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{"data": data}
}

func EncodeErrorResponse(err error) fiber.Map {
	return fiber.Map{
		"error": err.Error(),
	}
}
