package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/huypq02/secure-file-vault/internal/domain"
)

// S3API defines the interface for S3-compatible storage operations
type S3API interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
}

// NewS3Client creates a new S3-compatible client based on configuration
func NewS3Client(cfg *domain.StorageConfig) (S3API, error) {
	// Load AWS SDK configuration
	awsConfig, err := loadAWSConfig(cfg)
	if err != nil {
		return nil, err
	}

	// Create S3-compatible client with custom endpoint support (for MinIO, etc.)
	s3Client := s3.NewFromConfig(awsConfig, func(o *s3.Options) {
		if cfg.Endpoint != "" {
			o.BaseEndpoint = aws.String(cfg.Endpoint)
			o.UsePathStyle = cfg.ForcePathStyle
		}
		if cfg.ForcePathStyle {
			o.UsePathStyle = true
		}
	})

	return s3Client, nil
}

func loadAWSConfig(cfg *domain.StorageConfig) (aws.Config, error) {
	ctx := context.TODO()
	// Validate required fields
	if cfg.Region == "" || cfg.AccessKeyID == "" || cfg.SecretAccessKey == "" {
		return config.LoadDefaultConfig(ctx,
			config.WithRegion(cfg.Region))
	}

	// Use AWS SDK's default config loading mechanism
	return config.LoadDefaultConfig(ctx,
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AccessKeyID, cfg.SecretAccessKey, "",
		)),
	)
}
