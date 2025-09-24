package gui

import (
    "fmt"
    "log"
    "runtime"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"

    "ai-launcher/internal/project"
    "ai-launcher/internal/terminal"
)

type MainWindow struct {
    // Fyne 应用与窗口
    fyneApp fyne.App
    window  fyne.Window

    // 核心管理器
    projectManager  *project.ConfigManager
    terminalManager *terminal.TerminalManager

    // 主要 UI 组件
    menuBar      *fyne.MainMenu
    toolbar      *widget.Toolbar
    projectPanel *ProjectHistoryPanel
    terminalTabs *TerminalTabContainer
    statusBar    *StatusBar

    // 对话框组件
    projectDialog  *ProjectConfigDialog
    settingsDialog *SettingsDialog
    newTermDialog  *NewTerminalDialog

    // 窗口状态
    windowState *WindowState
}

// WindowState 保存窗口与主题等状态
type WindowState struct {
    Width          float32 `json:"width"`
    Height         float32 `json:"height"`
    X              float32 `json:"x"`
    Y              float32 `json:"y"`
    Maximized      bool    `json:"maximized"`
    Theme          string  `json:"theme"`
    LeftPanelWidth float32 `json:"left_panel_width"`
}

func NewMainWindow() *MainWindow {
    myApp := app.NewWithID("ai.launcher.desktop")
    myApp.SetIcon(theme.ComputerIcon())

    // Windows 下应用 CJK 字体与主题，避免中文显示为方框/乱码
    if runtime.GOOS == "windows" {
        EnsureCJKFont()
        if fp := SelectCJKFont(); fp != "" {
            _ = ApplyCJKTheme(myApp, fp)
        }
    }

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

func (mw *MainWindow) Run() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("GUI运行时错误: %v", r)
            panic(r)
        }
    }()

    if err := mw.projectManager.LoadProjects(); err != nil {
        log.Printf("加载项目配置失败: %v", err)
    }

    log.Println("创建主窗口...")
    mw.window = mw.fyneApp.NewWindow("AI 启动器 v2.0 - Desktop GUI 版")
    if mw.window == nil {
        panic("无法创建 Fyne 窗口")
    }

    mw.window.SetIcon(theme.ComputerIcon())
    mw.window.Resize(fyne.NewSize(mw.windowState.Width, mw.windowState.Height))
    mw.window.SetFixedSize(false)
    mw.window.CenterOnScreen()

    log.Println("窗口创建成功，设置属性...")
    mw.window.SetCloseIntercept(func() {
        mw.saveWindowState()
        mw.fyneApp.Quit()
    })

    if mw.windowState.Theme == "dark" {
        mw.fyneApp.Settings().SetTheme(theme.DarkTheme())
    } else {
        mw.fyneApp.Settings().SetTheme(theme.LightTheme())
    }

    log.Println("初始化 UI 组件...")
    mw.initializeComponents()

    log.Println("创建主布局...")
    content := mw.createMainLayout()
    mw.window.SetContent(content)

    log.Println("设置菜单...")
    mw.window.SetMainMenu(mw.createMenuBar())

    log.Println("显示窗口并开始事件循环...")
    mw.window.ShowAndRun()
}

func (mw *MainWindow) initializeComponents() {
    mw.toolbar = mw.createToolbar()
    mw.projectPanel = NewProjectHistoryPanel(mw.projectManager, mw.onProjectSelected)
    mw.terminalTabs = NewTerminalTabContainer(mw.terminalManager)
    mw.statusBar = NewStatusBar()
    mw.projectDialog = NewProjectConfigDialog(mw.window, mw.projectManager, mw.onProjectConfigured)
    mw.settingsDialog = NewSettingsDialog(mw.window, mw.onSettingsChanged)
    mw.newTermDialog = NewNewTerminalDialog(mw.window, mw.projectManager, mw.onNewTerminalRequested)
}

func (mw *MainWindow) createMainLayout() *fyne.Container {
    leftPanel := container.NewBorder(nil, nil, nil, nil, mw.projectPanel.GetContainer())
    leftPanel.Resize(fyne.NewSize(mw.windowState.LeftPanelWidth, 0))

    rightContent := container.NewBorder(
        mw.terminalTabs.GetTabHeader(),
        mw.statusBar.GetContainer(),
        nil, nil,
        mw.terminalTabs.GetContent(),
    )

    mainLayout := container.NewBorder(
        mw.toolbar,
        nil,
        leftPanel,
        nil,
        rightContent,
    )

    return mainLayout
}

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

func (mw *MainWindow) createMenuBar() *fyne.MainMenu {
    fileMenu := fyne.NewMenu("文件",
        fyne.NewMenuItem("打开项目", mw.onOpenProjectClicked),
        fyne.NewMenuItem("新建终端", mw.onNewTerminalClicked),
        fyne.NewMenuItemSeparator(),
        fyne.NewMenuItem("退出", func() { mw.fyneApp.Quit() }),
    )

    toolsMenu := fyne.NewMenu("工具",
        fyne.NewMenuItem("设置", mw.onSettingsClicked),
        fyne.NewMenuItem("监控", mw.onMonitorClicked),
        fyne.NewMenuItemSeparator(),
        fyne.NewMenuItem("清理缓存", mw.onClearCacheClicked),
    )

    helpMenu := fyne.NewMenu("帮助",
        fyne.NewMenuItem("使用说明", mw.onHelpClicked),
        fyne.NewMenuItem("关于", mw.onAboutClicked),
    )

    return fyne.NewMainMenu(fileMenu, toolsMenu, helpMenu)
}

// 事件处理

func (mw *MainWindow) onProjectSelected(proj project.ProjectConfig) {
    activeTab := mw.terminalTabs.GetActiveTab()
    if activeTab != nil {
        activeTab.SwitchProject(proj)
        mw.statusBar.SetMessage(fmt.Sprintf("已切换项目: %s", proj.Name))
    } else {
        mw.createNewTerminal(proj, proj.AIModel)
    }
}

func (mw *MainWindow) onProjectConfigured(proj project.ProjectConfig, aiModel project.AIModelType) {
    mw.createNewTerminal(proj, aiModel)
    mw.projectPanel.Refresh()
}

func (mw *MainWindow) onNewTerminalRequested(proj project.ProjectConfig, aiModel project.AIModelType, runInBackground bool) {
    tab := mw.createNewTerminal(proj, aiModel)
    if !runInBackground && tab != nil {
        mw.terminalTabs.SetActiveTab(tab.GetID())
    }
}

func (mw *MainWindow) onSettingsChanged(settings map[string]interface{}) {
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

func (mw *MainWindow) onOpenProjectClicked() { mw.projectDialog.Show() }
func (mw *MainWindow) onSettingsClicked()    { mw.settingsDialog.Show() }

func (mw *MainWindow) onMonitorClicked() {
    mw.statusBar.SetMessage("监控功能开发中...")
}

func (mw *MainWindow) onHelpClicked() {
    mw.statusBar.SetMessage("帮助功能开发中...")
}

func (mw *MainWindow) onNewTerminalClicked() { mw.newTermDialog.Show() }

func (mw *MainWindow) onClearCacheClicked() {
    mw.statusBar.SetMessage("缓存已清理")
}

func (mw *MainWindow) onAboutClicked() {
    mw.statusBar.SetMessage("AI 启动器 v2.0.0")
}

// 终端创建
func (mw *MainWindow) createNewTerminal(proj project.ProjectConfig, aiModel project.AIModelType) *TerminalTab {
    termName := fmt.Sprintf("%s(%s)", proj.Name, aiModel.String())
    termConfig := terminal.TerminalConfig{
        Type:       mw.getTerminalType(aiModel),
        Name:       termName,
        WorkingDir: proj.Path,
        Command:    aiModel.GetCommand(proj.YoloMode),
        YoloMode:   proj.YoloMode,
    }

    tab := mw.terminalTabs.CreateTab(termName, termConfig, proj)
    if tab != nil {
        mw.statusBar.SetMessage(fmt.Sprintf("已创建终端: %s", termName))
        return tab
    }
    mw.statusBar.SetMessage("创建终端失败")
    return nil
}

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

func (mw *MainWindow) saveWindowState() {
    log.Println("保存窗口状态...")
}

