package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func UploadFileHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Upload endpoint"})
}
