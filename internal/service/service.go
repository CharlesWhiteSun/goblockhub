package service

type IStatusService interface {
	GetStatus() (bool, error)
}

type ITimeService interface {
	GetTime() (bool, int64, error)
}
