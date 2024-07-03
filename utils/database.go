package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"myapp/models"
)

var mongoURI = "mongodb://localhost:27017"
var dbName = "financial_guru"
var collectionName = "user_responses"

var client *mongo.Client

func ConnectDB() (*mongo.Client, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return nil, err
	}

	return client, nil
}

func GetClient() (*mongo.Client, error) {
	if client != nil {
		return client, nil
	}
	return ConnectDB()
}

func GetUserResponses(userID string) (*models.UserResponses, error) {
	client, err := GetClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)
	var userResponses models.UserResponses
	err = collection.FindOne(context.Background(), bson.M{"user_id": userID}).Decode(&userResponses)
	if err != nil {
		log.Printf("Error fetching user responses for user ID %s: %v", userID, err)
		return nil, err
	}

	return &userResponses, nil
}

func SaveUserResponses(userResponses *models.UserResponses) error {
	client, err := GetClient()
	if err != nil {
		return err
	}

	collection := client.Database(dbName).Collection(collectionName)
	filter := bson.M{"user_id": userResponses.UserID}
	update := bson.M{"$set": userResponses}

	opts := options.Update().SetUpsert(true)
	_, err = collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		log.Printf("Error saving user responses for user ID %s: %v", userResponses.UserID, err)
		return err
	}

	return nil
}
