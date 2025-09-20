package gui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"ai-launcher/internal/project"
	"ai-launcher/internal/terminal"
)

// MainWindow 主窗口结构
type MainWindow struct {
	// Fyne应用和窗口
	fyneApp fyne.App
	window  fyne.Window

	// 核心管理器
	projectManager  *project.ConfigManager
	terminalManager *terminal.TerminalManager

	// 主要UI组件
	menuBar         *fyne.MainMenu
	toolbar         *widget.Toolbar
	projectPanel    *ProjectHistoryPanel
	terminalTabs    *TerminalTabContainer
	statusBar       *StatusBar

	// 弹窗组件
	projectDialog  *ProjectConfigDialog
	settingsDialog *SettingsDialog
	newTermDialog  *NewTerminalDialog

	// 窗口状态
	windowState *WindowState
}

// WindowState 窗口状态管理
type WindowState struct {
	Width      float32 `json:"width"`
	Height     float32 `json:"height"`
	X          float32 `json:"x"`
	Y          float32 `json:"y"`
	Maximized  bool    `json:"maximized"`
	Theme      string  `json:"theme"`
	LeftPanelWidth float32 `json:"left_panel_width"`
}

// NewMainWindow 创建主窗口
func NewMainWindow() *MainWindow {
	myApp := app.NewWithID("ai.launcher.desktop")
	myApp.SetIcon(theme.ComputerIcon()) // 使用内置图标，避免资源依赖问题

	// 設置應用元數據 - 在Fyne v2.4.5中需要通過不同方式設置
	// myApp.Metadata().Name 在此版本中不可直接賦值

	return &MainWindow{
		fyneApp:         myApp,
		projectManager:  project.NewConfigManager(),
		terminalManager: terminal.NewTerminalManager(),
		windowState: &WindowState{
			Width:          1200,
			Height:         800,
			Theme:          "dark",
			LeftPanelWidth: 250,
		},
	}
}

// Run 启动主窗口
func (mw *MainWindow) Run() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("GUI运行时错误: %v", r)
			panic(r) // 重新抛出，让main函数处理
		}
	}()

	// 加载配置
	if err := mw.projectManager.LoadProjects(); err != nil {
		log.Printf("加载项目配置失败: %v", err)
	}

	log.Println("创建主窗口...")
	// 创建主窗口
	mw.window = mw.fyneApp.NewWindow("AI启动器 v2.0 - Desktop GUI版本")
	if mw.window == nil {
		panic("无法创建Fyne窗口")
	}

	mw.window.SetIcon(theme.ComputerIcon())
	mw.window.Resize(fyne.NewSize(mw.windowState.Width, mw.windowState.Height))
	mw.window.SetFixedSize(false)
	mw.window.CenterOnScreen()

	log.Println("窗口创建成功，设置属性...")

	// 设置窗口关闭行为
	mw.window.SetCloseIntercept(func() {
		mw.saveWindowState()
		mw.fyneApp.Quit()
	})

	// 应用主题
	if mw.windowState.Theme == "dark" {
		mw.fyneApp.Settings().SetTheme(theme.DarkTheme())
	} else {
		mw.fyneApp.Settings().SetTheme(theme.LightTheme())
	}

	log.Println("初始化UI组件...")
	// 初始化UI组件
	mw.initializeComponents()

	log.Println("创建主布局...")
	// 创建主布局
	content := mw.createMainLayout()
	mw.window.SetContent(content)

	log.Println("设置菜单...")
	// 设置菜单
	mw.window.SetMainMenu(mw.createMenuBar())

	log.Println("显示窗口并开始事件循环...")
	// 显示窗口
	mw.window.ShowAndRun()
}

// initializeComponents 初始化UI组件
func (mw *MainWindow) initializeComponents() {
	// 创建工具栏
	mw.toolbar = mw.createToolbar()

	// 创建左侧项目历史面板
	mw.projectPanel = NewProjectHistoryPanel(mw.projectManager, mw.onProjectSelected)

	// 创建终端标签页容器
	mw.terminalTabs = NewTerminalTabContainer(mw.terminalManager)

	// 创建状态栏
	mw.statusBar = NewStatusBar()

	// 创建弹窗组件（延迟初始化）
	mw.projectDialog = NewProjectConfigDialog(mw.window, mw.projectManager, mw.onProjectConfigured)
	mw.settingsDialog = NewSettingsDialog(mw.window, mw.onSettingsChanged)
	mw.newTermDialog = NewNewTerminalDialog(mw.window, mw.projectManager, mw.onNewTerminalRequested)
}

// createMainLayout 创建主布局
func (mw *MainWindow) createMainLayout() *fyne.Container {
	// 左侧项目历史面板（固定宽度）
	leftPanel := container.NewBorder(nil, nil, nil, nil, mw.projectPanel.GetContainer())
	leftPanel.Resize(fyne.NewSize(mw.windowState.LeftPanelWidth, 0))

	// 右侧主要内容区
	rightContent := container.NewBorder(
		mw.terminalTabs.GetTabHeader(), // 顶部：标签页头部
		mw.statusBar.GetContainer(),    // 底部：状态栏
		nil, nil,                       // 左右留空
		mw.terminalTabs.GetContent(),   // 中央：终端内容
	)

	// 主布局：工具栏 + 左右分栏内容
	mainLayout := container.NewBorder(
		mw.toolbar, // 顶部：工具栏
		nil,        // 底部留空
		leftPanel,  // 左侧：项目面板
		nil,        // 右侧留空
		rightContent, // 中央：主要内容
	)

	return mainLayout
}

// createToolbar 创建工具栏
func (mw *MainWindow) createToolbar() *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderOpenIcon(), mw.onOpenProjectClicked),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.SettingsIcon(), mw.onSettingsClicked),
		widget.NewToolbarAction(theme.InfoIcon(), mw.onMonitorClicked),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.HelpIcon(), mw.onHelpClicked),
	)

	return toolbar
}

// createMenuBar 创建菜单栏
func (mw *MainWindow) createMenuBar() *fyne.MainMenu {
	// 文件菜单
	fileMenu := fyne.NewMenu("文件",
		fyne.NewMenuItem("开启项目", mw.onOpenProjectClicked),
		fyne.NewMenuItem("新建终端", mw.onNewTerminalClicked),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("退出", func() { mw.fyneApp.Quit() }),
	)

	// 工具菜单
	toolsMenu := fyne.NewMenu("工具",
		fyne.NewMenuItem("设置", mw.onSettingsClicked),
		fyne.NewMenuItem("监控", mw.onMonitorClicked),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("清理缓存", mw.onClearCacheClicked),
	)

	// 帮助菜单
	helpMenu := fyne.NewMenu("帮助",
		fyne.NewMenuItem("使用说明", mw.onHelpClicked),
		fyne.NewMenuItem("关于", mw.onAboutClicked),
	)

	return fyne.NewMainMenu(fileMenu, toolsMenu, helpMenu)
}

// 事件处理方法

// onProjectSelected 项目选择事件
func (mw *MainWindow) onProjectSelected(project project.ProjectConfig) {
	// 在当前活动终端中切换到选中的项目
	activeTab := mw.terminalTabs.GetActiveTab()
	if activeTab != nil {
		activeTab.SwitchProject(project)
		mw.statusBar.SetMessage(fmt.Sprintf("已切换到项目: %s", project.Name))
	} else {
		// 如果没有活动终端，创建新的终端
		mw.createNewTerminal(project, project.AIModel)
	}
}

// onProjectConfigured 项目配置完成事件
func (mw *MainWindow) onProjectConfigured(project project.ProjectConfig, aiModel project.AIModelType) {
	// 启动新的终端
	mw.createNewTerminal(project, aiModel)
	mw.projectPanel.Refresh() // 刷新项目列表
}

// onNewTerminalRequested 新建终端请求事件
func (mw *MainWindow) onNewTerminalRequested(project project.ProjectConfig, aiModel project.AIModelType, runInBackground bool) {
	tab := mw.createNewTerminal(project, aiModel)
	if !runInBackground && tab != nil {
		mw.terminalTabs.SetActiveTab(tab.GetID())
	}
}

// onSettingsChanged 设置变更事件
func (mw *MainWindow) onSettingsChanged(settings map[string]interface{}) {
	// 应用设置变更
	if themeChoice, ok := settings["theme"].(string); ok {
		mw.windowState.Theme = themeChoice
		if themeChoice == "dark" {
			mw.fyneApp.Settings().SetTheme(theme.DarkTheme())
		} else {
			mw.fyneApp.Settings().SetTheme(theme.LightTheme())
		}
	}
	mw.statusBar.SetMessage("设置已应用")
}

// 工具栏事件处理

func (mw *MainWindow) onOpenProjectClicked() {
	mw.projectDialog.Show()
}

func (mw *MainWindow) onSettingsClicked() {
	mw.settingsDialog.Show()
}

func (mw *MainWindow) onMonitorClicked() {
	// TODO: 实现监控功能
	mw.statusBar.SetMessage("监控功能开发中...")
}

func (mw *MainWindow) onHelpClicked() {
	// TODO: 实现帮助功能
	mw.statusBar.SetMessage("帮助功能开发中...")
}

func (mw *MainWindow) onNewTerminalClicked() {
	mw.newTermDialog.Show()
}

func (mw *MainWindow) onClearCacheClicked() {
	// TODO: 实现缓存清理
	mw.statusBar.SetMessage("缓存已清理")
}

func (mw *MainWindow) onAboutClicked() {
	// TODO: 显示关于对话框
	mw.statusBar.SetMessage("AI启动器 v2.0.0")
}

// 工具方法

// createNewTerminal 创建新终端
func (mw *MainWindow) createNewTerminal(project project.ProjectConfig, aiModel project.AIModelType) *TerminalTab {
	// 生成终端名称：项目名(AI工具)
	termName := fmt.Sprintf("%s(%s)", project.Name, aiModel.String())

	// 创建终端配置
	termConfig := terminal.TerminalConfig{
		Type:       mw.getTerminalType(aiModel),
		Name:       termName,
		WorkingDir: project.Path,
		Command:    aiModel.GetCommand(project.YoloMode),
		YoloMode:   project.YoloMode,
	}

	// 创建终端标签页
	tab := mw.terminalTabs.CreateTab(termName, termConfig, project)
	if tab != nil {
		mw.statusBar.SetMessage(fmt.Sprintf("已创建终端: %s", termName))
		return tab
	}

	mw.statusBar.SetMessage("创建终端失败")
	return nil
}

// getTerminalType 获取终端类型
func (mw *MainWindow) getTerminalType(model project.AIModelType) terminal.TerminalType {
	switch model {
	case project.ModelClaudeCode:
		return terminal.TypeClaudeCode
	case project.ModelGeminiCLI:
		return terminal.TypeGeminiCLI
	case project.ModelCodex:
		return terminal.TypeCustom
	default:
		return terminal.TypeCustom
	}
}

// saveWindowState 保存窗口状态
func (mw *MainWindow) saveWindowState() {
	// TODO: 实现窗口状态持久化
	log.Println("保存窗口状态...")
}