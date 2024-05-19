package util

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/SubhamMurarka/microService/Users/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmptyFields        = errors.New("email and username cannot be empty")
	ErrEmailExists        = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrTryLoginAgain      = errors.New("try to login again")
)

type tokenCreateParams struct {
	Username string
	Email    string
	UserID   string
	jwt.StandardClaims
}

var SECRET_KEY string = config.Config.JwtSecret

func GenerateAllTokens(userID string, username string, email string) (signedToken string, err error) {
	claims := &tokenCreateParams{
		Username: username,
		Email:    email,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Minute * time.Duration(10)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, err
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password %w", err)
	}

	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func EncodeSuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{"data": data}
}

func EncodeErrorResponse(err error) fiber.Map {
	return fiber.Map{
		"error": err.Error(),
	}
}
