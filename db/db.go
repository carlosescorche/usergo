package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect creates the mongo client
func Connect(ctx context.Context) *mongo.Client {
	fmt.Println(os.Getenv("MONGO_URI"))

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal("Has occurred an error creating mongo client: " + err.Error())
	}

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal("Has ocurred an error connecting with the server: " + err.Error())
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("Has ocurred an error connecting to the deployment: " + err.Error())
	}

	log.Println("Mongo connected successfully")

	return client
}

// SetDataBase sets the database for the given mongo client
func SetDatabase(client *mongo.Client) *mongo.Database {
	var db = os.Getenv("DB_NAME")
	return client.Database(db)
}
