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
	window             fyne.Window
	projectManager     *project.ConfigManager
	onTerminalRequested func(project.ProjectConfig, project.AIModelType, bool)

	// 对话框组件
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

// initializeUI 初始化对话框 UI
func (d *NewTerminalDialog) initializeUI() {
	// AI 工具选择（先创建控件，稍后再 SetSelected，避免回调早触发导致按钮为 nil）
	d.modelSelect = widget.NewRadioGroup([]string{
		"Claude Code（通用/推荐）",
		"Gemini CLI（分析/推荐）",
		"Codex（生成/推荐）",
		"Aider（重构/推荐）",
	}, d.onModelChanged)

	// 项目选择
	d.updateProjectOptions()
	d.projectSelect = widget.NewRadioGroup([]string{}, d.onProjectChanged)

	// 浏览按钮
	d.browseButton = widget.NewButtonWithIcon("浏览其他项目...", theme.FolderOpenIcon(), d.onBrowseClicked)
	d.inheritCheck = widget.NewCheck("继承项目设置（YOLO 模式等）", nil)
	inherit := d.inheritCheck; inherit.SetChecked(true)

	d.backgroundCheck = widget.NewCheck("在后台运行（不切换到新终端）", nil)

	// 按钮
	launch := widget.NewButtonWithIcon("启动", theme.MediaPlayIcon(), d.onLaunchClicked)
	launch.Importance = widget.HighImportance
	d.launchButton = launch

	d.cancelButton = widget.NewButtonWithIcon("取消", theme.CancelIcon(), d.onCancelClicked)

	// 设置默认选择（此时按钮已创建，避免 nil 回调崩溃）
	if len(d.modelSelect.Options) > 0 {
		d.modelSelect.SetSelected(d.modelSelect.Options[0])
	}

	// 创建表单布局
	form := d.createFormLayout()

	// 底部按钮行
	buttonRow := container.NewHBox(
		d.launchButton,
		layout.NewSpacer(),
		d.cancelButton,
	)

	// 主体内容
	content := container.NewVBox(
		form,
		widget.NewSeparator(),
		buttonRow,
	)

	// 创建自定义对话框
	d.dialog = dialog.NewCustom("新建终端", "", content, d.window)
	d.dialog.Resize(fyne.NewSize(450, 350))
}

// createFormLayout 创建表单布局
func (d *NewTerminalDialog) createFormLayout() fyne.CanvasObject {
	// AI 工具选择
	modelSection := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 选择 AI CLI 工具"),
		d.modelSelect,
	)

	// 项目选择
	projectSection := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 选择项目（从最近项目或手动选择）"),
		d.projectSelect,
		d.browseButton,
	)

	// 选项
	optionsSection := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 选项"),
		d.inheritCheck,
		d.backgroundCheck,
	)

	// 可滚动容器
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
		options = append(options, "• 当前项目: "+d.recentProjects[0].Name)
		d.selectedProject = &d.recentProjects[0]

		// 追加其他最近项目
		for i := 1; i < len(d.recentProjects); i++ {
			proj := d.recentProjects[i]
			options = append(options, "  • "+proj.Name)
		}
	} else {
		options = append(options, "• 暂无最近项目，请浏览选择")
	}

	// 更新 RadioGroup 选项
	if d.projectSelect != nil {
		d.projectSelect.Options = options
		if len(options) > 0 {
			d.projectSelect.SetSelected(options[0])
		}
	}
}

// 事件处理

func (d *NewTerminalDialog) onModelChanged(selected string) {
	d.updateButtonStates()
}

func (d *NewTerminalDialog) onProjectChanged(selected string) {
	if selected == "" {
		return
	}

	if len(d.recentProjects) > 0 && selected == d.projectSelect.Options[0] {
		d.selectedProject = &d.recentProjects[0]
		d.updateButtonStates()
		return
	}

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
			projectName := path
			if projectName == "" {
				projectName = "新项目"
			}

			tempProject := project.ProjectConfig{
				Name:     projectName,
				Path:     path,
				AIModel:  project.ModelClaudeCode,
				YoloMode: true,
			}

			d.selectedProject = &tempProject
			d.updateButtonStates()

			// 更新项目选择显示
			d.projectSelect.SetSelected("• 当前项目: " + projectName)
		}
	}, d.window)
}

func (d *NewTerminalDialog) onLaunchClicked() {
	if d.selectedProject == nil {
		dialog.ShowError(fmt.Errorf("请先选择项目"), d.window)
		return
	}

	aiModel := d.parseAIModel()
	if aiModel == "" {
		dialog.ShowError(fmt.Errorf("请选择 AI CLI 工具"), d.window)
		return
	}

	project := *d.selectedProject
	if !d.inheritCheck.Checked {
		project.YoloMode = false
	}
	project.AIModel = aiModel

	runInBackground := d.backgroundCheck.Checked
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
	case selected == d.modelSelect.Options[0]:
		return project.ModelClaudeCode
	case selected == d.modelSelect.Options[1]:
		return project.ModelGeminiCLI
	case selected == d.modelSelect.Options[2]:
		return project.ModelCodex
	case selected == d.modelSelect.Options[3]:
		return project.ModelAider
	default:
		return ""
	}
}

func (d *NewTerminalDialog) updateButtonStates() {
	// 组件可能尚未初始化，先做空指针防护
	if d.modelSelect == nil || d.launchButton == nil {
		return
	}
	canLaunch := d.modelSelect.Selected != "" && d.selectedProject != nil
	if canLaunch {
		d.launchButton.Enable()
	} else {
		d.launchButton.Disable()
	}
}
