#!/bin/bash
echo "Building AI Launcher for multiple platforms..."

echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o ai-launcher.exe ./cmd/launcher

echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o ai-launcher-linux ./cmd/launcher

echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o ai-launcher-darwin ./cmd/launcher

echo "Building for macOS (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -o ai-launcher-darwin-arm64 ./cmd/launcher

echo "Build completed!"
echo ""
echo "Windows: ai-launcher.exe"
echo "Linux:   ai-launcher-linux"
echo "macOS (Intel):   ai-launcher-darwin"
echo "macOS (Apple Silicon): ai-launcher-darwin-arm64"