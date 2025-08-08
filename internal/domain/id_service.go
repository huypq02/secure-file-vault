package domain

// Main ID service interface
type IDService interface {
	Generate() string
	// GenerateWithPrefix(prefix string) string
	// ValidateID(id string) error
}

// ID generation strategies
type IDGenerationStrategy string

const (
	UUIDStrategy   IDGenerationStrategy = "uuid"
	ULIDStrategy   IDGenerationStrategy = "ulid"
	NanoIDStrategy IDGenerationStrategy = "nanoid"
	CustomStrategy IDGenerationStrategy = "custom"
)

// ID metadata for analysis
type IDMetadata struct {
	Strategy  IDGenerationStrategy `json:"strategy"`
	Timestamp int64                `json:"timestamp,omitempty"`
	Prefix    string               `json:"prefix,omitempty"`
	Length    int                  `json:"length"`
	IsSecure  bool                 `json:"isSecure"`
}
