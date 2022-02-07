package db

import (
	"context"

	"github.com/carlosescorche/usergo/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserUpdate allows to update a registered user
func UserUpdate(client *mongo.Client, ctx context.Context, id string, u types.User) error {
	db := SetDatabase(client)
	col := db.Collection("users")

	ID, _ := primitive.ObjectIDFromHex(id)

	_, err := col.UpdateOne(ctx, bson.M{"_id": ID}, bson.M{"$set": u})

	return err
}
