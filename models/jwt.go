package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Guid string `json:"guid"`
	jwt.StandardClaims
}
