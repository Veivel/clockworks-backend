package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte(os.Getenv("KEY_STR"))

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
