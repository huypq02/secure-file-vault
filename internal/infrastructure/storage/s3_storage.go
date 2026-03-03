package storage

import (
	"bytes"
	"context"
	"io"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/huypq02/secure-file-vault/internal/domain"
)

// s3Storage implements domain.StorageService using S3-compatible backend
type s3Storage struct {
	client S3API
	config *domain.StorageConfig
}

// NewS3StorageService creates a storage service with S3-compatible backend
func NewS3StorageService(client S3API, cfg *domain.StorageConfig) domain.StorageService {
	return &s3Storage{
		client: client,
		config: cfg,
	}
}

func (s *s3Storage) Store(ctx context.Context, file *domain.FileMetadata, data []byte) (*domain.StorageResult, error) {
	// Validate file metadata
	if err := file.Validate(); err != nil {
		return nil, err
	}

	// Upload file to S3
	output, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:            aws.String(s.config.Bucket),
		Key:               aws.String(file.Filename),
		Body:              bytes.NewReader(data),
		ContentType:       aws.String(file.ContentType),
		ChecksumAlgorithm: types.ChecksumAlgorithmSha256,
	})
	if err != nil {
		return nil, err
	}

	// Create storage result
	if output.ETag == nil || output.ChecksumSHA256 == nil {
		return nil, domain.ErrUploadFailed
	}

	// Return storage result with metadata
	return &domain.StorageResult{
		Checksum: *output.ChecksumSHA256,
		StoredAt: time.Now(),
		Properties: map[string]string{
			"filename":    file.OriginalName,
			"size":        strconv.Itoa(int(file.Size)),
			"contentType": file.ContentType,
		},
	}, nil
}

func (s *s3Storage) Retrieve(ctx context.Context, fileID string) ([]byte, error) {
	// Retrieve file from S3
	output, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(fileID),
	})
	if err != nil {
		return nil, err
	}

	defer output.Body.Close()
	// Read the file data
	data, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *s3Storage) Delete(ctx context.Context, fileID string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(fileID),
	})
	if err != nil {
		return err
	}

	return nil
}
