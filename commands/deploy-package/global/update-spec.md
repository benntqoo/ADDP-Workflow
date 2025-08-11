---
arguments: optional
format: "[action] [content]"
actions: "review | version | backup | [section] [content]"
examples:
  - "/update-spec review  # 審查近期決策並建議更新"
  - "/update-spec version  # 查看 CLAUDE.md 版本歷史"
  - "/update-spec backup  # 備份當前 CLAUDE.md"
  - "/update-spec rules '新增 GraphQL API 命名規範'"
  - "/update-spec workflow '調整 PR review 流程'"
  - "/update-spec architecture '引入 Redis 緩存層'"
---

# 更新項目規範 (CLAUDE.md) - 單一職責命令

我是專門負責更新 CLAUDE.md 的命令，確保項目規範與演進保持同步。

## 🎯 核心職責

**我只做一件事**：智能更新 CLAUDE.md
- ✅ 讀取、分析、更新 CLAUDE.md
- ✅ 維護版本歷史和變更記錄
- ❌ 不會修改 DECISIONS.md（那是 `/learn` 的職責）
- ❌ 不會修改 PROJECT_CONTEXT.md（那是 `/start` 和 `/sync` 的職責）

## 📋 使用模式

### 模式 1：審查建議模式（無參數）
```bash
/update-spec review
```
**功能**：
1. 掃描最近的 DECISIONS.md 記錄
2. 分析哪些決策應該固化為規範
3. 提供具體的更新建議
4. 生成待執行的更新命令列表

### 模式 2：定向更新模式（帶參數）
```bash
/update-spec [section] "[content]"
```
**功能**：
1. 直接更新指定章節
2. 智能合併相似規則
3. 維護文檔一致性

## 🔧 支援的更新區域

| 參數 | 更新區域 | 說明 | 示例 |
|------|----------|------|------|
| `overview` | 項目概述 | 項目描述、目標、範圍 | 調整項目定位 |
| `architecture` | 技術架構 | 技術棧、架構模式、依賴 | 引入新框架 |
| `rules` | 開發規範 | 編碼規範、命名約定 | API 命名規則 |
| `workflow` | 工作流程 | 開發流程、協作模式 | PR 審查流程 |
| `behavior` | Claude 行為 | AI 協作原則、輸出風格 | 代碼風格偏好 |
| `testing` | 測試策略 | 測試要求、覆蓋率 | 單元測試規範 |
| `security` | 安全規範 | 安全要求、最佳實踐 | 認證授權規則 |
| `performance` | 性能標準 | 性能指標、優化原則 | 響應時間要求 |
| `custom` | 項目特定 | LOCAL 區域的內容 | 業務特定規則 |

## 📝 執行流程

### 1. 分析階段
```
讀取 CLAUDE.md
    ↓
解析文檔結構
    ↓
識別目標章節
    ↓
分析現有內容
```

### 2. 更新階段
```
智能判斷更新類型（新增/修改/替換）
    ↓
保持格式一致性
    ↓
添加版本標記
    ↓
執行更新
```

### 3. 驗證階段
```
檢查語法正確性
    ↓
確認沒有破壞結構
    ↓
保存備份（如需要）
```

## 🎨 智能特性

### 1. 重複檢測
- 自動識別相似規則
- 提示是否合併或替換
- 避免規範冗餘

### 2. 衝突預警
- 檢測矛盾的規則
- 提示潛在問題
- 建議解決方案

### 3. 版本管理
```markdown
<!-- 
Claude Constitution Version: 3.0.1 → 3.0.2
Last Updated: 2024-01-20
Change Log:
  - Added: GraphQL API naming convention
  - Modified: PR review process
-->
```

### 4. 智能分類
自動將更新內容放到最合適的章節，即使用戶指定錯誤

## 📝 使用示例

### 示例 1：審查模式（推薦每週執行）
```bash
/update-spec review
```
**輸出**：
```
📊 分析了最近 7 個決策，發現 3 個可固化為規範：

1. ✅ API 命名約定（2024-01-18）
   建議執行：/update-spec rules "所有 API endpoint 使用 kebab-case"
   
2. ✅ Redis 緩存策略（2024-01-17）
   建議執行：/update-spec architecture "引入 Redis 作為會話和熱數據緩存"
   
3. ⚠️ PR 審查流程（2024-01-16）
   已存在相似規則，建議修改而非新增
   建議執行：/update-spec workflow "PR 需要至少 2 個 approval"
```

### 示例 2：添加編碼規範
```bash
/update-spec rules "所有 API endpoint 使用 kebab-case 命名"
```
**結果**：
- ✅ 在"開發規範"章節添加 API 命名規則
- ✅ 更新版本號：3.0.1 → 3.0.2
- ✅ 添加變更記錄

### 示例 3：調整工作流程
```bash
/update-spec workflow "PR 需要至少一個 approval 才能合併"
```
**結果**：
- ✅ 更新"工作流程"章節的 PR 規則
- ⚠️ 檢測到與現有規則衝突，提示確認
- ✅ 保留原有流程的其他部分

### 示例 4：更新技術架構
```bash
/update-spec architecture "引入 Redis 作為會話存儲和緩存層"
```
**結果**：
- ✅ 在技術架構添加 Redis 說明
- ✅ 更新依賴列表
- ✅ 添加相關配置說明

## 🔄 版本管理

每次更新都會：
1. 在文檔頂部更新版本號
2. 添加變更註釋
3. 保留歷史記錄

```markdown
<!-- 
Claude Constitution Version: 3.0.1
Last Updated: 2024-01-20
Changes: 添加 GraphQL API 規範
-->
```

## 💡 最佳實踐

### 推薦工作流
```bash
# 每週/迭代結束時
1. /update-spec review        # 審查近期決策
2. 執行建議的更新命令          # 固化重要規範
3. git commit CLAUDE.md       # 版本控制
4. /context                   # 確認理解
```

### 使用原則
1. **職責單一**：只更新 CLAUDE.md，不碰其他文件
2. **小步快跑**：每次專注一個方面的更新
3. **及時固化**：重要決策及時轉為規範
4. **保持簡潔**：定期清理過時內容

## 🎯 與其他命令的協作

### 命令職責矩陣
| 命令 | CLAUDE.md | DECISIONS.md | PROJECT_CONTEXT.md |
|------|-----------|--------------|-------------------|
| `/meta` | ✅ 創建 | ❌ | ❌ |
| `/update-spec` | ✅ 更新 | ❌ | ❌ |
| `/learn` | ❌ | ✅ 更新 | ✅ 智能更新 |
| `/start` | ❌ | ❌ | ✅ 創建/更新 |
| `/sync` | ❌ | ❌ | ✅ 讀取 |
| `/context` | ✅ 讀取 | ✅ 讀取 | ✅ 讀取 |

### 協作流程
```
/learn "新決策"           # 記錄到 DECISIONS.md
      ↓
/update-spec review      # 分析哪些需要固化
      ↓
/update-spec rules "..." # 更新到 CLAUDE.md
      ↓
/context                 # 確認全部理解
```

## 💡 注意事項

1. **不會自動連動**：此命令專注於 CLAUDE.md，不會自動更新其他文件
2. **手動觸發**：需要你主動執行，不會被其他命令自動調用
3. **保持獨立**：與 `/learn` 分工明確，各司其職
4. **版本控制**：建議每次更新後 git commit

準備好管理你的項目規範了嗎？使用 `/update-spec review` 開始！