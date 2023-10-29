.PHONY: install
install:
	@echo "Installing Go 1.18"
	go install golang.org/dl/go1.18.10@latest
	go1.18.10 download
	export GOPHERJS_GOROOT="$(go1.18.10 env GOROOT)"
	
	@echo "Installing GopherJS"
	go install github.com/gopherjs/gopherjs@v1.18.0-beta3

.PHONY: build
build:
	gopherjs build .
