package boot

import (
	"goblockhub/internal/consts"
	"goblockhub/internal/manager"
	"goblockhub/internal/schedule"
	"goblockhub/internal/service"
	"time"

	"github.com/CharlesWhiteSun/gomodx/logger"
)

func Initial() {
	logger.InitLogger(logger.DebugLevel)

	timeManager := schedule.NewTimeManager()
	binanceSvc := service.NewBinanceService()
	okxSvc := service.NewOKXService()

	manager.InitOnce(&manager.Registry{
		TimeManager:    timeManager,
		BinanceService: binanceSvc,
		OKXService:     okxSvc,
	})

	timeManager.RegisterAndStart(consts.BINANCE, schedule.NewTimeSchedule(consts.BINANCE, binanceSvc), 1*time.Minute, 30*time.Minute)
}
