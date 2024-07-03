package controllers

import (
	"github.com/gin-gonic/gin"
	"myapp/models"
	"net/http"
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

	c.JSON(http.StatusOK, gin.H{"message": "Responses updated successfully"})
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
