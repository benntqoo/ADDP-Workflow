# 使用 Docker 编译 Fyne GUI

> 本项目已移除 Web GUI，仅保留本地 Fyne GUI。推荐通过 Docker 进行跨平台编译，避免在本机安装 CGO 与系统依赖。

## Windows 可执行文件 (.exe)

在任意主机（Windows/macOS/Linux）上构建 Windows 可执行文件：

```powershell
cd "D:\Code\fos\AI\claude"
docker run --rm -v D:\Code\fos\AI\claude:/workspace -w /workspace golang:1.23-bullseye bash -c "apt-get update -qq && apt-get install -y gcc-mingw-w64 pkg-config && go mod download && CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o ai-launcher.exe ./cmd/gui"
```

完成后，项目根目录会生成 ai-launcher.exe，可直接在 Windows 上运行。

## Linux 可执行文件

在容器内构建 Linux 二进制：

```bash
docker run --rm -v D:\Code\fos\AI\claude:/workspace -w /workspace golang:1.23-bullseye bash -c "apt-get update -qq && apt-get install -y gcc pkg-config libgl1-mesa-dev libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libxext-dev libxfixes-dev libx11-dev libxxf86vm-dev && go mod download && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v -o ai-launcher ./cmd/gui"
```

## 使用内置 Dockerfile（可选）

也可使用仓库内的 Dockerfile 完成构建：
- Linux: `build/docker/Dockerfile`
- Windows 交叉编译: `build/docker/Dockerfile.windows`

示例：

```bash
# Linux 构建镜像（示例标签）
docker build -f build/docker/Dockerfile -t ai-launcher:linux .

# Windows 交叉编译镜像（示例标签）
docker build -f build/docker/Dockerfile.windows -t ai-launcher:windows .
# 从镜像中拷贝生成的 exe（如镜像内产物为 /ai-launcher.exe）
docker create --name temp ai-launcher:windows && docker cp temp:/ai-launcher.exe ./ai-launcher.exe && docker rm temp
```

## 常见问题

- Fyne 需要 CGO 和系统图形依赖。使用 Docker 可以自动安装所需依赖，避免污染本机环境。
- 如果挂载路径出错，请确认 Docker Desktop 的文件共享设置已包含当前盘符/目录。
- 若公司网络限制导致 apt-get 失败，请在公司内源镜像上定制基础镜像，或下载依赖后缓存构建。

### 使用命名缓存卷（加速 Go 依赖与构建缓存）

```powershell
# 按项目 + Go 版本命名缓存卷，避免污染
$Project      = "ai-launcher"
$GoTag        = "1.23-bullseye"
$GoMajorMinor = "go1.23"
$ModVol   = "go-mod-cache-$Project-$GoMajorMinor"
$BuildVol = "go-build-cache-$Project-$GoMajorMinor"

$CurrentDir = Get-Location
$image = "golang:$GoTag"
$cmd = (
  "apt-get update -qq",
  "apt-get install -y gcc-mingw-w64 pkg-config",
  "go mod download",
  "CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -ldflags='-H windowsgui' -v -o ai-launcher.exe ./cmd/gui"
) -join " && "

docker run --rm `
  -v "${CurrentDir}:/workspace" `
  -v "${ModVol}:/go/pkg/mod" `
  -v "${BuildVol}:/root/.cache/go-build" `
  -w /workspace `
  $image bash -c $cmd
```

## 使用 Docker Compose 构建

项目内已提供 Compose 编排文件：`build/docker-compose.yml`

- 从仓库根目录执行（推荐）：

```powershell
# 构建 Windows 交叉编译的构建镜像（固化 mingw）
docker compose -f .\build\docker-compose.yml build win-build

# 编译产物（会在仓库根目录生成 ai-launcher.exe）
docker compose -f .\build\docker-compose.yml run --rm win-build
```

- 从 `build/` 目录执行：

```powershell
# 省略 -f，Compose 会默认读取当前目录下的 docker-compose.yml
docker compose build win-build
docker compose run --rm win-build
```

- 开发容器（可选）：

```powershell
# 进入带 Go 环境的 builder 容器，复用命名缓存卷
docker compose -f .\build\docker-compose.yml run --rm ai-launcher-dev
```

说明
- Compose 文件位于 `build/` 下，已将 `build.context` 设为仓库根（..），并使用命名缓存卷：
  - go-mod-cache-ai-launcher-go1.23 → /go/pkg/mod
  - go-build-cache-ai-launcher-go1.23 → /root/.cache/go-build
- 如需更换 Go 版本，可同步修改卷名后缀（例如 go1.24）与镜像标签。

