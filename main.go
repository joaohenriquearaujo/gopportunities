package main

import (
	"github.com/joaohenriquearaujo/gopportunities.git/config"
	"github.com/joaohenriquearaujo/gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
