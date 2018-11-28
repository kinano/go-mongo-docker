package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// GetDB gets the mongo database for our app
func GetDB() (*mongo.Database, error) {
	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		os.Getenv("MONGO_USERNAME"),
		os.Getenv("MONGO_PASSWORD"),
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_DB_NAME"),
	)
	client, err := mongo.NewClient(uri)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to Mongo DB: %v", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Mongo client couldn't connect with background context: %v", err)
	}
	return client.Database(os.Getenv("MONGO_DB_NAME")), nil
}

// GetByObjectID fetches a document by its id
func GetByObjectID(objectID string, c *mongo.Collection) (interface{}, error) {
	var result interface{}
	filter := bson.D{{"id", objectID}}
	if err := c.FindOne(context.Background(), filter).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
