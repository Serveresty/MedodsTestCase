package service

import (
	dbservice "MedodsTestCase/service/db_service"
	jwtservice "MedodsTestCase/service/jwt_service"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (db *DBCollection) RefreshKey(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh-token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rTokenStr := cookie.Value
	rClaims, err := jwtservice.ParseToken(rTokenStr)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	cookie, err = r.Cookie("access-token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	aTokenStr := cookie.Value
	aClaims, err := jwtservice.ParseToken(aTokenStr)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if rClaims.Guid != aClaims.Guid || rClaims.Id != aClaims.Id {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, user, err := dbservice.FindUser(rClaims.Guid, db.Collection)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var isExists bool
	for _, hash := range user.RefreshTokens {
		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(rTokenStr))
		if err == nil {
			isExists = true
			break
		}
	}
	if !isExists {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jwtToken, aExpTime, err := jwtservice.GenerateToken(user.Guid, rClaims.Id)
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
}
