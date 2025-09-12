package schedule

type ISchedule interface {
	Start()
	Get() (int64, bool)
}
