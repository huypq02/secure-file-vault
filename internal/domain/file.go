package domain

type FileMetadata struct {
    ID        string
    Filename  string
    Uploader  string
    ExpiresAt int64
    AccessCount int
}
