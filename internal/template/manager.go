package template

import (
	"fmt"
	"regexp"
	"strings"
	"text/template"
	"time"
)

// NewTemplateManager 創建新的模板管理器
func NewTemplateManager() *TemplateManager {
	manager := &TemplateManager{
		templates:     make(map[string]*QueryTemplate),
		templateCache: make(map[string]*template.Template),
		stats:         make(map[string]*TemplateStats),
		config: &TemplateConfig{
			EnableStats:    true,
			MaxCustom:     50,
			CacheSize:     100,
			ValidateStrict: false,
		},
	}

	// 初始化內建模板
	manager.initBuiltInTemplates()

	return manager
}

// GetAvailableTemplates 獲取所有可用模板
func (tm *TemplateManager) GetAvailableTemplates() []*QueryTemplate {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	templates := make([]*QueryTemplate, 0, len(tm.templates))
	for _, tmpl := range tm.templates {
		templates = append(templates, tmpl)
	}

	return templates
}

// GetTemplate 根據 ID 獲取模板
func (tm *TemplateManager) GetTemplate(templateID string) (*QueryTemplate, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	template, exists := tm.templates[templateID]
	return template, exists
}

// ApplyTemplate 應用模板到用戶查詢
func (tm *TemplateManager) ApplyTemplate(templateID, userQuery string, context map[string]string) (*ApplyResult, error) {
	if strings.TrimSpace(userQuery) == "" {
		return nil, fmt.Errorf("user query cannot be empty")
	}

	template, exists := tm.GetTemplate(templateID)
	if !exists {
		return nil, fmt.Errorf("template not found: %s", templateID)
	}

	// 準備模板數據
	data := make(map[string]string)
	data["query"] = userQuery

	// 添加用戶提供的上下文
	if context != nil {
		for key, value := range context {
			data[key] = value
		}
	}

	// 渲染模板
	optimizedPrompt, err := tm.renderTemplate(template.Prompt, data)
	if err != nil {
		return nil, fmt.Errorf("failed to render template: %w", err)
	}

	// 計算 token 估算
	tokensEstimate := tm.estimateTokens(optimizedPrompt)

	// 更新統計
	tm.updateStats(templateID, tokensEstimate)

	return &ApplyResult{
		TemplateID:      templateID,
		OriginalQuery:   userQuery,
		OptimizedPrompt: optimizedPrompt,
		Context:         context,
		AppliedAt:       time.Now(),
		TokensEstimate:  tokensEstimate,
	}, nil
}

// GetTemplatesByCategory 根據分類獲取模板
func (tm *TemplateManager) GetTemplatesByCategory(category string) []*QueryTemplate {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	var templates []*QueryTemplate
	for _, tmpl := range tm.templates {
		if tmpl.Category == category {
			templates = append(templates, tmpl)
		}
	}

	return templates
}

// ValidateTemplate 驗證模板
func (tm *TemplateManager) ValidateTemplate(tmpl *QueryTemplate) bool {
	if tmpl == nil {
		return false
	}

	// 檢查必需字段
	if strings.TrimSpace(tmpl.ID) == "" {
		return false
	}

	if strings.TrimSpace(tmpl.Prompt) == "" {
		return false
	}

	// 驗證模板語法
	_, err := template.New("test").Parse(tmpl.Prompt)
	if err != nil {
		return false
	}

	// 檢查變量一致性
	return tm.validateTemplateVariables(tmpl)
}

// AddCustomTemplate 添加自定義模板
func (tm *TemplateManager) AddCustomTemplate(tmpl *QueryTemplate) error {
	if !tm.ValidateTemplate(tmpl) {
		return fmt.Errorf("invalid template")
	}

	tm.mu.Lock()
	defer tm.mu.Unlock()

	// 檢查是否已存在
	if _, exists := tm.templates[tmpl.ID]; exists {
		return fmt.Errorf("template already exists: %s", tmpl.ID)
	}

	// 設置創建時間
	if tmpl.CreatedAt.IsZero() {
		tmpl.CreatedAt = time.Now()
	}
	tmpl.UpdatedAt = time.Now()
	tmpl.IsBuiltIn = false

	tm.templates[tmpl.ID] = tmpl
	return nil
}

// initBuiltInTemplates 初始化內建模板
func (tm *TemplateManager) initBuiltInTemplates() {
	builtInTemplates := []*QueryTemplate{
		{
			ID:          TemplateCoding,
			Name:        "編碼模板",
			Description: "用於代碼開發和實現的優化模板",
			Category:    CategoryDevelopment,
			Prompt: `作為一名資深{{.language}}開發者，請幫助我{{.query}}。

請考慮以下要求：
- 使用最佳實踐和編碼規範
- 提供清晰的代碼註釋
- 考慮錯誤處理和邊界情況
- 複雜度：{{.complexity}}
{{if .framework}}
- 使用框架：{{.framework}}
{{end}}

請提供詳細的實現方案和代碼示例。`,
			Variables: []string{"query", "language", "complexity", "framework"},
			CreatedAt: time.Now(),
			IsBuiltIn: true,
		},
		{
			ID:          TemplateDebug,
			Name:        "調試模板",
			Description: "用於問題診斷和錯誤修復的優化模板",
			Category:    CategoryMaintenance,
			Prompt: `作為調試專家，請幫助我分析和解決以下問題：

問題描述：{{.query}}

相關信息：
- 程式語言：{{.language}}
- 錯誤類型：{{.error_type}}
{{if .stack_trace}}
- 堆疊追蹤：{{.stack_trace}}
{{end}}

請提供：
1. 問題根因分析
2. 具體的解決方案
3. 預防措施建議`,
			Variables: []string{"query", "language", "error_type", "stack_trace"},
			CreatedAt: time.Now(),
			IsBuiltIn: true,
		},
		{
			ID:          TemplateRefactor,
			Name:        "重構模板",
			Description: "用於代碼重構和優化的模板",
			Category:    CategoryMaintenance,
			Prompt: `作為代碼重構專家，請幫助我{{.query}}。

重構目標：
- 提高代碼可讀性和可維護性
- 優化性能
- 遵循 SOLID 原則
- 改善代碼結構

{{if .focus_area}}
重點關注：{{.focus_area}}
{{end}}

請提供重構前後的對比和詳細的改進建議。`,
			Variables: []string{"query", "focus_area"},
			CreatedAt: time.Now(),
			IsBuiltIn: true,
		},
		{
			ID:          TemplateDocumentation,
			Name:        "文檔模板",
			Description: "用於生成和改進文檔的模板",
			Category:    CategoryDevelopment,
			Prompt: `作為技術文檔專家，請幫助我{{.query}}。

文檔要求：
- 清晰簡潔的語言
- 結構化的內容組織
- 包含實際使用示例
- 面向{{.target_audience}}

{{if .doc_type}}
文檔類型：{{.doc_type}}
{{end}}

請確保文檔易於理解和實用。`,
			Variables: []string{"query", "target_audience", "doc_type"},
			CreatedAt: time.Now(),
			IsBuiltIn: true,
		},
		{
			ID:          TemplateTesting,
			Name:        "測試模板",
			Description: "用於編寫和改進測試的模板",
			Category:    CategoryDevelopment,
			Prompt: `作為測試專家，請幫助我{{.query}}。

測試要求：
- 完整的測試覆蓋
- 包含正面和負面測試案例
- 使用適當的測試框架
- 清晰的測試意圖

{{if .test_type}}
測試類型：{{.test_type}}
{{end}}
{{if .coverage_target}}
覆蓋率目標：{{.coverage_target}}
{{end}}

請提供詳細的測試策略和實現代碼。`,
			Variables: []string{"query", "test_type", "coverage_target"},
			CreatedAt: time.Now(),
			IsBuiltIn: true,
		},
		{
			ID:          TemplatePerformance,
			Name:        "性能優化模板",
			Description: "用於性能分析和優化的模板",
			Category:    CategoryAnalysis,
			Prompt: `作為性能優化專家，請幫助我{{.query}}。

性能分析重點：
- 識別性能瓶頸
- 提供具體的優化方案
- 考慮內存和 CPU 使用
- 提供性能測試建議

{{if .performance_metric}}
關注指標：{{.performance_metric}}
{{end}}
{{if .current_performance}}
當前性能：{{.current_performance}}
{{end}}

請提供詳細的分析和優化建議。`,
			Variables: []string{"query", "performance_metric", "current_performance"},
			CreatedAt: time.Now(),
			IsBuiltIn: true,
		},
	}

	for _, tmpl := range builtInTemplates {
		tm.templates[tmpl.ID] = tmpl
	}
}

// renderTemplate 渲染模板
func (tm *TemplateManager) renderTemplate(promptTemplate string, data map[string]string) (string, error) {
	// 計算模板哈希作為緩存鍵
	cacheKey := tm.getTemplateHash(promptTemplate)

	// 嘗試從緩存獲取
	tm.mu.RLock()
	cachedTemplate, exists := tm.templateCache[cacheKey]
	tm.mu.RUnlock()

	var tmpl *template.Template
	var err error

	if exists {
		tmpl = cachedTemplate
	} else {
		// 編譯新模板
		tmpl, err = template.New("prompt").Parse(promptTemplate)
		if err != nil {
			return "", err
		}

		// 添加到緩存
		tm.mu.Lock()
		if len(tm.templateCache) < tm.config.CacheSize {
			tm.templateCache[cacheKey] = tmpl
		}
		tm.mu.Unlock()
	}

	var result strings.Builder
	if err := tmpl.Execute(&result, data); err != nil {
		return "", err
	}

	return result.String(), nil
}

// validateTemplateVariables 驗證模板變量
func (tm *TemplateManager) validateTemplateVariables(tmpl *QueryTemplate) bool {
	// 從模板中提取變量
	re := regexp.MustCompile(`\{\{\.(\w+)\}\}`)
	matches := re.FindAllStringSubmatch(tmpl.Prompt, -1)

	templateVars := make(map[string]bool)
	for _, match := range matches {
		if len(match) > 1 {
			templateVars[match[1]] = true
		}
	}

	// 檢查聲明的變量是否都在模板中使用
	declaredVars := make(map[string]bool)
	for _, variable := range tmpl.Variables {
		declaredVars[variable] = true
	}

	// 模板中的變量都應該在聲明列表中
	for variable := range templateVars {
		if !declaredVars[variable] {
			return false
		}
	}

	return true
}

// estimateTokens 估算 token 數量
func (tm *TemplateManager) estimateTokens(text string) int {
	// 簡單估算：中文字符 1 token，英文字符 0.25 token
	tokens := 0
	for _, r := range text {
		if r > 127 { // 非 ASCII 字符（主要是中文）
			tokens++
		} else {
			tokens += 1 // 英文字符
		}
	}
	return tokens / 3 // 平均化處理
}

// getTemplateHash 生成模板哈希值（簡單實現）
func (tm *TemplateManager) getTemplateHash(template string) string {
	// 簡單的哈希實現：使用字符串長度和前後字符
	length := len(template)
	if length == 0 {
		return "empty"
	}

	first := template[0]
	last := template[length-1]

	return fmt.Sprintf("%d_%c_%c", length, first, last)
}

// updateStats 更新模板使用統計
func (tm *TemplateManager) updateStats(templateID string, tokens int) {
	if !tm.config.EnableStats {
		return
	}

	tm.mu.Lock()
	defer tm.mu.Unlock()

	stats, exists := tm.stats[templateID]
	if !exists {
		stats = &TemplateStats{
			TemplateID: templateID,
		}
		tm.stats[templateID] = stats
	}

	stats.UsageCount++
	stats.LastUsed = time.Now()

	// 更新平均 token 數
	if stats.UsageCount == 1 {
		stats.AvgTokens = float64(tokens)
	} else {
		stats.AvgTokens = (stats.AvgTokens*float64(stats.UsageCount-1) + float64(tokens)) / float64(stats.UsageCount)
	}

	// 假設所有應用都成功（實際實現中可能需要追蹤失敗）
	stats.SuccessRate = 1.0
}

// GetStats 獲取模板使用統計
func (tm *TemplateManager) GetStats(templateID string) (*TemplateStats, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	stats, exists := tm.stats[templateID]
	return stats, exists
}