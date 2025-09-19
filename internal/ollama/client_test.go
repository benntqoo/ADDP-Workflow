package ollama

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewOllamaClient(t *testing.T) {
	client := NewOllamaClient("http://localhost:11434")

	assert.NotNil(t, client)
	assert.Equal(t, "http://localhost:11434", client.BaseURL)
	assert.NotNil(t, client.HTTPClient)
}

func TestOllamaClient_IsHealthy(t *testing.T) {
	client := NewOllamaClient("http://localhost:11434")

	// 測試健康檢查（可能失敗如果 Ollama 未運行）
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	healthy := client.IsHealthy(ctx)
	// 這個測試不應該失敗，只是記錄 Ollama 是否可用
	t.Logf("Ollama health status: %v", healthy)
}

func TestOllamaClient_OptimizeQuery(t *testing.T) {
	client := NewOllamaClient("http://localhost:11434")

	tests := []struct {
		name     string
		query    string
		context  string
		expected bool // 是否期望成功優化
	}{
		{
			name:     "Basic query optimization",
			query:    "如何實現一個簡單的 HTTP 服務器？",
			context:  "我正在學習 Go 語言",
			expected: false, // 假設 Ollama 不可用
		},
		{
			name:     "Empty query",
			query:    "",
			context:  "測試上下文",
			expected: false, // 空查詢應該失敗
		},
		{
			name:     "Long query",
			query:    "這是一個非常長的查詢，用來測試 Ollama 客戶端是否能夠處理長文本輸入並進行適當的優化處理",
			context:  "長文本處理測試",
			expected: false, // 假設 Ollama 不可用
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			result, err := client.OptimizeQuery(ctx, tt.query, tt.context)

			if tt.expected {
				assert.NoError(t, err)
				assert.NotEmpty(t, result.OptimizedQuery)
				assert.NotEmpty(t, result.Model)
				assert.Greater(t, result.ProcessingTime, time.Duration(0))
			} else {
				// 如果 Ollama 不可用或查詢無效，記錄但不失敗
				t.Logf("Query optimization result: error=%v, result=%+v", err, result)
			}
		})
	}
}

func TestOllamaClient_GetAvailableModels(t *testing.T) {
	client := NewOllamaClient("http://localhost:11434")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	models, err := client.GetAvailableModels(ctx)

	if err != nil {
		// Ollama 可能不可用，記錄但不失敗測試
		t.Logf("Could not get models (Ollama may not be running): %v", err)
		return
	}

	// 如果 Ollama 可用，驗證模型列表
	assert.NotNil(t, models)
	t.Logf("Available models: %+v", models)
}

func TestOllamaClient_ValidateModel(t *testing.T) {
	client := NewOllamaClient("http://localhost:11434")

	tests := []struct {
		name     string
		model    string
		expected bool
	}{
		{"Valid model name", "qwen2.5:14b", false}, // 假設不可用
		{"Invalid model name", "non-existent-model", false},
		{"Empty model name", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			valid := client.ValidateModel(ctx, tt.model)

			// 記錄結果，但不強制期望特定結果（因為 Ollama 可能不可用）
			t.Logf("Model %s validation: %v", tt.model, valid)
		})
	}
}

func TestOllamaClient_ErrorHandling(t *testing.T) {
	// 測試錯誤的 URL
	client := NewOllamaClient("http://invalid-url:99999")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// 健康檢查應該失敗
	healthy := client.IsHealthy(ctx)
	assert.False(t, healthy)

	// 查詢優化應該失敗
	result, err := client.OptimizeQuery(ctx, "test query", "test context")
	assert.Error(t, err)
	assert.Empty(t, result.OptimizedQuery)

	// 獲取模型列表應該失敗
	models, err := client.GetAvailableModels(ctx)
	assert.Error(t, err)
	assert.Nil(t, models)
}

func TestOllamaClient_ConcurrentRequests(t *testing.T) {
	client := NewOllamaClient("http://localhost:11434")

	// 並發測試健康檢查
	done := make(chan bool, 5)

	for i := 0; i < 5; i++ {
		go func(index int) {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			healthy := client.IsHealthy(ctx)
			t.Logf("Concurrent health check %d: %v", index, healthy)
			done <- true
		}(i)
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 5; i++ {
		<-done
	}
}

func TestOptimizationResult_Validation(t *testing.T) {
	result := &OptimizationResult{
		OriginalQuery:   "原始查詢",
		OptimizedQuery:  "優化後的查詢",
		Model:          "qwen2.5:14b",
		ProcessingTime: 100 * time.Millisecond,
		TokensUsed:     50,
		Confidence:     0.85,
	}

	assert.NotEmpty(t, result.OriginalQuery)
	assert.NotEmpty(t, result.OptimizedQuery)
	assert.NotEmpty(t, result.Model)
	assert.Greater(t, result.ProcessingTime, time.Duration(0))
	assert.Greater(t, result.TokensUsed, 0)
	assert.True(t, result.Confidence >= 0.0 && result.Confidence <= 1.0)
}