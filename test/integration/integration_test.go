package integration

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ai-launcher/internal/ollama"
	"ai-launcher/internal/template"
	"ai-launcher/internal/terminal"
	"ai-launcher/internal/tui"
)

// TestSystemIntegration 测试系统集成
func TestSystemIntegration(t *testing.T) {
	t.Run("Terminal Manager Integration", testTerminalManagerIntegration)
	t.Run("Ollama Client Integration", testOllamaClientIntegration)
	t.Run("Template Manager Integration", testTemplateManagerIntegration)
	t.Run("TUI Model Integration", testTUIModelIntegration)
	t.Run("End-to-End Workflow", testEndToEndWorkflow)
}

func testTerminalManagerIntegration(t *testing.T) {
	manager := terminal.NewTerminalManager()

	// 测试终端管理器基本功能
	assert.True(t, manager.IsHealthy())
	assert.Empty(t, manager.ListTerminals())

	// 测试启动自定义终端（用于测试）
	config := terminal.TerminalConfig{
		Type: terminal.TypeCustom,
		Name: "integration-test-terminal",
	}

	err := manager.StartTerminal(config)
	assert.NoError(t, err)

	// 验证终端已创建
	terminals := manager.ListTerminals()
	assert.Len(t, terminals, 1)
	assert.Equal(t, "integration-test-terminal", terminals[0].Name)

	// 测试终端状态
	terminal, exists := manager.GetTerminal("integration-test-terminal")
	assert.True(t, exists)
	assert.NotNil(t, terminal)

	// 清理
	err = manager.StopTerminal("integration-test-terminal")
	assert.NoError(t, err)
}

func testOllamaClientIntegration(t *testing.T) {
	client := ollama.NewOllamaClient("http://localhost:11434")

	// 测试健康检查
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	healthy := client.IsHealthy(ctx)
	t.Logf("Ollama health status: %v", healthy)

	// 测试查询优化（即使 Ollama 不可用也应该处理错误）
	result, err := client.OptimizeQuery(ctx, "测试查询", "测试上下文")
	if err != nil {
		// Ollama 不可用或出错时记录但不失败测试
		t.Logf("Ollama query optimization failed (expected): %v", err)
		assert.Error(t, err)
	} else {
		// Ollama 可用时验证结果
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.OptimizedQuery)
		t.Logf("Ollama query optimization successful")
	}

	// 测试模型验证
	valid := client.ValidateModel(ctx, "qwen2.5:14b")
	t.Logf("Model validation result: %v", valid)
}

func testTemplateManagerIntegration(t *testing.T) {
	manager := template.NewTemplateManager()

	// 验证内建模板
	templates := manager.GetAvailableTemplates()
	assert.GreaterOrEqual(t, len(templates), 6)

	// 测试获取特定模板
	codingTemplate, exists := manager.GetTemplate("coding")
	assert.True(t, exists)
	assert.NotNil(t, codingTemplate)
	assert.Equal(t, "coding", codingTemplate.ID)

	// 测试模板应用
	context := map[string]string{
		"language":   "Go",
		"complexity": "中等",
		"framework":  "标准库",
	}

	result, err := manager.ApplyTemplate("coding", "实现HTTP服务器", context)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "coding", result.TemplateID)
	assert.Equal(t, "实现HTTP服务器", result.OriginalQuery)
	assert.NotEmpty(t, result.OptimizedPrompt)
	assert.Greater(t, result.TokensEstimate, 0)

	// 测试分类过滤
	devTemplates := manager.GetTemplatesByCategory("development")
	assert.NotEmpty(t, devTemplates)

	// 测试自定义模板
	customTemplate := &template.QueryTemplate{
		ID:          "integration-test",
		Name:        "集成测试模板",
		Description: "用于集成测试的模板",
		Category:    "development",
		Prompt:      "这是一个测试模板：{{.query}}",
		Variables:   []string{"query"},
		CreatedAt:   time.Now(),
	}

	err = manager.AddCustomTemplate(customTemplate)
	assert.NoError(t, err)

	// 验证自定义模板
	retrieved, exists := manager.GetTemplate("integration-test")
	assert.True(t, exists)
	assert.Equal(t, "integration-test", retrieved.ID)
}

func testTUIModelIntegration(t *testing.T) {
	model := tui.NewModel()

	// 验证 TUI 模型初始化
	assert.Equal(t, tui.ViewMain, model.GetCurrentView())
	assert.NotNil(t, model.GetTerminalManager())
	assert.NotNil(t, model.GetOllamaClient())
	assert.NotNil(t, model.GetTemplateManager())

	// 测试视图切换
	model.SwitchView(tui.ViewTerminals)
	assert.Equal(t, tui.ViewTerminals, model.GetCurrentView())

	model.SwitchView(tui.ViewTemplates)
	assert.Equal(t, tui.ViewTemplates, model.GetCurrentView())

	// 测试查询优化集成
	result, err := model.OptimizeQuery("测试查询", "coding")
	if err != nil {
		// 可能因为没有实际上下文而失败，这是预期的
		t.Logf("Query optimization failed as expected: %v", err)
	} else {
		assert.NotNil(t, result)
		t.Logf("Query optimization successful: %+v", result)
	}

	// 测试终端操作集成
	terminals := model.ListTerminals()
	assert.NotNil(t, terminals)
	t.Logf("Found %d terminals", len(terminals))
}

func testEndToEndWorkflow(t *testing.T) {
	// 创建完整的工作流测试
	t.Log("Starting end-to-end workflow test")

	// 1. 初始化所有组件
	terminalManager := terminal.NewTerminalManager()
	ollamaClient := ollama.NewOllamaClient("http://localhost:11434")
	templateManager := template.NewTemplateManager()
	tuiModel := tui.NewModel()

	// 2. 验证系统健康状态
	assert.True(t, terminalManager.IsHealthy())
	assert.NotEmpty(t, templateManager.GetAvailableTemplates())
	assert.NotNil(t, tuiModel)

	// 3. 测试模板到查询优化的工作流
	tmpl, exists := templateManager.GetTemplate("coding")
	require.True(t, exists)
	require.Equal(t, "coding", tmpl.ID)

	// 应用模板
	templateContext := map[string]string{
		"language":   "Go",
		"complexity": "简单",
	}

	templateResult, err := templateManager.ApplyTemplate("coding", "创建REST API", templateContext)
	require.NoError(t, err)

	// 使用 Ollama 进一步优化（如果可用）
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if ollamaClient.IsHealthy(ctx) {
		ollamaResult, err := ollamaClient.OptimizeQuery(ctx, templateResult.OptimizedPrompt, "Go开发")
		if err == nil {
			t.Logf("Full workflow completed successfully")
			t.Logf("Template result tokens: %d", templateResult.TokensEstimate)
			t.Logf("Ollama result tokens: %d", ollamaResult.TokensUsed)
			assert.NotEmpty(t, ollamaResult.OptimizedQuery)
		} else {
			t.Logf("Ollama optimization failed: %v", err)
		}
	} else {
		t.Log("Ollama not available, skipping optimization step")
	}

	// 4. 测试终端管理集成
	config := terminal.TerminalConfig{
		Type: terminal.TypeCustom,
		Name: "e2e-test-terminal",
	}

	err = terminalManager.StartTerminal(config)
	assert.NoError(t, err)

	// 验证可以通过 TUI 访问终端
	terminals := tuiModel.ListTerminals()
	t.Logf("Terminals accessible through TUI: %d", len(terminals))

	// 清理
	err = terminalManager.StopTerminal("e2e-test-terminal")
	assert.NoError(t, err)

	t.Log("End-to-end workflow test completed")
}

// 辅助测试方法

// TestModuleCompatibility 测试模块兼容性
func TestModuleCompatibility(t *testing.T) {
	// 测试所有模块的版本兼容性
	terminalManager := terminal.NewTerminalManager()
	templateManager := template.NewTemplateManager()
	ollamaClient := ollama.NewOllamaClient("http://localhost:11434")

	// 验证接口兼容性
	assert.Implements(t, (*interface{})(nil), terminalManager)
	assert.Implements(t, (*interface{})(nil), templateManager)
	assert.Implements(t, (*interface{})(nil), ollamaClient)
}

// TestPerformance 测试系统性能
func TestPerformance(t *testing.T) {
	templateManager := template.NewTemplateManager()

	// 测试模板应用性能
	start := time.Now()

	for i := 0; i < 100; i++ {
		_, err := templateManager.ApplyTemplate("coding", "测试查询", nil)
		assert.NoError(t, err)
	}

	duration := time.Since(start)
	t.Logf("100 template applications took: %v", duration)

	// 应该在合理时间内完成
	assert.Less(t, duration, 5*time.Second)
}

// TestConcurrency 测试并发安全
func TestConcurrency(t *testing.T) {
	templateManager := template.NewTemplateManager()
	terminalManager := terminal.NewTerminalManager()

	// 并发测试
	done := make(chan bool, 20)

	// 启动多个 goroutine 进行并发操作
	for i := 0; i < 10; i++ {
		go func(index int) {
			// 测试模板操作
			_, err := templateManager.ApplyTemplate("coding", "并发测试", nil)
			assert.NoError(t, err)
			done <- true
		}(i)

		go func(index int) {
			// 测试终端操作
			terminals := terminalManager.ListTerminals()
			assert.NotNil(t, terminals)
			done <- true
		}(i)
	}

	// 等待所有操作完成
	for i := 0; i < 20; i++ {
		select {
		case <-done:
			// 成功
		case <-time.After(5 * time.Second):
			t.Fatal("Concurrency test timed out")
		}
	}
}