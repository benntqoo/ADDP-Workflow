# 🔧 依赖问题解决方案

## 🚨 当前问题
Fyne GUI框架依赖下载失败，导致编译无法完成。

## 💡 解决方案（按优先级）

### 方案一：网络环境优化（推荐）
```bash
# 1. 设置Go代理
go env -w GOPROXY=https://goproxy.cn,direct

# 2. 设置模块代理
go env -w GOSUMDB=sum.golang.google.cn

# 3. 清理并重新下载
go clean -modcache
go mod download

# 4. 重新整理依赖
go mod tidy
```

### 方案二：使用预编译版本
```bash
# 如果有预编译的二进制文件，直接使用
# 跳过编译步骤，直接运行
```

### 方案三：简化版GUI（临时方案）
创建一个基于HTML的本地GUI：
- 使用Go的net/http创建本地服务器
- HTML + CSS + JavaScript创建界面
- 通过REST API与Go后端通信

### 方案四：命令行版本
暂时回退到命令行界面：
- 移除GUI依赖
- 使用简单的命令行参数
- 保持核心功能不变

## 🛠️ 立即可用的临时解决方案

创建一个简化版本，不依赖外部GUI框架：

```go
// 使用标准库创建简单界面
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "os/exec"
)

// 简单的命令行启动器
func main() {
    if len(os.Args) < 2 {
        showUsage()
        return
    }

    switch os.Args[1] {
    case "launch":
        launchAI()
    case "list":
        listModels()
    default:
        showUsage()
    }
}
```

## 🌐 网络问题诊断

如果依赖下载失败，检查：

1. **网络连接**
   ```bash
   ping goproxy.cn
   ping github.com
   ```

2. **防火墙设置**
   - 检查企业防火墙
   - 检查本地防火墙设置

3. **代理设置**
   ```bash
   go env GOPROXY
   go env GOSUMDB
   ```

4. **DNS解析**
   ```bash
   nslookup proxy.golang.org
   nslookup goproxy.cn
   ```

## 🚀 快速恢复步骤

1. **检查Go环境**
   ```bash
   go version
   go env
   ```

2. **重置Go模块**
   ```bash
   go clean -modcache
   rm go.sum
   ```

3. **使用国内镜像**
   ```bash
   go env -w GOPROXY=https://goproxy.cn,direct
   go env -w GOSUMDB=sum.golang.google.cn
   ```

4. **重新构建**
   ```bash
   go mod download
   go mod tidy
   go build ./cmd/launcher
   ```

## 📱 备用方案：Web GUI

如果Fyne继续有问题，可以创建一个基于Web的GUI：

```go
// 使用标准库创建Web界面
func main() {
    http.HandleFunc("/", handleHome)
    http.HandleFunc("/launch", handleLaunch)

    fmt.Println("AI Launcher Web UI: http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

这样用户可以通过浏览器使用GUI界面，无需额外依赖。

## 🔄 后续计划

1. 解决网络问题后重新构建Fyne版本
2. 考虑使用更轻量的GUI框架
3. 提供多种启动方式供用户选择

---

**当前建议**：先解决网络问题，如果无法解决则使用Web GUI备用方案。