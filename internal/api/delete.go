package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteFile handles DELETE /files/*key
func (h *Handler) DeleteFile(c *gin.Context) {
	key := keyParam(c)
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "key is required"})
		return
	}

	if _, err := h.storage.Remove(key); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"key": key, "deleted": true})
}
