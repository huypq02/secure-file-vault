package dto

type HTTPDownloadFileResponse struct {
	Status  bool                 `json:"status"`
	Data    HTTPFileDownloadData `json:"data"`
	Message string               `json:"message,omitempty"`
}

type HTTPFileDownloadData struct {
	Filename     string `json:"filename"`
	Size         int64  `json:"size"`
	Content      string `json:"content"`
	ContentType  string `json:"contentType,omitempty"`
	Checksum     string `json:"checksum,omitempty"`
	DownloadedAt string `json:"downloadedAt,omitempty"`
	Encoding     string `json:"encoding,omitempty"`
}

type HTTPUploadFileRequest struct {
	OriginalName string `form:"originalName" json:"originalName"`
	Filename     string `form:"filename" json:"filename"` // The file name identified by the owner
	Description  string `form:"description" json:"description"`
	FileData     []byte `form:"fileData" json:"fileData"` // Raw file data
	Size         int64  `form:"size" json:"size"`
	ContentType  string `form:"contentType" json:"contentType"`
}

type HTTPUploadFileResponse struct {
	Status  bool               `json:"status"`
	Data    HTTPFileUploadData `json:"data"`
	Message string             `json:"message"`
}

type HTTPFileUploadData struct {
	FileID      string `json:"fileID"`
	Filename    string `json:"filename"`
	Size        int64  `json:"size"`
	UploadedAt  string `json:"uploadedAt,omitempty"`
	DownloadURL string `json:"downloadUrl,omitempty"`
}
