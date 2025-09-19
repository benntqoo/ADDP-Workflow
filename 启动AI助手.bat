@echo off
title AI启动器

REM 检查ai-launcher.exe是否存在
if not exist ai-launcher.exe (
    echo 错误：找不到 ai-launcher.exe 文件
    echo 请先运行 build.bat 构建程序
    echo.
    pause
    exit /b 1
)

REM 启动GUI程序
echo 正在启动AI启动器...
start "" ai-launcher.exe

REM 等待一秒后关闭命令行窗口
timeout /t 1 /nobreak > nul
exit