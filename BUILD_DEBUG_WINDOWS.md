# Windows 调试版本编译指南

## 编译调试版本（显示控制台）

### PowerShell 命令
```powershell
cd "D:\Code\fos\AI\claude"

# 编译调试版本（显示控制台窗口）
docker run --rm -v ${PWD}:/workspace -w /workspace golang:1.23-bullseye bash -c "apt-get update -qq && apt-get install -y gcc-mingw-w64 pkg-config && go mod download && CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -ldflags='-H windowsgui' -v -o ai-launcher-debug.exe ./cmd/gui"
```

### 运行调试版本
```cmd
# 在命令行中运行，可以看到所有日志输出
ai-launcher-debug.exe

# 或者带控制台参数
ai-launcher-debug.exe --console
```

## 编译普通版本（无控制台）

### PowerShell 命令
```powershell
cd "D:\Code\fos\AI\claude"

# 编译无控制台版本
docker run --rm -v ${PWD}:/workspace -w /workspace golang:1.23-bullseye bash -c "apt-get update -qq && apt-get install -y gcc-mingw-w64 pkg-config && go mod download && CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -ldflags='-H windowsgui' -v -o ai-launcher.exe ./cmd/gui"
```

## 故障排除步骤

1. **运行调试版本**
   ```cmd
   ai-launcher-debug.exe
   ```
   查看控制台输出，找到具体错误信息

2. **检查日志**
   - 查看是否有 "创建主窗口..." 等日志
   - 找到最后一条日志，确定失败位置

3. **常见错误及解决方案**

   **错误：OpenGL相关**
   ```
   无法初始化OpenGL上下文
   ```
   解决：更新显卡驱动

   **错误：DLL缺失**
   ```
   The program can't start because xxx.dll is missing
   ```
   解决：安装 Visual C++ Redistributable

   **错误：权限问题**
   ```
   Access denied
   ```
   解决：以管理员身份运行

## 快速测试命令

```powershell
# 一键编译调试版本并运行
cd "D:\Code\fos\AI\claude"
docker run --rm -v ${PWD}:/workspace -w /workspace golang:1.23-bullseye bash -c "apt-get update -qq && apt-get install -y gcc-mingw-w64 pkg-config && go mod download && CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o ai-launcher-debug.exe ./cmd/gui" && ./ai-launcher-debug.exe
```