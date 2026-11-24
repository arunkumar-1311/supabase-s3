package storage

import (
	"io"

	storage_go "github.com/supabase-community/storage-go"
)

// UploadOptions controls optional metadata sent with an upload.
type UploadOptions struct {
	ContentType  string
	CacheControl string
	Upsert       bool
}

func (o UploadOptions) toFileOptions() storage_go.FileOptions {
	opts := storage_go.FileOptions{}
	if o.ContentType != "" {
		opts.ContentType = &o.ContentType
	}
	if o.CacheControl != "" {
		opts.CacheControl = &o.CacheControl
	}
	opts.Upsert = &o.Upsert
	return opts
}

// Upload adds a new file at path in the client's bucket. Use Update to
// overwrite a file that already exists.
func (c *Client) Upload(path string, data io.Reader, opts UploadOptions) (storage_go.FileUploadResponse, error) {
	return c.sb.Storage.UploadFile(c.bucket, path, data, opts.toFileOptions())
}

// Update replaces an existing file at path with new data.
func (c *Client) Update(path string, data io.Reader, opts UploadOptions) (storage_go.FileUploadResponse, error) {
	return c.sb.Storage.UpdateFile(c.bucket, path, data, opts.toFileOptions())
}
