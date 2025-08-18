# Claude Code Agents - 實際使用指南

## ⚠️ 重要說明

本 Agents 系統是基於 Claude Code 官方的 Subagents 功能設計。請注意以下要點：

### 實際運作方式

1. **Agents 是 Subagents**：這些是 Claude Code 可以委派任務的專門 AI 助手
2. **自動委派**：Claude Code 會根據任務描述和 agent 的 description 自動選擇合適的 agent
3. **不是檔案觸發**：不會根據檔案類型自動觸發，而是根據任務內容

### ❌ 已棄用的功能

- `triggers.yaml` - Claude Code 不會解析此檔案
- `context-detector` agents - 不符合實際運作方式
- 檔案類型自動觸發 - 不支援

## 🚀 正確的安裝方式

### 1. 安裝 Agents

```bash
# 複製 agents 到 Claude 主目錄
cp -r agents/* ~/.claude/agents/

# Windows
xcopy /E /I agents\* %USERPROFILE%\.claude\agents\
```

### 2. 驗證安裝

```bash
ls ~/.claude/agents/
# 應該看到 .md 檔案
```

## 📝 Agent 配置格式

每個 agent 是一個 Markdown 檔案，包含 YAML frontmatter：

```yaml
---
name: agent-name           # 唯一識別碼
model: sonnet              # 可選: haiku/sonnet/opus
description: "何時使用此 agent 的描述，包含觸發關鍵字"
tools: Read, Write, Edit   # 可選: 限制工具訪問
---

# Agent 的系統提示詞
詳細定義 agent 的角色、能力和方法...
```

## 🎯 使用方式

### 方式 1：自動委派

Claude Code 會自動選擇合適的 agent：

```markdown
用戶: "Review this code for security issues"
# Claude 自動使用 code-reviewer agent

用戶: "This function is running slow"
# Claude 自動使用 performance-optimizer agent

用戶: "Write tests for this component"
# Claude 自動使用 test-automator agent
```

### 方式 2：明確指定

直接要求使用特定 agent：

```markdown
> Use the bug-hunter agent to find the crash issue
> Use the kotlin-expert for this Android code
```

## 📊 可用的 Agents

### 核心 Agents (Core)

| Agent | Model | 用途 | 觸發關鍵字 |
|-------|-------|------|------------|
| code-reviewer | sonnet | 代碼審查 | review, check, analyze, audit |
| performance-optimizer | opus | 性能優化 | slow, performance, optimize |
| test-automator | sonnet | 測試生成 | test, testing, coverage, TDD |

### 工作流 Agents (Workflow)

| Agent | Model | 用途 | 觸發關鍵字 |
|-------|-------|------|------------|
| bug-hunter | opus | 除錯修復 | bug, error, crash, exception |
| api-architect | sonnet | API 設計 | API, endpoint, REST, GraphQL |
| production-ready-coder | sonnet | 生產代碼 | implement, create, build |

### 語言專家 Agents (Languages)

| Agent | Model | 用途 | 適用場景 |
|-------|-------|------|----------|
| kotlin-expert | sonnet | Kotlin 全棧 | Android, Ktor, Spring Boot |
| python-ml-specialist | opus | Python ML/AI | 機器學習, 深度學習 |
| golang-systems-engineer | sonnet | Go 系統開發 | 後端服務, 系統工具 |
| typescript-fullstack-expert | sonnet | TypeScript | React, Node.js, Next.js |

### 品質保證 Agents (Quality)

| Agent | Model | 用途 | 特色 |
|-------|-------|------|------|
| jenny-validator | haiku | 規範驗證 | 檢查編碼規範 |
| karen-realist | haiku | 現實評估 | 時間和範圍評估 |
| senior-developer | opus | 資深審查 | 10+ 年經驗視角 |

## 🔧 Model 選擇指南

- **haiku**: 簡單任務、快速回應
- **sonnet**: 標準開發任務（預設）
- **opus**: 複雜任務、深度分析

## 💡 最佳實踐

### 1. 優化 Description

```yaml
# ❌ 不好的 description
description: "Code review agent"

# ✅ 好的 description  
description: "Review code for bugs, security issues, performance. Use when user mentions review, check, analyze, or audit."
```

### 2. 使用觸發詞

在 description 中包含常見觸發詞，提高自動委派準確性：
- 動作詞：review, optimize, test, debug, implement
- 問題詞：slow, error, bug, crash, failing
- 領域詞：security, performance, API, database

### 3. 限制工具訪問

只給 agent 必要的工具：
```yaml
tools: Read, Grep, Glob  # 只讀 agents
tools: Read, Write, Edit, Bash  # 完整開發 agents
```

## 🚨 常見問題

### Q: 為什麼 agent 沒有自動觸發？

A: 檢查：
1. Agent 是否在 `~/.claude/agents/` 目錄
2. Description 是否包含相關關鍵字
3. 任務描述是否明確

### Q: 如何確認 agent 被使用？

A: Claude Code 會在使用 subagent 時顯示訊息

### Q: 可以同時使用多個 agents 嗎？

A: 可以，Claude Code 會根據需要協調多個 agents

## 📚 參考資源

- [Claude Code Subagents 官方文檔](https://docs.anthropic.com/en/docs/claude-code/sub-agents)
- [wshobson/agents 參考實現](https://github.com/wshobson/agents)

## 🔄 從舊系統遷移

如果你之前使用 v4.0 的複雜配置：

1. 忽略 `triggers.yaml` - 已棄用
2. 簡化 context-detector agents 為專業 agents
3. 更新 descriptions 包含觸發關鍵字
4. 移除不必要的複雜配置

---

*本指南基於 Claude Code 實際功能編寫，而非理想化設計*