@echo off
echo 正在測試AI啟動器...
echo.

echo 系統信息:
systeminfo | findstr /C:"OS Name" /C:"OS Version" /C:"System Type"
echo.

echo 嘗試運行程序...
ai-launcher.exe > output.txt 2>&1

echo 程序已退出，檢查輸出:
if exist output.txt (
    type output.txt
) else (
    echo 沒有輸出文件生成
)

echo.
echo 檢查錯誤級別: %ERRORLEVEL%
echo.

echo 按任意鍵退出...
pause > nul