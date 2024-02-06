package main

import (
	"github.com/mnzsss/go-jobs/config"
	"github.com/mnzsss/go-jobs/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize config
	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	// Initialize router
	router.Init()
}
