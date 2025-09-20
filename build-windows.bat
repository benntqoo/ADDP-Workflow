@echo off
chcp 65001 >nul 2>&1
title AI Launcher Windows Builder

echo ================================
echo AI启动器 Windows版本 编译工具
echo ================================
echo.

echo [1/3] 检查 Docker 环境...
docker --version >nul 2>&1
if errorlevel 1 (
    echo ❌ Docker 未安装或未启动
    echo 请先安装并启动 Docker Desktop
    pause
    exit /b 1
)
echo ✅ Docker 环境正常

echo.
echo [2/3] 编译 Windows GUI 应用程序...
echo 这可能需要几分钟时间，请耐心等待...

REM 使用Docker进行Windows交叉编译
docker run --rm -v /d/Code/fos/AI/claude:/workspace -w /workspace golang:1.23-bullseye bash -c "apt-get update && apt-get install -y gcc-mingw-w64 pkg-config && go mod download && CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o ai-launcher.exe ./cmd/gui"

if errorlevel 1 (
    echo.
    echo ❌ 编译失败！
    echo 请检查错误信息并重试
    pause
    exit /b 1
)

echo.
echo [3/3] 验证编译结果...

if exist ai-launcher.exe (
    echo.
    echo ================================
    echo 编译成功！
    echo ================================
    echo 输出文件：ai-launcher.exe
    dir ai-launcher.exe
    echo.
    echo 使用方式：
    echo   1. 直接运行：ai-launcher.exe
    echo   2. 双击图标运行
    echo.
    echo 注意：首次运行可能需要Windows Defender扫描
    echo ================================
) else (
    echo ❌ 未找到编译输出文件
    echo 编译可能存在问题
)

echo.
echo 按任意键继续...
pause >nul