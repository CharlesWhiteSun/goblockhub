package schedule

import (
	"fmt"
	"sync"
	"time"

	"goblockhub/internal/consts"
	"goblockhub/internal/service"

	"github.com/CharlesWhiteSun/gomodx/logger"
)

type TimeSchedule struct {
	platform consts.Platform
	svc      service.ITimeService

	mu         sync.RWMutex
	serverTime int64
}

func NewTimeSchedule(platform consts.Platform, svc service.ITimeService) *TimeSchedule {
	return &TimeSchedule{
		platform: platform,
		svc:      svc,
	}
}

func (t *TimeSchedule) Start(retry time.Duration, interval time.Duration) {
	header := "TimeSchedule Start|"

	go func() {
		for {
			err := t.fetchTime()
			if err != nil {
				emsg := fmt.Errorf("%v [%s] fetch servertime fail, retrying in %v| error: %v", header, t.platform, retry, err.Error())
				logger.Error(emsg)
				time.Sleep(retry)
				continue
			}
			time.Sleep(interval)
		}
	}()
}

// fetchTime 抓取時間並存入 serverTime
func (t *TimeSchedule) fetchTime() error {
	ok, ts, err := t.svc.GetTime()
	if !ok {
		return err
	}

	t.mu.Lock()
	t.serverTime = ts
	t.mu.Unlock()

	logger.Infof("[%s] fetch servertime OK: %v", t.platform, ts)
	return nil
}

// Get 取得最後抓取到的時間，執行續安全
func (t *TimeSchedule) Get() (int64, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if t.serverTime == 0 {
		return 0, false
	}
	return t.serverTime, true
}
