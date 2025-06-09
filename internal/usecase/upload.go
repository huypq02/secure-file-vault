package usecase

import "secure-file-vault/internal/domain"

type FileUploader interface {
    UploadFile(meta domain.FileMetadata, data []byte) (string, error)
}
