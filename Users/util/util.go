package util

import (
	"fmt"
	"log"
	"time"

	"github.com/SubhamMurarka/microService/Users/config"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
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
