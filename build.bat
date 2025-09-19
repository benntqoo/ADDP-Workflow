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
set CGO_ENABLED=1
go build -o ai-launcher.exe ./cmd/launcher

echo [4/4] Checking build result...
if exist ai-launcher.exe (
    echo.
    echo ================================
    echo BUILD SUCCESSFUL!
    echo ================================
    echo Executable: ai-launcher.exe
    echo.
    echo Usage Instructions:
    echo   1. Double-click ai-launcher.exe to start GUI
    echo   2. Or run from command line:
    echo      ai-launcher.exe          ^(Start GUI^)
    echo      ai-launcher.exe version  ^(Show version^)
    echo      ai-launcher.exe list-models  ^(List AI models^)
    echo.
    echo TIP: Just double-click the exe file - no command line needed!
    echo ================================
    echo.
) else (
    echo.
    echo ================================
    echo BUILD FAILED!
    echo ================================
    echo Please check if dependencies are properly installed.
    echo NOTE: Fyne GUI requires CGO support.
    echo Install a C compiler like TDM-GCC or Visual Studio Build Tools.
    echo ================================
    echo.
)

echo Press any key to continue...
pause >nul