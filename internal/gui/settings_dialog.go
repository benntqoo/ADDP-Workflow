package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// SettingsDialog 璁剧疆瀵硅瘽妗?
type SettingsDialog struct {
	window     fyne.Window
	onChanged  func(map[string]interface{})

	// 寮圭獥缁勪欢
	dialog *dialog.CustomDialog

	// 宸︿晶瀵艰埅
	categoryList *widget.List
	categories   []string

	// 鍙充晶鍐呭鍖?
	contentArea *fyne.Container

	// 璁剧疆椤甸潰
	generalPage *GeneralSettingsPage
	ollamaPage  *OllamaSettingsPage
	aiModelPage *AIModelSettingsPage
	proxyPage   *ProxySettingsPage
	advancedPage *AdvancedSettingsPage

	// 鎸夐挳
	saveButton  *widget.Button
	resetButton *widget.Button
	closeButton *widget.Button

	// 褰撳墠璁剧疆
	currentSettings map[string]interface{}
}

// GeneralSettingsPage 涓€鑸缃〉闈?
type GeneralSettingsPage struct {
	container    fyne.CanvasObject
	languageSelect *widget.Select
	themeSelect    *widget.RadioGroup
	autoSaveCheck  *widget.Check
	startupCheck   *widget.Check
	trayCheck      *widget.Check
	welcomeCheck   *widget.Check
}

// OllamaSettingsPage Ollama璁剧疆椤甸潰
type OllamaSettingsPage struct {
	container     fyne.CanvasObject
	addressEntry  *widget.Entry
	statusLabel   *widget.Label
	reconnectBtn  *widget.Button
	modelList     *widget.List
	refreshBtn    *widget.Button
	downloadBtn   *widget.Button
	deleteBtn     *widget.Button
	defaultSelect *widget.Select
	optimizeCheck *widget.Check
	autoSwitchCheck *widget.Check
	preloadCheck  *widget.Check

	// 妯″瀷鏁版嵁
	models []OllamaModel
}

// OllamaModel Ollama妯″瀷淇℃伅
type OllamaModel struct {
	Name     string
	Size     string
	Status   string // "宸茶浇鍏?, "鏈浇鍏?
	Loaded   bool
}

// AIModelSettingsPage AI妯″瀷璁剧疆椤甸潰
type AIModelSettingsPage struct {
	container fyne.CanvasObject
	// TODO: AI妯″瀷鐩稿叧璁剧疆
}

// ProxySettingsPage 缃戠粶浠ｇ悊璁剧疆椤甸潰
type ProxySettingsPage struct {
	container fyne.CanvasObject
	// TODO: 浠ｇ悊璁剧疆
}

// AdvancedSettingsPage 杩涢樁璁剧疆椤甸潰
type AdvancedSettingsPage struct {
	container fyne.CanvasObject
	// TODO: 杩涢樁璁剧疆
}

// NewSettingsDialog 鍒涘缓璁剧疆瀵硅瘽妗?
func NewSettingsDialog(parent fyne.Window, onChanged func(map[string]interface{})) *SettingsDialog {
	d := &SettingsDialog{
		window:          parent,
		onChanged:       onChanged,
		currentSettings: make(map[string]interface{}),
		categories: []string{
			"通用设置",
			"Ollama",
			"AI 模型",
			"网络代理",
			"高级设置",
		},
	}

	d.initializeUI()
	return d
}

// initializeUI 鍒濆鍖栬缃璇濇UI
func (d *SettingsDialog) initializeUI() {
	// 鍒涘缓璁剧疆椤甸潰
	d.generalPage = d.createGeneralPage()
	d.ollamaPage = d.createOllamaPage()
	d.aiModelPage = d.createAIModelPage()
	d.proxyPage = d.createProxyPage()
	d.advancedPage = d.createAdvancedPage()

	// 宸︿晶鍒嗙被鍒楄〃
	d.categoryList = widget.NewList(
		func() int { return len(d.categories) },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			label.SetText(d.categories[id])
		},
	)

    d.categoryList.OnSelected = d.onCategorySelected

    // 初始化右侧内容区，避免 OnSelected 访问 nil
    d.contentArea = container.NewMax()

    // 选择默认分类，触发一次渲染
    d.categoryList.Select(0)


    leftPanel := container.NewBorder(
        widget.NewRichTextFromMarkdown("### 设置导航"),
        nil, nil, nil,
        d.categoryList,
    )

    // 初始化底部按钮
    d.saveButton = widget.NewButtonWithIcon("保存", theme.DocumentSaveIcon(), d.onSaveClicked)
    d.resetButton = widget.NewButtonWithIcon("重置", theme.ViewRefreshIcon(), d.onResetClicked)
    d.closeButton = widget.NewButtonWithIcon("关闭", theme.CancelIcon(), d.onCloseClicked)

    buttonRow := container.NewHBox(
        d.saveButton,
        d.resetButton,
        layout.NewSpacer(),
        d.closeButton,
    )
	// 涓诲竷灞€
	mainLayout := container.NewBorder(
		nil,        // 椤堕儴鐣欑┖
		buttonRow,  // 搴曢儴鎸夐挳
		leftPanel,  // 宸︿晶鍒嗙被
		nil,        // 鍙充晶鐣欑┖
		d.contentArea, // 涓ぎ鍐呭
	)

	// 鍒涘缓寮圭獥
    d.dialog = dialog.NewCustom("系统设置", "", mainLayout, d.window)
    d.dialog.Resize(fyne.NewSize(700, 600))
}

// createGeneralPage 鍒涘缓涓€鑸缃〉闈?
func (d *SettingsDialog) createGeneralPage() *GeneralSettingsPage {
    page := &GeneralSettingsPage{}

    // 界面语言
    page.languageSelect = widget.NewSelect([]string{"简体中文", "English", "日本語"}, nil)
    page.languageSelect.SetSelected("简体中文")

    // 主题选择
    page.themeSelect = widget.NewRadioGroup([]string{"浅色", "暗色", "系统主题"}, nil)
    page.themeSelect.SetSelected("暗色")

    // 自动保存
    page.autoSaveCheck = widget.NewCheck("自动保存项目", nil)
    page.autoSaveCheck.SetChecked(true)

    // 启动设置
    page.startupCheck = widget.NewCheck("开机自启动", nil)
    page.startupCheck.SetChecked(true)

    page.trayCheck = widget.NewCheck("最小化到托盘", nil)
    page.trayCheck.SetChecked(true)

    page.welcomeCheck = widget.NewCheck("显示欢迎页", nil)

    // 布局
    languageSection := container.NewVBox(
        widget.NewLabel("界面语言:"),
        page.languageSelect,
    )

    themeSection := container.NewVBox(
        widget.NewLabel("主题:"),
        page.themeSelect,
    )

    autoSaveSection := container.NewVBox(
        widget.NewLabel("自动保存:"),
        page.autoSaveCheck,
    )

    startupSection := container.NewVBox(
        widget.NewLabel("启动设置:"),
        page.startupCheck,
        page.trayCheck,
        page.welcomeCheck,
    )

    page.container = container.NewVBox(
        widget.NewRichTextFromMarkdown("### 通用设置"),
        languageSection,
        widget.NewSeparator(),
        themeSection,
        widget.NewSeparator(),
        autoSaveSection,
        widget.NewSeparator(),
        startupSection,
    )

    return page
}

// createOllamaPage 鍒涘缓Ollama璁剧疆椤甸潰
func (d *SettingsDialog) createOllamaPage() *OllamaSettingsPage {
    page := &OllamaSettingsPage{
        models: []OllamaModel{
            {Name: "qwen2.5:7b", Size: "4.0GB", Status: "已加载", Loaded: true},
            {Name: "llama3:8b",  Size: "4.5GB", Status: "未加载", Loaded: false},
        },
    }

    // 服务设置
    page.addressEntry = widget.NewEntry()
    page.addressEntry.SetText("http://localhost:11434")

    page.statusLabel = widget.NewLabel("就绪")
    page.reconnectBtn = widget.NewButtonWithIcon("重新连接", theme.ViewRefreshIcon(), func() {
        page.statusLabel.SetText("连接中...")
        // TODO: 实际重连逻辑
    })

    serviceSection := container.NewVBox(
        widget.NewRichTextFromMarkdown("### Ollama 服务设置"),
        container.NewVBox(
            widget.NewLabel("服务地址:"),
            page.addressEntry,
            container.NewHBox(widget.NewLabel("连接状态:"), page.statusLabel, page.reconnectBtn),
        ),
    )

    // 模型列表
    page.modelList = widget.NewList(
        func() int { return len(page.models) },
        func() fyne.CanvasObject {
            name := widget.NewLabel("")
            size := widget.NewLabel("")
            status := widget.NewLabel("")
            return container.NewHBox(name, layout.NewSpacer(), size, status)
        },
        func(id widget.ListItemID, obj fyne.CanvasObject) {
            if id >= len(page.models) { return }
            m := page.models[id]
            row := obj.(*fyne.Container)
            name := row.Objects[0].(*widget.Label)
            size := row.Objects[2].(*widget.Label)
            status := row.Objects[3].(*widget.Label)
            name.SetText(m.Name)
            size.SetText(fmt.Sprintf("(%s)", m.Size))
            status.SetText(m.Status)
        },
    )

    page.refreshBtn = widget.NewButtonWithIcon("刷新", theme.ViewRefreshIcon(), func() {
        page.modelList.Refresh()
    })
    page.downloadBtn = widget.NewButtonWithIcon("下载模型", theme.DownloadIcon(), func() {})
    page.deleteBtn = widget.NewButtonWithIcon("删除模型", theme.DeleteIcon(), func() {})

    modelButtons := container.NewHBox(page.refreshBtn, page.downloadBtn, page.deleteBtn)
    modelSection := container.NewVBox(
        widget.NewRichTextFromMarkdown("### 已安装模型"),
        page.modelList,
        modelButtons,
    )

    // 使用设置
    page.defaultSelect = widget.NewSelect([]string{"qwen2.5:7b", "llama3:8b"}, nil)
    page.defaultSelect.SetSelected("qwen2.5:7b")
    page.optimizeCheck = widget.NewCheck("启用查询优化", nil)
    page.optimizeCheck.SetChecked(true)
    page.autoSwitchCheck = widget.NewCheck("根据任务自动切换模型", nil)
    page.autoSwitchCheck.SetChecked(true)
    page.preloadCheck = widget.NewCheck("预加载常用模型", nil)

    usageSection := container.NewVBox(
        widget.NewRichTextFromMarkdown("### 模型使用设置"),
        container.NewVBox(widget.NewLabel("默认模型:"), page.defaultSelect),
        page.optimizeCheck,
        page.autoSwitchCheck,
        page.preloadCheck,
    )

    page.container = container.NewScroll(container.NewVBox(
        serviceSection,
        widget.NewSeparator(),
        modelSection,
        widget.NewSeparator(),
        usageSection,
    ))

    return page
}

// createAIModelPage 鍒涘缓AI妯″瀷璁剧疆椤甸潰
func (d *SettingsDialog) createAIModelPage() *AIModelSettingsPage {
	page := &AIModelSettingsPage{}

	content := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 馃敡 AI 妯″瀷瑷疆"),
		widget.NewLabel("AI妯″瀷瑷疆鍔熻兘闁嬬櫦涓?.."),
	)

	page.container = content
	return page
}

// createProxyPage 鍒涘缓浠ｇ悊璁剧疆椤甸潰
func (d *SettingsDialog) createProxyPage() *ProxySettingsPage {
    page := &ProxySettingsPage{}

    content := container.NewVBox(
        widget.NewRichTextFromMarkdown("### 网络代理设置"),
        widget.NewLabel("网络代理配置功能开发中..."),
    )

    page.container = content
    return page
}

// createAdvancedPage 鍒涘缓杩涢樁璁剧疆椤甸潰
func (d *SettingsDialog) createAdvancedPage() *AdvancedSettingsPage {
    page := &AdvancedSettingsPage{}

    content := container.NewVBox(
        widget.NewRichTextFromMarkdown("### 高级设置"),
        widget.NewLabel("高级设置功能开发中..."),
    )

    page.container = content
    return page
}

// Show 鏄剧ず璁剧疆瀵硅瘽妗?
func (d *SettingsDialog) Show() {
	d.loadCurrentSettings()
	d.dialog.Show()
}

// Hide 闅愯棌璁剧疆瀵硅瘽妗?
func (d *SettingsDialog) Hide() {
	d.dialog.Hide()
}

// 浜嬩欢澶勭悊鏂规硶

func (d *SettingsDialog) onCategorySelected(id widget.ListItemID) {
	switch id {
	case 0: // 涓€鑸缃?
		d.contentArea.Objects = []fyne.CanvasObject{d.generalPage.container}
	case 1: // Ollama
		d.contentArea.Objects = []fyne.CanvasObject{d.ollamaPage.container}
	case 2: // AI妯″瀷
		d.contentArea.Objects = []fyne.CanvasObject{d.aiModelPage.container}
	case 3: // 缃戠粶浠ｇ悊
		d.contentArea.Objects = []fyne.CanvasObject{d.proxyPage.container}
	case 4: // 杩涢樁璁剧疆
		d.contentArea.Objects = []fyne.CanvasObject{d.advancedPage.container}
	}
	d.contentArea.Refresh()
}

func (d *SettingsDialog) onSaveClicked() {
	settings := d.collectSettings()
	if d.onChanged != nil {
		d.onChanged(settings)
	}
	d.Hide()
}

func (d *SettingsDialog) onResetClicked() {
    dialog.ShowConfirm(
        "重置设置",
        "确定要将所有设置重置为默认值吗？",
        func(confirmed bool) {
            if confirmed {
                d.resetToDefaults()
            }
        },
        d.window,
    )
}

func (d *SettingsDialog) onCloseClicked() {
	d.Hide()
}

// 宸ュ叿鏂规硶

func (d *SettingsDialog) loadCurrentSettings() {
	// TODO: 浠庨厤缃枃浠跺姞杞藉綋鍓嶈缃?
}

func (d *SettingsDialog) collectSettings() map[string]interface{} {
	settings := make(map[string]interface{})

	// 涓€鑸缃?
	settings["language"] = d.generalPage.languageSelect.Selected
	settings["theme"] = d.generalPage.themeSelect.Selected
	settings["auto_save"] = d.generalPage.autoSaveCheck.Checked
	settings["startup"] = d.generalPage.startupCheck.Checked
	settings["tray"] = d.generalPage.trayCheck.Checked
	settings["welcome"] = d.generalPage.welcomeCheck.Checked

	// Ollama璁剧疆
	settings["ollama_address"] = d.ollamaPage.addressEntry.Text
	settings["ollama_default_model"] = d.ollamaPage.defaultSelect.Selected
	settings["ollama_optimize"] = d.ollamaPage.optimizeCheck.Checked
	settings["ollama_auto_switch"] = d.ollamaPage.autoSwitchCheck.Checked
	settings["ollama_preload"] = d.ollamaPage.preloadCheck.Checked

	return settings
}

func (d *SettingsDialog) resetToDefaults() {
    // 通用设置默认值
    d.generalPage.languageSelect.SetSelected("简体中文")
    d.generalPage.themeSelect.SetSelected("暗色")
    d.generalPage.autoSaveCheck.SetChecked(true)
    d.generalPage.startupCheck.SetChecked(true)
    d.generalPage.trayCheck.SetChecked(true)
    d.generalPage.welcomeCheck.SetChecked(false)

    // Ollama 设置默认值
    d.ollamaPage.addressEntry.SetText("http://localhost:11434")
    d.ollamaPage.defaultSelect.SetSelected("qwen2.5:7b")
    d.ollamaPage.optimizeCheck.SetChecked(true)
    d.ollamaPage.autoSwitchCheck.SetChecked(true)
    d.ollamaPage.preloadCheck.SetChecked(false)
}

