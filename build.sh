#!/bin/bash

# Build script for Linux/macOS
set -e

APP_NAME="nzb-monkey-go"
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS="-ldflags \"-X main.appVersion=$VERSION -s -w\""

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
MAGENTA='\033[0;35m'
NC='\033[0m' # No Color

# Create bin directory
mkdir -p bin

function install_deps() {
    echo -e "${GREEN}Installing dependencies...${NC}"
    go mod download
    go get fyne.io/fyne/v2@latest
    go mod tidy
    echo -e "${CYAN}Dependencies installed!${NC}"
}

function build_current() {
    echo -e "${GREEN}Building for current platform...${NC}"
    eval go build $LDFLAGS -o "bin/$APP_NAME"
    echo -e "${CYAN}Binary created: bin/$APP_NAME${NC}"
}

function build_linux() {
    echo -e "${GREEN}Building for Linux...${NC}"
    GOOS=linux GOARCH=amd64 eval go build $LDFLAGS -o "bin/$APP_NAME-linux-amd64"
    echo -e "${CYAN}Linux binary created: bin/$APP_NAME-linux-amd64${NC}"
}

function build_linux_arm64() {
    echo -e "${GREEN}Building for Linux ARM64...${NC}"
    GOOS=linux GOARCH=arm64 eval go build $LDFLAGS -o "bin/$APP_NAME-linux-arm64"
    echo -e "${CYAN}Linux ARM64 binary created: bin/$APP_NAME-linux-arm64${NC}"
}

function build_windows() {
    echo -e "${GREEN}Building for Windows...${NC}"
    GOOS=windows GOARCH=amd64 eval go build $LDFLAGS -o "bin/$APP_NAME-windows-amd64.exe"
    echo -e "${CYAN}Windows binary created: bin/$APP_NAME-windows-amd64.exe${NC}"
}

function build_darwin() {
    echo -e "${GREEN}Building for macOS...${NC}"
    GOOS=darwin GOARCH=amd64 eval go build $LDFLAGS -o "bin/$APP_NAME-darwin-amd64"
    GOOS=darwin GOARCH=arm64 eval go build $LDFLAGS -o "bin/$APP_NAME-darwin-arm64"
    echo -e "${CYAN}macOS binaries created: bin/$APP_NAME-darwin-*${NC}"
}

function build_all() {
    echo -e "${YELLOW}Building for all platforms...${NC}"
    build_linux
    build_linux_arm64
    build_windows
    build_darwin
    echo ""
    echo -e "${GREEN}All platform binaries built successfully!${NC}"
}

function clean() {
    echo -e "${YELLOW}Cleaning build artifacts...${NC}"
    rm -rf bin/
    rm -f "$APP_NAME"
    echo -e "${CYAN}Build artifacts cleaned!${NC}"
}

function install_local() {
    build_linux
    echo -e "${GREEN}Installing to /usr/local/bin...${NC}"
    sudo install -m 755 "bin/$APP_NAME-linux-amd64" "/usr/local/bin/$APP_NAME"
    echo -e "${CYAN}Installed to /usr/local/bin/$APP_NAME${NC}"
}

function show_help() {
    echo "NZB Monkey Go Build Script"
    echo ""
    echo "Usage: $0 [command]"
    echo ""
    echo "Commands:"
    echo "  current      Build for current platform (default)"
    echo "  linux        Build for Linux (amd64)"
    echo "  linux-arm64  Build for Linux (ARM64)"
    echo "  windows      Build for Windows"
    echo "  darwin       Build for macOS (Intel + Apple Silicon)"
    echo "  all          Build for all platforms"
    echo "  clean        Remove build artifacts"
    echo "  install      Build and install to /usr/local/bin (Linux only)"
    echo "  deps         Install/update dependencies"
    echo "  help         Show this help message"
    echo ""
}

# Main execution
echo -e "${MAGENTA}NZB Monkey Go Build Script${NC}"
echo -e "${MAGENTA}Version: $VERSION${NC}"
echo ""

# Install dependencies first
install_deps
echo ""

case "${1:-current}" in
    current)
        build_current
        ;;
    linux)
        build_linux
        ;;
    linux-arm64)
        build_linux_arm64
        ;;
    windows)
        build_windows
        ;;
    darwin|macos)
        build_darwin
        ;;
    all)
        build_all
        ;;
    clean)
        clean
        ;;
    install)
        install_local
        ;;
    deps)
        # Already installed at the start
        ;;
    help|--help|-h)
        show_help
        ;;
    *)
        echo -e "${RED}Unknown command: $1${NC}"
        show_help
        exit 1
        ;;
esac

echo ""
echo -e "${GREEN}Build completed!${NC}"
