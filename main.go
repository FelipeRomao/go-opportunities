package main

import "github.com/FelipeRomao/go-opportunities/config"

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("error initializing config: %v", err)
		return
	}
}
