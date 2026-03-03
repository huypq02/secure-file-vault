package scheduler

import (
	"context"
	"fmt"

	"github.com/huypq02/secure-file-vault/internal/domain"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron           *cron.Cron
	fileRepo       domain.FileRepository
	storageService domain.StorageService
}

func NewScheduler(
	fileRepo domain.FileRepository,
	storage domain.StorageService,
) domain.Scheduler {
	return &Scheduler{
		cron:           cron.New(),
		fileRepo:       fileRepo,
		storageService: storage,
	}
}

func (s *Scheduler) Start() {
	c := s.cron
	c.AddFunc("0 2 * * *", func() {
		// Runs at 2 AM daily
		s.cleanupExpiredFiles()
	})

	c.Start()
	fmt.Println("Scheduler started")
}

func (s *Scheduler) cleanupExpiredFiles() {

	// TODO: Check expiration file exceed current date
	var fileIDList []string

	// Delete files saving s3 storage
	for _, fileID := range fileIDList {
		if err := s.storageService.Delete(context.TODO(), fileID); err != nil {
			fmt.Printf("Unexpected error while cleaning expired file ID %s: %v\n",
				fileID, err)
		}
	}

	fmt.Println("Cleanup expired files successfully")
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
}
