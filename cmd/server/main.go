// server runs an HTTP server that receives URLs from remote clients and opens them
// in the default browser using Bearer token authentication.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ramaro/xdg-remote/internal/server"
)

var (
	addr      = flag.String("addr", ":8787", "Server address (host:port)")
	tokenFile = flag.String("token-file", "", "Path to file containing bearer token")
	debug     = flag.Bool("debug", false, "Enable debug logging")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Runs an HTTP server to open URL requests in the default browser.\n\n")
		fmt.Fprintf(os.Stderr, "Usage: xdg-remote-server [options]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *tokenFile == "" {
		fmt.Fprintf(os.Stderr, "Error: -token-file is required\n")
		os.Exit(1)
	}

	server.Run(*addr, *tokenFile, *debug)
}
