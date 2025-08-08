package external

import (
	"github.com/google/uuid"
)

type idService struct{}

func NewIDService() *idService {
	return &idService{}
}

func (s *idService) Generate() string {
	return uuid.New().String()
}
