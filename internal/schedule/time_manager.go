package schedule

import (
	"goblockhub/internal/consts"
	"sync"
	"time"
)

type TimeManager struct {
	schedules map[consts.Platform]*TimeSchedule
	mu        sync.RWMutex
}

func NewTimeManager() *TimeManager {
	return &TimeManager{
		schedules: make(map[consts.Platform]*TimeSchedule),
	}
}

func (tm *TimeManager) RegisterAndStart(platform consts.Platform, schedule *TimeSchedule, retry time.Duration, interval time.Duration) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.schedules[platform] = schedule
	schedule.Start(retry, interval)
}

func (tm *TimeManager) GetTime(platform consts.Platform) (int64, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	if s, ok := tm.schedules[platform]; ok {
		return s.Get()
	}
	return 0, false
}
