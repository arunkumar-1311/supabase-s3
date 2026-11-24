package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const defaultExpirySeconds = 300

// PresignedURL handles GET /presigned-url/*key?expires_in=<seconds>
func (h *Handler) PresignedURL(c *gin.Context) {
	key := keyParam(c)
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "key is required"})
		return
	}

	expiresIn := defaultExpirySeconds
	if raw := c.Query("expires_in"); raw != "" {
		v, err := strconv.Atoi(raw)
		if err != nil || v <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "expires_in must be a positive integer"})
			return
		}
		expiresIn = v
	}

	url, err := h.storage.SignedURL(key, expiresIn)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"key":        key,
		"url":        url,
		"expires_in": expiresIn,
	})
}
