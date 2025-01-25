package main

import (
	"context"
	"log"
	"time"

	"go-project/config"
	"go-project/db"
	"go-project/server"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	if err := db.Initialize(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Test database connection
	if err := testDatabaseConnection(); err != nil {
		log.Fatalf("Database connection test failed: %v", err)
	}

	// Start the server
	server.Start(cfg)
}

func testDatabaseConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := db.Client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	log.Println("Successfully connected to the database")
	return nil
}

