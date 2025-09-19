# 🚀 AI启动器使用指南

## 快速開始

### 1. 環境準備

#### 必要條件
- **Go 1.24+**: 用於編譯主程序
- **Ollama**: 本地 LLM 服務（推薦 qwen2.5:14b 模型）
- **終端環境**: 支援 Windows/Linux/macOS

#### 安裝 Ollama 和模型
```bash
# 1. 安裝 Ollama (參考 https://ollama.ai)
curl -fsSL https://ollama.ai/install.sh | sh

# 2. 下載推薦模型
ollama pull qwen2.5:14b

# 3. 驗證安裝
ollama list
```

### 2. 編譯和安裝

#### 從源碼編譯
```bash
# 克隆項目
git clone <repository-url>
cd claude

# 編譯程序
go build -o ai-launcher ./cmd/proxy

# 運行程序
./ai-launcher
```

#### 使用 Docker
```bash
# 構建鏡像
docker build -t ai-launcher .

# 運行容器
docker run -p 8080:8080 ai-launcher
```

## 核心功能使用

### 1. TUI 界面操作

啟動程序後，您將看到現代化的終端界面：

```
┌─ AI启动器 ─────────────────────────────────┐
│                                                  │
│  [1] 啟動 Claude Code                            │
│  [2] 啟動 Gemini CLI                             │
│  [3] 啟動 Cursor                                 │
│  [4] 啟動 Aider                                  │
│  [5] 查詢優化                                    │
│  [6] 終端管理                                    │
│  [q] 退出                                        │
│                                                  │
└──────────────────────────────────────────────────┘
```

#### 基本操作快捷鍵
- **數字鍵 1-6**: 選擇功能
- **ESC**: 返回主選單
- **q**: 退出程序
- **↑/↓**: 在列表中導航
- **Enter**: 確認選擇

### 2. 終端管理功能

#### 啟動不同類型的終端
```bash
# 選擇對應選項啟動終端
[1] Claude Code    # 啟動 Claude Code 工作環境
[2] Gemini CLI     # 啟動 Google Gemini CLI
[3] Cursor         # 啟動 Cursor AI IDE
[4] Aider          # 啟動 Aider 編程助手
```

#### 終端狀態監控
- **運行中**: 🟢 綠色指示器
- **已停止**: 🔴 紅色指示器
- **錯誤**: ⚠️ 黃色警告

### 3. 查詢優化系統

#### 可用模板類型
1. **coding** - 編程任務優化
2. **debug** - 調試問題優化
3. **review** - 代碼審查優化
4. **refactor** - 重構任務優化
5. **test** - 測試相關優化
6. **doc** - 文檔編寫優化

#### 使用查詢優化
```
選擇 [5] 查詢優化 → 輸入原始查詢 → 選擇模板 → 獲得優化結果
```

#### 優化前後對比示例
```
原始查詢: "修復這個 bug"
↓ 使用 debug 模板優化
優化查詢: "請分析這個 JavaScript 錯誤的根本原因，提供具體的修復步驟和預防措施，包括相關的測試用例"
```

## 實際使用場景

### 場景 1: 多終端協作開發

```bash
# 1. 啟動 Claude Code (主要 AI 協作)
選擇 [1] → 配置項目路徑 → 開始 AI 編程

# 2. 啟動 Cursor (IDE 環境)
選擇 [3] → 打開同一項目 → 實時編輯

# 3. 啟動 Aider (代碼審查)
選擇 [4] → 進行代碼審查 → 自動化重構
```

### 場景 2: AI 輔助調試流程

```bash
# 1. 遇到問題時使用查詢優化
選擇 [5] → 輸入: "程序崩潰了"
→ 選擇 debug 模板
→ 獲得: "請提供詳細的錯誤堆棧、復現步驟、環境信息..."

# 2. 將優化後的查詢發送給 AI 助手
複製優化查詢 → 發送給 Claude/Gemini → 獲得精準答案
```

### 場景 3: 團隊協作規範

```bash
# 1. 統一查詢模板
團隊成員都使用相同的優化模板 → 確保問題描述標準化

# 2. 知識積累
優化查詢產生的結果 → 形成團隊知識庫 → 提升整體效率
```

## 高級配置

### 1. 自定義終端配置

創建配置文件 `config.yaml`:
```yaml
terminals:
  claude_code:
    command: "claude"
    args: ["--project", "current"]
    env:
      CLAUDE_API_KEY: "your-api-key"

  gemini_cli:
    command: "gemini"
    args: ["--interactive"]

  cursor:
    command: "cursor"
    args: ["."]
```

### 2. Ollama 模型配置

```bash
# 使用不同模型
export OLLAMA_MODEL="qwen2.5:32b"  # 更大模型，更好效果
export OLLAMA_MODEL="qwen2.5:7b"   # 較小模型，更快響應

# 調整生成參數
export OLLAMA_TEMPERATURE="0.7"    # 創造性程度
export OLLAMA_MAX_TOKENS="2048"    # 最大輸出長度
```

### 3. 模板自定義

在代碼中修改 `internal/template/templates.go`:
```go
// 添加自定義模板
"my_template": {
    ID: "my_template",
    Name: "我的模板",
    Description: "自定義查詢優化模板",
    Template: "根據 {{.domain}} 領域的最佳實踐，請 {{.query}}",
    Variables: []string{"domain", "query"},
}
```

## 故障排除

### 常見問題

#### 1. Ollama 連接失敗
```bash
# 檢查 Ollama 服務狀態
ollama list
curl http://localhost:11434/api/version

# 重啟 Ollama 服務
systemctl restart ollama  # Linux
brew services restart ollama  # macOS
```

#### 2. 終端啟動失敗
```bash
# 檢查終端程序是否安裝
which claude  # 檢查 Claude Code
which cursor  # 檢查 Cursor
which aider   # 檢查 Aider

# 更新 PATH 環境變量
export PATH=$PATH:/path/to/your/tools
```

#### 3. TUI 顯示異常
```bash
# 檢查終端支援
echo $TERM
export TERM=xterm-256color

# 調整終端大小
resize  # Linux
```

### 日誌調試

```bash
# 啟用詳細日誌
./ai-launcher --verbose

# 查看日誌文件
tail -f ~/.ai-launcher/logs/app.log
```

## 性能優化建議

### 1. 硬件要求
- **RAM**: 建議 8GB+ （Ollama 模型需要 4-6GB）
- **CPU**: 多核心處理器（並行終端處理）
- **存儲**: SSD 推薦（快速模型載入）

### 2. 網絡配置
```bash
# 如果使用遠端 Ollama
export OLLAMA_HOST="http://remote-server:11434"

# 設置代理（如需要）
export HTTP_PROXY="http://proxy:8080"
export HTTPS_PROXY="http://proxy:8080"
```

### 3. 模型選擇策略
- **開發階段**: 使用 qwen2.5:7b （快速響應）
- **生產階段**: 使用 qwen2.5:14b （平衡性能）
- **高品質需求**: 使用 qwen2.5:32b （最佳效果）

## 進階使用技巧

### 1. 批量查詢優化
```bash
# 準備查詢文件
echo "修復登錄問題" > queries.txt
echo "優化數據庫查詢" >> queries.txt
echo "添加單元測試" >> queries.txt

# 批量處理（需要腳本支援）
./ai-launcher batch-optimize --input queries.txt --template coding
```

### 2. 整合到 CI/CD
```yaml
# GitHub Actions 示例
- name: Optimize commit messages
  run: |
    message=$(git log -1 --pretty=%B)
    optimized=$(ai-launcher optimize --template doc --query "$message")
    echo "Optimized: $optimized"
```

### 3. 團隊共享配置
```bash
# 版本控制團隊配置
git add .ai-launcher/
git commit -m "Add team ai-launcher configuration"
```

---

## 📚 更多資源

- **項目文檔**: `ARCHITECTURE.md`, `DEVELOPMENT_PLAN.md`
- **API 參考**: `internal/` 目錄下的代碼文檔
- **測試示例**: `test/` 目錄
- **CI/CD 配置**: `.github/workflows/ci.yml`

開始您的 AI 輔助開發之旅！🚀