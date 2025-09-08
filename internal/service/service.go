package service

// 所有平台 Service 都實作這個介面
type IPlatformService interface {
	GetStatus() string
}
