package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListFiles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of files"})
}

func GetFile(c *gin.Context) {
	fileID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "File details", "id": fileID})
}

func DeleteFile(c *gin.Context) {
	fileID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "File deleted", "id": fileID})
}
