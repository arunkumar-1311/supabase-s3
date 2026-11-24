package storage

import (
	storage_go "github.com/supabase-community/storage-go"
)

// ListOptions controls pagination and sorting for List.
type ListOptions struct {
	Limit  int
	Offset int
}

// List returns the files under prefix in the client's bucket.
func (c *Client) List(prefix string, opts ListOptions) ([]storage_go.FileObject, error) {
	return c.sb.Storage.ListFiles(c.bucket, prefix, storage_go.FileSearchOptions{
		Limit:  opts.Limit,
		Offset: opts.Offset,
	})
}
