package service

type IPlatformService interface {
	GetStatus() (bool, error)
}
