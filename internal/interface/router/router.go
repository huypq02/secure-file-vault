package router

import (
	"github.com/huypq02/secure-file-vault/internal/domain"
	"github.com/huypq02/secure-file-vault/internal/interface/handler"
	"github.com/huypq02/secure-file-vault/internal/interface/middleware"

	// Import the gin package for routing
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	file   *handler.FileHandler
}

func NewRouter(file *handler.FileHandler) domain.Router {
	return &Router{
		engine: gin.Default(),
		file:   file,
	}
}

func (r *Router) RegisterRoutes() {
	// Set up the router with middleware
	r.engine.Use(middleware.AuthMiddleware())

	// File routes
	v1 := r.engine.Group("/api/v1")
	{
		fileGroup := v1.Group("file")
		{
			fileGroup.POST("/upload", r.file.UploadFileHandler)
			fileGroup.GET("/download/:fileID", r.file.DownloadFileHandler)
			// fileGroup.GET("/list", file.ListFilesHandler)
		}
	}
}

func (r *Router) Run(addr string) error {
	r.RegisterRoutes()
	return r.engine.Run(addr)
}
