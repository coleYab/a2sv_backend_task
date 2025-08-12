// Package utils: utitlity module
package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

// ConnectMongo connects to MongoDB and returns the client instance.
func ConnectMongo(uri string) (*mongo.Client, error) {
    if mongoClient != nil {
        return mongoClient, nil
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI(uri)

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
    }

    // Ping to ensure connection is established
    if err := client.Ping(ctx, nil); err != nil {
        return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
    }

    log.Println("Connected to MongoDB")
    mongoClient = client
    return client, nil
}

func CreateCollection(db *mongo.Database, collectionName string) *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := db.CreateCollection(ctx, collectionName)
	if err != nil {
		// If the collection already exists, we still return it
		commandErr, ok := err.(mongo.CommandError)
		if ok && commandErr.Code == 48 {
			log.Printf("Collection '%s' already exists\n", collectionName)
			return db.Collection(collectionName)
		}
		log.Fatalf("Failed to creaet a collection %v\n", err.Error())
	}

	log.Printf("Collection '%s' created successfully\n", collectionName)
	return db.Collection(collectionName)
}

