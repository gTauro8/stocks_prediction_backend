package main

import (
	"log"
	"time"

	"stock_prediction_backend/config"
	"stock_prediction_backend/routes"
	"stock_prediction_backend/routines"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setup application configuration
	config.Setup()

	// Start the daily update routine in a separate goroutine
	go routines.StartDailyUpdateRoutine()

	// Create a new Gin router
	r := gin.Default()

	// Add logging middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Configure CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup routes
	routes.SetupRoutes(r)

	// Start the HTTP server on port 8080
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
