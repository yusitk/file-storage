package main

import (
	"app/internal/handlers"
	"app/internal/storage"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	storage.InitDB()

	storage.InitMinIO()

	router := gin.Default()

	router.POST("/upload", handlers.UploadFile)
	router.GET("/files", handlers.ListFiles)
	router.GET("/files/:id", handlers.GetFile)
	router.DELETE("/files/:id", handlers.DeleteFile)
	router.GET("/download/:id", handlers.DownloadFile)

	log.Println("Starting API server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
