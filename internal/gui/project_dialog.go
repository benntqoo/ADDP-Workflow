package gui

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"ai-launcher/internal/project"
)

// ProjectConfigDialog 项目配置弹窗
type ProjectConfigDialog struct {
	window         fyne.Window
	projectManager *project.ConfigManager
	onConfigured   func(project.ProjectConfig, project.AIModelType)

	// 弹窗组件
	dialog *dialog.CustomDialog

	// 表单组件
	pathEntry     *widget.Entry
	browseButton  *widget.Button
	nameLabel     *widget.Label
	modelSelect   *widget.RadioGroup
	modeSelect    *widget.RadioGroup
	envStatus     *widget.RichText

	// 按钮
	launchButton *widget.Button
	saveButton   *widget.Button
	cancelButton *widget.Button

	// 状态
	selectedProject *project.ProjectConfig
}

// NewProjectConfigDialog 创建项目配置弹窗
func NewProjectConfigDialog(parent fyne.Window, pm *project.ConfigManager, onConfigured func(project.ProjectConfig, project.AIModelType)) *ProjectConfigDialog {
	d := &ProjectConfigDialog{
		window:         parent,
		projectManager: pm,
		onConfigured:   onConfigured,
	}

	d.initializeUI()
	return d
}

// initializeUI 初始化弹窗UI
func (d *ProjectConfigDialog) initializeUI() {
	// 项目路径选择
	d.pathEntry = widget.NewEntry()
	d.pathEntry.SetPlaceHolder("选择项目目录...")
	d.pathEntry.OnChanged = d.onPathChanged

	d.browseButton = widget.NewButtonWithIcon("", theme.FolderOpenIcon(), d.onBrowseClicked)

	pathRow := container.NewBorder(nil, nil, nil, d.browseButton, d.pathEntry)

	// 项目名称（自动从目录获取）
	d.nameLabel = widget.NewLabel("(自动从目录名称获取)")
	d.nameLabel.TextStyle = fyne.TextStyle{Italic: true}

	// AI工具选择
	d.modelSelect = widget.NewRadioGroup([]string{
		"🤖 Claude Code    (推荐用於通用開發)",
		"💎 Gemini CLI     (推荐用於創意和分析)",
		"🔧 Codex          (推荐用於代碼生成)",
		"🔬 Aider          (推荐用於代碼重構)",
	}, d.onModelChanged)
	d.modelSelect.SetSelected(d.modelSelect.Options[0]) // 默认选择第一个

	// 运行模式选择
	d.modeSelect = widget.NewRadioGroup([]string{
		"🛡️ 普通模式 (需要確認操作，更安全)",
		"⚡ YOLO模式 (跳過安全確認，快速開發)",
	}, d.onModeChanged)
	d.modeSelect.SetSelected(d.modeSelect.Options[1]) // 默认YOLO模式

	// 环境检测状态
	d.envStatus = widget.NewRichText()
	d.envStatus.Wrapping = fyne.TextWrapWord

	// 按钮
	d.launchButton = widget.NewButtonWithIcon("🚀啟動", theme.MediaPlayIcon(), d.onLaunchClicked)
	d.launchButton.Importance = widget.HighImportance

	d.saveButton = widget.NewButtonWithIcon("💾儲存", theme.DocumentSaveIcon(), d.onSaveClicked)

	d.cancelButton = widget.NewButtonWithIcon("❌取消", theme.CancelIcon(), d.onCancelClicked)

	// 创建表单布局
	form := d.createFormLayout(pathRow)

	// 按钮行
	buttonRow := container.NewHBox(
		d.launchButton,
		d.saveButton,
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
	d.dialog = dialog.NewCustom("📂 開啟/新建項目", "", content, d.window)
	d.dialog.Resize(fyne.NewSize(600, 450))
}

// createFormLayout 创建表单布局
func (d *ProjectConfigDialog) createFormLayout(pathRow *fyne.Container) fyne.CanvasObject {
	// 项目信息区域
	projectInfo := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 📁 項目信息"),
		container.NewVBox(
			widget.NewLabel("項目路徑:"),
			pathRow,
			widget.NewLabel("項目名稱:"),
			d.nameLabel,
		),
	)

	// AI工具选择区域
	modelInfo := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 🤖 選擇 AI CLI 工具"),
		d.modelSelect,
	)

	// 运行模式区域
	modeInfo := container.NewVBox(
		widget.NewRichTextFromMarkdown("### ⚡ 運行模式"),
		d.modeSelect,
	)

	// 环境检测区域
	envInfo := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 🔧 環境檢測 (自動掃描項目)"),
		d.envStatus,
	)

	// 滚动容器
	scroll := container.NewScroll(container.NewVBox(
		projectInfo,
		widget.NewSeparator(),
		modelInfo,
		widget.NewSeparator(),
		modeInfo,
		widget.NewSeparator(),
		envInfo,
	))

	return scroll
}

// Show 显示弹窗
func (d *ProjectConfigDialog) Show() {
	d.resetForm()
	d.dialog.Show()
}

// Hide 隐藏弹窗
func (d *ProjectConfigDialog) Hide() {
	d.dialog.Hide()
}

// resetForm 重置表单
func (d *ProjectConfigDialog) resetForm() {
	d.pathEntry.SetText("")
	d.nameLabel.SetText("(自動從目錄名稱獲取)")
	d.modelSelect.SetSelected(d.modelSelect.Options[0])
	d.modeSelect.SetSelected(d.modeSelect.Options[1])
	d.envStatus.ParseMarkdown("請先選擇項目目錄...")
	d.updateButtonStates()
}

// 事件处理方法

func (d *ProjectConfigDialog) onPathChanged(path string) {
	if path == "" {
		d.nameLabel.SetText("(自動從目錄名稱獲取)")
		d.envStatus.ParseMarkdown("請選擇項目目錄...")
		d.updateButtonStates()
		return
	}

	// 验证路径是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		d.nameLabel.SetText("(路徑不存在)")
		d.envStatus.ParseMarkdown("❌ 選擇的路徑不存在")
		d.updateButtonStates()
		return
	}

	// 设置项目名称
	projectName := filepath.Base(path)
	d.nameLabel.SetText(projectName)

	// 执行环境检测
	d.performEnvironmentDetection(path)
	d.updateButtonStates()
}

func (d *ProjectConfigDialog) onBrowseClicked() {
	dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
		if err == nil && uri != nil {
			d.pathEntry.SetText(uri.Path())
		}
	}, d.window)
}

func (d *ProjectConfigDialog) onModelChanged(selected string) {
	// AI工具选择变更
	d.updateButtonStates()
}

func (d *ProjectConfigDialog) onModeChanged(selected string) {
	// 运行模式变更
	d.updateButtonStates()
}

func (d *ProjectConfigDialog) onLaunchClicked() {
	config, model := d.buildProjectConfig()
	if config != nil && d.onConfigured != nil {
		d.onConfigured(*config, model)
		d.Hide()
	}
}

func (d *ProjectConfigDialog) onSaveClicked() {
	config, _ := d.buildProjectConfig()
	if config != nil {
		if err := d.projectManager.AddProject(*config); err != nil {
			dialog.ShowError(fmt.Errorf("保存失败: %v", err), d.window)
		} else {
			dialog.ShowInformation("保存成功", fmt.Sprintf("項目 '%s' 已保存到配置", config.Name), d.window)
		}
	}
}

func (d *ProjectConfigDialog) onCancelClicked() {
	d.Hide()
}

// 工具方法

func (d *ProjectConfigDialog) performEnvironmentDetection(path string) {
	var detections []string

	// 检测各种项目类型
	if d.fileExists(filepath.Join(path, "package.json")) {
		detections = append(detections, "✅ Node.js 項目 (檢測到 package.json)")
	}

	if d.fileExists(filepath.Join(path, "go.mod")) {
		detections = append(detections, "✅ Go 項目 (檢測到 go.mod)")
	}

	if d.fileExists(filepath.Join(path, "requirements.txt")) || d.fileExists(filepath.Join(path, "pyproject.toml")) {
		detections = append(detections, "✅ Python 項目 (檢測到依賴文件)")
	}

	if d.fileExists(filepath.Join(path, ".git")) {
		detections = append(detections, "✅ Git 初始化完成")
	} else {
		detections = append(detections, "⚠️ 未初始化 Git")
	}

	if d.fileExists(filepath.Join(path, "tsconfig.json")) {
		detections = append(detections, "✅ TypeScript 配置正確")
	}

	if !d.fileExists(filepath.Join(path, ".env")) {
		detections = append(detections, "⚠️ 缺少 .env 檔案")
	}

	if len(detections) == 0 {
		detections = append(detections, "📁 通用項目目錄")
	}

	// 更新环境状态显示
	statusText := ""
	for _, detection := range detections {
		statusText += detection + "\n"
	}
	d.envStatus.ParseMarkdown(statusText)
}

func (d *ProjectConfigDialog) fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (d *ProjectConfigDialog) buildProjectConfig() (*project.ProjectConfig, project.AIModelType) {
	path := d.pathEntry.Text
	if path == "" {
		dialog.ShowError(fmt.Errorf("請選擇項目目錄"), d.window)
		return nil, ""
	}

	// 验证路径
	if _, err := os.Stat(path); os.IsNotExist(err) {
		dialog.ShowError(fmt.Errorf("選擇的路徑不存在"), d.window)
		return nil, ""
	}

	projectName := filepath.Base(path)
	if projectName == "" || projectName == "." {
		dialog.ShowError(fmt.Errorf("無效的項目名稱"), d.window)
		return nil, ""
	}

	// 解析AI模型
	aiModel := d.parseAIModel()
	if aiModel == "" {
		dialog.ShowError(fmt.Errorf("請選擇 AI CLI 工具"), d.window)
		return nil, ""
	}

	// 解析运行模式
	yoloMode := d.parseRunMode()

	config := &project.ProjectConfig{
		Name:     projectName,
		Path:     path,
		AIModel:  aiModel,
		YoloMode: yoloMode,
	}

	return config, aiModel
}

func (d *ProjectConfigDialog) parseAIModel() project.AIModelType {
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

func (d *ProjectConfigDialog) parseRunMode() bool {
	selected := d.modeSelect.Selected
	return selected == d.modeSelect.Options[1] // YOLO模式
}

func (d *ProjectConfigDialog) updateButtonStates() {
	// 检查是否可以启动
	canLaunch := d.pathEntry.Text != "" &&
		d.modelSelect.Selected != "" &&
		d.modeSelect.Selected != ""

	d.launchButton.Enable()
	d.saveButton.Enable()

	if !canLaunch {
		d.launchButton.Disable()
		d.saveButton.Disable()
	}
}