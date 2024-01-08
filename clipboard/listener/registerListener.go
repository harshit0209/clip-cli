package clipboard

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/atotto/clipboard"
)

func (c *Client) RegisterListener(ctx context.Context, s string) error {
	fmt.Println("Registering Clipbaord Listener")
	// Check if clipboard is supported on the current platform
	if !clipboard.Unsupported {
		fmt.Println("Clipboard operations are supported.")

		// Start a goroutine to continuously check for changes in the clipboard
		go c.clipboardCopyListener()

	} else {
		fmt.Println("Clipboard operations are not supported on this platform.")
	}
	return nil
}

// Avoid Race with mutex, I dont think race is achiveable in this case xD
func (c *Client) clipboardCopyListener() {
	var lastClipboardContent string

	for {
		// Get the current clipboard content
		clipboardContent, err := clipboard.ReadAll()
		if err != nil {
			fmt.Println("Error reading clipboard:", err)
			os.Exit(1)
		}

		// Check if the clipboard content has changed
		if clipboardContent != lastClipboardContent && clipboardContent != c.FlagClip {
			fmt.Print("Clipboard content changed:")
			fmt.Println(clipboardContent)
			c.Clips = append(c.Clips, clipboardContent)
			// Update the last clipboard content
			lastClipboardContent = clipboardContent
		}
		// Sleep for a short duration before checking again
		time.Sleep(500 * time.Millisecond)
	}
}
