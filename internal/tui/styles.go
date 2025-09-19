package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// 顏色定義
var (
	Primary   = lipgloss.Color("#007ACC")
	Secondary = lipgloss.Color("#6C7B7F")
	Success   = lipgloss.Color("#28A745")
	Warning   = lipgloss.Color("#FFC107")
	Error     = lipgloss.Color("#DC3545")
	Muted     = lipgloss.Color("#6C757D")
	Border    = lipgloss.Color("#E9ECEF")
	Highlight = lipgloss.Color("#FFF3CD")
)

// Styles TUI 樣式集合
type Styles struct {
	// 容器樣式
	AppStyle    lipgloss.Style
	HeaderStyle lipgloss.Style
	ContentStyle lipgloss.Style
	StatusStyle lipgloss.Style
	HelpStyle   lipgloss.Style

	// 組件樣式
	MenuStyle         lipgloss.Style
	MenuItemStyle     lipgloss.Style
	SelectedItemStyle lipgloss.Style
	DisabledItemStyle lipgloss.Style

	// 狀態樣式
	SuccessStyle lipgloss.Style
	WarningStyle lipgloss.Style
	ErrorStyle   lipgloss.Style
	InfoStyle    lipgloss.Style

	// 特殊樣式
	TitleStyle      lipgloss.Style
	SubtitleStyle   lipgloss.Style
	DescriptionStyle lipgloss.Style
	KeyStyle        lipgloss.Style
	ValueStyle      lipgloss.Style
}

// NewStyles 創建新的樣式集合
func NewStyles() *Styles {
	return &Styles{
		// 容器樣式
		AppStyle: lipgloss.NewStyle().
			Padding(1, 2),

		HeaderStyle: lipgloss.NewStyle().
			Foreground(Primary).
			Bold(true).
			Padding(0, 1).
			MarginBottom(1),

		ContentStyle: lipgloss.NewStyle().
			Padding(0, 1).
			MarginBottom(1),

		StatusStyle: lipgloss.NewStyle().
			Foreground(Muted).
			Padding(0, 1).
			Border(lipgloss.NormalBorder(), true, false, false, false).
			BorderForeground(Border),

		HelpStyle: lipgloss.NewStyle().
			Foreground(Muted).
			Italic(true).
			Padding(0, 1),

		// 組件樣式
		MenuStyle: lipgloss.NewStyle().
			Padding(0, 1),

		MenuItemStyle: lipgloss.NewStyle().
			Padding(0, 1),

		SelectedItemStyle: lipgloss.NewStyle().
			Foreground(Primary).
			Bold(true).
			Padding(0, 1).
			Background(Highlight),

		DisabledItemStyle: lipgloss.NewStyle().
			Foreground(Muted).
			Padding(0, 1),

		// 狀態樣式
		SuccessStyle: lipgloss.NewStyle().
			Foreground(Success).
			Bold(true),

		WarningStyle: lipgloss.NewStyle().
			Foreground(Warning).
			Bold(true),

		ErrorStyle: lipgloss.NewStyle().
			Foreground(Error).
			Bold(true),

		InfoStyle: lipgloss.NewStyle().
			Foreground(Primary),

		// 特殊樣式
		TitleStyle: lipgloss.NewStyle().
			Foreground(Primary).
			Bold(true).
			Underline(true),

		SubtitleStyle: lipgloss.NewStyle().
			Foreground(Secondary).
			Bold(true),

		DescriptionStyle: lipgloss.NewStyle().
			Foreground(Muted),

		KeyStyle: lipgloss.NewStyle().
			Foreground(Primary).
			Bold(true),

		ValueStyle: lipgloss.NewStyle().
			Foreground(Secondary),
	}
}

// RenderBox 渲染帶邊框的盒子
func (s *Styles) RenderBox(title, content string, width int) string {
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Border).
		Padding(1).
		Width(width)

	if title != "" {
		titleStyle := lipgloss.NewStyle().
			Foreground(Primary).
			Bold(true).
			Padding(0, 1)

		return boxStyle.Render(
			titleStyle.Render(title) + "\n\n" + content,
		)
	}

	return boxStyle.Render(content)
}

// RenderList 渲染列表
func (s *Styles) RenderList(items []string, selectedIndex int) string {
	var result string

	for i, item := range items {
		style := s.MenuItemStyle
		prefix := "  "

		if i == selectedIndex {
			style = s.SelectedItemStyle
			prefix = "▶ "
		}

		result += style.Render(prefix + item) + "\n"
	}

	return result
}

// RenderStatus 渲染狀態信息
func (s *Styles) RenderStatus(status, statusType string) string {
	var style lipgloss.Style

	switch statusType {
	case "success":
		style = s.SuccessStyle
	case "warning":
		style = s.WarningStyle
	case "error":
		style = s.ErrorStyle
	default:
		style = s.InfoStyle
	}

	return style.Render(status)
}

// RenderKeyValue 渲染鍵值對
func (s *Styles) RenderKeyValue(key, value string) string {
	return s.KeyStyle.Render(key) + ": " + s.ValueStyle.Render(value)
}

// RenderProgress 渲染進度條
func (s *Styles) RenderProgress(current, total int, width int) string {
	if total == 0 {
		return ""
	}

	progress := float64(current) / float64(total)
	filledWidth := int(progress * float64(width))

	filled := lipgloss.NewStyle().
		Background(Success).
		Render(lipgloss.PlaceHorizontal(filledWidth, lipgloss.Left, ""))

	empty := lipgloss.NewStyle().
		Background(Muted).
		Render(lipgloss.PlaceHorizontal(width-filledWidth, lipgloss.Left, ""))

	return lipgloss.JoinHorizontal(lipgloss.Left, filled, empty)
}

// RenderTable 渲染簡單表格
func (s *Styles) RenderTable(headers []string, rows [][]string) string {
	if len(headers) == 0 || len(rows) == 0 {
		return ""
	}

	// 計算列寬
	colWidths := make([]int, len(headers))
	for i, header := range headers {
		colWidths[i] = len(header)
	}

	for _, row := range rows {
		for i, cell := range row {
			if i < len(colWidths) && len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	var result string

	// 渲染標題
	headerRow := ""
	for i, header := range headers {
		cellStyle := s.KeyStyle.Copy().
			Width(colWidths[i]).
			Align(lipgloss.Left)
		headerRow += cellStyle.Render(header) + " "
	}
	result += headerRow + "\n"

	// 渲染分隔線
	separator := ""
	for _, width := range colWidths {
		separator += lipgloss.NewStyle().
			Width(width).
			Render(lipgloss.PlaceHorizontal(width, lipgloss.Left, "─")) + " "
	}
	result += separator + "\n"

	// 渲染數據行
	for _, row := range rows {
		dataRow := ""
		for i, cell := range row {
			if i < len(colWidths) {
				cellStyle := s.ValueStyle.Copy().
					Width(colWidths[i]).
					Align(lipgloss.Left)
				dataRow += cellStyle.Render(cell) + " "
			}
		}
		result += dataRow + "\n"
	}

	return result
}

// GetStatusIcon 獲取狀態圖標
func GetStatusIcon(status string) string {
	switch status {
	case "running":
		return "🟢"
	case "starting":
		return "🟡"
	case "stopping":
		return "🟠"
	case "stopped":
		return "🔴"
	case "error":
		return "❌"
	default:
		return "⚪"
	}
}

// GetCategoryIcon 獲取分類圖標
func GetCategoryIcon(category string) string {
	switch category {
	case "development":
		return "💻"
	case "maintenance":
		return "🔧"
	case "analysis":
		return "📊"
	case "testing":
		return "🧪"
	case "documentation":
		return "📝"
	default:
		return "📂"
	}
}