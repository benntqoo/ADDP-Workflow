package template

import (
	"sync"
	"text/template"
	"time"
)

// 預定義模板 ID
const (
	TemplateCoding        = "coding"
	TemplateDebug         = "debug"
	TemplateRefactor      = "refactor"
	TemplateDocumentation = "documentation"
	TemplateTesting       = "testing"
	TemplatePerformance   = "performance"
)

// 模板分類
const (
	CategoryDevelopment = "development"
	CategoryMaintenance = "maintenance"
	CategoryAnalysis    = "analysis"
)

// QueryTemplate 查詢優化模板
type QueryTemplate struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Category    string            `json:"category"`
	Prompt      string            `json:"prompt"`       // 模板提示詞，支持變量替換
	Variables   []string          `json:"variables"`    // 模板中的變量列表
	Metadata    map[string]string `json:"metadata"`     // 額外的元數據
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	IsBuiltIn   bool              `json:"is_built_in"`  // 是否為內建模板
}

// ApplyResult 模板應用結果
type ApplyResult struct {
	TemplateID      string            `json:"template_id"`
	OriginalQuery   string            `json:"original_query"`
	OptimizedPrompt string            `json:"optimized_prompt"`
	Context         map[string]string `json:"context"`
	AppliedAt       time.Time         `json:"applied_at"`
	TokensEstimate  int               `json:"tokens_estimate"`
}

// TemplateManager 模板管理器
type TemplateManager struct {
	templates    map[string]*QueryTemplate
	templateCache map[string]*template.Template // 已編譯的模板緩存
	stats        map[string]*TemplateStats     // 使用統計
	config       *TemplateConfig               // 配置
	mu           sync.RWMutex
}

// TemplateStats 模板使用統計
type TemplateStats struct {
	TemplateID  string    `json:"template_id"`
	UsageCount  int       `json:"usage_count"`
	LastUsed    time.Time `json:"last_used"`
	AvgTokens   float64   `json:"avg_tokens"`
	SuccessRate float64   `json:"success_rate"`
}

// TemplateConfig 模板配置
type TemplateConfig struct {
	EnableStats     bool `json:"enable_stats"`
	MaxCustom       int  `json:"max_custom"`        // 最大自定義模板數量
	CacheSize       int  `json:"cache_size"`        // 緩存大小
	ValidateStrict  bool `json:"validate_strict"`   // 嚴格驗證模式
}

// VariableInfo 變量信息
type VariableInfo struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Default     string `json:"default"`
	Description string `json:"description"`
}

// TemplateValidationError 模板驗證錯誤
type TemplateValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *TemplateValidationError) Error() string {
	return e.Message
}