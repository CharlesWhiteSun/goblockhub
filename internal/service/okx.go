package service

type OKXService struct{}

func NewOKXService() IPlatformService {
	return &OKXService{}
}

func (s *OKXService) GetStatus() (bool, error) {
	return true, nil
}
