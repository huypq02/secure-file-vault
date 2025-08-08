package db

import (
	"fmt"

	"github.com/huypq02/secure-file-vault/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type connection struct {
	DB     *gorm.DB
	config *domain.DatabaseConfig
}

func NewConnection(cfg *domain.DatabaseConfig) (*connection, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode, cfg.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &connection{
		DB:     db,
		config: cfg,
	}, nil
}

func (c *connection) Connect() error {
	if c.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	return nil
}
