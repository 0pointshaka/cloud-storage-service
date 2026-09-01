package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/0pointshaka/cloud-storage-service/backend/config"
	"github.com/0pointshaka/cloud-storage-service/backend/models"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	userID := c.GetUint("user_id")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	// Create file path
	userDir := filepath.Join("uploads", strconv.FormatUint(uint64(userID), 10))
	if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user directory"})
		return
	}

	filePath := filepath.Join(userDir, file.Filename)

	// Save file
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Get file info
	fileInfo, _ := os.Stat(filePath)

	// Store in database
	newFile := models.File{
		UserID:    userID,
		Filename:  file.Filename,
		Path:      filePath,
		Size:      fileInfo.Size(),
		MimeType:  file.Header.Get("Content-Type"),
		CreatedAt: time.Now().Unix(),
	}

	if result := config.DB.Create(&newFile); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store file metadata"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "File uploaded successfully", "file_id": newFile.ID})
}

func ListFiles(c *gin.Context) {
	userID := c.GetUint("user_id")

	var files []models.File
	if result := config.DB.Where("user_id = ?", userID).Find(&files); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}

	c.JSON(http.StatusOK, files)
}

func DownloadFile(c *gin.Context) {
	userID := c.GetUint("user_id")
	fileID := c.Param("id")

	var file models.File
	if result := config.DB.Where("id = ? AND user_id = ?", fileID, userID).First(&file); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	if _, err := os.Stat(file.Path); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found on disk"})
		return
	}

	c.FileAttachment(file.Path, file.Filename)
}

func DeleteFile(c *gin.Context) {
	userID := c.GetUint("user_id")
	fileID := c.Param("id")

	var file models.File
	if result := config.DB.Where("id = ? AND user_id = ?", fileID, userID).First(&file); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Delete file from disk
	if err := os.Remove(file.Path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file"})
		return
	}

	// Delete from database
	if result := config.DB.Delete(&file); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}
