package main

import (
	"clip-cli/cli-interface/cli"
	httpHelper "clip-cli/cli-interface/http-helpers"
	"os"
)

func main() {
	helperHttp := httpHelper.New(httpHelper.WithUrl("http://localhost:58080"))
	// helperHttp.GetClips()
	cliClient := &cli.Client{HttpHelper: helperHttp}

	if len(os.Args) == 1 {
		cliClient.PromotList()
		return
	}
	command := os.Args[1]
	if command == "-h" {
		cliClient.PromptHelp()
	}
	if command == "-v" {
		cliClient.PromptVersion()
	}
}
