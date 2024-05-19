package auth

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/SubhamMurarka/microService/Products/config"
	"github.com/SubhamMurarka/microService/Products/models"
	"github.com/SubhamMurarka/microService/Products/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var SECRET_KEY string = config.Config.JwtSecret

func Authorise(c *fiber.Ctx) error {
	token := c.Get("token")
	if token == "" {
		return c.Status(http.StatusBadRequest).JSON(utils.EncodeErrorResponse(utils.ErrEmptyToken))
	}
	fmt.Println(token)
	claims, err := ValidateToken(token)
	if err != nil {
		fmt.Println("error occured: ", err)
		switch {
		case errors.Is(err, utils.ErrExpiredToken):
			return c.Status(http.StatusUnauthorized).JSON(utils.EncodeErrorResponse(err))
		default:
			return c.Status(http.StatusInternalServerError).JSON(utils.EncodeErrorResponse(err))
		}
	}

	c.Locals("ID", claims.UserID)
	c.Locals("Username", claims.Username)
	c.Locals("Email", claims.Email)

	return c.Next()
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
		return nil, utils.ErrInvalidToken
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, utils.ErrExpiredToken
	}

	return claims, nil
}
