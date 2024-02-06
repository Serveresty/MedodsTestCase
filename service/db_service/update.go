package dbservice

import (
	"MedodsTestCase/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUser(user *models.User, tokenHash string, coll *mongo.Collection) error {
	user.RefreshTokens = append(user.RefreshTokens, tokenHash)

	updateFilter := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "refreshtokens", Value: user.RefreshTokens}}}}

	findFilter := bson.M{"guid": user.Guid}
	_, err := coll.UpdateOne(context.Background(), findFilter, updateFilter)
	if err != nil {
		return err
	}

	return nil
}
