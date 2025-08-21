# 變更日誌

## v3.1.0 (2025-01-21) - 狀態管理完善

### 🚀 新功能
- `/plan` 命令現在會創建/更新 `last-session.yml`，標記開發週期開始
- `/context` 命令新增項目進展對比功能，使用 `initial-scan.json` 作為基線
- `/update-spec` 命令升級為週期終結儀式，完整保存所有狀態
- `/sync` 命令增強錯誤處理，優雅處理文件缺失情況

### ⚡ 改進
- 修復了命令體系中的斷鏈問題（last-session.yml 更新鏈）
- 激活了孤立文件（initial-scan.json 現在被 /context 使用）
- 完善了工作流程：sync → context/plan → 開發 → update-spec → commit
- 所有狀態文件現在能正確流轉

### 📝 文檔更新
- 更新 SIMPLE_COMMANDS_SUMMARY.md 反映新的狀態管理功能
- 各命令文檔添加了狀態文件交互說明
- 創建 COMMAND_SYSTEM_AUDIT.md 記錄體系分析

### 🐛 修復
- 解決了 last-session.yml 只創建不更新的問題
- 修復了 initial-scan.json 完全未使用的問題
- 改進了文件缺失時的錯誤處理邏輯

---

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
