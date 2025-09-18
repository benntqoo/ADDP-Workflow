# CLAUDE.md - AI Coding 統一協作框架

## 項目概述

**Universal AI Coding Framework** - 基於 MCP + Ollama 的新一代 AI 編程協作解決方案

本項目旨在解決當前 AI coding 工具的核心痛點：
- 🔄 **多工具同步問題** - 不同 AI coding 工具間無法共享上下文和工作狀態
- 🎯 **提問精準度低** - 用戶模糊輸入導致 AI 回應發散，浪費 token
- 🏗️ **缺乏標準化工作流程** - 無紀律性開發導致反覆修改和低效率
- 🔒 **隱私和廠商綁定** - 敏感項目信息外洩風險，被特定工具綁定

## 🎯 核心理念

### ADDP 框架 (AI-Driven Disciplined Programming)
- **A**nalysis - 需求分析階段
- **D**esign - 架構設計階段
- **D**evelopment - 開發實施階段
- **P**ersistence - 持久化驗證階段

### MCP + Ollama 雙層架構
```
用戶輸入 → Ollama優化層 → MCP統一服務層 → AI工具執行層 → 統一回饋
```

## 🏗️ 技術架構

### 第一層：Ollama 本地優化器
- **功能**：將用戶模糊輸入轉換為精準技術指令
- **模型**：Qwen2.5:14b 或其他適合的本地模型
- **優勢**：100% 隱私保護，3秒內響應，300% 精準度提升

### 第二層：MCP 統一服務層
- **功能**：標準化工作流程，跨工具記憶同步
- **協議**：Model Context Protocol (MCP)
- **覆蓋**：Claude Code, Gemini CLI, Codex, Cursor 等

### 第三層：AI 工具執行層
- **支援工具**：所有主流 AI coding CLI 和 GUI 工具
- **統一體驗**：相同的命令、記憶、工作流程
- **無縫切換**：項目狀態完全同步

## 📁 項目結構

```
claude/
├── .claude/                    # Claude Code 配置目錄
├── workflow-legacy/            # 舊版本工作流程系統存檔
├── CLAUDE.md                   # 項目協作規範（本文件）
├── README.md                   # 項目說明文檔
└── TARGET.md                   # MCP架構規劃和目標
```

## 🚀 預期效果

### 解決多工具同步
- ✅ 統一記憶系統，所有工具共享項目上下文
- ✅ MCP 協議確保跨工具兼容性
- ✅ 實時狀態同步，無縫工具切換

### 提高提問精準度
- ✅ 本地 Ollama 智能優化用戶輸入
- ✅ 40-60% 精準度提升
- ✅ 30-50% token 節省

### 標準化工作流程
- ✅ 強制性 TDD 約束機制
- ✅ ADDP 四階段標準流程
- ✅ 智能質量守護者

## 🎯 開發優先級

### Phase 1：核心 MVP（1-2個月）
1. 實現 Ollama 查詢優化器
2. 建立基礎 MCP 服務
3. 支援 Claude Code + Gemini CLI
4. 驗證架構可行性

### Phase 2：完整功能（2-4個月）
1. 完善 ADDP 工作流程
2. 實現統一記憶系統
3. 支援更多 AI 工具
4. 建立質量保證機制

### Phase 3：生態建設（4-6個月）
1. 開源社群推廣
2. 建立最佳實踐庫
3. 與工具廠商合作
4. 推動行業標準化

## 💡 關鍵創新點

1. **本地隱私優化** - Ollama 確保敏感信息不外洩
2. **零廠商綁定** - MCP 開放標準支援所有工具
3. **智能查詢優化** - 本地 AI 預處理提高執行精準度
4. **統一工作流程** - 跨工具一致的開發體驗
5. **標準化記憶** - 項目知識永久積累和重用

## 🔄 使用工作流程

```bash
# 1. 啟動本地優化器
ollama serve

# 2. 啟動 MCP 統一服務
mcp-server --config unified-coding-config.json

# 3. 配置 AI 工具
claude config mcp-servers add unified-coding-assistant

# 4. 開始協作
claude "我想實現用戶登錄功能"
# → Ollama 優化 → MCP 處理 → 精準執行 → 統一記憶
```

---

**這是一個重新定義 AI coding 工具使用方式的革命性框架！**

*最後更新：2025-09-18*