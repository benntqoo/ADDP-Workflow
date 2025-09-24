# Windows Debug 编译（在 Docker 内构建 + 命名缓存卷）

$Project      = "ai-launcher"
$GoTag        = "1.23-bullseye"
$GoMajorMinor = "go1.23"

$ModVol   = "go-mod-cache-$Project-$GoMajorMinor"
$BuildVol = "go-build-cache-$Project-$GoMajorMinor"

$RepoRoot = (Resolve-Path (Join-Path $PSScriptRoot "..")).Path
Write-Host "仓库根目录: $RepoRoot"
Write-Host "使用缓存卷: $ModVol / $BuildVol"

$image = "golang:$GoTag"
$cmd = @(
  "apt-get update -qq",
  "apt-get install -y gcc-mingw-w64 pkg-config",
  "go mod download",
  "CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o ai-launcher-debug.exe ./cmd/gui"
) -join " && "

docker run --rm `
  -v "${RepoRoot}:/workspace" `
  -v "${ModVol}:/go/pkg/mod" `
  -v "${BuildVol}:/root/.cache/go-build" `
  -w /workspace `
  $image bash -c $cmd

if ($LASTEXITCODE -eq 0) {
    Write-Host "编译成功: .\ai-launcher-debug.exe" -ForegroundColor Green
} else {
    Write-Host "编译失败" -ForegroundColor Red
    exit 1
}

