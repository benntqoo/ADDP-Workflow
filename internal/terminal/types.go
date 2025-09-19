package terminal

import (
	"bufio"
	"os/exec"
	"sync"
)

// TerminalType 表示支援的 AI 終端類型
type TerminalType int

const (
	TypeClaudeCode TerminalType = iota // Claude Code CLI
	TypeGeminiCLI                      // Gemini CLI
	TypeCursor                         // Cursor CLI
	TypeAider                          // Aider CLI
	TypeCustom                         // 自定義終端
)

// String 返回終端類型的字符串表示
func (t TerminalType) String() string {
	switch t {
	case TypeClaudeCode:
		return "claude"
	case TypeGeminiCLI:
		return "gemini"
	case TypeCursor:
		return "cursor"
	case TypeAider:
		return "aider"
	case TypeCustom:
		return "custom"
	default:
		return "unknown"
	}
}

// CommandName 返回終端類型對應的命令名
func (t TerminalType) CommandName() string {
	switch t {
	case TypeClaudeCode:
		return "claude"
	case TypeGeminiCLI:
		return "gemini"
	case TypeCursor:
		return "cursor"
	case TypeAider:
		return "aider"
	default:
		return "bash"
	}
}

// TerminalStatus 表示終端狀態
type TerminalStatus int

const (
	StatusStopped TerminalStatus = iota // 已停止
	StatusStarting                      // 啟動中
	StatusRunning                       // 運行中
	StatusStopping                      // 停止中
	StatusError                         // 錯誤狀態
)

// String 返回終端狀態的字符串表示
func (s TerminalStatus) String() string {
	switch s {
	case StatusStopped:
		return "stopped"
	case StatusStarting:
		return "starting"
	case StatusRunning:
		return "running"
	case StatusStopping:
		return "stopping"
	case StatusError:
		return "error"
	default:
		return "unknown"
	}
}

// Terminal 表示一個 AI 終端實例
type Terminal struct {
	Name     string         // 終端名稱
	Type     TerminalType   // 終端類型
	Status   TerminalStatus // 終端狀態
	Process  *exec.Cmd      // 底層進程
	Stdin    *bufio.Writer  // 標準輸入寫入器
	Stdout   *bufio.Scanner // 標準輸出掃描器
	LastUsed int64          // 最後使用時間戳
	mu       sync.RWMutex   // 保護並發訪問的鎖
}

// GetStatus 安全地獲取終端狀態
func (t *Terminal) GetStatus() TerminalStatus {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.Status
}

// SetStatus 安全地設置終端狀態
func (t *Terminal) SetStatus(status TerminalStatus) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Status = status
}

// IsRunning 檢查終端是否正在運行
func (t *Terminal) IsRunning() bool {
	return t.GetStatus() == StatusRunning
}

// TerminalConfig 終端配置
type TerminalConfig struct {
	Type        TerminalType          // 終端類型
	Name        string                // 終端名稱
	WorkingDir  string                // 工作目錄
	Environment map[string]string     // 環境變量
	Args        []string              // 額外參數
}

// Manager 介面定義終端管理器的行為
type Manager interface {
	// StartTerminal 啟動指定的終端
	StartTerminal(config TerminalConfig) error

	// StopTerminal 停止指定名稱的終端
	StopTerminal(name string) error

	// SendCommand 向指定終端發送命令
	SendCommand(name string, command string) error

	// GetTerminal 獲取指定名稱的終端
	GetTerminal(name string) (*Terminal, bool)

	// ListTerminals 列出所有終端
	ListTerminals() []*Terminal

	// IsHealthy 檢查終端管理器是否健康
	IsHealthy() bool
}