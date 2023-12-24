package core

import "context"

func (c *Client) GetClipboard(ctx context.Context) ([]string, error) {
	return c.clipboard.GetClipboard(ctx)
}
