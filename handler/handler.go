package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Corer interface {
	GetClipboard(context.Context) ([]string, error)
	SetClipboardValue(ctx context.Context, s string) error
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
		c.echo.GET("/list", c.getClips)
		c.echo.GET("/list:index", c.getClips)

		return nil
	}
}
func (c *Client) Start() error {
	return c.echo.Start(c.port)
}

func (c *Client) getClips(echoCtx echo.Context) error {
	clps, err := c.core.GetClipboard(context.Background())
	if err != nil {
		return err
	}

	if index := echoCtx.Param("index"); index != "" {
		indexNum, err := strconv.Atoi(index[1:])
		if err != nil {
			return echoCtx.String(http.StatusBadRequest, fmt.Sprint("Invalid index parameter", err))
		}
		return echoCtx.JSON(http.StatusOK, clps[indexNum])

	}
	return echoCtx.JSON(http.StatusOK, clps)
}
