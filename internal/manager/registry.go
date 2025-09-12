package manager

import (
	"goblockhub/internal/schedule"
	"goblockhub/internal/service"
	"sync"
)

type Registry struct {
	TimeManager    *schedule.TimeManager
	BinanceService *service.BinanceService
	OKXService     *service.OKXService
}

var (
	mu       sync.RWMutex
	once     sync.Once
	registry *Registry
)

func InitOnce(r *Registry) {
	once.Do(func() {
		registry = r
	})
}

func Get() *Registry {
	mu.RLock()
	defer mu.RUnlock()
	return registry
}
