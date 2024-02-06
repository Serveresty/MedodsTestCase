package jwtservice

import (
	"MedodsTestCase/env"
	"MedodsTestCase/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(guid string, signedKey string) (string, time.Time, error) {
	exTime := time.Now().Add(2 * time.Hour)

	claims := &models.Claims{
		Guid: guid,
		StandardClaims: jwt.StandardClaims{
			Id:        signedKey,
			ExpiresAt: exTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(env.GetEnv("SECRET_KEY")))
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, exTime, nil
}
