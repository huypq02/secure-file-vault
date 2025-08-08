package domain

type FileRepository interface {
	Save(file *FileMetadata) error
	FindByID(id string) (*FileMetadata, error)
	// Delete(id string) error
	// List(limit, offset int) ([]*FileMetadata, error)
	Update(file *FileMetadata) error
}
