#!/usr/bin/env bash
set -euo pipefail

OS="$(uname -s)"
if [[ "$OS" != "Darwin" ]]; then
    echo "This template requires macOS."
    exit 1
fi

echo "Setting up TUI project..."

if ! command -v brew &>/dev/null; then
    echo "Error: Homebrew not found. Install from https://brew.sh"
    exit 1
fi

if ! command -v go &>/dev/null; then
    echo "Installing Go..."
    brew install go
fi

echo "Downloading Go dependencies..."
go mod download

echo ""
echo "Setup complete! Try these commands:"
echo "  go run .              Launch the TUI application"
echo "  DEBUG=1 go run .      Launch with debug logging"
echo "  go build -o bin/myapp Build a binary"
echo ""
