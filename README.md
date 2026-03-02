# xdg-remote

Open URLs in your local browser when working on a remote machine via SSH.

## Problem

When connected via SSH to a remote server, running `xdg-open https://example.com` attempts to open the URL on the remote machine rather than your local machine. This often fails because remote servers typically lack a graphical environment. This tool solves that by forwarding URLs to open in your local browser.

## How it works

1. **Server** - Runs on your local machine. It listens for HTTP requests and opens URLs in your local browser.

2. **Client** - Runs on the remote machine. Sends an HTTP POST request to the server with the URL to open.

## Usage

### 1. Start the server (on your local machine)

```bash
xdg-remote-server -token-file /path/to/token
```

The server listens on port 8787 by default. Use `-addr` to change the address.

### 2. SSH with port forwarding

```bash
ssh -R 8787:localhost:8787 user@remote
```

This forwards port 8787 on the remote machine to port 8787 on your local machine.

### 3. Use from remote machine (via SSH)

```bash
xdg-remote-client https://example.com -server-url http://localhost:8787 -token-file /path/to/token
```

## Setup

1. Build the binaries:
   ```bash
   go build -o xdg-remote-server ./cmd/server
   go build -o xdg-remote-client ./cmd/client
   ```

2. Create a token file (used for authentication):
   ```bash
   mkdir -p ~/.config/xdg-remote
   echo "your-secret-token" > ~/.config/xdg-remote/token
   ```

3. Copy `xdg-remote-server` to your local machine and `xdg-remote-client` to the remote machine.

4. Start the server on your local machine.

5. SSH with port forwarding: `ssh -R 8787:localhost:8787 user@remote`

6. From SSH, use the client to open URLs in your local browser.

## Setting up the client as your default xdg-open

Most remote hosts already have an `xdg-open` command. To intercept URL-opening requests and forward them to your local browser, install the wrapper script:

```bash
# On the remote machine
mkdir -p ~/.local/bin
cp xdg-remote-client ~/.local/bin/
cp scripts/xdg-open.sh ~/.local/bin/xdg-open
chmod 700 ~/.local/bin/xdg-open
```

To ensure the new `xdg-open` takes precedence, update your PATH in `.bashrc` or `.zshrc`:

```bash
export PATH=~/.local/bin:$PATH
```

Now any program that calls `xdg-open` will open URLs in your local browser instead.

## Options

### Server
- `-addr` - Address to listen on (default: `:8787`)
- `-token-file` - Path to file containing bearer token (required)
- `-debug` - Enable debug logging

### Client
- `-server-url` - Server HTTP URL (default: `http://localhost:8787`)
- `-token-file` - Path to file containing bearer token (required)
- `<url>` - URL to open in the browser (positional argument, required)

## Security

The server uses bearer token authentication. Set appropriate permissions on the token file:
```bash
chmod 600 ~/.config/xdg-remote/token
```
