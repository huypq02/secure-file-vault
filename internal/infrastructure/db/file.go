package db

import (
	"fmt"

	"github.com/huypq02/secure-file-vault/internal/domain"
)

type fileRepository struct {
	conn *connection
}

func NewFileRepository(conn *connection) domain.FileRepository {
	return &fileRepository{
		conn: conn,
	}
}

func (f *fileRepository) FindByID(fileID string) (*domain.FileMetadata, error) {
	if fileID == "" {
		return nil, fmt.Errorf("fileID cannot be empty")
	}

	var fileModel fileModel
	if err := f.conn.DB.Debug().Find(&fileModel, "id = ?", fileID).Error; err != nil {
		return nil, fmt.Errorf("failed to find file: %w", err)
	}

	// Convert fileModel to domain.FileMetadata
	fileEntity := fileModel.ToEntity()
	if fileEntity.ID == "" {
		return nil, fmt.Errorf("file not found with ID: %s", fileID)
	}

	return &fileEntity, nil
}

func (f *fileRepository) Update(file *domain.FileMetadata) error {
	if file == nil {
		return fmt.Errorf("file cannot be nil")
	}

	if file.ID == "" {
		return fmt.Errorf("file ID cannot be empty")
	}

	// Replace with actual logic to update file metadata in the database.
	var fileModel fileModel
	fileModel.FromEntity(file)
	if err := f.conn.DB.Save(&fileModel).Error; err != nil {
		return fmt.Errorf("failed to update file: %w", err)
	}

	fmt.Printf("File %s updated successfully\n", file.ID)
	return nil
}

func (f *fileRepository) Save(file *domain.FileMetadata) error {
	if file == nil {
		return fmt.Errorf("file cannot be nil")
	}

	if file.ID == "" {
		return fmt.Errorf("file ID cannot be empty")
	}
	// Replace with actual logic to save file metadata in the database.
	var fileModel fileModel
	fileModel.FromEntity(file)
	if err := f.conn.DB.Create(&fileModel).Error; err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}
