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

// SettingsDialog 设置对话框
type SettingsDialog struct {
	window     fyne.Window
	onChanged  func(map[string]interface{})

	// 弹窗组件
	dialog *dialog.CustomDialog

	// 左侧导航
	categoryList *widget.List
	categories   []string

	// 右侧内容区
	contentArea *fyne.Container

	// 设置页面
	generalPage *GeneralSettingsPage
	ollamaPage  *OllamaSettingsPage
	aiModelPage *AIModelSettingsPage
	proxyPage   *ProxySettingsPage
	advancedPage *AdvancedSettingsPage

	// 按钮
	saveButton  *widget.Button
	resetButton *widget.Button
	closeButton *widget.Button

	// 当前设置
	currentSettings map[string]interface{}
}

// GeneralSettingsPage 一般设置页面
type GeneralSettingsPage struct {
	container    fyne.CanvasObject
	languageSelect *widget.Select
	themeSelect    *widget.RadioGroup
	autoSaveCheck  *widget.Check
	startupCheck   *widget.Check
	trayCheck      *widget.Check
	welcomeCheck   *widget.Check
}

// OllamaSettingsPage Ollama设置页面
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

	// 模型数据
	models []OllamaModel
}

// OllamaModel Ollama模型信息
type OllamaModel struct {
	Name     string
	Size     string
	Status   string // "已载入", "未载入"
	Loaded   bool
}

// AIModelSettingsPage AI模型设置页面
type AIModelSettingsPage struct {
	container fyne.CanvasObject
	// TODO: AI模型相关设置
}

// ProxySettingsPage 网络代理设置页面
type ProxySettingsPage struct {
	container fyne.CanvasObject
	// TODO: 代理设置
}

// AdvancedSettingsPage 进阶设置页面
type AdvancedSettingsPage struct {
	container fyne.CanvasObject
	// TODO: 进阶设置
}

// NewSettingsDialog 创建设置对话框
func NewSettingsDialog(parent fyne.Window, onChanged func(map[string]interface{})) *SettingsDialog {
	d := &SettingsDialog{
		window:          parent,
		onChanged:       onChanged,
		currentSettings: make(map[string]interface{}),
		categories: []string{
			"📋 一般設置",
			"🤖 Ollama",
			"🔧 AI 模型",
			"🌐 網路代理",
			"⚙️ 進階設置",
		},
	}

	d.initializeUI()
	return d
}

// initializeUI 初始化设置对话框UI
func (d *SettingsDialog) initializeUI() {
	// 创建设置页面
	d.generalPage = d.createGeneralPage()
	d.ollamaPage = d.createOllamaPage()
	d.aiModelPage = d.createAIModelPage()
	d.proxyPage = d.createProxyPage()
	d.advancedPage = d.createAdvancedPage()

	// 左侧分类列表
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
	d.categoryList.Select(0) // 默认选择第一个

	// 右侧内容区
	d.contentArea = container.NewMax()

	// 主要布局（左右分栏）
	leftPanel := container.NewBorder(
		widget.NewRichTextFromMarkdown("### 設置分類"), // 顶部标题
		nil, nil, nil,
		d.categoryList,
	)

	// 按钮
	d.saveButton = widget.NewButtonWithIcon("💾儲存", theme.DocumentSaveIcon(), d.onSaveClicked)
	d.resetButton = widget.NewButtonWithIcon("🔄重置", theme.ViewRefreshIcon(), d.onResetClicked)
	d.closeButton = widget.NewButtonWithIcon("❌關閉", theme.CancelIcon(), d.onCloseClicked)

	buttonRow := container.NewHBox(
		d.saveButton,
		d.resetButton,
		layout.NewSpacer(),
		d.closeButton,
	)

	// 主布局
	mainLayout := container.NewBorder(
		nil,        // 顶部留空
		buttonRow,  // 底部按钮
		leftPanel,  // 左侧分类
		nil,        // 右侧留空
		d.contentArea, // 中央内容
	)

	// 创建弹窗
	d.dialog = dialog.NewCustom("⚙️ 系統設置", "", mainLayout, d.window)
	d.dialog.Resize(fyne.NewSize(700, 600))
}

// createGeneralPage 创建一般设置页面
func (d *SettingsDialog) createGeneralPage() *GeneralSettingsPage {
	page := &GeneralSettingsPage{}

	// 界面语言
	page.languageSelect = widget.NewSelect([]string{
		"繁體中文", "简体中文", "English", "日本語",
	}, nil)
	page.languageSelect.SetSelected("繁體中文")

	// 主题选择
	page.themeSelect = widget.NewRadioGroup([]string{
		"亮色", "暗色", "跟隨系統",
	}, nil)
	page.themeSelect.SetSelected("暗色")

	// 自动保存
	page.autoSaveCheck = widget.NewCheck("啟用項目自動儲存", nil)
	page.autoSaveCheck.SetChecked(true)

	// 启动设置
	page.startupCheck = widget.NewCheck("開機時自動啟動", nil)
	page.startupCheck.SetChecked(true)

	page.trayCheck = widget.NewCheck("最小化到系統托盤", nil)
	page.trayCheck.SetChecked(true)

	page.welcomeCheck = widget.NewCheck("啟動時顯示歡迎畫面", nil)

	// 布局
	languageSection := container.NewVBox(
		widget.NewLabel("🌍 介面語言:"),
		page.languageSelect,
	)

	themeSection := container.NewVBox(
		widget.NewLabel("🎨 主題:"),
		page.themeSelect,
	)

	autoSaveSection := container.NewVBox(
		widget.NewLabel("💾 自動儲存:"),
		page.autoSaveCheck,
	)

	startupSection := container.NewVBox(
		widget.NewLabel("🚀 啟動設置:"),
		page.startupCheck,
		page.trayCheck,
		page.welcomeCheck,
	)

	page.container = container.NewVBox(
		widget.NewRichTextFromMarkdown("### 🌐 一般設置"),
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

// createOllamaPage 创建Ollama设置页面
func (d *SettingsDialog) createOllamaPage() *OllamaSettingsPage {
	page := &OllamaSettingsPage{
		models: []OllamaModel{
			{Name: "qwen2.5:7b", Size: "4.0GB", Status: "已載入", Loaded: true},
			{Name: "llama3:8b", Size: "4.5GB", Status: "未載入", Loaded: false},
			{Name: "codellama:7b", Size: "3.8GB", Status: "未載入", Loaded: false},
			{Name: "mistral:7b", Size: "4.1GB", Status: "未載入", Loaded: false},
		},
	}

	// Ollama服务设置
	page.addressEntry = widget.NewEntry()
	page.addressEntry.SetText("http://localhost:11434")

	page.statusLabel = widget.NewLabel("🟢 已連線")
	page.reconnectBtn = widget.NewButtonWithIcon("🔄 重新連線", theme.ViewRefreshIcon(), func() {
		page.statusLabel.SetText("🔄 連線中...")
		// TODO: 实际重连逻辑
	})

	serviceSection := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 🌐 Ollama 服務設置"),
		container.NewVBox(
			widget.NewLabel("服務地址:"),
			page.addressEntry,
			container.NewHBox(
				widget.NewLabel("連線狀態:"),
				page.statusLabel,
				page.reconnectBtn,
			),
		),
	)

	// 模型列表
	page.modelList = widget.NewList(
		func() int { return len(page.models) },
		func() fyne.CanvasObject {
			status := widget.NewLabel("")
			name := widget.NewLabel("")
			name.TextStyle = fyne.TextStyle{Bold: true}
			size := widget.NewLabel("")
			indicator := widget.NewLabel("")

			return container.NewHBox(
				indicator, name, layout.NewSpacer(), size, status,
			)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			if id >= len(page.models) {
				return
			}
			model := page.models[id]
			row := obj.(*fyne.Container)

			indicator := row.Objects[0].(*widget.Label)
			name := row.Objects[1].(*widget.Label)
			size := row.Objects[3].(*widget.Label)
			status := row.Objects[4].(*widget.Label)

			if model.Loaded {
				indicator.SetText("●")
			} else {
				indicator.SetText("○")
			}
			name.SetText(model.Name)
			size.SetText(fmt.Sprintf("(%s)", model.Size))
			if model.Loaded {
				status.SetText("✅ 已載入")
			} else {
				status.SetText("💤 未載入")
			}
		},
	)

	// 模型管理按钮
	page.refreshBtn = widget.NewButtonWithIcon("🔄 重新整理", theme.ViewRefreshIcon(), func() {
		page.modelList.Refresh()
	})

	page.downloadBtn = widget.NewButtonWithIcon("➕ 下載新模型", theme.DownloadIcon(), func() {
		// TODO: 下载模型对话框
	})

	page.deleteBtn = widget.NewButtonWithIcon("🗑️ 刪除模型", theme.DeleteIcon(), func() {
		// TODO: 删除选中的模型
	})

	modelButtons := container.NewHBox(
		page.refreshBtn, page.downloadBtn, page.deleteBtn,
	)

	modelSection := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 📦 已安裝模型"),
		page.modelList,
		modelButtons,
	)

	// 模型使用设置
	page.defaultSelect = widget.NewSelect([]string{
		"qwen2.5:7b", "llama3:8b", "codellama:7b", "mistral:7b",
	}, nil)
	page.defaultSelect.SetSelected("qwen2.5:7b")

	page.optimizeCheck = widget.NewCheck("啟用智能查詢優化", nil)
	page.optimizeCheck.SetChecked(true)

	page.autoSwitchCheck = widget.NewCheck("自動模型切換 (根據項目類型)", nil)
	page.autoSwitchCheck.SetChecked(true)

	page.preloadCheck = widget.NewCheck("預載入常用模型", nil)

	usageSection := container.NewVBox(
		widget.NewRichTextFromMarkdown("### ⚡ 模型使用設置"),
		container.NewVBox(
			widget.NewLabel("預設查詢優化模型:"),
			page.defaultSelect,
		),
		page.optimizeCheck,
		page.autoSwitchCheck,
		page.preloadCheck,
	)

	// 滚动容器
	scroll := container.NewScroll(container.NewVBox(
		serviceSection,
		widget.NewSeparator(),
		modelSection,
		widget.NewSeparator(),
		usageSection,
	))

	page.container = scroll
	return page
}

// createAIModelPage 创建AI模型设置页面
func (d *SettingsDialog) createAIModelPage() *AIModelSettingsPage {
	page := &AIModelSettingsPage{}

	content := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 🔧 AI 模型設置"),
		widget.NewLabel("AI模型設置功能開發中..."),
	)

	page.container = content
	return page
}

// createProxyPage 创建代理设置页面
func (d *SettingsDialog) createProxyPage() *ProxySettingsPage {
	page := &ProxySettingsPage{}

	content := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 🌐 網路代理設置"),
		widget.NewLabel("網路代理設置功能開發中..."),
	)

	page.container = content
	return page
}

// createAdvancedPage 创建进阶设置页面
func (d *SettingsDialog) createAdvancedPage() *AdvancedSettingsPage {
	page := &AdvancedSettingsPage{}

	content := container.NewVBox(
		widget.NewRichTextFromMarkdown("### ⚙️ 進階設置"),
		widget.NewLabel("進階設置功能開發中..."),
	)

	page.container = content
	return page
}

// Show 显示设置对话框
func (d *SettingsDialog) Show() {
	d.loadCurrentSettings()
	d.dialog.Show()
}

// Hide 隐藏设置对话框
func (d *SettingsDialog) Hide() {
	d.dialog.Hide()
}

// 事件处理方法

func (d *SettingsDialog) onCategorySelected(id widget.ListItemID) {
	switch id {
	case 0: // 一般设置
		d.contentArea.Objects = []fyne.CanvasObject{d.generalPage.container}
	case 1: // Ollama
		d.contentArea.Objects = []fyne.CanvasObject{d.ollamaPage.container}
	case 2: // AI模型
		d.contentArea.Objects = []fyne.CanvasObject{d.aiModelPage.container}
	case 3: // 网络代理
		d.contentArea.Objects = []fyne.CanvasObject{d.proxyPage.container}
	case 4: // 进阶设置
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
	// TODO: 重置到默认设置
	dialog.ShowConfirm("確認重置", "確定要重置所有設置到默認值嗎？", func(confirmed bool) {
		if confirmed {
			d.resetToDefaults()
		}
	}, d.window)
}

func (d *SettingsDialog) onCloseClicked() {
	d.Hide()
}

// 工具方法

func (d *SettingsDialog) loadCurrentSettings() {
	// TODO: 从配置文件加载当前设置
}

func (d *SettingsDialog) collectSettings() map[string]interface{} {
	settings := make(map[string]interface{})

	// 一般设置
	settings["language"] = d.generalPage.languageSelect.Selected
	settings["theme"] = d.generalPage.themeSelect.Selected
	settings["auto_save"] = d.generalPage.autoSaveCheck.Checked
	settings["startup"] = d.generalPage.startupCheck.Checked
	settings["tray"] = d.generalPage.trayCheck.Checked
	settings["welcome"] = d.generalPage.welcomeCheck.Checked

	// Ollama设置
	settings["ollama_address"] = d.ollamaPage.addressEntry.Text
	settings["ollama_default_model"] = d.ollamaPage.defaultSelect.Selected
	settings["ollama_optimize"] = d.ollamaPage.optimizeCheck.Checked
	settings["ollama_auto_switch"] = d.ollamaPage.autoSwitchCheck.Checked
	settings["ollama_preload"] = d.ollamaPage.preloadCheck.Checked

	return settings
}

func (d *SettingsDialog) resetToDefaults() {
	// 重置一般设置
	d.generalPage.languageSelect.SetSelected("繁體中文")
	d.generalPage.themeSelect.SetSelected("暗色")
	d.generalPage.autoSaveCheck.SetChecked(true)
	d.generalPage.startupCheck.SetChecked(true)
	d.generalPage.trayCheck.SetChecked(true)
	d.generalPage.welcomeCheck.SetChecked(false)

	// 重置Ollama设置
	d.ollamaPage.addressEntry.SetText("http://localhost:11434")
	d.ollamaPage.defaultSelect.SetSelected("qwen2.5:7b")
	d.ollamaPage.optimizeCheck.SetChecked(true)
	d.ollamaPage.autoSwitchCheck.SetChecked(true)
	d.ollamaPage.preloadCheck.SetChecked(false)
}