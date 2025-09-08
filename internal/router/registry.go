package router

import (
	"goblockhub/internal/handler"
	"goblockhub/internal/service"
)

func getPlatformHandlers() []handler.IPlatformHandler {
	return []handler.IPlatformHandler{
		handler.NewBinanceHandler(service.NewBinanceService()),
		handler.NewOKXHandler(service.NewOKXService()),
	}
}
