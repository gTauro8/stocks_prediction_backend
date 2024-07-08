package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stock_prediction_backend/models"
)

func AddToWallet(c *gin.Context) {
	userID := c.Param("user_id")

	var requestData struct {
		Tickers        []string           `json:"tickers"`
		AmountInvested float64            `json:"amount_invested"`
		ExpectedGain   map[string]float64 `json:"expected_gain"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var tickerPredictions []models.TickerPrediction
	for _, ticker := range requestData.Tickers {
		tickerPredictions = append(tickerPredictions, models.TickerPrediction{
			Ticker:      ticker,
			Predictions: []models.Prediction{},
		})
	}

	wallet := models.Wallet{
		UserID:         userID,
		Tickers:        tickerPredictions,
		AmountInvested: requestData.AmountInvested,
		ExpectedGain:   requestData.ExpectedGain,
	}

	if err := models.AddToWallet(userID, wallet.Tickers, wallet.AmountInvested, wallet.ExpectedGain); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to wallet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Added to wallet"})
}

func GetWallet(c *gin.Context) {
	userID := c.Param("user_id")

	wallet, err := models.GetWallet(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get wallet"})
		return
	}

	if wallet == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	c.JSON(http.StatusOK, wallet)
}
