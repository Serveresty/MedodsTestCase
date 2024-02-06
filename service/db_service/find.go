package dbservice

import (
	"MedodsTestCase/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUser(guid string, tokenHash string, coll *mongo.Collection) (bool, *models.User, error) {
	findFilter := bson.M{"guid": guid}

	var user models.User

	err := coll.FindOne(context.Background(), findFilter).Decode(&user)
	isNewUser := err == mongo.ErrNoDocuments
	if err != nil && !isNewUser {
		return false, nil, err
	}
	return isNewUser, &user, nil
}
