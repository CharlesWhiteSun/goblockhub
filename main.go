package main

import (
	"goblockhub/internal/cmd/boot"
	"goblockhub/internal/router"
	"goblockhub/internal/server"

	"github.com/CharlesWhiteSun/gomodx/logger"
)

func main() {
	boot.Initial()

	s := server.NewGinServer(":8080", router.SetupRoutes)
	if err := s.Start(); err != nil {
		logger.Errorf("server start error: %v", err)
	}
}
