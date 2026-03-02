// Package server provides an HTTP server that receives URLs from remote clients
// and opens them in the system's default browser using Bearer token authentication.
package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ramaro/xdg-remote/internal/util"
)

// Handler returns an HTTP handler that validates Bearer tokens and opens URLs
// from the request body in the system's default browser.
func Handler(token string, debug bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		providedToken := strings.TrimPrefix(authHeader, "Bearer ")
		if providedToken != token {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusBadRequest)
			return
		}
		urlToOpen := string(body)
		urlToOpen = strings.TrimSpace(urlToOpen)

		if urlToOpen == "" {
			http.Error(w, "URL required", http.StatusBadRequest)
			return
		}

		if debug {
			log.Printf("Opening URL: %s", urlToOpen)
		}

		if err := util.OpenURL(urlToOpen); err != nil {
			http.Error(w, fmt.Sprintf("Failed to open URL: %v", err), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}

// Run starts an HTTP server on addr that validates Bearer tokens against the token
// stored in tokenFile, then opens received URLs in the default browser.
func Run(addr, tokenFile string, debug bool) {
	if debug {
		log.SetOutput(os.Stdout)
		log.SetFlags(0)
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

	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler(token, debug))

	fmt.Printf("Server starting on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
