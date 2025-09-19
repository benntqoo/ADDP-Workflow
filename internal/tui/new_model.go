package tui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"ai-launcher/internal/project"
	"ai-launcher/internal/terminal"
)

// NewUIModel 新的优化UI模型
type NewUIModel struct {
	// 核心组件
	configManager   *project.ConfigManager
	terminalManager *terminal.TerminalManager

	// UI状态
	currentStep  UIStep
	width        int
	height       int
	selectedIndex int

	// 配置数据
	projectPath     string
	projectName     string
	selectedModel   project.AIModelType
	yoloMode        bool
	recentProjects  []project.ProjectConfig
	isPathValid     bool

	// 状态信息
	status          string
	errorMessage    string
	isLaunching     bool

	// 样式
	styles          Styles
}

// UIStep UI步骤
type UIStep int

const (
	StepWelcome     UIStep = iota  // 欢迎界面
	StepProjectPath                // 项目路径选择
	StepModelSelect                // 模型选择
	StepYoloMode                   // YOLO模式选择
	StepConfirm                    // 确认配置
	StepLaunching                  // 启动中
	StepComplete                   // 完成
)

// String 返回步骤的字符串表示
func (s UIStep) String() string {
	switch s {
	case StepWelcome:
		return "欢迎"
	case StepProjectPath:
		return "项目路径"
	case StepModelSelect:
		return "模型选择"
	case StepYoloMode:
		return "YOLO模式"
	case StepConfirm:
		return "确认配置"
	case StepLaunching:
		return "启动中"
	case StepComplete:
		return "完成"
	default:
		return "未知"
	}
}

// NewNewUIModel 创建新的UI模型
func NewNewUIModel() *NewUIModel {
	configManager := project.NewConfigManager()

	// 加载项目配置
	if err := configManager.LoadProjects(); err != nil {
		// 如果加载失败，继续使用空配置
	}

	return &NewUIModel{
		configManager:   configManager,
		terminalManager: terminal.NewTerminalManager(),
		currentStep:     StepWelcome,
		selectedModel:   project.ModelClaudeCode,
		yoloMode:        false,
		recentProjects:  configManager.GetRecentProjects(5),
		styles:          *NewStyles(),
		status:          "准备开始",
	}
}

// Init 初始化模型
func (m *NewUIModel) Init() tea.Cmd {
	return nil
}

// Update 更新模型
func (m *NewUIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case LaunchResultMsg:
		return m.handleLaunchResult(msg)

	default:
		return m, nil
	}
}

// View 渲染视图
func (m *NewUIModel) View() string {
	var content strings.Builder

	// 标题栏
	content.WriteString(m.renderHeader())
	content.WriteString("\n\n")

	// 主内容区域
	switch m.currentStep {
	case StepWelcome:
		content.WriteString(m.renderWelcomeView())
	case StepProjectPath:
		content.WriteString(m.renderProjectPathView())
	case StepModelSelect:
		content.WriteString(m.renderModelSelectView())
	case StepYoloMode:
		content.WriteString(m.renderYoloModeView())
	case StepConfirm:
		content.WriteString(m.renderConfirmView())
	case StepLaunching:
		content.WriteString(m.renderLaunchingView())
	case StepComplete:
		content.WriteString(m.renderCompleteView())
	}

	content.WriteString("\n\n")

	// 状态栏
	content.WriteString(m.renderStatusBar())

	// 帮助信息
	content.WriteString("\n")
	content.WriteString(m.renderHelp())

	return content.String()
}

// handleKeyMsg 处理按键消息
func (m *NewUIModel) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		if m.currentStep == StepLaunching {
			return m, nil // 启动中不允许退出
		}
		return m, tea.Quit

	case "enter":
		return m.handleEnterKey()

	case "up", "k":
		if m.selectedIndex > 0 {
			m.selectedIndex--
		}
		return m, nil

	case "down", "j":
		maxIndex := m.getMaxIndex()
		if m.selectedIndex < maxIndex {
			m.selectedIndex++
		}
		return m, nil

	case "esc":
		return m.handleBackStep()

	case "tab":
		if m.currentStep == StepYoloMode {
			m.yoloMode = !m.yoloMode
			m.status = fmt.Sprintf("YOLO模式: %v", m.yoloMode)
		}
		return m, nil

	default:
		// 在项目路径输入时处理字符输入
		if m.currentStep == StepProjectPath && len(msg.String()) == 1 {
			return m.handlePathInput(msg.String())
		}
		return m, nil
	}
}

// handleEnterKey 处理回车键
func (m *NewUIModel) handleEnterKey() (tea.Model, tea.Cmd) {
	switch m.currentStep {
	case StepWelcome:
		m.currentStep = StepProjectPath
		m.selectedIndex = 0
		m.status = "请选择或输入项目路径"
		return m, nil

	case StepProjectPath:
		return m.handleProjectPathSelection()

	case StepModelSelect:
		return m.handleModelSelection()

	case StepYoloMode:
		m.currentStep = StepConfirm
		m.selectedIndex = 0
		m.status = "请确认配置"
		return m, nil

	case StepConfirm:
		if m.selectedIndex == 0 { // 确认启动
			return m.startLaunching()
		} else { // 返回修改
			m.currentStep = StepModelSelect
			m.selectedIndex = 0
			m.status = "返回模型选择"
			return m, nil
		}

	case StepComplete:
		return m, tea.Quit

	default:
		return m, nil
	}
}

// handleBackStep 处理返回步骤
func (m *NewUIModel) handleBackStep() (tea.Model, tea.Cmd) {
	switch m.currentStep {
	case StepProjectPath:
		m.currentStep = StepWelcome
	case StepModelSelect:
		m.currentStep = StepProjectPath
	case StepYoloMode:
		m.currentStep = StepModelSelect
	case StepConfirm:
		m.currentStep = StepYoloMode
	default:
		return m, nil
	}

	m.selectedIndex = 0
	m.status = "返回上一步"
	return m, nil
}

// handleProjectPathSelection 处理项目路径选择
func (m *NewUIModel) handleProjectPathSelection() (tea.Model, tea.Cmd) {
	if len(m.recentProjects) > 0 && m.selectedIndex < len(m.recentProjects) {
		// 选择最近的项目
		selected := m.recentProjects[m.selectedIndex]
		m.projectPath = selected.Path
		m.projectName = selected.Name
		m.selectedModel = selected.AIModel
		m.yoloMode = selected.YoloMode
	} else {
		// 使用当前输入的路径
		if !m.isPathValid {
			m.errorMessage = "请输入有效的项目路径"
			return m, nil
		}
		m.projectName = filepath.Base(m.projectPath)
	}

	// 验证路径
	if err := m.configManager.ValidateProjectPath(m.projectPath); err != nil {
		m.errorMessage = err.Error()
		return m, nil
	}

	m.currentStep = StepModelSelect
	// 将selectedModel转换为索引
	models := m.configManager.GetAvailableModels()
	for i, model := range models {
		if model == m.selectedModel {
			m.selectedIndex = i
			break
		}
	}
	m.status = "请选择AI模型"
	m.errorMessage = ""
	return m, nil
}

// handleModelSelection 处理模型选择
func (m *NewUIModel) handleModelSelection() (tea.Model, tea.Cmd) {
	models := m.configManager.GetAvailableModels()
	if m.selectedIndex < len(models) {
		m.selectedModel = models[m.selectedIndex]
		m.currentStep = StepYoloMode
		m.selectedIndex = 0
		m.status = "请选择运行模式"
	}
	return m, nil
}

// startLaunching 开始启动
func (m *NewUIModel) startLaunching() (tea.Model, tea.Cmd) {
	m.currentStep = StepLaunching
	m.isLaunching = true
	m.status = "正在启动AI工具..."

	// 保存项目配置
	projectConfig := project.ProjectConfig{
		Name:     m.projectName,
		Path:     m.projectPath,
		AIModel:  m.selectedModel,
		YoloMode: m.yoloMode,
		LastUsed: time.Now(),
	}

	if err := m.configManager.AddProject(projectConfig); err != nil {
		m.errorMessage = fmt.Sprintf("保存配置失败: %v", err)
	}

	// 返回启动命令
	return m, m.launchTerminalCmd()
}

// launchTerminalCmd 启动终端命令
func (m *NewUIModel) launchTerminalCmd() tea.Cmd {
	return func() tea.Msg {
		// 创建终端配置
		config := terminal.TerminalConfig{
			Type:        m.getTerminalType(),
			Name:        fmt.Sprintf("%s-%s", m.selectedModel, m.projectName),
			WorkingDir:  m.projectPath,
			Command:     m.selectedModel.GetCommand(m.yoloMode),
			YoloMode:    m.yoloMode,
		}

		// 启动终端
		err := m.terminalManager.StartTerminal(config)

		return LaunchResultMsg{
			Success: err == nil,
			Error:   err,
			Config:  config,
		}
	}
}

// handleLaunchResult 处理启动结果
func (m *NewUIModel) handleLaunchResult(msg LaunchResultMsg) (tea.Model, tea.Cmd) {
	m.isLaunching = false

	if msg.Success {
		m.currentStep = StepComplete
		m.status = fmt.Sprintf("✅ 成功启动 %s", msg.Config.Name)
	} else {
		m.errorMessage = fmt.Sprintf("启动失败: %v", msg.Error)
		m.currentStep = StepConfirm // 返回确认页面
	}

	return m, nil
}

// handlePathInput 处理路径输入
func (m *NewUIModel) handlePathInput(char string) (tea.Model, tea.Cmd) {
	// 这里可以实现路径输入逻辑
	// 暂时使用当前工作目录作为默认值
	if m.projectPath == "" {
		if wd, err := os.Getwd(); err == nil {
			m.projectPath = wd
			m.isPathValid = true
			m.status = fmt.Sprintf("当前路径: %s", m.projectPath)
		}
	}
	return m, nil
}

// 渲染方法

// renderHeader 渲染标题
func (m *NewUIModel) renderHeader() string {
	step := fmt.Sprintf("[%d/6]", int(m.currentStep)+1)
	title := fmt.Sprintf("🚀 AI启动器 %s %s", step, m.currentStep.String())
	return m.styles.TitleStyle.Render(title)
}

// renderWelcomeView 渲染欢迎界面
func (m *NewUIModel) renderWelcomeView() string {
	var content strings.Builder

	content.WriteString(m.styles.TitleStyle.Render("欢迎使用AI启动器"))
	content.WriteString("\n\n")

	content.WriteString("✨ 功能特性:\n")
	content.WriteString("• 支持多种AI模型 (Claude Code, Gemini CLI, Codex)\n")
	content.WriteString("• 快速项目目录切换\n")
	content.WriteString("• YOLO模式支持\n")
	content.WriteString("• 配置自动保存\n\n")

	content.WriteString("🎯 使用流程:\n")
	content.WriteString("1. 选择项目目录\n")
	content.WriteString("2. 选择AI模型\n")
	content.WriteString("3. 配置运行模式\n")
	content.WriteString("4. 一键启动终端\n\n")

	content.WriteString(m.styles.SelectedItemStyle.Render("按 Enter 开始配置"))

	return content.String()
}

// renderProjectPathView 渲染项目路径选择界面
func (m *NewUIModel) renderProjectPathView() string {
	var content strings.Builder

	content.WriteString(m.styles.SubtitleStyle.Render("📁 选择项目目录"))
	content.WriteString("\n\n")

	// 显示最近使用的项目
	if len(m.recentProjects) > 0 {
		content.WriteString("最近使用的项目:\n\n")

		for i, proj := range m.recentProjects {
			style := m.styles.MenuItemStyle
			prefix := "  "
			if i == m.selectedIndex {
				style = m.styles.SelectedItemStyle
				prefix = "▶ "
			}

			lastUsed := proj.LastUsed.Format("01-02 15:04")
			line := fmt.Sprintf("%s%s %s (%s) - %s",
				prefix, proj.AIModel.GetIcon(), proj.Name, proj.AIModel.String(), lastUsed)

			content.WriteString(style.Render(line))
			content.WriteString("\n")
		}
		content.WriteString("\n")
	}

	// 当前路径输入
	content.WriteString("当前路径:\n")
	if m.projectPath != "" {
		pathStyle := m.styles.InfoStyle
		if m.isPathValid {
			pathStyle = m.styles.SuccessStyle
		} else {
			pathStyle = m.styles.ErrorStyle
		}
		content.WriteString(pathStyle.Render(fmt.Sprintf("📍 %s", m.projectPath)))
	} else {
		content.WriteString(m.styles.DescriptionStyle.Render("📍 请选择项目或按任意键使用当前目录"))
	}

	return content.String()
}

// renderModelSelectView 渲染模型选择界面
func (m *NewUIModel) renderModelSelectView() string {
	var content strings.Builder

	content.WriteString(m.styles.SubtitleStyle.Render("🤖 选择AI模型"))
	content.WriteString("\n\n")

	content.WriteString(fmt.Sprintf("项目: %s (%s)\n\n", m.projectName, m.projectPath))

	models := m.configManager.GetAvailableModels()

	for i, model := range models {
		style := m.styles.MenuItemStyle
		prefix := "  "
		if i == m.selectedIndex {
			style = m.styles.SelectedItemStyle
			prefix = "▶ "
		}

		line := fmt.Sprintf("%s%s %s", prefix, model.GetIcon(), model.String())
		content.WriteString(style.Render(line))
		content.WriteString("\n")
	}

	content.WriteString("\n")
	content.WriteString(m.styles.DescriptionStyle.Render("选择您想要使用的AI编程助手"))

	return content.String()
}

// renderYoloModeView 渲染YOLO模式选择界面
func (m *NewUIModel) renderYoloModeView() string {
	var content strings.Builder

	content.WriteString(m.styles.SubtitleStyle.Render("⚡ 运行模式配置"))
	content.WriteString("\n\n")

	content.WriteString(fmt.Sprintf("项目: %s\n", m.projectName))
	content.WriteString(fmt.Sprintf("模型: %s %s\n\n", m.selectedModel.GetIcon(), m.selectedModel.String()))

	// YOLO模式说明
	content.WriteString("🎯 运行模式:\n\n")

	// 普通模式
	normalStyle := m.styles.MenuItemStyle
	normalPrefix := "  "
	if !m.yoloMode {
		normalStyle = m.styles.SelectedItemStyle
		normalPrefix = "▶ "
	}

	content.WriteString(normalStyle.Render(fmt.Sprintf("%s🛡️  普通模式 (安全模式)", normalPrefix)))
	content.WriteString("\n")
	content.WriteString("     • 需要用户确认重要操作\n")
	content.WriteString("     • 适合生产环境和重要项目\n\n")

	// YOLO模式
	yoloStyle := m.styles.MenuItemStyle
	yoloPrefix := "  "
	if m.yoloMode {
		yoloStyle = m.styles.SelectedItemStyle
		yoloPrefix = "▶ "
	}

	content.WriteString(yoloStyle.Render(fmt.Sprintf("%s🚀 YOLO模式 (极速模式)", yoloPrefix)))
	content.WriteString("\n")
	content.WriteString("     • 跳过大部分安全检查\n")
	content.WriteString("     • 自动执行AI建议的操作\n")
	content.WriteString("     • 适合实验和快速原型\n\n")

	content.WriteString(m.styles.SelectedItemStyle.Render("按 Tab 切换模式，Enter 继续"))

	return content.String()
}

// renderConfirmView 渲染确认配置界面
func (m *NewUIModel) renderConfirmView() string {
	var content strings.Builder

	content.WriteString(m.styles.SubtitleStyle.Render("✅ 确认配置"))
	content.WriteString("\n\n")

	// 配置摘要
	content.WriteString("🔧 启动配置:\n\n")
	content.WriteString(fmt.Sprintf("📁 项目路径: %s\n", m.projectPath))
	content.WriteString(fmt.Sprintf("📝 项目名称: %s\n", m.projectName))
	content.WriteString(fmt.Sprintf("🤖 AI模型: %s %s\n", m.selectedModel.GetIcon(), m.selectedModel.String()))

	modeText := "🛡️ 普通模式"
	if m.yoloMode {
		modeText = "🚀 YOLO模式"
	}
	content.WriteString(fmt.Sprintf("⚡ 运行模式: %s\n\n", modeText))

	// 启动命令预览
	command := strings.Join(m.selectedModel.GetCommand(m.yoloMode), " ")
	content.WriteString(fmt.Sprintf("💻 启动命令: %s\n\n", command))

	// 选项
	options := []string{"🚀 确认启动", "📝 返回修改"}

	for i, option := range options {
		style := m.styles.MenuItemStyle
		prefix := "  "
		if i == m.selectedIndex {
			style = m.styles.SelectedItemStyle
			prefix = "▶ "
		}

		content.WriteString(style.Render(prefix + option))
		content.WriteString("\n")
	}

	return content.String()
}

// renderLaunchingView 渲染启动中界面
func (m *NewUIModel) renderLaunchingView() string {
	var content strings.Builder

	content.WriteString(m.styles.SubtitleStyle.Render("🚀 正在启动AI工具..."))
	content.WriteString("\n\n")

	// 启动动画效果
	dots := strings.Repeat(".", (int(time.Now().Unix())%4)+1)
	content.WriteString(fmt.Sprintf("启动中%s\n\n", dots))

	content.WriteString(fmt.Sprintf("📁 项目: %s\n", m.projectName))
	content.WriteString(fmt.Sprintf("🤖 模型: %s %s\n", m.selectedModel.GetIcon(), m.selectedModel.String()))
	content.WriteString(fmt.Sprintf("📍 路径: %s\n\n", m.projectPath))

	content.WriteString("⏳ 请稍等，正在初始化终端环境...")

	return content.String()
}

// renderCompleteView 渲染完成界面
func (m *NewUIModel) renderCompleteView() string {
	var content strings.Builder

	content.WriteString(m.styles.SubtitleStyle.Render("🎉 启动完成！"))
	content.WriteString("\n\n")

	content.WriteString("✅ AI工具已成功启动\n\n")

	content.WriteString(fmt.Sprintf("📁 项目: %s\n", m.projectName))
	content.WriteString(fmt.Sprintf("🤖 模型: %s %s\n", m.selectedModel.GetIcon(), m.selectedModel.String()))
	content.WriteString(fmt.Sprintf("📍 路径: %s\n\n", m.projectPath))

	content.WriteString("🎯 接下来您可以:\n")
	content.WriteString("• 在新开的终端窗口中开始AI协作\n")
	content.WriteString("• 使用配置的AI模型进行编程\n")
	if m.yoloMode {
		content.WriteString("• 享受YOLO模式的极速体验\n")
	}
	content.WriteString("\n")

	content.WriteString(m.styles.SelectedItemStyle.Render("按 Enter 退出启动器"))

	return content.String()
}

// renderStatusBar 渲染状态栏
func (m *NewUIModel) renderStatusBar() string {
	status := m.status
	if m.errorMessage != "" {
		status = m.styles.ErrorStyle.Render("❌ " + m.errorMessage)
	}

	return fmt.Sprintf("状态: %s", status)
}

// renderHelp 渲染帮助信息
func (m *NewUIModel) renderHelp() string {
	switch m.currentStep {
	case StepWelcome:
		return "按键: Enter 开始 | Ctrl+C 退出"
	case StepProjectPath:
		return "按键: ↑/↓ 选择 | Enter 确认 | Esc 返回 | 任意键使用当前路径"
	case StepModelSelect, StepConfirm:
		return "按键: ↑/↓ 选择 | Enter 确认 | Esc 返回"
	case StepYoloMode:
		return "按键: Tab 切换模式 | Enter 继续 | Esc 返回"
	case StepLaunching:
		return "请等待启动完成..."
	case StepComplete:
		return "按键: Enter 退出"
	default:
		return ""
	}
}

// 辅助方法

// getMaxIndex 获取当前步骤的最大索引
func (m *NewUIModel) getMaxIndex() int {
	switch m.currentStep {
	case StepProjectPath:
		return len(m.recentProjects) - 1
	case StepModelSelect:
		return len(m.configManager.GetAvailableModels()) - 1
	case StepYoloMode:
		return 1 // 两个选项
	case StepConfirm:
		return 1 // 两个选项
	default:
		return 0
	}
}

// getTerminalType 获取终端类型
func (m *NewUIModel) getTerminalType() terminal.TerminalType {
	switch m.selectedModel {
	case project.ModelClaudeCode:
		return terminal.TypeClaudeCode
	case project.ModelGeminiCLI:
		return terminal.TypeGeminiCLI
	case project.ModelCodex:
		return terminal.TypeCustom // 使用自定义类型
	default:
		return terminal.TypeCustom
	}
}

// 消息类型

// LaunchResultMsg 启动结果消息
type LaunchResultMsg struct {
	Success bool
	Error   error
	Config  terminal.TerminalConfig
}

// 测试辅助方法

// GetSelectedIndex 获取当前选择索引（用于测试）
func (m *NewUIModel) GetSelectedIndex() int {
	return m.selectedIndex
}

// GetRecentProjects 获取最近项目列表（用于测试）
func (m *NewUIModel) GetRecentProjects() []project.ProjectConfig {
	return m.recentProjects
}

// GetYoloMode 获取YOLO模式状态（用于测试）
func (m *NewUIModel) GetYoloMode() bool {
	return m.yoloMode
}

// SetProjectPath 设置项目路径（用于测试）
func (m *NewUIModel) SetProjectPath(path string) {
	m.projectPath = path
	m.isPathValid = true
}

// SetProjectName 设置项目名称（用于测试）
func (m *NewUIModel) SetProjectName(name string) {
	m.projectName = name
}

// SetSelectedModel 设置选择的模型（用于测试）
func (m *NewUIModel) SetSelectedModel(model project.AIModelType) {
	m.selectedModel = model
}

// SetYoloMode 设置YOLO模式（用于测试）
func (m *NewUIModel) SetYoloMode(yolo bool) {
	m.yoloMode = yolo
}

// SetCurrentStep 设置当前步骤（用于测试）
func (m *NewUIModel) SetCurrentStep(step UIStep) {
	m.currentStep = step
}

// SetErrorMessage 设置错误消息（用于测试）
func (m *NewUIModel) SetErrorMessage(msg string) {
	m.errorMessage = msg
}

// GetConfigManager 获取配置管理器（用于测试）
func (m *NewUIModel) GetConfigManager() *project.ConfigManager {
	return m.configManager
}

// GetTerminalManager 获取终端管理器（用于测试）
func (m *NewUIModel) GetTerminalManager() *terminal.TerminalManager {
	return m.terminalManager
}

// GetWindowSize 获取窗口大小（用于测试）
func (m *NewUIModel) GetWindowSize() (int, int) {
	return m.width, m.height
}

// RefreshRecentProjects 刷新最近项目列表（用于测试）
func (m *NewUIModel) RefreshRecentProjects() {
	m.recentProjects = m.configManager.GetRecentProjects(5)
}