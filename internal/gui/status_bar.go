package gui

import (
	"fmt"
	"runtime"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// StatusBar 状态栏组件
type StatusBar struct {
	// UI组件
	container    *fyne.Container
	messageLabel *widget.Label
	statusLabel  *widget.Label
	timeLabel    *widget.Label
	cpuLabel     *widget.Label
	memoryLabel  *widget.Label
	networkLabel *widget.Label
	ollamaLabel  *widget.Label

	// 状态数据
	currentMessage string
	isRunning      bool
	startTime      time.Time
}

// NewStatusBar 创建状态栏
func NewStatusBar() *StatusBar {
	sb := &StatusBar{
		startTime: time.Now(),
		isRunning: true,
	}

	sb.initializeUI()
	sb.startStatusUpdater()

	return sb
}

// initializeUI 初始化状态栏UI
func (sb *StatusBar) initializeUI() {
	// 主状态消息
	sb.messageLabel = widget.NewLabel("准备就绪")
	sb.messageLabel.TextStyle = fyne.TextStyle{Bold: true}

	// 运行状态指示器
	sb.statusLabel = widget.NewLabel("🟢 AI启动器: 运行中")

	// 运行时间
	sb.timeLabel = widget.NewLabel("🕐 运行时间: 00:00:00")

	// 系统资源监控
	sb.cpuLabel = widget.NewLabel("📊 CPU: 0%")
	sb.memoryLabel = widget.NewLabel("💾 内存: 0MB")
	sb.networkLabel = widget.NewLabel("🌐 网路: 未知")
	sb.ollamaLabel = widget.NewLabel("🤖 Ollama: 未连接")

	// 第一行：主要状态信息
	topRow := container.NewHBox(
		sb.messageLabel,
		layout.NewSpacer(),
		sb.statusLabel,
		widget.NewSeparator(),
		sb.timeLabel,
	)

	// 第二行：系统监控信息
	bottomRow := container.NewHBox(
		sb.cpuLabel,
		widget.NewSeparator(),
		sb.memoryLabel,
		widget.NewSeparator(),
		sb.networkLabel,
		widget.NewSeparator(),
		sb.ollamaLabel,
	)

	// 状态栏容器
	sb.container = container.NewVBox(
		widget.NewSeparator(),
		topRow,
		bottomRow,
	)
}

// GetContainer 获取状态栏容器
func (sb *StatusBar) GetContainer() *fyne.Container {
	return sb.container
}

// SetMessage 设置状态消息
func (sb *StatusBar) SetMessage(message string) {
	sb.currentMessage = message
	sb.messageLabel.SetText(message)
}

// SetRunningStatus 设置运行状态
func (sb *StatusBar) SetRunningStatus(running bool, component string) {
	sb.isRunning = running
	if running {
		sb.statusLabel.SetText(fmt.Sprintf("🟢 %s: 运行中", component))
	} else {
		sb.statusLabel.SetText(fmt.Sprintf("🔴 %s: 已停止", component))
	}
}

// startStatusUpdater 启动状态更新器
func (sb *StatusBar) startStatusUpdater() {
	// 启动定时更新协程
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			sb.updateSystemStatus()
		}
	}()
}

// updateSystemStatus 更新系统状态
func (sb *StatusBar) updateSystemStatus() {
	// 更新运行时间
	duration := time.Since(sb.startTime)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	sb.timeLabel.SetText(fmt.Sprintf("🕐 运行时间: %02d:%02d:%02d", hours, minutes, seconds))

	// 更新CPU使用率（简化版本）
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// 计算内存使用量（MB）
	memoryMB := m.Alloc / 1024 / 1024
	sb.memoryLabel.SetText(fmt.Sprintf("💾 内存: %dMB", memoryMB))

	// CPU使用率（模拟数据，实际应该使用系统API）
	cpuPercent := sb.getCPUUsage()
	sb.cpuLabel.SetText(fmt.Sprintf("📊 CPU: %d%%", cpuPercent))

	// 网络状态（模拟检测）
	networkStatus := sb.getNetworkStatus()
	sb.networkLabel.SetText(fmt.Sprintf("🌐 网路: %s", networkStatus))

	// Ollama状态（模拟检测）
	ollamaStatus := sb.getOllamaStatus()
	sb.ollamaLabel.SetText(fmt.Sprintf("🤖 Ollama: %s", ollamaStatus))
}

// getCPUUsage 获取CPU使用率（简化版本）
func (sb *StatusBar) getCPUUsage() int {
	// 这里应该使用实际的系统API来获取CPU使用率
	// 目前返回模拟数据
	return int(time.Now().Unix() % 20) + 5 // 5-25%的模拟数据
}

// getNetworkStatus 获取网络状态
func (sb *StatusBar) getNetworkStatus() string {
	// 这里应该实际检测网络连接状态
	// 目前返回模拟数据
	statuses := []string{"连线正常", "连线中", "离线"}
	return statuses[int(time.Now().Unix())%len(statuses)]
}

// getOllamaStatus 获取Ollama状态
func (sb *StatusBar) getOllamaStatus() string {
	// 这里应该实际检测Ollama服务状态
	// 目前返回模拟数据
	statuses := []string{"qwen2.5:7b 已载入", "未连接", "启动中"}
	return statuses[int(time.Now().Unix()/10)%len(statuses)]
}

// SetProjectInfo 设置当前项目信息
func (sb *StatusBar) SetProjectInfo(projectName string, aiModel string, mode string) {
	statusText := fmt.Sprintf("🟢 %s: 运行中 | 📁 项目: %s", aiModel, projectName)
	if mode != "" {
		statusText += fmt.Sprintf(" | %s", mode)
	}
	sb.statusLabel.SetText(statusText)
}

// ShowProgress 显示进度信息
func (sb *StatusBar) ShowProgress(message string) {
	// 可以在这里添加进度条或加载动画
	sb.SetMessage(fmt.Sprintf("🔄 %s", message))
}

// ShowSuccess 显示成功消息
func (sb *StatusBar) ShowSuccess(message string) {
	sb.SetMessage(fmt.Sprintf("✅ %s", message))
}

// ShowError 显示错误消息
func (sb *StatusBar) ShowError(message string) {
	sb.SetMessage(fmt.Sprintf("❌ %s", message))
}

// ShowWarning 显示警告消息
func (sb *StatusBar) ShowWarning(message string) {
	sb.SetMessage(fmt.Sprintf("⚠️ %s", message))
}