// interface/handler/file_handler.go
package handler

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/huypq02/secure-file-vault/internal/domain"
	httpDTO "github.com/huypq02/secure-file-vault/internal/interface/dto"
	usecaseDTO "github.com/huypq02/secure-file-vault/internal/usecase/dto"
	"github.com/huypq02/secure-file-vault/internal/usecase/file"
)

type FileHandler struct {
	downloadFileUsecase file.DownloadFileUsecase
	uploadFileUsecase   file.UploadFileUsecase
}

func NewFileHandler(
	downloadFileUsecase file.DownloadFileUsecase,
	uploadFileUsecase file.UploadFileUsecase,
) *FileHandler {
	return &FileHandler{
		downloadFileUsecase: downloadFileUsecase,
		uploadFileUsecase:   uploadFileUsecase,
	}
}

// DownloadFile with proper DTO separation
func (f *FileHandler) DownloadFileHandler(c *gin.Context) {
	ctx := c.Request.Context()
	downloadCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	fileID := c.Param("fileID")
	if fileID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fileID is required"})
		return
	}

	// Call use case - get internal DTO
	usecaseResp, err := f.downloadFileUsecase.DownloadFile(downloadCtx, fileID)
	if err != nil {
		f.handleError(c, err)
		return
	}

	if usecaseResp == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Convert to HTTP DTO with encoding
	httpResp := f.convertToHTTPDownloadResponse(usecaseResp)

	c.JSON(http.StatusOK, httpResp)
}

// UploadFile with proper DTO separation
func (f *FileHandler) UploadFileHandler(c *gin.Context) {
	ctx := c.Request.Context()
	uploadCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Parse HTTP request
	httpReq, err := f.parseUploadRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert to use case DTO
	usecaseReq := f.convertToUsecaseUploadRequest(httpReq)

	// Call use case
	usecaseResp, err := f.uploadFileUsecase.UploadFile(uploadCtx, usecaseReq)
	if err != nil {
		f.handleError(c, err)
		return
	}

	// Convert to HTTP response DTO
	httpResp := f.convertToHTTPUploadResponse(usecaseResp)

	c.JSON(http.StatusCreated, httpResp)
}

// Conversion methods
func (f *FileHandler) convertToHTTPDownloadResponse(usecaseResp *usecaseDTO.DownloadFileResponse) *httpDTO.HTTPDownloadFileResponse {
	// Encoding happens here, not mutating original DTO
	encodedContent := base64.StdEncoding.EncodeToString(usecaseResp.Data)

	return &httpDTO.HTTPDownloadFileResponse{
		Status: true,
		Data: httpDTO.HTTPFileDownloadData{
			Filename:    usecaseResp.Filename,
			Size:        usecaseResp.Size,
			ContentType: usecaseResp.ContentType,
			Content:     encodedContent, // Encoded content
			Encoding:    "base64",
		},
		Message: usecaseResp.Message,
	}
}

func (f *FileHandler) parseUploadRequest(c *gin.Context) (*httpDTO.HTTPUploadFileRequest, error) {
	// Parse multipart form
	fileBody, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("failed to get file from form: %w", err)
	}
	defer fileBody.Close()

	fileData, err := io.ReadAll(fileBody)
	if err != nil {
		return nil, fmt.Errorf("failed to read file data: %w", err)
	}

	return &httpDTO.HTTPUploadFileRequest{
		OriginalName: fileHeader.Filename,
		Filename:     c.PostForm("filename"),
		Description:  c.PostForm("description"),
		FileData:     fileData, // Raw file data
		Size:         int64(len(fileData)),
		ContentType:  fileHeader.Header.Get("Content-Type"),
	}, nil
}

func (f *FileHandler) convertToUsecaseUploadRequest(httpReq *httpDTO.HTTPUploadFileRequest) *usecaseDTO.UploadFileRequest {
	return &usecaseDTO.UploadFileRequest{
		FileItem: usecaseDTO.FileItem{
			OriginalName: httpReq.OriginalName,
			Filename:     httpReq.Filename,
			Size:         httpReq.Size,
			ContentType:  httpReq.ContentType,
			Data:         httpReq.FileData,
		},
		Description: httpReq.Description,
	}
}

func (f *FileHandler) convertToHTTPUploadResponse(usecaseResp *usecaseDTO.UploadFileResponse) *httpDTO.HTTPUploadFileResponse {
	return &httpDTO.HTTPUploadFileResponse{
		Status: true,
		Data: httpDTO.HTTPFileUploadData{
			FileID:   usecaseResp.FileID,
			Filename: usecaseResp.Filename,
			Size:     usecaseResp.Size,
		},
		Message: usecaseResp.Message,
	}
}

func (f *FileHandler) handleError(c *gin.Context, err error) {
	// Proper error handling with different HTTP status codes
	switch {
	case errors.Is(err, domain.ErrFileNotFound):
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  "File not found",
		})
	case errors.Is(err, domain.ErrAccessDenied):
		c.JSON(http.StatusForbidden, gin.H{
			"status": false,
			"error":  "Access denied",
		})
	case errors.Is(err, domain.ErrFileTooLarge):
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{
			"status": false,
			"error":  "File too large",
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Internal server error",
		})
	}
}
