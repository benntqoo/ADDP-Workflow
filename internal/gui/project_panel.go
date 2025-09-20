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

// ProjectHistoryPanel 项目历史面板
type ProjectHistoryPanel struct {
	projectManager   *project.ConfigManager
	onProjectSelect  func(project.ProjectConfig)

	// UI组件
	container        *fyne.Container
	projectList      *widget.List
	refreshButton    *widget.Button
	autoRefreshCheck *widget.Check

	// 状态
	projects         []project.ProjectConfig
	selectedProject  *project.ProjectConfig
}

// NewProjectHistoryPanel 创建项目历史面板
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

// initializeUI 初始化UI
func (p *ProjectHistoryPanel) initializeUI() {
	// 标题
	title := widget.NewRichTextFromMarkdown("## 📂 项目历史")
	title.Wrapping = fyne.TextWrapWord

	// 项目列表
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

	// 设置列表选择事件
	p.projectList.OnSelected = func(id widget.ListItemID) {
		if id < len(p.projects) {
			p.selectedProject = &p.projects[id]
			if p.onProjectSelect != nil {
				p.onProjectSelect(p.projects[id])
			}
		}
	}

	// 自动刷新复选框
	p.autoRefreshCheck = widget.NewCheck("🔄 自动刷新", func(checked bool) {
		if checked {
			// TODO: 启动自动刷新定时器
		} else {
			// TODO: 停止自动刷新定时器
		}
	})
	p.autoRefreshCheck.SetChecked(true)

	// 快速操作按钮
	p.refreshButton = widget.NewButtonWithIcon("刷新", theme.ViewRefreshIcon(), func() {
		p.refreshProjects()
	})

	newTerminalBtn := widget.NewButtonWithIcon("🚀 新终端", theme.ContentAddIcon(), func() {
		// TODO: 触发新建终端对话框
	})

	settingsBtn := widget.NewButtonWithIcon("⚙️ 设置", theme.SettingsIcon(), func() {
		// TODO: 触发设置对话框
	})

	monitorBtn := widget.NewButtonWithIcon("📊 监控", theme.VisibilityIcon(), func() {
		// TODO: 触发监控界面
	})

	// 快速操作区域
	quickActions := container.NewVBox(
		widget.NewLabel("快速操作"),
		newTerminalBtn,
		settingsBtn,
		monitorBtn,
	)

	// 主容器
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

// createProjectCard 创建项目卡片模板
func (p *ProjectHistoryPanel) createProjectCard() fyne.CanvasObject {
	// 项目名称（粗体）
	nameLabel := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	nameLabel.Truncation = fyne.TextTruncateEllipsis

	// AI工具标签
	aiLabel := widget.NewLabel("")
	aiLabel.TextStyle = fyne.TextStyle{Italic: true}

	// 运行模式标签
	modeLabel := widget.NewLabel("")

	// 最后使用时间
	timeLabel := widget.NewLabel("")
	timeLabel.TextStyle = fyne.TextStyle{Italic: true}

	// 卡片布局
	card := container.NewVBox(
		// 第一行：项目名称
		nameLabel,
		// 第二行：AI工具和模式
		container.NewHBox(
			aiLabel,
			layout.NewSpacer(),
			modeLabel,
		),
		// 第三行：时间
		timeLabel,
		// 分隔线
		widget.NewSeparator(),
	)

	// 添加内边距
	return container.NewPadded(card)
}

// updateProjectCard 更新项目卡片内容
func (p *ProjectHistoryPanel) updateProjectCard(id widget.ListItemID, obj fyne.CanvasObject) {
	if id >= len(p.projects) {
		return
	}

	proj := p.projects[id]
	card := obj.(*fyne.Container).Objects[0].(*fyne.Container)

	// 获取标签组件
	nameLabel := card.Objects[0].(*widget.Label)
	aiModeRow := card.Objects[1].(*fyne.Container)
	aiLabel := aiModeRow.Objects[0].(*widget.Label)
	modeLabel := aiModeRow.Objects[2].(*widget.Label)
	timeLabel := card.Objects[2].(*widget.Label)

	// 设置项目名称
	nameLabel.SetText(fmt.Sprintf("• %s", proj.Name))

	// 设置AI工具
	aiLabel.SetText(proj.AIModel.String())

	// 设置运行模式
	if proj.YoloMode {
		modeLabel.SetText("⚡YOLO")
	} else {
		modeLabel.SetText("🛡️普通")
	}

	// 设置时间（相对时间）
	timeLabel.SetText(p.formatRelativeTime(proj.LastUsed))
}

// formatRelativeTime 格式化相对时间
func (p *ProjectHistoryPanel) formatRelativeTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	switch {
	case diff < time.Minute:
		return "刚刚"
	case diff < time.Hour:
		return fmt.Sprintf("%d分钟前", int(diff.Minutes()))
	case diff < 24*time.Hour:
		return fmt.Sprintf("%d小时前", int(diff.Hours()))
	case diff < 7*24*time.Hour:
		return fmt.Sprintf("%d天前", int(diff.Hours()/24))
	default:
		return t.Format("01-02")
	}
}

// GetContainer 获取容器组件
func (p *ProjectHistoryPanel) GetContainer() *fyne.Container {
	return p.container
}

// Refresh 刷新项目列表
func (p *ProjectHistoryPanel) Refresh() {
	p.refreshProjects()
}

// refreshProjects 从配置管理器刷新项目列表
func (p *ProjectHistoryPanel) refreshProjects() {
	// 获取最近项目（最多显示10个）
	p.projects = p.projectManager.GetRecentProjects(10)

	// 刷新列表UI
	p.projectList.Refresh()
}

// GetSelectedProject 获取当前选中的项目
func (p *ProjectHistoryPanel) GetSelectedProject() *project.ProjectConfig {
	return p.selectedProject
}

// SelectProject 程序化选择项目
func (p *ProjectHistoryPanel) SelectProject(projectName string) {
	for i, proj := range p.projects {
		if proj.Name == projectName {
			p.projectList.Select(i)
			break
		}
	}
}

// ClearSelection 清除选择
func (p *ProjectHistoryPanel) ClearSelection() {
	p.projectList.UnselectAll()
	p.selectedProject = nil
}