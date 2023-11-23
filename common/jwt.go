package common

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
