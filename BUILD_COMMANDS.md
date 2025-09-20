# AI Launcher Docker 编译命令

## 手动 Docker 编译命令

### 1. Linux 版本编译
```bash
# 构建Linux版本镜像
docker build -t ai-launcher:linux .

# 提取Linux二进制文件
docker create --name temp-linux ai-launcher:linux
docker cp temp-linux:/app/ai-launcher ./ai-launcher-linux
docker rm temp-linux
```

### 2. Windows 版本编译（推荐）
```bash
# 构建Windows版本镜像
docker build -f Dockerfile.windows -t ai-launcher:windows .

# 提取Windows exe文件
docker create --name temp-windows ai-launcher:windows
docker cp temp-windows:/ai-launcher.exe ./ai-launcher.exe
docker rm temp-windows
```

### 3. 交互式编译（调试用）
```bash
# 进入构建环境进行手动编译
docker run -it --rm -v "%cd%":/workspace -w /workspace golang:1.23-bullseye bash

# 在容器内手动执行编译
apt-get update && apt-get install -y gcc-mingw-w64 pkg-config
go mod download
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o ai-launcher.exe ./cmd/gui
```

### 4. 一键命令（Windows PowerShell）
```powershell
# 构建并提取Windows exe
docker build -f Dockerfile.windows -t ai-launcher:windows . && docker create --name temp-windows ai-launcher:windows && docker cp temp-windows:/ai-launcher.exe ./ai-launcher.exe && docker rm temp-windows
```

### 5. 一键命令（Linux/macOS）
```bash
# 构建并提取Windows exe
docker build -f Dockerfile.windows -t ai-launcher:windows . && \
docker create --name temp-windows ai-launcher:windows && \
docker cp temp-windows:/ai-launcher.exe ./ai-launcher.exe && \
docker rm temp-windows
```

## 清理命令
```bash
# 清理所有相关镜像和容器
docker rmi ai-launcher:linux ai-launcher:windows 2>/dev/null || true
docker rm temp-linux temp-windows 2>/dev/null || true

# 清理悬空镜像
docker image prune -f
```

## 验证编译结果
```bash
# 检查生成的文件
ls -la ai-launcher*

# Windows用户检查文件属性
file ai-launcher.exe
```

## 注意事项
1. **Fyne GUI依赖**：需要CGO支持，Windows版本使用mingw交叉编译
2. **文件大小**：编译后的exe文件大约25-30MB
3. **运行环境**：Windows exe可以在Windows 10/11上直接运行，无需额外依赖
4. **调试模式**：如果编译失败，使用交互式模式进行调试