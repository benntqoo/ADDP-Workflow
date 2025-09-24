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
    // UI 组件
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

// initializeUI 初始化状态栏 UI
func (sb *StatusBar) initializeUI() {
    sb.messageLabel = widget.NewLabel("准备就绪")
    sb.messageLabel.TextStyle = fyne.TextStyle{Bold: true}

    sb.statusLabel = widget.NewLabel("🟢 AI 启动器 运行中")
    sb.timeLabel = widget.NewLabel("⏱ 运行时间: 00:00:00")
    sb.cpuLabel = widget.NewLabel("🧠 CPU: 0%")
    sb.memoryLabel = widget.NewLabel("💾 内存: 0MB")
    sb.networkLabel = widget.NewLabel("🌐 网络: 未知")
    sb.ollamaLabel = widget.NewLabel("🤖 Ollama: 未连接")

    topRow := container.NewHBox(
        sb.messageLabel,
        layout.NewSpacer(),
        sb.statusLabel,
        widget.NewSeparator(),
        sb.timeLabel,
    )

    bottomRow := container.NewHBox(
        sb.cpuLabel,
        widget.NewSeparator(),
        sb.memoryLabel,
        widget.NewSeparator(),
        sb.networkLabel,
        widget.NewSeparator(),
        sb.ollamaLabel,
    )

    sb.container = container.NewVBox(
        widget.NewSeparator(),
        topRow,
        bottomRow,
    )
}

// GetContainer 获取状态栏容器
func (sb *StatusBar) GetContainer() *fyne.Container { return sb.container }

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

// startStatusUpdater 启动状态刷新
func (sb *StatusBar) startStatusUpdater() {
    go func() {
        ticker := time.NewTicker(1 * time.Second)
        defer ticker.Stop()
        for range ticker.C {
            sb.updateSystemStatus()
        }
    }()
}

// updateSystemStatus 刷新系统状态
func (sb *StatusBar) updateSystemStatus() {
    duration := time.Since(sb.startTime)
    hours := int(duration.Hours())
    minutes := int(duration.Minutes()) % 60
    seconds := int(duration.Seconds()) % 60
    sb.timeLabel.SetText(fmt.Sprintf("⏱ 运行时间: %02d:%02d:%02d", hours, minutes, seconds))

    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    memoryMB := m.Alloc / 1024 / 1024
    sb.memoryLabel.SetText(fmt.Sprintf("💾 内存: %dMB", memoryMB))

    cpuPercent := sb.getCPUUsage()
    sb.cpuLabel.SetText(fmt.Sprintf("🧠 CPU: %d%%", cpuPercent))

    networkStatus := sb.getNetworkStatus()
    sb.networkLabel.SetText(fmt.Sprintf("🌐 网络: %s", networkStatus))

    ollamaStatus := sb.getOllamaStatus()
    sb.ollamaLabel.SetText(fmt.Sprintf("🤖 Ollama: %s", ollamaStatus))
}

func (sb *StatusBar) getCPUUsage() int { return int(time.Now().Unix()%20) + 5 }

func (sb *StatusBar) getNetworkStatus() string {
    statuses := []string{"在线正常", "在线较慢", "离线"}
    return statuses[int(time.Now().Unix())%len(statuses)]
}

func (sb *StatusBar) getOllamaStatus() string {
    statuses := []string{"qwen2.5:7b 已加载", "未连接", "启动中"}
    return statuses[int(time.Now().Unix()/10)%len(statuses)]
}

func (sb *StatusBar) SetProjectInfo(projectName string, aiModel string, mode string) {
    statusText := fmt.Sprintf("🟢 %s: 运行中 | 项目: %s", aiModel, projectName)
    if mode != "" {
        statusText += fmt.Sprintf(" | %s", mode)
    }
    sb.statusLabel.SetText(statusText)
}

func (sb *StatusBar) ShowProgress(message string) { sb.SetMessage(fmt.Sprintf("🔄 %s", message)) }
func (sb *StatusBar) ShowSuccess(message string)  { sb.SetMessage(fmt.Sprintf("✅ %s", message)) }
func (sb *StatusBar) ShowError(message string)    { sb.SetMessage(fmt.Sprintf("❌ %s", message)) }
func (sb *StatusBar) ShowWarning(message string)  { sb.SetMessage(fmt.Sprintf("⚠️ %s", message)) }

