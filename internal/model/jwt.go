package model

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	Username string `json:"username"`
	Hostname string `json:"hostname"`
	jwt.RegisteredClaims
}
