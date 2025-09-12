//go:build slow
// +build slow

package service

import (
	"sync"
	"testing"
	"time"
)

// 測試首次呼叫 CurrentTime，serverTime 為 0
func TestCurrentTimeFirstCall(t *testing.T) {
	svc := NewBinanceService()

	ok, ts := svc.CurrentTime()
	if ok {
		t.Errorf("Expected ok=false for initial serverTime=0, got true")
	}
	if ts != 0 {
		t.Errorf("Expected timestamp=0 for initial serverTime=0, got %d", ts)
	}
}

// 測試呼叫後會根據本地單調鐘時間累加
func TestCurrentTimeIncrement(t *testing.T) {
	svc := NewBinanceService()

	// 模擬伺服器時間
	svc.mu.Lock()
	svc.serverTime = 1_000_000
	svc.lastUpdate = time.Now().Add(-2 * time.Second) // 模擬 2 秒前同步
	svc.mu.Unlock()

	ok, ts := svc.CurrentTime()
	if !ok {
		t.Errorf("Expected ok=true, got false")
	}

	if ts < 1_000_000+1990 || ts > 1_000_000+2100 { // 允許誤差
		t.Errorf("Expected serverTime to increase by ~2000ms, got %d", ts)
	}
}

// 測試多執行緒安全
func TestCurrentTimeConcurrent(t *testing.T) {
	svc := NewBinanceService()
	svc.mu.Lock()
	svc.serverTime = 500_000
	svc.lastUpdate = time.Now().Add(-1 * time.Second)
	svc.mu.Unlock()

	var wg sync.WaitGroup
	const numGoroutines = 10000
	results := make([]int64, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			_, ts := svc.CurrentTime()
			results[idx] = ts
		}(i)
	}
	wg.Wait()

	// 檢查所有結果不為 0，並大於初始值
	for i, ts := range results {
		if ts <= 500_000 {
			t.Errorf("Concurrent call %d returned timestamp <= initial, got %d", i, ts)
		}
	}
}

// 測試實際經過的時間會反應在 CurrentTime
func TestCurrentTimeRealElapsed(t *testing.T) {
	svc := &BinanceService{
		serverTime: time.Now().UnixMilli(), // 模擬伺服器時間
		lastUpdate: time.Now(),
	}

	for i := 0; i < 5; i++ {
		ok, ts := svc.CurrentTime()
		if !ok {
			t.Errorf("CurrentTime 尚未初始化正確")
			continue
		}

		t.Logf("第 %d 秒: CurrentTime = %d", i+1, ts)
		time.Sleep(1 * time.Second)
	}
}
