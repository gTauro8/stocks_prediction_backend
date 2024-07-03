package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"myapp/models"
	"net/http"
)

type RecommendationResponse struct {
	Tickers     []string                            `json:"tickers"`
	Predictions map[string][]map[string]interface{} `json:"predictions"`
}

func GetRecommendations(userResponses models.UserResponses) (RecommendationResponse, error) {
	url := "http://127.0.0.1:8000/recommend"

	jsonData, err := json.Marshal(userResponses)
	if err != nil {
		return RecommendationResponse{}, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return RecommendationResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RecommendationResponse{}, errors.New("failed to get recommendations")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RecommendationResponse{}, err
	}

	fmt.Printf("Response body: %s\n", string(body)) // Debug logging

	var recommendations RecommendationResponse
	err = json.Unmarshal(body, &recommendations)
	if err != nil {
		return RecommendationResponse{}, err
	}

	fmt.Printf("Tickers: %v\n", recommendations.Tickers)         // Debug logging
	fmt.Printf("Predictions: %v\n", recommendations.Predictions) // Debug logging

	return recommendations, nil
}
