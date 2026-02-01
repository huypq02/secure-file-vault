package file

import (
	"context"

	"github.com/huypq02/secure-file-vault/internal/domain"
	"github.com/huypq02/secure-file-vault/internal/usecase/dto"
)

type UploadFileUsecase interface {
	UploadFile(ctx context.Context, file *dto.UploadFileRequest) (*dto.UploadFileResponse, error)
	UploadMultipleFiles(ctx context.Context, file *dto.BatchUploadRequest) (*dto.BatchUploadResponse, error)
}

type uploadFileUsecase struct {
	fileRepo       domain.FileRepository
	storageService domain.StorageService
	idService      domain.IDService
}

func NewUploadFileUsecase(fileRepo domain.FileRepository, storageService domain.StorageService, idService domain.IDService) UploadFileUsecase {
	return &uploadFileUsecase{
		fileRepo:       fileRepo,
		storageService: storageService,
		idService:      idService,
	}
}
