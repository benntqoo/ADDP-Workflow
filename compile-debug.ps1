# PowerShell 编译脚本
$CurrentDir = Get-Location
Write-Host "当前目录: $CurrentDir"

# 编译调试版本
Write-Host "开始编译调试版本..."
docker run --rm -v "${CurrentDir}:/workspace" -w /workspace golang:1.23-bullseye bash -c "apt-get update -qq && apt-get install -y gcc-mingw-w64 pkg-config && go mod download && CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o ai-launcher-debug.exe ./cmd/gui"

if ($LASTEXITCODE -eq 0) {
    Write-Host "编译成功！"
    Write-Host "运行调试版本: .\ai-launcher-debug.exe"

    # 询问是否立即运行
    $run = Read-Host "是否立即运行调试版本？(y/n)"
    if ($run -eq "y" -or $run -eq "Y") {
        .\ai-launcher-debug.exe
    }
} else {
    Write-Host "编译失败！"
    exit 1
}