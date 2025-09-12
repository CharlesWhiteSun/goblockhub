package manager

type IJob interface {
	Name() string
	Run() error
}
