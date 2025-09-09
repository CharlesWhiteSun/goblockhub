package main

import (
	"goblockhub/internal/logger"
	"goblockhub/internal/router"
	"goblockhub/internal/server"
)

func main() {
    logger.InitLogger("logs", logger.DebugLevel)
	
	s := server.NewGinServer(":8080", router.SetupRoutes)
	if err := s.Start(); err != nil {
		logger.Errorf("server start error: %v", err)
	}
}
