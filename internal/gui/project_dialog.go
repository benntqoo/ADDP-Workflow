//go:build !windows
// +build !windows

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

// ProjectConfigDialog 椤圭洰閰嶇疆寮圭獥
type ProjectConfigDialog struct {
	window         fyne.Window
	projectManager *project.ConfigManager
	onConfigured   func(project.ProjectConfig, project.AIModelType)

	// 寮圭獥缁勪欢
	dialog *dialog.CustomDialog

	// 琛ㄥ崟缁勪欢
	pathEntry     *widget.Entry
	browseButton  *widget.Button
	nameLabel     *widget.Label
	modelSelect   *widget.RadioGroup
	modeSelect    *widget.RadioGroup
	envStatus     *widget.RichText

	// 鎸夐挳
	launchButton *widget.Button
	saveButton   *widget.Button
	cancelButton *widget.Button

	// 鐘舵€?
	selectedProject *project.ProjectConfig
}

// NewProjectConfigDialog 鍒涘缓椤圭洰閰嶇疆寮圭獥
func NewProjectConfigDialog(parent fyne.Window, pm *project.ConfigManager, onConfigured func(project.ProjectConfig, project.AIModelType)) *ProjectConfigDialog {
	d := &ProjectConfigDialog{
		window:         parent,
		projectManager: pm,
		onConfigured:   onConfigured,
	}

	d.initializeUI()
	return d
}

// initializeUI 鍒濆鍖栧脊绐桿I
func (d *ProjectConfigDialog) initializeUI() {
	// 椤圭洰璺緞閫夋嫨
	d.pathEntry = widget.NewEntry()
	d.pathEntry.SetPlaceHolder("閫夋嫨椤圭洰鐩綍...")
	d.pathEntry.OnChanged = d.onPathChanged

	d.browseButton = widget.NewButtonWithIcon("", theme.FolderOpenIcon(), d.onBrowseClicked)

	pathRow := container.NewBorder(nil, nil, nil, d.browseButton, d.pathEntry)

	// 椤圭洰鍚嶇О锛堣嚜鍔ㄤ粠鐩綍鑾峰彇锛?
	d.nameLabel = widget.NewLabel("(鑷姩浠庣洰褰曞悕绉拌幏鍙?")
	d.nameLabel.TextStyle = fyne.TextStyle{Italic: true}

	// AI宸ュ叿閫夋嫨
	d.modelSelect = widget.NewRadioGroup([]string{
		"馃 Claude Code    (鎺ㄨ崘鐢ㄦ柤閫氱敤闁嬬櫦)",
		"馃拵 Gemini CLI     (鎺ㄨ崘鐢ㄦ柤鍓垫剰鍜屽垎鏋?",
		"馃敡 Codex          (鎺ㄨ崘鐢ㄦ柤浠ｇ⒓鐢熸垚)",
		"馃敩 Aider          (鎺ㄨ崘鐢ㄦ柤浠ｇ⒓閲嶆)",
	}, d.onModelChanged)
	d.modelSelect.SetSelected(d.modelSelect.Options[0]) // 榛樿閫夋嫨绗竴涓?

	// 杩愯妯″紡閫夋嫨
	d.modeSelect = widget.NewRadioGroup([]string{
		"馃洝锔?鏅€氭ā寮?(闇€瑕佺⒑瑾嶆搷浣滐紝鏇村畨鍏?",
		"鈿?YOLO妯″紡 (璺抽亷瀹夊叏纰鸿獚锛屽揩閫熼枊鐧?",
	}, d.onModeChanged)
	d.modeSelect.SetSelected(d.modeSelect.Options[1]) // 榛樿YOLO妯″紡

	// 鐜妫€娴嬬姸鎬?
	d.envStatus = widget.NewRichText()
	d.envStatus.Wrapping = fyne.TextWrapWord

	// 鎸夐挳
	d.launchButton = widget.NewButtonWithIcon("馃殌鍟熷嫊", theme.MediaPlayIcon(), d.onLaunchClicked)
	d.launchButton.Importance = widget.HighImportance

	d.saveButton = widget.NewButtonWithIcon("馃捑鍎插瓨", theme.DocumentSaveIcon(), d.onSaveClicked)

	d.cancelButton = widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), d.onCancelClicked)

	// 鍒涘缓琛ㄥ崟甯冨眬
	form := d.createFormLayout(pathRow)

	// 鎸夐挳琛?
	buttonRow := container.NewHBox(
		d.launchButton,
		d.saveButton,
		layout.NewSpacer(),
		d.cancelButton,
	)

	// 涓昏鍐呭
	content := container.NewVBox(
		form,
		widget.NewSeparator(),
		buttonRow,
	)

	// 鍒涘缓鑷畾涔夊脊绐?
	d.dialog = dialog.NewCustom("馃搨 闁嬪暉/鏂板缓闋呯洰", "", content, d.window)
	d.dialog.Resize(fyne.NewSize(600, 450))
}

// createFormLayout 鍒涘缓琛ㄥ崟甯冨眬
func (d *ProjectConfigDialog) createFormLayout(pathRow *fyne.Container) fyne.CanvasObject {
	// 椤圭洰淇℃伅鍖哄煙
	projectInfo := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 馃搧 闋呯洰淇℃伅"),
		container.NewVBox(
			widget.NewLabel("闋呯洰璺緫:"),
			pathRow,
			widget.NewLabel("闋呯洰鍚嶇ū:"),
			d.nameLabel,
		),
	)

	// AI宸ュ叿閫夋嫨鍖哄煙
	modelInfo := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 馃 閬告搰 AI CLI 宸ュ叿"),
		d.modelSelect,
	)

	// 杩愯妯″紡鍖哄煙
	modeInfo := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 鈿?閬嬭妯″紡"),
		d.modeSelect,
	)

	// 鐜妫€娴嬪尯鍩?
	envInfo := container.NewVBox(
		widget.NewRichTextFromMarkdown("### 馃敡 鐠板妾㈡脯 (鑷嫊鎺冩弿闋呯洰)"),
		d.envStatus,
	)

	// 婊氬姩瀹瑰櫒
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

// Show 鏄剧ず寮圭獥
func (d *ProjectConfigDialog) Show() {
	d.resetForm()
	d.dialog.Show()
}

// Hide 闅愯棌寮圭獥
func (d *ProjectConfigDialog) Hide() {
	d.dialog.Hide()
}

// resetForm 閲嶇疆琛ㄥ崟
func (d *ProjectConfigDialog) resetForm() {
	d.pathEntry.SetText("")
	d.nameLabel.SetText("(鑷嫊寰炵洰閷勫悕绋辩嵅鍙?")
	d.modelSelect.SetSelected(d.modelSelect.Options[0])
	d.modeSelect.SetSelected(d.modeSelect.Options[1])
	d.envStatus.ParseMarkdown("璜嬪厛閬告搰闋呯洰鐩寗...")
	d.updateButtonStates()
}

// 浜嬩欢澶勭悊鏂规硶

func (d *ProjectConfigDialog) onPathChanged(path string) {
	if path == "" {
		d.nameLabel.SetText("(鑷嫊寰炵洰閷勫悕绋辩嵅鍙?")
		d.envStatus.ParseMarkdown("璜嬮伕鎿囬爡鐩洰閷?..")
		d.updateButtonStates()
		return
	}

	// 楠岃瘉璺緞鏄惁瀛樺湪
	if _, err := os.Stat(path); os.IsNotExist(err) {
		d.nameLabel.SetText("(璺緫涓嶅瓨鍦?")
		d.envStatus.ParseMarkdown("鉂?閬告搰鐨勮矾寰戜笉瀛樺湪")
		d.updateButtonStates()
		return
	}

	// 璁剧疆椤圭洰鍚嶇О
	projectName := filepath.Base(path)
	d.nameLabel.SetText(projectName)

	// 鎵ц鐜妫€娴?
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
	// AI宸ュ叿閫夋嫨鍙樻洿
	d.updateButtonStates()
}

func (d *ProjectConfigDialog) onModeChanged(selected string) {
	// 杩愯妯″紡鍙樻洿
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
			dialog.ShowError(fmt.Errorf("淇濆瓨澶辫触: %v", err), d.window)
		} else {
			dialog.ShowInformation("淇濆瓨鎴愬姛", fmt.Sprintf("闋呯洰 '%s' 宸蹭繚瀛樺埌閰嶇疆", config.Name), d.window)
		}
	}
}

func (d *ProjectConfigDialog) onCancelClicked() {
	d.Hide()
}

// 宸ュ叿鏂规硶

func (d *ProjectConfigDialog) performEnvironmentDetection(path string) {
	var detections []string

	if d.fileExists(filepath.Join(path, "package.json")) {
		detections = append(detections, "Detected Node.js project (package.json)")
	}
	if d.fileExists(filepath.Join(path, "go.mod")) {
		detections = append(detections, "Detected Go project (go.mod)")
	}
	if d.fileExists(filepath.Join(path, "requirements.txt")) || d.fileExists(filepath.Join(path, "pyproject.toml")) {
		detections = append(detections, "Detected Python project (requirements/pyproject)")
	}
	if d.fileExists(filepath.Join(path, ".git")) {
		detections = append(detections, "Detected Git repository")
	} else {
		detections = append(detections, "Git not initialized")
	}
	if d.fileExists(filepath.Join(path, "tsconfig.json")) {
		detections = append(detections, "Detected TypeScript config")
	}
	if !d.fileExists(filepath.Join(path, ".env")) {
		detections = append(detections, "Missing .env file")
	}
	if len(detections) == 0 {
		detections = append(detections, "Generic project folder")
	}

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
		dialog.ShowError(fmt.Errorf("Please select a project folder"), d.window)
		return nil, ""
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		dialog.ShowError(fmt.Errorf("Selected path does not exist"), d.window)
		return nil, ""
	}
	projectName := filepath.Base(path)
	if projectName == "" || projectName == "." {
		dialog.ShowError(fmt.Errorf("Invalid project name"), d.window)
		return nil, ""
	}
	aiModel := d.parseAIModel()
	if aiModel == "" {
		dialog.ShowError(fmt.Errorf("Please select an AI CLI tool"), d.window)
		return nil, ""
	}
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

func (d *ProjectConfigDialog) parseRunMode() bool {
	selected := d.modeSelect.Selected
	return selected == d.modeSelect.Options[1]
}

func (d *ProjectConfigDialog) updateButtonStates() {
	if d.pathEntry == nil || d.modelSelect == nil || d.modeSelect == nil || d.launchButton == nil || d.saveButton == nil {
		return
	}
	canLaunch := d.pathEntry.Text != "" && d.modelSelect.Selected != "" && d.modeSelect.Selected != ""
	d.launchButton.Enable()
	d.saveButton.Enable()
	if !canLaunch {
		d.launchButton.Disable()
		d.saveButton.Disable()
	}
}

