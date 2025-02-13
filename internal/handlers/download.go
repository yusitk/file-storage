package handlers

import (
	"app/internal/storage"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func DownloadFile(c *gin.Context) {
	fileID := c.Param("id")

	//мета
	var filename, storageURL string
	query := `SELECT filename, storage_url FROM files WHERE id = $1`
	err := storage.DB.QueryRow(query, fileID).Scan(&filename, &storageURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	ctx := context.Background()
	object, err := storage.MinioClient.GetObject(ctx, "file-storage", storageURL, minio.GetObjectOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch file"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.DataFromReader(http.StatusOK, -1, "application/octet-stream", object, nil)
}
