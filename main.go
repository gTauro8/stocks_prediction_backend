// Package main is the entry point for the myapp application.
// It sets up the configuration, starts routines, and configures the web server.
package main

import (
	"stock_prediction_backend/config"
	"stock_prediction_backend/routes"
	"stock_prediction_backend/routines"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// main is the entry point of the application.
// It sets up the application configuration, starts background routines,
// and configures the HTTP server with CORS and routes.
func main() {
	// Setup application configuration
	config.Setup()

	// Start the daily update routine in a separate goroutine
	go routines.StartDailyUpdateRoutine()

	// Create a new Gin router
	r := gin.Default()

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
		return
	}
}
