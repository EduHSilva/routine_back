package main

import (
	"github.com/Eduardo681/go_routine/config"
	"github.com/Eduardo681/go_routine/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.ErrF("Configuration initialization failed with error: %v", err)
		return
	}
	router.Init()
}
