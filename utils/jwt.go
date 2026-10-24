package utils

import (
	"time"

	"github.com/davadinata/project-management/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateJWT(roleID int64, userID int64, publicID uuid.UUID) (string, error) {
	secret := config.AppConfig.JWTSecret

	duration, _ := time.ParseDuration(config.AppConfig.JWTExpire)

	claim := jwt.MapClaims{
		"user_id":   userID,
		"role_id":   roleID,
		"public_id": publicID,
		"exp":       time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(secret))
}
