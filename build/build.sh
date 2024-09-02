#!/bin/bash

# Define output directory
OUT_DIR="bin"

# Create the output directory if it doesn't exist
mkdir -p $OUT_DIR

# Build for Linux (amd64)
GOOS=linux GOARCH=amd64 go build -o $OUT_DIR/urlcheck-linux-amd64

# Build for Linux (arm64)
GOOS=linux GOARCH=arm64 go build -o $OUT_DIR/urlcheck-linux-arm64

# Build for Windows (amd64)
GOOS=windows GOARCH=amd64 go build -o $OUT_DIR/urlcheck-windows-amd64.exe

# Build for macOS (amd64)
GOOS=darwin GOARCH=amd64 go build -o $OUT_DIR/urlcheck-darwin-amd64

# Build for macOS (arm64)
GOOS=darwin GOARCH=arm64 go build -o $OUT_DIR/urlcheck-darwin-arm64

echo "Builds completed and placed in $OUT_DIR"