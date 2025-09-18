# 🚀 部署指南 - Memory System 統一升級

## 📋 升級概述

此升級將記憶文件統一到 `.claude/memory/` 目錄，並移除了無效的 memory-system agent 引用。

### 🎯 主要變更
- ✅ 統一記憶目錄：`.claude/memory/`
- ✅ 智能文件遷移：自動搜索並遷移舊文件
- ✅ 移除無效引用：刪除 memory-system agent
- ✅ 向後兼容：支持舊版本文件位置

## 📁 新的目錄結構

```
.claude/
├── memory/                      # 🆕 統一記憶目錄
│   ├── PROJECT_CONTEXT.md       # 項目上下文
│   ├── DECISIONS.md             # 技術決策
│   └── last-session.yml         # 會話狀態
├── CLAUDE.md                    # 用戶偏好（位置不變）
└── settings.local.json          # 本地設置（位置不變）
```

## 🔄 部署步驟

### 步驟 1：更新命令文件
**已完成** - 所有命令文件已更新：
- `/sync` - 加入智能搜索和遷移邏輯
- `/start` - 使用新的 memory 目錄
- `/learn` - 直接操作新目錄
- `/plan` - 更新至新路徑
- `/context` - 智能搜索文件
- `/update-spec` - 批量更新新目錄

### 步驟 2：部署到本地環境

```bash
# 1. 進入項目目錄
cd D:\Code\ai\claude

# 2. 部署更新的命令
# Windows
xcopy /E /Y commands\deploy-package\global\* %USERPROFILE%\.claude\commands\

# macOS/Linux  
cp -r commands/deploy-package/global/* ~/.claude/commands/

# 3. 重新啟動 Claude Code
# 關閉並重新打開 Claude Code CLI
```

### 步驟 3：測試升級

```bash
# 測試 /sync 命令的遷移功能
/sync

# 應該看到：
# 📊 文件遷移報告：
# ✅ PROJECT_CONTEXT.md: 從 .claude/ 遷移到 .claude/memory/
# ✅ DECISIONS.md: 從 .claude/ 遷移到 .claude/memory/  
# ✅ last-session.yml: 從 .claude/ 遷移到 .claude/memory/
```

## 🔍 智能遷移機制

### 文件搜索優先級
1. **新位置**：`.claude/memory/{file}` 
2. **舊位置1**：`.claude/{file}`
3. **舊位置2**：`./{file}` 
4. **舊位置3**：`docs/{file}`

### 遷移邏輯
```yaml
搜索策略:
  - 如果新位置存在 → 使用它
  - 如果舊位置存在 → 自動遷移到新位置
  - 如果都不存在 → 在新位置創建

衝突處理:
  - 多處存在文件 → 選擇最新版本
  - 遷移成功 → 刪除舊位置文件  
  - 保持數據完整性
```

## ⚠️ 注意事項

### 升級前檢查
- ✅ 確保沒有未保存的工作
- ✅ 檢查 `.claude/` 目錄下的重要文件
- ✅ 備份重要配置（可選）

### 升級後驗證
```bash
# 檢查新目錄
ls -la .claude/memory/

# 測試命令
/sync
/context
/learn "測試升級成功"
```

### 回滾方案
如果需要回滾到舊版本：
```bash
# 1. 恢復舊命令文件（如果有備份）
# 2. 手動移動文件回舊位置
mv .claude/memory/*.md .claude/
mv .claude/memory/*.yml .claude/
```

## 🎉 升級完成驗證

### 功能測試清單
- [ ] `/sync` 能正常讀取項目狀態
- [ ] `/start` 能在新目錄創建文件  
- [ ] `/learn` 能正常記錄決策
- [ ] `/plan` 能保存會話狀態
- [ ] `/context` 能讀取完整上下文
- [ ] `/update-spec` 能批量更新文件

### 預期效果
- 🏃‍♂️ **更快的文件操作**：統一目錄減少搜索時間
- 🔒 **更好的組織性**：清晰的文件結構
- ⚡ **向後兼容**：自動處理舊版本文件
- 🐛 **移除錯誤**：不再嘗試調用不存在的 memory-system agent

## 📞 問題排查

### 常見問題

**Q1: /sync 提示找不到文件**
```bash
A: 正常情況，首次運行會創建新文件
   檢查 .claude/memory/ 目錄是否創建成功
```

**Q2: 舊文件沒有自動遷移**  
```bash
A: 手動檢查文件位置：
   ls -la .claude/PROJECT_CONTEXT.md
   如果存在，再次運行 /sync
```

**Q3: 命令執行報錯**
```bash
A: 確認命令文件部署成功：
   ls ~/.claude/commands/sync.md
   檢查文件是否為最新版本
```

## 🚀 下一步

升級完成後：
1. 正常使用所有命令
2. 享受更統一的文件管理體驗  
3. 舊位置的文件會自動清理
4. 新項目會直接使用新的目錄結構

---

**升級完成！** 🎊 現在你可以享受更簡潔、統一的記憶文件管理體驗了！