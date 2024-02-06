package service

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type DBCollection struct {
	*mongo.Collection
}

func Routs(router *http.ServeMux, db *mongo.Collection) {
	dtb := DBCollection{db}
	router.HandleFunc("/receive-keys", dtb.GetKeys)
	router.HandleFunc("/refresh-key", dtb.RefreshKey)
}
