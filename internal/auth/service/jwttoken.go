package service

import (
	"time"

	"github.com/RickDred/ozinse/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("secret_key")

func generateJWT(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
