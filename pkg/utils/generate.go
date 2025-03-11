package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/model"
)

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func GenerateSecret(length int) (string, error) {
	if length < 8 {
		return "", fmt.Errorf("secret length must be at least 8 characters")
	}

	secret := make([]byte, length)
	for i := range secret {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", fmt.Errorf("failed to generate secret: %v", err)
		}
		secret[i] = letterBytes[index.Int64()]
	}
	return string(secret), nil
}

func GenerateJWT(username string, hostKey string) (string, error) {
	host, _ := os.Hostname()
	claims := model.JWTClaims{
		Username: username,
		Hostname: host,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(hostKey))
}
