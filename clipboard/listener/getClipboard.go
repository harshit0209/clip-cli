package clipboard

import "context"

func (c *Client) GetClipboard(ctx context.Context) ([]string, error) {
	return c.Clips, nil
}
