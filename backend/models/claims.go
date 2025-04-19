package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}
