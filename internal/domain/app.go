package domain

type Application interface {
	Run(addr string) error
	Shutdown() error
}
