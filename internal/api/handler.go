// Package api exposes Supabase Storage operations over HTTP using Gin.
package api

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/arunkumar-1311/supabase-s3/internal/storage"
)

// Handler holds dependencies shared by the API's route handlers.
type Handler struct {
	storage *storage.Client
}

// NewHandler creates a Handler backed by the given storage client.
func NewHandler(storageClient *storage.Client) *Handler {
	return &Handler{storage: storageClient}
}

// keyParam extracts the object key from a wildcard route param (":key" or
// "*key"), stripping the leading slash Gin includes for "*" params.
func keyParam(c *gin.Context) string {
	key := c.Param("key")
	return strings.TrimPrefix(key, "/")
}
