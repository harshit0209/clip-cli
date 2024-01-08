package cli

import (
	httpHelper "clip-cli/cli-interface/http-helpers"
	"fmt"
	"strings"
	"unicode/utf8"
)

type Client struct {
	HttpHelper *httpHelper.Client
}

func (c *Client) PromotList() {
	clips, err := c.HttpHelper.GetClips()
	if err != nil {
		return
	}

	fmt.Println("You have 📋")

	for i, clip := range clips {
		fmt.Printf("%v. %v 📎\n", i, truncateString(clip, 50))
	}
	var clip int
	fmt.Println("Enter index that you want to load-clipboard")
	fmt.Scanf("%d", &clip)
	fmt.Println("Copied to clipboard!", clips[clip])
	c.HttpHelper.SetClips(clip)
}

func (c *Client) PromptHelp() {
	fmt.Println("Usage Clip-Cli")
	fmt.Println("-h or --help\tshows this")
	fmt.Println("-v or --version\tshow version")
}
func (c *Client) PromptVersion() {
	fmt.Println("0.1")
}
func truncateString(s string, maxLen int) string {
	s = strings.ReplaceAll(s, "\n", "")
	if utf8.RuneCountInString(s) > maxLen {
		runes := []rune(s)
		return string(runes[:maxLen]) + "..."
	}
	return s
}
