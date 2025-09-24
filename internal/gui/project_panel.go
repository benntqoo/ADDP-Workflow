package gui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"ai-launcher/internal/project"
)

// ProjectHistoryPanel 椤圭洰鍘嗗彶闈㈡澘
type ProjectHistoryPanel struct {
	projectManager   *project.ConfigManager
	onProjectSelect  func(project.ProjectConfig)

	// UI缁勪欢
	container        *fyne.Container
	projectList      *widget.List
	refreshButton    *widget.Button
	autoRefreshCheck *widget.Check

	// 鐘舵€?
	projects         []project.ProjectConfig
	selectedProject  *project.ProjectConfig
}

// NewProjectHistoryPanel 鍒涘缓椤圭洰鍘嗗彶闈㈡澘
func NewProjectHistoryPanel(pm *project.ConfigManager, onSelect func(project.ProjectConfig)) *ProjectHistoryPanel {
	panel := &ProjectHistoryPanel{
		projectManager:  pm,
		onProjectSelect: onSelect,
		projects:        []project.ProjectConfig{},
	}

	panel.initializeUI()
	panel.refreshProjects()

	return panel
}

// initializeUI 鍒濆鍖朥I
func (p *ProjectHistoryPanel) initializeUI() {
	// 鏍囬
	title := widget.NewRichTextFromMarkdown("## 项目历史")
	title.Wrapping = fyne.TextWrapWord

	// 椤圭洰鍒楄〃
	p.projectList = widget.NewList(
		func() int {
			return len(p.projects)
		},
		func() fyne.CanvasObject {
			return p.createProjectCard()
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			p.updateProjectCard(id, obj)
		},
	)

	// 璁剧疆鍒楄〃閫夋嫨浜嬩欢
	p.projectList.OnSelected = func(id widget.ListItemID) {
		if id < len(p.projects) {
			p.selectedProject = &p.projects[id]
			if p.onProjectSelect != nil {
				p.onProjectSelect(p.projects[id])
			}
		}
	}

    // 自动刷新复选框
    p.autoRefreshCheck = widget.NewCheck("自动刷新", func(checked bool) {
        if checked {
            // TODO: 鍚姩鑷姩鍒锋柊瀹氭椂鍣?
        } else {
            // TODO: 鍋滄鑷姩鍒锋柊瀹氭椂鍣?
        }
    })
	p.autoRefreshCheck.SetChecked(true)

    // 快捷操作按钮：刷新
    p.refreshButton = widget.NewButtonWithIcon("刷新", theme.ViewRefreshIcon(), func() {
        p.refreshProjects()
    })

    // 快捷操作按钮：新建终端
    newTerminalBtn := widget.NewButtonWithIcon("新建终端", theme.ContentAddIcon(), func() {
        // TODO: 瑙﹀彂鏂板缓缁堢瀵硅瘽妗?
    })

    // 快捷操作按钮：设置
    settingsBtn := widget.NewButtonWithIcon("设置", theme.SettingsIcon(), func() {
        // TODO: 瑙﹀彂璁剧疆瀵硅瘽妗?
    })

    // 快捷操作按钮：监控
    monitorBtn := widget.NewButtonWithIcon("监控", theme.VisibilityIcon(), func() {
        // TODO: 瑙﹀彂鐩戞帶鐣岄潰
    })

    // 快捷操作区
    quickActions := container.NewVBox(
        widget.NewLabel("快捷操作"),
        newTerminalBtn,
        settingsBtn,
        monitorBtn,
    )

	// 涓诲鍣?
	p.container = container.NewVBox(
		title,
		widget.NewSeparator(),
		p.projectList,
		widget.NewSeparator(),
		p.autoRefreshCheck,
		p.refreshButton,
		widget.NewSeparator(),
		quickActions,
	)
}

// createProjectCard 鍒涘缓椤圭洰鍗＄墖妯℃澘
func (p *ProjectHistoryPanel) createProjectCard() fyne.CanvasObject {
	// 椤圭洰鍚嶇О锛堢矖浣擄級
	nameLabel := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	nameLabel.Truncation = fyne.TextTruncateEllipsis

	// AI宸ュ叿鏍囩
	aiLabel := widget.NewLabel("")
	aiLabel.TextStyle = fyne.TextStyle{Italic: true}

	// 杩愯妯″紡鏍囩
	modeLabel := widget.NewLabel("")

	// 鏈€鍚庝娇鐢ㄦ椂闂?
	timeLabel := widget.NewLabel("")
	timeLabel.TextStyle = fyne.TextStyle{Italic: true}

	// 鍗＄墖甯冨眬
	card := container.NewVBox(
		// 绗竴琛岋細椤圭洰鍚嶇О
		nameLabel,
		// 绗簩琛岋細AI宸ュ叿鍜屾ā寮?
		container.NewHBox(
			aiLabel,
			layout.NewSpacer(),
			modeLabel,
		),
		// 绗笁琛岋細鏃堕棿
		timeLabel,
		// 鍒嗛殧绾?
		widget.NewSeparator(),
	)

	// 娣诲姞鍐呰竟璺?
	return container.NewPadded(card)
}

// updateProjectCard 鏇存柊椤圭洰鍗＄墖鍐呭
func (p *ProjectHistoryPanel) updateProjectCard(id widget.ListItemID, obj fyne.CanvasObject) {
    if id >= len(p.projects) || obj == nil {
        return
    }

    proj := p.projects[id]

    // 防御式获取容器层级
    root, ok := obj.(*fyne.Container)
    if !ok || len(root.Objects) == 0 {
        return
    }
    card, ok := root.Objects[0].(*fyne.Container)
    if !ok || len(card.Objects) < 3 {
        return
    }

    nameLabel, _ := card.Objects[0].(*widget.Label)
    aiModeRow, _ := card.Objects[1].(*fyne.Container)
    var aiLabel, modeLabel *widget.Label
    if aiModeRow != nil && len(aiModeRow.Objects) >= 3 {
        aiLabel, _ = aiModeRow.Objects[0].(*widget.Label)
        modeLabel, _ = aiModeRow.Objects[2].(*widget.Label)
    }
    timeLabel, _ := card.Objects[2].(*widget.Label)

    if nameLabel != nil {
        nameLabel.SetText(fmt.Sprintf("• %s", proj.Name))
    }
    if aiLabel != nil {
        aiLabel.SetText(proj.AIModel.String())
    }
    if modeLabel != nil {
        if proj.YoloMode {
            modeLabel.SetText("⚡YOLO")
        } else {
            modeLabel.SetText("普通")
        }
    }
    if timeLabel != nil {
        timeLabel.SetText(p.formatRelativeTime(proj.LastUsed))
    }
}

// formatRelativeTime 鏍煎紡鍖栫浉瀵规椂闂?
func (p *ProjectHistoryPanel) formatRelativeTime(t time.Time) string {
    now := time.Now()
    diff := now.Sub(t)

    switch {
    case diff < time.Minute:
        return "刚刚"
    case diff < time.Hour:
        return fmt.Sprintf("%d 分钟前", int(diff.Minutes()))
    case diff < 24*time.Hour:
        return fmt.Sprintf("%d 小时前", int(diff.Hours()))
    case diff < 7*24*time.Hour:
        return fmt.Sprintf("%d 天前", int(diff.Hours()/24))
    default:
        return t.Format("01-02")
    }
}

// GetContainer 鑾峰彇瀹瑰櫒缁勪欢
func (p *ProjectHistoryPanel) GetContainer() *fyne.Container {
	return p.container
}

// Refresh 鍒锋柊椤圭洰鍒楄〃
func (p *ProjectHistoryPanel) Refresh() {
	p.refreshProjects()
}

// refreshProjects 浠庨厤缃鐞嗗櫒鍒锋柊椤圭洰鍒楄〃
func (p *ProjectHistoryPanel) refreshProjects() {
	// 鑾峰彇鏈€杩戦」鐩紙鏈€澶氭樉绀?0涓級
	p.projects = p.projectManager.GetRecentProjects(10)

	// 鍒锋柊鍒楄〃UI
	p.projectList.Refresh()
}

// GetSelectedProject 鑾峰彇褰撳墠閫変腑鐨勯」鐩?
func (p *ProjectHistoryPanel) GetSelectedProject() *project.ProjectConfig {
	return p.selectedProject
}

// SelectProject 绋嬪簭鍖栭€夋嫨椤圭洰
func (p *ProjectHistoryPanel) SelectProject(projectName string) {
	for i, proj := range p.projects {
		if proj.Name == projectName {
			p.projectList.Select(i)
			break
		}
	}
}

// ClearSelection 娓呴櫎閫夋嫨
func (p *ProjectHistoryPanel) ClearSelection() {
	p.projectList.UnselectAll()
	p.selectedProject = nil
}
