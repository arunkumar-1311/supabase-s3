package storage

// Download fetches the raw bytes of the file at path in the client's bucket.
func (c *Client) Download(path string) ([]byte, error) {
	return c.sb.Storage.DownloadFile(c.bucket, path)
}
