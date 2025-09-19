package tui

import (
	"context"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"ai-launcher/internal/ollama"
	"ai-launcher/internal/template"
	"ai-launcher/internal/terminal"
)

// NewModel 創建新的 TUI 模型
func NewModel() Model {
	return Model{
		terminalManager: terminal.NewTerminalManager(),
		ollamaClient:    ollama.NewOllamaClient("http://localhost:11434"),
		templateManager: template.NewTemplateManager(),
		currentView:     ViewMain,
		status:          "準備就緒",
		lastUpdate:      time.Now(),
		terminals:       make([]TerminalInfo, 0),
		selectedIndex:   0,
		maxItems:        10,
		styles:          NewStyles(),
	}
}

// Init 初始化模型
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tickCmd(),
		m.updateTerminalsCmd(),
	)
}

// Update 更新模型
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)

	case TickMsg:
		m.lastUpdate = msg.Time
		m.updateStatus()
		return m, tickCmd()

	case TerminalStatusMsg:
		m.updateTerminalStatus(msg.Name, msg.Status)
		return m, nil

	case QueryResultMsg:
		m.handleQueryResult(msg)
		return m, nil

	case ErrorMsg:
		m.handleError(msg.Err.Error())
		return m, nil

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	default:
		return m, nil
	}
}

// View 渲染視圖
func (m Model) View() string {
	var content strings.Builder

	// 標題
	content.WriteString(m.renderHeader())
	content.WriteString("\n\n")

	// 主內容區域
	switch m.currentView {
	case ViewMain:
		content.WriteString(m.renderMainView())
	case ViewTerminals:
		content.WriteString(m.renderTerminalsView())
	case ViewTemplates:
		content.WriteString(m.renderTemplatesView())
	case ViewSettings:
		content.WriteString(m.renderSettingsView())
	}

	content.WriteString("\n\n")

	// 狀態欄
	content.WriteString(m.renderStatusBar())

	// 幫助信息
	content.WriteString("\n")
	content.WriteString(m.renderHelp())

	return content.String()
}

// handleKeyMsg 處理按鍵消息
func (m Model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit

	case "1":
		m.switchView(ViewTerminals)
		return m, m.updateTerminalsCmd()

	case "2":
		m.switchView(ViewTemplates)
		return m, nil

	case "3":
		m.switchView(ViewSettings)
		return m, nil

	case "up", "k":
		if m.selectedIndex > 0 {
			m.selectedIndex--
		}
		return m, nil

	case "down", "j":
		if m.selectedIndex < m.maxItems-1 {
			m.selectedIndex++
		}
		return m, nil

	case "enter":
		return m.handleEnterKey()

	case "esc":
		m.switchView(ViewMain)
		return m, nil

	default:
		return m, nil
	}
}

// handleEnterKey 處理回車鍵
func (m Model) handleEnterKey() (tea.Model, tea.Cmd) {
	switch m.currentView {
	case ViewTerminals:
		// 在終端視圖中，回車可以啟動/停止終端
		if len(m.terminals) > m.selectedIndex {
			terminal := m.terminals[m.selectedIndex]
			if terminal.Status == "stopped" {
				err := m.startTerminal(terminal.Name, terminal.Type)
				if err != nil {
					m.handleError(fmt.Sprintf("啟動終端失敗: %v", err))
				}
			} else {
				err := m.stopTerminal(terminal.Name)
				if err != nil {
					m.handleError(fmt.Sprintf("停止終端失敗: %v", err))
				}
			}
		}
		return m, m.updateTerminalsCmd()

	case ViewTemplates:
		// 在模板視圖中，回車可以應用模板
		return m, nil

	default:
		return m, nil
	}
}

// switchView 切換視圖
func (m *Model) switchView(view ViewType) {
	m.currentView = view
	m.selectedIndex = 0
	m.status = fmt.Sprintf("切換到 %s", view.String())
}

// renderHeader 渲染標題
func (m Model) renderHeader() string {
	return fmt.Sprintf("🤖 AI 終端代理 - %s", m.currentView.String())
}

// renderMainView 渲染主視圖
func (m Model) renderMainView() string {
	var content strings.Builder

	content.WriteString("歡迎使用 AI 終端代理系統\n\n")
	content.WriteString("功能概覽:\n")
	content.WriteString("• 多終端管理 (Claude Code, Gemini CLI, Cursor, Aider)\n")
	content.WriteString("• Ollama 本地 LLM 查詢優化\n")
	content.WriteString("• 智能模板系統\n")
	content.WriteString("• 跨平台支持\n\n")

	content.WriteString("快速操作:\n")
	content.WriteString("1️⃣  終端管理  2️⃣  模板管理  3️⃣  系統設置\n")

	return content.String()
}

// renderTerminalsView 渲染終端視圖
func (m Model) renderTerminalsView() string {
	var content strings.Builder

	// 標題
	content.WriteString(m.styles.TitleStyle.Render("📺 終端管理"))
	content.WriteString("\n\n")

	if len(m.terminals) == 0 {
		emptyMsg := m.styles.DescriptionStyle.Render("暫無活動終端")
		helpMsg := m.styles.InfoStyle.Render("使用 Enter 鍵創建新終端")
		content.WriteString(emptyMsg + "\n" + helpMsg)
		return content.String()
	}

	content.WriteString(m.styles.SubtitleStyle.Render("活動終端列表:"))
	content.WriteString("\n\n")

	// 渲染終端列表
	for i, term := range m.terminals {
		style := m.styles.MenuItemStyle
		prefix := "  "
		if i == m.selectedIndex {
			style = m.styles.SelectedItemStyle
			prefix = "▶ "
		}

		statusIcon := GetStatusIcon(term.Status)
		line := fmt.Sprintf("%s%s %s (%s) %s",
			prefix, statusIcon, term.Name, term.Type, term.Status)

		content.WriteString(style.Render(line))
		content.WriteString("\n")
	}

	content.WriteString("\n")
	content.WriteString(m.styles.HelpStyle.Render("💡 按 Enter 啟動/停止選中的終端"))

	return content.String()
}

// renderTemplatesView 渲染模板視圖
func (m Model) renderTemplatesView() string {
	var content strings.Builder

	content.WriteString("📝 模板管理\n\n")

	templates := m.templateManager.GetAvailableTemplates()

	if len(templates) == 0 {
		content.WriteString("暫無可用模板")
		return content.String()
	}

	content.WriteString("可用模板:\n\n")

	for i, tmpl := range templates {
		if i >= m.maxItems {
			break
		}

		prefix := "  "
		if i == m.selectedIndex {
			prefix = "▶ "
		}

		categoryIcon := "📂"
		switch tmpl.Category {
		case "development":
			categoryIcon = "💻"
		case "maintenance":
			categoryIcon = "🔧"
		case "analysis":
			categoryIcon = "📊"
		}

		content.WriteString(fmt.Sprintf("%s%s %s - %s\n",
			prefix, categoryIcon, tmpl.Name, tmpl.Description))
	}

	content.WriteString("\n💡 按 Enter 使用選中的模板")

	return content.String()
}

// renderSettingsView 渲染設置視圖
func (m Model) renderSettingsView() string {
	var content strings.Builder

	content.WriteString("⚙️ 系統設置\n\n")

	// Ollama 連接狀態
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ollamaHealthy := m.ollamaClient.IsHealthy(ctx)
	ollamaStatus := "🔴 離線"
	if ollamaHealthy {
		ollamaStatus = "🟢 在線"
	}

	content.WriteString(fmt.Sprintf("Ollama 狀態: %s\n", ollamaStatus))
	content.WriteString(fmt.Sprintf("終端管理器: 🟢 就緒\n"))
	content.WriteString(fmt.Sprintf("模板系統: 🟢 已載入 %d 個模板\n",
		len(m.templateManager.GetAvailableTemplates())))

	content.WriteString("\n系統信息:\n")
	content.WriteString(fmt.Sprintf("• 運行時間: %v\n",
		time.Since(m.lastUpdate).Truncate(time.Second)))
	content.WriteString(fmt.Sprintf("• 活動終端: %d\n", len(m.terminals)))

	return content.String()
}

// renderStatusBar 渲染狀態欄
func (m Model) renderStatusBar() string {
	return fmt.Sprintf("狀態: %s | 更新時間: %s",
		m.status, m.lastUpdate.Format("15:04:05"))
}

// renderHelp 渲染幫助信息
func (m Model) renderHelp() string {
	return "按鍵: ↑/k ↓/j 選擇 | Enter 確認 | 1-3 切換視圖 | q/Ctrl+C 退出"
}

// updateStatus 更新狀態
func (m *Model) updateStatus() {
	m.status = "系統運行中"
}

// updateTerminalStatus 更新終端狀態
func (m *Model) updateTerminalStatus(name, status string) {
	for i := range m.terminals {
		if m.terminals[i].Name == name {
			m.terminals[i].Status = status
			break
		}
	}
}

// handleQueryResult 處理查詢結果
func (m *Model) handleQueryResult(msg QueryResultMsg) {
	if msg.Success {
		m.status = fmt.Sprintf("查詢完成: %s", msg.TemplateID)
	} else {
		m.status = "查詢失敗"
	}
}

// handleError 處理錯誤
func (m *Model) handleError(err string) {
	m.status = fmt.Sprintf("錯誤: %s", err)
}

// startTerminal 啟動終端
func (m *Model) startTerminal(name, termType string) error {
	config := terminal.TerminalConfig{
		Type: m.getTerminalType(termType),
		Name: name,
	}

	return m.terminalManager.StartTerminal(config)
}

// stopTerminal 停止終端
func (m *Model) stopTerminal(name string) error {
	return m.terminalManager.StopTerminal(name)
}

// listTerminals 列出終端
func (m *Model) listTerminals() []TerminalInfo {
	terminals := m.terminalManager.ListTerminals()
	result := make([]TerminalInfo, len(terminals))

	for i, term := range terminals {
		result[i] = TerminalInfo{
			Name:     term.Name,
			Type:     term.Type.String(),
			Status:   term.GetStatus().String(),
			LastUsed: time.Unix(term.LastUsed, 0),
		}
	}

	return result
}

// optimizeQuery 優化查詢
func (m *Model) optimizeQuery(query, templateID string) (*template.ApplyResult, error) {
	context := make(map[string]string)
	context["language"] = "Go"
	context["complexity"] = "中等"

	return m.templateManager.ApplyTemplate(templateID, query, context)
}

// getTerminalType 獲取終端類型
func (m *Model) getTerminalType(typeStr string) terminal.TerminalType {
	switch typeStr {
	case "claude":
		return terminal.TypeClaudeCode
	case "gemini":
		return terminal.TypeGeminiCLI
	case "cursor":
		return terminal.TypeCursor
	case "aider":
		return terminal.TypeAider
	default:
		return terminal.TypeCustom
	}
}

// updateTerminalsCmd 更新終端列表命令
func (m *Model) updateTerminalsCmd() tea.Cmd {
	return func() tea.Msg {
		// 更新終端列表
		_ = m.listTerminals()

		// 創建一個模擬的終端狀態更新
		return TerminalStatusMsg{
			Name:   "system",
			Status: "updated",
		}
	}
}

// tickCmd 創建定時器命令
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg{Time: t}
	})
}

// Getter 方法用於測試和集成

// GetCurrentView 獲取當前視圖
func (m Model) GetCurrentView() ViewType {
	return m.currentView
}

// GetTerminalManager 獲取終端管理器
func (m Model) GetTerminalManager() *terminal.TerminalManager {
	return m.terminalManager
}

// GetOllamaClient 獲取 Ollama 客戶端
func (m Model) GetOllamaClient() *ollama.OllamaClient {
	return m.ollamaClient
}

// GetTemplateManager 獲取模板管理器
func (m Model) GetTemplateManager() *template.TemplateManager {
	return m.templateManager
}

// SwitchView 公開的視圖切換方法
func (m *Model) SwitchView(view ViewType) {
	m.switchView(view)
}

// OptimizeQuery 公開的查詢優化方法
func (m *Model) OptimizeQuery(query, templateID string) (*template.ApplyResult, error) {
	return m.optimizeQuery(query, templateID)
}

// ListTerminals 公開的終端列表方法
func (m *Model) ListTerminals() []TerminalInfo {
	return m.listTerminals()
}