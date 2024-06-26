package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"myapp/config"
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
