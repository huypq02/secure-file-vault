package external

import (
	"github.com/google/uuid"
	"github.com/huypq02/secure-file-vault/internal/domain"
)

type idService struct{}

func NewIDService() domain.IDService {
	return &idService{}
}

func (s *idService) Generate() string {
	return uuid.New().String()
}
