package file

import (
	"context"

	"github.com/huypq02/secure-file-vault/internal/domain"
	"github.com/huypq02/secure-file-vault/internal/usecase/dto"
)

type DownloadFileUsecase interface {
	DownloadFile(ctx context.Context, fileID string) (*dto.DownloadFileResponse, error)
	// DownloadFileData(ctx context.Context, fileID string) (io.Reader, error)
	// GenerateAccessURL(ctx context.Context, fileID string, expiration time.Duration) (string, error)
}

type downloadFileUsecase struct {
	fileRepo       domain.FileRepository
	storageService domain.StorageService
}

func NewDownloadFileUsecase(fileRepo domain.FileRepository, storageService domain.StorageService) DownloadFileUsecase {
	return &downloadFileUsecase{
		fileRepo:       fileRepo,
		storageService: storageService,
	}
}
