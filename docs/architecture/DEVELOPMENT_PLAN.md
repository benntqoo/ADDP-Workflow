# 糾正版開發計劃
*雙軌架構：Golang 終端代理 + Python MCP 獨立服務*

## 🎯 架構糾正後的開發策略

### 雙軌並行開發
```
軌道1: Golang 終端代理 (主軌)
├─ GUI 界面管理
├─ 多終端代理 (claude/gemini/cursor/aider)
├─ Ollama 查詢優化
└─ 命令路由分發

軌道2: Python MCP 獨立服務 (輔軌)
├─ TDD 流程化工具
├─ 工作流管理
├─ 用戶可選配置
└─ 與 Golang 完全解耦
```

## 📅 重新規劃的時程安排

### Week 1-2: Golang 終端代理核心 (主要精力)
**目標**: 實現終端管理 + Ollama 優化的基礎框架

#### Week 1: 終端管理器 + Ollama 集成
```bash
# 可演示產物
□ claude-proxy --list-terminals          # ✅ 列出可用終端類型
□ claude-proxy --start claude            # ✅ 啟動 Claude Code 終端
□ claude-proxy --optimize "修復bug"      # ✅ Ollama 查詢優化演示
□ claude-proxy --send claude "優化後命令" # ✅ 發送命令到終端

# 硬性驗收標準
□ 成功啟動 4 種終端 (claude/gemini/cursor/aider)
□ Ollama 優化響應時間 < 2秒
□ 終端進程管理穩定（啟動/停止/重啟）
□ 跨平台終端命令適配 (Windows/Linux/macOS)
```

#### Week 2: TUI 界面 + 命令路由
```bash
# 可演示產物
□ claude-proxy                          # ✅ 啟動完整 TUI 界面
□ Tab 切換面板 (輸入/終端選擇/輸出)      # ✅ 界面導航
□ 1-4 數字鍵快速切換終端                # ✅ 終端選擇
□ Enter 發送優化後命令                  # ✅ 端到端流程

# 硬性驗收標準
□ TUI 界面響應流暢（無卡頓）
□ 多終端並行處理（最多4個同時活躍）
□ 命令歷史記錄和回溯
□ 錯誤處理和用戶友好提示
```

---

### Week 3-4: 多終端深度適配 (主要精力)
**目標**: 針對每個 AI 工具優化命令格式和交互方式

#### Week 3: Claude Code + Gemini CLI 適配
```bash
# Claude Code 專項適配
□ claude-proxy --demo claude-commands   # ✅ 演示 Claude 命令優化
□ /sync → 會話同步優化
□ /plan → 規劃命令優化
□ /context → 上下文分析優化

# Gemini CLI 專項適配
□ claude-proxy --demo gemini-prompts    # ✅ 演示 Gemini 提示優化
□ 自然語言查詢結構化
□ 代碼請求格式標準化
□ 多輪對話上下文保持

# 驗收標準
□ Claude Code 命令識別準確率 > 90%
□ Gemini 自然語言優化效果可感知
□ 兩個終端可以同時運行互不干擾
□ 查詢優化針對性強（不同工具不同模板）
```

#### Week 4: Cursor + Aider 適配 + 穩定性測試
```bash
# Cursor 專項適配
□ claude-proxy --demo cursor-integration # ✅ Cursor 集成演示
□ 代碼編輯請求優化
□ 文件定位和修改指令
□ Cursor Chat 模式適配

# Aider 專項適配
□ claude-proxy --demo aider-commands     # ✅ Aider 命令演示
□ 編程指令自然語言化
□ 文件操作命令優化
□ Git 集成工作流

# 穩定性測試
□ 長時間運行測試 (2小時)
□ 終端崩潰自動恢復
□ 記憶體洩漏檢查
□ 併發命令處理壓力測試
```

---

### Week 5-6: Python MCP 獨立開發 (次要精力)
**目標**: 實現 TDD 流程化工具，完全獨立於 Golang

#### Week 5: MCP 服務框架 + 核心工具
```bash
# Python MCP 獨立服務
□ python -m src.mcp_server --stdio       # ✅ STDIO 模式運行
□ python -m src.mcp_server --http        # ✅ HTTP 模式運行
□ claude-mcp-test                        # ✅ 與 Claude 集成測試

# 最小可行 MCP 工具集（與 ARCHITECTURE.md 對齊）
□ query.optimize - Ollama 驅動的查詢優化
□ plan.update - 項目計劃更新和狀態管理
□ fs.apply_patch - 安全的文件系統修補（白名單保護）
□ test.run - TDD 測試執行和驗證
□ guard.validate_flow - 開發流程門禁驗證
□ mem.save_phase - 項目記憶和狀態持久化
□ prompt.render - 提示詞模板渲染
□ context.pack - 上下文打包和優化

# DoD 對齊驗收標準
□ MCP 協議標準合規（stdio/JSON-RPC 運行）
□ 工具執行成功率 > 95%，結構化 IO
□ fs.apply_patch 拒絕非白名單路徑並記錄
□ test.run 返回機器可讀報告
□ guard.validate_flow 在測試失敗時阻止
□ 初始化 .addp/{specifications,workflows,memory,queries,gates,sync,configs}
□ 與主流 AI 工具 MCP 配置兼容
```

#### Week 6: 工作流管理 + 用戶配置
```bash
# 工作流管理
□ ADDP 四階段流程管理
□ TDD Red-Green-Refactor 驗證
□ 代碼質量門禁
□ 項目狀態持久化

# 用戶配置系統
□ Claude Code 配置模板
□ Cursor 配置模板
□ 通用 MCP 配置指南
□ 故障排除文檔

# DoD 對齊驗收標準
□ 強制執行階段順序（ADDP），違規時報告修復提示
□ 每次調用寫入 MCP_RUN_ID、輸入、輸出到 .addp/**
□ 記錄基本分析（延遲 P50/P95、緩存命中率、失敗代碼）
□ 配置模板在各 AI 工具中可用性 > 90%
□ 故障排除文檔完整，常見問題有明確解決方案
□ 端到端工作流演示：query.optimize → prompt.render → context.pack
```

---

### Week 7: 集成測試 + 生產部署
**目標**: 兩個獨立系統的最終打包和部署

#### 系統集成驗證
```bash
# Golang 終端代理驗證
□ claude-proxy --benchmark              # ✅ 性能基準測試
□ claude-proxy --test-all-terminals     # ✅ 所有終端兼容性測試
□ claude-proxy --install                # ✅ 一鍵安裝腳本

# Python MCP 獨立驗證
□ 在 Claude Code 中配置和使用 MCP 工具
□ 在 Cursor 中配置和使用 MCP 工具
□ MCP 工具獨立運行測試

# 用戶體驗測試
□ 新用戶 5 分鐘上手測試
□ 跨平台部署測試 (Windows/Linux/macOS)
□ 錯誤場景恢復測試
```

---

## 🎯 核心產物清晰定位

### Golang 終端代理 (主產物)
```bash
# 核心功能
- 統一的 GUI 界面管理多個 AI 終端
- Ollama 本地查詢優化 (針對不同工具)
- 智能命令路由和格式轉換
- 終端會話管理和歷史記錄

# 目標用戶
- 需要同時使用多個 AI 工具的開發者
- 希望本地優化查詢節省 token 的用戶
- 喜歡終端界面的效率用戶

# 獨立價值
即使不使用 Python MCP，Golang 代理本身就是完整可用的產品
```

### Python MCP 獨立服務 (可選產物)
```bash
# 核心功能
- TDD 流程化工具集
- 開發工作流管理
- 代碼質量門禁
- 項目狀態記憶

# 目標用戶
- 使用支持 MCP 的 AI 工具的開發者
- 需要標準化開發流程的團隊
- 重視代碼質量和 TDD 的開發者

# 獨立價值
可以單獨配置給任何支持 MCP 的 AI 工具使用，與終端代理完全解耦
```

---

## 🔧 技術決策簡化

### Golang 技術棧
```yaml
core_dependencies:
  - github.com/charmbracelet/bubbletea  # TUI 框架
  - github.com/ollama/ollama/api        # Ollama 集成
  - github.com/creack/pty               # 終端 PTY 控制

architecture:
  - internal/gui: TUI 界面管理
  - internal/terminal: 終端進程管理
  - internal/optimizer: Ollama 查詢優化
  - internal/router: 命令路由分發

deployment:
  - 單一二進制文件
  - 跨平台編譯 (Goreleaser)
  - 一鍵安裝腳本
```

### Python 技術棧
```yaml
core_dependencies:
  - mcp: MCP 協議標準實現
  - pytest: TDD 測試框架
  - pydantic: 數據驗證

architecture:
  - src/mcp_server: MCP 服務主體
  - src/tools: TDD 工具集實現
  - src/workflow: 工作流管理

deployment:
  - Python 包形式分發
  - 用戶根據 AI 工具自行配置
  - 配置模板和文檔
```

---

## 📊 成功指標重新定義

### Golang 終端代理成功指標
```yaml
functionality:
  - 支持 4+ 主流 AI 終端工具
  - Ollama 優化響應時間 < 2s
  - 終端切換延遲 < 100ms
  - Token 節省率 > 30%

usability:
  - 新用戶 5 分鐘上手
  - 跨平台一鍵安裝
  - 終端崩潰自動恢復
  - 離線模式可用 (僅 Ollama 功能)

performance:
  - 記憶體使用 < 100MB
  - CPU 使用率 < 5% (空閒時)
  - 支持同時管理 4+ 終端
  - 24小時穩定運行
```

### Python MCP 成功指標
```yaml
compatibility:
  - 支持主流 MCP 兼容 AI 工具
  - 標準 MCP 協議合規
  - 配置模板可用性 > 90%

workflow:
  - TDD 流程驗證準確率 > 95%
  - 開發門禁誤報率 < 5%
  - 工作流狀態轉換正確性 100%

adoption:
  - 配置文檔完整度
  - 用戶集成成功率 > 80%
  - 工具執行穩定性 > 95%
```

---

這個糾正版開發計劃現在完全符合您的架構理念：

1. **Golang 終端代理**：主要產物，GUI + 多終端 + Ollama 優化
2. **Python MCP**：獨立可選產物，TDD 工具 + 用戶自配置
3. **完全解耦**：兩者各自獨立開發、部署、使用

您覺得這個方向是否正確？我們可以開始專注於 Golang 終端代理的開發了嗎？