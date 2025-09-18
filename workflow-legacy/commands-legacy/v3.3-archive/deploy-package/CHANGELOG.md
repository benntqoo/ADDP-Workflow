# 變更日誌

## v3.0.0 (2024-01-15) - 重大重構

### 🚀 新功能
- 精簡命令系統從 31 個減少到 11 個核心命令
- 智能命令整合，一個命令完成多項相關任務
- 自動化記憶管理系統
- 結構化項目上下文（PROJECT_CONTEXT.md）
- 決策記錄系統（DECISIONS.md）

### 💡 核心命令
1. `/start` - 項目快速啟動與理解
2. `/context` - 上下文同步檢查點
3. `/sync` - 狀態同步器
4. `/plan` - 任務規劃與設計
5. `/check` - 智能代碼審查
6. `/test` - 測試生成與執行
7. `/learn` - 學習並記錄決策
8. `/doc` - 智能文檔維護
9. `/review` - PR 準備助手
10. `/debug` - 智能調試助手
11. `/meta` - 項目規範定制

### ⚡ 改進
- 命令職責更明確，無功能重疊
- 大部分命令無需參數，智能推斷
- 更好的上下文保持能力
- 簡化的學習曲線
- 提升人機協作效率

### 🔄 整合的命令
- `analyze`, `discover`, `explore` → `/start`
- `audit`, `coverage`, `perf` → `/check`
- `doc-api`, `doc-arch`, `changelog`, `readme` → `/doc`
- `deploy-check`, `rollback`, `config` → `/review`
- 其他專業命令整合到相關核心命令

### 📝 文檔更新
- 新增 SIMPLE_COMMANDS.md 詳細說明
- 更新 COMMANDS_SUMMARY.md 為 v3.0
- 重寫 DEPLOY_GUIDE.md 適配新系統
- 創建 PROJECT_CONTEXT.md 模板

---

## v2.1.0 (2024-01-14)

### 新增功能
- 參數規範化：所有命令都有明確的 format 和 examples
- 命令協調機制：支持狀態共享和智能推薦
- 項目命令增強：從簡單功能升級為完整系統

### 優化
- 解決了功能重疊問題
- 明確了命令職責邊界
- 優化了命令參數定義

---

## v2.0.0 (2024-01-13)

### 初始版本
- 31 個命令覆蓋完整開發生命週期
- 包含全局命令和項目命令
- 自動部署腳本支持 Windows/macOS/Linux
