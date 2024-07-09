package controllers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"stock_prediction_backend/models"

	"github.com/gin-gonic/gin"
)

func CreatePortfolio(c *gin.Context) {
	var requestData struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	// Recupera i dati dell'utente dal database
	userResponses, err := models.GetUserResponses(requestData.UserID)
	if err != nil {
		fmt.Printf("Error retrieving user responses: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving user responses", "error": err.Error()})
		return
	}

	if userResponses == nil {
		fmt.Println("User responses not found")
		c.JSON(http.StatusNotFound, gin.H{"message": "User responses not found"})
		return
	}

	// Prepara i parametri per il comando Python
	investmentAmount := fmt.Sprintf("%f", userResponses.InvestmentAmount)
	riskProfile := userResponses.RiskProfile

	fmt.Printf("Running Python script with parameters: UserID=%s, InvestmentAmount=%s, RiskProfile=%s\n", requestData.UserID, investmentAmount, riskProfile)

	// Usa il percorso del sorgente del progetto
	projectDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting project directory: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get project directory", "error": err.Error()})
		return
	}
	scriptPath := filepath.Join(projectDir, "test.py")

	// Comando per eseguire il file Python con l'ID utente, `investmentAmount` e `riskProfile` come argomenti
	cmd := exec.Command("python", scriptPath, requestData.Username, requestData.Password, requestData.UserID, investmentAmount, riskProfile)

	// Esegui il comando e cattura l'output
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing Python script: %v\n", err)
		fmt.Printf("Python script output: %s\n", string(output))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create portfolio", "error": err.Error(), "output": string(output)})
		return
	}

	fmt.Printf("Python script output: %s\n", string(output))
	fmt.Printf("Parameters received: Username=%s, Password=%s, UserID=%s\n", requestData.Username, requestData.Password, requestData.UserID)
	fmt.Printf("Running Python script with parameters: UserID=%s, InvestmentAmount=%s, RiskProfile=%s\n", requestData.UserID, investmentAmount, riskProfile)

	c.JSON(http.StatusOK, gin.H{"message": "Portfolio created successfully", "output": string(output)})
}
