# Agent 系統測試計劃

## 測試前準備

### 1. 同步配置到全局目錄
```bash
# Windows - 在 PowerShell 或 CMD 中執行
xcopy /E /I /Y "D:\Code\ai\claude\agents\*.md" "%USERPROFILE%\.claude\agents\"
```

### 2. 確認文件已同步
```bash
dir %USERPROFILE%\.claude\agents\
```

## 測試場景

### 🧪 測試 1：代碼審查（code-reviewer）
```
請求：Review this code for security issues
```
預期：自動使用 code-reviewer agent

### 🧪 測試 2：性能優化（performance-optimizer）
```
請求：This function is running slow, help optimize it
```
預期：自動使用 performance-optimizer agent

### 🧪 測試 3：錯誤調試（bug-hunter）
```
請求：I'm getting an error in my code, help me debug
```
預期：自動使用 bug-hunter agent

### 🧪 測試 4：測試生成（test-automator）
```
請求：Write unit tests for this function
```
預期：自動使用 test-automator agent

### 🧪 測試 5：顯式調用
```
請求：Use the kotlin-expert agent to review this Android code
```
預期：明確使用 kotlin-expert agent

## 驗證點

- [ ] Agent 是否被觸發？
- [ ] 觸發的是否是正確的 agent？
- [ ] 回應是否符合 agent 的專業領域？
- [ ] 開發體驗是否改善？

## 已完成的修改總結

### ✅ 刪除的文件/目錄
1. `config/triggers.yaml` - 不被支援的觸發配置
2. `config/workflows.yaml` - 不需要的工作流配置
3. `agents/specialized/` - 空目錄

### ✅ 修改的 agents
1. `code-reviewer.md` - 優化 description 添加觸發關鍵字
2. `performance-optimizer.md` - 優化 description
3. `test-automator.md` - 優化 description
4. `bug-hunter.md` - 優化 description

### ✅ 新增的文件
1. `agents/kotlin-expert.md` - 取代 context-detector
2. `agents/README_ACTUAL_USAGE.md` - 實際使用指南
3. `config/README.md` - 配置說明

### ✅ 更新的文檔
1. `README.md` - 標記棄用功能，說明刪除原因
2. `README.zh.md` - 同步中文版本

## 測試提示

1. 先測試自動觸發，看是否改善
2. 如果自動觸發不理想，使用顯式調用
3. 注意觀察 Claude Code 的回應中是否有提到使用了哪個 agent