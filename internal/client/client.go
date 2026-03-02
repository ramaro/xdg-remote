// Package client provides functionality for sending URLs to a remote xdg-remote server.
package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ramaro/xdg-remote/internal/util"
)

// Run sends the URL to serverUrl using HTTP POST with Bearer token authentication.
// The token is read from tokenFile and included in the Authorization header.
func Run(serverUrl string, url string, tokenFile string) {
	if url == "" {
		fmt.Fprintf(os.Stderr, "Error: URL is required\n")
		os.Exit(1)
	}

	if tokenFile == "" {
		fmt.Fprintf(os.Stderr, "Error: -token-file is required\n")
		os.Exit(1)
	}

	token, err := util.ReadToken(tokenFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading token file: %v\n", err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", serverUrl, strings.NewReader(url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create request: %v\n", err)
		os.Exit(1)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Server error: %s %s\n", resp.Status, body)
		os.Exit(1)
	}

	fmt.Printf("Opened: %s\n", url)
}
