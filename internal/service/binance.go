package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/CharlesWhiteSun/gomodx/logger"
)

type BinanceService struct {
	mu         sync.Mutex
	serverTime int64     // 最近一次從伺服器抓到的時間
	lastUpdate time.Time // 本地時間對應上次 serverTime 的時間
}

func NewBinanceService() *BinanceService {
	return &BinanceService{}
}

func (s *BinanceService) GetStatus() (bool, error) {
	url := "https://api.binance.com/api/v3/ping"
	header := "Binance ping|"

	resp, err := http.Get(url)
	if err != nil {
		emsg := fmt.Errorf("%v get URL error: %v", header, err.Error())
		logger.Error(emsg)
		return false, emsg
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		emsg := fmt.Errorf("%v reading response error: %v", header, err.Error())
		logger.Error(emsg)
		return false, emsg
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		emsg := fmt.Errorf("%v parsing JSON error: %v", header, err.Error())
		logger.Error(emsg)
		return false, emsg
	}

	// Binance /ping 回傳空物件 {}，可以直接表示成功
	if len(result) == 0 {
		return true, nil
	}

	// 若有其他內容，回傳原始 JSON
	emsg := fmt.Errorf("%v unexpected body: %v", header, string(body))
	logger.Info(emsg)
	return true, emsg
}

func (s *BinanceService) GetTime() (bool, int64, error) {
	url := "https://api.binance.com/api/v3/time"
	header := "Binance time|"

	resp, err := http.Get(url)
	if err != nil {
		emsg := fmt.Errorf("%v get URL error: %v", header, err.Error())
		logger.Error(emsg)
		return false, 0, emsg
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		emsg := fmt.Errorf("%v reading response error: %v", header, err.Error())
		logger.Error(emsg)
		return false, 0, emsg
	}

	// 解析 JSON
	var result struct {
		ServerTime int64 `json:"serverTime"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		emsg := fmt.Errorf("%v parsing JSON error: %v", header, err.Error())
		logger.Error(emsg)
		return false, 0, emsg
	}

	s.mu.Lock()
	s.serverTime = result.ServerTime
	s.lastUpdate = time.Now()
	s.mu.Unlock()

	return true, result.ServerTime, nil
}

// CurrentTime 回傳隨本地時間增長的伺服器時間戳
func (s *BinanceService) CurrentTime() (bool, int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.serverTime == 0 {
		return false, 0
	}

	elapsed := time.Since(s.lastUpdate).Milliseconds()
	s.serverTime += elapsed
	s.lastUpdate = time.Now()
	return true, s.serverTime
}
