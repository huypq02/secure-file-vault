-- Database Script: Create files table for Secure File Vault
-- Version: 1.0.0
-- Author: Development Team
-- Date: 2025-01-23
-- Purpose: Initial database setup

-- =============================================================================
-- SECURE FILE VAULT - DATABASE SETUP SCRIPT v1.0.0
-- =============================================================================

-- Create files table
CREATE TABLE files (
    -- Core identification
    id VARCHAR(36) PRIMARY KEY,
    original_name VARCHAR(255) NOT NULL,
    filename VARCHAR(255) NOT NULL,
    path VARCHAR(500) NOT NULL,
    
    -- File properties
    size BIGINT DEFAULT 0 CHECK (size >= 0),
    content_type VARCHAR(100),
    mime_type VARCHAR(100),
    description TEXT,
    
    -- S3 Information (current structure)
    s3_bucket VARCHAR(100),
    s3_key VARCHAR(500),
    s3_region VARCHAR(50),
    s3_etag VARCHAR(100),
    s3_version_id VARCHAR(100),
    
    -- Security & Access
    checksum VARCHAR(64),
    access_policy VARCHAR(20) DEFAULT 'PRIVATE' CHECK (access_policy IN ('PRIVATE', 'PUBLIC_READ', 'OWNER_ONLY')),
    is_public BOOLEAN DEFAULT FALSE,
    
    -- Ownership
    owner VARCHAR(36) NOT NULL,
    uploader VARCHAR(36) NOT NULL,
    access_count INTEGER DEFAULT 0 CHECK (access_count >= 0),
    download_count INTEGER DEFAULT 0 CHECK (download_count >= 0),
    
    -- Status
    upload_status VARCHAR(20) DEFAULT 'PENDING' CHECK (upload_status IN ('PENDING', 'UPLOADING', 'COMPLETED', 'FAILED')),
    
    -- Timestamps
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NULL,
    last_accessed_at TIMESTAMP NULL
);

-- Add table and column comments for documentation
COMMENT ON TABLE files IS 'File metadata and storage information for Secure File Vault v1.0.0';

-- Core field comments
COMMENT ON COLUMN files.id IS 'Unique file identifier (UUID format)';
COMMENT ON COLUMN files.original_name IS 'Original filename as uploaded by user';
COMMENT ON COLUMN files.filename IS 'Sanitized filename used for storage';
COMMENT ON COLUMN files.path IS 'Storage path or reference identifier';

-- Property comments  
COMMENT ON COLUMN files.size IS 'File size in bytes (must be >= 0)';
COMMENT ON COLUMN files.content_type IS 'HTTP Content-Type header value';
COMMENT ON COLUMN files.mime_type IS 'MIME type for file validation';

-- Storage comments
COMMENT ON COLUMN files.s3_bucket IS 'AWS S3 bucket name where file is stored';
COMMENT ON COLUMN files.s3_key IS 'S3 object key/path for file location';
COMMENT ON COLUMN files.s3_etag IS 'S3 ETag for file integrity verification';

-- Security comments
COMMENT ON COLUMN files.checksum IS 'SHA-256 checksum for file integrity verification';
COMMENT ON COLUMN files.access_policy IS 'Access control policy (PRIVATE/PUBLIC_READ/OWNER_ONLY)';
COMMENT ON COLUMN files.is_public IS 'Quick flag for public access (denormalized for performance)';

-- Ownership comments
COMMENT ON COLUMN files.owner IS 'UUID of the user who owns this file';
COMMENT ON COLUMN files.uploader IS 'UUID of the user who uploaded this file (may differ from owner)';
COMMENT ON COLUMN files.access_count IS 'Total number of times file has been accessed';
COMMENT ON COLUMN files.download_count IS 'Total number of times file has been downloaded';

-- Status comments
COMMENT ON COLUMN files.upload_status IS 'Current upload/processing status';

-- Timestamp comments
COMMENT ON COLUMN files.created_at IS 'Timestamp when file record was created';
COMMENT ON COLUMN files.updated_at IS 'Timestamp when file record was last updated (auto-updated)';
COMMENT ON COLUMN files.expires_at IS 'Timestamp when file expires (NULL = never expires)';
COMMENT ON COLUMN files.last_accessed_at IS 'Timestamp of last file access/download';
