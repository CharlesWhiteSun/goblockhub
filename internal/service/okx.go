package service

type OKXService struct{}

func NewOKXService() IPlatformService {
	return &OKXService{}
}

func (s *OKXService) GetStatus() string {
	return "OKX OK"
}
