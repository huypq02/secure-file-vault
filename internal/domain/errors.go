package domain

import "errors"

// Core business errors
var (
	ErrFileNotFound       = errors.New("file not found")
	ErrFileExpired        = errors.New("file has expired")
	ErrInvalidFilename    = errors.New("invalid filename")
	ErrFileTooLarge       = errors.New("file too large")
	ErrUnauthorized       = errors.New("unauthorized access")
	ErrInvalidFileID      = errors.New("invalid file ID")
	ErrNilFile            = errors.New("file cannot be nil")
	ErrAccessDenied       = errors.New("access denied")
	ErrFileNotReady       = errors.New("file upload not completed")
	ErrInvalidOwner       = errors.New("invalid file owner")
	ErrChecksumMismatch   = errors.New("file checksum mismatch")
	ErrFileAlreadyExists  = errors.New("file already exists")
	ErrQuotaExceeded      = errors.New("storage quota exceeded")
	ErrFileCorrupted      = errors.New("file is corrupted")
	ErrUnsupportedFormat  = errors.New("unsupported file format")
	ErrFailedToSaveFile   = errors.New("failed to save file")
	ErrFailedToDeleteFile = errors.New("failed to delete file")
	ErrFailedToUpdateFile = errors.New("failed to update file")
	ErrFailedToListFiles  = errors.New("failed to list files")
)

// Validation errors
var (
	ErrEmptyFilename      = errors.New("filename cannot be empty")
	ErrInvalidFileSize    = errors.New("invalid file size")
	ErrInvalidContentType = errors.New("invalid content type")
	ErrInvalidOwnerID     = errors.New("invalid owner ID")
)

// Storage errors
var (
	ErrInvalidS3Config    = errors.New("invalid S3 configuration")
	ErrInvalidStorageType = errors.New("invalid storage type")
	ErrUploadFailed       = errors.New("file upload failed")
	ErrDownloadFailed     = errors.New("file download failed")
)

// Domain constants
const (
	MaxFileSize    = 100 * 1024 * 1024 // 100MB
	MaxFilenameLen = 255
	DefaultExpiry  = 24 // hours
)
