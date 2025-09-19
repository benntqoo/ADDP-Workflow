@echo off
chcp 936 >nul 2>&1
title AI启动器

REM 检查ai-launcher.exe是否存在
if not exist ai-launcher.exe (
    echo 错误：找不到 ai-launcher.exe 文件
    echo 请先运行 build.bat 构建程序
    echo.
    echo 按任意键继续...
    pause >nul
    exit /b 1
)

REM 启动GUI程序
echo 正在启动AI启动器...
echo 请稍等，程序正在加载中...
start "" ai-launcher.exe

REM 等待一秒后关闭命令行窗口
timeout /t 2 /nobreak > nul
exit