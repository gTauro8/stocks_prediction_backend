package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stock_prediction_backend/models"
	"stock_prediction_backend/utils"
)

func SaveUserResponses(c *gin.Context) {
	var responses models.UserResponses
	userID := c.Param("user_id")

	if err := c.ShouldBindJSON(&responses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	responses.UserID = userID

	if err := models.SaveUserResponses(&responses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving responses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Responses saved successfully"})
}

func UpdateUserResponses(c *gin.Context) {
	var responses models.UserResponses
	userID := c.Param("user_id")

	if err := c.ShouldBindJSON(&responses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	responses.UserID = userID

	if err := models.UpdateUserResponses(&responses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating responses"})
		return
	}

	recommendations, err := utils.GetRecommendations(responses)
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

	amountInvested := responses.InvestmentAmount
	if err := models.AddToWallet(userID, tickerPredictions, amountInvested, recommendations.ExpectedGain); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update wallet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Responses and wallet updated successfully"})
}

func GetUserResponses(c *gin.Context) {
	userID := c.Param("user_id")

	responses, err := models.GetUserResponses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving responses"})
		return
	}

	if responses == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Responses not found"})
		return
	}

	c.JSON(http.StatusOK, responses)
}

func GetProfile(c *gin.Context) {
	userID := c.Param("user_id")

	responses, err := models.GetUserResponses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving responses"})
		return
	}

	if responses == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Responses not found"})
		return
	}

	wallet, err := models.GetWallet(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving wallet"})
		return
	}

	if wallet == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_responses": responses,
		"wallet":         wallet,
	})
}
