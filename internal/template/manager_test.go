package template

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTemplateManager(t *testing.T) {
	manager := NewTemplateManager()

	assert.NotNil(t, manager)
	assert.NotEmpty(t, manager.GetAvailableTemplates())
}

func TestTemplateManager_GetTemplate(t *testing.T) {
	manager := NewTemplateManager()

	tests := []struct {
		name        string
		templateID  string
		shouldExist bool
	}{
		{"Coding template", TemplateCoding, true},
		{"Debug template", TemplateDebug, true},
		{"Refactor template", TemplateRefactor, true},
		{"Documentation template", TemplateDocumentation, true},
		{"Testing template", TemplateTesting, true},
		{"Performance template", TemplatePerformance, true},
		{"Non-existent template", "non-existent", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			template, exists := manager.GetTemplate(tt.templateID)

			if tt.shouldExist {
				assert.True(t, exists)
				assert.NotNil(t, template)
				assert.Equal(t, tt.templateID, template.ID)
				assert.NotEmpty(t, template.Name)
				assert.NotEmpty(t, template.Prompt)
				assert.NotEmpty(t, template.Category)
			} else {
				assert.False(t, exists)
				assert.Nil(t, template)
			}
		})
	}
}

func TestTemplateManager_ApplyTemplate(t *testing.T) {
	manager := NewTemplateManager()

	tests := []struct {
		name       string
		templateID string
		userQuery  string
		context    map[string]string
		expectErr  bool
	}{
		{
			name:       "Coding template with valid query",
			templateID: TemplateCoding,
			userQuery:  "如何實現一個 HTTP 服務器？",
			context: map[string]string{
				"language":    "Go",
				"framework":   "標準庫",
				"complexity":  "簡單",
			},
			expectErr: false,
		},
		{
			name:       "Debug template with error description",
			templateID: TemplateDebug,
			userQuery:  "我的程序崩潰了",
			context: map[string]string{
				"error_type": "panic",
				"language":   "Go",
				"stack_trace": "goroutine 1 [running]",
			},
			expectErr: false,
		},
		{
			name:       "Empty query",
			templateID: TemplateCoding,
			userQuery:  "",
			context:    nil,
			expectErr:  true,
		},
		{
			name:       "Non-existent template",
			templateID: "non-existent",
			userQuery:  "test query",
			context:    nil,
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := manager.ApplyTemplate(tt.templateID, tt.userQuery, tt.context)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.templateID, result.TemplateID)
				assert.Equal(t, tt.userQuery, result.OriginalQuery)
				assert.NotEmpty(t, result.OptimizedPrompt)
				assert.NotZero(t, result.AppliedAt)
			}
		})
	}
}

func TestTemplateManager_GetTemplatesByCategory(t *testing.T) {
	manager := NewTemplateManager()

	tests := []struct {
		name           string
		category       string
		expectedCount  int
		minExpected    int
	}{
		{"Development category", CategoryDevelopment, 0, 3}, // 至少 3 個開發模板
		{"Maintenance category", CategoryMaintenance, 0, 2}, // 至少 2 個維護模板
		{"Analysis category", CategoryAnalysis, 0, 1},       // 至少 1 個分析模板
		{"Non-existent category", "non-existent", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			templates := manager.GetTemplatesByCategory(tt.category)

			if tt.minExpected > 0 {
				assert.GreaterOrEqual(t, len(templates), tt.minExpected)
				for _, template := range templates {
					assert.Equal(t, tt.category, template.Category)
				}
			} else {
				assert.Len(t, templates, 0)
			}
		})
	}
}

func TestTemplateManager_ValidateTemplate(t *testing.T) {
	manager := NewTemplateManager()

	tests := []struct {
		name     string
		template *QueryTemplate
		isValid  bool
	}{
		{
			name: "Valid template",
			template: &QueryTemplate{
				ID:          "test-template",
				Name:        "測試模板",
				Description: "測試用途",
				Category:    CategoryDevelopment,
				Prompt:      "這是一個測試提示詞：{{.query}}",
				Variables:   []string{"query"},
				CreatedAt:   time.Now(),
			},
			isValid: true,
		},
		{
			name: "Missing ID",
			template: &QueryTemplate{
				Name:        "測試模板",
				Description: "測試用途",
				Category:    CategoryDevelopment,
				Prompt:      "這是一個測試提示詞：{{.query}}",
				Variables:   []string{"query"},
				CreatedAt:   time.Now(),
			},
			isValid: false,
		},
		{
			name: "Missing prompt",
			template: &QueryTemplate{
				ID:          "test-template",
				Name:        "測試模板",
				Description: "測試用途",
				Category:    CategoryDevelopment,
				Variables:   []string{"query"},
				CreatedAt:   time.Now(),
			},
			isValid: false,
		},
		{
			name: "Invalid variables",
			template: &QueryTemplate{
				ID:          "test-template",
				Name:        "測試模板",
				Description: "測試用途",
				Category:    CategoryDevelopment,
				Prompt:      "這是一個測試提示詞：{{.query}} {{.context}}",
				Variables:   []string{"query"}, // 缺少 context 變量
				CreatedAt:   time.Now(),
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := manager.ValidateTemplate(tt.template)
			assert.Equal(t, tt.isValid, valid)
		})
	}
}

func TestTemplateManager_CustomTemplate(t *testing.T) {
	manager := NewTemplateManager()

	// 添加自定義模板
	customTemplate := &QueryTemplate{
		ID:          "custom-test",
		Name:        "自定義測試模板",
		Description: "用戶自定義的測試模板",
		Category:    CategoryDevelopment,
		Prompt:      "請幫我 {{.action}} 關於 {{.topic}} 的代碼，使用 {{.language}} 語言",
		Variables:   []string{"action", "topic", "language"},
		CreatedAt:   time.Now(),
	}

	err := manager.AddCustomTemplate(customTemplate)
	require.NoError(t, err)

	// 驗證模板已添加
	retrieved, exists := manager.GetTemplate("custom-test")
	assert.True(t, exists)
	assert.Equal(t, customTemplate.ID, retrieved.ID)

	// 應用自定義模板
	context := map[string]string{
		"action":   "重構",
		"topic":    "HTTP 處理器",
		"language": "Go",
	}

	result, err := manager.ApplyTemplate("custom-test", "優化我的代碼", context)
	require.NoError(t, err)
	assert.Contains(t, result.OptimizedPrompt, "重構")
	assert.Contains(t, result.OptimizedPrompt, "HTTP 處理器")
	assert.Contains(t, result.OptimizedPrompt, "Go")
}

func TestTemplateManager_ConcurrentAccess(t *testing.T) {
	manager := NewTemplateManager()

	// 並發獲取模板
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(index int) {
			templateID := TemplateCoding
			if index%2 == 0 {
				templateID = TemplateDebug
			}

			template, exists := manager.GetTemplate(templateID)
			assert.True(t, exists)
			assert.NotNil(t, template)

			done <- true
		}(i)
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 10; i++ {
		select {
		case <-done:
			// 成功
		case <-time.After(5 * time.Second):
			t.Fatal("Concurrent access test timed out")
		}
	}
}

func TestApplyResult_Validation(t *testing.T) {
	result := &ApplyResult{
		TemplateID:      TemplateCoding,
		OriginalQuery:   "如何實現 REST API？",
		OptimizedPrompt: "請提供一個詳細的 REST API 實現指南...",
		Context:         map[string]string{"language": "Go"},
		AppliedAt:       time.Now(),
		TokensEstimate:  150,
	}

	assert.NotEmpty(t, result.TemplateID)
	assert.NotEmpty(t, result.OriginalQuery)
	assert.NotEmpty(t, result.OptimizedPrompt)
	assert.NotZero(t, result.AppliedAt)
	assert.Greater(t, result.TokensEstimate, 0)
}