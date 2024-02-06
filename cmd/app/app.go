package main

import (
	"MedodsTestCase/env"
	"MedodsTestCase/service"
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Инициализация пути к параметрам
func init() {
	if err := godotenv.Load("../../env/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	// Получение хоста и порта к mongodb
	host := env.GetEnv("HOST")
	port := env.GetEnv("PORT")

	// Коннект к mongodb
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+host+":"+port))
	if err != nil {
		return err
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Print(err)
		}
	}()

	// Создать бд + коллекцию
	collection := client.Database("Medods").Collection("users")

	// Создать хендлер запросов и передать коллекции
	router := http.NewServeMux()
	service.Routs(router, collection)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return err
	}

	return nil
}
