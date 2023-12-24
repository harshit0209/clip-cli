package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Corer interface {
	// RegisterListener(ctx context.Context, s string) error
	// GetClipboard(context.Context) ([]string, error)
	// SetClipboardValue(ctx context.Context, s string) error
}
type Client struct {
	core        Corer
	port        string
	serviceName string
	echo        *echo.Echo
}
type ClientOptions func(*Client) error

func New(opts ...ClientOptions) *Client {
	c := Client{}
	for _, opt := range opts {
		opt(&c)
	}
	return &c
}
func WithCore(co Corer) ClientOptions {
	return func(c *Client) error {
		c.core = co
		return nil
	}
}
func WithServiceName(s string) ClientOptions {
	return func(c *Client) error {
		c.serviceName = s
		return nil
	}
}
func WithPort(s string) ClientOptions {
	return func(c *Client) error {
		c.port = s
		return nil
	}
}

func WithEchoFramework(e *echo.Echo) ClientOptions {
	return func(c *Client) error {
		c.echo = e
		return nil
	}

}
func WithRoutesAndMiddlewares() ClientOptions {
	return func(c *Client) error {
		c.echo.Use(middleware.Logger())
		c.echo.Use(middleware.Recover())
		c.echo.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, Echo with Middleware!")
		})

		return nil
	}
}
func (c *Client) Start() error {
	return c.echo.Start(c.port)
}
