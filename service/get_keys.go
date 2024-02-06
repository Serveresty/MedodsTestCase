package service

import (
	jwtservice "MedodsTestCase/service/jwt_service"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Router struct {
	*mongo.Database
}

func (db *Router) GetKeys(w http.ResponseWriter, r *http.Request) {
	guid := r.URL.Query().Get("guid")

	signedKey := jwtservice.GenerateSeed()

	jwtToken, aExpTime, err := jwtservice.GenerateToken(guid, signedKey)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	refreshToken, rExpTime, err := jwtservice.GenerateToken(guid, signedKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rTokenHash, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access-token",
		Value:    jwtToken,
		Expires:  aExpTime,
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh-token",
		Value:    refreshToken,
		Expires:  rExpTime,
		HttpOnly: true,
	})
}
