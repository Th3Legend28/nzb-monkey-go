.PHONY: all build build-linux build-windows build-darwin clean install deps

APP_NAME=nzb-monkey-go
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X main.appVersion=$(VERSION) -s -w"

# Default target
all: deps build

# Install dependencies
deps:
	go mod download
	go get fyne.io/fyne/v2@latest
	go mod tidy

# Build for current platform
build:
	@echo "Building for current platform..."
	go build $(LDFLAGS) -o bin/$(APP_NAME)

# Build for Linux (primary target)
build-linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/$(APP_NAME)-linux-amd64
	@echo "Linux binary created: bin/$(APP_NAME)-linux-amd64"

# Build for Linux ARM64 (e.g., Raspberry Pi)
build-linux-arm64:
	@echo "Building for Linux ARM64..."
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o bin/$(APP_NAME)-linux-arm64
	@echo "Linux ARM64 binary created: bin/$(APP_NAME)-linux-arm64"

# Build for Windows
build-windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/$(APP_NAME)-windows-amd64.exe
	@echo "Windows binary created: bin/$(APP_NAME)-windows-amd64.exe"

# Build for macOS
build-darwin:
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/$(APP_NAME)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o bin/$(APP_NAME)-darwin-arm64
	@echo "macOS binaries created: bin/$(APP_NAME)-darwin-*"

# Build for all platforms
build-all: build-linux build-linux-arm64 build-windows build-darwin
	@echo "All platform binaries built successfully!"

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f $(APP_NAME)
	@echo "Build artifacts cleaned"

# Install to system (Linux)
install: build-linux
	install -m 755 bin/$(APP_NAME)-linux-amd64 /usr/local/bin/$(APP_NAME)
	@echo "Installed to /usr/local/bin/$(APP_NAME)"

# Run the application (with GUI)
run-gui:
	go run . --gui

# Run the application (CLI mode with test NZBLNK)
run-cli:
	go run . --help
