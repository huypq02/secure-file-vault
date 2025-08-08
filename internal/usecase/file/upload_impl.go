package file

import (
	"context"
	"fmt"
	"time"

	"github.com/huypq02/secure-file-vault/internal/domain"
	"github.com/huypq02/secure-file-vault/internal/usecase/dto"
)

func (uc *uploadFileUsecase) UploadFile(ctx context.Context, req *dto.UploadFileRequest) (*dto.UploadFileResponse, error) {
	// Validate request
	if req == nil || req.Filename == "" || len(req.File) == 0 {
		return nil, fmt.Errorf("invalid file request or data")
	}
	if req.Size <= 0 {
		return nil, fmt.Errorf("invalid file size: %d", req.Size)
	}

	// Generate unique ID for the file
	fileID := uc.idService.Generate()

	// Create domain entity
	meta := &domain.FileMetadata{
		ID:           fileID,
		Filename:     req.Filename,
		OriginalName: req.OriginalName,
		Size:         req.Size,
		ContentType:  req.ContentType,
		Description:  req.Description,
		UploadStatus: domain.FileStatusUploading,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Business logic: Set file as uploading status
	meta.MarkAsUploading()

	// Save file data to storage
	storageResult, err := uc.storageService.Store(ctx, meta, req.File)
	if err != nil {
		// Mark as failed on storage error
		meta.MarkAsFailed()
		return nil, fmt.Errorf("failed to store file: %w", err)
	}

	// Business logic: Mark as completed after successful storage
	meta.MarkAsCompleted()

	// Persist file metadata to repository
	if err := uc.fileRepo.Save(meta); err != nil {
		return nil, fmt.Errorf("failed to save file metadata: %w", err)
	}

	// Create response DTO
	response := &dto.UploadFileResponse{
		FileID:       meta.ID,
		Filename:     meta.Filename,
		OriginalName: meta.OriginalName,
		Size:         meta.Size,
		Checksum:     storageResult.Checksum,
		Message:      "file uploaded successfully",
		UploadedAt:   meta.CreatedAt,
	}

	return response, nil
}
