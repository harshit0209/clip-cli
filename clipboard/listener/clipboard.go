package clipboard

type Client struct {
	ListenerName string
}
type ClientOptions func(*Client) error

func New(opts ...ClientOptions) *Client {
	c := Client{}
	for _, f := range opts {
		f(&c)
	}
	return &c
}
func WithClipboardListenerName(listenerName string) ClientOptions {
	return func(c *Client) error {
		c.ListenerName = listenerName
		return nil
	}
}
