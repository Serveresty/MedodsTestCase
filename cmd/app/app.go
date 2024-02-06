package main

import (
	"MedodsTestCase/env"
	"MedodsTestCase/service"
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	host := env.GetEnv("HOST")
	port := env.GetEnv("PORT")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+host+":"+port))
	if err != nil {
		return err
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Print(err)
		}
	}()

	database := client.Database("Medods")
	router := http.NewServeMux()

	service.Routs(router, database)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return err
	}

	return nil
}
