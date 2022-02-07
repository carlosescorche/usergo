package db

import (
	"context"

	"github.com/carlosescorche/usergo/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserAdd records an user in the database
func UserAdd(client *mongo.Client, ctx context.Context, u types.User) (string, error) {
	db := SetDatabase(client)
	col := db.Collection("users")

	u.Password, _ = PasswordEncrypt(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", err
	}

	id, _ := result.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}
