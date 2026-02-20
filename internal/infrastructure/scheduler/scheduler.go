package scheduler

import (
	"fmt"

	"github.com/huypq02/secure-file-vault/internal/domain"
	"github.com/robfig/cron/v3"
)

type scheduler struct {
	cron           *cron.Cron
	fileRepo       domain.FileRepository
	storageService domain.StorageService
}

func NewScheduler(
	fileRepo domain.FileRepository,
	storage domain.StorageService,
) domain.Scheduler {
	return &scheduler{
		cron:           cron.New(),
		fileRepo:       fileRepo,
		storageService: storage,
	}
}

func (s *scheduler) Start() {
	c := s.cron
	c.AddFunc("0 2 * * *", func() {
		// Runs at 2 AM daily
		s.cleanupExpiredFiles()
	})

	c.Start()
	fmt.Println("Scheduler started")
}

func (s *scheduler) cleanupExpiredFiles() {
	// TODO: Implement: delete expired file automatically
}

func (s *scheduler) Stop() {
	s.cron.Stop()
}
