package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"go-project/config"
)

func TestDatabaseConnection(t *testing.T) {
	// Move up to the root directory where .env is located
	if err := moveToRootDir(); err != nil {
		t.Fatalf("Failed to move to root directory: %v", err)
	}

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

func moveToRootDir() error {
	// Keep moving up until we find the go.mod file
	for {
		if _, err := os.Stat("go.mod"); err == nil {
			return nil
		}
		if err := os.Chdir(".."); err != nil {
			return err
		}
		// Prevent infinite loop
		if wd, _ := os.Getwd(); wd == filepath.Dir(wd) {
			return fmt.Errorf("reached root directory without finding go.mod")
		}
	}
}