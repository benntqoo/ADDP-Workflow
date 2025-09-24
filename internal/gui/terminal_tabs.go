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

// 终端标签容器
type TerminalTabContainer struct {
    terminalManager *terminal.TerminalManager

    // UI
    tabContainer *container.AppTabs
    tabHeader    *fyne.Container
    content      *fyne.Container

    // 状态
    tabs        map[string]*TerminalTab
    activeTabID string
    nextTabID   int
    mutex       sync.RWMutex
}

// 单个终端标签
type TerminalTab struct {
    id           string
    name         string
    terminalType terminal.TerminalType
    project      project.ProjectConfig

    // UI
    content     *fyne.Container
    outputArea  *widget.RichText
    inputArea   *widget.Entry
    statusLabel *widget.Label

    // 状态
    active  bool
    running bool
}

func NewTerminalTabContainer(tm *terminal.TerminalManager) *TerminalTabContainer {
    tc := &TerminalTabContainer{
        terminalManager: tm,
        tabs:            make(map[string]*TerminalTab),
        nextTabID:       1,
    }
    tc.initializeUI()
    return tc
}

func (tc *TerminalTabContainer) initializeUI() {
    tc.tabContainer = container.NewAppTabs()
    tc.addNewTabButton()
    tc.tabHeader = container.NewBorder(nil, nil, nil, nil, tc.tabContainer)
    tc.content = container.NewMax()
    tc.tabContainer.OnChanged = tc.onTabChanged
}

func (tc *TerminalTabContainer) addNewTabButton() {
    newTabContent := container.NewCenter(
        widget.NewButtonWithIcon("新建终端", theme.ContentAddIcon(), func() {
            // TODO: 打开新建终端对话框
        }),
    )
    newTab := &container.TabItem{
        Text:    "+ 新建",
        Icon:    theme.ContentAddIcon(),
        Content: newTabContent,
    }
    tc.tabContainer.Append(newTab)
}

// CreateTab 创建新的终端标签页
func (tc *TerminalTabContainer) CreateTab(name string, termConfig terminal.TerminalConfig, proj project.ProjectConfig) *TerminalTab {
    tc.mutex.Lock()
    defer tc.mutex.Unlock()

    tabID := fmt.Sprintf("tab_%d", tc.nextTabID)
    tc.nextTabID++

    tab := &TerminalTab{
        id:           tabID,
        name:         name,
        terminalType: termConfig.Type,
        project:      proj,
    }
    tab.initializeUI()
    go tab.startTerminal(termConfig)

    appTab := &container.TabItem{Text: name, Content: tab.GetContent()}
    tc.tabContainer.Append(appTab)

    tc.tabs[tabID] = tab
    tc.SetActiveTab(tabID)
    return tab
}

// RemoveTab 移除标签
func (tc *TerminalTabContainer) RemoveTab(tabID string) {
    tc.mutex.Lock()
    defer tc.mutex.Unlock()

    tab, ok := tc.tabs[tabID]
    if !ok {
        return
    }
    tab.stopTerminal()
    for i := 0; i < len(tc.tabContainer.Items); i++ {
        if tc.tabContainer.Items[i].Text == tab.name {
            tc.tabContainer.RemoveIndex(i)
            break
        }
    }
    delete(tc.tabs, tabID)
    if tc.activeTabID == tabID {
        // 激活剩余任意标签，否则清空内容
        activated := false
        for id := range tc.tabs {
            tc.activeTabID = ""
            tc.SetActiveTab(id)
            activated = true
            break
        }
        if !activated {
            tc.activeTabID = ""
            tc.content.Objects = []fyne.CanvasObject{}
            tc.content.Refresh()
        }
    }
}

func (tc *TerminalTabContainer) SetActiveTab(tabID string) {
    tc.mutex.Lock()
    defer tc.mutex.Unlock()

    tab, ok := tc.tabs[tabID]
    if !ok {
        return
    }
    if tc.activeTabID != "" {
        if prev := tc.tabs[tc.activeTabID]; prev != nil {
            prev.active = false
        }
    }
    tc.activeTabID = tabID
    tab.active = true
    tc.content.Objects = []fyne.CanvasObject{tab.GetContent()}
    tc.content.Refresh()

    for i := 0; i < len(tc.tabContainer.Items); i++ {
        if tc.tabContainer.Items[i].Text == tab.name {
            tc.tabContainer.SelectTabIndex(i)
            break
        }
    }
}

func (tc *TerminalTabContainer) GetActiveTab() *TerminalTab {
    tc.mutex.RLock()
    defer tc.mutex.RUnlock()
    if tc.activeTabID == "" {
        return nil
    }
    return tc.tabs[tc.activeTabID]
}

func (tc *TerminalTabContainer) GetTabHeader() *fyne.Container { return tc.tabHeader }
func (tc *TerminalTabContainer) GetContent() *fyne.Container  { return tc.content }

func (tc *TerminalTabContainer) onTabChanged(tabItem *container.TabItem) {
    // 来自用户点击的选项卡切换，不要在此再次调用 SelectTabIndex 以避免递归。
    for id, t := range tc.tabs {
        if t.name == tabItem.Text {
            tc.activeTabID = id
            tc.content.Objects = []fyne.CanvasObject{t.GetContent()}
            tc.content.Refresh()
            break
        }
    }
}

// FindTabByProjectPath 根据项目路径查找标签ID（若存在多个，返回第一个）
func (tc *TerminalTabContainer) FindTabByProjectPath(path string) string {
    tc.mutex.RLock()
    defer tc.mutex.RUnlock()
    for id, t := range tc.tabs {
        if t.project.Path == path {
            return id
        }
    }
    return ""
}

// TerminalTab UI 初始化
func (tab *TerminalTab) initializeUI() {
    tab.outputArea = widget.NewRichText()
    tab.outputArea.Wrapping = fyne.TextWrapWord
    tab.outputArea.Scroll = container.ScrollBoth

    tab.inputArea = widget.NewEntry()
    tab.inputArea.SetPlaceHolder("输入命令后回车执行...")
    tab.inputArea.OnSubmitted = tab.onInputSubmitted

    tab.statusLabel = widget.NewLabel("空闲")

    toolbar := widget.NewToolbar(
        widget.NewToolbarAction(theme.MediaPlayIcon(), tab.onStartTerminal),
        widget.NewToolbarAction(theme.MediaStopIcon(), tab.onStopTerminal),
        widget.NewToolbarSeparator(),
        widget.NewToolbarAction(theme.ViewRefreshIcon(), tab.onClearOutput),
        widget.NewToolbarAction(theme.SettingsIcon(), tab.onTerminalSettings),
    )

    statusBar := container.NewBorder(nil, nil, tab.statusLabel, nil, tab.statusLabel)

    tab.content = container.NewBorder(
        toolbar,
        container.NewVBox(tab.inputArea, statusBar),
        nil, nil,
        container.NewScroll(tab.outputArea),
    )
}

func (tab *TerminalTab) startTerminal(config terminal.TerminalConfig) {
    tab.running = true
    tab.statusLabel.SetText("运行中...")
    tab.appendOutput(fmt.Sprintf("正在启动 %s...\n", config.Type))
    tab.appendOutput(fmt.Sprintf("工作目录: %s\n", config.WorkingDir))
    tab.appendOutput(fmt.Sprintf("模式: %s\n", map[bool]string{true: "YOLO", false: "普通"}[config.YoloMode]))
    tab.appendOutput("终端已启动，准备接收命令\n\n")
}

func (tab *TerminalTab) stopTerminal() {
    tab.running = false
    tab.statusLabel.SetText("已停止")
    tab.appendOutput("\n终端已停止\n")
}

func (tab *TerminalTab) GetContent() *fyne.Container { return tab.content }
func (tab *TerminalTab) GetID() string              { return tab.id }

func (tab *TerminalTab) SwitchProject(proj project.ProjectConfig) {
    tab.project = proj
    tab.appendOutput(fmt.Sprintf("\n切换项目: %s\n", proj.Name))
    tab.appendOutput(fmt.Sprintf("路径: %s\n\n", proj.Path))
}

func (tab *TerminalTab) appendOutput(text string) {
    current := tab.outputArea.String() + text
    tab.outputArea.ParseMarkdown(current)
    tab.outputArea.Refresh()
}

func (tab *TerminalTab) onInputSubmitted(input string) {
    if input == "" {
        return
    }
    tab.appendOutput(fmt.Sprintf("> %s\n", input))
    tab.appendOutput("命令已提交，执行中...\n")
    tab.inputArea.SetText("")
}

func (tab *TerminalTab) onStartTerminal() {
    if !tab.running {
        tab.appendOutput("正在启动终端...\n")
    }
}

func (tab *TerminalTab) onStopTerminal() { if tab.running { tab.stopTerminal() } }
func (tab *TerminalTab) onClearOutput()   { tab.outputArea.ParseMarkdown(""); tab.outputArea.Refresh(); tab.appendOutput("已清空输出\n") }
func (tab *TerminalTab) onTerminalSettings() { tab.appendOutput("终端设置暂未实现...\n") }

