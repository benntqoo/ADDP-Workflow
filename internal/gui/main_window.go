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
	menuBar         *fyne.MainMenu
	toolbar         *widget.Toolbar
	projectPanel    *ProjectHistoryPanel
	terminalTabs    *TerminalTabContainer
	statusBar       *StatusBar

	// 瀵硅瘽妗嗙粍浠?
	projectDialog  *ProjectConfigDialog
	settingsDialog *SettingsDialog
	newTermDialog  *NewTerminalDialog

	// 绐楀彛鐘舵€?
	windowState *WindowState
}

// WindowState 缂佹劖顨呰ぐ娑㈡偐閼哥鍋撴担渚悁闁?
type WindowState struct {
	Width      float32 `json:"width"`
	Height     float32 `json:"height"`
	X          float32 `json:"x"`
	Y          float32 `json:"y"`
	Maximized  bool    `json:"maximized"`
	Theme      string  `json:"theme"`
	LeftPanelWidth float32 `json:"left_panel_width"`
}

// NewMainWindow 闁告帗绋戠紓鎾寸▔閼姐倗宕堕柛?
func NewMainWindow() *MainWindow {
	myApp := app.NewWithID("ai.launcher.desktop")
	myApp.SetIcon(theme.ComputerIcon())

	// Windows 娑撳绨查悽?CJK 娑撳顣介敍宀勪缉閸忓秳鑵戦弬鍥ㄦ煙濡?娑旇京鐖?
	if runtime.GOOS == "windows" {
		EnsureCJKFont()
		if fp := SelectCJKFont(); fp != "" {
			_ = ApplyCJKTheme(myApp, fp)
		}
	}

	// 閻熷鍘鹃悿鍡涘箠婢跺本鏆忛柛蹇撳暞閺嗏晠骞?- 闁革腹鈧ne v2.4.5濞戞搩鍙冨〒鍓佹啺娓氣偓閳ь剚宀告禍閿嬬▔瀹ュ懏鍊遍柡鍌滄嚀缁憋紕鎳涢鐘垫瀭
	// myApp.Metadata().Name 闁革负鍔嶉婵嬫偋閸喐鎷卞☉鎿冨幒缁楀宕ｉ婊勭函闁规亽鍎撮懖鐘诲磹?

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

// Run 闁告凹鍨版慨鈺傜▔閼姐倗宕堕柛?
func (mw *MainWindow) Run() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("GUI閺夆晜鍔橀、鎴﹀籍閸洘鏅╅悹? %v", r)
			panic(r) // 闂佹彃绉甸弻濠囧箮濞戞ê姣夐柨娑樼焷椤斺偓main闁告垼濮ら弳鐔稿緞閸曨厽鍊?
		}
	}()

	// 闁告梻濮惧ù鍥煀瀹ュ洨鏋?
	if err := mw.projectManager.LoadProjects(); err != nil {
		log.Printf("闁告梻濮惧ù鍥ㄣ亜閸︻厽绐楅梺鏉跨Ф閻ゅ棙寰勬潏顐バ? %v", err)
	}

	log.Println("閸掓稑缂撴稉鑽ょ崶閸?..")
	// 闁告帗绋戠紓鎾寸▔閼姐倗宕堕柛?
	mw.window = mw.fyneApp.NewWindow("AI 启动器 v2.0 - Desktop GUI 版")
	if mw.window == nil {
		panic("无法创建 Fyne 窗口")
	}

	mw.window.SetIcon(theme.ComputerIcon())
	mw.window.Resize(fyne.NewSize(mw.windowState.Width, mw.windowState.Height))
	mw.window.SetFixedSize(false)
	mw.window.CenterOnScreen()

	log.Println("缁愭褰涢崚娑樼紦閹存劕濮涢敍宀冾啎缂冾喖鐫橀幀?..")

	// 閻犱礁澧介悿鍡欑玻濡も偓瑜版盯宕楅幎鑺ワ紨閻炴稑濂旂拹?
	mw.window.SetCloseIntercept(func() {
		mw.saveWindowState()
		mw.fyneApp.Quit()
	})

	// 閹煎瓨姊婚弫銈嗙▔婵犳凹鏆?
	if mw.windowState.Theme == "dark" {
		mw.fyneApp.Settings().SetTheme(theme.DarkTheme())
	} else {
		mw.fyneApp.Settings().SetTheme(theme.LightTheme())
	}

	log.Println("閸掓繂顫愰崠?UI 缂佸嫪娆?..")
	// 闁告帗绻傞～鎰板礌閺堫檹缂備礁瀚▎?
	mw.initializeComponents()

	log.Println("閸掓稑缂撴稉璇茬鐏炩偓...")
	// 闁告帗绋戠紓鎾寸▔鐠囪尙顏撮悘鐐╁亾
	content := mw.createMainLayout()
	mw.window.SetContent(content)

	log.Println("鐠佸墽鐤嗛懣婊冨礋...")
	// 閻犱礁澧介悿鍡涙嚕濠婂啫绀?
	mw.window.SetMainMenu(mw.createMenuBar())

	log.Println("闁哄嫬澧介妵姘辩玻濡も偓瑜版盯鐛捄铏圭；濠殿喖顑勭花銊︾鐠虹儤鍎曢柣?..")
	// 闁哄嫬澧介妵姘辩玻濡も偓瑜?
	mw.window.ShowAndRun()
}

// initializeComponents 闁告帗绻傞～鎰板礌閺堫檹缂備礁瀚▎?
func (mw *MainWindow) initializeComponents() {
	// 闁告帗绋戠紓鎾愁啅閵夈儱寰旈柡?
	mw.toolbar = mw.createToolbar()

	// 闁告帗绋戠紓鎾愁啅閿旀寧娅犲銈呮贡濞蹭即宕㈤崱妤€钑夐梻鍫涘灪濠?
	mw.projectPanel = NewProjectHistoryPanel(mw.projectManager, mw.onProjectSelected)

	// 闁告帗绋戠紓鎾剁磼閸埄浼傞柡宥呮川椤掗攱銇勯棃娑卞晣闁?
	mw.terminalTabs = NewTerminalTabContainer(mw.terminalManager)

	// 闁告帗绋戠紓鎾绘偐閼哥鍋撴担鍦焿
	mw.statusBar = NewStatusBar()

	// 闁告帗绋戠紓鎾愁嚕閸︻厾宕剁紓浣稿濞嗐垽鏁嶉崼婵囶偨閺夆晝鍠庨崹鍨叏鐎ｎ亜顕ч柨?
	mw.projectDialog = NewProjectConfigDialog(mw.window, mw.projectManager, mw.onProjectConfigured)
	mw.settingsDialog = NewSettingsDialog(mw.window, mw.onSettingsChanged)
	mw.newTermDialog = NewNewTerminalDialog(mw.window, mw.projectManager, mw.onNewTerminalRequested)
}

// createMainLayout 闁告帗绋戠紓鎾寸▔鐠囪尙顏撮悘鐐╁亾
func (mw *MainWindow) createMainLayout() *fyne.Container {
	// 鐎归潻缂氶弲鑸点亜閸︻厽绐楅柛妯烘瑜板爼妫冮姀鈩冪凡闁挎稑鐗嗗ù鎰偓瑙勮壘椤旀梹鎯旈敂鑲╃
	leftPanel := container.NewBorder(nil, nil, nil, nil, mw.projectPanel.GetContainer())
	leftPanel.Resize(fyne.NewSize(mw.windowState.LeftPanelWidth, 0))

	// 闁告瑥鍘栭弲鑸电▔閺勫浚娲ｉ柛鎰噹椤旀劙宕?
	rightContent := container.NewBorder(
		mw.terminalTabs.GetTabHeader(), // 濡炪倕鐖奸崕鎾晬濮橆厾鍨肩紒娑㈢畺閵嗗寰勯幘顔煎姤
		mw.statusBar.GetContainer(),    // 閹煎瓨娲熼崕鎾晬濮樺崬笑闁诡兛鐒﹂悥?
		nil, nil,                       // 鐎归潻绠戣ぐ鎼佹偩濞嗘垟鏁?
		mw.terminalTabs.GetContent(),   // 濞戞搩鍘奸妵搴ㄦ晬濮樿京鐭掔紒鏃戝灠閸炲鈧?
	)

	// 濞戞捁顕х粩椋庝沪閳ь剟鏁嶅顑跨矗闁稿繘鏀遍悥?+ 鐎归潻绠戣ぐ鎼佸礆閸℃鍩夐柛鎰噹椤?
	mainLayout := container.NewBorder(
		mw.toolbar, // 濡炪倕鐖奸崕鎾晬濮橆兛绱ｉ柛蹇涙敱閻?
		nil,        // 閹煎瓨娲熼崕鎾偩濞嗘垟鏁?
		leftPanel,  // 鐎归潻缂氶弲鍫曟晬濮樿翰鈧秹鎯勯鈧浼村级?
		nil,        // 闁告瑥鍘栭弲鍫曟偩濞嗘垟鏁?
		rightContent, // 濞戞搩鍘奸妵搴ㄦ晬濮橆偄鐦滈悷鏇氱閸炲鈧?
	)

	return mainLayout
}

// createToolbar 闁告帗绋戠紓鎾愁啅閵夈儱寰旈柡?
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

// createMenuBar 闁告帗绋戠紓鎾绘嚕濠婂啫绀嬮柡?
func (mw *MainWindow) createMenuBar() *fyne.MainMenu {
	// 闁哄倸娲ｅ▎銏ゆ嚕濠婂啫绀?
	fileMenu := fyne.NewMenu("文件",
		fyne.NewMenuItem("打开项目", mw.onOpenProjectClicked),
		fyne.NewMenuItem("新建终端", mw.onNewTerminalClicked),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("退出", func() { mw.fyneApp.Quit() }),
	)

	// 鐎规悶鍎遍崣鍧楁嚕濠婂啫绀?
	toolsMenu := fyne.NewMenu("工具",
		fyne.NewMenuItem("设置", mw.onSettingsClicked),
		fyne.NewMenuItem("监控", mw.onMonitorClicked),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("清理缓存", mw.onClearCacheClicked),
	)

	// 閻㈩垼鍠栨慨顏堟嚕濠婂啫绀?
	helpMenu := fyne.NewMenu("帮助",
		fyne.NewMenuItem("使用说明", mw.onHelpClicked),
		fyne.NewMenuItem("关于", mw.onAboutClicked),
	)

	return fyne.NewMainMenu(fileMenu, toolsMenu, helpMenu)
}

// 濞存粌顑勫▎銏″緞閸曨厽鍊為柡鍌濐潐绾?

// onProjectSelected 濡炪倕婀卞ú浼存焻婢跺顏ュù婊冾儎濞?
func (mw *MainWindow) onProjectSelected(project project.ProjectConfig) {
	// 闁革负鍔岀紞瀣礈瀹ュ棙銇熼柛鏂诲妿缁挾绮╅娆掑幀闁告帒娲﹀畷鏌ュ礆娴肩补鍋撴径澶庡幀闁汇劌瀚伴妴宥夋儎?
	activeTab := mw.terminalTabs.GetActiveTab()
	if activeTab != nil {
		activeTab.SwitchProject(project)
		mw.statusBar.SetMessage(fmt.Sprintf("鐎瑰憡褰冮崹蹇涘箲閵忕姴鐓傚銈呮贡濞? %s", project.Name))
	} else {
		// 濠碘€冲€归悘澶娾柦閳╁啯绠掓繛鑼额嚙婵晝绱掗崼銏╀紓闁挎稑鑻崹鍗烆嚈閻戞ɑ鐓€闁汇劌瀚划鎾剁博?
		mw.createNewTerminal(project, project.AIModel)
	}
}

// onProjectConfigured 濡炪倕婀卞ú浼存煀瀹ュ洨鏋傞悗鐟版湰閸ㄦ碍绂嶇€ｂ晜顐?
func (mw *MainWindow) onProjectConfigured(project project.ProjectConfig, aiModel project.AIModelType) {
	// 闁告凹鍨版慨鈺呭棘閹殿喗鐣辩紓浣哥墢椤?
	mw.createNewTerminal(project, aiModel)
	mw.projectPanel.Refresh() // 闁告帡鏀遍弻濠冦亜閸︻厽绐楅柛鎺擃殙閵?
}

// onNewTerminalRequested 闁哄倹婢樼紓鎾剁磼閸埄浼傞悹鍥敱閻増绂嶇€ｂ晜顐?
func (mw *MainWindow) onNewTerminalRequested(project project.ProjectConfig, aiModel project.AIModelType, runInBackground bool) {
	tab := mw.createNewTerminal(project, aiModel)
	if !runInBackground && tab != nil {
		mw.terminalTabs.SetActiveTab(tab.GetID())
	}
}

// onSettingsChanged 閻犱礁澧介悿鍡涘矗濡粯绾ù婊冾儎濞?
func (mw *MainWindow) onSettingsChanged(settings map[string]interface{}) {
	// 閹煎瓨姊婚弫銈囨媼閸撗呮瀭闁告瑦蓱濞?
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

// 鐎规悶鍎遍崣鍧楀冀韫囧海鐨戝ù鐘烘硾椤︹晠鎮?

func (mw *MainWindow) onOpenProjectClicked() {
	mw.projectDialog.Show()
}

func (mw *MainWindow) onSettingsClicked() {
	mw.settingsDialog.Show()
}

func (mw *MainWindow) onMonitorClicked() {
	// TODO: 閻庡湱鍋熼獮鍥儎閹寸偛浠橀柛鏃傚枙閸?
	mw.statusBar.SetMessage("监控功能开发中...")
}

func (mw *MainWindow) onHelpClicked() {
	// TODO: 閻庡湱鍋熼獮鍥╂暜椤旂厧袠闁告梻鍠曢崗?
	mw.statusBar.SetMessage("帮助功能开发中...")
}

func (mw *MainWindow) onNewTerminalClicked() {
	mw.newTermDialog.Show()
}

func (mw *MainWindow) onClearCacheClicked() {
	// TODO: 閻庡湱鍋熼獮鍥╃磽閹惧磭鎽犳繛鎾虫噽閹?
	mw.statusBar.SetMessage("缓存已清理")
}

func (mw *MainWindow) onAboutClicked() {
	// TODO: 闁哄嫬澧介妵姘跺礂閸忓懐鑹鹃悗鐢殿攰閻﹁棄顩?
	mw.statusBar.SetMessage("AI 启动器 v2.0.0")
}

// 鐎规悶鍎遍崣鍧楀棘鐟欏嫮銆?

// createNewTerminal 闁告帗绋戠紓鎾诲棘閹殿喚鐭掔紒?
func (mw *MainWindow) createNewTerminal(project project.ProjectConfig, aiModel project.AIModelType) *TerminalTab {
	// 闁汇垻鍠愰崹姘辩磼閸埄浼傞柛姘Ф琚ㄩ柨娑欏哺閵嗗秹鎯勯鍏煎€?AI鐎规悶鍎遍崣?
	termName := fmt.Sprintf("%s(%s)", project.Name, aiModel.String())

	// 闁告帗绋戠紓鎾剁磼閸埄浼傞梺鏉跨Ф閻?
	termConfig := terminal.TerminalConfig{
		Type:       mw.getTerminalType(aiModel),
		Name:       termName,
		WorkingDir: project.Path,
		Command:    aiModel.GetCommand(project.YoloMode),
		YoloMode:   project.YoloMode,
	}

	// 闁告帗绋戠紓鎾剁磼閸埄浼傞柡宥呮川椤掗攱銇?
	tab := mw.terminalTabs.CreateTab(termName, termConfig, project)
	if tab != nil {
	mw.statusBar.SetMessage(fmt.Sprintf("已创建终端: %s", termName))
		return tab
	}

	mw.statusBar.SetMessage("创建终端失败")
	return nil
}

// getTerminalType 闁兼儳鍢茶ぐ鍥╃磼閸埄浼傜紒顐ヮ嚙閻?
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

// saveWindowState 濞ｅ洦绻傞悺銊х玻濡も偓瑜版盯鎮╅懜纰樺亾?
func (mw *MainWindow) saveWindowState() {
	// TODO: 閻庡湱鍋熼獮鍥╃玻濡も偓瑜版盯鎮╅懜纰樺亾娴ｇ懓鐦☉鏂挎噹鐎?
	log.Println("濞ｅ洦绻傞悺銊х玻濡も偓瑜版盯鎮╅懜纰樺亾?..")
}




