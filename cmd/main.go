package main

import (
	"fmt"

	"github.com/huypq02/secure-file-vault/internal/config"
	"github.com/huypq02/secure-file-vault/internal/infrastructure/db"
	"github.com/huypq02/secure-file-vault/internal/infrastructure/external"
	"github.com/huypq02/secure-file-vault/internal/infrastructure/scheduler"
	"github.com/huypq02/secure-file-vault/internal/infrastructure/storage"
	"github.com/huypq02/secure-file-vault/internal/interface/handler"
	"github.com/huypq02/secure-file-vault/internal/interface/router"
	"github.com/huypq02/secure-file-vault/internal/usecase/file"
)

func main() {
	fmt.Println("Secure File Vault - Entry point")

	// Export env for database and storage service
	cfg := config.NewConfig()

	// Infrastructure layer
	// Initialize database connection
	dbConfig, err := db.NewConnection(cfg.GetDatabaseConfig())
	if err != nil {
		fmt.Printf("Failed to create database connection: %v\n", err)
		return
	}
	// Check database connection
	if err := dbConfig.Connect(); err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}
	// Initialize the file repository
	fileRepo := db.NewFileRepository(dbConfig)
	// Initialize the s3 service
	s3Service, err := storage.NewS3Storage(cfg.GetStorageConfig())
	if err != nil {
		fmt.Printf("Failed to create storage service: %v\n", err)
		return
	}
	// Initialize the storage services
	storageService := storage.NewServiceStorage(s3Service)
	if storageService == nil {
		fmt.Println("Failed to create storage service")
		return
	}
	// Initialize ID service
	idService := external.NewIDService()
	if idService == nil {
		fmt.Println("Failed to create ID service")
		return
	}
	// Initialize scheduler
	scheduler := scheduler.NewScheduler(fileRepo, storageService)

	// Application logic layer
	// Initialize use cases
	downloadFileUsecase := file.NewDownloadFileUsecase(fileRepo, storageService)
	uploadFileUsecase := file.NewUploadFileUsecase(fileRepo, storageService, idService)

	// Interfaces layer
	// Create file handler
	fileHandler := handler.NewFileHandler(downloadFileUsecase, uploadFileUsecase)
	// Initialize the Gin router and register routes
	r := router.NewRouter(fileHandler)
	if r == nil {
		fmt.Println("Failed to create router")
		return
	}

	// Start the scheduler
	scheduler.Start()
	// Start the server
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
	// Print success message
	fmt.Println("Server is running on port 8080")
}
