package db

import (
	"context"

	"github.com/carlosescorche/usergo/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserCheckEmail Checks if the given email is registered in the database
func UserCheckEmail(client *mongo.Client, ctx context.Context, email string, id string) (types.User, bool) {

	db := SetDatabase(client)
	col := db.Collection("users")

	filter := bson.M{"email": email}

	if len(id) > 0 {
		ID, _ := primitive.ObjectIDFromHex(id)
		filter["_id"] = bson.M{"$ne": ID}
	}

	var result types.User
	err := col.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		return result, false
	}

	return result, true
}
