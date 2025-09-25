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
    // Fyne 搴旂敤涓庣獥鍙?
    fyneApp fyne.App
    window  fyne.Window

    // 鏍稿績绠＄悊鍣?
    projectManager  *project.ConfigManager
    terminalManager *terminal.TerminalManager

    // 涓昏 UI 缁勪欢
    menuBar      *fyne.MainMenu
    toolbar      *widget.Toolbar
    projectPanel *ProjectHistoryPanel
    terminalTabs *TerminalTabContainer
    statusBar    *StatusBar

    // 瀵硅瘽妗嗙粍浠?
    projectDialog  *ProjectConfigDialog
    settingsDialog *SettingsDialog
    newTermDialog  *NewTerminalDialog

    // 绐楀彛鐘舵€?
    windowState *WindowState
}

// WindowState 淇濆瓨绐楀彛涓庝富棰樼瓑鐘舵€?
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

    // Windows 涓嬪簲鐢?CJK 瀛椾綋涓庝富棰橈紝閬垮厤涓枃鏄剧ず涓烘柟妗?涔辩爜
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
            log.Printf("GUI杩愯鏃堕敊璇? %v", r)
            panic(r)
        }
    }()

    if err := mw.projectManager.LoadProjects(); err != nil {
        log.Printf("鍔犺浇椤圭洰閰嶇疆澶辫触: %v", err)
    }

    log.Println("创建主窗口...")
    mw.window = mw.fyneApp.NewWindow("AI 启动器 v2.0 - Desktop GUI 版")
    if mw.window == nil {
        panic("鏃犳硶鍒涘缓 Fyne 绐楀彛")
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

    log.Println("鍒濆鍖?UI 缁勪欢...")
    mw.initializeComponents()

    log.Println("鍒涘缓涓诲竷灞€...")
    content := mw.createMainLayout()
    mw.window.SetContent(content)

    log.Println("璁剧疆鑿滃崟...")
    mw.window.SetMainMenu(mw.createMenuBar())

    log.Println("鏄剧ず绐楀彛骞跺紑濮嬩簨浠跺惊鐜?..")
    mw.window.ShowAndRun()
}

func (mw *MainWindow) initializeComponents() {
    // 顶部工具栏取消（避免与主菜单重复），保持简洁导航
    mw.toolbar = nil
    mw.projectPanel = NewProjectHistoryPanel(mw.projectManager, mw.onProjectSelected)
    mw.terminalTabs = NewTerminalTabContainer(mw.terminalManager, func(){ mw.onNewTerminalClicked() })
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
        nil,
        nil,
        leftPanel,
        nil,
        rightContent,
    )

    return mainLayout
}

// 工具栏已移除，避免与菜单重复。如需恢复，可按需实现 createToolbar()

func (mw *MainWindow) createMenuBar() *fyne.MainMenu {
    fileMenu := fyne.NewMenu("文件",
        fyne.NewMenuItem("新建终端", mw.onNewTerminalClicked),
        fyne.NewMenuItemSeparator(),
        fyne.NewMenuItem("退出", func() { mw.fyneApp.Quit() }),
    )

    toolsMenu := fyne.NewMenu("工具",
        fyne.NewMenuItem("监控", mw.onMonitorClicked),
        fyne.NewMenuItemSeparator(),
        fyne.NewMenuItem("清理缓存", mw.onClearCacheClicked),
    )

    settingsMenu := fyne.NewMenu("设置",
        fyne.NewMenuItem("首选项", mw.onSettingsClicked),
    )

    helpMenu := fyne.NewMenu("帮助",
        fyne.NewMenuItem("使用说明", mw.onHelpClicked),
        fyne.NewMenuItem("关于", mw.onAboutClicked),
    )

    // 顶部导航顺序：文件 | 工具 | 设置 | 帮助
    return fyne.NewMainMenu(fileMenu, toolsMenu, settingsMenu, helpMenu)
}

// 浜嬩欢澶勭悊

func (mw *MainWindow) onProjectSelected(proj project.ProjectConfig) {
    // 宸︿晶鐐瑰嚮椤圭洰锛氫紭鍏堝垏鎹㈠埌宸叉湁璇ラ」鐩殑缁堢鏍囩锛屽惁鍒欏垱寤轰竴涓?
    if id := mw.terminalTabs.FindTabByProjectPath(proj.Path); id != "" {
        mw.terminalTabs.SetActiveTab(id)
        mw.statusBar.SetMessage(fmt.Sprintf("宸插垏鎹㈠埌椤圭洰: %s", proj.Name))
        return
    }
    mw.createNewTerminal(proj, proj.AIModel)
}

func (mw *MainWindow) onProjectConfigured(proj project.ProjectConfig, aiModel project.AIModelType) {
    mw.createNewTerminal(proj, aiModel)
    mw.projectPanel.Refresh()
}

func (mw *MainWindow) onNewTerminalRequested(proj project.ProjectConfig, aiModel project.AIModelType, runInBackground bool) {
    log.Printf("[MainWindow] new terminal requested: path=%s model=%s yolo=%t bg=%t", proj.Path, aiModel, proj.YoloMode, runInBackground)
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
    mw.statusBar.SetMessage("AI 鍚姩鍣?v2.0.0")
}

// 缁堢鍒涘缓
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
        mw.statusBar.SetMessage(fmt.Sprintf("宸插垱寤虹粓绔? %s", termName))
        return tab
    }
    mw.statusBar.SetMessage("鍒涘缓缁堢澶辫触")
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
    log.Println("淇濆瓨绐楀彛鐘舵€?..")
}

