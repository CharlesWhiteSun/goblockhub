package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/CharlesWhiteSun/gomodx/logger"
)

type BinanceService struct{}

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

	return true, result.ServerTime, nil
}
