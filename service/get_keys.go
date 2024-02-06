package service

import (
	dbservice "MedodsTestCase/service/db_service"
	jwtservice "MedodsTestCase/service/jwt_service"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Получение токенов
func (db *DBCollection) GetKeys(w http.ResponseWriter, r *http.Request) {
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
	rTokenHashStr := string(rTokenHash)

	isNewUser, user, err := dbservice.FindUser(guid, db.Collection)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isNewUser {
		err := dbservice.InsertNewUser(guid, rTokenHashStr, db.Collection)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		err := dbservice.UpdateUser(user, rTokenHashStr, db.Collection)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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
