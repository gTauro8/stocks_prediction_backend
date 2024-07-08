package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"stock_prediction_backend/config"
)

type UserResponses struct {
	UserID                  string   `json:"user_id" bson:"user_id"`
	Age                     int      `json:"age" bson:"age"`
	EmploymentStatus        string   `json:"employment_status" bson:"employment_status"`
	AnnualIncome            string   `json:"annual_income" bson:"annual_income"`
	Wealth                  string   `json:"wealth" bson:"wealth"`
	InvestmentObjectives    []string `json:"investment_objectives" bson:"investment_objectives"`
	InvestmentExperience    string   `json:"investment_experience" bson:"investment_experience"`
	MarketKnowledge         string   `json:"market_knowledge" bson:"market_knowledge"`
	ReactionToMarketDrop    string   `json:"reaction_to_market_drop" bson:"reaction_to_market_drop"`
	RiskTolerancePercentage string   `json:"risk_tolerance_percentage" bson:"risk_tolerance_percentage"`
	CapitalSafetyImportance string   `json:"capital_safety_importance" bson:"capital_safety_importance"`
	RiskForHigherReturns    string   `json:"risk_for_higher_returns" bson:"risk_for_higher_returns"`
	KnownInvestments        []string `json:"known_investments" bson:"known_investments"`
	PastInvestments         []string `json:"past_investments" bson:"past_investments"`
	PreferredInvestments    []string `json:"preferred_investments" bson:"preferred_investments"`
	RiskProfile             string   `json:"risk_profile" bson:"risk_profile"`
	InvestmentAmount        float64  `json:"investment_amount" bson:"investment_amount"`
}

func SaveUserResponses(responses *UserResponses) error {
	collection := config.DB.Collection("user_responses")
	_, err := collection.InsertOne(context.Background(), responses)
	return err
}

func UpdateUserResponses(responses *UserResponses) error {
	collection := config.DB.Collection("user_responses")
	filter := bson.M{"user_id": responses.UserID}
	update := bson.M{"$set": responses}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

func GetUserResponses(userID string) (*UserResponses, error) {
	collection := config.DB.Collection("user_responses")
	filter := bson.M{"user_id": userID}

	var responses UserResponses
	err := collection.FindOne(context.Background(), filter).Decode(&responses)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &responses, err
}
