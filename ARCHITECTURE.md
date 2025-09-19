# 正確架構設計文檔
*AI Terminal Proxy - 終端代理 + MCP 獨立雙軌架構*

## 🎯 架構概覽（糾正版）

```
┌─────────────────────────────────────────────────────────────┐
│                    Golang GUI 終端代理                      │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐ │
│  │ Bubble Tea  │  │   Ollama    │  │    Terminal         │ │
│  │     TUI     │  │ 查詢優化器    │  │     Manager         │ │
│  └─────────────┘  └─────────────┘  └─────────────────────┘ │
└─────────────────────┬───────────────────────────────────────┘
                      │ 優化後命令發送
┌─────────────────────┴───────────────────────────────────────┐
│                   多終端代理層                               │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐ │
│  │ Claude Code │  │ Gemini CLI  │  │  Cursor/Aider       │ │
│  │   Terminal  │  │   Terminal  │  │    Terminals        │ │
│  └─────────────┘  └─────────────┘  └─────────────────────┘ │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│              Python MCP 獨立服務（用戶可選配置）             │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐ │
│  │     TDD     │  │  Workflow   │  │      Memory         │ │
│  │   Tools     │  │   Manager   │  │     System          │ │
│  └─────────────┘  └─────────────┘  └─────────────────────┘ │
└─────────────────────────────────────────────────────────────┘
```

## 🔧 Golang 終端代理設計

### 核心職責
1. **GUI 界面管理**：提供統一的終端操作界面
2. **多終端代理**：管理 claude/gemini/cursor/aider 等多個終端進程
3. **查詢優化**：使用 Ollama 優化用戶輸入
4. **命令路由**：將優化後的命令發送到對應的 AI 工具終端

### 主要組件

#### 1. 終端管理器
```go
// internal/terminal/manager.go
package terminal

import (
    "bufio"
    "context"
    "os/exec"
    "sync"
)

type TerminalManager struct {
    terminals map[string]*Terminal
    mu        sync.RWMutex
}

type Terminal struct {
    Name    string
    Type    TerminalType
    Process *exec.Cmd
    Stdin   *bufio.Writer
    Stdout  *bufio.Scanner
    Active  bool
}

type TerminalType int
const (
    TypeClaudeCode TerminalType = iota
    TypeGeminiCLI
    TypeCursor
    TypeAider
    TypeCustom
)

func NewTerminalManager() *TerminalManager {
    return &TerminalManager{
        terminals: make(map[string]*Terminal),
    }
}

func (tm *TerminalManager) StartTerminal(name string, termType TerminalType) error {
    cmd := tm.getTerminalCommand(termType)

    terminal := &Terminal{
        Name:    name,
        Type:    termType,
        Process: cmd,
        Active:  true,
    }

    // 設置輸入輸出管道
    stdin, err := cmd.StdinPipe()
    if err != nil {
        return err
    }
    terminal.Stdin = bufio.NewWriter(stdin)

    stdout, err := cmd.StdoutPipe()
    if err != nil {
        return err
    }
    terminal.Stdout = bufio.NewScanner(stdout)

    // 啟動進程
    if err := cmd.Start(); err != nil {
        return err
    }

    tm.mu.Lock()
    tm.terminals[name] = terminal
    tm.mu.Unlock()

    return nil
}

func (tm *TerminalManager) SendCommand(terminalName string, command string) error {
    tm.mu.RLock()
    terminal, exists := tm.terminals[terminalName]
    tm.mu.RUnlock()

    if !exists {
        return fmt.Errorf("terminal %s not found", terminalName)
    }

    _, err := terminal.Stdin.WriteString(command + "\n")
    if err != nil {
        return err
    }

    return terminal.Stdin.Flush()
}

func (tm *TerminalManager) getTerminalCommand(termType TerminalType) *exec.Cmd {
    switch termType {
    case TypeClaudeCode:
        return exec.Command("claude")
    case TypeGeminiCLI:
        return exec.Command("gemini")
    case TypeCursor:
        return exec.Command("cursor", "--cli")
    case TypeAider:
        return exec.Command("aider")
    default:
        return exec.Command("bash")
    }
}
```

#### 2. Ollama 查詢優化器
```go
// internal/optimizer/ollama.go
package optimizer

import (
    "context"
    "github.com/ollama/ollama/api"
)

type QueryOptimizer struct {
    client *api.Client
    model  string
}

func NewQueryOptimizer() (*QueryOptimizer, error) {
    client, err := api.ClientFromEnvironment()
    if err != nil {
        return nil, err
    }

    return &QueryOptimizer{
        client: client,
        model:  "qwen2.5:14b",
    }, nil
}

func (qo *QueryOptimizer) OptimizeForTerminal(ctx context.Context, query string, targetTerminal TerminalType) (string, error) {
    template := qo.getOptimizationTemplate(targetTerminal)

    req := &api.ChatRequest{
        Model: qo.model,
        Messages: []api.Message{
            {Role: "system", Content: template},
            {Role: "user", Content: query},
        },
        Stream: false,
    }

    resp, err := qo.client.Chat(ctx, req)
    if err != nil {
        return "", err
    }

    return resp.Message.Content, nil
}

func (qo *QueryOptimizer) getOptimizationTemplate(termType TerminalType) string {
    switch termType {
    case TypeClaudeCode:
        return `你是 Claude Code 命令優化專家。將用戶的自然語言請求轉換為最適合的 Claude Code 命令格式。

規則：
1. 使用 Claude Code 的內建命令格式 (/sync, /plan, /context 等)
2. 保持簡潔明確
3. 如果是代碼相關，添加具體的文件或模塊信息
4. 輸出格式：直接輸出優化後的命令，不要額外解釋

用戶輸入：`

    case TypeGeminiCLI:
        return `你是 Gemini CLI 查詢優化專家。將用戶請求優化為適合 Gemini 的自然語言提示。

規則：
1. 使用清晰的自然語言描述
2. 如果涉及代碼，要求提供具體示例
3. 結構化表達需求
4. 輸出格式：優化後的自然語言查詢

用戶輸入：`

    case TypeCursor:
        return `你是 Cursor AI 提示優化專家。將用戶請求轉換為 Cursor 最容易理解的代碼相關查詢。

規則：
1. 專注於代碼編輯和重構
2. 明確指出要修改的文件和位置
3. 提供期望的輸出格式
4. 輸出格式：優化後的 Cursor 查詢

用戶輸入：`

    case TypeAider:
        return `你是 Aider 命令優化專家。將用戶請求轉換為 Aider 的自然語言編程指令。

規則：
1. 使用 Aider 偏好的編程語言描述
2. 明確要修改或創建的文件
3. 提供清晰的功能描述
4. 輸出格式：優化後的 Aider 指令

用戶輸入：`

    default:
        return "優化以下查詢使其更加清晰和具體："
    }
}
```

#### 3. GUI 界面
```go
// internal/gui/model.go
package gui

import (
    "github.com/charmbracelet/bubbles/list"
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
    // 界面狀態
    width, height   int
    activePane      Pane

    // 輸入組件
    textInput       textinput.Model
    terminalList    list.Model

    // 終端管理
    terminalManager *terminal.TerminalManager
    queryOptimizer  *optimizer.QueryOptimizer

    // 終端輸出
    terminals       map[string]*TerminalOutput
    selectedTerminal string
}

type Pane int
const (
    PaneInput Pane = iota
    PaneTerminalList
    PaneOutput
)

type TerminalOutput struct {
    Name     string
    Type     terminal.TerminalType
    Content  []string
    Active   bool
}

func (m Model) Init() tea.Cmd {
    return tea.Batch(
        textinput.Blink,
        m.loadAvailableTerminals(),
    )
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "tab":
            return m.switchPane()
        case "enter":
            if m.activePane == PaneInput {
                return m.handleUserInput()
            }
        case "1", "2", "3", "4":
            return m.selectTerminal(msg.String())
        }
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
    }

    // 更新子組件
    var cmd tea.Cmd
    m.textInput, cmd = m.textInput.Update(msg)
    cmds = append(cmds, cmd)

    return m, tea.Batch(cmds...)
}

func (m Model) View() string {
    var sections []string

    // 頂部：輸入區域
    sections = append(sections, m.renderInputSection())

    // 中部：終端選擇
    sections = append(sections, m.renderTerminalSelection())

    // 底部：輸出區域
    sections = append(sections, m.renderOutputSection())

    // 狀態欄
    sections = append(sections, m.renderStatusBar())

    return strings.Join(sections, "\n")
}

func (m Model) handleUserInput() (Model, tea.Cmd) {
    userInput := m.textInput.Value()
    if userInput == "" {
        return m, nil
    }

    // 獲取當前選中的終端
    selectedTerminal := m.getSelectedTerminal()
    if selectedTerminal == nil {
        return m, nil
    }

    // 使用 Ollama 優化查詢
    optimizedQuery, err := m.queryOptimizer.OptimizeForTerminal(
        context.Background(),
        userInput,
        selectedTerminal.Type,
    )
    if err != nil {
        return m, m.showError(err)
    }

    // 發送到終端
    err = m.terminalManager.SendCommand(selectedTerminal.Name, optimizedQuery)
    if err != nil {
        return m, m.showError(err)
    }

    // 清空輸入
    m.textInput.SetValue("")

    // 記錄操作
    m.addToOutput(selectedTerminal.Name, fmt.Sprintf("> %s", userInput))
    m.addToOutput(selectedTerminal.Name, fmt.Sprintf("優化後: %s", optimizedQuery))

    return m, nil
}
```

## 🐍 Python MCP 獨立服務設計

### 核心理念
- **完全獨立**：與 Golang GUI 無任何聯動
- **用戶可選**：根據使用的 AI 工具自行配置
- **TDD 驅動**：專注於開發流程標準化

### MCP 服務架構
```python
# src/mcp_server/__main__.py
import asyncio
import argparse
from mcp.server import Server

class ADDPMCPServer:
    """獨立的 ADDP MCP 服務器，用戶可選配置"""

    def __init__(self):
        self.server = Server("addp-tdd-tools")
        self.setup_tools()

    def setup_tools(self):
        """設置 TDD 流程化工具"""
        from .tools import (
            InitProjectTool,
            PlanUpdateTool,
            TestRunTool,
            GuardValidateTool,
            MemoryManagerTool,
            SpecGeneratorTool
        )

        # 註冊工具
        self.server.register_tool(InitProjectTool())
        self.server.register_tool(PlanUpdateTool())
        self.server.register_tool(TestRunTool())
        self.server.register_tool(GuardValidateTool())
        self.server.register_tool(MemoryManagerTool())
        self.server.register_tool(SpecGeneratorTool())

    async def run(self):
        """運行 MCP 服務器"""
        if "--stdio" in sys.argv:
            await self.server.run_stdio()
        else:
            await self.server.run_sse("http://localhost:8000")

async def main():
    parser = argparse.ArgumentParser(description="ADDP TDD MCP Server")
    parser.add_argument("--stdio", action="store_true", help="Run in STDIO mode")
    parser.add_argument("--port", type=int, default=8000, help="HTTP server port")

    args = parser.parse_args()

    server = ADDPMCPServer()
    await server.run()

if __name__ == "__main__":
    asyncio.run(main())
```

### TDD 工具集
```python
# src/mcp_server/tools/tdd_tools.py
from mcp.types import Tool

class TestRunTool(Tool):
    """TDD 測試執行工具"""

    def __init__(self):
        super().__init__(
            name="test.run",
            description="Run tests with TDD validation"
        )

    async def execute(self, test_type="all", **kwargs):
        """執行測試並驗證 TDD 流程"""
        result = {
            "test_type": test_type,
            "status": "running",
            "results": [],
            "coverage": {},
            "tdd_validation": {}
        }

        # 1. 執行測試
        if test_type in ["all", "unit"]:
            unit_results = await self._run_unit_tests()
            result["results"].append(unit_results)

        if test_type in ["all", "integration"]:
            integration_results = await self._run_integration_tests()
            result["results"].append(integration_results)

        # 2. 生成覆蓋率報告
        result["coverage"] = await self._generate_coverage_report()

        # 3. TDD 流程驗證
        result["tdd_validation"] = await self._validate_tdd_process()

        return result

class GuardValidateTool(Tool):
    """開發門禁驗證工具"""

    async def execute(self, phase, artifacts, **kwargs):
        """執行開發門禁驗證"""
        validation = {
            "phase": phase,
            "valid": True,
            "gates": {},
            "violations": []
        }

        # TDD 門禁
        if phase == "development":
            tdd_gate = await self._validate_tdd_gate(artifacts)
            validation["gates"]["tdd"] = tdd_gate
            if not tdd_gate["passed"]:
                validation["valid"] = False

        # 測試覆蓋率門禁
        coverage_gate = await self._validate_coverage_gate(artifacts)
        validation["gates"]["coverage"] = coverage_gate
        if coverage_gate["coverage"] < 0.8:
            validation["violations"].append("Coverage below 80%")

        return validation
```

## 🔄 用戶配置方式

### 1. Claude Code 用戶配置
```json
// ~/.claude/mcp.json
{
  "servers": {
    "addp-tdd": {
      "command": "python",
      "args": ["-m", "src.mcp_server", "--stdio"],
      "cwd": "/path/to/claude-terminal-proxy"
    }
  }
}
```

### 2. Cursor 用戶配置
```json
// .cursor/mcp_servers.json
{
  "addp-tdd-tools": {
    "command": "python",
    "args": ["-m", "src.mcp_server", "--stdio"],
    "env": {
      "ADDP_MODE": "cursor"
    }
  }
}
```

### 3. 其他 AI 工具
用戶根據各自 AI 工具的 MCP 配置方式自行設置。

## 📁 正確的項目結構

```
D:\Code\fos\AI\claude\
├── cmd/
│   └── main.go              # Golang 主程序入口
├── internal/
│   ├── gui/                 # TUI 界面
│   ├── terminal/            # 終端管理
│   ├── optimizer/           # Ollama 查詢優化
│   └── config/              # 配置管理
├── src/                     # Python MCP 服務（獨立）
│   └── mcp_server/
│       ├── __main__.py      # MCP 服務入口
│       ├── tools/           # TDD 工具集
│       └── workflow/        # 工作流管理
├── templates/               # 查詢優化模板
├── configs/                 # 配置文件模板
├── scripts/                 # 部署腳本
├── go.mod                   # Go 模塊定義
└── pyproject.toml          # Python 包定義
```

## 🎯 開發優先級重新排列

### Week 1-2: Golang 終端代理核心
1. 終端管理器實現
2. Ollama 集成和查詢優化
3. 基礎 TUI 界面

### Week 3-4: 多終端適配
1. Claude Code 終端適配
2. Gemini CLI 終端適配
3. Cursor/Aider 終端適配

### Week 5-7: Python MCP 獨立開發
1. TDD 工具集實現
2. 工作流管理
3. 用戶配置文檔

---

這個糾正版架構現在符合您的真實需求：

1. **Golang GUI** = 終端代理 + Ollama 優化 + 多終端管理
2. **Python MCP** = 獨立的 TDD 流程化工具，用戶可選配置
3. **兩者完全解耦**，各自獨立開發和部署

## 🎯 目標與驗收標準 (Objectives & DoD)

### 項目目標 (Objectives)

#### Golang 終端代理目標
- **本地智能代理**：提供終端版 AI IDE，解析斜線命令，執行本地操作（Ollama、模板化、查詢優化），僅在需要時將優化後的最小提示移交給外部 AI CLI 工具
- **統一終端管理**：透過單一 GUI 界面管理多個 AI CLI 工具（Claude Code、Gemini CLI、Cursor、Aider），提供一致的操作體驗
- **Token 經濟優化**：通過本地提示優化、上下文打包、緩存和預算護欄減少往返和冗餘，節省 30-50% 的 token 使用
- **跨 CLI 橋接**：通過"提示+上下文移交"和按工具適配器標準化 Claude/Gemini/Cursor 等工具之間的差異

#### Python MCP 工具目標
- **紀律性交付**：強制執行 TDD 和質量門禁（白名單修補、階段順序、覆蓋率閾值）作為合併阻止條件
- **可審計狀態**：在 `.addp/**` 下持久化運行記錄、決策、產物和分析，使用 MCP_RUN_ID 端到端追蹤流程
- **統一 MCP 工具**：通過 stdio/JSON-RPC 在任何 AI CLI 中暴露 ADDP 一致工具集，具有清晰的契約和安全的副作用

### 完成定義 (Definition of Done)

#### Golang 終端代理 DoD

**基礎功能驗收**：
- [ ] TUI 界面運行流暢，支援 Tab 鍵切換面板，1-4 數字鍵選擇終端
- [ ] 成功啟動和管理 4 種終端類型：Claude Code、Gemini CLI、Cursor、Aider
- [ ] Ollama 集成工作正常，查詢優化響應時間 < 2 秒
- [ ] 終端進程管理穩定：啟動/停止/重啟功能無錯誤
- [ ] 跨平台兼容：Windows、Linux、macOS 一鍵編譯和運行

**優化效果驗收**：
- [ ] 針對不同 AI 工具的查詢優化模板完整且有效
- [ ] Token 節省率達到 30% 以上（與直接使用 AI 工具對比）
- [ ] 查詢優化準確率 > 85%（用戶滿意度評估）
- [ ] 本地緩存命中率 > 60%，重複查詢響應時間 < 1 秒

**穩定性驗收**：
- [ ] 連續運行 2+ 小時無崩潰
- [ ] 記憶體使用 < 100MB（空閒時）
- [ ] CPU 使用率 < 5%（空閒時）
- [ ] 終端異常時自動恢復機制有效

#### Python MCP 獨立服務 DoD

**MCP 協議合規**：
- [ ] 通過 stdio/JSON-RPC 運行，具有文檔化的請求/響應模式和超時設置
- [ ] 可用工具集：`query.optimize`、`plan.update`、`fs.apply_patch`（白名單+dry-run）、`test.run`、`guard.validate_flow`、`mem.save_phase`、`prompt.render`、`context.pack`
- [ ] 所有工具具有結構化 IO：明確的參數、返回值、錯誤處理

**安全與門禁**：
- [ ] `fs.apply_patch` 拒絕非白名單路徑並記錄拒絕操作
- [ ] `test.run` 返回機器可讀報告；`guard.validate_flow` 在測試失敗、缺少 Red→Green 或覆蓋率 < 配置閾值時阻止
- [ ] 強制執行階段順序（ADDP）；違規時報告並提供修復提示

**本地優先流程**：
- [ ] `/init` 初始化 `.addp/{specifications,workflows,memory,queries,gates,sync,configs}` 並種子化模板
- [ ] 完整工作流：`query.optimize` → `prompt.render` → `context.pack` 產生最小化提示包準備移交
- [ ] 與至少一個 AI CLI 的移交演示工作，並將結果往返回本地記憶/狀態

**持久化與可觀測性**：
- [ ] 每次調用寫入 MCP_RUN_ID、輸入、輸出和門禁結果到 `.addp/**`
- [ ] 記錄基本分析（延遲 P50/P95、緩存命中率、失敗代碼、token 估計）
- [ ] 工具執行成功率 > 95%，錯誤有清晰的診斷信息

**配置與文檔**：
- [ ] 提供 Claude Code、Cursor 的 MCP 配置模板
- [ ] 配置模板在各 AI 工具中可用性 > 90%
- [ ] 故障排除文檔完整，常見問題有明確解決方案

#### 整體項目 DoD

**部署與用戶體驗**：
- [ ] README 快速開始流程端到端運行成功
- [ ] 新用戶 5 分鐘內成功啟動並使用基本功能
- [ ] 一鍵安裝腳本支援主流平台
- [ ] Windows 控制台引號和編碼行為已記錄

**文檔完整性**：
- [ ] 架構接口目錄列出工具名稱、參數、返回值、錯誤和副作用
- [ ] 開發計劃將里程碑與可演示命令和 CI 門禁綁定
- [ ] 包含實際使用案例的最佳實踐文檔

**質量保證**：
- [ ] 煙霧測試在 Windows/macOS/Linux 通過
- [ ] 核心功能單元測試覆蓋率 > 80%
- [ ] 集成測試覆蓋主要用戶流程
- [ ] 性能基準測試建立並可重現

### 最小可行工具集 (Minimum Viable Toolset)

#### Python MCP 核心工具
1. **query.optimize** - Ollama 驅動的查詢優化
2. **plan.update** - 項目計劃更新和狀態管理
3. **fs.apply_patch** - 安全的文件系統修補（白名單保護）
4. **test.run** - TDD 測試執行和驗證
5. **guard.validate_flow** - 開發流程門禁驗證
6. **mem.save_phase** - 項目記憶和狀態持久化
7. **prompt.render** - 提示詞模板渲染
8. **context.pack** - 上下文打包和優化

#### 持久化目錄結構
```
.addp/
├── specifications/     # 項目規格和需求文檔
├── workflows/         # ADDP 四階段工作流狀態
├── memory/           # 跨工具記憶和上下文
├── queries/          # 查詢優化緩存和模板
├── gates/            # 質量門禁檢查記錄
├── sync/             # 工具狀態同步數據
└── configs/          # 配置管理和模板
```

#### 驗收演示場景
1. **Golang 代理演示**：啟動代理 → 切換終端 → 輸入查詢 → 顯示 Ollama 優化 → 發送到選定 AI 工具
2. **Python MCP 演示**：配置 MCP → 調用 test.run → 驗證 TDD 流程 → 保存到 .addp/ → 查看結構化報告
3. **端到端流程**：本地優化 → MCP 工具處理 → 移交到 AI CLI → 結果返回 → 狀態持久化

這個 DoD 確保兩個軌道都有明確的完成標準，並且完全遵循獨立架構原則。