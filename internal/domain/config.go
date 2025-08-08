package domain

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

type StorageConfig struct {
	Provider        string
	Region          string
	Bucket          string
	AccessKeyID     string
	SecretAccessKey string
	Endpoint        string
	ForcePathStyle  bool
}

type ConfigProvider interface {
	GetDatabaseConfig() *DatabaseConfig
	GetStorageConfig() *StorageConfig
}
