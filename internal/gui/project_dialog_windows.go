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
    d.pathEntry.SetPlaceHolder("闁瀚ㄦい鍦窗閻╊喖缍?..")
    d.pathEntry.OnChanged = d.onPathChanged

    d.browseButton = widget.NewButtonWithIcon("濞村繗顫?..", theme.FolderOpenIcon(), d.onBrowseClicked)
    pathRow := container.NewBorder(nil, nil, nil, d.browseButton, d.pathEntry)

    d.nameLabel = widget.NewLabel("(閼奉亜濮╂禒搴ｆ窗瑜版洖鎮曢懢宄板絿)")
    d.nameLabel.TextStyle = fyne.TextStyle{Italic: true}

    d.modelSelect = widget.NewRadioGroup([]string{
        "Claude Code閿涘牓鈧氨鏁?閹恒劏宕橀敍?,
        "Gemini CLI閿涘牆鍨庨弸?閹恒劏宕橀敍?,
        "Codex閿涘牏鏁撻幋?閹恒劏宕橀敍?,
        "Aider閿涘牓鍣搁弸?閹恒劏宕橀敍?,
    }, d.onModelChanged)\n    // 默认选中延后到控件完全创建后再设置\n    // 榛樿閫変腑寤跺悗鍒版帶浠跺畬鍏ㄥ垱寤哄悗鍐嶈缃?
    d.modeSelect = widget.NewRadioGroup([]string{
        "閺咁噣鈧碍膩瀵骏绱欓棁鈧憰浣衡€樼拋銈忕礉閺囨潙鐣ㄩ崗顭掔礆",
        "YOLO 濡€崇础閿涘牐鐑︽潻鍥┾€樼拋銈忕礉闁喎瀹虫导妯哄帥閿?,
    }, d.onModeChanged)\n    // 默认选中延后到控件完全创建后再设置\n    // 榛樿閫変腑寤跺悗鍒版帶浠跺畬鍏ㄥ垱寤哄悗鍐嶈缃?
    d.envStatus = widget.NewRichText()
    d.envStatus.Wrapping = fyne.TextWrapWord

    d.launchButton = widget.NewButtonWithIcon("閸氼垰濮?, theme.MediaPlayIcon(), d.onLaunchClicked)
    d.launchButton.Importance = widget.HighImportance
    d.saveButton = widget.NewButtonWithIcon("娣囨繂鐡?, theme.DocumentSaveIcon(), d.onSaveClicked)
    d.cancelButton = widget.NewButtonWithIcon("閸欐牗绉?, theme.CancelIcon(), d.onCancelClicked)

    form := d.createFormLayout(pathRow)
    buttons := container.NewHBox(d.launchButton, d.saveButton, layout.NewSpacer(), d.cancelButton)
    content := container.NewVBox(form, widget.NewSeparator(), buttons)

    d.dialog = dialog.NewCustom("閹垫挸绱?閺傛澘缂撴い鍦窗", "", content, d.window)
    \n    // 组件创建完成后再设置默认选项，避免回调访问未初始化的按钮\n    if len(d.modelSelect.Options) > 0 { d.modelSelect.SetSelected(d.modelSelect.Options[0]) }\n    if len(d.modeSelect.Options) > 1 { d.modeSelect.SetSelected(d.modeSelect.Options[1]) }\n    d.updateButtonStates()\n)
}

func (d *ProjectConfigDialog) createFormLayout(pathRow *fyne.Container) fyne.CanvasObject {
    projectInfo := container.NewVBox(
        widget.NewRichTextFromMarkdown("### 妞ゅ湱娲扮捄顖氱窞"),
        pathRow,
        widget.NewSeparator(),
        widget.NewRichTextFromMarkdown("**妞ゅ湱娲伴崥宥囆?*"),
        d.nameLabel,
    )

    selects := container.NewVBox(
        widget.NewRichTextFromMarkdown("### AI 瀹搞儱鍙挎稉搴⒛佸?),
        d.modelSelect,
        widget.NewSeparator(),
        d.modeSelect,
    )

    envInfo := container.NewVBox(
        widget.NewRichTextFromMarkdown("### 閻滎垰顣ㄥΛ鈧ù?),
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
    d.nameLabel.SetText("(閼奉亜濮╂禒搴ｆ窗瑜版洖鎮曢懢宄板絿)")\n    // 默认选中延后到控件完全创建后再设置\n    d.modeSelect.SetSelected(d.modeSelect.Options[1])
    d.envStatus.ParseMarkdown("鐠囧嘲鍘涢柅澶嬪妞ゅ湱娲伴惄顔肩秿...")
    d.updateButtonStates()
}

func (d *ProjectConfigDialog) onPathChanged(path string) {
    if path == "" {
        d.nameLabel.SetText("(閼奉亜濮╂禒搴ｆ窗瑜版洖鎮曢懢宄板絿)")
        d.envStatus.ParseMarkdown("鐠囧嘲鍘涢柅澶嬪妞ゅ湱娲伴惄顔肩秿...")
        d.updateButtonStates()
        return
    }
    if _, err := os.Stat(path); os.IsNotExist(err) {
        d.nameLabel.SetText("(鐠侯垰绶炴稉宥呯摠閸?")
        d.envStatus.ParseMarkdown("鐠侯垰绶炴稉宥呯摠閸︻煉绱濈拠鐑藉櫢閺備即鈧瀚?)
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
            dialog.ShowError(fmt.Errorf("娣囨繂鐡ㄦ径杈Е: %v", err), d.window)
        } else {
            dialog.ShowInformation("娣囨繂鐡ㄩ幋鎰", fmt.Sprintf("妞ゅ湱娲?'%s' 瀹歌弓绻氱€?, cfg.Name), d.window)
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
        dialog.ShowError(fmt.Errorf("鐠囩兘鈧瀚ㄦい鍦窗閻╊喖缍?), d.window)
        return nil, ""
    }
    if _, err := os.Stat(path); os.IsNotExist(err) {
        dialog.ShowError(fmt.Errorf("鐠侯垰绶炴稉宥呯摠閸?), d.window)
        return nil, ""
    }
    projectName := filepath.Base(path)
    if projectName == "" || projectName == "." { dialog.ShowError(fmt.Errorf("妞ゅ湱娲伴崥宥嗘￥閺?), d.window); return nil, "" }

    aiModel := d.parseAIModel()
    if aiModel == "" { dialog.ShowError(fmt.Errorf("鐠囩兘鈧瀚?AI CLI 瀹搞儱鍙?), d.window); return nil, "" }

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

func (d *ProjectConfigDialog) updateButtonStates() {\n    if d.launchButton == nil || d.saveButton == nil || d.modelSelect == nil || d.modeSelect == nil || d.pathEntry == nil { return }\n
    can := d.pathEntry.Text != "" && d.modelSelect.Selected != "" && d.modeSelect.Selected != ""
    if can { d.launchButton.Enable(); d.saveButton.Enable() } else { d.launchButton.Disable(); d.saveButton.Disable() }
}


