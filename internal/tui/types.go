package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"claude-proxy/internal/ollama"
	"claude-proxy/internal/template"
	"claude-proxy/internal/terminal"
)

// ViewType 視圖類型
type ViewType int

const (
	ViewMain ViewType = iota
	ViewTerminals
	ViewTemplates
	ViewSettings
)

// String 返回視圖類型的字符串表示
func (v ViewType) String() string {
	switch v {
	case ViewMain:
		return "主畫面"
	case ViewTerminals:
		return "終端管理"
	case ViewTemplates:
		return "模板管理"
	case ViewSettings:
		return "設置"
	default:
		return "未知視圖"
	}
}

// Model TUI 主模型
type Model struct {
	// 核心組件
	terminalManager *terminal.TerminalManager
	ollamaClient    *ollama.OllamaClient
	templateManager *template.TemplateManager

	// UI 狀態
	currentView ViewType
	width       int
	height      int
	styles      *Styles

	// 應用狀態
	status     string
	lastUpdate time.Time
	terminals  []TerminalInfo

	// 輸入狀態
	inputMode   bool
	inputBuffer string
	inputPrompt string

	// 選擇狀態
	selectedIndex int
	maxItems      int
}

// TerminalInfo 終端信息顯示結構
type TerminalInfo struct {
	Name     string
	Type     string
	Status   string
	LastUsed time.Time
	PID      int
}

// TemplateInfo 模板信息顯示結構
type TemplateInfo struct {
	ID          string
	Name        string
	Category    string
	Description string
	UsageCount  int
}

// TickMsg 定時更新消息
type TickMsg struct {
	Time time.Time
}

// TerminalStatusMsg 終端狀態更新消息
type TerminalStatusMsg struct {
	Name   string
	Status string
}

// QueryResultMsg 查詢結果消息
type QueryResultMsg struct {
	Query      string
	Result     string
	TemplateID string
	Success    bool
}

// ErrorMsg 錯誤消息
type ErrorMsg struct {
	Err error
}

// InputCompleteMsg 輸入完成消息
type InputCompleteMsg struct {
	Input string
	Mode  string
}

// WindowSizeMsg 窗口大小變化消息
type WindowSizeMsg struct {
	Width  int
	Height int
}

// Config TUI 配置
type Config struct {
	// 外觀設置
	ShowBorder    bool `json:"show_border"`
	ShowStatus    bool `json:"show_status"`
	ShowHelp      bool `json:"show_help"`

	// 更新間隔
	UpdateInterval time.Duration `json:"update_interval"`

	// 顏色主題
	Theme string `json:"theme"`

	// 快捷鍵設置
	KeyBindings map[string]string `json:"key_bindings"`
}

// DefaultConfig 返回默認配置
func DefaultConfig() *Config {
	return &Config{
		ShowBorder:     true,
		ShowStatus:     true,
		ShowHelp:       true,
		UpdateInterval: 1 * time.Second,
		Theme:          "default",
		KeyBindings: map[string]string{
			"quit":         "q",
			"terminals":    "1",
			"templates":    "2",
			"settings":     "3",
			"up":           "k",
			"down":         "j",
			"enter":        "enter",
			"escape":       "esc",
		},
	}
}

// AppState 應用狀態
type AppState struct {
	Running        bool
	TerminalCount  int
	ActiveTerminal string
	LastQuery      string
	SystemHealth   string
}

// KeyMap 按鍵映射
type KeyMap struct {
	Up     tea.KeyMsg
	Down   tea.KeyMsg
	Enter  tea.KeyMsg
	Escape tea.KeyMsg
	Quit   tea.KeyMsg
	Help   tea.KeyMsg
}

// DefaultKeyMap 返回默認按鍵映射
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up:     tea.KeyMsg{Type: tea.KeyUp},
		Down:   tea.KeyMsg{Type: tea.KeyDown},
		Enter:  tea.KeyMsg{Type: tea.KeyEnter},
		Escape: tea.KeyMsg{Type: tea.KeyEsc},
		Quit:   tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}},
		Help:   tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'?'}},
	}
}

// MenuItem 菜單項
type MenuItem struct {
	Title       string
	Description string
	Action      string
	Shortcut    string
	Enabled     bool
}

// Menu 菜單結構
type Menu struct {
	Title        string
	Items        []MenuItem
	SelectedItem int
	MaxDisplay   int
}

// Theme 主題配置
type Theme struct {
	Primary   string
	Secondary string
	Success   string
	Warning   string
	Error     string
	Muted     string
	Border    string
	Highlight string
}

// DefaultTheme 返回默認主題
func DefaultTheme() Theme {
	return Theme{
		Primary:   "#007ACC",
		Secondary: "#6C7B7F",
		Success:   "#28A745",
		Warning:   "#FFC107",
		Error:     "#DC3545",
		Muted:     "#6C757D",
		Border:    "#E9ECEF",
		Highlight: "#FFF3CD",
	}
}