package main

import (
	"github.com/joho/godotenv"
	"github.com/mnzsss/go-jobs/config"
	"github.com/mnzsss/go-jobs/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Load .env variables
	if err := godotenv.Load(); err != nil {
		logger.Errorf("Error loading .env file")
		return
	}

	// Initialize config
	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	// Initialize router
	router.Init()
}
