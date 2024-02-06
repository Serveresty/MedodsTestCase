package jwtservice

import (
	"MedodsTestCase/env"
	"MedodsTestCase/models"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(token string) (*models.Claims, error) {
	secretKey := []byte(env.GetEnv("SECRET_KEY"))

	parsedToken, err := jwt.ParseWithClaims(token, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Wrong algorithm: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*models.Claims)
	if !ok {
		return nil, errors.New("Error segment of token")
	}

	if !parsedToken.Valid {
		return claims, errors.New("Token is not valid")
	}

	return claims, nil
}
