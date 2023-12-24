package core

import (
	"context"
)

func (c *Client) SetClipboardValue(ctx context.Context, s string) error {
	return c.clipboard.SetClipboardValue(ctx, s)
}
