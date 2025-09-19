package ollama

import (
	"net/http"
	"time"
)

// OllamaClient Ollama API 客戶端
type OllamaClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Model      string // 默認使用的模型
}

// OptimizationResult 查詢優化結果
type OptimizationResult struct {
	OriginalQuery   string        `json:"original_query"`
	OptimizedQuery  string        `json:"optimized_query"`
	Model          string        `json:"model"`
	ProcessingTime time.Duration `json:"processing_time"`
	TokensUsed     int           `json:"tokens_used"`
	Confidence     float64       `json:"confidence"`
}

// ModelInfo 模型信息
type ModelInfo struct {
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Modified  string `json:"modified"`
	Available bool   `json:"available"`
}

// HealthStatus Ollama 服務健康狀態
type HealthStatus struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// GenerateRequest Ollama 生成請求
type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

// GenerateResponse Ollama 生成響應
type GenerateResponse struct {
	Model     string `json:"model"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
	Context   []int  `json:"context,omitempty"`
	TotalTime int64  `json:"total_duration,omitempty"`
}

// ListModelsResponse 模型列表響應
type ListModelsResponse struct {
	Models []ModelInfo `json:"models"`
}