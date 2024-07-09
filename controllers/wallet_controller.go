package controllers

import (
	"fmt"
	"net/http"
	"stock_prediction_backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func AddToWallet(c *gin.Context) {
	userID := c.Param("user_id")

	var requestData struct {
		Tickers []struct {
			Ticker         string              `json:"ticker"`
			AmountInvested float64             `json:"amount_invested"`
			Predictions    []models.Prediction `json:"predictions"`
		} `json:"tickers"`
		ExpectedGain map[string]float64 `json:"expected_gain"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Printf("Request data: %+v\n", requestData)

	var tickerPredictions []models.TickerPrediction
	for _, t := range requestData.Tickers {
		tickerPredictions = append(tickerPredictions, models.TickerPrediction{
			Ticker:         t.Ticker,
			Predictions:    t.Predictions,
			AmountInvested: t.AmountInvested,
		})
	}

	wallet := models.Wallet{
		UserID:       userID,
		Tickers:      tickerPredictions,
		ExpectedGain: requestData.ExpectedGain,
		DateAdded:    time.Now(), // Assicurati di aggiungere il campo DateAdded qui
	}

	fmt.Printf("Wallet to be added: %+v\n", wallet)

	if err := models.AddToWallet(userID, wallet.Tickers, wallet.ExpectedGain); err != nil {
		fmt.Printf("Error adding to wallet: %v\n", err)
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
