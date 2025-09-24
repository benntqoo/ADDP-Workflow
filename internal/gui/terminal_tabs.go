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

// TerminalTabContainer 缂佸牏顏弽鍥╊劮妞ら潧顔愰崳?
type TerminalTabContainer struct {
	terminalManager *terminal.TerminalManager

	// UI缂佸嫪娆?
	tabContainer *container.AppTabs
	tabHeader    *fyne.Container
	content      *fyne.Container

	// 閺嶅洨顒锋い鐢殿吀閻?
	tabs         map[string]*TerminalTab
	activeTabID  string
	nextTabID    int
	mutex        sync.RWMutex
}

// TerminalTab 缂佸牏顏弽鍥╊劮妞?
type TerminalTab struct {
	id           string
	name         string
	terminalType terminal.TerminalType
	project      project.ProjectConfig

	// UI缂佸嫪娆?
	content      *fyne.Container
	terminal     *widget.Entry // 缁犫偓閸栨牜澧楅張顒婄礉閸氬海鐢婚崣顖欎簰閺囨寧宕叉稉铏规埂濮濓絿娈戠紒鍫㈩伂缂佸嫪娆?
	outputArea   *widget.RichText
	inputArea    *widget.Entry
	statusLabel  *widget.Label

	// 閻樿埖鈧?
	active       bool
	running      bool
}

// NewTerminalTabContainer 閸掓稑缂撶紒鍫㈩伂閺嶅洨顒锋い闈涱啇閸?
func NewTerminalTabContainer(tm *terminal.TerminalManager) *TerminalTabContainer {
	container := &TerminalTabContainer{
		terminalManager: tm,
		tabs:           make(map[string]*TerminalTab),
		nextTabID:      1,
	}

	container.initializeUI()
	return container
}

// initializeUI 閸掓繂顫愰崠鏈
func (tc *TerminalTabContainer) initializeUI() {
	// 閸掓稑缂撻弽鍥╊劮妞ら潧顔愰崳?
	tc.tabContainer = container.NewAppTabs()

	// 閸掓稑缂?閺傛澘缂?閺嶅洨顒锋い?
	tc.addNewTabButton()

	// 閺嶅洨顒锋い闈涖仈闁劌顔愰崳?
	tc.tabHeader = container.NewBorder(nil, nil, nil, nil, tc.tabContainer)

	// 閸愬懎顔愰崠鍝勭厵閿涘牊妯夌粈鍝勭秼閸撳秵妞块崝銊︾垼缁涢箖銆夐惃鍕敶鐎圭櫢绱?
	tc.content = container.NewMax()

	// 鐠佸墽鐤嗛弽鍥╊劮妞ら潧鍨忛幑顫皑娴?
	tc.tabContainer.OnChanged = tc.onTabChanged
}

// addNewTabButton 濞ｈ濮?閺傛澘缂?閺嶅洨顒锋い鍨瘻闁?
func (tc *TerminalTabContainer) addNewTabButton() {
	newTabContent := container.NewCenter(
		widget.NewButtonWithIcon("新建终端", theme.ContentAddIcon(), func() {
			// TODO: 打开新建终端对话框
		}),
	)

	newTab := &container.TabItem{
		Text: "+ 新建",
		Icon: theme.ContentAddIcon(),
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

    appTab := &container.TabItem{
        Text:    name,
        Content: tab.GetContent(),
    }
    tc.tabContainer.Append(appTab)

    tc.tabs[tabID] = tab
    tc.SetActiveTab(tabID)
    return tab
}

// RemoveTab 缁夊娅庨弽鍥╊劮妞?
func (tc *TerminalTabContainer) RemoveTab(tabID string) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	tab, exists := tc.tabs[tabID]
	if !exists {
		return
	}

	// 閸嬫粍顒涚紒鍫㈩伂
	tab.stopTerminal()

	// 娴犲孩鐖ｇ粵楣冦€夌€圭懓娅掓稉顓犘╅梽?
	for i := 0; i < len(tc.tabContainer.Items)-1; i++ { // 閹烘帡娅?閺傛澘缂?閺嶅洨顒锋い?
		if tc.tabContainer.Items[i].Text == tab.name {
			tc.tabContainer.RemoveIndex(i)
			break
		}
	}

	// 娴犲孩妲х亸鍕厬閸掔娀娅?
	delete(tc.tabs, tabID)

	// 婵″倹鐏夐崚鐘绘珟閻ㄥ嫭妲稿ú璇插З閺嶅洨顒锋い纰夌礉閸掑洦宕查崚棰佺瑓娑撯偓娑?
	if tc.activeTabID == tabID {
		tc.activateNextTab()
	}
}

// SetActiveTab 鐠佸墽鐤嗗ú璇插З閺嶅洨顒锋い?
func (tc *TerminalTabContainer) SetActiveTab(tabID string) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	tab, exists := tc.tabs[tabID]
	if !exists {
		return
	}

	// 閸欐牗绉锋稊瀣閻ㄥ嫭妞块崝銊уЦ閹?
	if tc.activeTabID != "" {
		if prevTab := tc.tabs[tc.activeTabID]; prevTab != nil {
			prevTab.active = false
		}
	}

	// 鐠佸墽鐤嗛弬鎵畱濞茶濮╅弽鍥╊劮妞?
	tc.activeTabID = tabID
	tab.active = true

	// 閺囧瓨鏌婇崘鍛啇閸栧搫鐓?
	tc.content.Objects = []fyne.CanvasObject{tab.GetContent()}
	tc.content.Refresh()

	// 閸掑洦宕查弽鍥╊劮妞ら潧顔愰崳銊ф畱闁鑵戦悩鑸碘偓?
	for i := 0; i < len(tc.tabContainer.Items)-1; i++ {
		if tc.tabContainer.Items[i].Text == tab.name {
			tc.tabContainer.SelectTabIndex(i)
			break
		}
	}
}

// GetActiveTab 閼惧嘲褰囪ぐ鎾冲濞茶濮╅惃鍕垼缁涢箖銆?
func (tc *TerminalTabContainer) GetActiveTab() *TerminalTab {
	tc.mutex.RLock()
	defer tc.mutex.RUnlock()

	if tc.activeTabID == "" {
		return nil
	}

	return tc.tabs[tc.activeTabID]
}

// GetTabHeader 閼惧嘲褰囬弽鍥╊劮妞ら潧銇旈柈銊ヮ啇閸?
func (tc *TerminalTabContainer) GetTabHeader() *fyne.Container {
	return tc.tabHeader
}

// GetContent 閼惧嘲褰囬崘鍛啇閸栧搫鐓欑€圭懓娅?
func (tc *TerminalTabContainer) GetContent() *fyne.Container {
	return tc.content
}

// onTabChanged 閺嶅洨顒锋い闈涘瀼閹诡澀绨ㄦ禒璺侯槱閻?
func (tc *TerminalTabContainer) onTabChanged(tab *container.TabItem) {
	// 閺屻儲澹樼€电懓绨查惃鍕矒缁旑垱鐖ｇ粵楣冦€?
	for tabID, termTab := range tc.tabs {
		if termTab.name == tab.Text {
			tc.SetActiveTab(tabID)
			return
		}
	}
}

// activateNextTab 濠碘偓濞茶绗呮稉鈧稉顏呯垼缁涢箖銆?
func (tc *TerminalTabContainer) activateNextTab() {
	if len(tc.tabs) == 0 {
		tc.activeTabID = ""
		tc.content.Objects = []fyne.CanvasObject{}
		tc.content.Refresh()
		return
	}

	// 濠碘偓濞茶崵顑囨稉鈧稉顏勫讲閻劎娈戦弽鍥╊劮妞?
	for tabID := range tc.tabs {
		tc.SetActiveTab(tabID)
		return
	}
}

// TerminalTab 閺傝纭剁€圭偟骞?

// initializeUI 閸掓繂顫愰崠鏍矒缁旑垱鐖ｇ粵楣冦€塙I
func (tab *TerminalTab) initializeUI() {
	// 鏉堟挸鍤崠鍝勭厵閿涘牆褰х拠浼欑礆
	tab.outputArea = widget.NewRichText()
	tab.outputArea.Wrapping = fyne.TextWrapWord
	tab.outputArea.Scroll = container.ScrollBoth

	// 鏉堟挸鍙嗛崠鍝勭厵
	tab.inputArea = widget.NewEntry()
	tab.inputArea.SetPlaceHolder("鏉堟挸鍙嗛崨鎴掓姢閹存牔绗孉I娴溿倓绨?..")
	tab.inputArea.OnSubmitted = tab.onInputSubmitted

	// 閻樿埖鈧焦鐖ｇ粵?
	tab.statusLabel = widget.NewLabel("閸戝棗顦亸杈╁崕")

	// 瀹搞儱鍙块弽?
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.MediaPlayIcon(), tab.onStartTerminal),
		widget.NewToolbarAction(theme.MediaStopIcon(), tab.onStopTerminal),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), tab.onClearOutput),
		widget.NewToolbarAction(theme.SettingsIcon(), tab.onTerminalSettings),
	)

	// 鎼存洟鍎撮悩鑸碘偓浣圭埉
	statusBar := container.NewBorder(nil, nil, tab.statusLabel, nil, tab.statusLabel)

	// 娑撴槒顩﹂崘鍛啇鐢啫鐪?
	tab.content = container.NewBorder(
		toolbar,           // 妞ゅ爼鍎撮敍姘紣閸忛攱鐖?
		container.NewVBox( // 鎼存洟鍎撮敍姘崇翻閸忋儱灏崺鐔锋嫲閻樿埖鈧焦鐖?
			tab.inputArea,
			statusBar,
		),
		nil, nil,          // 瀹革箑褰搁悾娆戔敄
		container.NewScroll(tab.outputArea), // 娑擃厼銇庨敍姘崇翻閸戝搫灏崺?
	)
}

// startTerminal 閸氼垰濮╃紒鍫㈩伂
func (tab *TerminalTab) startTerminal(config terminal.TerminalConfig) {
	tab.running = true
	tab.statusLabel.SetText("运行中...")

	tab.appendOutput(fmt.Sprintf("正在启动 %s...\n", config.Type))
	tab.appendOutput(fmt.Sprintf("工作目录: %s\n", config.WorkingDir))
	tab.appendOutput(fmt.Sprintf("模式: %s\n", map[bool]string{true: "YOLO", false: "普通"}[config.YoloMode]))
	tab.appendOutput("终端已启动，准备接收命令\n\n")
}

// stopTerminal 閸嬫粍顒涚紒鍫㈩伂
func (tab *TerminalTab) stopTerminal() {
	tab.running = false
	tab.statusLabel.SetText("已停止")
	tab.appendOutput("\n终端已停止\n")
}

// GetContent 閼惧嘲褰囬弽鍥╊劮妞ら潧鍞寸€?
func (tab *TerminalTab) GetContent() *fyne.Container {
	return tab.content
}

// GetID 閼惧嘲褰囬弽鍥╊劮妞ょD
func (tab *TerminalTab) GetID() string {
	return tab.id
}

// SwitchProject 閸掑洦宕叉い鍦窗
func (tab *TerminalTab) SwitchProject(proj project.ProjectConfig) {
	tab.project = proj
	tab.appendOutput(fmt.Sprintf("\n棣冩惂 瀹告彃鍨忛幑銏犲煂妞ゅ湱娲? %s\n", proj.Name))
	tab.appendOutput(fmt.Sprintf("棣冩惙 鐠侯垰绶? %s\n\n", proj.Path))
}

// appendOutput 鏉╄棄濮炴潏鎾冲毉閸愬懎顔?
func (tab *TerminalTab) appendOutput(text string) {
	currentText := tab.outputArea.String() + text
	tab.outputArea.ParseMarkdown(currentText)
	tab.outputArea.Refresh()
}

// 娴滃娆㈡径鍕倞閺傝纭?

func (tab *TerminalTab) onInputSubmitted(input string) {
	if input == "" {
		return
	}

	// 閺勫墽銇氶悽銊﹀煕鏉堟挸鍙?
	tab.appendOutput(fmt.Sprintf("> %s\n", input))

	// TODO: 閸欐垿鈧礁鍩岀紒鍫㈩伂缁狅紕鎮婇崳銊ヮ槱閻?
	tab.appendOutput("棣冩憫 閸涙垝鎶ゅ鍙夊复閺€璁圭礉濮濓絽婀径鍕倞...\n")

	// 濞撳懐鈹栨潏鎾冲弳
	tab.inputArea.SetText("")
}

func (tab *TerminalTab) onStartTerminal() {
	if !tab.running {
		// TODO: 闁插秵鏌婇崥顖氬З缂佸牏顏?
		tab.appendOutput("棣冩敡 闁插秵鏌婇崥顖氬З缂佸牏顏?..\n")
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
	tab.appendOutput("棣冩惈 鏉堟挸鍤鍙夌缁岀n\n")
}

func (tab *TerminalTab) onTerminalSettings() {
	// TODO: 閺勫墽銇氱紒鍫㈩伂鐠佸墽鐤?
	tab.appendOutput("閳挎瑱绗?缂佸牏顏拋鍓х枂閸旂喕鍏樺鈧崣鎴滆厬...\n")
}



