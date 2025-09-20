# AI Launcher EXE 调试指南

## 问题诊断

### 第一步：通过命令行运行以查看错误信息

```cmd
# 在 cmd 中运行
cd "D:\Code\fos\AI\claude"
ai-launcher.exe

# 或者在 PowerShell 中运行
cd "D:\Code\fos\AI\claude"
.\ai-launcher.exe

# 查看详细错误信息
.\ai-launcher.exe 2>&1
```

### 第二步：检查依赖项

```powershell
# 检查文件是否完整
Get-ItemProperty .\ai-launcher.exe | Select-Object Name, Length
file .\ai-launcher.exe

# 使用 Dependency Walker 或 Dependencies.exe 查看DLL依赖
# 下载 https://github.com/lucasg/Dependencies
```

### 第三步：常见问题排查

1. **缺少运行时库**
   - Visual C++ Redistributable
   - Windows 10/11 兼容性

2. **防火墙/杀毒软件阻止**
   - Windows Defender
   - 第三方杀毒软件

3. **GUI相关问题**
   - OpenGL 驱动
   - 显卡驱动

## 快速修复

### 临时解决方案
在程序开始处添加错误捕获和延迟退出：

```go
// 在 main 函数开始处添加
defer func() {
    if r := recover(); r != nil {
        fmt.Printf("程序发生错误: %v\n", r)
        fmt.Println("按回车键退出...")
        fmt.Scanln()
    }
}()
```

### 调试版本编译
编译一个带控制台输出的调试版本：

```bash
# 编译调试版本（显示控制台）
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -ldflags="-H windowsgui" -v -o ai-launcher-debug.exe ./cmd/gui
```

## 立即可尝试的解决方案

1. **以管理员身份运行**
2. **关闭杀毒软件暂时测试**
3. **在兼容模式下运行**
4. **检查Windows事件日志**