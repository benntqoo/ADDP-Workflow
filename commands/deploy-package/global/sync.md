---
arguments: none
---

# 狀態同步器 - 恢復工作上下文

新會話開始時，我將幫你快速恢復到上次的工作狀態：

## 🔄 同步內容

<!-- File Operations: Direct Read/Write -->
<!-- Memory Directory: .claude/memory/ -->

### 1. 智能文件搜索與遷移

我將執行智能搜索和遷移，確保記憶文件在正確位置：

```yaml
# 記憶文件管理策略
target_directory: .claude/memory/
files_to_manage:
  - PROJECT_CONTEXT.md    # 項目上下文
  - DECISIONS.md          # 技術決策
  - last-session.yml      # 會話狀態

search_locations:
  priority_1: .claude/memory/    # 新標準位置
  priority_2: .claude/           # 舊版位置
  priority_3: ./                 # 根目錄
  priority_4: docs/              # 文檔目錄

migration_strategy:
  - if_found_in_old: 遷移到新位置
  - if_not_found: 在新位置創建
  - if_multiple: 選擇最新版本
```

**執行步驟**：

1. **檢查/創建目標目錄**
   ```bash
   mkdir -p .claude/memory
   ```

2. **智能文件搜索**（對每個文件）：
   - 檢查 `.claude/memory/{file}` 是否存在
   - 如不存在，搜索舊位置：
     - `.claude/{file}`
     - `./{file}`
     - `docs/{file}`
   
3. **文件遷移**（如需要）：
   - 找到舊文件 → 複製到新位置
   - 驗證遷移成功 → 刪除舊文件
   - 生成遷移報告

4. **缺失文件處理**：
   - PROJECT_CONTEXT.md → 創建基礎模板
   - DECISIONS.md → 創建空框架
   - last-session.yml → 從 git 歷史推斷

5. **語言偏好檢查**：
   - 讀取 `.claude/CLAUDE.md` 的 `preferred_language`
   - 如未設置，詢問用戶偏好

### 2. 檢查項目狀態
```bash
# 自動執行
- git status          # 未提交的更改
- git log -5         # 最近的提交
- 查看最近修改的文件
- 檢查測試狀態
```

### 3. 智能分析工作重點
**從多個數據源綜合分析**：
- 從 git log 提取最近工作內容
- 從 TODO 註釋找出待辦事項
- 從 DECISIONS.md 識別重要決策
- 從最近修改文件推斷當前任務

**智能建議**：
- 根據未完成的工作推薦下一步
- 識別潛在問題和風險
- 提供具體的命令建議

### 4. 檢查語言偏好
**自動語言設置**：
- 從 `.claude/CLAUDE.md` 讀取 `preferred_language` 設置
- 如果文件不存在或未設置，詢問用戶：「您希望我使用什麼語言回應？(zh-TW/zh-CN/en/ja/ko/其他)」
- 記錄用戶選擇到 `.claude/CLAUDE.md` 的語言偏好部分
- 根據設置調整回應語言（技術術語保持英文）
- 智能適應：根據用戶輸入語言動態調整

### 5. 生成狀態報告

```markdown
📊 工作狀態同步報告

## 🌐 語言設置
- 當前語言：[從 .claude/CLAUDE.md 的 preferred_language 讀取]
- 自動適應：根據用戶輸入動態調整

## 🔄 會話恢復
- 上次工作：[從 last-session.yml 讀取]
- 週期狀態：[planning/development/reviewing]
- 中斷時間：[計算時間差]

## 🎯 當前焦點
- 正在開發：[功能名稱]
- 上次進度：[完成情況]
- 待解決：[問題列表]

## 📝 最近變更
- [文件1]：[變更說明]
- [文件2]：[變更說明]

## ⏳ 待辦事項
1. [高優先級任務]
2. [中優先級任務]
3. [低優先級任務]

## 💡 建議行動
基於當前狀態，建議你：
- [具體建議]

## ⚠️ 注意事項
- [需要注意的風險或問題]
```

### 5. 智能提醒
- 是否有未運行的測試
- 是否有未提交的代碼
- 是否有待更新的文檔
- 是否有待處理的 PR 反饋

## 🎯 使用場景

1. **每天開始工作時**
   - 快速恢復昨天的進度
   - 確認今天的任務

2. **中斷後繼續**
   - 會話超時後重新開始
   - 切換設備後繼續工作

3. **任務切換**
   - 從一個功能切換到另一個
   - 處理緊急問題後回到原任務

## 💾 錯誤處理機制

### 優雅降級策略
如果 last-session.yml 不存在：
1. 從 git log 提取最近工作
2. 從 PROJECT_CONTEXT.md 讀取狀態
3. 從未提交文件推斷進度
4. 提供智能恢復建議

### 恢復建議
- 有遺留工作 → 建議使用 `/context` 確認
- 無遺留工作 → 建議使用 `/plan` 開始新任務
- 文檔缺失 → 引導使用對應命令創建

## 🔗 命令連動機制

**初始化流程**：
```
/start → /meta → /sync
 創建      定制     同步
 文檔      規範     狀態
```

**日常工作流**：
```
/sync → /context → /plan → [coding] → /update-spec
 恢復      確認      規劃      實現        終結
 狀態      理解      任務                 週期
```

**文檔關係**：
- `/start` 創建 → `/sync` 讀取
- `/learn` 更新 → `/sync` 同步
- `/meta` 定制 → `/sync` 應用

這確保你永遠不會失去工作上下文！