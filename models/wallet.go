package models

import (
	"context"
	"fmt"
	"stock_prediction_backend/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Prediction struct {
	Date  string  `json:"date" bson:"date"`
	Value float64 `json:"value" bson:"value"`
}

type TickerPrediction struct {
	Ticker         string       `json:"ticker" bson:"ticker"`
	Predictions    []Prediction `json:"predictions" bson:"predictions"`
	AmountInvested float64      `json:"amount_invested" bson:"amount_invested"`
}

type Wallet struct {
	UserID       string             `json:"user_id" bson:"user_id"`
	Tickers      []TickerPrediction `json:"tickers" bson:"tickers"`
	ExpectedGain map[string]float64 `json:"expected_gain" bson:"expected_gain"`
	DateAdded    time.Time          `json:"date_added" bson:"date_added"`
}

func AddToWallet(userID string, tickerPredictions []TickerPrediction, expectedGain map[string]float64) error {
	collection := config.DB.Collection("wallets")

	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$push": bson.M{"tickers": bson.M{"$each": tickerPredictions}},
		"$set": bson.M{
			"expected_gain": expectedGain,
			"date_added":    time.Now(),
		},
	}
	opts := options.Update().SetUpsert(true)

	fmt.Printf("Updating wallet for user: %s with filter: %+v and update: %+v\n", userID, filter, update)

	_, err := collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		fmt.Printf("Error updating wallet: %v\n", err)
	}
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

func UpdateWallet(userID string, tickerPredictions []TickerPrediction, expectedGain map[string]float64) error {
	collection := config.DB.Collection("wallets")

	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$push": bson.M{"tickers": bson.M{"$each": tickerPredictions}},
		"$set": bson.M{
			"expected_gain": expectedGain,
			"date_added":    time.Now(),
		},
	}
	opts := options.Update()

	_, err := collection.UpdateOne(context.Background(), filter, update, opts)
	return err
}
