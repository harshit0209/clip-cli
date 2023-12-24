package clipboard

import (
	"context"

	"github.com/atotto/clipboard"
)

func (c *Client) SetClipboardValue(ctx context.Context, s string) error {
	c.FlagClip = s
	err := clipboard.WriteAll(s)
	if err != nil {
		return err
	}
	return nil
}
