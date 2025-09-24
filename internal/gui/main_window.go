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

    log.Println("鍒涘缓涓荤獥鍙?..")
    mw.window = mw.fyneApp.NewWindow("AI 鍚姩鍣?v2.0 - Desktop GUI 鐗?)
    if mw.window == nil {
        panic("鏃犳硶鍒涘缓 Fyne 绐楀彛")
    }

    mw.window.SetIcon(theme.ComputerIcon())
    mw.window.Resize(fyne.NewSize(mw.windowState.Width, mw.windowState.Height))
    mw.window.SetFixedSize(false)
    mw.window.CenterOnScreen()

    log.Println("绐楀彛鍒涘缓鎴愬姛锛岃缃睘鎬?..")
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
    mw.toolbar = mw.createToolbar()
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
    fileMenu := fyne.NewMenu("鏂囦欢",
        fyne.NewMenuItem("鎵撳紑椤圭洰", mw.onOpenProjectClicked),
        fyne.NewMenuItem("鏂板缓缁堢", mw.onNewTerminalClicked),
        fyne.NewMenuItemSeparator(),
        fyne.NewMenuItem("閫€鍑?, func() { mw.fyneApp.Quit() }),
    )

    toolsMenu := fyne.NewMenu("宸ュ叿",
        fyne.NewMenuItem("璁剧疆", mw.onSettingsClicked),
        fyne.NewMenuItem("鐩戞帶", mw.onMonitorClicked),
        fyne.NewMenuItemSeparator(),
        fyne.NewMenuItem("娓呯悊缂撳瓨", mw.onClearCacheClicked),
    )

    helpMenu := fyne.NewMenu("甯姪",
        fyne.NewMenuItem("浣跨敤璇存槑", mw.onHelpClicked),
        fyne.NewMenuItem("鍏充簬", mw.onAboutClicked),
    )

    return fyne.NewMainMenu(fileMenu, toolsMenu, helpMenu)
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
    mw.statusBar.SetMessage("璁剧疆宸插簲鐢?)
}

func (mw *MainWindow) onOpenProjectClicked() { mw.projectDialog.Show() }
func (mw *MainWindow) onSettingsClicked()    { mw.settingsDialog.Show() }

func (mw *MainWindow) onMonitorClicked() {
    mw.statusBar.SetMessage("鐩戞帶鍔熻兘寮€鍙戜腑...")
}

func (mw *MainWindow) onHelpClicked() {
    mw.statusBar.SetMessage("甯姪鍔熻兘寮€鍙戜腑...")
}

func (mw *MainWindow) onNewTerminalClicked() { mw.newTermDialog.Show() }

func (mw *MainWindow) onClearCacheClicked() {
    mw.statusBar.SetMessage("缂撳瓨宸叉竻鐞?)
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

