# 快速修复Windows启动问题

## 立即可尝试的解决方案

### 1. 在命令行中运行现有程序
```cmd
cd "D:\Code\fos\AI\claude"
ai-launcher.exe
```
观察输出的错误信息。

### 2. 检查依赖项
```cmd
# 检查是否缺少DLL
where ai-launcher.exe
file ai-launcher.exe
```

### 3. 以管理员身份运行
右键 -> "以管理员身份运行"

### 4. 使用兼容模式
右键程序 -> 属性 -> 兼容性 -> Windows 10

## 如果编译有问题，使用备用方案

### 方案A: 使用PowerShell ISE
```powershell
# 在 PowerShell ISE 中执行
Set-Location "D:\Code\fos\AI\claude"
.\build\docker\build.ps1 -Target windows -OutputDir dist -KeepImage
```

### 方案B: 使用WSL（如果已安装）
```bash
cd /mnt/d/Code/fos/AI/claude
./build/docker/build.sh windows dist
```

## 最可能的错误及解决方案

### 1. OpenGL/图形驱动问题
**症状**: 程序启动立即关闭，无错误信息
**解决**: 更新显卡驱动

### 2. 缺少Visual C++ Redistributable
**症状**: 提示缺少DLL文件
**解决**: 安装最新的 VC++ Redistributable

### 3. 防火墙/杀毒软件阻止
**症状**: 程序被阻止运行
**解决**: 临时关闭杀毒软件测试

### 4. 权限问题
**症状**: Access denied错误
**解决**: 以管理员身份运行

## 立即测试步骤

1. 打开命令行（cmd）
2. 导航到程序目录：`cd "D:\Code\fos\AI\claude"`
3. 运行程序：`ai-launcher.exe`
4. 观察输出，报告具体错误信息

如果看到任何错误信息，请告诉我具体内容，我可以提供针对性的解决方案。