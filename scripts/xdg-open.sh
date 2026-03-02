#!/bin/bash
# Wrapper script for xdg-remote-client
# This script wraps xdg-remote-client calls with the required parameters
# Usage: xdg-open.sh <url>

# Default token file location - adjust as needed
TOKEN_FILE="${XDG_REMOTE_TOKEN_FILE:-$HOME/.config/xdg-remote/token}"

# Default server URL - can be overridden via environment variable
SERVER_URL="${XDG_REMOTE_SERVER_URL:-http://localhost:8787}"

# Get the URL from the first argument
URL="$1"

if [ -z "$URL" ]; then
    echo "Error: URL is required" >&2
    echo "Usage: $0 <url>" >&2
    exit 1
fi

# Call xdg-remote-client with the wrapped parameters
exec xdg-remote-client -server-url "$SERVER_URL" -token-file "$TOKEN_FILE" "$URL"
