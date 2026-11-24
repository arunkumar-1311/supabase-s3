// Package config loads application configuration from environment variables.
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the settings needed to talk to Supabase Storage.
type Config struct {
	SupabaseURL string
	SupabaseKey string
	BucketName  string
}

// Load reads a .env file (if present) and builds a Config from environment
// variables. A missing .env file is not an error since env vars may already
// be set by the shell or deployment environment.
func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		SupabaseURL: os.Getenv("SUPABASE_URL"),
		SupabaseKey: os.Getenv("SUPABASE_KEY"),
		BucketName:  os.Getenv("SUPABASE_BUCKET"),
	}

	if cfg.SupabaseURL == "" {
		return nil, fmt.Errorf("SUPABASE_URL is required")
	}
	if cfg.SupabaseKey == "" {
		return nil, fmt.Errorf("SUPABASE_KEY is required")
	}
	if cfg.BucketName == "" {
		return nil, fmt.Errorf("SUPABASE_BUCKET is required")
	}

	return cfg, nil
}
