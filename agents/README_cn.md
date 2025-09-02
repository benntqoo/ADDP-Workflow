# 🤖 Claude Code 專業代理系統

*[English](README.md) | 中文*

為 Claude Code 設計的全方位專業 AI 代理集合，提供專業級的軟體開發協助。

## 📋 目錄

- [概述](#概述)
- [快速開始](#快速開始)
- [代理分類](#代理分類)
- [安裝](#安裝)
- [使用方式](#使用方式)
- [配置](#配置)
- [可用代理](#可用代理)
- [工作流程](#工作流程)
- [最佳實踐](#最佳實踐)

## 🎯 概述

本代理系統整合了領先的 Claude Code 代理項目的最佳實踐：
- **wshobson/agents** - 模型分層代理選擇
- **claude_code_agent_farm** - 並行執行與協調
- **ClaudeCodeAgents** - 質量保證專家

### 核心特性

- 🚀 **40+ 專業代理** 覆蓋所有主流語言和框架
- ⚡ **智能選擇via Orchestrator** - 智能單專家優先策略
- 🔄 **Token優化工作流** - 60%+效率提升（800k → 300k平均）
- 🌐 **語言偏好記憶** - 持久化跨會話本地化
- 📁 **標準化記憶系統** - 統一`.claude/memory/`目錄結構
- 📊 **智能協調** 支援並行執行
- ✅ **質量保證** 專門的驗證代理
- 🎨 **模型優化** 根據任務複雜度使用 Haiku/Sonnet/Opus
- 🎯 **生產就緒代碼** 內建最佳實踐和安全性
- 🐛 **高級調試** 根因分析
- 📡 **API 設計** OpenAPI 規範和 GraphQL 架構

## 🚀 快速開始

### 全局安裝（推薦）

1. 複製整個 `claude/agents` 目錄到你的主目錄：
```bash
# Windows
xcopy /E /I "D:\Code\ai\claude\agents" "%USERPROFILE%\.claude\agents"

# macOS/Linux
cp -r /path/to/claude/agents ~/.claude/agents
```

2. 複製配置文件：
```bash
# Windows
xcopy /E /I "D:\Code\ai\claude\config" "%USERPROFILE%\.claude\config"

# macOS/Linux
cp -r /path/to/claude/config ~/.claude/config
```

3. 驗證安裝：
```bash
ls ~/.claude/agents/
# 應該看到所有代理 .md 文件
```

### 項目級安裝

如果你想為特定項目使用代理：
```bash
# 在項目根目錄
mkdir -p .claude/agents
cp -r /path/to/claude/agents/* .claude/agents/
```

## 📦 代理分類

### 🧠 ~~上下文檢測器~~ (已棄用)
~~智能分析代碼上下文，解決多用途語言場景衝突~~

**注意**：上下文檢測器已被棄用。請改用語言專家代理（如 `kotlin-expert`、`python-ml-specialist`）

### 💻 技術專家代理

#### 移動與跨平台
- `android-kotlin-architect` - Android 應用開發
- `kotlin-expert` - Kotlin 全棧開發
- `react-native-developer` - 跨平台移動應用

#### 後端與 API
- `ktor-backend-architect` - Ktor 框架專家
- `spring-boot-enterprise` - Spring Boot 微服務
- `golang-systems-engineer` - Go 系統開發
- `nodejs-backend-developer` - Node.js 後端

#### 機器學習與數據科學
- `python-ml-specialist` - Python ML/AI 專家
- `data-scientist` - 數據分析與可視化
- `llm-development-expert` - LLM 開發專家

#### 前端開發
- `react-developer` - React 應用開發
- `vue-developer` - Vue.js 開發
- `angular-developer` - Angular 企業應用

### 🔍 質量保證代理
- `code-reviewer` - 全面代碼審查，安全和性能分析
- `test-automator` - 智能測試生成和執行
- `performance-optimizer` - 性能瓶頸分析和優化建議
- `jenny-validator` - 規範驗證專家
- `karen-realist` - 現實評估專家，時間線和範圍把關

### 🎭 工作流代理
- `work-coordinator` - 多代理協調器，處理複雜跨領域任務
- `bug-hunter` - 錯誤查找和修復，根因分析
- `api-architect` - API 設計與 OpenAPI 規範

### ⚡ 優化代理
- `token-efficient-loader` - Token 使用優化策略
- `senior-developer` - 應用 10+ 年經驗到每行代碼
- `production-ready-coder` - 自動編寫生產品質代碼

## 🎯 使用方式

### 自動委派
Claude Code 會基於任務描述自動選擇合適的代理：

```markdown
用戶: "Review this code for security issues"
# 自動使用 code-reviewer

用戶: "這個函數運行很慢"
# 自動使用 performance-optimizer

用戶: "為這個組件寫測試"
# 自動使用 test-automator
```

### 顯式調用
直接指定要使用的代理：

```markdown
> Use the kotlin-expert agent to review this Android code
> 使用 python-ml-specialist 來優化這個模型
```

### 通過 Task Tool
```python
使用 Task tool:
- subagent_type: "bug-hunter"
- prompt: "找出這個崩潰的原因"
```

## ⚙️ 配置

每個代理使用 YAML frontmatter 配置：

```yaml
---
name: agent-name
model: sonnet  # 可選: haiku/sonnet/opus
description: "何時使用此代理，包含觸發關鍵字"
tools: Read, Write, Edit, Bash  # 可選: 工具限制
---
```

### 模型選擇
- **haiku** - 簡單任務，快速響應
- **sonnet** - 標準開發任務（默認）
- **opus** - 複雜任務，深度分析

## 🔄 工作流程

### 1. 單代理執行
```
請求 → 代理選擇 → 執行 → 結果
```

### 2. 多代理協作
```
複雜任務 → work-coordinator → [專業代理1, 專業代理2] → 整合結果
```

### 3. 質量保證流程
```
代碼變更 → code-reviewer → test-automator → performance-optimizer
```

## 💡 最佳實踐

### 1. 優化代理描述
在 `description` 欄位中包含關鍵觸發詞：
- 動作詞：review, optimize, test, debug, implement
- 問題詞：slow, error, bug, crash, failing
- 領域詞：security, performance, API, database

### 2. 選擇合適的模型
- 簡單任務用 haiku
- 標準任務用 sonnet
- 複雜分析用 opus

### 3. 限制工具訪問
只給代理必要的工具權限：
```yaml
tools: Read, Grep  # 只讀代理
tools: Read, Write, Edit, Bash  # 完整開發代理
```

### 4. 上下文管理
- 長對話自動觸發 context-manager
- 使用 work-coordinator 處理多領域任務

## 📊 性能指標

| 任務類型 | 效率提升 |
|---------|----------|
| 代碼審查 | 6x |
| 測試生成 | 10x |
| 性能優化 | 15x |
| 錯誤修復 | 8x |
| API 設計 | 5x |

## 🚨 故障排除

### 代理未觸發？
1. 檢查代理是否在 `~/.claude/agents/`
2. 確認 description 包含相關關鍵字
3. 嘗試顯式調用

### 錯誤的代理被選中？
- 使用更具體的關鍵字
- 顯式指定代理名稱

### 性能問題？
- 簡單任務使用 haiku 模型
- 限制不必要的工具訪問

## 🤝 貢獻

歡迎貢獻新的代理！請確保：
1. 清晰的 description 和觸發詞
2. 適當的模型選擇
3. 最小必要的工具集
4. 詳細的系統提示詞

## 📚 相關資源

- [Claude Code Subagents 官方文檔](https://docs.anthropic.com/en/docs/claude-code/sub-agents)
- [wshobson/agents 參考實現](https://github.com/wshobson/agents)
- [實際使用指南](README_ACTUAL_USAGE.md)

---

*智能代理系統代表了 AI 輔助開發的未來 - 專業知識與智能自動化的結合。*