@echo off
echo Building AI Launcher for multiple platforms...

echo Building for Windows...
set GOOS=windows
set GOARCH=amd64
go build -o ai-launcher.exe ./cmd/launcher

echo Building for Linux...
set GOOS=linux
set GOARCH=amd64
go build -o ai-launcher-linux ./cmd/launcher

echo Building for macOS...
set GOOS=darwin
set GOARCH=amd64
go build -o ai-launcher-darwin ./cmd/launcher

echo Build completed!
echo.
echo Windows: ai-launcher.exe
echo Linux:   ai-launcher-linux
echo macOS:   ai-launcher-darwin