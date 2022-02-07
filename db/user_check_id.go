package db

import (
	"context"

	"github.com/carlosescorche/usergo/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserCheckId checks if the given id is registered in the database
func UserCheckId(client *mongo.Client, ctx context.Context, id string) (types.User, bool) {
	var result types.User

	db := SetDatabase(client)
	col := db.Collection("users")

	ID, _ := primitive.ObjectIDFromHex(id)
	err := col.FindOne(ctx, bson.M{"_id": ID}).Decode(&result)

	if err != nil {
		return result, false
	}

	return result, true
}
