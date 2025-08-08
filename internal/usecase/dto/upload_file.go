package dto

import (
	"mime/multipart"
	"time"
)

// UploadFileRequest represents the request structure for file upload
type UploadFileRequest struct {
	// File identification
	OriginalName string `json:"originalName" form:"originalName" validate:"required"`
	Filename     string `json:"filename" form:"filename" validate:"required"`
	ContentType  string `json:"contentType" form:"contentType" validate:"required"`
	Size         int64  `json:"size" form:"size" validate:"required,min=1"`

	// Optional metadata
	Description string            `json:"description" form:"description"`
	Tags        []string          `json:"tags" form:"tags"`
	Metadata    map[string]string `json:"metadata" form:"metadata"`

	// File content
	File []byte `json:"-" form:"-"` // Raw file data

	// For multipart form uploads
	FileHeader *multipart.FileHeader `json:"-" form:"-"` // Multipart file header

	// Context data
	UserID string `json:"-" form:"-"` // From authentication
}

// UploadFileResponse represents the response structure for file upload
type UploadFileResponse struct {
	FileID       string    `json:"fileID,omitempty"`
	OriginalName string    `json:"originalName,omitempty"`
	Filename     string    `json:"filename,omitempty"`
	Size         int64     `json:"size,omitempty"`
	UploadedAt   time.Time `json:"uploadAt,omitempty"`
	Checksum     string    `json:"checksum,omitempty"`
	Message      string    `json:"message,omitempty"` // Success message
}
