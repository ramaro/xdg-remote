// client is the command-line tool for sending URLs to a remote xdg-remote server.
// It authenticates with a bearer token from a file and sends URLs via HTTP POST to the server.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ramaro/xdg-remote/internal/client"
)

var (
	serverUrl = flag.String("server-url", "http://localhost:8787", "Server HTTP URL")
	tokenFile = flag.String("token-file", "", "Path to file containing bearer token")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Opens an URL in a remote browser via xdg-remote-server url.\n\n")
		fmt.Fprintf(os.Stderr, "Usage: xdg-remote-client [options] <url>\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	url := flag.Arg(0)
	if url == "" {
		fmt.Fprintf(os.Stderr, "Error: URL is required\n")
		os.Exit(1)
	}

	if *tokenFile == "" {
		fmt.Fprintf(os.Stderr, "Error: -token-file is required\n")
		os.Exit(1)
	}

	client.Run(*serverUrl, url, *tokenFile)
}
