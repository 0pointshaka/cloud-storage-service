package routes

import (
	"github.com/0pointshaka/cloud-storage-service/backend/handlers"
	"github.com/0pointshaka/cloud-storage-service/backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Health check
	router.GET("/health", handlers.HealthCheck)

	// Auth routes
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/register", handlers.Register)
		authGroup.POST("/login", handlers.Login)
	}

	// File routes
	fileGroup := router.Group("/api/files")
	fileGroup.Use(middleware.AuthMiddleware())
	{
		fileGroup.POST("/upload", handlers.UploadFile)
		fileGroup.GET("/list", handlers.ListFiles)
		fileGroup.GET("/download/:id", handlers.DownloadFile)
		fileGroup.DELETE("/:id", handlers.DeleteFile)
	}

	// Sharing routes
	sharingGroup := router.Group("/api/share")
	sharingGroup.Use(middleware.AuthMiddleware())
	{
		sharingGroup.POST("/:file_id", handlers.ShareFile)
		sharingGroup.GET("/list", handlers.ListSharedFiles)
		sharingGroup.DELETE("/:share_id", handlers.UnshareFile)
	}
}
