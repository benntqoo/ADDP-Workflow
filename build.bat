@echo off
echo 构建AI启动器...

echo 正在清理旧的构建文件...
if exist ai-launcher.exe del ai-launcher.exe

echo 正在下载依赖...
go mod download

echo 正在编译程序...
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1
go build -o ai-launcher.exe ./cmd/launcher

if exist ai-launcher.exe (
    echo.
    echo ✅ 构建成功！
    echo 可执行文件：ai-launcher.exe
    echo.
    echo 使用方法：
    echo   双击 ai-launcher.exe 启动GUI界面
    echo   或在命令行运行：
    echo     ai-launcher.exe          启动GUI界面
    echo     ai-launcher.exe version  显示版本信息
    echo     ai-launcher.exe list-models  列出支持的AI模型
    echo.
    echo 💡 提示：直接双击exe文件即可使用，无需命令行操作！
    echo.
) else (
    echo.
    echo ❌ 构建失败！
    echo 请检查依赖是否正确安装。
    echo 注意：Fyne GUI需要CGO支持，请确保安装了C编译器（如TDM-GCC）。
    echo.
)

pause