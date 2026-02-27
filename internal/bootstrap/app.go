package bootstrap

import (
	"github.com/huypq02/secure-file-vault/internal/domain"
)

type App struct {
	router    domain.Router
	scheduler domain.Scheduler
}

func NewApp(
	r domain.Router,
	s domain.Scheduler,
) domain.Application {
	return &App{
		router:    r,
		scheduler: s,
	}
}

func (a *App) Run(addr string) error {
	a.scheduler.Start()
	return a.router.Run(addr)
}

func (a *App) Shutdown() error {
	a.scheduler.Stop()
	return nil
}
