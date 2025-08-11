# Claude Code 協作規範體系 v3.3

本文檔整合了 Claude Code 的所有核心內容，包括命令系統、協作規範、使用指南等。

## 📚 目錄

1. [快速開始](#-快速開始)
2. [命令系統 v3.0](#-命令系統-v30)
3. [命令詳細使用手冊](#-命令詳細使用手冊) ⭐ **新增**
4. [項目結構](#-項目結構)
5. [協作宪法](#-協作宪法)
6. [使用指南](#-使用指南)
7. [工作流程](#-工作流程)
8. [最佳實踐](#-最佳實踐)
9. [版本歷史](#-版本歷史)

### 🔍 命令快速索引

| 類別 | 命令 | 功能簡述 | 跳轉 |
|------|------|---------|------|
| **項目理解** | `/start` | 快速理解新項目 | [詳情](#1-start---項目快速啟動與理解) |
| | `/sync` | 恢復工作狀態 | [詳情](#2-sync---狀態同步器) |
| | `/context` | 確認理解一致 | [詳情](#3-context---上下文同步檢查點) |
| **開發輔助** | `/plan` | 任務規劃分解 | [詳情](#4-plan---任務規劃與設計) |
| | `/check` | 完整質量檢查 | [詳情](#5-check---完整質量檢查) |
| | `/watch` | 監察模式 | [詳情](#6-watch---監察模式) |
| | `/test` | 測試生成執行 | [詳情](#7-test---測試生成與執行) |
| **問題解決** | `/debug` | 錯誤快速定位 | [詳情](#8-debug---智能調試助手) |
| | `/analyze` | 深度風險分析 | [詳情](#13-analyze---深度分析與驗證) |
| **知識管理** | `/learn` | 記錄技術決策 | [詳情](#9-learn---學習並記錄決策) |
| | `/doc` | 文檔智能維護 | [詳情](#10-doc---智能文檔維護) |
| **工作流程** | `/review` | PR 準備助手 | [詳情](#11-review---pr-準備助手) |
| | `/meta` | 項目規範定制 | [詳情](#12-meta---項目規範定制) |
| | `/update-spec` | 規範更新管理 | [詳情](#14-update-spec---claudemd-更新專用) |
| **SDK 開發** | `/sdk-design` | API 設計指導 | [詳情](#1-sdk-design---api-設計助手) |
| | `/sdk-example` | 示例代碼生成 | [詳情](#2-sdk-example---示例代碼生成) |
| | `/sdk-test` | 測試套件生成 | [詳情](#3-sdk-test---sdk-測試套件) |
| | `/sdk-doc` | SDK 文檔生成 | [詳情](#4-sdk-doc---sdk-文檔生成) |
| | `/sdk-release` | 發布準備檢查 | [詳情](#5-sdk-release---發布準備助手) |

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

## 🎯 命令系統 v3.3 + SDK 擴展

### 核心理念
- **少即是多**：通用命令 14 個 + SDK 專用 5 個
- **智能整合**：每個命令完成多項相關任務
- **上下文感知**：自動管理記憶和狀態
- **場景適配**：應用開發和 SDK 開發雙軌支援

### 14 個通用命令（全局通用）

#### 項目理解與管理（3個）
| 命令 | 功能 | 使用時機 | 參數 |
|------|------|----------|------|
| `/start` | 項目快速啟動與理解 | 初次接觸項目 | 無 |
| `/context` | 上下文同步檢查點 | 確保理解一致 | 無 |
| `/sync` | 狀態同步器 | 新會話開始 | 無 |

#### 開發輔助（4個）
| 命令 | 功能 | 使用時機 | 參數 |
|------|------|----------|------|
| `/plan` | 任務規劃與設計 | 開始新功能前 | [任務描述] |
| `/check` | 完整質量檢查 | 提交代碼前 | 無 |
| `/watch` | 監察模式 | 編碼過程中 | [on\|off\|status\|report] |
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

#### 質量保證（2個）
| 命令 | 功能 | 使用時機 | 參數 |
|------|------|----------|------|
| `/analyze` | 深度分析與驗證 | 基於經驗直覺的風險分析 | [功能/模組] [疑慮或"deep"] |
| `/update-spec` | CLAUDE.md 更新專用 | 固化決策為規範 | [review\|section "content"] |

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
│       ├── global/             # 14 個通用命令
│       │   ├── analyze.md
│       │   ├── check.md
│       │   ├── context.md
│       │   ├── debug.md
│       │   ├── doc.md
│       │   ├── watch.md
│       │   ├── learn.md
│       │   ├── meta.md
│       │   ├── plan.md
│       │   ├── review.md
│       │   ├── start.md
│       │   ├── sync.md
│       │   ├── test.md
│       │   └── update-spec.md
│       └── sdk/                # 5 個 SDK 專用命令
│           ├── sdk-design.md
│           ├── sdk-doc.md
│           ├── sdk-example.md
│           ├── sdk-release.md
│           └── sdk-test.md
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

### 📘 命令詳細使用手冊

#### 1. `/start` - 項目快速啟動與理解

**使用場景**：
- 首次接觸一個新項目
- 需要快速理解項目結構
- 接手他人的代碼庫

**具體用法**：
```bash
/start
```

**預期輸出**：
- 項目類型識別（Web/API/SDK/工具等）
- 技術棧分析（語言、框架、依賴）
- 目錄結構解析
- 關鍵文件定位
- 自動創建 `.claude/PROJECT_CONTEXT.md`

**實際案例**：
```bash
# 接手一個 React 項目
/start
> 識別為：React Web 應用
> 技術棧：React 18, TypeScript, Vite
> 主要模組：components/, pages/, services/
> 入口文件：src/main.tsx
> 已創建項目上下文文件
```

---

#### 2. `/sync` - 狀態同步器

**使用場景**：
- 開始新的工作會話
- 切換到不同分支後
- 長時間中斷後恢復工作

**具體用法**：
```bash
/sync
```

**預期輸出**：
- 載入項目上下文
- 恢復決策記錄
- 顯示當前工作狀態
- 提醒待完成任務

**實際案例**：
```bash
/sync
> 恢復項目：電商平台 v2.0
> 當前分支：feature/payment
> 最近決策：選用 Stripe 支付
> 待完成：支付回調處理
```

---

#### 3. `/context` - 上下文同步檢查點

**使用場景**：
- 確認 Claude 理解是否正確
- 重大變更後同步認知
- 團隊成員交接時

**具體用法**：
```bash
/context
```

**預期輸出**：
- 當前理解的項目狀態
- 最近的變更摘要
- 待確認的假設

**實際案例**：
```bash
/context
> 當前理解：
> - 正在實現用戶認證模組
> - 使用 JWT token 方案
> - 需要支援 OAuth2.0
> 請確認以上理解是否正確？
```

---

#### 4. `/plan` - 任務規劃與設計

**使用場景**：
- 開始新功能開發
- 重構現有代碼
- 解決複雜問題

**具體用法**：
```bash
/plan "任務描述"
```

**預期輸出**：
- 分解的子任務列表
- 實施順序建議
- 潛在風險提示
- 時間估算

**實際案例**：
```bash
/plan "實現購物車功能"
> 任務分解：
> 1. 設計購物車數據模型 (2h)
> 2. 實現添加商品 API (1h)
> 3. 實現數量修改 API (1h)
> 4. 實現刪除商品 API (0.5h)
> 5. 添加庫存檢查邏輯 (1h)
> 6. 實現價格計算 (1.5h)
> 7. 編寫單元測試 (2h)
> 風險：併發修改可能導致超賣
```

---

#### 5. `/check` - 完整質量檢查

**使用場景**：
- 提交代碼前的全面檢查
- Code Review 前的自檢
- 定期的質量審計
- 與 `/guardian` 配合使用

**具體用法**：
```bash
/check
```

**預期輸出**：
- 代碼風格問題
- 潛在 bug
- 性能優化建議
- 安全漏洞警告
- 質量評分報告

**實際案例**：
```bash
/check
> ✅ 代碼風格：符合 ESLint 規範
> ⚠️ 性能：發現 N+1 查詢 (user.service.ts:45)
> ❌ 安全：SQL 注入風險 (db.query.ts:23)
> 💡 建議：使用參數化查詢
```

---

#### 6. `/watch` - 監察模式（協作式質量守護）

**使用場景**：
- 編碼過程中的持續關注
- 需要主動提交代碼觸發檢查
- 建立良好的檢查習慣
- 與 `/check` 形成完整質量保證

**具體用法**：
```bash
/watch on      # 開啟監察模式
/watch off     # 關閉監察模式
/watch status  # 查看當前狀態
/watch report  # 生成監察報告
```

**預期輸出**：
- 即時安全警告
- 代碼質量提醒
- 性能風險提示
- 最佳實踐建議

**實際案例**：
```bash
/watch on
> 監察模式已開啟，請在編碼時定期提交代碼片段

"我剛寫了用戶驗證函數：[貼上代碼]"
> 🔴 發現安全問題：密碼未加密存儲
> 💡 建議使用 bcrypt 加密

"已修復，現在使用 bcrypt：[貼上修復後代碼]"
> ✅ 安全問題已解決

/watch status
> 監察時長：45分鐘
> 發現問題：3個（已修復2個）
```

**與 /check 的配合**：
```bash
/watch on       # 開始編碼前開啟監察
[編碼過程中持續提交和獲得反饋...]
/check          # 完成後執行完整檢查
/watch off      # 結束工作關閉監察
```

**重要說明**：
ℹ️ `/watch` 並非真正的實時監控，需要您主動提交代碼片段來觸發檢查。這是一個協作式的工作模式，幫助您在編碼過程中保持質量意識。

---

#### 7. `/test` - 測試生成與執行

**使用場景**：
- 為新功能編寫測試
- 補充測試覆蓋率
- 驗證 bug 修復

**具體用法**：
```bash
/test [文件|功能]
```

**預期輸出**：
- 生成的測試代碼
- 測試執行結果
- 覆蓋率報告
- 邊界條件測試

**實際案例**：
```bash
/test "UserService"
> 生成 5 個測試用例：
> ✅ should create user successfully
> ✅ should validate email format
> ✅ should hash password
> ❌ should handle duplicate email
> ✅ should update last login time
> 覆蓋率：87%
```

---

#### 8. `/debug` - 智能調試助手

**使用場景**：
- 遇到錯誤需要定位
- 性能問題排查
- 異常行為分析

**具體用法**：
```bash
/debug "錯誤信息或問題描述"
```

**預期輸出**：
- 問題原因分析
- 可能的解決方案
- 調試步驟建議
- 相關代碼定位

**實際案例**：
```bash
/debug "TypeError: Cannot read property 'id' of undefined"
> 問題定位：components/UserCard.tsx:15
> 原因：未檢查 user 對象是否存在
> 解決方案：
> 1. 添加可選鏈：user?.id
> 2. 或添加條件檢查：if (user && user.id)
> 3. 設置默認值：user = user || {}
```

---

#### 9. `/learn` - 學習並記錄決策

**使用場景**：
- 重要技術選型後
- 解決關鍵問題後
- 發現最佳實踐時

**具體用法**：
```bash
/learn "決策內容或經驗"
```

**預期輸出**：
- 更新 DECISIONS.md
- 智能分類標籤
- 關聯影響分析

**實際案例**：
```bash
/learn "選擇 PostgreSQL 而非 MongoDB，因為需要強一致性事務"
> ✅ 已記錄到 DECISIONS.md
> 標籤：#數據庫 #架構決策
> 影響：需要設計關係型表結構
> 相關文件：database/schema.sql
```

---

#### 10. `/doc` - 智能文檔維護

**使用場景**：
- 更新 API 文檔
- 維護 README
- 生成變更日誌

**具體用法**：
```bash
/doc [api|readme|changelog|arch]
```

**預期輸出**：
- 自動更新指定文檔
- 保留手動編輯內容
- 生成缺失部分

**實際案例**：
```bash
/doc api
> 掃描到 15 個 API 端點
> 更新 3 個新端點文檔
> 修正 2 個過時的參數說明
> 生成 Postman 集合
> 文檔位置：docs/api/README.md
```

---

#### 11. `/review` - PR 準備助手

**使用場景**：
- 創建 Pull Request 前
- 需要自我審查時
- 準備 code review

**具體用法**：
```bash
/review
```

**預期輸出**：
- 變更摘要
- PR 描述模板
- 檢查清單
- 潛在問題提醒

**實際案例**：
```bash
/review
> 變更摘要：
> - 新增：購物車功能模組
> - 修改：訂單服務接口
> - 刪除：廢棄的支付方法
> 
> PR 描述：
> ## 功能
> 實現完整購物車功能
> 
> ## 測試
> - [x] 單元測試通過
> - [x] 集成測試通過
> 
> ⚠️ 提醒：記得更新 API 文檔
```

---

#### 12. `/meta` - 項目規範定制

**使用場景**：
- 新項目初始化
- 團隊規範制定
- 技術棧變更時

**具體用法**：
```bash
/meta
```

**預期輸出**：
- 生成 CLAUDE.md
- 項目特定規範
- 工作流程定義

**實際案例**：
```bash
/meta
> 分析項目特徵...
> 識別為：Node.js 微服務
> 生成規範：
> - API 設計原則
> - 錯誤處理規範
> - 日誌記錄標準
> - 測試要求
> 已創建 CLAUDE.md
```

---

#### 13. `/analyze` - 深度分析與驗證

**使用場景**：
- 功能完成但有疑慮時
- 需要風險評估時
- 性能瓶頸分析

**具體用法**：
```bash
/analyze "功能/模組" ["具體疑慮"或"deep"]
```

**預期輸出**：
- 風險等級評估
- 邊界條件分析
- 改進建議
- 測試場景

**實際案例**：
```bash
/analyze "支付系統" "併發安全"
> 🚨 高風險：
> - 重複支付可能（無冪等性保證）
> - 金額計算存在浮點誤差
> 
> ⚠️ 中風險：
> - 超時處理不完善
> 
> 建議：
> 1. 添加分佈式鎖
> 2. 使用 BigDecimal
> 3. 實現冪等性檢查
```

---

#### 14. `/update-spec` - CLAUDE.md 更新專用

**使用場景**：
- 固化重要決策為規範
- 更新項目規則
- 週期性規範審查

**具體用法**：
```bash
/update-spec [review|section "content"]
```

**預期輸出**：
- 規範更新建議
- 版本變更記錄
- 衝突檢測

**實際案例**：
```bash
/update-spec review
> 發現 3 個可固化的決策：
> 1. API 命名使用 kebab-case
> 2. 所有異步函數使用 async/await
> 3. 測試覆蓋率不低於 80%
> 
> 執行更新：
/update-spec rules "API 使用 kebab-case 命名"
> ✅ 已更新開發規範章節
```

---

#### 📦 SDK 專用命令詳解

#### 1. `/sdk-design` - API 設計助手

**使用場景**：
- 設計新的 SDK 接口
- 重構現有 API
- 制定設計規範

**具體用法**：
```bash
/sdk-design "功能描述"
```

**預期輸出**：
- API 結構建議
- 命名規範
- 參數設計
- 錯誤處理策略

**實際案例**：
```bash
/sdk-design "文件上傳 SDK"
> API 設計建議：
> 
> 核心接口：
> - upload(file, options)
> - uploadMultiple(files, options)
> - resumeUpload(uploadId)
> 
> 配置選項：
> - chunkSize: 分片大小
> - retryTimes: 重試次數
> - onProgress: 進度回調
> 
> 錯誤碼設計：
> - FILE_TOO_LARGE
> - NETWORK_ERROR
> - INVALID_FORMAT
```

---

#### 2. `/sdk-example` - 示例代碼生成

**使用場景**：
- 為 SDK 創建使用示例
- 編寫快速開始指南
- 展示最佳實踐

**具體用法**：
```bash
/sdk-example [basic|advanced|integration|all]
```

**預期輸出**：
- 基礎使用示例
- 高級特性展示
- 集成場景代碼
- 完整示例套件

**實際案例**：
```bash
/sdk-example basic
> 生成基礎示例：
> 
> // 初始化
> const sdk = new MySDK({
>   apiKey: 'your-api-key'
> });
> 
> // 基本使用
> const result = await sdk.doSomething();
> 
> // 錯誤處理
> try {
>   await sdk.riskyOperation();
> } catch (error) {
>   console.error(error.code);
> }
```

---

#### 3. `/sdk-test` - SDK 測試套件

**使用場景**：
- 生成完整測試套件
- 兼容性測試
- 性能基準測試

**具體用法**：
```bash
/sdk-test [unit|integration|compat|performance|all]
```

**預期輸出**：
- 單元測試代碼
- 集成測試場景
- 多版本兼容測試
- 性能測試基準

**實際案例**：
```bash
/sdk-test compat
> 生成兼容性測試：
> 
> ✅ Node.js 14.x 兼容性
> ✅ Node.js 16.x 兼容性
> ✅ Node.js 18.x 兼容性
> ✅ TypeScript 4.x 支援
> ✅ CommonJS 導入
> ✅ ES Module 導入
> 
> 測試文件：tests/compatibility/
```

---

#### 4. `/sdk-doc` - SDK 文檔生成

**使用場景**：
- 生成 API 參考文檔
- 編寫使用指南
- 創建遷移文檔

**具體用法**：
```bash
/sdk-doc [api|guide|migration|all]
```

**預期輸出**：
- API 參考手冊
- 入門指南
- 遷移指南
- 完整文檔套件

**實際案例**：
```bash
/sdk-doc api
> 生成 API 文檔：
> 
> ## Class: MySDK
> ### Constructor
> - new MySDK(config)
> 
> ### Methods
> - async connect()
> - async disconnect()
> - async query(params)
> 
> ### Events
> - 'connected'
> - 'error'
> - 'data'
> 
> 文檔位置：docs/api-reference.md
```

---

#### 5. `/sdk-release` - 發布準備助手

**使用場景**：
- 準備新版本發布
- 檢查發布清單
- 生成發布說明

**具體用法**：
```bash
/sdk-release [major|minor|patch|check]
```

**預期輸出**：
- 版本號建議
- 破壞性變更檢查
- 發布清單
- 變更日誌

**實際案例**：
```bash
/sdk-release check
> 發布檢查清單：
> ✅ 版本號更新
> ✅ CHANGELOG 更新
> ✅ 文檔同步
> ❌ 破壞性變更未標記
> ✅ 測試全部通過
> ⚠️ 建議：添加遷移指南
> 
> 建議版本：2.1.0 (minor)
> 原因：新增功能，向後兼容
```

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

#### 5. 深度驗證（新增）
```bash
/analyze "核心功能" deep       # 全面風險分析
/analyze "支付模組" "併發安全"  # 特定場景驗證
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
| 驗證 | 憑經驗判斷 | `/analyze` 深度分析 |
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

# 5. 深度驗證（新增）
/analyze "JWT認證" "token洩露和重放攻擊"
> 風險評估：發現 3 個潛在問題
> 提供加固方案

# 6. 記錄決策
/learn "選擇 JWT 而非 session 因為需要支援分佈式"
```

### 🎯 命令組合使用場景

#### 場景 1：從零開始新項目
```bash
/meta                      # 建立項目規範
/start                     # 理解項目結構
/plan "MVP 功能列表"       # 規劃開發路線
/learn "技術選型決策"      # 記錄重要決定
```

#### 場景 2：接手遺留項目
```bash
/start                     # 快速理解項目
/analyze "整體架構" deep   # 深度分析潛在問題
/plan "重構計劃"          # 制定改進方案
/update-spec review        # 建立規範基線
```

#### 場景 3：日常功能開發
```bash
/sync                      # 恢復工作狀態
/plan "新功能"            # 分解任務
# ... 編碼 ...
/test                      # 生成測試
/check                     # 質量檢查
/analyze "新功能" deep    # 風險評估
/doc api                   # 更新文檔
/review                    # 準備 PR
```

#### 場景 4：緊急 Bug 修復
```bash
/debug "錯誤描述"         # 快速定位
/analyze "相關模組" "影響範圍"  # 評估影響
# ... 修復 ...
/test "修復驗證"          # 驗證修復
/learn "bug 原因和解決方案"  # 記錄經驗
```

#### 場景 5：性能優化
```bash
/analyze "瓶頸模組" "性能"  # 分析瓶頸
/plan "優化方案"          # 制定計劃
# ... 優化 ...
/test performance         # 性能測試
/check                    # 確保不破壞功能
/learn "優化技巧"        # 記錄經驗
```

#### 場景 6：SDK 完整開發流程
```bash
/meta                           # SDK 規範
/sdk-design "核心功能"          # API 設計
/plan "實現計劃"               # 任務分解
# ... 開發 ...
/sdk-test all                  # 完整測試
/sdk-example all               # 生成示例
/sdk-doc all                   # 生成文檔
/sdk-release check             # 發布檢查
```

### 🚀 高級技巧

#### 1. 命令鏈式使用
```bash
# 完整的質量保證鏈
/check && /test && /analyze "核心模組" deep

# 文檔更新鏈
/doc api && /doc readme && /doc changelog
```

#### 2. 定期維護流程
```bash
# 每週執行
/update-spec review     # 審查可固化的決策
/doc changelog         # 更新變更日誌
/analyze "核心系統" deep  # 深度健康檢查

# 每個迭代執行
/review               # PR 準備
/sdk-release check    # 版本檢查
```

#### 3. 知識管理最佳實踐
```bash
# 立即記錄
/learn "任何重要決定"    # 不要等待，立即記錄

# 定期固化
/update-spec review      # 將決策轉為規範
```

#### 4. 風險預防策略
```bash
# 預防性分析
/analyze "新功能" deep   # 開發前分析
/analyze "重構目標" "影響"  # 重構前評估

# 持續監控
/check                  # 每次提交前
/test                   # 每個功能完成後
```

### 📈 效率提升對比

| 傳統開發 | 使用命令系統 | 效率提升 |
|---------|------------|---------|
| 手動分析項目 30min | `/start` 2min | 15x |
| 編寫測試 2h | `/test` 10min | 12x |
| 代碼審查 1h | `/check` 5min | 12x |
| 文檔更新 2h | `/doc` 10min | 12x |
| 問題定位 1h | `/debug` 5min | 12x |
| 風險評估 2h | `/analyze` 15min | 8x |
| PR 準備 30min | `/review` 3min | 10x |

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
- **深度驗證**：用 `/analyze` 驗證直覺
- **規範管理**：用 `/update-spec` 固化重要決策

### 4. 團隊協作
- 共享 `.claude/` 目錄
- 統一使用命令系統
- 定期更新 PROJECT_CONTEXT.md
- 重要決策都記錄在 DECISIONS.md

---

## 📊 版本歷史

### v3.3.0 (2025-08-10) - 當前版本
- **重大調整**：
  - 移除所有 project 層級命令，專注於全局通用命令
  - 讓開發者自行建立項目專屬命令
  - 統一使用 13 個 global 通用命令 + 5 個 SDK 專用命令

### v3.2.2 (2025-08-10)
- **架構修正**：
  - 將 `/analyze` 和 `/update-spec` 正確歸類為 global 命令
  - 通用命令數量更正為 13 個
  - 清理重複的命令文件
  - 修正項目結構說明

### v3.2.1 (2025-08-10)
- **文檔大幅增強**：
  - 新增完整的命令使用手冊（18個命令詳解）
  - 每個命令包含：使用場景、具體用法、預期輸出、實際案例
  - 添加 6 個典型命令組合場景
  - 提供高級使用技巧和效率對比
- **用戶體驗優化**：
  - 命令說明更加清晰直觀
  - 提供豐富的實際使用案例
  - 新增命令鏈式使用指南

### v3.2.0 (2025-08-10)
- **新增功能**：深度分析驗證命令
  - 創建 `/analyze` 命令，填補"功能完成"到"生產就緒"的驗證空白
  - 支援基於經驗直覺的風險分析
  - 提供量化風險評估和優先級建議
- **質量保證增強**：
  - 邊界條件自動分析
  - 特殊場景推演
  - 架構層面審視
  - 業務邏輯驗證

### v3.1.0 (2024-01-20)
- **新增功能**：專用的 CLAUDE.md 更新命令
  - 創建 `/update-spec` 命令，專門負責更新項目規範
  - 支援兩種模式：審查模式（review）和定向更新模式
  - 實現命令職責單一化設計
- **架構優化**：
  - 明確各命令的職責邊界
  - `/learn` 只更新 DECISIONS.md 和 PROJECT_CONTEXT.md
  - `/update-spec` 只更新 CLAUDE.md
  - 建立清晰的命令職責矩陣
- **工作流改進**：
  - 決策記錄與規範固化分離
  - 支援週期性規範審查和更新
  - 提供智能建議和衝突檢測

### v3.0.0 (2024-01-15)
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