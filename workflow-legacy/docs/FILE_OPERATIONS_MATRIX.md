# 核心文件操作矩陣

本文檔詳細記錄了所有 commands、agents、styles 對三個核心文件的操作情況。

## 📊 操作矩陣總覽

### 圖例說明
- 📖 READ: 讀取文件
- ✏️ WRITE: 寫入文件  
- 🔄 UPDATE: 更新文件
- ✅ CREATE: 創建文件（如果不存在）
- ❌ 無操作

## 1️⃣ PROJECT_CONTEXT.md 操作矩陣

| 組件 | 操作類型 | 具體行為 |
|------|---------|---------|
| **Commands** | | |
| `/start` | ✅🔄 CREATE/UPDATE | 初始化項目上下文，創建完整的項目信息 |
| `/sync` | 📖 READ | 讀取項目狀態用於恢復會話 |
| `/context` | 📖 READ | 讀取並顯示當前項目上下文 |
| `/learn` | 🔄 UPDATE | 智能更新項目狀態（如果有重大變更） |
| `/meta` | 📖 READ | 參考項目信息來創建規範 |
| `/update-spec` | 🔄 UPDATE | 更新項目狀態和進度信息 |
| **Agents** | ❌ | 無直接操作 |
| **Styles** | ❌ | 無直接操作 |

## 2️⃣ DECISIONS.md 操作矩陣

| 組件 | 操作類型 | 具體行為 |
|------|---------|---------|
| **Commands** | | |
| `/start` | ✅ CREATE | 創建基礎決策記錄框架 |
| `/sync` | 📖 READ | 讀取技術決策用於理解項目 |
| `/context` | 📖 READ | 讀取決策記錄提供上下文 |
| `/learn` | ✏️🔄 WRITE/UPDATE | 記錄新的技術決策（主要寫入者） |
| `/meta` | 📖 READ | 參考決策來生成規範 |
| `/update-spec` | 🔄 UPDATE | 記錄週期內的技術決策 |
| **Agents** | ❌ | 無直接操作 |
| **Styles** | ❌ | 無直接操作 |

## 3️⃣ last-session.yml 操作矩陣

| 組件 | 操作類型 | 具體行為 |
|------|---------|---------|
| **Commands** | | |
| `/plan` | ✏️ WRITE | 創建/更新會話狀態，記錄任務規劃 |
| `/sync` | 📖 READ | 讀取上次會話狀態用於恢復 |
| `/context` | 📖 READ | 讀取會話狀態了解當前進度 |
| `/update-spec` | ✏️ WRITE | 保存完整的週期狀態 |
| **Agents** | ❌ | 無直接操作 |
| **Styles** | ❌ | 無直接操作 |

## 🔍 詳細操作分析

### 主要寫入者
1. **PROJECT_CONTEXT.md**
   - 主要寫入：`/start`（初始化）、`/update-spec`（更新）
   - 智能更新：`/learn`（僅在有重大變更時）

2. **DECISIONS.md**
   - 主要寫入：`/learn`（記錄決策）
   - 週期更新：`/update-spec`（批量記錄）
   - 初始化：`/start`（創建框架）

3. **last-session.yml**
   - 主要寫入：`/plan`（任務規劃）、`/update-spec`（週期結束）
   - 無初始化命令（按需創建）

### 主要讀取者
- `/sync`: 讀取所有三個文件（完整狀態恢復）
- `/context`: 讀取所有三個文件（理解上下文）
- `/meta`: 讀取 PROJECT_CONTEXT.md 和 DECISIONS.md（生成規範）

### 無操作組件
- **所有 Agents**: 不直接操作這些文件
- **所有 Output Styles**: 不直接操作這些文件
- **SDK Commands**: 不操作這些文件

## 💡 關鍵發現

### 1. 職責分離清晰
- **寫入職責**：集中在少數命令（start、learn、plan、update-spec）
- **讀取職責**：主要是狀態恢復命令（sync、context）

### 2. 生命週期管理
```
初始化：/start → 創建 PROJECT_CONTEXT.md、DECISIONS.md
日常記錄：/learn → 更新 DECISIONS.md
任務規劃：/plan → 更新 last-session.yml
週期結束：/update-spec → 更新所有三個文件
新會話：/sync → 讀取所有三個文件
```

### 3. 文件依賴關係
- PROJECT_CONTEXT.md：項目級別，變更最少
- DECISIONS.md：持續累積，只增不減
- last-session.yml：會話級別，頻繁更新

## 🎯 整合建議

基於這個矩陣，如果要整合這三個文件的操作，建議考慮：

1. **統一寫入接口**：創建一個核心服務來管理所有寫入操作
2. **版本控制**：為每個文件添加版本號，便於追踪變更
3. **原子操作**：確保多文件更新的原子性（要麼全部成功，要麼全部失敗）
4. **緩存機制**：頻繁讀取的文件可以考慮緩存
5. **事件驅動**：文件變更時觸發相關命令更新

## 📝 命令操作詳情

### /start（初始化專家）
- PROJECT_CONTEXT.md: CREATE/UPDATE - 創建完整項目信息
- DECISIONS.md: CREATE - 創建決策記錄框架
- last-session.yml: ❌

### /sync（狀態恢復專家）
- PROJECT_CONTEXT.md: READ - 恢復項目信息
- DECISIONS.md: READ - 載入決策歷史
- last-session.yml: READ - 恢復會話狀態

### /learn（知識管理專家）
- PROJECT_CONTEXT.md: UPDATE（智能）- 僅重大變更時更新
- DECISIONS.md: WRITE - 主要寫入者，記錄決策
- last-session.yml: ❌

### /plan（任務規劃專家）
- PROJECT_CONTEXT.md: ❌
- DECISIONS.md: ❌
- last-session.yml: WRITE - 記錄任務規劃

### /update-spec（週期終結者）
- PROJECT_CONTEXT.md: UPDATE - 更新項目進度
- DECISIONS.md: UPDATE - 批量記錄決策
- last-session.yml: WRITE - 保存完整狀態

### /context（上下文同步）
- PROJECT_CONTEXT.md: READ
- DECISIONS.md: READ
- last-session.yml: READ

### /meta（規範制定）
- PROJECT_CONTEXT.md: READ - 參考項目信息
- DECISIONS.md: READ - 參考技術決策
- last-session.yml: ❌

---

*最後更新：2025-01-26*
*用於支持核心文件操作整合計劃*