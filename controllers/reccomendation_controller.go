package controllers

import (
	"net/http"
	"stock_prediction_backend/models"
	"stock_prediction_backend/utils"

	"github.com/gin-gonic/gin"
)

func GetRecommendations(c *gin.Context) {
	userID := c.Param("user_id")

	userResponses, err := models.GetUserResponses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user responses"})
		return
	}

	if userResponses == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User responses not found"})
		return
	}

	recommendations, err := utils.GetRecommendations(*userResponses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get recommendations"})
		return
	}

	var tickerPredictions []models.TickerPrediction
	for ticker, preds := range recommendations.Predictions {
		var predictions []models.Prediction
		for _, pred := range preds {
			date, ok := pred["ds"].(string)
			if !ok {
				continue
			}
			value, ok := pred["yhat"].(float64)
			if !ok {
				continue
			}
			predictions = append(predictions, models.Prediction{
				Date:  date,
				Value: value,
			})
		}
		tickerPredictions = append(tickerPredictions, models.TickerPrediction{
			Ticker:      ticker,
			Predictions: predictions,
		})
	}

	if err := models.AddToWallet(userID, tickerPredictions, recommendations.ExpectedGain); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save recommendations to wallet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tickers":         recommendations.Tickers,
		"predictions":     recommendations.Predictions,
		"amount_invested": userResponses.InvestmentAmount,
		"expected_gain":   recommendations.ExpectedGain,
	})
}
