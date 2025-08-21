# 命令職責矩陣 v3.1

## 📊 命令與文件交互矩陣

| 命令 | CLAUDE.md | DECISIONS.md | PROJECT_CONTEXT.md | last-session.yml | initial-scan.json |
|------|-----------|--------------|-------------------|-----------------|-------------------|
| `/start` | ❌ | ✅ 創建 | ✅ 創建 | ❌ | ✅ 創建 |
| `/meta` | ✅ 創建 | ❌ | ❌ | ❌ | ❌ |
| `/sync` | ✅ 讀取 | ✅ 讀取 | ✅ 讀取 | ✅ 讀取 | ❌ |
| `/context` | ✅ 讀取 | ✅ 讀取 | ✅ 讀取 | ✅ 讀取 | ✅ 讀取 |
| `/plan` | ❌ | ❌ | ❌ | ✅ 創建/更新 | ❌ |
| `/learn` | ❌ | ✅ 更新 | ✅ 更新 | ❌ | ❌ |
| `/update-spec` | ✅ 更新 | ✅ 更新 | ✅ 更新 | ✅ 創建/更新 | ❌ |
| `/check` | ❌ | ❌ | ❌ | ❌ | ❌ |
| `/test` | ❌ | ❌ | ❌ | ❌ | ❌ |
| `/doc` | ❌ | ❌ | ❌ | ❌ | ❌ |
| `/review` | ❌ | ❌ | ❌ | ❌ | ❌ |
| `/debug` | ❌ | ❌ | ❌ | ❌ | ❌ |

## 🔄 標準工作流程

### 1. 項目初始化
```
/meta → 創建 CLAUDE.md（項目規範）
/start → 創建 PROJECT_CONTEXT.md, DECISIONS.md, initial-scan.json
```

### 2. 日常開發週期
```
/sync → 讀取所有狀態，恢復工作
  ↓
/context → 確認理解，查看進展（對比 initial-scan.json）
  ↓
/plan "任務" → 開始新週期（創建/更新 last-session.yml）
  ↓
[開發工作]
  ↓
/learn "決策" → 記錄重要決策（更新 DECISIONS.md）
  ↓
/update-spec → 週期終結（更新所有狀態文件）
  ↓
git commit
```

## 📁 文件職責說明

### 核心規範文件
- **CLAUDE.md** - 項目協作規範（AI行為指導）
  - 創建：`/meta`
  - 更新：`/update-spec`
  - 讀取：`/sync`, `/context`

### 狀態管理文件
- **PROJECT_CONTEXT.md** - 項目當前狀態
  - 創建：`/start`
  - 更新：`/learn`, `/update-spec`
  - 讀取：`/sync`, `/context`

- **DECISIONS.md** - 技術決策記錄
  - 創建：`/start`
  - 更新：`/learn`, `/update-spec`
  - 讀取：`/sync`, `/context`

### 會話狀態文件
- **last-session.yml** - 上次會話狀態
  - 創建/更新：`/plan`, `/update-spec`
  - 讀取：`/sync`, `/context`

- **initial-scan.json** - 項目初始基線
  - 創建：`/start`
  - 讀取：`/context`（用於進展對比）

## 🎯 關鍵改進 (v3.1)

1. **斷鏈修復**
   - `/plan` 現在會更新 last-session.yml
   - 週期開始和結束都有狀態記錄

2. **孤立文件激活**
   - initial-scan.json 被 `/context` 用於顯示項目進展
   - 所有創建的文件都有明確用途

3. **錯誤處理增強**
   - `/sync` 優雅處理文件缺失
   - 智能降級和恢復建議

## ⚠️ 未來改進建議

1. **測試結果持久化**
   - `/test` 和 `/check` 可以保存結果到 `.claude/state/test-results.yml`

2. **命令歷史追踪**
   - 可以添加 `.claude/state/command-history.yml` 記錄所有命令執行

3. **增量更新優化**
   - `/learn` 可以輕量更新 last-session.yml（只添加決策）

---
*版本：3.1.0*
*更新日期：2025-01-21*
*作者：Claude Code 命令系統團隊*