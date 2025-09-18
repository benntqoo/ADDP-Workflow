---
arguments: optional
format: "[summary]"
examples:
  - "/update-spec  # 結束當前開發週期，更新所有規範"
  - "/update-spec 'Repository層優化完成'  # 帶總結的週期終結"
---

# 開發週期終結與規範更新

我是開發週期的終結儀式，負責總結成果、更新規範、保存狀態，為下一個週期做準備。

## 🎯 核心職責

**週期終結的完整儀式**：
- ✅ 總結本次開發週期成果
- ✅ 更新 CLAUDE.md（項目規範）
- ✅ 更新 PROJECT_CONTEXT.md（項目狀態）
- ✅ 更新 DECISIONS.md（技術決策）
- ✅ 創建/更新 last-session.yml（會話狀態）
- ✅ 生成週期報告並歸檔

## 📋 執行流程

### 1. 收集週期信息
```yaml
# 自動分析獲取
- 週期起止時間（從上次 update-spec 或 /plan）
- Git 提交歷史和變更統計
- TodoWrite 完成情況
- 測試和代碼質量指標
```

### 2. 總結開發成果
```yaml
# 自動生成
achievements:
  features_added: [從 git 和 todos 提取]
  issues_fixed: [從 commit messages 提取]
  tests_added: [統計新增測試]
  quality_metrics: [測試覆蓋率、Detekt 變化]
```

### 3. 更新項目文檔

<!-- File Operations: Direct Read/Write -->
<!-- Target: .claude/memory/ directory -->

#### 直接批量更新記憶文件

```bash
# 確保目錄存在
mkdir -p .claude/memory
```

**更新流程**：

1. **PROJECT_CONTEXT.md 更新**
   - 讀取現有文件（智能搜索）
   - 更新版本號、進度狀態
   - 添加新增功能
   - 寫回 `.claude/memory/PROJECT_CONTEXT.md`

2. **DECISIONS.md 更新**
   - 讀取現有決策
   - 添加本週期新決策
   - 保持時間順序
   - 寫回 `.claude/memory/DECISIONS.md`

3. **last-session.yml 更新**
   - 記錄週期完成狀態
   - 總結成就和下一步
   - 寫回 `.claude/memory/last-session.yml`

**文件操作保證**：
- ✅ 按順序更新，確保一致性
- ✅ 每個文件先讀後寫
- ✅ 保持格式一致
- ✅ 智能搜索舊位置

### 4. 生成週期報告
創建 `.claude/reports/[date]-[summary].md` 包含：
- 週期概覽和關鍵指標
- 達成目標和遇到的挑戰
- 經驗總結和下一步建議

### 5. 歸檔當前週期
```bash
.claude/archive/cycles/[date]/
├── session.yml      # 會話狀態快照
├── report.md       # 週期報告
├── git-diff.patch  # 代碼變更
└── metrics.json    # 度量數據
```

## 💾 last-session.yml 結構

```yaml
session:
  cycle:
    id: "[date]-[topic]"
    type: "feature|bugfix|refactor|optimization"
    description: "週期主要工作"
    started_at: "開始時間"
    completed_at: "結束時間"
    
  summary:
    main_achievement: "主要成果"
    key_decisions: ["決策列表"]
    challenges_faced: ["遇到的挑戰"]
    solutions_applied: ["解決方案"]
    
  changes:
    new_files: ["新建文件列表"]
    modified_files: ["修改文件列表"]
    deleted_files: ["刪除文件列表"]
    
  knowledge_gained:
    - topic: "學到的知識點"
      insight: "具體見解"
      
  next_cycle:
    status: "ready|blocked|pending_review"
    suggestions: ["下一步建議"]
    pending_tasks: ["遺留任務"]
    
  metrics:
    commits_count: 數量
    test_coverage: "百分比"
    code_quality: "指標變化"
    performance_gain: "性能提升"
```

## 🎨 智能特性

### 1. 自動週期識別
- 從上次 update-spec 或 /plan 計算週期
- 自動收集期間所有變更
- 智能分類成果類型

### 2. 規範提取
- 從 DECISIONS.md 識別可固化的規範
- 從代碼變更推斷最佳實踐
- 自動生成規範建議

### 3. 狀態連續性
- 讀取上個週期的遺留任務
- 追踪長期目標進展
- 維護知識積累鏈條

### 4. 智能歸檔
- 自動創建週期目錄結構
- 保存所有相關快照
- 生成可追溯的歷史記錄

## 📝 使用示例

### 示例 1：無參數執行（標準週期終結）
```bash
/update-spec
```
**輸出**：
```
📊 開發週期總結 [2025-01-21]

✅ 本週期成果：
- 完成：Repository層N+1查詢優化
- 新增：15個測試用例
- 性能：查詢速度提升300%
- 質量：Detekt問題 -20

📝 已更新文檔：
- CLAUDE.md（新增緩存策略規範）
- PROJECT_CONTEXT.md（更新項目進度）
- DECISIONS.md（記錄3個技術決策）
- last-session.yml（保存週期狀態）

📁 週期報告：
.claude/reports/2025-01-21-optimization.md

💡 下一週期建議：
- 可以開始新功能開發
- 建議先運行完整測試套件
- Detekt問題仍需關注（當前420個）

使用 /sync 開始新的工作週期
```

### 示例 2：帶總結的執行
```bash
/update-spec "完成用戶認證系統重構"
```
**輸出**：
```
📊 開發週期總結：完成用戶認證系統重構

✅ 識別的主要變更：
- 重構：認證流程從Session改為JWT
- 新增：設備管理功能
- 優化：Token刷新機制

[其餘輸出同上]
```

### 示例 3：週期中的異常情況
```bash
/update-spec "緊急修復生產環境bug"
```
**輸出**：
```
⚠️ 檢測到非正常週期（僅2小時）

📊 緊急修復總結：
- 問題：生產環境內存洩漏
- 解決：修復Repository層資源釋放
- 影響：已部署並驗證

📝 已記錄到 DECISIONS.md：
- 緊急修復流程
- 資源管理最佳實踐

💡 建議：
- 添加內存洩漏檢測測試
- 完善監控告警機制
```

## 🔄 與工作流程的整合

### 你的日常工作流程
```mermaid
graph LR
    A[/sync 開始] --> B{有遺留?}
    B -->|是| C[/context 繼續]
    B -->|否| D[/plan 新任務]
    C --> E[開發工作]
    D --> E
    E --> F[/update-spec 終結]
    F --> G[commit 提交]
    G --> H[週期完成]
```

### 命令職責矩陣（更新後）
| 命令 | CLAUDE.md | DECISIONS.md | PROJECT_CONTEXT.md | last-session.yml |
|------|-----------|--------------|-------------------|-----------------|
| `/start` | ❌ | ❌ | ✅ 創建 | ❌ |
| `/sync` | ❌ | ❌ | ✅ 讀取 | ✅ 讀取 |
| `/context` | ✅ 讀取 | ✅ 讀取 | ✅ 讀取 | ✅ 讀取 |
| `/plan` | ❌ | ❌ | ❌ | ✅ 記錄開始 |
| `/learn` | ❌ | ✅ 更新 | ✅ 智能更新 | ❌ |
| `/update-spec` | ✅ 更新 | ✅ 更新 | ✅ 更新 | ✅ 創建/更新 |

## 💡 最佳實踐

### 推薦工作流
```bash
# 每日開始
/sync                    # 讀取 last-session.yml，恢復狀態

# 開發過程
/context                 # 確認理解
/plan "新功能"           # 開始新週期（如需要）
[開發工作...]
/learn "重要決策"        # 記錄決策（如需要）

# 週期結束
/update-spec            # 終結週期，更新所有文檔
git commit              # 提交變更
```

### 使用原則
1. **週期終結**：把 update-spec 當作開發週期的結束儀式
2. **自動保存**：所有狀態自動保存到 last-session.yml
3. **持續積累**：每個週期的知識都會被保留
4. **可追溯性**：所有週期都有完整歸檔

## 💡 注意事項

1. **自動創建**：首次執行會自動創建 last-session.yml
2. **智能識別**：自動識別週期類型（feature/bugfix/refactor）
3. **連續追踪**：維護週期間的連續性
4. **知識管理**：自動提取和保存經驗教訓

準備好使用新的週期管理系統了嗎？在完成當前工作後使用 `/update-spec` 終結週期！