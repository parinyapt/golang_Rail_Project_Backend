package models

import (
	"github.com/golang-jwt/jwt/v4"
)

type JWTCustomClaims struct {
	VerifyID         string `json:"verify_id"`
	jwt.RegisteredClaims
}
