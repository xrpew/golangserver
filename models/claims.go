package models

import "github.com/golang-jwt/jwt"

type AppClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
