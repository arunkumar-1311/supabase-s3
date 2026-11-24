package storage

import (
	storage_go "github.com/supabase-community/storage-go"
)

// Remove deletes the given paths from the client's bucket.
func (c *Client) Remove(paths ...string) ([]storage_go.FileUploadResponse, error) {
	return c.sb.Storage.RemoveFile(c.bucket, paths)
}

// Move relocates a file from srcPath to dstPath within the client's bucket.
func (c *Client) Move(srcPath, dstPath string) (storage_go.FileUploadResponse, error) {
	return c.sb.Storage.MoveFile(c.bucket, srcPath, dstPath)
}
