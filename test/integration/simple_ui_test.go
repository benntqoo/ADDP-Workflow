package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ai-launcher/internal/project"
	"ai-launcher/internal/tui"
)

func TestNewUIModel_BasicFunctionality(t *testing.T) {
	model := tui.NewNewUIModel()

	require.NotNil(t, model, "模型应该被正确初始化")

	// 测试基本的getter方法
	assert.Equal(t, 0, model.GetSelectedIndex(), "初始选择索引应该为0")
	assert.False(t, model.GetYoloMode(), "初始YOLO模式应该为false")

	// 测试配置管理器
	configManager := model.GetConfigManager()
	require.NotNil(t, configManager, "配置管理器应该被初始化")

	// 测试终端管理器
	terminalManager := model.GetTerminalManager()
	require.NotNil(t, terminalManager, "终端管理器应该被初始化")
}

func TestNewUIModel_ConfigurationMethods(t *testing.T) {
	model := tui.NewNewUIModel()

	// 测试设置项目路径
	testPath := "/test/project/path"
	model.SetProjectPath(testPath)
	// 注意：我们无法直接获取项目路径，但可以通过其他方式验证

	// 测试设置项目名称
	testName := "test-project"
	model.SetProjectName(testName)

	// 测试设置AI模型
	model.SetSelectedModel(project.ModelGeminiCLI)

	// 测试设置YOLO模式
	model.SetYoloMode(true)
	assert.True(t, model.GetYoloMode(), "YOLO模式应该被设置为true")

	// 测试设置错误消息
	testError := "测试错误"
	model.SetErrorMessage(testError)

	// 测试UI步骤设置
	model.SetCurrentStep(tui.StepConfirm)
}

func TestNewUIModel_ProjectConfigIntegration(t *testing.T) {
	model := tui.NewNewUIModel()
	configManager := model.GetConfigManager()

	// 测试获取可用模型
	models := configManager.GetAvailableModels()
	assert.Len(t, models, 3, "应该有3个可用模型")

	expectedModels := []project.AIModelType{
		project.ModelClaudeCode,
		project.ModelGeminiCLI,
		project.ModelCodex,
	}

	for _, expectedModel := range expectedModels {
		assert.Contains(t, models, expectedModel, "应该包含%s模型", expectedModel)
	}

	// 测试模型验证
	for _, model := range models {
		assert.True(t, configManager.IsValidModel(model), "模型%s应该是有效的", model)
	}

	assert.False(t, configManager.IsValidModel(project.ModelCustom), "ModelCustom应该是无效的")
}

func TestNewUIModel_ModelTypeOperations(t *testing.T) {
	// 测试模型类型的字符串表示
	assert.Equal(t, "Claude Code", project.ModelClaudeCode.String())
	assert.Equal(t, "Gemini CLI", project.ModelGeminiCLI.String())
	assert.Equal(t, "Codex", project.ModelCodex.String())

	// 测试模型图标
	assert.Equal(t, "🤖", project.ModelClaudeCode.GetIcon())
	assert.Equal(t, "💎", project.ModelGeminiCLI.GetIcon())
	assert.Equal(t, "🔧", project.ModelCodex.GetIcon())

	// 测试命令生成（普通模式）
	claudeCmd := project.ModelClaudeCode.GetCommand(false)
	assert.Equal(t, []string{"claude"}, claudeCmd)

	geminiCmd := project.ModelGeminiCLI.GetCommand(false)
	assert.Equal(t, []string{"gemini"}, geminiCmd)

	codexCmd := project.ModelCodex.GetCommand(false)
	assert.Equal(t, []string{"codex"}, codexCmd)

	// 测试命令生成（YOLO模式）
	claudeCmdYolo := project.ModelClaudeCode.GetCommand(true)
	assert.Equal(t, []string{"claude", "--dangerously-skip-permissions"}, claudeCmdYolo)

	geminiCmdYolo := project.ModelGeminiCLI.GetCommand(true)
	assert.Equal(t, []string{"gemini", "--yolo"}, geminiCmdYolo)

	codexCmdYolo := project.ModelCodex.GetCommand(true)
	assert.Equal(t, []string{"codex", "--dangerously-bypass-approvals-and-sandbox"}, codexCmdYolo)
}

func TestNewUIModel_WindowSizeHandling(t *testing.T) {
	model := tui.NewNewUIModel()

	// 获取初始窗口大小
	width, height := model.GetWindowSize()
	assert.Equal(t, 0, width, "初始宽度应该为0")
	assert.Equal(t, 0, height, "初始高度应该为0")
}

func TestNewUIModel_ViewRendering(t *testing.T) {
	model := tui.NewNewUIModel()

	// 测试初始视图渲染
	view := model.View()
	assert.NotEmpty(t, view, "视图应该不为空")
	assert.Contains(t, view, "欢迎使用AI启动器", "应该包含欢迎信息")

	// 测试不同步骤的视图
	steps := []tui.UIStep{
		tui.StepWelcome,
		tui.StepProjectPath,
		tui.StepModelSelect,
		tui.StepYoloMode,
		tui.StepConfirm,
		tui.StepLaunching,
		tui.StepComplete,
	}

	for _, step := range steps {
		model.SetCurrentStep(step)
		view := model.View()
		assert.NotEmpty(t, view, "步骤%s的视图应该不为空", step.String())

		// 验证每个步骤都有适当的内容
		switch step {
		case tui.StepWelcome:
			assert.Contains(t, view, "欢迎", "欢迎步骤应该包含欢迎信息")
		case tui.StepProjectPath:
			assert.Contains(t, view, "项目目录", "项目路径步骤应该包含目录选择")
		case tui.StepModelSelect:
			assert.Contains(t, view, "AI模型", "模型选择步骤应该包含模型选择")
		case tui.StepYoloMode:
			assert.Contains(t, view, "运行模式", "YOLO模式步骤应该包含模式选择")
		case tui.StepConfirm:
			assert.Contains(t, view, "确认", "确认步骤应该包含确认信息")
		case tui.StepLaunching:
			assert.Contains(t, view, "启动", "启动步骤应该包含启动信息")
		case tui.StepComplete:
			assert.Contains(t, view, "完成", "完成步骤应该包含完成信息")
		}
	}
}

func TestNewUIModel_RecentProjectsHandling(t *testing.T) {
	model := tui.NewNewUIModel()

	// 测试获取最近项目（初始应该为空）
	recentProjects := model.GetRecentProjects()
	assert.NotNil(t, recentProjects, "最近项目列表应该不为nil")

	// 测试刷新最近项目
	model.RefreshRecentProjects()
	recentProjectsAfterRefresh := model.GetRecentProjects()
	assert.NotNil(t, recentProjectsAfterRefresh, "刷新后的最近项目列表应该不为nil")
}

func TestNewUIModel_StepProgression(t *testing.T) {
	// 测试步骤字符串表示
	steps := []struct {
		step     tui.UIStep
		expected string
	}{
		{tui.StepWelcome, "欢迎"},
		{tui.StepProjectPath, "项目路径"},
		{tui.StepModelSelect, "模型选择"},
		{tui.StepYoloMode, "YOLO模式"},
		{tui.StepConfirm, "确认配置"},
		{tui.StepLaunching, "启动中"},
		{tui.StepComplete, "完成"},
	}

	for _, test := range steps {
		assert.Equal(t, test.expected, test.step.String(), "步骤字符串应该匹配")
	}
}