package dto

// DownloadFileResponse represents the response structure for file download
type DownloadFileResponse struct {
	Filename    string `json:"filename"`
	Size        int64  `json:"size"`
	ContentType string `json:"contentType"`
	ExpiresAt   string `json:"expiresAt,omitempty"`
	Data        []byte `json:"data,omitempty"`
	Message     string `json:"message,omitempty"`
}
