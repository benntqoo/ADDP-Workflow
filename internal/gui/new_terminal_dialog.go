package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"ai-launcher/internal/project"
)

// NewTerminalDialog 新建终端对话框
type NewTerminalDialog struct {
	window          fyne.Window
	projectManager  *project.ConfigManager
	onTerminalRequested func(project.ProjectConfig, project.AIModelType, bool)

	// 弹窗组件
	dialog *dialog.CustomDialog

	// 表单组件
	modelSelect     *widget.RadioGroup
	projectSelect   *widget.RadioGroup
	browseButton    *widget.Button
	inheritCheck    *widget.Check
	backgroundCheck *widget.Check

	// 按钮
	launchButton *widget.Button
	cancelButton *widget.Button

	// 数据
	recentProjects  []project.ProjectConfig
	selectedProject *project.ProjectConfig
}

// NewNewTerminalDialog 创建新建终端对话框
func NewNewTerminalDialog(parent fyne.Window, pm *project.ConfigManager, onRequested func(project.ProjectConfig, project.AIModelType, bool)) *NewTerminalDialog {
	d := &NewTerminalDialog{
		window:             parent,
		projectManager:     pm,
		onTerminalRequested: onRequested,
	}

	d.initializeUI()
	return d
}

// initializeUI 初始化对话框UI
func (d *NewTerminalDialog) initializeUI() {
	// AI工具选择
	d.modelSelect = widget.NewRadioGroup([]string{
		"🤖 Claude Code    (推薦用於通用開發)",
		"💎 Gemini CLI     (推薦用於創意和分析)",
		"🔧 Codex          (推薦用於代碼生成)",
		"🔬 Aider          (推薦用於代碼重構)",
	}, d.onModelChanged)
	d.modelSelect.SetSelected(d.modelSelect.Options[0]) // 默认选择第一个

	// 项目选择
	d.updateProjectOptions()

	d.projectSelect = widget.NewRadioGroup([]string{}, d.onProjectChanged)

	// 浏览按钮
	d.browseButton = widget.NewButtonWithIcon("📂 瀏覽選擇其他項目...", theme.FolderOpenIcon(), d.onBrowseClicked)

	// 快速选项
	d.inheritCheck = widget.NewCheck("繼承項目設置 (YOLO模式等)", nil)
	d.inheritCheck.SetChecked(true)

	d.backgroundCheck = widget.NewCheck("在背景運行 (不切換到新標籤頁)", nil)

	// 按钮
	d.launchButton = widget.NewButtonWithIcon("🚀啟動", theme.MediaPlayIcon(), d.onLaunchClicked)
	d.launchButton.Importance = widget.HighImportance

	d.cancelButton = widget.NewButtonWithIcon("❌取消", theme.CancelIcon(), d.onCancelClicked)

	// 创建表单布局
	form := d.createFormLayout()

	// 按钮行
	buttonRow := container.NewHBox(
		d.launchButton,
		layout.NewSpacer(),
		d.cancelButton,
	)

	// 主要内容
	content := container.NewVBox(
		form,
		widget.NewSeparator(),
		buttonRow,
	)

	// 创建自定义弹窗
	d.dialog = dialog.NewCustom("➕ 新建終端", "", content, d.window)
	d.dialog.Resize(fyne.NewSize(450, 350))
}

// createFormLayout 创建表单布局
func (d *NewTerminalDialog) createFormLayout() fyne.CanvasObject {
	// AI工具选择区域
	modelSection := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 🤖 選擇 AI CLI 工具"),
		d.modelSelect,
	)

	// 项目选择区域
	projectSection := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 📁 目標項目 (終端標籤頁名稱將使用項目名稱+AI工具)"),
		d.projectSelect,
		d.browseButton,
	)

	// 快速选项区域
	optionsSection := container.NewVBox(
		widget.NewRichTextFromMarkdown("### ⚡ 快速選項"),
		d.inheritCheck,
		d.backgroundCheck,
	)

	// 滚动容器
	scroll := container.NewScroll(container.NewVBox(
		modelSection,
		widget.NewSeparator(),
		projectSection,
		widget.NewSeparator(),
		optionsSection,
	))

	return scroll
}

// Show 显示对话框
func (d *NewTerminalDialog) Show() {
	d.updateProjectOptions()
	d.updateButtonStates()
	d.dialog.Show()
}

// Hide 隐藏对话框
func (d *NewTerminalDialog) Hide() {
	d.dialog.Hide()
}

// updateProjectOptions 更新项目选项
func (d *NewTerminalDialog) updateProjectOptions() {
	// 获取最近项目
	d.recentProjects = d.projectManager.GetRecentProjects(5)

	// 构建项目选项
	options := []string{}
	if len(d.recentProjects) > 0 {
		// 默认选择第一个项目
		options = append(options, "● 當前項目: "+d.recentProjects[0].Name)
		d.selectedProject = &d.recentProjects[0]

		// 添加其他最近项目
		if len(d.recentProjects) > 1 {
			options = append(options, "○ 從左側歷史項目選擇")
			for i := 1; i < len(d.recentProjects); i++ {
				proj := d.recentProjects[i]
				options = append(options, "  • "+proj.Name)
			}
		}
	} else {
		options = append(options, "○ 無最近項目，請瀏覽選擇")
	}

	// 更新RadioGroup选项
	if d.projectSelect != nil {
		d.projectSelect.Options = options
		if len(options) > 0 {
			d.projectSelect.SetSelected(options[0])
		}
	}
}

// 事件处理方法

func (d *NewTerminalDialog) onModelChanged(selected string) {
	d.updateButtonStates()
}

func (d *NewTerminalDialog) onProjectChanged(selected string) {
	// 解析选中的项目
	if selected == "" {
		return
	}

	// 如果选择的是"当前项目"
	if len(d.recentProjects) > 0 && selected == d.projectSelect.Options[0] {
		d.selectedProject = &d.recentProjects[0]
		d.updateButtonStates()
		return
	}

	// 如果选择的是历史项目中的某一个
	for _, proj := range d.recentProjects {
		if selected == "  • "+proj.Name {
			d.selectedProject = &proj
			d.updateButtonStates()
			return
		}
	}

	d.updateButtonStates()
}

func (d *NewTerminalDialog) onBrowseClicked() {
	dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
		if err == nil && uri != nil {
			// 创建临时项目配置
			path := uri.Path()
			projectName := path[len(path)-1:]
			if projectName == "" {
				projectName = "新項目"
			}

			tempProject := project.ProjectConfig{
				Name:     projectName,
				Path:     path,
				AIModel:  project.ModelClaudeCode, // 默认模型
				YoloMode: true,                    // 默认YOLO模式
			}

			d.selectedProject = &tempProject
			d.updateButtonStates()

			// 更新项目选择显示
			d.projectSelect.SetSelected("● 選中項目: " + projectName)
		}
	}, d.window)
}

func (d *NewTerminalDialog) onLaunchClicked() {
	if d.selectedProject == nil {
		dialog.ShowError(fmt.Errorf("請選擇目標項目"), d.window)
		return
	}

	// 获取选择的AI模型
	aiModel := d.parseAIModel()
	if aiModel == "" {
		dialog.ShowError(fmt.Errorf("請選擇 AI CLI 工具"), d.window)
		return
	}

	// 应用继承设置
	project := *d.selectedProject
	if !d.inheritCheck.Checked {
		// 不继承设置时，使用默认配置
		project.YoloMode = false
	}

	// 更新AI模型
	project.AIModel = aiModel

	// 获取背景运行选项
	runInBackground := d.backgroundCheck.Checked

	// 触发回调
	if d.onTerminalRequested != nil {
		d.onTerminalRequested(project, aiModel, runInBackground)
	}

	d.Hide()
}

func (d *NewTerminalDialog) onCancelClicked() {
	d.Hide()
}

// 工具方法

func (d *NewTerminalDialog) parseAIModel() project.AIModelType {
	selected := d.modelSelect.Selected
	switch {
	case selected == d.modelSelect.Options[0]: // Claude Code
		return project.ModelClaudeCode
	case selected == d.modelSelect.Options[1]: // Gemini CLI
		return project.ModelGeminiCLI
	case selected == d.modelSelect.Options[2]: // Codex
		return project.ModelCodex
	case selected == d.modelSelect.Options[3]: // Aider
		return project.ModelAider
	default:
		return ""
	}
}

func (d *NewTerminalDialog) updateButtonStates() {
	// 检查是否可以启动
	canLaunch := d.modelSelect.Selected != "" && d.selectedProject != nil

	if canLaunch {
		d.launchButton.Enable()
	} else {
		d.launchButton.Disable()
	}
}