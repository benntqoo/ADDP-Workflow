@echo off
echo 构建AI启动器GUI版本...

echo 正在清理旧的构建文件...
if exist ai-launcher-gui.exe del ai-launcher-gui.exe

echo 正在下载依赖...
go mod download

echo 正在编译GUI版本...
set GOOS=windows
set GOARCH=amd64
go build -o ai-launcher-gui.exe ./cmd/launcher

if exist ai-launcher-gui.exe (
    echo.
    echo ✅ 构建成功！
    echo 可执行文件：ai-launcher-gui.exe
    echo.
    echo 使用方法：
    echo   ai-launcher-gui.exe         启动TUI界面
    echo   ai-launcher-gui.exe gui     启动GUI界面
    echo   ai-launcher-gui.exe version 显示版本信息
    echo.
) else (
    echo.
    echo ❌ 构建失败！
    echo 请检查依赖是否正确安装。
    echo.
)

pause