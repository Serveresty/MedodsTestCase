package service

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func Routs(router *http.ServeMux, db *mongo.Database) {
	dtb := Router{db}
	router.HandleFunc("/receive-keys", dtb.GetKeys)
	router.HandleFunc("/refresh-key", dtb.RefreshKey)
}
