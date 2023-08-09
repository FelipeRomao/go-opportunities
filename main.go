package main

import (
	"github.com/FelipeRomao/go-opportunities/config"
	"github.com/FelipeRomao/go-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize the config
	err := config.Init()

	if err != nil {
		logger.Errorf("error initializing config: %v", err)
		return
	}

	// Initialize the server
	router.Initialize()
}
