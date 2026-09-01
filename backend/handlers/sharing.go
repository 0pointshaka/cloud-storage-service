package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/0pointshaka/cloud-storage-service/backend/config"
	"github.com/0pointshaka/cloud-storage-service/backend/models"
	"github.com/gin-gonic/gin"
)

type ShareRequest struct {
	Username string `json:"username" binding:"required"`
}

func ShareFile(c *gin.Context) {
	userID := c.GetUint("user_id")
	fileID := c.Param("file_id")

	var req ShareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify file belongs to user
	var file models.File
	if result := config.DB.Where("id = ? AND user_id = ?", fileID, userID).First(&file); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Find user to share with
	var targetUser models.User
	if result := config.DB.Where("username = ?", req.Username).First(&targetUser); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Create shared file record
	sharedFile := models.SharedFile{
		FileID:    file.ID,
		SharedBy:  userID,
		SharedTo:  targetUser.ID,
		CreatedAt: time.Now().Unix(),
	}

	if result := config.DB.Create(&sharedFile); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to share file"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "File shared successfully", "share_id": sharedFile.ID})
}

func ListSharedFiles(c *gin.Context) {
	userID := c.GetUint("user_id")

	var sharedFiles []struct {
		FileID   uint   `json:"file_id"`
		Filename string `json:"filename"`
		Size     int64  `json:"size"`
		SharedBy uint   `json:"shared_by"`
		CreatedAt int64  `json:"created_at"`
	}

	if result := config.DB.Table("shared_files").
		Select("files.id as file_id, files.filename, files.size, shared_files.shared_by, shared_files.created_at").
		Joins("JOIN files ON shared_files.file_id = files.id").
		Where("shared_files.shared_to = ?", userID).
		Find(&sharedFiles); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shared files"})
		return
	}

	c.JSON(http.StatusOK, sharedFiles)
}

func UnshareFile(c *gin.Context) {
	userID := c.GetUint("user_id")
	shareID := c.Param("share_id")

	var sharedFile models.SharedFile
	if result := config.DB.Where("id = ? AND shared_by = ?", shareID, userID).First(&sharedFile); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Share not found"})
		return
	}

	if result := config.DB.Delete(&sharedFile); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unshare file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File unshared successfully"})
}
