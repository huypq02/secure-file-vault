package config

import (
	"fmt"
	"strings"

	"github.com/huypq02/secure-file-vault/internal/domain"
	"github.com/spf13/viper"
)

type Config struct {
	Database Database `mapstructure:"database"`
	Storage  Storage  `mapstructure:"storage"`
}

func NewConfig() domain.ConfigProvider {
	cfg := &Config{}
	return cfg.SetConfigFile()
}

func (c *Config) SetConfigFile() domain.ConfigProvider {
	v := viper.New()

	// Set default configuration paths
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AutomaticEnv()                                   // Enable reading from environment variables
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // database.host -> DATABASE_HOST

	// Try to read the config file
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Warning: error reading config file: %v\n", err)
	}

	// Unmarshal directly into the Config struct
	if err := v.Unmarshal(c); err != nil {
		panic(fmt.Errorf("unable to decode config into struct: %w", err))
	}

	fmt.Printf("Loaded config: %+v\n", c)
	return c
}

func (c *Config) GetDatabaseConfig() *domain.DatabaseConfig {
	return &domain.DatabaseConfig{
		Host:     c.Database.Host,
		Port:     c.Database.Port,
		User:     c.Database.User,
		Password: c.Database.Password,
		Name:     c.Database.Name,
		SSLMode:  c.Database.SSLMode,
		TimeZone: c.Database.TimeZone,
	}
}

func (c *Config) GetStorageConfig() *domain.StorageConfig {
	return &domain.StorageConfig{
		Provider:        c.Storage.Provider,
		Bucket:          c.Storage.Bucket,
		Region:          c.Storage.Region,
		AccessKeyID:     c.Storage.Credentials.AccessKeyID,
		SecretAccessKey: c.Storage.Credentials.SecretAccessKey,
		Endpoint:        c.Storage.Endpoint,
		ForcePathStyle:  c.Storage.ForcePathStyle,
	}
}
