# Output Styles System

## ⚠️ 重要聲明 (IMPORTANT NOTICE)

**Output Styles 定義專業角色和輸出類型，每個角色有明確的職責邊界！**

核心原則：
- ✅ 每個Style專注其專業領域
- ✅ 明確的輸出類型（文檔/代碼/報告）
- ✅ 清晰的角色協作流程
- ✅ 修改優先於創建新文件

## 📊 Style角色矩陣

| Style | 主要輸出 | 實際行為 | 適用場景 |
|-------|---------|---------|---------|
| **architect** | 📄 架構文檔、設計決策 | 分析→設計→文檔 | 系統設計、技術選型 |
| **concise-developer** | 💻 生產級代碼 | 讀碼→改碼→測試 | 功能實現、bug修復 |
| **devops-engineer** | 🔧 腳本、配置 | 分析→直接實施 | CI/CD、部署、監控 |
| **security-analyst** | 📊 安全報告 | 審查→分析→建議 | 安全審計、威脅評估 |
| **product-expert** | 📝 PRD文檔 | 需求→規格→文檔 | 產品規劃、需求定義 |
| **educational-mentor** | 📚 教學內容 | 解釋→示例→引導 | 學習新技術、理解概念 |
| **sdk-design-expert** | 📐 API設計 | 設計→規範→文檔 | API設計、接口定義 |
| **sdk-prd-expert** | 📖 SDK文檔 | 規劃→文檔→示例 | SDK需求、開發者文檔 |

## 🎯 使用指南

### 三類使用場景

#### 1. 探索/設計類（不確定時）
```bash
# 產品規劃
/output-style:set product-expert

# 技術架構
/output-style:set architect

# 學習理解
/output-style:set educational-mentor
```

#### 2. 執行/開發類（確定要做什麼時）
```bash
# 寫代碼
/output-style:set concise-developer

# 部署配置
/output-style:set devops-engineer

# API設計
/output-style:set sdk-design-expert
```

#### 3. 審查/分析類（需要檢查時）
```bash
# 安全審查
/output-style:set security-analyst
```

## 🔄 標準協作流程

### 完整開發流程
```
1. product-expert      → 輸出PRD文檔
2. architect          → 輸出技術設計
3. concise-developer  → 實現代碼
4. security-analyst   → 安全審查報告
5. devops-engineer    → 部署和監控
```

### 問題修復流程
```
1. security-analyst   → 發現問題，輸出報告
2. architect         → 設計修復方案（如需要）
3. concise-developer → 基於報告實施修復
4. devops-engineer   → 部署修復
```

## 🎨 Style + Command 最佳組合

### 核心概念
- **Style定義角色性格**：如何思考、如何表達、輸出什麼
- **Command定義具體動作**：做什麼任務、執行什麼流程
- **組合產生協同效應**：不同組合適用不同場景

### 📋 推薦組合矩陣

| 場景 | Style | Command | 效果 |
|------|-------|---------|------|
| **項目啟動** | architect | `/start` → `/plan` | 理解項目並設計架構 |
| **日常開發** | concise-developer | `/sync` → `/plan` | 恢復狀態並規劃任務 |
| **功能設計** | architect | `/plan` → `/learn` | 設計方案並記錄決策 |
| **代碼實現** | concise-developer | `/context` → 編碼 | 確認理解後實現 |
| **學習新技術** | educational-mentor | `/start` → `/doc` | 理解並記錄知識 |
| **安全審計** | security-analyst | `/context` → 分析 | 理解系統後審查 |
| **部署準備** | devops-engineer | `/plan` → 實施 | 規劃並執行部署 |
| **項目交接** | architect | `/doc` → `/update-spec` | 更新文檔和規範 |

### 🚀 典型工作流範例

#### 1. 新功能開發（從零開始）
```bash
# 階段1：需求理解
/output-style:set product-expert
/start  # 理解項目背景
"描述用戶認證需求"

# 階段2：技術設計
/output-style:set architect
/plan "用戶認證系統"  # 架構設計
/learn "選擇JWT作為認證方案"  # 記錄決策

# 階段3：代碼實現
/output-style:set concise-developer
/context  # 確認技術方案理解
"實現JWT認證模塊"  # 開始編碼

# 階段4：安全檢查
/output-style:set security-analyst
"審查認證實現的安全性"  # 輸出報告

# 階段5：修復問題
/output-style:set concise-developer
"基於安全報告修復token過期處理"  # 修復

# 階段6：部署
/output-style:set devops-engineer
"配置CI/CD pipeline"  # 部署設置
```

#### 2. 日常維護開發
```bash
# 早上開始工作
/output-style:set concise-developer
/sync  # 恢復昨天的進度
/context  # 確認今天要做什麼
/plan "完成用戶管理模塊"  # 規劃今天的任務

# 開發過程中
"實現用戶CRUD操作"  # 編碼
/learn "使用Repository模式"  # 記錄重要決策

# 完成後
/update-spec  # 更新項目規範
/doc  # 更新相關文檔
```

#### 3. 緊急Bug修復
```bash
# 問題分析
/output-style:set security-analyst
/context  # 理解問題背景
"分析SQL注入漏洞"  # 輸出分析報告

# 快速修復
/output-style:set concise-developer
/sync  # 獲取最新代碼狀態
"基於報告修復SQL注入"  # 立即修復

# 部署修復
/output-style:set devops-engineer
"部署hotfix到生產環境"  # 緊急部署
```

#### 4. 學習和探索
```bash
# 學習新框架
/output-style:set educational-mentor
/start  # 了解基礎概念
"解釋React Server Components"

# 嘗試實現
/output-style:set concise-developer
/plan "創建RSC示例"  # 規劃學習項目
"實現一個簡單的RSC應用"

# 記錄學習
/learn "RSC適合SEO優化場景"  # 記錄心得
/doc  # 創建學習筆記
```

### 💡 組合使用技巧

#### Style切換時機
- **開始新任務**：根據任務類型選擇Style
- **角色轉換時**：從設計轉到實現時切換
- **需要不同視角**：想要不同專業意見時

#### Command使用時機
- **開始工作**：`/sync` 恢復狀態
- **任務規劃**：`/plan` 組織思路
- **重要決策**：`/learn` 記錄下來
- **完成任務**：`/update-spec` 更新規範

#### 高效組合原則
1. **Style穩定，Command頻繁**：在一個任務中保持Style，頻繁使用Commands
2. **先規劃後執行**：先用 `/plan`，再開始實際工作
3. **及時記錄**：用 `/learn` 記錄重要決策
4. **保持同步**：定期 `/sync` 和 `/context` 確保理解一致

## 💡 關鍵特性

### 1. 明確的輸出邊界
- **architect**: 只出設計文檔，不寫生產代碼
- **security-analyst**: 只出報告，不實施修復
- **product-expert**: 只寫需求，不寫代碼
- **concise-developer**: 負責所有代碼實現

### 2. 修改優先原則
所有開發相關的Styles都遵循：
- 先搜索現有代碼
- 優先修改而非創建
- 擴展而非重寫
- 創建前需確認

### 3. 專業分工
- 每個角色都有其專業深度
- 不試圖讓一個角色做所有事
- 通過切換Style實現完整工作流

## 📁 可用風格列表

### 設計與規劃
- **architect** - 系統架構設計（合併了senior和system架構師）
- **product-expert** - 產品需求規劃
- **sdk-design-expert** - API設計規範

### 開發與實施
- **concise-developer** - 高效代碼開發
- **devops-engineer** - 基礎設施自動化

### 分析與審查
- **security-analyst** - 安全威脅分析
- **educational-mentor** - 教學與解釋

### 文檔與規範
- **sdk-prd-expert** - SDK產品文檔

## 🚀 快速開始

### 查看可用風格
```bash
/output-style
```

### 切換風格
```bash
/output-style:set architect
/output-style:set concise-developer
```

### 查看當前風格
```bash
/output-style:current
```

## 📋 使用建議

### 在你擅長的領域
- 直接使用 `concise-developer`
- 給出精確指令
- 快速迭代

### 在你不熟悉的領域
1. 先用相關expert探索（architect/product-expert）
2. 理解後用developer實施
3. 最後用analyst審查

### 學習新技術時
- 使用 `educational-mentor` 理解概念
- 切換到 `concise-developer` 實踐
- 用 `security-analyst` 檢查最佳實踐

## ⚙️ 配置方法

### 項目級別配置
```json
// .claude/settings.local.json
{
  "outputStyle": "concise-developer"
}
```

### 全局配置
```json
// ~/.claude/settings.json
{
  "defaultOutputStyle": "architect"
}
```

## 🎓 最佳實踐

### DO ✅
- 為不同任務選擇合適的Style
- 遵循協作流程
- 利用每個角色的專業性
- 明確指令配合Style使用

### DON'T ❌
- 不要期望architect寫生產代碼
- 不要讓security-analyst修復代碼
- 不要用product-expert寫實現
- 不要在一個Style中做所有事

## 🔧 問題排查

### Q: Style沒有生效？
檢查：
1. 文件是否存在於 `~/.claude/output-styles/`
2. 使用 `/output-style:set <name>` 切換
3. 確認Style名稱正確

### Q: 不確定用哪個Style？
- 需要設計 → architect
- 需要編碼 → concise-developer
- 需要學習 → educational-mentor
- 需要審查 → security-analyst

### Q: 如何處理安全問題？
1. security-analyst 分析和報告
2. architect 設計修復方案（如需要）
3. concise-developer 實施修復

## 🌍 國際化支持

所有Style文件都支持多語言描述：
```yaml
---
description:
  en: English description
  zh: 中文描述
---
```

## 📜 版本歷史

- **v2.0** - 重構為專業分工系統
  - 合併重複的架構師角色
  - 明確每個角色的輸出類型
  - 加入修改優先原則
  - 優化協作流程

- **v1.0** - 初始版本
  - 9個專業Style
  - 基礎功能實現

---

*Output Styles系統持續優化中，歡迎反饋使用體驗！*