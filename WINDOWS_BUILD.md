# Windows EXE 构建指南

## 问题说明
在Windows环境下使用Docker进行交叉编译时，路径映射存在问题。以下提供几种解决方案：

## 方案1：直接使用PowerShell（推荐）

```powershell
# 在PowerShell中执行以下命令
cd "D:\Code\fos\AI\claude"
docker run --rm -v ${PWD}:/workspace -w /workspace golang:1.23-bullseye bash -c "apt-get update -qq && apt-get install -y gcc-mingw-w64 pkg-config && go mod download && CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o ai-launcher.exe ./cmd/gui"
```

## 方案2：使用WSL2（如果可用）

```bash
# 在WSL2中执行
cd /mnt/d/Code/fos/AI/claude
docker run --rm -v $(pwd):/workspace -w /workspace golang:1.23-bullseye bash -c "apt-get update -qq && apt-get install -y gcc-mingw-w64 pkg-config && go mod download && CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o ai-launcher.exe ./cmd/gui"
```

## 方案3：使用Docker Desktop的文件共享

1. 打开Docker Desktop设置
2. 进入"Resources" -> "File sharing"
3. 添加 `D:\Code\fos\AI\claude` 目录
4. 重启Docker Desktop
5. 执行：

```bash
docker run --rm -v "/d/Code/fos/AI/claude":/workspace -w /workspace golang:1.23-bullseye bash -c "apt-get update -qq && apt-get install -y gcc-mingw-w64 pkg-config && go mod download && CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o ai-launcher.exe ./cmd/gui"
```

## 方案4：本地Go环境编译（如果已安装Go和MinGW）

```bash
# 确保已安装MinGW-w64和Go
go env -w CGO_ENABLED=1
go env -w GOOS=windows
go env -w GOARCH=amd64
go env -w CC=x86_64-w64-mingw32-gcc

# 编译
go build -v -o ai-launcher.exe ./cmd/gui
```

## 方案5：分步构建

### 步骤1：启动构建容器
```bash
docker run -it --rm --name windows-builder -v "/d/Code/fos/AI/claude":/workspace golang:1.23-bullseye bash
```

### 步骤2：在容器内安装依赖
```bash
apt-get update -qq
apt-get install -y gcc-mingw-w64 pkg-config
cd /workspace
go mod download
```

### 步骤3：编译
```bash
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o ai-launcher.exe ./cmd/gui
```

### 步骤4：退出容器
```bash
exit
```

## 快速命令（复制粘贴执行）

### PowerShell版本：
```powershell
cd "D:\Code\fos\AI\claude"; docker run --rm -v ${PWD}:/workspace -w /workspace golang:1.23-bullseye bash -c "apt-get update -qq && apt-get install -y gcc-mingw-w64 pkg-config && go mod download && CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o ai-launcher.exe ./cmd/gui"
```

### Git Bash版本：
```bash
cd /d/Code/fos/AI/claude && docker run --rm -v "/$(pwd -W | tr '\\' '/'):/workspace" -w /workspace golang:1.23-bullseye bash -c "apt-get update -qq && apt-get install -y gcc-mingw-w64 pkg-config && go mod download && CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o ai-launcher.exe ./cmd/gui"
```

## 验证结果

编译成功后，你应该看到：
- `ai-launcher.exe` 文件，大小约 25-30MB
- 可以在Windows上直接双击运行

## 故障排除

1. **路径问题**：确保使用正确的路径格式
2. **Docker权限**：确保Docker有访问文件夹的权限
3. **防火墙**：Windows Defender可能会扫描新生成的exe文件
4. **依赖问题**：如果编译失败，检查CGO和MinGW安装

## 成功标志

如果看到以下输出，说明编译成功：
```
ai-launcher/internal/gui
ai-launcher/cmd/gui
```

然后在当前目录下会生成 `ai-launcher.exe` 文件。