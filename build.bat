@echo off
chcp 65001 >nul 2>&1
title AI Launcher Builder

echo Building AI Launcher...
echo.

echo [1/4] Cleaning old build files...
if exist ai-launcher.exe del ai-launcher.exe

echo [2/4] Downloading dependencies...
go mod download

echo [3/4] Compiling program...
set GOOS=windows
set GOARCH=amd64
go build -o ai-launcher.exe ./cmd/simple

echo [4/4] Checking build result...
if exist ai-launcher.exe (
    echo.
    echo ================================
    echo BUILD SUCCESSFUL!
    echo ================================
    echo Executable: ai-launcher.exe
    echo.
    echo Usage Instructions:
    echo   1. Double-click ai-launcher.exe to start Web GUI
    echo   2. Or run from command line:
    echo      ai-launcher.exe          ^(Start Web GUI^)
    echo      ai-launcher.exe version  ^(Show version^)
    echo      ai-launcher.exe help     ^(Show help^)
    echo.
    echo TIP: Web interface will open in your browser automatically!
    echo ================================
    echo.
) else (
    echo.
    echo ================================
    echo BUILD FAILED!
    echo ================================
    echo Please check if Go is properly installed and accessible.
    echo This version uses only Go standard library - no external dependencies.
    echo Try: go version
    echo ================================
    echo.
)

echo Press any key to continue...
pause >nul