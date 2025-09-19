package gui

import (
	"fmt"
	"image/color"
	"log"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"ai-launcher/internal/project"
	"ai-launcher/internal/terminal"
)

// App GUI应用程序结构
type App struct {
	fyneApp         fyne.App
	window          fyne.Window
	configManager   *project.ConfigManager
	terminalManager *terminal.TerminalManager

	// UI组件
	projectPathEntry   *widget.Entry
	projectNameEntry   *widget.Entry
	modelSelect        *widget.Select
	yoloModeCheck      *widget.Check
	recentProjectsList *widget.List
	statusLabel        *widget.Label
	launchButton       *widget.Button

	// 数据绑定
	selectedProject binding.String
	statusMessage   binding.String
}

// NewApp 创建新的GUI应用
func NewApp() *App {
	myApp := app.NewWithID("ai.launcher.app")
	myApp.SetIcon(resourceAppIconPng) // 需要创建图标资源

	// 设置应用元数据
	myApp.Metadata().Name = "AI启动器"
	myApp.Metadata().Icon = resourceAppIconPng

	return &App{
		fyneApp:         myApp,
		configManager:   project.NewConfigManager(),
		terminalManager: terminal.NewTerminalManager(),
		selectedProject: binding.NewString(),
		statusMessage:   binding.NewString(),
	}
}

// Run 启动GUI应用
func (a *App) Run() {
	// 加载配置
	if err := a.configManager.LoadProjects(); err != nil {
		log.Printf("加载配置失败: %v", err)
	}

	// 创建主窗口
	a.window = a.fyneApp.NewWindow("AI启动器 - 智能多AI工具启动器")
	a.window.SetIcon(resourceAppIconPng)
	a.window.Resize(fyne.NewSize(800, 600))
	a.window.SetFixedSize(false) // 允许调整大小
	a.window.CenterOnScreen()

	// 设置窗口关闭时的行为
	a.window.SetCloseIntercept(func() {
		a.fyneApp.Quit()
	})

	// 应用主题
	a.fyneApp.Settings().SetTheme(&customTheme{})

	// 创建界面
	content := a.createMainLayout()
	a.window.SetContent(content)

	// 初始化数据
	a.refreshRecentProjects()
	a.statusMessage.Set("准备就绪")

	// 显示窗口
	a.window.ShowAndRun()
}

// createMainLayout 创建主布局
func (a *App) createMainLayout() *fyne.Container {
	// 左侧面板 - 最近项目
	leftPanel := a.createRecentProjectsPanel()

	// 右侧面板 - 配置选项
	rightPanel := a.createConfigPanel()

	// 底部面板 - 状态和操作
	bottomPanel := a.createBottomPanel()

	// 主要内容区域
	mainContent := container.New(
		layout.NewBorderLayout(nil, bottomPanel, leftPanel, nil),
		leftPanel,
		rightPanel,
		bottomPanel,
	)

	return mainContent
}

// createRecentProjectsPanel 创建最近项目面板
func (a *App) createRecentProjectsPanel() *fyne.Container {
	// 标题
	title := widget.NewRichTextFromMarkdown("## 📁 最近项目")
	title.Wrapping = fyne.TextWrapWord

	// 最近项目列表
	a.recentProjectsList = widget.NewList(
		func() int {
			projects := a.configManager.GetRecentProjects(10)
			return len(projects)
		},
		func() fyne.CanvasObject {
			// 项目卡片模板 - 更美观的设计
			nameLabel := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
			nameLabel.Truncation = fyne.TextTruncateEllipsis

			pathLabel := widget.NewLabel("")
			pathLabel.TextStyle = fyne.TextStyle{}
			pathLabel.Truncation = fyne.TextTruncateEllipsis

			modelIcon := widget.NewLabel("")
			modelLabel := widget.NewLabel("")
			modelLabel.TextStyle = fyne.TextStyle{Italic: true}

			timeLabel := widget.NewLabel("")
			timeLabel.TextStyle = fyne.TextStyle{Italic: true}

			// 模式指示器
			modeLabel := widget.NewLabel("")

			card := container.NewVBox(
				// 第一行：项目名 + 模型图标和名称
				container.NewHBox(nameLabel, layout.NewSpacer(), modelIcon, modelLabel),
				// 第二行：路径
				container.NewHBox(widget.NewIcon(theme.FolderIcon()), pathLabel),
				// 第三行：时间 + 模式
				container.NewHBox(timeLabel, layout.NewSpacer(), modeLabel),
			)

			// 添加内边距和背景
			cardWithPadding := container.NewPadded(card)
			return cardWithPadding
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			projects := a.configManager.GetRecentProjects(10)
			if id >= len(projects) {
				return
			}

			proj := projects[id]
			card := obj.(*fyne.Container).Objects[0].(*fyne.Container)

			// 更新项目信息 - 匹配新的卡片结构
			// 第一行：项目名 + 模型信息
			headerRow := card.Objects[0].(*fyne.Container)
			nameLabel := headerRow.Objects[0].(*widget.Label)
			modelIcon := headerRow.Objects[2].(*widget.Label)
			modelLabel := headerRow.Objects[3].(*widget.Label)

			// 第二行：路径
			pathRow := card.Objects[1].(*fyne.Container)
			pathLabel := pathRow.Objects[1].(*widget.Label)

			// 第三行：时间和模式
			timeRow := card.Objects[2].(*fyne.Container)
			timeLabel := timeRow.Objects[0].(*widget.Label)
			modeLabel := timeRow.Objects[2].(*widget.Label)

			// 设置数据
			nameLabel.SetText(proj.Name)
			pathLabel.SetText(proj.Path)
			timeLabel.SetText(proj.LastUsed.Format("01-02 15:04"))

			modelIcon.SetText(proj.AIModel.GetIcon())
			modelLabel.SetText(proj.AIModel.String())

			// 设置模式指示
			if proj.YoloMode {
				modeLabel.SetText("🚀 YOLO")
			} else {
				modeLabel.SetText("🛡️ 普通")
			}
		},
	)

	// 设置选择处理
	a.recentProjectsList.OnSelected = func(id widget.ListItemID) {
		projects := a.configManager.GetRecentProjects(10)
		if id < len(projects) {
			proj := projects[id]
			a.loadProject(proj)
		}
	}

	// 新项目按钮
	newProjectBtn := widget.NewButtonWithIcon("新建项目", theme.FolderNewIcon(), func() {
		a.clearForm()
	})

	panel := container.NewVBox(
		title,
		widget.NewSeparator(),
		a.recentProjectsList,
		widget.NewSeparator(),
		newProjectBtn,
	)

	// 设置面板大小和边距
	panelWithPadding := container.NewPadded(panel)
	return container.NewBorder(nil, nil, nil, nil, panelWithPadding)
}

// createConfigPanel 创建配置面板
func (a *App) createConfigPanel() *fyne.Container {
	// 项目配置部分
	projectConfigTitle := widget.NewRichTextFromMarkdown("## ⚙️ 项目配置")

	// 项目路径
	a.projectPathEntry = widget.NewEntry()
	a.projectPathEntry.SetPlaceHolder("选择项目目录...")

	browseBtn := widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI) {
			if uri != nil {
				path := uri.Path()
				a.projectPathEntry.SetText(path)
				a.projectNameEntry.SetText(filepath.Base(path))
			}
		}, a.window)
	})

	projectPathRow := container.NewBorder(nil, nil, nil, browseBtn, a.projectPathEntry)

	// 项目名称
	a.projectNameEntry = widget.NewEntry()
	a.projectNameEntry.SetPlaceHolder("项目名称")

	// AI模型选择
	modelTitle := widget.NewRichTextFromMarkdown("## 🤖 AI模型")

	models := a.configManager.GetAvailableModels()
	modelOptions := make([]string, len(models))
	for i, model := range models {
		modelOptions[i] = fmt.Sprintf("%s %s", model.GetIcon(), model.String())
	}

	a.modelSelect = widget.NewSelect(modelOptions, func(value string) {
		// 模型选择回调
	})
	a.modelSelect.SetSelected(modelOptions[0]) // 默认选择第一个

	// YOLO模式
	yoloTitle := widget.NewRichTextFromMarkdown("## ⚡ 运行模式")

	a.yoloModeCheck = widget.NewCheck("启用YOLO模式 (跳过安全确认)", func(checked bool) {
		if checked {
			a.statusMessage.Set("⚠️ YOLO模式已启用 - 将跳过安全确认")
		} else {
			a.statusMessage.Set("🛡️ 普通模式 - 需要用户确认操作")
		}
	})

	// 模式说明卡片
	normalModeCard := container.NewVBox(
		widget.NewLabelWithStyle("🛡️ 普通模式", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel("• 需要用户确认重要操作"),
		widget.NewLabel("• 适合生产环境和重要项目"),
		widget.NewLabel("• 更加安全可靠"),
	)

	yoloModeCard := container.NewVBox(
		widget.NewLabelWithStyle("🚀 YOLO模式", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel("• 跳过大部分安全检查"),
		widget.NewLabel("• 自动执行AI建议的操作"),
		widget.NewLabel("• 适合实验和快速原型"),
	)

	modeDescription := container.NewHBox(
		normalModeCard,
		widget.NewSeparator(),
		yoloModeCard,
	)

	// 组装配置面板 - 使用Form容器提供更好的布局
	projectForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "项目路径", Widget: projectPathRow},
			{Text: "项目名称", Widget: a.projectNameEntry},
		},
	}

	configPanel := container.NewVBox(
		projectConfigTitle,
		projectForm,

		widget.NewSeparator(),
		modelTitle,
		container.NewPadded(a.modelSelect),

		widget.NewSeparator(),
		yoloTitle,
		container.NewPadded(a.yoloModeCheck),
		container.NewPadded(modeDescription),
	)

	return container.NewPadded(configPanel)
}

// createBottomPanel 创建底部面板
func (a *App) createBottomPanel() *fyne.Container {
	// 状态标签 - 带图标和颜色指示
	statusIcon := widget.NewIcon(theme.InfoIcon())
	a.statusLabel = widget.NewLabel("")
	a.statusLabel.Bind(a.statusMessage)

	statusRow := container.NewHBox(
		statusIcon,
		a.statusLabel,
	)

	// 主要动作按钮
	a.launchButton = widget.NewButtonWithIcon("🚀 启动AI工具", theme.MediaPlayIcon(), func() {
		a.launchAI()
	})
	a.launchButton.Importance = widget.HighImportance

	// 保存配置按钮
	saveBtn := widget.NewButtonWithIcon("💾 保存配置", theme.DocumentSaveIcon(), func() {
		a.saveCurrentConfig()
	})

	// 其他功能按钮
	settingsBtn := widget.NewButtonWithIcon("⚙️ 设置", theme.SettingsIcon(), func() {
		a.showSettings()
	})

	aboutBtn := widget.NewButtonWithIcon("📖 关于", theme.InfoIcon(), func() {
		a.showAbout()
	})

	refreshBtn := widget.NewButtonWithIcon("🔄 刷新", theme.ViewRefreshIcon(), func() {
		a.refreshRecentProjects()
		a.statusMessage.Set("项目列表已刷新")
	})

	exitBtn := widget.NewButtonWithIcon("❌ 退出", theme.CancelIcon(), func() {
		a.fyneApp.Quit()
	})

	// 按钮组 - 分为主要操作和辅助操作
	primaryButtons := container.NewHBox(
		a.launchButton,
		saveBtn,
	)

	secondaryButtons := container.NewHBox(
		refreshBtn,
		settingsBtn,
		aboutBtn,
		exitBtn,
	)

	buttonRow := container.NewHBox(
		primaryButtons,
		layout.NewSpacer(),
		secondaryButtons,
	)

	// 底部面板
	bottomPanel := container.NewVBox(
		widget.NewSeparator(),
		statusRow,
		buttonRow,
	)

	return container.NewPadded(bottomPanel)
}

// 辅助方法

// loadProject 加载项目到表单
func (a *App) loadProject(proj project.ProjectConfig) {
	a.projectPathEntry.SetText(proj.Path)
	a.projectNameEntry.SetText(proj.Name)
	a.yoloModeCheck.SetChecked(proj.YoloMode)

	// 设置模型选择
	models := a.configManager.GetAvailableModels()
	for i, model := range models {
		if model == proj.AIModel {
			modelText := fmt.Sprintf("%s %s", model.GetIcon(), model.String())
			a.modelSelect.SetSelected(modelText)
			break
		}
	}

	a.statusMessage.Set(fmt.Sprintf("已加载项目: %s", proj.Name))
}

// clearForm 清空表单
func (a *App) clearForm() {
	a.projectPathEntry.SetText("")
	a.projectNameEntry.SetText("")
	a.yoloModeCheck.SetChecked(false)
	a.modelSelect.SetSelectedIndex(0)
	a.statusMessage.Set("准备创建新项目")
}

// refreshRecentProjects 刷新最近项目列表
func (a *App) refreshRecentProjects() {
	a.recentProjectsList.Refresh()
}

// launchAI 启动AI工具
func (a *App) launchAI() {
	// 验证输入
	projectPath := a.projectPathEntry.Text
	projectName := a.projectNameEntry.Text

	if projectPath == "" {
		dialog.ShowError(fmt.Errorf("请选择项目目录"), a.window)
		return
	}

	if projectName == "" {
		projectName = filepath.Base(projectPath)
		a.projectNameEntry.SetText(projectName)
	}

	// 验证路径
	if err := a.configManager.ValidateProjectPath(projectPath); err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	// 获取选择的模型
	selectedIndex := a.modelSelect.SelectedIndex()
	models := a.configManager.GetAvailableModels()
	if selectedIndex < 0 || selectedIndex >= len(models) {
		dialog.ShowError(fmt.Errorf("请选择AI模型"), a.window)
		return
	}

	selectedModel := models[selectedIndex]
	yoloMode := a.yoloModeCheck.Checked

	// 保存项目配置
	projectConfig := project.ProjectConfig{
		Name:     projectName,
		Path:     projectPath,
		AIModel:  selectedModel,
		YoloMode: yoloMode,
	}

	if err := a.configManager.AddProject(projectConfig); err != nil {
		dialog.ShowError(fmt.Errorf("保存配置失败: %v", err), a.window)
		return
	}

	// 显示启动确认对话框
	confirmText := fmt.Sprintf(
		"将要启动：\n\n"+
			"📁 项目: %s\n"+
			"📍 路径: %s\n"+
			"🤖 模型: %s %s\n"+
			"⚡ 模式: %s\n\n"+
			"确认启动吗？",
		projectName,
		projectPath,
		selectedModel.GetIcon(),
		selectedModel.String(),
		map[bool]string{true: "🚀 YOLO模式", false: "🛡️ 普通模式"}[yoloMode],
	)

	dialog.ShowConfirm("确认启动", confirmText, func(confirmed bool) {
		if confirmed {
			a.executeAILaunch(projectConfig)
		}
	}, a.window)
}

// executeAILaunch 执行AI工具启动
func (a *App) executeAILaunch(config project.ProjectConfig) {
	a.statusMessage.Set("正在启动AI工具...")
	a.launchButton.SetText("启动中...")
	a.launchButton.Disable()

	// 创建终端配置
	terminalConfig := terminal.TerminalConfig{
		Type:       a.getTerminalType(config.AIModel),
		Name:       fmt.Sprintf("%s-%s", config.AIModel, config.Name),
		WorkingDir: config.Path,
		Command:    config.AIModel.GetCommand(config.YoloMode),
		YoloMode:   config.YoloMode,
	}

	// 启动终端（在后台goroutine中）
	go func() {
		err := a.terminalManager.StartTerminal(terminalConfig)

		// 更新UI（需要在主线程中）
		if err != nil {
			a.statusMessage.Set(fmt.Sprintf("启动失败: %v", err))
			dialog.ShowError(err, a.window)
		} else {
			a.statusMessage.Set("✅ AI工具启动成功！")
			// 刷新最近项目列表
			a.refreshRecentProjects()
		}

		a.launchButton.SetText("🚀 启动AI工具")
		a.launchButton.Enable()
	}()
}

// getTerminalType 获取终端类型
func (a *App) getTerminalType(model project.AIModelType) terminal.TerminalType {
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

// showSettings 显示设置对话框
func (a *App) showSettings() {
	dialog.ShowInformation("设置", "设置功能开发中...", a.window)
}

// saveCurrentConfig 保存当前配置
func (a *App) saveCurrentConfig() {
	projectPath := a.projectPathEntry.Text
	projectName := a.projectNameEntry.Text

	if projectPath == "" || projectName == "" {
		dialog.ShowError(fmt.Errorf("请填写完整的项目信息"), a.window)
		return
	}

	// 验证路径
	if err := a.configManager.ValidateProjectPath(projectPath); err != nil {
		dialog.ShowError(err, a.window)
		return
	}

	// 获取选择的模型
	selectedIndex := a.modelSelect.SelectedIndex()
	models := a.configManager.GetAvailableModels()
	if selectedIndex < 0 || selectedIndex >= len(models) {
		dialog.ShowError(fmt.Errorf("请选择AI模型"), a.window)
		return
	}

	selectedModel := models[selectedIndex]
	yoloMode := a.yoloModeCheck.Checked

	// 创建项目配置
	projectConfig := project.ProjectConfig{
		Name:     projectName,
		Path:     projectPath,
		AIModel:  selectedModel,
		YoloMode: yoloMode,
	}

	// 保存配置
	if err := a.configManager.AddProject(projectConfig); err != nil {
		dialog.ShowError(fmt.Errorf("保存配置失败: %v", err), a.window)
		return
	}

	a.statusMessage.Set("✅ 配置已保存")
	a.refreshRecentProjects()

	dialog.ShowInformation("配置保存成功",
		fmt.Sprintf("项目 '%s' 的配置已保存", projectName), a.window)
}

// showAbout 显示关于对话框
func (a *App) showAbout() {
	about := fmt.Sprintf(
		"AI启动器 v2.0.0\n\n"+
			"智能多AI工具启动器\n"+
			"支持 Claude Code、Gemini CLI、Codex\n\n"+
			"功能特性：\n"+
			"• 项目管理和快速切换\n"+
			"• 多种AI模型支持\n"+
			"• YOLO模式快速启动\n"+
			"• 图形化用户界面\n\n"+
			"开发：AI Assistant\n"+
			"基于：Go + Fyne GUI",
	)

	dialog.ShowInformation("关于 AI启动器", about, a.window)
}

// customTheme 自定义主题
type customTheme struct{}

func (t *customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNamePrimary:
		return color.RGBA{0, 122, 204, 255} // 蓝色主题
	case theme.ColorNameButton:
		return color.RGBA{0, 122, 204, 255}
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (t *customTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t *customTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *customTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}