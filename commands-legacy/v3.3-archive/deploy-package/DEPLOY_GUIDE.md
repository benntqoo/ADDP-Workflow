# Claude Code 精簡命令系統部署指南 v3.0

## 🚀 快速部署

### Windows
```powershell
cd claude\commands\deploy-package
.\deploy.ps1
```

### macOS/Linux
```bash
cd claude/commands/deploy-package
chmod +x deploy.sh
./deploy.sh
```

## 📦 包含內容

### 通用命令（13個）
1. **start** - 項目快速啟動與理解
2. **context** - 上下文同步檢查點
3. **sync** - 狀態同步器
4. **plan** - 任務規劃與設計
5. **check** - 智能代碼審查
6. **test** - 測試生成與執行
7. **learn** - 學習並記錄決策
8. **doc** - 智能文檔維護
9. **review** - PR 準備助手
10. **debug** - 智能調試助手
11. **meta** - 項目規範定制
12. **analyze** - 深度風險分析
13. **update-spec** - 更新項目規範

### 部署位置
- **Windows**: `%USERPROFILE%\.claude\commands\`
- **macOS/Linux**: `~/.claude/commands/`

## 🔧 手動部署

如果自動腳本失敗，可手動複製：

```bash
# 1. 創建目標目錄
mkdir -p ~/.claude/commands

# 2. 複製全局命令
cp global/*.md ~/.claude/commands/

# 3. 創建項目專屬命令（開發者自定義）
# 開發者可以在 .claude/commands/ 中創建自己的項目命令
mkdir -p YOUR_PROJECT/.claude/commands
# 參考 global 命令格式創建專屬命令
```

## 📋 部署後驗證

1. 在 Claude Code 中輸入 `/` 查看可用命令
2. 測試核心命令：
   ```bash
   /start    # 應該開始分析項目
   /context  # 應該顯示當前理解
   ```

## ⚙️ 自定義配置

### 選擇性部署
如果只需要部分命令，可以只複製需要的 `.md` 文件。

### 項目特定命令
在項目的 `.claude/commands/` 目錄創建自定義命令。

### 命令優先級
1. 項目命令（`.claude/commands/`）
2. 全局命令（`~/.claude/commands/`）
3. 內建命令

## 🔄 更新說明

從舊版本（v2.x）升級：
1. 備份現有自定義命令
2. 刪除舊的全局命令：
   ```bash
   rm ~/.claude/commands/*.md
   ```
3. 運行新的部署腳本
4. 恢復自定義命令（如有）

## 💡 使用建議

### 新項目啟動
```bash
/meta      # 建立項目規範
/start     # 理解項目結構
/plan      # 規劃第一個功能
```

### 日常開發流程
```bash
/sync      # 恢復工作狀態
/context   # 確認理解正確
/plan      # 規劃新任務
/check     # 代碼質量檢查
/test      # 執行測試
/learn     # 記錄重要決策
```

### 提交代碼
```bash
/check     # 最終檢查
/doc       # 更新文檔
/review    # 準備 PR
```

## ❓ 常見問題

### Q: 命令不顯示？
A: 
- 確保文件擴展名是 `.md`
- 檢查文件是否在正確目錄
- 重啟 Claude Code

### Q: 命令衝突？
A: 項目命令會覆蓋同名的全局命令。

### Q: 如何完全卸載？
A: 刪除 `~/.claude/commands/` 中的相應文件。

### Q: 與舊命令的區別？
A: v3.0 將 31 個命令整合為 13 個通用命令，功能更智能，使用更簡單。

## 📝 版本說明

### v3.3 (2025-08-10)
- **精簡**：從 31 個命令減少到 13 個通用命令
- **專注**：移除 project 層級命令，讓開發者自行定義
- **智能**：每個命令整合多項相關功能
- **自動**：記憶管理和狀態同步自動化
- **簡化**：大部分命令無需參數

### 主要改進
- 命令數量減少 65%
- 功能覆蓋率保持 100%
- 學習成本大幅降低
- 人機協作效率提升

### 移除的命令
以下功能已整合到新命令中：
- analyze, discover, explore → `/start`
- audit, coverage, perf → `/check`
- doc-api, doc-arch, changelog, readme → `/doc`
- deploy-check, rollback, config → `/review`
- 其他專業命令整合到相關核心命令

---

*簡單高效，讓 Claude Code 成為你的最佳開發夥伴！*

*部署包版本：3.0.0*  
*發布日期：2024-01-15*