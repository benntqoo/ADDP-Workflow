@echo off
chcp 65001 >nul 2>&1
title Fix AI Launcher Dependencies

echo Fixing AI Launcher dependencies...
echo.

echo [1/5] Setting up Go proxy...
go env -w GOPROXY=https://goproxy.cn,direct

echo [2/5] Cleaning module cache...
go clean -modcache

echo [3/5] Removing go.sum...
if exist go.sum del go.sum

echo [4/5] Downloading dependencies step by step...
go mod download fyne.io/fyne/v2@v2.4.5
go mod download github.com/spf13/cobra@v1.8.0
go mod download github.com/stretchr/testify@v1.8.4

echo [5/5] Running go mod tidy...
go mod tidy

echo.
echo ================================
echo Dependencies fixed!
echo ================================
echo Now try building with: go build ./cmd/launcher
echo.

pause