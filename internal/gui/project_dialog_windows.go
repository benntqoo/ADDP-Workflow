//go:build windows

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

type ProjectConfigDialog struct {
    window         fyne.Window
    projectManager *project.ConfigManager
    onConfigured   func(project.ProjectConfig, project.AIModelType)

    dialog *dialog.CustomDialog

    pathEntry    *widget.Entry
    browseButton *widget.Button
    nameLabel    *widget.Label
    modelSelect  *widget.RadioGroup
    modeSelect   *widget.RadioGroup
    envStatus    *widget.RichText

    launchButton *widget.Button
    saveButton   *widget.Button
    cancelButton *widget.Button

    selectedProject *project.ProjectConfig
}

func NewProjectConfigDialog(parent fyne.Window, pm *project.ConfigManager, onConfigured func(project.ProjectConfig, project.AIModelType)) *ProjectConfigDialog {
    d := &ProjectConfigDialog{
        window:         parent,
        projectManager: pm,
        onConfigured:   onConfigured,
    }
    d.initializeUI()
    return d
}

func (d *ProjectConfigDialog) initializeUI() {
    d.pathEntry = widget.NewEntry()
    d.pathEntry.SetPlaceHolder("选择项目目录...")
    d.pathEntry.OnChanged = d.onPathChanged

    d.browseButton = widget.NewButtonWithIcon("浏览...", theme.FolderOpenIcon(), d.onBrowseClicked)
    pathRow := container.NewBorder(nil, nil, nil, d.browseButton, d.pathEntry)

    d.nameLabel = widget.NewLabel("(自动从目录名获取)")
    d.nameLabel.TextStyle = fyne.TextStyle{Italic: true}

    d.modelSelect = widget.NewRadioGroup([]string{
        "Claude Code（通用/推荐）",
        "Gemini CLI（分析/推荐）",
        "Codex（生成/推荐）",
        "Aider（重构/推荐）",
    }, d.onModelChanged)
    d.modelSelect.SetSelected(d.modelSelect.Options[0])

    d.modeSelect = widget.NewRadioGroup([]string{
        "普通模式（需要确认，更安全）",
        "YOLO 模式（跳过确认，速度优先）",
    }, d.onModeChanged)
    d.modeSelect.SetSelected(d.modeSelect.Options[1])

    d.envStatus = widget.NewRichText()
    d.envStatus.Wrapping = fyne.TextWrapWord

    d.launchButton = widget.NewButtonWithIcon("启动", theme.MediaPlayIcon(), d.onLaunchClicked)
    d.launchButton.Importance = widget.HighImportance
    d.saveButton = widget.NewButtonWithIcon("保存", theme.DocumentSaveIcon(), d.onSaveClicked)
    d.cancelButton = widget.NewButtonWithIcon("取消", theme.CancelIcon(), d.onCancelClicked)

    form := d.createFormLayout(pathRow)
    buttons := container.NewHBox(d.launchButton, d.saveButton, layout.NewSpacer(), d.cancelButton)
    content := container.NewVBox(form, widget.NewSeparator(), buttons)

    d.dialog = dialog.NewCustom("打开/新建项目", "", content, d.window)
    d.dialog.Resize(fyne.NewSize(600, 450))
}

func (d *ProjectConfigDialog) createFormLayout(pathRow *fyne.Container) fyne.CanvasObject {
    projectInfo := container.NewVBox(
        widget.NewRichTextFromMarkdown("### 项目路径"),
        pathRow,
        widget.NewSeparator(),
        widget.NewRichTextFromMarkdown("**项目名称**"),
        d.nameLabel,
    )

    selects := container.NewVBox(
        widget.NewRichTextFromMarkdown("### AI 工具与模式"),
        d.modelSelect,
        widget.NewSeparator(),
        d.modeSelect,
    )

    envInfo := container.NewVBox(
        widget.NewRichTextFromMarkdown("### 环境检测"),
        d.envStatus,
    )

    return container.NewScroll(container.NewVBox(
        projectInfo,
        widget.NewSeparator(),
        selects,
        widget.NewSeparator(),
        envInfo,
    ))
}

func (d *ProjectConfigDialog) Show() { d.resetForm(); d.dialog.Show() }
func (d *ProjectConfigDialog) Hide() { d.dialog.Hide() }

func (d *ProjectConfigDialog) resetForm() {
    d.pathEntry.SetText("")
    d.nameLabel.SetText("(自动从目录名获取)")
    d.modelSelect.SetSelected(d.modelSelect.Options[0])
    d.modeSelect.SetSelected(d.modeSelect.Options[1])
    d.envStatus.ParseMarkdown("请先选择项目目录...")
    d.updateButtonStates()
}

func (d *ProjectConfigDialog) onPathChanged(path string) {
    if path == "" {
        d.nameLabel.SetText("(自动从目录名获取)")
        d.envStatus.ParseMarkdown("请先选择项目目录...")
        d.updateButtonStates()
        return
    }
    if _, err := os.Stat(path); os.IsNotExist(err) {
        d.nameLabel.SetText("(路径不存在)")
        d.envStatus.ParseMarkdown("路径不存在，请重新选择")
        d.updateButtonStates()
        return
    }
    d.nameLabel.SetText(filepath.Base(path))
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

func (d *ProjectConfigDialog) onModelChanged(_ string) { d.updateButtonStates() }
func (d *ProjectConfigDialog) onModeChanged(_ string)  { d.updateButtonStates() }

func (d *ProjectConfigDialog) onLaunchClicked() {
    cfg, model := d.buildProjectConfig()
    if cfg != nil && d.onConfigured != nil {
        d.onConfigured(*cfg, model)
        d.Hide()
    }
}

func (d *ProjectConfigDialog) onSaveClicked() {
    cfg, _ := d.buildProjectConfig()
    if cfg != nil {
        if err := d.projectManager.AddProject(*cfg); err != nil {
            dialog.ShowError(fmt.Errorf("保存失败: %v", err), d.window)
        } else {
            dialog.ShowInformation("保存成功", fmt.Sprintf("项目 '%s' 已保存", cfg.Name), d.window)
        }
    }
}

func (d *ProjectConfigDialog) onCancelClicked() { d.Hide() }

func (d *ProjectConfigDialog) performEnvironmentDetection(path string) {
    var detections []string
    if d.fileExists(filepath.Join(path, "package.json")) { detections = append(detections, "Detected Node.js project") }
    if d.fileExists(filepath.Join(path, "go.mod")) { detections = append(detections, "Detected Go project") }
    if d.fileExists(filepath.Join(path, "requirements.txt")) || d.fileExists(filepath.Join(path, "pyproject.toml")) { detections = append(detections, "Detected Python project") }
    if d.fileExists(filepath.Join(path, ".git")) { detections = append(detections, "Detected Git repository") } else { detections = append(detections, "Git not initialized") }
    if d.fileExists(filepath.Join(path, "tsconfig.json")) { detections = append(detections, "Detected TypeScript config") }
    if !d.fileExists(filepath.Join(path, ".env")) { detections = append(detections, "Missing .env file") }
    if len(detections) == 0 { detections = append(detections, "Generic project folder") }
    statusText := ""
    for _, t := range detections { statusText += t + "\n" }
    d.envStatus.ParseMarkdown(statusText)
}

func (d *ProjectConfigDialog) fileExists(path string) bool { _, err := os.Stat(path); return !os.IsNotExist(err) }

func (d *ProjectConfigDialog) buildProjectConfig() (*project.ProjectConfig, project.AIModelType) {
    path := d.pathEntry.Text
    if path == "" {
        dialog.ShowError(fmt.Errorf("请选择项目目录"), d.window)
        return nil, ""
    }
    if _, err := os.Stat(path); os.IsNotExist(err) {
        dialog.ShowError(fmt.Errorf("路径不存在"), d.window)
        return nil, ""
    }
    projectName := filepath.Base(path)
    if projectName == "" || projectName == "." { dialog.ShowError(fmt.Errorf("项目名无效"), d.window); return nil, "" }

    aiModel := d.parseAIModel()
    if aiModel == "" { dialog.ShowError(fmt.Errorf("请选择 AI CLI 工具"), d.window); return nil, "" }

    yolo := d.parseRunMode()
    cfg := &project.ProjectConfig{ Name: projectName, Path: path, AIModel: aiModel, YoloMode: yolo }
    return cfg, aiModel
}

func (d *ProjectConfigDialog) parseAIModel() project.AIModelType {
    s := d.modelSelect.Selected
    switch {
    case s == d.modelSelect.Options[0]:
        return project.ModelClaudeCode
    case s == d.modelSelect.Options[1]:
        return project.ModelGeminiCLI
    case s == d.modelSelect.Options[2]:
        return project.ModelCodex
    case s == d.modelSelect.Options[3]:
        return project.ModelAider
    default:
        return ""
    }
}

func (d *ProjectConfigDialog) parseRunMode() bool {
    return d.modeSelect.Selected == d.modeSelect.Options[1]
}

func (d *ProjectConfigDialog) updateButtonStates() {
    can := d.pathEntry.Text != "" && d.modelSelect.Selected != "" && d.modeSelect.Selected != ""
    if can { d.launchButton.Enable(); d.saveButton.Enable() } else { d.launchButton.Disable(); d.saveButton.Disable() }
}

