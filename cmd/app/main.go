// Command app runs an HTTP API exposing S3-style Supabase Storage
// operations: upload, delete, and presigned URL generation.
package main

import (
	"log"
	"os"

	"github.com/arunkumar-1311/supabase-s3/internal/api"
	"github.com/arunkumar-1311/supabase-s3/internal/config"
	"github.com/arunkumar-1311/supabase-s3/internal/storage"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	client, err := storage.New(cfg)
	if err != nil {
		log.Fatalf("storage client: %v", err)
	}

	handler := api.NewHandler(client)
	router := api.NewRouter(handler)

	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("listening on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("server: %v", err)
	}
}
