# Claude Code 協作規範體系 v3.0

本文檔整合了 Claude Code 的所有核心內容，包括命令系統、協作規範、使用指南等。

## 📚 目錄

1. [快速開始](#-快速開始)
2. [命令系統 v3.0](#-命令系統-v30)
3. [項目結構](#-項目結構)
4. [協作宪法](#-協作宪法)
5. [使用指南](#-使用指南)
6. [工作流程](#-工作流程)
7. [最佳實踐](#-最佳實踐)
8. [版本歷史](#-版本歷史)

---

## 🚀 快速開始

### 5 分鐘內開始使用 Claude Code

#### A. 新項目
```bash
# 使用元工作流命令
/meta

# Claude 會：
# 1. 分析項目特徵
# 2. 詢問關鍵信息
# 3. 自動生成 CLAUDE.md
# 4. 設置文檔結構
```

#### B. 現有項目
```bash
# 直接理解項目
/start

# 或恢復之前的工作狀態
/sync
```

#### C. 安裝命令系統
```bash
# Windows
cd claude\commands\deploy-package
.\deploy.ps1

# macOS/Linux
cd claude/commands/deploy-package
./deploy.sh
```

---

## 🎯 命令系統 v3.0 + SDK 擴展

### 核心理念
- **少即是多**：核心命令 11 個 + SDK 專用 5 個
- **智能整合**：每個命令完成多項相關任務
- **上下文感知**：自動管理記憶和狀態
- **場景適配**：應用開發和 SDK 開發雙軌支援

### 11 個核心命令

#### 項目理解與管理（3個）
| 命令 | 功能 | 使用時機 | 參數 |
|------|------|----------|------|
| `/start` | 項目快速啟動與理解 | 初次接觸項目 | 無 |
| `/context` | 上下文同步檢查點 | 確保理解一致 | 無 |
| `/sync` | 狀態同步器 | 新會話開始 | 無 |

#### 開發輔助（3個）
| 命令 | 功能 | 使用時機 | 參數 |
|------|------|----------|------|
| `/plan` | 任務規劃與設計 | 開始新功能前 | [任務描述] |
| `/check` | 智能代碼審查 | 提交代碼前 | 無 |
| `/test` | 測試生成與執行 | 確保代碼質量 | [文件\|功能] |

#### 知識管理（2個）
| 命令 | 功能 | 使用時機 | 參數 |
|------|------|----------|------|
| `/learn` | 學習並記錄決策 | 重要決定後 | [決策內容] |
| `/doc` | 智能文檔維護 | 更新項目文檔 | [api\|readme\|changelog\|arch] |

#### 工作流優化（3個）
| 命令 | 功能 | 使用時機 | 參數 |
|------|------|----------|------|
| `/review` | PR 準備助手 | 創建 PR 前 | 無 |
| `/debug` | 智能調試助手 | 遇到問題時 | [錯誤信息] |
| `/meta` | 項目規範定制 | 新項目或重大變更 | 無 |

### 🆕 SDK 開發專用命令（5個）

專為 SDK/Library 開發設計的命令集：

#### SDK 專用命令
| 命令 | 功能 | 使用時機 | 參數 |
|------|------|----------|------|
| `/sdk-design` | API 設計助手 | 設計新 API 時 | [功能描述] |
| `/sdk-example` | 示例代碼生成 | 創建使用示例 | [basic\|advanced\|integration\|all] |
| `/sdk-test` | SDK 測試套件 | 生成專業測試 | [unit\|integration\|compat\|performance\|all] |
| `/sdk-doc` | SDK 文檔生成 | 編寫文檔時 | [api\|guide\|migration\|all] |
| `/sdk-release` | 發布準備助手 | 準備新版本 | [major\|minor\|patch\|check] |

### 項目專用命令（可選）
這些命令在 `.claude/commands/project/` 中：

| 命令 | 功能 | 說明 |
|------|------|------|
| `/ai-rules` | AI 規則管理 | check/apply/validate/report |
| `/guardian` | 主動監控系統 | 實時質量/安全/性能監控 |
| `/sync` | 憲法同步系統 | 版本管理/衝突解決 |

---

## 📁 項目結構

```
claude/
├── README.md                    # 本文檔（所有核心內容整合版）
│
├── constitution/               # 憲法體系（參考）
│   └── CLAUDE_CONSTITUTION.md  # Claude 協作憲法完整版
│
├── commands/                   # 命令系統
│   └── deploy-package/         # 命令部署包 v3.0
│       ├── DEPLOY_GUIDE.md     # 部署指南
│       ├── CHANGELOG.md        # 版本歷史
│       ├── SIMPLE_COMMANDS_SUMMARY.md # 命令系統說明
│       ├── deploy.ps1          # Windows 部署腳本
│       ├── deploy.sh           # macOS/Linux 部署腳本
│       ├── global/             # 11 個全局命令
│       │   ├── check.md
│       │   ├── context.md
│       │   ├── debug.md
│       │   ├── doc.md
│       │   ├── learn.md
│       │   ├── meta.md
│       │   ├── plan.md
│       │   ├── review.md
│       │   ├── start.md
│       │   ├── sync.md
│       │   └── test.md
│       └── project/            # 3 個項目命令
│           ├── ai-rules.md
│           ├── guardian.md
│           └── sync.md
│
├── guides/                     # 深入指南（進階參考）
│   ├── AI_ASSISTANT_COMPARISON.md    # AI 助手對比
│   ├── COMMAND_WRITING_GUIDE.md       # 命令編寫指南
│   ├── CONSTITUTION_SYNC_GUIDE.md     # 憲法同步指南
│   ├── CONSTITUTION_USAGE_GUIDE.md    # 憲法使用指南
│   ├── DOCUMENT_STRUCTURE_STANDARD.md # 文檔結構規範
│   ├── LEGACY_PROJECT_ONBOARDING.md   # 遺留項目接入
│   ├── MARKET_ANALYSIS.md             # 市場分析
│   ├── NEW_VS_LEGACY_PROJECT.md       # 新舊項目對比
│   └── SDK_DEVELOPMENT_WORKFLOW.md    # SDK 開發工作流
│
└── templates/                  # 模板文件
    ├── CLAUDE_MD_TEMPLATE.md   # CLAUDE.md 通用模板
    └── SDK_PROJECT_TEMPLATE.md # SDK 項目專用模板
```

### 📝 文檔說明
- **README.md**：包含所有核心內容，日常使用只需查看此文件
- **commands/deploy-package/**：用於部署命令系統
- **guides/**：深入的專題指南，需要時參考
- **templates/**：項目初始化時使用的模板

---

## 🏛️ 協作憲法

### 核心理念
Claude Code 的協作基於以下原則：

1. **上下文優先**：保持理解的連續性
2. **知識積累**：記錄決策，避免重複
3. **漸進改進**：小步快跑，持續優化
4. **人機協作**：明確分工，發揮各自優勢

### 工作模式

#### 開發工作流
1. **理解階段**：`/start` 或 `/sync`
2. **規劃階段**：`/plan` 任務分解
3. **實施階段**：編碼實現
4. **驗證階段**：`/check` 和 `/test`
5. **知識沉澱**：`/learn` 記錄決策

#### 元工作流
用於建立和更新項目規範：
1. **評估項目**：技術棧、團隊、複雜度
2. **定制規範**：生成項目特定的 CLAUDE.md
3. **持續優化**：根據實踐調整規範

---

## 📖 使用指南

### 典型工作流程

#### 1. 新項目初始化

**應用開發**：
```bash
/meta               # 建立項目規範
/start              # 理解項目結構
/plan "核心功能"     # 規劃第一個任務
```

**SDK 開發**：
```bash
/meta               # 建立 SDK 規範
/sdk-design "核心 API"  # 設計接口
/plan "實現核心功能"    # 規劃實現
```

#### 2. 日常開發循環
```bash
/sync               # 恢復工作狀態
/context            # 確認理解正確
/plan "新功能"       # 規劃實現方案
# ... 編碼 ...
/check              # 代碼質量檢查
/test               # 生成並運行測試
/learn "技術決策"    # 記錄重要決定
```

#### 3. 提交與發布
```bash
/check              # 最終質量檢查
/doc                # 更新相關文檔
/review             # 準備 PR
```

#### 4. 問題解決
```bash
/debug "錯誤信息"    # 快速定位問題
/test feature       # 驗證修復效果
```

### 記憶管理策略

#### 項目上下文 (`.claude/PROJECT_CONTEXT.md`)
```markdown
# 項目上下文

## 🎯 項目願景
- 我想做什麼
- 為什麼要做
- 最終效果

## 🏗️ 技術架構
- 技術棧選擇理由
- 核心設計決策
- 已知限制

## 📊 當前狀態
- 已完成
- 進行中
- 待實現
```

#### 決策記錄 (`.claude/DECISIONS.md`)
```markdown
# 技術決策記錄

## 日期：決策標題
**決策**：具體內容
**原因**：為什麼這樣決定
**影響**：會帶來什麼變化
```

---

## 🔄 工作流程

### 開發流程對比

#### 應用開發流程
| 階段 | 傳統方式 | Claude Code v3.0 |
|------|----------|------------------|
| 開始 | 手動說明背景 | `/sync` 自動恢復 |
| 理解 | 重複解釋 | `/context` 同步確認 |
| 規劃 | 自由討論 | `/plan` 結構化設計 |
| 開發 | 獨立編碼 | AI 協助實現 |
| 測試 | 手動編寫 | `/test` 智能生成 |
| 審查 | 人工檢查 | `/check` 自動審查 |
| 文檔 | 事後補充 | `/doc` 同步更新 |
| 知識 | 容易遺忘 | `/learn` 持久記錄 |

#### SDK 開發流程
| 階段 | 傳統方式 | Claude Code + SDK 命令 |
|------|----------|------------------------|
| API 設計 | 憑經驗設計 | `/sdk-design` 專業指導 |
| 示例編寫 | 手動創建 | `/sdk-example` 自動生成 |
| 測試策略 | 基礎測試 | `/sdk-test` 全面覆蓋 |
| 文檔編寫 | 耗時費力 | `/sdk-doc` 結構化生成 |
| 版本發布 | 容易遺漏 | `/sdk-release` 完整檢查 |

### 命令協作示例

```bash
# 場景：開發用戶認證功能

# 1. 開始工作
/sync
> 恢復上次狀態：正在開發用戶系統

# 2. 規劃新功能
/plan "實現 JWT 認證"
> 生成實施方案：
> - 設計 token 結構
> - 實現登錄接口
> - 添加中間件
> - 編寫測試

# 3. 開發過程中遇到問題
/debug "JWT token 驗證失敗"
> 分析原因：token 過期時間配置錯誤
> 提供解決方案

# 4. 完成開發
/check
> 代碼風格 ✓
> 安全檢查 ✓
> 性能分析 ✓

# 5. 記錄決策
/learn "選擇 JWT 而非 session 因為需要支援分佈式"
```

---

## 💡 最佳實踐

### 1. 溝通技巧
- **明確邊界**：告訴 Claude 不要修改什麼
- **提供示例**：給出期望的代碼風格
- **分段確認**：複雜任務分成多個檢查點
- **記錄決策**：重要選擇要寫入文檔

### 2. 項目組織
```
your-project/
├── .claude/
│   ├── commands/           # 項目特定命令
│   ├── PROJECT_CONTEXT.md  # 項目上下文
│   ├── DECISIONS.md        # 決策記錄
│   └── state/              # 狀態文件
├── CLAUDE.md               # 項目規範
└── ... 項目文件
```

### 3. 效率提升
- **開始即同步**：每次都用 `/sync` 開始
- **及時記錄**：用 `/learn` 避免知識流失
- **結構化規劃**：用 `/plan` 而非自由討論
- **自動化檢查**：用 `/check` 保證質量

### 4. 團隊協作
- 共享 `.claude/` 目錄
- 統一使用命令系統
- 定期更新 PROJECT_CONTEXT.md
- 重要決策都記錄在 DECISIONS.md

---

## 📊 版本歷史

### v3.0.0 (2024-01-15) - 當前版本
- **重大重構**：命令系統精簡優化
- **核心改進**：
  - 命令從 31 個減少到 11 個
  - 智能命令整合
  - 自動化記憶管理
  - 結構化項目上下文
- **新增功能**：
  - PROJECT_CONTEXT.md 系統
  - DECISIONS.md 決策記錄
  - 智能調試助手
  - PR 準備助手

### v2.1.0 (2024-01-14)
- 參數規範化
- 命令協調機制
- 項目命令增強

### v2.0.0 (2024-01-13)
- 初始版本
- 31 個命令
- 基礎部署系統

---

## 🤝 貢獻指南

歡迎貢獻！請：
1. Fork 本項目
2. 創建功能分支
3. 提交變更
4. 發起 Pull Request

### 報告問題
- 使用 GitHub Issues
- 提供詳細的復現步驟
- 說明期望行為

---

## 📚 相關資源

- [Claude Code 官方文檔](https://docs.anthropic.com/en/docs/claude-code)
- [命令系統文檔](https://docs.anthropic.com/en/docs/claude-code/slash-commands)
- [MCP 協議](https://docs.anthropic.com/en/docs/claude-code/mcp)
- [項目 Hook 模板](PROJECT_HOOKS_TEMPLATE.md)

---

## 🎯 核心價值

1. **簡單高效**：更少的命令，更高的效率
2. **智能協作**：AI 理解你，你引導 AI
3. **知識積累**：每個決策都是財富
4. **持續進化**：根據使用不斷優化

---

*本規範體系由 Claude 與人類開發者共同創建，持續演進中。*

*簡單、智能、高效 - 讓 Claude Code 成為你的最佳開發夥伴！*