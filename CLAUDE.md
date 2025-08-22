# CLAUDE.md - 項目協作規範

## 項目概述
Claude Code 優化項目，提供工作流命令、Agent和Output Styles系統，為開發者提供高效的AI協作體驗。

## Output Styles 系統

### 使用內建風格
項目提供了5種專業的Output Styles，可通過以下方式使用：

```bash
# 查看可用風格
/output-style

# 切換到架構師風格進行系統設計
/output-style:set senior-architect
/plan 設計微服務架構

# 切換到簡潔開發風格進行編碼
/output-style:set concise-developer

# 進行安全審查
/output-style:set security-analyst
/review
```

### 內建風格說明
- **senior-architect**: 戰略系統設計和架構討論
- **concise-developer**: 簡潔高效的編碼風格
- **educational-mentor**: 詳細解釋的教學風格
- **devops-engineer**: 基礎設施自動化專注
- **security-analyst**: 安全分析和威脅建模

### 自定義風格
創建文件 `~/.claude/output-styles/my-style.md`：
```markdown
---
description: 我的自定義風格
---

# My Custom Style

[風格定義內容]
```

## 開發工作流

### 1. 開始新週期
```bash
/sync                    # 恢復上次狀態
/plan "feature name"     # 開始新功能開發
```

### 2. 開發過程
```bash
/context                 # 確認上下文理解
# 進行開發工作...
/learn "重要決策"        # 記錄關鍵決策
```

### 3. 結束週期
```bash
/update-spec            # 更新規範和文檔
git commit              # 提交變更
```

## 測試規範
- 始終運行lint和typecheck
- 測試失敗前不提交代碼
- 使用項目定義的測試命令

## 代碼規範
- 遵循現有代碼風格
- 不添加不必要的註釋
- 優先修改現有文件而非創建新文件
- 永不主動創建文檔除非明確要求