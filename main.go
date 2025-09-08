package main

import (
	"goblockhub/internal/router"
	"goblockhub/internal/server"
)

func main() {
	s := server.NewGinServer(":8080", router.SetupRoutes)
	if err := s.Start(); err != nil {
		panic(err)
	}
}
