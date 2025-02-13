package handlers

import (
	"app/internal/storage"
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

var allowedExtensions = map[string]bool{
	".csv":  true,
	".pdf":  true,
	".docx": true,
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	ext := filepath.Ext(file.Filename)
	if !allowedExtensions[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File type not allowed"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	fileID := uuid.New().String()
	objectName := fmt.Sprintf("%s%s", fileID, ext)

	// метаданные
	ctx := context.Background()
	_, err = storage.MinioClient.PutObject(ctx, "file-storage", objectName, src, file.Size, minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})

		return

	}
	log.Println("Storage URL:", objectName)
	log.Println("Filename:", file.Filename, "Filetype:", ext, "Size:", file.Size)

	query := `INSERT INTO files (id, filename, filetype, filesize, storage_url) VALUES ($1, $2, $3, $4, $5)`
	_, err = storage.DB.Exec(query, fileID, file.Filename, ext, file.Size, objectName)
	if err != nil {
		log.Println("Error inserting into database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file metadata"})
		return
	}
}
