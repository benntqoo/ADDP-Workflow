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

// TerminalTabContainer 缁堢鏍囩椤靛鍣?
type TerminalTabContainer struct {
	terminalManager *terminal.TerminalManager

	// UI缁勪欢
	tabContainer *container.AppTabs
	tabHeader    *fyne.Container
	content      *fyne.Container

	// 鏍囩椤电鐞?
	tabs         map[string]*TerminalTab
	activeTabID  string
	nextTabID    int
	mutex        sync.RWMutex
}

// TerminalTab 缁堢鏍囩椤?
type TerminalTab struct {
	id           string
	name         string
	terminalType terminal.TerminalType
	project      project.ProjectConfig

	// UI缁勪欢
	content      *fyne.Container
	terminal     *widget.Entry // 绠€鍖栫増鏈紝鍚庣画鍙互鏇挎崲涓虹湡姝ｇ殑缁堢缁勪欢
	outputArea   *widget.RichText
	inputArea    *widget.Entry
	statusLabel  *widget.Label

	// 鐘舵€?
	active       bool
	running      bool
}

// NewTerminalTabContainer 鍒涘缓缁堢鏍囩椤靛鍣?
func NewTerminalTabContainer(tm *terminal.TerminalManager) *TerminalTabContainer {
	container := &TerminalTabContainer{
		terminalManager: tm,
		tabs:           make(map[string]*TerminalTab),
		nextTabID:      1,
	}

	container.initializeUI()
	return container
}

// initializeUI 鍒濆鍖朥I
func (tc *TerminalTabContainer) initializeUI() {
	// 鍒涘缓鏍囩椤靛鍣?
	tc.tabContainer = container.NewAppTabs()

	// 鍒涘缓"鏂板缓"鏍囩椤?
	tc.addNewTabButton()

	// 鏍囩椤靛ご閮ㄥ鍣?
	tc.tabHeader = container.NewBorder(nil, nil, nil, nil, tc.tabContainer)

	// 鍐呭鍖哄煙锛堟樉绀哄綋鍓嶆椿鍔ㄦ爣绛鹃〉鐨勫唴瀹癸級
	tc.content = container.NewMax()

	// 璁剧疆鏍囩椤靛垏鎹簨浠?
	tc.tabContainer.OnChanged = tc.onTabChanged
}

// addNewTabButton 娣诲姞"鏂板缓"鏍囩椤垫寜閽?
func (tc *TerminalTabContainer) addNewTabButton() {
	newTabContent := container.NewCenter(
		widget.NewButtonWithIcon("新建终端", theme.ContentAddIcon(), func() {
			// TODO: 瑙﹀彂鏂板缓缁堢瀵硅瘽妗?
		}),
	)

	newTab := &container.TabItem{
		Text: "+ 新建",
		Icon: theme.ContentAddIcon(),
		Icon: theme.ContentAddIcon(),
		Content: newTabContent,
	}
	tc.tabContainer.Append(newTab)
// CreateTab 鍒涘缓鏂扮殑缁堢鏍囩椤?
func (tc *TerminalTabContainer) CreateTab(name string, termConfig terminal.TerminalConfig, proj project.ProjectConfig) *TerminalTab {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	// 鐢熸垚鍞竴ID
	tabID := fmt.Sprintf("tab_%d", tc.nextTabID)
	tc.nextTabID++

	// 鍒涘缓缁堢鏍囩椤?
	tab := &TerminalTab{
		id:           tabID,
		name:         name,
		terminalType: termConfig.Type,
		project:      proj,
	}

	// 鍒濆鍖栨爣绛鹃〉UI
	tab.initializeUI()

	// 鍚姩缁堢
	go tab.startTerminal(termConfig)

	// 娣诲姞鍒版爣绛鹃〉瀹瑰櫒锛堝湪"鏂板缓"鏍囩椤典箣鍓嶏級
	tabContent := tab.GetContent()
	appTab := &container.TabItem{
		Text: name,
		Content: tabContent,
	}

	// 鎻掑叆鍒版渶鍚庝竴涓綅缃紙"鏂板缓"鏍囩椤典箣鍓嶏級
	tc.tabContainer.Append(appTab)

	// 淇濆瓨鏍囩椤靛紩鐢?
	tc.tabs[tabID] = tab

	// 璁剧疆涓烘椿鍔ㄦ爣绛鹃〉
	tc.SetActiveTab(tabID)

	return tab
}

// RemoveTab 绉婚櫎鏍囩椤?
func (tc *TerminalTabContainer) RemoveTab(tabID string) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	tab, exists := tc.tabs[tabID]
	if !exists {
		return
	}

	// 鍋滄缁堢
	tab.stopTerminal()

	// 浠庢爣绛鹃〉瀹瑰櫒涓Щ闄?
	for i := 0; i < len(tc.tabContainer.Items)-1; i++ { // 鎺掗櫎"鏂板缓"鏍囩椤?
		if tc.tabContainer.Items[i].Text == tab.name {
			tc.tabContainer.RemoveIndex(i)
			break
		}
	}

	// 浠庢槧灏勪腑鍒犻櫎
	delete(tc.tabs, tabID)

	// 濡傛灉鍒犻櫎鐨勬槸娲诲姩鏍囩椤碉紝鍒囨崲鍒颁笅涓€涓?
	if tc.activeTabID == tabID {
		tc.activateNextTab()
	}
}

// SetActiveTab 璁剧疆娲诲姩鏍囩椤?
func (tc *TerminalTabContainer) SetActiveTab(tabID string) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	tab, exists := tc.tabs[tabID]
	if !exists {
		return
	}

	// 鍙栨秷涔嬪墠鐨勬椿鍔ㄧ姸鎬?
	if tc.activeTabID != "" {
		if prevTab := tc.tabs[tc.activeTabID]; prevTab != nil {
			prevTab.active = false
		}
	}

	// 璁剧疆鏂扮殑娲诲姩鏍囩椤?
	tc.activeTabID = tabID
	tab.active = true

	// 鏇存柊鍐呭鍖哄煙
	tc.content.Objects = []fyne.CanvasObject{tab.GetContent()}
	tc.content.Refresh()

	// 鍒囨崲鏍囩椤靛鍣ㄧ殑閫変腑鐘舵€?
	for i := 0; i < len(tc.tabContainer.Items)-1; i++ {
		if tc.tabContainer.Items[i].Text == tab.name {
			tc.tabContainer.SelectTabIndex(i)
			break
		}
	}
}

// GetActiveTab 鑾峰彇褰撳墠娲诲姩鐨勬爣绛鹃〉
func (tc *TerminalTabContainer) GetActiveTab() *TerminalTab {
	tc.mutex.RLock()
	defer tc.mutex.RUnlock()

	if tc.activeTabID == "" {
		return nil
	}

	return tc.tabs[tc.activeTabID]
}

// GetTabHeader 鑾峰彇鏍囩椤靛ご閮ㄥ鍣?
func (tc *TerminalTabContainer) GetTabHeader() *fyne.Container {
	return tc.tabHeader
}

// GetContent 鑾峰彇鍐呭鍖哄煙瀹瑰櫒
func (tc *TerminalTabContainer) GetContent() *fyne.Container {
	return tc.content
}

// onTabChanged 鏍囩椤靛垏鎹簨浠跺鐞?
func (tc *TerminalTabContainer) onTabChanged(tab *container.TabItem) {
	// 鏌ユ壘瀵瑰簲鐨勭粓绔爣绛鹃〉
	for tabID, termTab := range tc.tabs {
		if termTab.name == tab.Text {
			tc.SetActiveTab(tabID)
			return
		}
	}
}

// activateNextTab 婵€娲讳笅涓€涓爣绛鹃〉
func (tc *TerminalTabContainer) activateNextTab() {
	if len(tc.tabs) == 0 {
		tc.activeTabID = ""
		tc.content.Objects = []fyne.CanvasObject{}
		tc.content.Refresh()
		return
	}

	// 婵€娲荤涓€涓彲鐢ㄧ殑鏍囩椤?
	for tabID := range tc.tabs {
		tc.SetActiveTab(tabID)
		return
	}
}

// TerminalTab 鏂规硶瀹炵幇

// initializeUI 鍒濆鍖栫粓绔爣绛鹃〉UI
func (tab *TerminalTab) initializeUI() {
	// 杈撳嚭鍖哄煙锛堝彧璇伙級
	tab.outputArea = widget.NewRichText()
	tab.outputArea.Wrapping = fyne.TextWrapWord
	tab.outputArea.Scroll = container.ScrollBoth

	// 杈撳叆鍖哄煙
	tab.inputArea = widget.NewEntry()
	tab.inputArea.SetPlaceHolder("杈撳叆鍛戒护鎴栦笌AI浜や簰...")
	tab.inputArea.OnSubmitted = tab.onInputSubmitted

	// 鐘舵€佹爣绛?
	tab.statusLabel = widget.NewLabel("鍑嗗灏辩华")

	// 宸ュ叿鏍?
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.MediaPlayIcon(), tab.onStartTerminal),
		widget.NewToolbarAction(theme.MediaStopIcon(), tab.onStopTerminal),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), tab.onClearOutput),
		widget.NewToolbarAction(theme.SettingsIcon(), tab.onTerminalSettings),
	)

	// 搴曢儴鐘舵€佹爮
	statusBar := container.NewBorder(nil, nil, tab.statusLabel, nil, tab.statusLabel)

	// 涓昏鍐呭甯冨眬
	tab.content = container.NewBorder(
		toolbar,           // 椤堕儴锛氬伐鍏锋爮
		container.NewVBox( // 搴曢儴锛氳緭鍏ュ尯鍩熷拰鐘舵€佹爮
			tab.inputArea,
			statusBar,
		),
		nil, nil,          // 宸﹀彸鐣欑┖
		container.NewScroll(tab.outputArea), // 涓ぎ锛氳緭鍑哄尯鍩?
	)
}

// startTerminal 鍚姩缁堢
func (tab *TerminalTab) startTerminal(config terminal.TerminalConfig) {
	tab.running = true
	tab.statusLabel.SetText("运行中...")

	// TODO: 闆嗘垚鐪熸鐨勭粓绔鐞嗗櫒
	tab.appendOutput(fmt.Sprintf("正在启动 %s...\n", config.Type))
	tab.appendOutput(fmt.Sprintf("工作目录: %s\n", config.WorkingDir))
	tab.appendOutput(fmt.Sprintf("模式: %s\n", map[bool]string{true: "YOLO", false: "普通"}[config.YoloMode]))
	tab.appendOutput("终端已启动，准备接收命令\n\n")
}

// stopTerminal 鍋滄缁堢
func (tab *TerminalTab) stopTerminal() {
	tab.running = false
	tab.statusLabel.SetText("已停止")
	tab.appendOutput("\n终端已停止\n")
}

// GetContent 鑾峰彇鏍囩椤靛唴瀹?
func (tab *TerminalTab) GetContent() *fyne.Container {
	return tab.content
}

// GetID 鑾峰彇鏍囩椤礗D
func (tab *TerminalTab) GetID() string {
	return tab.id
}

// SwitchProject 鍒囨崲椤圭洰
func (tab *TerminalTab) SwitchProject(proj project.ProjectConfig) {
	tab.project = proj
	tab.appendOutput(fmt.Sprintf("\n馃搧 宸插垏鎹㈠埌椤圭洰: %s\n", proj.Name))
	tab.appendOutput(fmt.Sprintf("馃搷 璺緞: %s\n\n", proj.Path))
}

// appendOutput 杩藉姞杈撳嚭鍐呭
func (tab *TerminalTab) appendOutput(text string) {
	currentText := tab.outputArea.String() + text
	tab.outputArea.ParseMarkdown(currentText)
	tab.outputArea.Refresh()
}

// 浜嬩欢澶勭悊鏂规硶

func (tab *TerminalTab) onInputSubmitted(input string) {
	if input == "" {
		return
	}

	// 鏄剧ず鐢ㄦ埛杈撳叆
	tab.appendOutput(fmt.Sprintf("> %s\n", input))

	// TODO: 鍙戦€佸埌缁堢绠＄悊鍣ㄥ鐞?
	tab.appendOutput("馃摑 鍛戒护宸叉帴鏀讹紝姝ｅ湪澶勭悊...\n")

	// 娓呯┖杈撳叆
	tab.inputArea.SetText("")
}

func (tab *TerminalTab) onStartTerminal() {
	if !tab.running {
		// TODO: 閲嶆柊鍚姩缁堢
		tab.appendOutput("馃攧 閲嶆柊鍚姩缁堢...\n")
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
	tab.appendOutput("馃搫 杈撳嚭宸叉竻绌篭n\n")
}

func (tab *TerminalTab) onTerminalSettings() {
	// TODO: 鏄剧ず缁堢璁剧疆
	tab.appendOutput("鈿欙笍 缁堢璁剧疆鍔熻兘寮€鍙戜腑...\n")
}

