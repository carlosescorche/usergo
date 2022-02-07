package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserDelete allows to delete users from the database
func UserDelete(client *mongo.Client, ctx context.Context, id string) error {
	db := SetDatabase(client)
	col := db.Collection("users")

	ID, _ := primitive.ObjectIDFromHex(id)

	_, err := col.DeleteOne(ctx, bson.M{
		"_id": ID,
	})

	return err
}
