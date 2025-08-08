package file

import (
	"context"

	"github.com/huypq02/secure-file-vault/internal/domain"
	"github.com/huypq02/secure-file-vault/internal/usecase/dto"
)

func (uc *downloadFileUsecase) DownloadFile(ctx context.Context, fileID string) (*dto.DownloadFileResponse, error) {
	// Get file metadata from repository
	meta, err := uc.fileRepo.FindByID(fileID)
	if err != nil {
		return nil, err
	}
	if meta == nil {
		return nil, domain.ErrFileNotFound // File not found
	}

	// Download file data from storage
	data, err := uc.storageService.Retrieve(ctx, meta.Filename)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, domain.ErrFileNotFound // File not found in storage
	}

	// Business logic: increment download count using domain method
	meta.IncrementDownload()

	// Persist updated metadata
	if err := uc.fileRepo.Update(meta); err != nil {
		return nil, err
	}

	// Create response DTO
	response := &dto.DownloadFileResponse{
		Filename: meta.Filename,
		Size:     meta.Size,
		Data:     data,
		Message:  "download file successfully",
	}

	// Return file metadata and data
	return response, nil
}
