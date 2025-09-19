# 🚀 AI启动器 (AI Launcher)

*智能AI工具启动器 - 一键启动各种AI编程助手*

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org/)
[![GUI](https://img.shields.io/badge/GUI-Fyne-orange.svg)](https://fyne.io/)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-green.svg)](https://github.com/fyne-io/fyne)

## 🎯 核心理念：一键启动，智能管理 AI 开发工具

### 解决开发者痛点
- **🔄 工具切换繁琐** - Claude Code、Gemini CLI、Codex 各有不同启动方式
- **📁 目录管理困难** - 每次都要手动切换到项目目录
- **⚙️ 配置复杂** - YOLO模式、安全参数等需要记忆复杂命令
- **🎯 缺乏统一界面** - 没有统一的项目管理和快速启动方案

### 🌟 我们的解决方案

**直观GUI界面 + 智能项目管理 = 极简AI开发体验**

```
双击启动 → GUI界面 → 选择项目 → 一键启动AI工具
   ↓         ↓        ↓         ↓
 ai-launcher.exe → 项目管理 → Claude/Gemini/Codex
```

### 💡 核心特性：GUI界面 + 智能配置管理

基于现代GUI框架，提供最佳的用户体验：

- **🖥️ 直观GUI界面**：800x600可视化界面，支持拖拽和点击操作
- **📁 项目管理**：自动保存最近项目，一键加载配置
- **🤖 多AI支持**：Claude Code、Gemini CLI、Codex完整支持
- **⚡ YOLO模式**：安全模式vs快速模式，适应不同开发需求

## ✨ 界面展示

### 🖥️ GUI主界面 (800x600)
```
┌─────────────────────────────────────────────────────────────────────────────┐
│ 🚀 AI启动器 - 智能多AI工具启动器                    │ 🔄 📖 ⚙️ ❌ │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────────────────────┐  ┌───────────────────────────────────────────┐ │
│  │  📁 最近使用的项目       │  │  ⚙️ 项目配置                              │ │
│  │  ┌─────────────────────┐ │  │                                           │ │
│  │  │ 🤖 my-web-app       │ │  │  项目路径: ┌─────────────────────────┐🔍  │ │
│  │  │ /path/to/project    │ │  │           │ /Users/dev/my-project   │     │ │
│  │  │ Claude Code • 12:34 │ │  │           └─────────────────────────┘     │ │
│  │  └─────────────────────┘ │  │                                           │ │
│  │  ┌─────────────────────┐ │  │  AI模型: ┌─────────────────────────────┐   │ │
│  │  │ 💎 data-analyzer    │ │  │         │ 🤖 Claude Code              │▼ │ │
│  │  │ /work/analyzer      │ │  │         └─────────────────────────────┘   │ │
│  │  │ Gemini CLI • 10:15  │ │  │                                           │ │
│  │  └─────────────────────┘ │  │  ⚡ ☐ 启用YOLO模式 (跳过安全确认)          │ │
│  │                         │  │                                           │ │
│  │  ┌─────────────────────┐ │  │                                           │ │
│  │  │ ➕ 新建项目          │ │  │                                           │ │
│  │  └─────────────────────┘ │  │                                           │ │
│  └─────────────────────────┘  └───────────────────────────────────────────┘ │
├─────────────────────────────────────────────────────────────────────────────┤
│ 状态: 准备启动 my-project                                                    │
│ ┌─────────────────┐  ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐              │
│ │ 🚀 启动AI工具    │  │ 💾保存 │ │ ⚙️设置 │ │ 📖关于 │ │ ❌退出 │              │
│ └─────────────────┘  └──────┘ └──────┘ └──────┘ └──────┘              │
└─────────────────────────────────────────────────────────────────────────────┘
```
## 🚀 快速开始

### 📦 安装方式

#### Windows用户（推荐）
```bash
# 1. 克隆项目
git clone https://github.com/your-repo/ai-launcher.git
cd ai-launcher

# 2. 构建程序
./build.bat

# 3. 双击启动
# 直接双击 ai-launcher.exe 即可使用
```

#### 手动构建（所有平台）
```bash
# 确保安装了Go 1.24+
go version

# 克隆并构建
git clone https://github.com/your-repo/ai-launcher.git
cd ai-launcher
go mod download
go build -o ai-launcher ./cmd/launcher
```

### 🎯 使用方法

#### 方式一：双击启动（最简单）
1. 双击 `ai-launcher.exe`
2. GUI界面自动打开
3. 选择项目目录和AI模型
4. 点击"🚀 启动AI工具"

#### 方式二：命令行启动
```bash
# 启动GUI界面
./ai-launcher.exe

# 查看版本信息
./ai-launcher.exe version

# 列出支持的AI模型
./ai-launcher.exe list-models
```

## 🛠️ 支持的AI工具

### 🤖 Claude Code
```bash
# 普通模式
claude

# YOLO模式（跳过权限确认）
claude --dangerously-skip-permissions
```

### 💎 Gemini CLI
```bash
# 普通模式
gemini

# YOLO模式
gemini --yolo
```

### 🔧 Codex
```bash
# 普通模式
codex

# YOLO模式（跳过审批和沙盒）
codex --dangerously-bypass-approvals-and-sandbox
```

## ⚡ YOLO模式说明

### 🛡️ 普通模式（推荐）
- ✅ 需要用户确认重要操作
- ✅ 适合生产环境和重要项目
- ✅ 更加安全可靠

### 🚀 YOLO模式（实验性）
- ⚠️ 跳过大部分安全检查
- ⚠️ 自动执行AI建议的操作
- ⚠️ 适合实验和快速原型

## 📁 项目管理

- **最近项目**：自动保存最近使用的10个项目
- **一键加载**：点击项目卡片快速加载配置
- **智能保存**：自动保存项目配置到用户目录
- **配置持久化**：配置保存在 `~/.ai-launcher/` 目录

## 🔧 技术架构

### 🏗️ 核心组件
- **GUI界面**：基于 [Fyne](https://fyne.io/) 框架的跨平台GUI
- **项目管理**：Go语言实现的配置管理系统
- **终端管理**：智能终端启动和管理
- **跨平台支持**：Windows、Linux、macOS全平台支持

### 📂 目录结构
```
ai-launcher/
├── cmd/launcher/          # 主程序入口
├── internal/
│   ├── gui/              # GUI界面实现
│   ├── project/          # 项目配置管理
│   └── terminal/         # 终端管理
├── build.bat             # Windows构建脚本
├── 启动AI助手.bat          # 便捷启动脚本
└── GUI_DESIGN.md         # 界面设计文档
```

## 📋 系统要求

### 💻 运行环境
- **Windows**：Windows 10+ (推荐)
- **Linux**：任何支持GTK+的发行版
- **macOS**：macOS 10.12+

### 🔧 开发环境
- **Go**：1.24+ (构建需要)
- **CGO**：需要C编译器支持Fyne GUI
- **Git**：克隆代码库

## 🤝 贡献指南

欢迎贡献代码和想法！请参考以下步骤：

1. **Fork** 本项目
2. **创建功能分支** (`git checkout -b feature/amazing-feature`)
3. **提交更改** (`git commit -m 'Add amazing feature'`)
4. **推送分支** (`git push origin feature/amazing-feature`)
5. **创建Pull Request**

## 📝 更新日志

### v2.0.0 (当前版本)
- ✅ 完整的GUI界面实现
- ✅ 项目配置管理系统
- ✅ 多AI模型支持 (Claude/Gemini/Codex)
- ✅ YOLO模式支持
- ✅ 双击启动体验

### v1.0.0 (历史版本)
- ✅ 基础TUI界面 (已移除)
- ✅ 终端代理功能
- ✅ 项目架构设计

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

- [Fyne](https://fyne.io/) - 优秀的Go GUI框架
- [Cobra](https://cobra.dev/) - 强大的CLI框架
- 所有AI工具开发者们的杰出工作

---

**💡 提示**：如果你觉得这个项目有用，请给我们一个 ⭐！
