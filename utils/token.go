package utils

import (
	"github.com/dgrijalva/jwt-go"
	"stock_prediction_backend/config"
	"stock_prediction_backend/models"
	"time"
)

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}
