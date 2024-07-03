package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"myapp/config"
)

type Prediction struct {
	Date  string  `json:"date" bson:"date"`
	Value float64 `json:"value" bson:"value"`
}

type TickerPrediction struct {
	Ticker      string       `json:"ticker" bson:"ticker"`
	Predictions []Prediction `json:"predictions" bson:"predictions"`
}

type Wallet struct {
	UserID  string             `json:"user_id" bson:"user_id"`
	Tickers []TickerPrediction `json:"tickers" bson:"tickers"`
}

func AddToWallet(userID string, tickerPredictions []TickerPrediction) error {
	collection := config.DB.Collection("wallets")

	// Find the wallet for the user
	filter := bson.M{"user_id": userID}
	update := bson.M{"$push": bson.M{"tickers": bson.M{"$each": tickerPredictions}}}
	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.Background(), filter, update, opts)
	return err
}

func GetWallet(userID string) (*Wallet, error) {
	collection := config.DB.Collection("wallets")
	filter := bson.M{"user_id": userID}

	var wallet Wallet
	err := collection.FindOne(context.Background(), filter).Decode(&wallet)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &wallet, err
}
