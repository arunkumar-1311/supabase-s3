package storage

import (
	storage_go "github.com/supabase-community/storage-go"
)

// BucketOptions controls visibility and constraints when creating or
// updating a bucket.
type BucketOptions struct {
	Public           bool
	FileSizeLimit    string
	AllowedMimeTypes []string
}

func (o BucketOptions) toStorageOptions() storage_go.BucketOptions {
	return storage_go.BucketOptions{
		Public:           o.Public,
		FileSizeLimit:    o.FileSizeLimit,
		AllowedMimeTypes: o.AllowedMimeTypes,
	}
}

// CreateBucket creates a new bucket with the given id.
func (c *Client) CreateBucket(id string, opts BucketOptions) (storage_go.Bucket, error) {
	return c.sb.Storage.CreateBucket(id, opts.toStorageOptions())
}

// GetBucket retrieves details of an existing bucket.
func (c *Client) GetBucket(id string) (storage_go.Bucket, error) {
	return c.sb.Storage.GetBucket(id)
}

// ListBuckets returns all buckets in the project.
func (c *Client) ListBuckets() ([]storage_go.Bucket, error) {
	return c.sb.Storage.ListBuckets()
}

// DeleteBucket removes an existing bucket. The bucket must be empty.
func (c *Client) DeleteBucket(id string) error {
	_, err := c.sb.Storage.DeleteBucket(id)
	return err
}
