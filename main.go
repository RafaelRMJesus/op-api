package main

import (
	"github.com/RafaelRMJesus/op-api/config"
	"github.com/RafaelRMJesus/op-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.InitDb()
	if err != nil {
		logger.Errorf("CONFIG_INITIALIZATION_ERROR: %v", err)
		return
	}

	router.InitializeRouter()
}
