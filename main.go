package main

import (
	clipboard "clip-cli/clipboard/listener"
	"clip-cli/handler"
	"clip-cli/internal/core"

	"github.com/labstack/echo/v4"
)

func main() {

	// init Clipboard
	clipboard := clipboard.New(clipboard.WithClipboardListenerName("Listener#1"))

	// init core
	core := core.New(
		core.WithClipboardListener(clipboard),
	)

	// init echo framework with handlers and core
	handleClient := handler.New(
		handler.WithCore(core),
		handler.WithPort(":58080"),
		handler.WithServiceName("ClipServer"),
		handler.WithEchoFramework(echo.New()),
		handler.WithRoutesAndMiddlewares(),
	)

	handleClient.Start()
}
