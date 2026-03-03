package domain

import (
	"context"
	"time"
)

// Storage Service interface for file storage operations
type StorageService interface {
	// Business operations only
	Store(ctx context.Context, file *FileMetadata, data []byte) (*StorageResult, error)
	Retrieve(ctx context.Context, fileID string) ([]byte, error)
	Delete(ctx context.Context, fileID string) error

	// // Business-oriented methods
	// GenerateAccessURL(ctx context.Context, fileID string, expiration time.Duration) (string, error)
}

// Business-focused result
type StorageResult struct {
	StorageID  string
	Checksum   string
	StoredAt   time.Time
	Properties map[string]string
}
