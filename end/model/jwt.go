package model

import "github.com/dgrijalva/jwt-go"

// Claims jwt
type Claims struct {
	UserID uint
	jwt.StandardClaims
}
