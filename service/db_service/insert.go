package dbservice

import (
	"MedodsTestCase/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Записать нового пользователя в коллекцию
func InsertNewUser(guid string, tokenHash string, coll *mongo.Collection) error {
	newRefresh := make([]string, 0, 1)
	newRefresh = append(newRefresh, tokenHash)
	newUser := models.User{
		Guid:          guid,
		RefreshTokens: newRefresh,
	}

	_, err := coll.InsertOne(context.Background(), newUser)
	if err != nil {
		return err
	}

	return nil
}
