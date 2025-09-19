package ollama

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// NewOllamaClient 創建新的 Ollama 客戶端
func NewOllamaClient(baseURL string) *OllamaClient {
	return &OllamaClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		Model: "qwen2.5:14b", // 默認模型
	}
}

// IsHealthy 檢查 Ollama 服務是否健康
func (c *OllamaClient) IsHealthy(ctx context.Context) bool {
	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+"/api/tags", nil)
	if err != nil {
		return false
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

// OptimizeQuery 使用 Ollama 優化查詢
func (c *OllamaClient) OptimizeQuery(ctx context.Context, query, queryContext string) (*OptimizationResult, error) {
	if strings.TrimSpace(query) == "" {
		return &OptimizationResult{}, fmt.Errorf("query cannot be empty")
	}

	startTime := time.Now()

	// 構建優化提示詞
	prompt := c.buildOptimizationPrompt(query, queryContext)

	// 發送生成請求
	generateResp, err := c.sendGenerateRequest(ctx, prompt)
	if err != nil {
		return &OptimizationResult{}, err
	}

	processingTime := time.Since(startTime)

	// 計算估算的 token 數量
	tokensUsed := c.calculateTokens(generateResp.Response)

	return &OptimizationResult{
		OriginalQuery:   query,
		OptimizedQuery:  strings.TrimSpace(generateResp.Response),
		Model:          generateResp.Model,
		ProcessingTime: processingTime,
		TokensUsed:     tokensUsed,
		Confidence:     0.8, // 固定信心度，實際實現中可能會更復雜
	}, nil
}

// GetAvailableModels 獲取可用模型列表
func (c *OllamaClient) GetAvailableModels(ctx context.Context) ([]ModelInfo, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+"/api/tags", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ollama returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var listResp ListModelsResponse
	if err := json.Unmarshal(body, &listResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// 標記所有模型為可用
	for i := range listResp.Models {
		listResp.Models[i].Available = true
	}

	return listResp.Models, nil
}

// ValidateModel 驗證模型是否可用
func (c *OllamaClient) ValidateModel(ctx context.Context, modelName string) bool {
	if strings.TrimSpace(modelName) == "" {
		return false
	}

	models, err := c.GetAvailableModels(ctx)
	if err != nil {
		return false
	}

	for _, model := range models {
		if model.Name == modelName {
			return model.Available
		}
	}

	return false
}

// SetModel 設置默認模型
func (c *OllamaClient) SetModel(model string) {
	c.Model = model
}

// GetModel 獲取當前默認模型
func (c *OllamaClient) GetModel() string {
	return c.Model
}

// buildOptimizationPrompt 構建查詢優化提示詞
func (c *OllamaClient) buildOptimizationPrompt(query, queryContext string) string {
	return fmt.Sprintf(`請優化以下查詢，讓它更清晰、具體和有效：

原始查詢: %s
上下文: %s

請提供優化後的查詢，只返回優化後的查詢內容，不要包含其他解釋文字。`, query, queryContext)
}

// calculateTokens 估算 token 數量
func (c *OllamaClient) calculateTokens(text string) int {
	// 簡單估算：中文字符按 1 個 token，英文字符按 1/4 個 token
	tokens := 0
	for _, r := range text {
		if r > 127 { // 非 ASCII 字符（主要是中文）
			tokens++
		} else {
			tokens += 1 // 英文字符計為 0.25 token，但至少 1
		}
	}
	return tokens / 4 // 平均化處理
}

// sendGenerateRequest 發送生成請求的通用方法
func (c *OllamaClient) sendGenerateRequest(ctx context.Context, prompt string) (*GenerateResponse, error) {
	generateReq := GenerateRequest{
		Model:  c.Model,
		Prompt: prompt,
		Stream: false,
	}

	reqBody, err := json.Marshal(generateReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.BaseURL+"/api/generate", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ollama returned status %d", resp.StatusCode)
	}

	var generateResp GenerateResponse
	if err := json.NewDecoder(resp.Body).Decode(&generateResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &generateResp, nil
}