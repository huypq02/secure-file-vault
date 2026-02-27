//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"github.com/huypq02/secure-file-vault/internal/config"
	"github.com/huypq02/secure-file-vault/internal/domain"
	"github.com/huypq02/secure-file-vault/internal/infrastructure/db"
	"github.com/huypq02/secure-file-vault/internal/infrastructure/external"
	"github.com/huypq02/secure-file-vault/internal/infrastructure/scheduler"
	"github.com/huypq02/secure-file-vault/internal/infrastructure/storage"
	"github.com/huypq02/secure-file-vault/internal/interface/handler"
	"github.com/huypq02/secure-file-vault/internal/interface/router"
	"github.com/huypq02/secure-file-vault/internal/usecase/file"
)

func InitializeApp() (domain.Application, error) {
	wire.Build(
		// Configuration
		config.NewConfig,
		config.ProvideDatabaseConfig,
		config.ProvideStorageConfig,
		// Infrastructure layer
		db.NewConnection,
		db.NewFileRepository,
		storage.NewS3Client,
		storage.NewS3StorageService,
		external.NewIDService,
		scheduler.NewScheduler,
		// Application logic layer
		file.NewDownloadFileUsecase,
		file.NewUploadFileUsecase,
		// Interfaces layer
		handler.NewFileHandler,
		router.NewRouter,

		NewApp,
	)
	return nil, nil
}
