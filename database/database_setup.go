package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var Client *mongo.Client
var UserCollection *mongo.Collection
var ProdCollection *mongo.Collection

func init() {
	var err error
	Client, err = DBSet() // DBSet connects and pings, return *mongo.Client
	ProdCollection = OpenCollection(Client, "Products")
	UserCollection = OpenCollection(Client, "Users")
	if err != nil {
		log.Fatalf("Fatal: Could not establish MongoDB connection: %v", err)
	}
}

func DBSet() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB primary: %w", err)
	}
	fmt.Println("Successfully connected and pinged MongoDB!")
	return client, nil
}

// OpenCollection is the helper that gets a specific *mongo.Collection.
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("Ecommerce").Collection(collectionName)
}
