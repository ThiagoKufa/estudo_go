package main

import (
	"github.com/ThiagoKufa/estudo_go/config"
	"github.com/ThiagoKufa/estudo_go/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.ErrorF("Config initialization erro : %v", err)
		return
	}

	router.Initialize()
}
