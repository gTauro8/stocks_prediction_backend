package controllers

import (
	"github.com/gin-gonic/gin"
	"myapp/models"
	"net/http"
)

func AddToWallet(c *gin.Context) {
	userID := c.Param("user_id")

	var tickers []string
	if err := c.ShouldBindJSON(&tickers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var tickerPredictions []models.TickerPrediction
	for _, ticker := range tickers {
		tickerPredictions = append(tickerPredictions, models.TickerPrediction{
			Ticker: ticker,
			// For now, we're not adding predictions here. Adjust as needed.
			Predictions: []models.Prediction{},
		})
	}

	if err := models.AddToWallet(userID, tickerPredictions); err != nil {
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
