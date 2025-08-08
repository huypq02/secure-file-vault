package external

import "encoding/base64"

type cryptoService struct{}

func NewCryptoService() *cryptoService {
	return &cryptoService{}
}

func (s *cryptoService) Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func (s *cryptoService) Base64Decode(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}
