package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/arunkumar-1311/supabase-s3/internal/storage"
)

// UploadFile handles POST /files
// Expects a multipart form with a "file" field and an optional "key" field
// (defaults to the uploaded filename).
func (h *Handler) UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	key := c.PostForm("key")
	if key == "" {
		key = fileHeader.Filename
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	fmt.Println("key", key)
	resp, err := h.storage.Upload(key, file, storage.UploadOptions{
		ContentType: fileHeader.Header.Get("Content-Type"),
		Upsert:      true,
	})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"key":     key,
		"message": resp.Message,
	})
}
