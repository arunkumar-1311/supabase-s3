// Package storage wraps the supabase-go Storage client with a smaller,
// bucket-scoped API for uploading, downloading, listing, and managing files.
package storage

import (
	supabase "github.com/supabase-community/supabase-go"

	"github.com/arunkumar-1311/supabase-s3/internal/config"
)

// Client is a thin, bucket-scoped wrapper around the Supabase Storage API.
type Client struct {
	sb     *supabase.Client
	bucket string
}

// New creates a storage Client from the given config, targeting cfg.BucketName
// by default for file operations.
func New(cfg *config.Config) (*Client, error) {
	sb, err := supabase.NewClient(cfg.SupabaseURL, cfg.SupabaseKey, &supabase.ClientOptions{})
	if err != nil {
		return nil, err
	}

	return &Client{
		sb:     sb,
		bucket: cfg.BucketName,
	}, nil
}
