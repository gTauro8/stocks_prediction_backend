// Package config provides the configuration setup for the application,
// including loading configuration from a file and setting up a MongoDB connection.
package config

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"time"
)

// Config holds the application configuration values.
type Config struct {
	JWTSecret string `json:"JWT_SECRET"`
}

// AppConfig is the global configuration instance.
var AppConfig Config

// DB is the global MongoDB database instance.
var DB *mongo.Database

// LoadConfig reads configuration from the specified JSON file and unmarshals it into AppConfig.
//
// Parameters:
// - filename: The path to the JSON configuration file.
//
// Returns:
// - error: An error if reading the file or unmarshalling the JSON fails, otherwise nil.
func LoadConfig(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &AppConfig)
	if err != nil {
		return err
	}

	return nil
}

// Setup initializes the application configuration by loading it from a file
// and sets up a connection to the MongoDB database. It terminates the application
// if there is any error during the setup process.
func Setup() {
	// Load the application configuration
	err := LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create a new MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	// Create a context with a timeout for the connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Assign the connected database instance to the global variable
	DB = client.Database("financial_guru")
}
