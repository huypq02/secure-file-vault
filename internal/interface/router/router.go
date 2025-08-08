package router

import (
	"github.com/huypq02/secure-file-vault/internal/interface/handler"
	"github.com/huypq02/secure-file-vault/internal/interface/middleware"

	// Import the gin package for routing
	"github.com/gin-gonic/gin"
)

func NewRouter(file *handler.FileHandler) *gin.Engine {
	r := gin.Default()

	// Set up the router with middleware
	r.Use(middleware.AuthMiddleware())

	// File routes
	v1 := r.Group("/api/v1")
	{
		fileGroup := v1.Group("file")
		{
			fileGroup.POST("/upload", file.UploadFileHandler)
			fileGroup.GET("/download/:fileID", file.DownloadFileHandler)
			// fileGroup.GET("/list", file.ListFilesHandler)
		}
	}

	return r
}
