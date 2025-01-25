package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	MongoURI   string `mapstructure:"MONGO_URI"`
	DBName     string `mapstructure:"DB_NAME"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
}

func Load() (*Config, error) {
	// Get the current working directory
	workDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %v", err)
	}

	// Configure Viper
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(workDir) // Look in current directory
	viper.AddConfigPath(".")     // Also look in root directory
	viper.AutomaticEnv()         // Read environment variables

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return &config, nil
}

