package core

import "context"

type Clipboard interface {
	RegisterListener(ctx context.Context, s string) error
	GetClipboard(context.Context) ([]string, error)
	SetClipboardValue(ctx context.Context, s string) error
}
type Client struct {
	clipboard Clipboard
}
type ClientOptions func(*Client) error

func New(opts ...ClientOptions) *Client {
	c := Client{}
	for _, f := range opts {
		f(&c)
	}
	return &c
}
func WithClipboardListener(clipboard Clipboard) ClientOptions {
	return func(c *Client) error {
		c.clipboard = clipboard
		return nil
	}
}
func (c *Client) RegisterListener() error {
	return c.clipboard.RegisterListener(context.Background(), "Service1")
}
