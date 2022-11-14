package models

import (
	"github.com/golang-jwt/jwt/v4"
)

type JWTCustomClaims struct {
	jwt.StandardClaims
	VerifyID string `json:"verify_id"`
}