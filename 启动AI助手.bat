@echo off
chcp 65001 >nul 2>&1
title AI Launcher

REM Check if ai-launcher.exe exists
if not exist ai-launcher.exe (
    echo ERROR: ai-launcher.exe not found
    echo Please run build.bat first to build the program
    echo.
    echo Press any key to continue...
    pause >nul
    exit /b 1
)

REM Start GUI program
echo Starting AI Launcher...
start "" ai-launcher.exe

REM Wait 1 second then close command window
timeout /t 1 /nobreak > nul
exit