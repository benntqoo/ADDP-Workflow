package gui

import (
	"testing"

	"ai-launcher/internal/project"
	"ai-launcher/internal/terminal"
)

// TestNewApp 测试应用创建
func TestNewApp(t *testing.T) {
	app := NewApp()

	if app == nil {
		t.Fatal("NewApp() returned nil")
	}

	if app.fyneApp == nil {
		t.Error("fyneApp not initialized")
	}

	if app.configManager == nil {
		t.Error("configManager not initialized")
	}

	if app.terminalManager == nil {
		t.Error("terminalManager not initialized")
	}

	if app.selectedProject == nil {
		t.Error("selectedProject binding not initialized")
	}

	if app.statusMessage == nil {
		t.Error("statusMessage binding not initialized")
	}
}

// TestLoadProject 测试项目加载功能
func TestLoadProject(t *testing.T) {
	app := NewApp()

	// 创建模拟项目配置
	testProject := project.ProjectConfig{
		Name:     "test-project",
		Path:     "/test/path",
		AIModel:  project.ModelClaudeCode,
		YoloMode: true,
	}

	// 初始化UI组件（模拟）
	app.initUIComponents()

	// 加载项目
	app.loadProject(testProject)

	// 验证加载结果
	if app.projectPathEntry.Text != testProject.Path {
		t.Errorf("Expected path %s, got %s", testProject.Path, app.projectPathEntry.Text)
	}

	if app.projectNameEntry.Text != testProject.Name {
		t.Errorf("Expected name %s, got %s", testProject.Name, app.projectNameEntry.Text)
	}

	if app.yoloModeCheck.Checked != testProject.YoloMode {
		t.Errorf("Expected YOLO mode %v, got %v", testProject.YoloMode, app.yoloModeCheck.Checked)
	}
}

// TestGetTerminalType 测试终端类型映射
func TestGetTerminalType(t *testing.T) {
	app := NewApp()

	tests := []struct {
		model    project.AIModelType
		expected terminal.TerminalType
	}{
		{project.ModelClaudeCode, terminal.TypeClaudeCode},
		{project.ModelGeminiCLI, terminal.TypeGeminiCLI},
		{project.ModelCodex, terminal.TypeCustom},
		{project.ModelCustom, terminal.TypeCustom},
	}

	for _, test := range tests {
		result := app.getTerminalType(test.model)
		if result != test.expected {
			t.Errorf("For model %s, expected %s, got %s",
				test.model, test.expected, result)
		}
	}
}

// TestClearForm 测试表单清空功能
func TestClearForm(t *testing.T) {
	app := NewApp()

	// 初始化UI组件
	app.initUIComponents()

	// 设置一些数据
	app.projectPathEntry.SetText("/some/path")
	app.projectNameEntry.SetText("some-project")
	app.yoloModeCheck.SetChecked(true)
	app.modelSelect.SetSelectedIndex(1)

	// 清空表单
	app.clearForm()

	// 验证清空结果
	if app.projectPathEntry.Text != "" {
		t.Error("Path entry not cleared")
	}

	if app.projectNameEntry.Text != "" {
		t.Error("Name entry not cleared")
	}

	if app.yoloModeCheck.Checked {
		t.Error("YOLO mode not cleared")
	}

	if app.modelSelect.SelectedIndex() != 0 {
		t.Error("Model selection not reset to first option")
	}
}

// initUIComponents 初始化UI组件用于测试
func (a *App) initUIComponents() {
	// 创建模拟的UI组件
	a.projectPathEntry = &mockEntry{}
	a.projectNameEntry = &mockEntry{}
	a.yoloModeCheck = &mockCheck{}
	a.modelSelect = &mockSelect{}
}

// 模拟组件用于测试

type mockEntry struct {
	Text string
}

func (e *mockEntry) SetText(text string) {
	e.Text = text
}

func (e *mockEntry) SetPlaceHolder(text string) {}

type mockCheck struct {
	Checked bool
}

func (c *mockCheck) SetChecked(checked bool) {
	c.Checked = checked
}

type mockSelect struct {
	index int
}

func (s *mockSelect) SetSelectedIndex(index int) {
	s.index = index
}

func (s *mockSelect) SelectedIndex() int {
	return s.index
}

func (s *mockSelect) SetSelected(option string) {}