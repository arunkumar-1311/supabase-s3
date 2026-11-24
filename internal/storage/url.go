package storage

// PublicURL returns the public URL for path in the client's bucket. The
// bucket must be marked public for this URL to be accessible.
func (c *Client) PublicURL(path string) string {
	return c.sb.Storage.GetPublicUrl(c.bucket, path).SignedURL
}

// SignedURL returns a time-limited download URL for path, valid for
// expiresInSeconds seconds.
func (c *Client) SignedURL(path string, expiresInSeconds int) (string, error) {
	resp, err := c.sb.Storage.CreateSignedUrl(c.bucket, path, expiresInSeconds)
	if err != nil {
		return "", err
	}
	return resp.SignedURL, nil
}
