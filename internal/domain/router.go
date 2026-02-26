package domain

type Router interface {
	RegisterRoutes()
	Run(addr string) error
}
