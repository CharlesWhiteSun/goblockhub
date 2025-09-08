package service

type BinanceService struct{}

func NewBinanceService() IPlatformService {
	return &BinanceService{}
}

func (s *BinanceService) GetStatus() string {
	return "Binance OK"
}
