package domain

import (
	"strings"
	"time"
)

// FileStatus represents upload/download status
type FileStatus string

const (
	FileStatusPending   FileStatus = "PENDING"
	FileStatusUploading FileStatus = "UPLOADING"
	FileStatusCompleted FileStatus = "COMPLETED"
	FileStatusFailed    FileStatus = "FAILED"
)

// AccessPolicy defines file access permissions
type AccessPolicy string

const (
	AccessPolicyPrivate    AccessPolicy = "PRIVATE"
	AccessPolicyPublicRead AccessPolicy = "PUBLIC_READ"
	AccessPolicyOwnerOnly  AccessPolicy = "OWNER_ONLY"
)

type FileMetadata struct {
	// Core identification (pure business)
	ID           string
	OriginalName string
	Filename     string
	Description  string

	// File properties (business attributes)
	Size        int64
	ContentType string
	MimeType    string
	Encoding    string

	// Security & Integrity (business rules)
	ChecksumSHA256 string
	AccessPolicy   AccessPolicy
	IsPublic       bool

	// Ownership & Access (business concepts)
	Owner         string
	Uploader      string
	AccessCount   int
	DownloadCount int

	// Status tracking (business state)
	UploadStatus FileStatus

	// Timestamps (business audit trail)
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ExpiresAt      *time.Time
	LastAccessedAt *time.Time
}

// BUSINESS VALIDATION METHODS
func (f *FileMetadata) Validate() error {
	if f == nil {
		return ErrNilFile
	}

	if strings.TrimSpace(f.Filename) == "" {
		return ErrInvalidFilename
	}

	if len(f.Filename) > MaxFilenameLen {
		return ErrInvalidFilename
	}

	if f.Size <= 0 {
		return ErrInvalidFileSize
	}

	if f.Size > MaxFileSize {
		return ErrFileTooLarge
	}

	return nil
}

// BUSINESS LOGIC METHODS
func (f *FileMetadata) IsExpired() bool {
	if f.ExpiresAt == nil {
		return false
	}
	return time.Now().After(*f.ExpiresAt)
}

func (f *FileMetadata) CanDownload() error {
	if f.UploadStatus != FileStatusCompleted {
		return ErrFileNotReady
	}

	if f.IsExpired() {
		return ErrFileExpired
	}

	return nil
}

func (f *FileMetadata) IncrementDownload() {
	f.DownloadCount++
	f.AccessCount++
	now := time.Now()
	f.LastAccessedAt = &now
	f.UpdatedAt = now
}

func (f *FileMetadata) IsAccessibleBy(userID string) bool {
	if f.AccessPolicy == AccessPolicyPublicRead || f.IsPublic {
		return true
	}

	return f.Owner == userID || f.Uploader == userID
}

func (f *FileMetadata) SetExpiry(hours int) {
	if hours > 0 {
		expiry := time.Now().Add(time.Duration(hours) * time.Hour)
		f.ExpiresAt = &expiry
	} else {
		f.ExpiresAt = nil
	}
}

func (f *FileMetadata) MarkAsUploading() {
	f.UploadStatus = FileStatusUploading
	f.UpdatedAt = time.Now()
}

func (f *FileMetadata) MarkAsCompleted() {
	f.UploadStatus = FileStatusCompleted
	f.UpdatedAt = time.Now()
}

func (f *FileMetadata) MarkAsFailed() {
	f.UploadStatus = FileStatusFailed
	f.UpdatedAt = time.Now()
}

func (f *FileMetadata) IsOwnedBy(userID string) bool {
	return f.Owner == userID
}
