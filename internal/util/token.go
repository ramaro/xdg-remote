// Package util provides helper functions for xdg-remote, including browser launching
// and secure token management.
package util

import (
	"os"
	"strings"
)

// ReadToken reads the bearer token from the specified file and returns it with
// leading and trailing whitespace removed.
func ReadToken(tokenPath string) (string, error) {
	data, err := os.ReadFile(tokenPath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
