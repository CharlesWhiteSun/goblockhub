package main

import (
	"goblockhub/internal/router"
	"goblockhub/internal/server"

	"github.com/CharlesWhiteSun/gomodx/logger"
)

func main() {
    logger.InitLogger(logger.DebugLevel)
	
	s := server.NewGinServer(":8080", router.SetupRoutes)
	if err := s.Start(); err != nil {
		logger.Errorf("server start error: %v", err)
	}
}
