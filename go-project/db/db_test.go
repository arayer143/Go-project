package db

import (
	"context"
	"log"
	"testing"
	"time"

	"go-project/config"
)

func TestDatabaseConnection(t *testing.T) {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	err = Initialize(cfg)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer Close()

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = Client.Ping(ctx, nil)
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the database")
}

