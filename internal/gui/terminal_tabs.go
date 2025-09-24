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

// TerminalTabContainer 缂傚倸鍊风粈渚€鎮ц箛娑辨晜妞ゆ帒瀚崘鈧梺鎼炲劗閺呪晠宕濋纰辨富闁靛牆顦板☉褎銇勯弮鈧崝娆忕暦?
type TerminalTabContainer struct {
	terminalManager *terminal.TerminalManager

	// UI缂傚倸鍊风粈浣衡偓姘间簻閳?
	tabContainer *container.AppTabs
	tabHeader    *fyne.Container
	content      *fyne.Container

	// 闂備礁鎼粔鏉懨洪妸鈺婃晢闂佸灝顑冩禍婊堟煟閵忊槅鍟忛柛姘ｅ亾闂?
	tabs         map[string]*TerminalTab
	activeTabID  string
	nextTabID    int
	mutex        sync.RWMutex
}

// TerminalTab 缂傚倸鍊风粈渚€鎮ц箛娑辨晜妞ゆ帒瀚崘鈧梺鎼炲劗閺呪晠宕濋纰辨富?
type TerminalTab struct {
	id           string
	name         string
	terminalType terminal.TerminalType
	project      project.ProjectConfig

	// UI缂傚倸鍊风粈浣衡偓姘间簻閳?
	content      *fyne.Container
	terminal     *widget.Entry // 缂傚倷鑳舵慨顓㈠磻閹剧粯鐓曢柡宥冨妿婢у崬顭胯椤ㄥ﹤顕ｉ悽鍓叉晢濠电姴瀚ˇ顕€姊洪崨濠呭婵炲鍏橀幃鑺ョ節濮橆剛鐓戝銈嗙墬閻熴倗鑺辨导瀛樼厸闁搞儯鍔岄婊呪偓瑙勬礀瀵墎绮欐径鎰優閻熸瑥瀚悞绋库攽椤旂晫绠扮紒鎻掔仢閳绘捇骞嬮悩鐢碉紲闂佸壊鍋侀崺鍕閸屾粎纾藉ù锝囶焾椤徰勭箾?
	outputArea   *widget.RichText
	inputArea    *widget.Entry
	statusLabel  *widget.Label

	// 闂備胶绮…鍫ュ春閺嶎厼鐒?
	active       bool
	running      bool
}

// NewTerminalTabContainer 闂備礁鎲＄敮妤冪矙閹寸姷纾介柟鎯у绾惧ジ鏌涢鐘茬労濞寸厧鍊块弻鈥愁吋閸涱喖绐涘銈嗗浮閺€閬嶅Φ閸曨垱顥堟繛鎴濆船閺咃綁姊?
func NewTerminalTabContainer(tm *terminal.TerminalManager) *TerminalTabContainer {
	container := &TerminalTabContainer{
		terminalManager: tm,
		tabs:           make(map[string]*TerminalTab),
		nextTabID:      1,
	}

	container.initializeUI()
	return container
}

// initializeUI 闂備礁鎲＄敮妤冩崲閸岀儑缍栭柟鐗堟緲缁€宀勬煛閸偅顎?
func (tc *TerminalTabContainer) initializeUI() {
	// 闂備礁鎲＄敮妤冪矙閹寸姷纾介柟鎹愵嚙閸愨偓闂佹悶鍎弲鈺呭礉椤旂⒈娓婚柕鍫濐槹濞懷勩亜閺冣偓閸旀瑥鐣?
	tc.tabContainer = container.NewAppTabs()

	// 闂備礁鎲＄敮妤冪矙閹寸姷纾?闂備礁鎼崐鐟邦熆濡偐纾?闂備礁鎼粔鏉懨洪妸鈺婃晢闂佸灝顑冩禍?
	tc.addNewTabButton()

	// 闂備礁鎼粔鏉懨洪妸鈺婃晢闂佸灝顑冩禍婊堟⒒閸喓鈼ゅù鐘茬墦濮婃椽顢欓崫鍕瀳濡炪倖姊归崝娆忕暦?
	tc.tabHeader = container.NewBorder(nil, nil, nil, nil, tc.tabContainer)

	// 闂備礁鎲￠崝鏇㈠箠鎼搭煈鏁婇柟鐗堟緲缁€宀勬煕濠靛棗顏柛妯兼暬閺屻劌鈽夊Ο鐓庮杸濠碘槅鍨伴ˇ杈╁垝閸儱浼犻柛鏇ㄥ幘姝囬梻浣告啞閹稿摜绮旈棃娑辨富闁秆勵殔缁€澶愭煏婵犲繐顩柛銊ㄥ亹缁辨帒鈽夐姀銏㈡毇闂侀潧妫楅ˇ鐢稿箚閸愵喖绀嬫い鎰╁灪閺嗐儵鏌ｆ惔鈩冭础婵炲懌鍨诲Σ?
	tc.content = container.NewMax()

	// 闂佽崵濮崇粈浣规櫠娴犲鍋柛鈩冪☉閸愨偓闂佹悶鍎弲鈺呭礉椤旂⒈娓婚柕鍫濐槹濞懷囨煕閵娿儳绠绘鐐村灴閿濈偤顢橀悩鍨吇濠?
	tc.tabContainer.OnChanged = tc.onTabChanged
}

// addNewTabButton 婵犵數鍎戠紞鈧い鏇嗗嫭鍙?闂備礁鎼崐鐟邦熆濡偐纾?闂備礁鎼粔鏉懨洪妸鈺婃晢闂佸灝顑冩禍婊堟煕椤垵浜滈柣锕€顭峰?
func (tc *TerminalTabContainer) addNewTabButton() {
	newTabContent := container.NewCenter(
		widget.NewButtonWithIcon("闁哄倹婢樼紓鎾剁磼閸埄浼?, theme.ContentAddIcon(), func() {
		widget.NewButtonWithIcon("鏂板缓缁堢", theme.ContentAddIcon(), func() {
			// TODO: 鎵撳紑鏂板缓缁堢瀵硅瘽妗?
		}),
	newTab := &container.TabItem{
		Text: "+ 闁哄倹婢樼紓?,
		Text: "+ 鏂板缓",
		Content: newTabContent,
	}
    tc.tabContainer.Append(newTab)
}

// CreateTab 闁告帗绋戠紓鎾诲棘閹殿喗鐣辩紓浣哥墢椤忣剟寮介崶鈺婂姰濡?func (tc *TerminalTabContainer) CreateTab(name string, termConfig terminal.TerminalConfig, proj project.ProjectConfig) *TerminalTab {
// CreateTab 鍒涘缓鏂扮殑缁堢鏍囩椤?
func (tc *TerminalTabContainer) CreateTab(name string, termConfig terminal.TerminalConfig, proj project.ProjectConfig) *TerminalTab {

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

// RemoveTab 缂傚倷绀侀ˇ顖炩€﹀畡鎵虫瀺閹兼番鍔岄崘鈧梺鎼炲劗閺呪晠宕濋纰辨富?
func (tc *TerminalTabContainer) RemoveTab(tabID string) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	tab, exists := tc.tabs[tabID]
	if !exists {
		return
	}

	// 闂備胶顭堥鍡欏垝瀹ュ鏁嗘繛鎴炴皑绾惧ジ鏌涢鐘茬労濞?
	tab.stopTerminal()

	// 濠电偛顕慨瀵糕偓娑掓櫊閹儵鏁愰崶鈺呪攺婵°倧绲介崯銊╁焵椤掆偓椤﹂亶鍩€椤掆偓濠€閬嶅箠閹炬枼鏋嶉柟鐑樻尵閳绘梹銇勯幘璺侯潙闁冲搫鎳忛埛?
	for i := 0; i < len(tc.tabContainer.Items)-1; i++ { // 闂備礁婀遍崕銈囨暜閳ユ枼鏋?闂備礁鎼崐鐟邦熆濡偐纾?闂備礁鎼粔鏉懨洪妸鈺婃晢闂佸灝顑冩禍?
		if tc.tabContainer.Items[i].Text == tab.name {
			tc.tabContainer.RemoveIndex(i)
			break
		}
	}

	// 濠电偛顕慨瀵糕偓娑掓櫅鍗辩憸蹇旂閹间礁绀嬫い蹇撴搐楠炩偓闂備礁鎲＄敮鐐寸箾閳ь剚绻?
	delete(tc.tabs, tabID)

	// 濠电姷顣介埀顒€鍟块埀顒€缍婇幃妯诲緞閹邦剛顦梺缁橆焽缁垶鎮甸悢鍏肩厽闁靛鍎遍顏呬繆閼碱剦鐒鹃悡銈夋偣閸ャ劌绲荤悮鐔兼⒑閸濆嫮孝婵炲眰鍔戦、妤呮煥鐎ｃ劋绨荤紒鐐緲椤﹁京绮堟径鎰厱闁圭儤鍨舵径鍕偓瑙勬礃閻撯€崇暦濮橆叏绱ｅù锝堟閹差喗绻涢幋鐐村皑闁稿鎸搁埥?
	if tc.activeTabID == tabID {
		tc.activateNextTab()
	}
}

// SetActiveTab 闂佽崵濮崇粈浣规櫠娴犲鍋柛鈩冾樅閾忚瀚氶柟缁樺醇濡ゅ懏鐓涚€广儱鎳忓婵囥亜閹烘鏁遍柕?
func (tc *TerminalTabContainer) SetActiveTab(tabID string) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	tab, exists := tc.tabs[tabID]
	if !exists {
		return
	}

	// 闂備礁鎲￠悷锕傛偋濡ゅ啰鐭撻梺鍨儑閳绘洟鎮楅敐搴濈盎妞ゃ倕鎳橀弻锝夊Ω閵夈儺浠惧┑锛勫仜濞尖€崇暦濠靛惟鐟滃秶鐟ч梻?
	if tc.activeTabID != "" {
		if prevTab := tc.tabs[tc.activeTabID]; prevTab != nil {
			prevTab.active = false
		}
	}

	// 闂佽崵濮崇粈浣规櫠娴犲鍋柛鈩冪☉濡﹢鏌熷▓鍨灍闁伙綀椴哥换娑㈡嚑妫版繂娈繝娈垮枓閺呯姴顕ｆ禒瀣倞闁冲搫锕ゆ慨鐗堜繆?
	tc.activeTabID = tabID
	tab.active = true

	// 闂備礁鎼ú銈夋偤閵娾晛钃熷┑鐘叉搐缁€鍐煕濞戝崬寮鹃柛鐔锋喘閺屾盯寮借閹牓鏌?
	tc.content.Objects = []fyne.CanvasObject{tab.GetContent()}
	tc.content.Refresh()

	// 闂備礁鎲＄敮鎺懨洪敃鈧悾鐑藉蓟閵夈儱鍞ㄩ梺鎼炲劗閺呪晠宕濋纰辨富闁靛牆顦板☉褎銇勯弮鈧崝娆忕暦閹惰棄惟鐟滃酣鎮鹃柆宥嗏拺妞ゆ劑鍩勫Σ鐑芥煠閸偄鐏撮柟顔规櫊閹虫粎鍠婂Ο杞板?
	for i := 0; i < len(tc.tabContainer.Items)-1; i++ {
		if tc.tabContainer.Items[i].Text == tab.name {
			tc.tabContainer.SelectTabIndex(i)
			break
		}
	}
}

// GetActiveTab 闂備礁鍚嬮崕鎶藉床閼艰翰浜归柛銉簵娴滃綊鏌熼幆褍鏆辨い銈呮噺缁绘盯鎳犳０婵嗘婵犳鍠掗弲鐘诲箚閸愵喖绀嬫い鎺戝€搁悘杈╃磽娴ｅ湱娲寸紒鐘崇墵婵?
func (tc *TerminalTabContainer) GetActiveTab() *TerminalTab {
	tc.mutex.RLock()
	defer tc.mutex.RUnlock()

	if tc.activeTabID == "" {
		return nil
	}

	return tc.tabs[tc.activeTabID]
}

// GetTabHeader 闂備礁鍚嬮崕鎶藉床閼艰翰浜归柛銉墮閸愨偓闂佹悶鍎弲鈺呭礉椤旂⒈娓婚柕鍫濐槹濞懷囨煏閸ャ劍顥堥柡灞界墦婵℃悂濡烽鍏兼珝闂?
func (tc *TerminalTabContainer) GetTabHeader() *fyne.Container {
	return tc.tabHeader
}

// GetContent 闂備礁鍚嬮崕鎶藉床閼艰翰浜归柛銉墮缁€鍐煕濞戝崬寮鹃柛鐔锋喘閺屾盯寮借閹牓鏌ｉ幘瀵告噮闁逞屽墮濠€閬嶅箠閹炬枼鏋?
func (tc *TerminalTabContainer) GetContent() *fyne.Container {
	return tc.content
}

// onTabChanged 闂備礁鎼粔鏉懨洪妸鈺婃晢闂佸灝顑冩禍婊堟⒒閸喓鈽夐柣搴°偢閺岀喓鎷犻埄鍐獢缂備降鍔婇崝宥囩矉閹烘梹宕夊〒姘煎灡琚╅梻?
func (tc *TerminalTabContainer) onTabChanged(tab *container.TabItem) {
	// 闂備礁鎼悮顐﹀磿閸欏鐝舵俊顖氱毞閸嬫捇鎮介崹顐㈡畬缂備降鍔嶉悡锟犲箚閸愵喖绀嬫い鎾跺Х閻撳绱撴担瑙勵棓闁搞劑浜堕幃銉╂晲閸モ晠鈹忔俊銈忕到閸熴劑鍩€?
	for tabID, termTab := range tc.tabs {
		if termTab.name == tab.Text {
			tc.SetActiveTab(tabID)
			return
		}
	}
}

// activateNextTab 婵犵數濮烽。浠嬪磻閹惧绠鹃柤濂割杺閸炶櫣绱掑Δ鈧幊蹇曠矙婢舵劕鐒垫い鎺嶈兌閳绘梹銇勮箛鎾村櫧闁搞劏鍋愮槐鎺戔槈閵忋垻鏆梺?
func (tc *TerminalTabContainer) activateNextTab() {
	if len(tc.tabs) == 0 {
		tc.activeTabID = ""
		tc.content.Objects = []fyne.CanvasObject{}
		tc.content.Refresh()
		return
	}

	// 婵犵數濮烽。浠嬪磻閹惧绠鹃柤鑹版硾閸氳銇勯幋婵囧缂佸顦甸崺鈧い鎺嶈兌閳绘梹銇勮箛鎾愁仼閻犱焦鐓￠弻锝夛綖椤掆偓婵℃寧绻涢崼鐔风伌鐎殿喕绮欏畷鍫曞煛婵犲倸袨濠?
	for tabID := range tc.tabs {
		tc.SetActiveTab(tabID)
		return
	}
}

// TerminalTab 闂備礁鎼崐浠嬶綖婢跺本鍏滈柛鎾茶閸嬫捇宕烽鐐版埛濡?

// initializeUI 闂備礁鎲＄敮妤冩崲閸岀儑缍栭柟鐗堟緲缁€宀勬煛瀹ュ骸浜為柣顓熷笧缁辨帡寮幋婵堜画闂佺粯鐗徊璺ㄥ垝閸偁鍋呴柛鎰惰吂閸嬫挻绻呭?
func (tab *TerminalTab) initializeUI() {
	// 闂佸搫顦悧濠囧箰閹间礁鍚规い鎾卞灩缁€宀勬煕濠靛棗顏柛妯兼暬閺屻劌鈽夊Ο鐓庮暫閻熸粍濡搁崨顔肩彴婵炴潙鍚嬮悷褏绮?
	tab.outputArea = widget.NewRichText()
	tab.outputArea.Wrapping = fyne.TextWrapWord
	tab.outputArea.Scroll = container.ScrollBoth

	// 闂佸搫顦悧濠囧箰閹间礁鐭楅柛鈩冪☉缁€宀勬煕濠靛棗顏柛?
	tab.inputArea = widget.NewEntry()
	tab.inputArea.SetPlaceHolder("鏉堟挸鍙嗛崨鎴掓姢閸氬骸娲栨潪锔藉⒔鐞?..")
	tab.inputArea.OnSubmitted = tab.onInputSubmitted

	// 闂備胶绮…鍫ュ春閺嶎厼鐒垫い鎴ｆ硶閸斿秹鏌ｉ弽鐢垫偧缂?
	// 閻樿埖鈧焦鏋冮張?
	tab.statusLabel = widget.NewLabel("缁屾椽妫?)
	// 闁诲氦顫夐幃鍫曞磿闁秴鐭楅柛褎顨呴崘鈧?
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.MediaPlayIcon(), tab.onStartTerminal),
		widget.NewToolbarAction(theme.MediaStopIcon(), tab.onStopTerminal),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), tab.onClearOutput),
		widget.NewToolbarAction(theme.SettingsIcon(), tab.onTerminalSettings),
	)

	// 闂佸湱鍘ч悺銊ッ洪悢鐓庣？闁圭虎鍠楅崑鎰版煠閸濄儺鏆柛瀣尰閹峰懘宕烽鐐靛姽
	statusBar := container.NewBorder(nil, nil, tab.statusLabel, nil, tab.statusLabel)

	// 濠电偞鍨堕幐璇参ｉ幒妞尖偓鍛存晝閸屾氨顦ㄩ梺鍛婁緱閸欏酣宕崶顒佸仺妞ゆ牗鑹鹃弳娆撴煟?
	tab.content = container.NewBorder(
		toolbar,           // 濠碉紕鍋戦崐鏇㈡偉婵傜纾块柟缁㈠枟閺咁剚鎱ㄥ鍡楀幐缂佹唻缍侀弻娑滅疀濞戞瑦鏆犻梺?
		container.NewVBox( // 闂佸湱鍘ч悺銊ッ洪悢鐓庣？闁圭虎鍠楅弲顒佹叏濡櫣锛嶇紓鍫濐煼閺屾稖绠涚€ｎ亜濮庨柣蹇撴禋閸ㄨ泛鐣峰ú顏呭亜闂佸灝顑呴顖炴⒑缂佹﹩鐒鹃柛鈺傜墵閸┾偓妞ゆ垼娉曢崝宥夋煟?
			tab.inputArea,
			statusBar,
		),
		nil, nil,          // 闁诲骸缍婂鑽ょ不閹达絻浜归柟闂寸劍閸嬧晜绻涢崱妯虹亶闁?
		container.NewScroll(tab.outputArea), // 濠电偞鍨堕幖鈺呭储婵傜桅閹兼番鍔嶉弲顒佹叏濡櫣锛嶇紓鍫濐煼閺屾盯骞嬪┑鍥ㄥ創闁诲繐娴氶崹璺虹暦?
	)
}

// startTerminal 闂備礁鎲￠崙褰掑垂閻楀牊鍙忛柍鍝勫暟绾惧ジ鏌涢鐘茬労濞?
func (tab *TerminalTab) startTerminal(config terminal.TerminalConfig) {
	tab.running = true
	tab.statusLabel.SetText("閺夆晜鍔橀、鎴炵▔?..")
	tab.statusLabel.SetText("运行中...")
	tab.appendOutput(fmt.Sprintf("婵繐绲藉﹢顏堝触椤栨艾袟 %s...\n", config.Type))
	tab.appendOutput(fmt.Sprintf("正在启动 %s...\\n", config.Type))
	tab.appendOutput(fmt.Sprintf("工作目录: %s\\n", config.WorkingDir))
	tab.appendOutput(fmt.Sprintf("模式: %s\\n", map[bool]string{true: "YOLO", false: "普通"}[config.YoloMode]))
	tab.appendOutput("终端已启动，准备接收命令\\n\\n")

// stopTerminal 闂備胶顭堥鍡欏垝瀹ュ鏁嗘繛鎴炴皑绾惧ジ鏌涢鐘茬労濞?
func (tab *TerminalTab) stopTerminal() {
	tab.running = false
	tab.statusLabel.SetText("鐎瑰憡褰冩禒鐘差潰?)
	tab.statusLabel.SetText("已停止")
	tab.appendOutput("\\n终端已停止\\n")

// GetContent 闂備礁鍚嬮崕鎶藉床閼艰翰浜归柛銉墮閸愨偓闂佹悶鍎弲鈺呭礉椤旂⒈娓婚柕鍫濐槹濞懷囨煕閻愬樊鍤熼柍?
func (tab *TerminalTab) GetContent() *fyne.Container {
	return tab.content
}

// GetID 闂備礁鍚嬮崕鎶藉床閼艰翰浜归柛銉墮閸愨偓闂佹悶鍎弲鈺呭礉椤旂⒈娓婚柕鍫濇祩濡槏
func (tab *TerminalTab) GetID() string {
	return tab.id
}

// SwitchProject 闂備礁鎲＄敮鎺懨洪敃鈧悾鐑藉矗婢跺牅绨婚梺闈╁瘜閸樼晫绮?
func (tab *TerminalTab) SwitchProject(proj project.ProjectConfig) {
	tab.project = proj
	tab.appendOutput(fmt.Sprintf("\n閸掑洦宕叉い鍦窗: %s\n", proj.Name))
	tab.appendOutput(fmt.Sprintf("鐠侯垰绶? %s\n\n", proj.Path))
}

// appendOutput 闂佸搫顦弲鐐参涢崟顒佸弿闁绘劕鐡ㄧ紞鍥煙閹冩毐婵絽顦甸弻娑㈠箛椤掍礁娅ｅ?
func (tab *TerminalTab) appendOutput(text string) {
	currentText := tab.outputArea.String() + text
	tab.outputArea.ParseMarkdown(currentText)
	tab.outputArea.Refresh()
}

// 濠电偛鐡ㄧ划宀勵敄閸曨偀鏋庨柕蹇娾偓宕囩獮闂佸憡娲﹂崢浠嬪磹閻愮儤鐓涢柛灞剧箥濞兼劗鐥?

func (tab *TerminalTab) onInputSubmitted(input string) {
	if input == "" {
		return
	}

	// 闂備礁鎼€氼剚鏅舵禒瀣︽慨妯垮煐閸嬨劑鏌曟繝蹇曠暠闁绘挻娲熷鍫曞醇閻旂纰嶉梺?
	tab.appendOutput(fmt.Sprintf("> %s\n", input))

	// TODO: 闂備礁鎲￠悷锕傚垂閸ф鐒垫い鎴ｆ硶椤︼箓鏌涢埡浣虹伇缂佽鲸甯″畷濂稿閳衡偓缁辨挾绱撴担鍝勵€撶紒杈ㄦ礋楠炲啯绻濋崶褔妫烽梺闈涱焾閸庡搫危闁秵鐓?
	tab.appendOutput("婵☆偓绲介崯顖炲箟?闂備礁鎲＄粙鎺楀垂濠靛绠柕鍫濇媼閸熷懘鏌涘▎蹇ｆЦ濠㈣泛绉归弻锟犲焵椤掑倹濯奸柛锔诲幘椤︻喖鈹戦鐣岀缂佽尪妫勯敃銏ゎ敂閸涱厾绐為梺鍛婃处閸樹粙宕?..\n")

	// 婵犵數鍋為幐鎼佸箠閹版澘鐓橀柡宥冨妽缂嶅洭鏌熼幆褍鏆辩€?
	tab.inputArea.SetText("")
}

func (tab *TerminalTab) onStartTerminal() {
	if !tab.running {
		// TODO: 闂傚倷鐒﹁ぐ鍐矓閻㈢钃熷┑鐘叉搐鐟欙附銇勯弽銊ㄥ鐞氱喓绱撻崒娆戭槮闁绘绻橀、?
		tab.appendOutput("婵☆偓绲介崯顖炲极?闂傚倷鐒﹁ぐ鍐矓閻㈢钃熷┑鐘叉搐鐟欙附銇勯弽銊ㄥ鐞氱喓绱撻崒娆戭槮闁绘绻橀、?..\n")
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
	tab.appendOutput("婵☆偓绲介崯顖炲箚?闂佸搫顦悧濠囧箰閹间礁鍚规い鎾跺剱閸熷懘鏌涘▎蹇ｆШ妞ゅ繐鎽滅槐鎺戠暆閳ь剟顢氶。鎿籲")
}

func (tab *TerminalTab) onTerminalSettings() {
	// TODO: 闂備礁鎼€氼剚鏅舵禒瀣︽慨妯夸含绾惧ジ鏌涢鐘茬労濞寸厧鍊块幃瑙勬媴缁嬪簱鎸冮梺?
	tab.appendOutput("闂備浇娅曠€笛囨偡鏉堚晝绠?缂傚倸鍊风粈渚€鎮ц箛娑辨晜妞ゆ帒鍊规刊濂告煕閹炬鎳忛悗顓㈡⒑閸涘﹥鈷掗柛鐘虫礋瀹曟螣娓氼垰娈ㄩ梺閫炲苯澧寸€规洩缍侀獮瀣煥閸℃绠?..\n")
}





