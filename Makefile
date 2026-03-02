.PHONY: build clean

build:
	go build -o xdg-remote-server ./cmd/server
	go build -o xdg-remote-client ./cmd/client

clean:
	rm -f xdg-remote-server xdg-remote-client
