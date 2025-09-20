package gui

import (
	"fmt"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"ai-launcher/internal/project"
	"ai-launcher/internal/terminal"
)

// TerminalTabContainer 终端标签页容器
type TerminalTabContainer struct {
	terminalManager *terminal.TerminalManager

	// UI组件
	tabContainer *container.AppTabs
	tabHeader    *fyne.Container
	content      *fyne.Container

	// 标签页管理
	tabs         map[string]*TerminalTab
	activeTabID  string
	nextTabID    int
	mutex        sync.RWMutex
}

// TerminalTab 终端标签页
type TerminalTab struct {
	id           string
	name         string
	terminalType terminal.TerminalType
	project      project.ProjectConfig

	// UI组件
	content      *fyne.Container
	terminal     *widget.Entry // 简化版本，后续可以替换为真正的终端组件
	outputArea   *widget.RichText
	inputArea    *widget.Entry
	statusLabel  *widget.Label

	// 状态
	active       bool
	running      bool
}

// NewTerminalTabContainer 创建终端标签页容器
func NewTerminalTabContainer(tm *terminal.TerminalManager) *TerminalTabContainer {
	container := &TerminalTabContainer{
		terminalManager: tm,
		tabs:           make(map[string]*TerminalTab),
		nextTabID:      1,
	}

	container.initializeUI()
	return container
}

// initializeUI 初始化UI
func (tc *TerminalTabContainer) initializeUI() {
	// 创建标签页容器
	tc.tabContainer = container.NewAppTabs()

	// 创建"新建"标签页
	tc.addNewTabButton()

	// 标签页头部容器
	tc.tabHeader = container.NewBorder(nil, nil, nil, nil, tc.tabContainer)

	// 内容区域（显示当前活动标签页的内容）
	tc.content = container.NewMax()

	// 设置标签页切换事件
	tc.tabContainer.OnChanged = tc.onTabChanged
}

// addNewTabButton 添加"新建"标签页按钮
func (tc *TerminalTabContainer) addNewTabButton() {
	newTabContent := container.NewCenter(
		widget.NewButtonWithIcon("新建终端", theme.ContentAddIcon(), func() {
			// TODO: 触发新建终端对话框
		}),
	)

	newTab := &container.TabItem{
		Text: "➕ 新建",
		Icon: theme.ContentAddIcon(),
		Content: newTabContent,
	}
	tc.tabContainer.Append(newTab)
}

// CreateTab 创建新的终端标签页
func (tc *TerminalTabContainer) CreateTab(name string, termConfig terminal.TerminalConfig, proj project.ProjectConfig) *TerminalTab {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	// 生成唯一ID
	tabID := fmt.Sprintf("tab_%d", tc.nextTabID)
	tc.nextTabID++

	// 创建终端标签页
	tab := &TerminalTab{
		id:           tabID,
		name:         name,
		terminalType: termConfig.Type,
		project:      proj,
	}

	// 初始化标签页UI
	tab.initializeUI()

	// 启动终端
	go tab.startTerminal(termConfig)

	// 添加到标签页容器（在"新建"标签页之前）
	tabContent := tab.GetContent()
	appTab := &container.TabItem{
		Text: name,
		Content: tabContent,
	}

	// 插入到最后一个位置（"新建"标签页之前）
	tc.tabContainer.Append(appTab)

	// 保存标签页引用
	tc.tabs[tabID] = tab

	// 设置为活动标签页
	tc.SetActiveTab(tabID)

	return tab
}

// RemoveTab 移除标签页
func (tc *TerminalTabContainer) RemoveTab(tabID string) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	tab, exists := tc.tabs[tabID]
	if !exists {
		return
	}

	// 停止终端
	tab.stopTerminal()

	// 从标签页容器中移除
	for i := 0; i < len(tc.tabContainer.Items)-1; i++ { // 排除"新建"标签页
		if tc.tabContainer.Items[i].Text == tab.name {
			tc.tabContainer.RemoveIndex(i)
			break
		}
	}

	// 从映射中删除
	delete(tc.tabs, tabID)

	// 如果删除的是活动标签页，切换到下一个
	if tc.activeTabID == tabID {
		tc.activateNextTab()
	}
}

// SetActiveTab 设置活动标签页
func (tc *TerminalTabContainer) SetActiveTab(tabID string) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	tab, exists := tc.tabs[tabID]
	if !exists {
		return
	}

	// 取消之前的活动状态
	if tc.activeTabID != "" {
		if prevTab := tc.tabs[tc.activeTabID]; prevTab != nil {
			prevTab.active = false
		}
	}

	// 设置新的活动标签页
	tc.activeTabID = tabID
	tab.active = true

	// 更新内容区域
	tc.content.Objects = []fyne.CanvasObject{tab.GetContent()}
	tc.content.Refresh()

	// 切换标签页容器的选中状态
	for i := 0; i < len(tc.tabContainer.Items)-1; i++ {
		if tc.tabContainer.Items[i].Text == tab.name {
			tc.tabContainer.SelectTabIndex(i)
			break
		}
	}
}

// GetActiveTab 获取当前活动的标签页
func (tc *TerminalTabContainer) GetActiveTab() *TerminalTab {
	tc.mutex.RLock()
	defer tc.mutex.RUnlock()

	if tc.activeTabID == "" {
		return nil
	}

	return tc.tabs[tc.activeTabID]
}

// GetTabHeader 获取标签页头部容器
func (tc *TerminalTabContainer) GetTabHeader() *fyne.Container {
	return tc.tabHeader
}

// GetContent 获取内容区域容器
func (tc *TerminalTabContainer) GetContent() *fyne.Container {
	return tc.content
}

// onTabChanged 标签页切换事件处理
func (tc *TerminalTabContainer) onTabChanged(tab *container.TabItem) {
	// 查找对应的终端标签页
	for tabID, termTab := range tc.tabs {
		if termTab.name == tab.Text {
			tc.SetActiveTab(tabID)
			return
		}
	}
}

// activateNextTab 激活下一个标签页
func (tc *TerminalTabContainer) activateNextTab() {
	if len(tc.tabs) == 0 {
		tc.activeTabID = ""
		tc.content.Objects = []fyne.CanvasObject{}
		tc.content.Refresh()
		return
	}

	// 激活第一个可用的标签页
	for tabID := range tc.tabs {
		tc.SetActiveTab(tabID)
		return
	}
}

// TerminalTab 方法实现

// initializeUI 初始化终端标签页UI
func (tab *TerminalTab) initializeUI() {
	// 输出区域（只读）
	tab.outputArea = widget.NewRichText()
	tab.outputArea.Wrapping = fyne.TextWrapWord
	tab.outputArea.Scroll = container.ScrollBoth

	// 输入区域
	tab.inputArea = widget.NewEntry()
	tab.inputArea.SetPlaceHolder("输入命令或与AI交互...")
	tab.inputArea.OnSubmitted = tab.onInputSubmitted

	// 状态标签
	tab.statusLabel = widget.NewLabel("准备就绪")

	// 工具栏
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.MediaPlayIcon(), tab.onStartTerminal),
		widget.NewToolbarAction(theme.MediaStopIcon(), tab.onStopTerminal),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), tab.onClearOutput),
		widget.NewToolbarAction(theme.SettingsIcon(), tab.onTerminalSettings),
	)

	// 底部状态栏
	statusBar := container.NewBorder(nil, nil, tab.statusLabel, nil, tab.statusLabel)

	// 主要内容布局
	tab.content = container.NewBorder(
		toolbar,           // 顶部：工具栏
		container.NewVBox( // 底部：输入区域和状态栏
			tab.inputArea,
			statusBar,
		),
		nil, nil,          // 左右留空
		container.NewScroll(tab.outputArea), // 中央：输出区域
	)
}

// startTerminal 启动终端
func (tab *TerminalTab) startTerminal(config terminal.TerminalConfig) {
	tab.running = true
	tab.statusLabel.SetText("运行中...")

	// TODO: 集成真正的终端管理器
	tab.appendOutput(fmt.Sprintf("🚀 正在启动 %s...\n", config.Type))
	tab.appendOutput(fmt.Sprintf("📁 工作目录: %s\n", config.WorkingDir))
	tab.appendOutput(fmt.Sprintf("⚡ 模式: %s\n", map[bool]string{true: "YOLO", false: "普通"}[config.YoloMode]))
	tab.appendOutput("✅ 终端已启动，准备接收命令\n\n")
}

// stopTerminal 停止终端
func (tab *TerminalTab) stopTerminal() {
	tab.running = false
	tab.statusLabel.SetText("已停止")
	tab.appendOutput("\n🛑 终端已停止\n")
}

// GetContent 获取标签页内容
func (tab *TerminalTab) GetContent() *fyne.Container {
	return tab.content
}

// GetID 获取标签页ID
func (tab *TerminalTab) GetID() string {
	return tab.id
}

// SwitchProject 切换项目
func (tab *TerminalTab) SwitchProject(proj project.ProjectConfig) {
	tab.project = proj
	tab.appendOutput(fmt.Sprintf("\n📁 已切换到项目: %s\n", proj.Name))
	tab.appendOutput(fmt.Sprintf("📍 路径: %s\n\n", proj.Path))
}

// appendOutput 追加输出内容
func (tab *TerminalTab) appendOutput(text string) {
	currentText := tab.outputArea.String() + text
	tab.outputArea.ParseMarkdown(currentText)
	tab.outputArea.Refresh()
}

// 事件处理方法

func (tab *TerminalTab) onInputSubmitted(input string) {
	if input == "" {
		return
	}

	// 显示用户输入
	tab.appendOutput(fmt.Sprintf("> %s\n", input))

	// TODO: 发送到终端管理器处理
	tab.appendOutput("📝 命令已接收，正在处理...\n")

	// 清空输入
	tab.inputArea.SetText("")
}

func (tab *TerminalTab) onStartTerminal() {
	if !tab.running {
		// TODO: 重新启动终端
		tab.appendOutput("🔄 重新启动终端...\n")
	}
}

func (tab *TerminalTab) onStopTerminal() {
	if tab.running {
		tab.stopTerminal()
	}
}

func (tab *TerminalTab) onClearOutput() {
	tab.outputArea.ParseMarkdown("")
	tab.outputArea.Refresh()
	tab.appendOutput("📄 输出已清空\n\n")
}

func (tab *TerminalTab) onTerminalSettings() {
	// TODO: 显示终端设置
	tab.appendOutput("⚙️ 终端设置功能开发中...\n")
}