// Package util provides helper functions for xdg-remote, including browser launching
// and secure token management.
package util

import (
	"os"
	"os/exec"
	"strings"
)

// OpenURL launches the given URL in the system's default browser.
// It detects the desktop environment and uses the appropriate command (xdg-open on Linux, open on macOS).
func OpenURL(url string) error {
	var cmd *exec.Cmd
	switch strings.ToLower(os.Getenv("XDG_CURRENT_DESKTOP")) {
	case "gnome", "unity", "cinnamon", "mate", "lxde", "xfce", "kde", "plasma":
		cmd = exec.Command("xdg-open", url)
	case "sway", "wayfire", "hyprland", "river", "weston", "enlightenment":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	return cmd.Start()
}
