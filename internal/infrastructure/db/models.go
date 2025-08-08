package db

import (
	"time"

	"github.com/huypq02/secure-file-vault/internal/domain"
)


type fileModel struct {
	ID           string `gorm:"primaryKey;column:id;type:varchar(36)"`
	OriginalName string `gorm:"column:original_name;type:varchar(255);not null"`
	Filename     string `gorm:"column:filename;type:varchar(255);not null"`
	Path         string `gorm:"column:path;type:varchar(500);not null"`
	Size         int64  `gorm:"column:size;type:bigint;default:0"`
	ContentType  string `gorm:"column:content_type;type:varchar(100)"`
	MimeType     string `gorm:"column:mime_type;type:varchar(100)"`
	Description  string `gorm:"column:description;type:text"`

	// S3 Information
	S3Bucket    string `gorm:"column:s3_bucket;type:varchar(100)"`
	S3Key       string `gorm:"column:s3_key;type:varchar(500)"`
	S3Region    string `gorm:"column:s3_region;type:varchar(50)"`
	S3ETag      string `gorm:"column:s3_etag;type:varchar(100)"`
	S3VersionID string `gorm:"column:s3_version_id;type:varchar(100)"`

	// Security & Access
	Checksum     string `gorm:"column:checksum;type:varchar(64)"`
	AccessPolicy string `gorm:"column:access_policy;type:varchar(20);default:'PRIVATE'"`
	IsPublic     bool   `gorm:"column:is_public;type:boolean;default:false"`

	// Ownership
	Owner         string `gorm:"column:owner;type:varchar(36);not null"`
	Uploader      string `gorm:"column:uploader;type:varchar(36);not null"`
	AccessCount   int    `gorm:"column:access_count;type:integer;default:0"`
	DownloadCount int    `gorm:"column:download_count;type:integer;default:0"`

	// Status
	UploadStatus string `gorm:"column:upload_status;type:varchar(20);default:'PENDING'"`

	// Timestamps
	CreatedAt      time.Time  `gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt      time.Time  `gorm:"column:updated_at;type:timestamp;not null"`
	ExpiresAt      *time.Time `gorm:"column:expires_at;type:timestamp"`
	LastAccessedAt *time.Time `gorm:"column:last_accessed_at;type:timestamp"`
}

// TableName specifies table name in PostgreSQL
func (fileModel) TableName() string {
	return "files"
}

// ToEntity converts FileModel to domain.FileMetadata
func (f *fileModel) ToEntity() domain.FileMetadata {
	return domain.FileMetadata{
		ID:             f.ID,
		OriginalName:   f.OriginalName,
		Filename:       f.Filename,
		Description:    f.Description,
		Size:           f.Size,
		ContentType:    f.ContentType,
		MimeType:       f.MimeType,
		ChecksumSHA256: f.Checksum,
		AccessPolicy:   domain.AccessPolicy(f.AccessPolicy),
		IsPublic:       f.IsPublic,
		Owner:          f.Owner,
		Uploader:       f.Uploader,
		AccessCount:    f.AccessCount,
		DownloadCount:  f.DownloadCount,
		UploadStatus:   domain.FileStatus(f.UploadStatus),
		CreatedAt:      f.CreatedAt,
		UpdatedAt:      f.UpdatedAt,
		ExpiresAt:      f.ExpiresAt,
		LastAccessedAt: f.LastAccessedAt,
	}
}

// FromEntity converts domain.FileMetadata to FileModel
func (f *fileModel) FromEntity(file *domain.FileMetadata) {
	f.ID = file.ID
	f.OriginalName = file.OriginalName
	f.Filename = file.Filename
	f.Description = file.Description
	f.Size = file.Size
	f.ContentType = file.ContentType
	f.MimeType = file.MimeType
	f.Checksum = file.ChecksumSHA256
	f.AccessPolicy = string(file.AccessPolicy)
	f.IsPublic = file.IsPublic
	f.Owner = file.Owner
	f.Uploader = file.Uploader
	f.AccessCount = file.AccessCount
	f.DownloadCount = file.DownloadCount
	f.UploadStatus = string(file.UploadStatus)
	f.CreatedAt = file.CreatedAt
	f.UpdatedAt = file.UpdatedAt
	if file.ExpiresAt != nil {
		expiresAt := *file.ExpiresAt
		f.ExpiresAt = &expiresAt
	} else {
		f.ExpiresAt = nil
	}
	if file.LastAccessedAt != nil {
		lastAccessedAt := *file.LastAccessedAt
		f.LastAccessedAt = &lastAccessedAt
	} else {
		f.LastAccessedAt = nil
	}
}
