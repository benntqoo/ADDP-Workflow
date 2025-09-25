package gui

import (
    "fmt"
    "path/filepath"
    "log"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"

    "ai-launcher/internal/project"
)

// NewTerminalDialog 新建终端对话框（目录 + AI CLI + YOLO）
type NewTerminalDialog struct {
    window              fyne.Window
    projectManager      *project.ConfigManager
    onTerminalRequested func(project.ProjectConfig, project.AIModelType, bool)

    // 对话框
    dialog *dialog.CustomDialog

    // 表单控件
    pathEntry   *widget.Entry
    browseBtn   *widget.Button
    modelSelect *widget.RadioGroup
    yoloCheck   *widget.Check

    // 按钮
    launchButton *widget.Button
    cancelButton *widget.Button
}

// NewNewTerminalDialog 创建新建终端对话框
func NewNewTerminalDialog(parent fyne.Window, pm *project.ConfigManager, onRequested func(project.ProjectConfig, project.AIModelType, bool)) *NewTerminalDialog {
    d := &NewTerminalDialog{
        window:              parent,
        projectManager:      pm,
        onTerminalRequested: onRequested,
    }
    d.initializeUI()
    return d
}

func (d *NewTerminalDialog) initializeUI() {
    // 目录选择
    d.pathEntry = widget.NewEntry()
    d.pathEntry.SetPlaceHolder("选择要开启的目录...")
    d.browseBtn = widget.NewButtonWithIcon("浏览...", theme.FolderOpenIcon(), d.onBrowseClicked)
    // 路径变更即刷新按钮状态
    d.pathEntry.OnChanged = func(string){ d.updateButtonStates() }
    // 使用两列网格避免 Border 在某些主题下高度计算为 0 的问题
    pathRow := container.NewGridWithColumns(2, d.pathEntry, d.browseBtn)

    // AI CLI 选择
    d.modelSelect = widget.NewRadioGroup([]string{
        "Claude Code（通用/推荐）",
        "Gemini CLI（分析/推荐）",
        "Codex（生成/推荐）",
        "Aider（重构/推荐）",
    }, func(string) { d.updateButtonStates() })
    if len(d.modelSelect.Options) > 0 {
        d.modelSelect.SetSelected(d.modelSelect.Options[0])
    }

    // YOLO
    d.yoloCheck = widget.NewCheck("YOLO 模式（跳过确认，速度优先）", nil)
    d.yoloCheck.SetChecked(true)

    // 底部按钮
    d.launchButton = widget.NewButtonWithIcon("确定", theme.ConfirmIcon(), d.onConfirmClicked)
    d.launchButton.Importance = widget.HighImportance
    d.cancelButton = widget.NewButtonWithIcon("取消", theme.CancelIcon(), d.onCancelClicked)

    // 表单
    form := container.NewVBox(
        widget.NewRichTextFromMarkdown("### 开启目录"),
        pathRow,
        widget.NewSeparator(),
        widget.NewRichTextFromMarkdown("### 选择 AI CLI 工具"),
        d.modelSelect,
        widget.NewSeparator(),
        d.yoloCheck,
    )
    // 右对齐按钮，去掉中间空位
    buttons := container.NewHBox(layout.NewSpacer(), d.launchButton, d.cancelButton)
    content := container.NewVBox(form, widget.NewSeparator(), buttons)

    d.dialog = dialog.NewCustom("新建终端", "", content, d.window)
    d.dialog.Resize(fyne.NewSize(600, 420))
    d.updateButtonStates()
}

func (d *NewTerminalDialog) Show()  { d.reset(); d.dialog.Show() }
func (d *NewTerminalDialog) Hide()  { d.dialog.Hide() }

func (d *NewTerminalDialog) reset() {
    d.pathEntry.SetText("")
    if len(d.modelSelect.Options) > 0 {
        d.modelSelect.SetSelected(d.modelSelect.Options[0])
    }
    d.yoloCheck.SetChecked(true)
    d.updateButtonStates()
}

func (d *NewTerminalDialog) onBrowseClicked() {
    dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
        if err == nil && uri != nil {
            d.pathEntry.SetText(uri.Path())
            d.updateButtonStates()
        }
    }, d.window)
}

func (d *NewTerminalDialog) onConfirmClicked() {
    path := d.pathEntry.Text
    if path == "" {
        dialog.ShowError(fmt.Errorf("请选择项目目录"), d.window)
        return
    }
    aiModel := d.parseAIModel()
    if aiModel == "" {
        dialog.ShowError(fmt.Errorf("请选择 AI CLI 工具"), d.window)
        return
    }

    proj := project.ProjectConfig{
        Name:     filepath.Base(path),
        Path:     path,
        AIModel:  aiModel,
        YoloMode: d.yoloCheck.Checked,
    }

    log.Printf("[NewTerminalDialog] confirm path=%s model=%s yolo=%t", proj.Path, proj.AIModel, proj.YoloMode)
    // 先关闭对话框，避免遮罩未关闭造成界面看似“无响应”
    d.Hide()
    if d.onTerminalRequested != nil {
        d.onTerminalRequested(proj, aiModel, false)
    } else {
        log.Printf("[NewTerminalDialog] onTerminalRequested is nil")
    }
}

func (d *NewTerminalDialog) onCancelClicked() { d.Hide() }

func (d *NewTerminalDialog) parseAIModel() project.AIModelType {
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

func (d *NewTerminalDialog) updateButtonStates() {
    if d.launchButton == nil || d.modelSelect == nil || d.pathEntry == nil { return }
    can := d.modelSelect.Selected != "" && d.pathEntry.Text != ""
    if can { d.launchButton.Enable() } else { d.launchButton.Disable() }
}
