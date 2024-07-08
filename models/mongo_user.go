package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"stock_prediction_backend/config"
	"time"
)

func GetUserByUsername(username string) (*User, error) {
	var user User
	collection := config.DB.Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *User) error {
	collection := config.DB.Collection("users")
	user.ID = primitive.NewObjectID().Hex()
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func GetAllUserIDs() ([]string, error) {
	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetProjection(bson.M{"_id": 1}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var userIDs []string
	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		userIDs = append(userIDs, user.ID)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return userIDs, nil
}
