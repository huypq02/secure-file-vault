package dto

import (
	"mime/multipart"
	"time"
)

// FileItem represents a single file
type FileItem struct {
	// File identification
	OriginalName string `json:"originalName" form:"originalName" validate:"required"`
	Filename     string `json:"filename" form:"filename" validate:"required"`
	ContentType  string `json:"contentType" form:"contentType" validate:"required"`
	Size         int64  `json:"size" form:"size" validate:"required,min=1"`

	// File content
	Data []byte `json:"-"` // Raw file data
}

// FileUploadResult represents the result of a single file upload within a batch
type FileUploadResult struct {
	FileID       string    `json:"fileId,omitempty"`
	OriginalName string    `json:"originalName"`
	Filename     string    `json:"filename,omitempty"`
	Size         int64     `json:"size,omitempty"`
	UploadedAt   time.Time `json:"uploadedAt,omitempty"`
	Checksum     string    `json:"checksum,omitempty"`
}

// UploadFileRequest represents the request structure for file upload
type UploadFileRequest struct {
	// File identification
	FileItem

	// Optional metadata
	Description string            `json:"description" form:"description"`
	Tags        []string          `json:"tags" form:"tags"`
	Metadata    map[string]string `json:"metadata" form:"metadata"`

	// For multipart form uploads
	FileHeader *multipart.FileHeader `json:"-" form:"-"` // Multipart file header

	// Context data
	UserID string `json:"-" form:"-"` // From authentication
}

// UploadFileResponse represents the response structure for file upload
type UploadFileResponse struct {
	FileUploadResult
	Message string `json:"message,omitempty"` // Success message
}

// BatchUploadRequest represents a request to upload multiple files
type BatchUploadRequest struct {
	// Shared metadata for all files
	Description string            `json:"description" form:"description"`
	Tags        []string          `json:"tags" form:"tags"`
	Metadata    map[string]string `json:"metadata" form:"metadata"`

	// Collection of files
	Files []FileItem `json:"-"`

	// Context data
	UserID string `json:"-"` // From authentication
}

// BatchUploadResponse represents the response for a batch upload
type BatchUploadResponse struct {
	// Summary information
	TotalFiles   int       `json:"totalFiles"`
	SuccessCount int       `json:"successCount"`
	FailureCount int       `json:"failureCount"`
	BatchID      string    `json:"batchId,omitempty"`
	ProcessedAt  time.Time `json:"processedAt"`

	// Detailed results for each file
	Results []FileUploadResult `json:"results"`

	// Overall status
	Message string `json:"message"`
}
