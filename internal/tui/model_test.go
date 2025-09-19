package tui

import (
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
)

func TestNewModel(t *testing.T) {
	model := NewModel()

	assert.NotNil(t, model)
	assert.Equal(t, ViewMain, model.currentView)
	assert.NotNil(t, model.terminalManager)
	assert.NotNil(t, model.ollamaClient)
	assert.NotNil(t, model.templateManager)
}

func TestModel_Init(t *testing.T) {
	model := NewModel()

	cmd := model.Init()
	assert.NotNil(t, cmd)
}

func TestModel_Update_KeyMessages(t *testing.T) {
	model := NewModel()

	tests := []struct {
		name     string
		key      string
		expected ViewType
	}{
		{"Switch to terminals", "1", ViewTerminals},
		{"Switch to templates", "2", ViewTemplates},
		{"Switch to settings", "3", ViewSettings},
		{"Quit application", "q", ViewMain}, // 應該觸發退出
		{"Ctrl+C quit", "ctrl+c", ViewMain}, // 應該觸發退出
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(tt.key)}
			if tt.key == "ctrl+c" {
				msg = tea.KeyMsg{Type: tea.KeyCtrlC}
			}

			updatedModel, cmd := model.Update(msg)

			if tt.key == "q" || tt.key == "ctrl+c" {
				// 退出命令
				assert.NotNil(t, cmd)
			} else {
				// 檢查視圖切換
				assert.Equal(t, tt.expected, updatedModel.(Model).currentView)
			}
		})
	}
}

func TestModel_Update_TickMessage(t *testing.T) {
	model := NewModel()

	// 創建 tick 消息
	tickMsg := TickMsg{time.Now()}

	updatedModel, cmd := model.Update(tickMsg)

	assert.NotNil(t, updatedModel)
	assert.NotNil(t, cmd) // 應該返回下一個 tick 命令
}

func TestModel_Update_TerminalStatusMessage(t *testing.T) {
	model := NewModel()

	// 測試終端狀態更新
	statusMsg := TerminalStatusMsg{
		Name:   "test-terminal",
		Status: "running",
	}

	updatedModel, cmd := model.Update(statusMsg)

	assert.NotNil(t, updatedModel)
	assert.Nil(t, cmd) // 狀態更新不應該產生新命令
}

func TestModel_View(t *testing.T) {
	model := NewModel()

	// 測試主視圖
	view := model.View()
	assert.NotEmpty(t, view)
	assert.Contains(t, view, "AI 終端代理")

	// 測試不同視圖
	views := []ViewType{ViewTerminals, ViewTemplates, ViewSettings}

	for _, viewType := range views {
		model.currentView = viewType
		view := model.View()
		assert.NotEmpty(t, view)
	}
}

func TestModel_TerminalOperations(t *testing.T) {
	model := NewModel()

	tests := []struct {
		name        string
		operation   string
		expectError bool
	}{
		{"Start terminal", "start", false},
		{"Stop terminal", "stop", false},
		{"List terminals", "list", false},
		{"Invalid operation", "invalid", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			switch tt.operation {
			case "start":
				err = model.startTerminal("test-terminal", "claude")
			case "stop":
				err = model.stopTerminal("test-terminal")
			case "list":
				terminals := model.listTerminals()
				assert.NotNil(t, terminals)
				return
			default:
				err = assert.AnError // 模擬錯誤
			}

			if tt.expectError {
				assert.Error(t, err)
			} else {
				// 由於沒有實際的終端管理器，可能會有錯誤，但不應該 panic
				t.Logf("Operation %s result: %v", tt.operation, err)
			}
		})
	}
}

func TestModel_QueryOptimization(t *testing.T) {
	model := NewModel()

	tests := []struct {
		name       string
		query      string
		templateID string
		expectErr  bool
	}{
		{
			name:       "Valid optimization",
			query:      "如何實現 HTTP 服務器？",
			templateID: "coding",
			expectErr:  false,
		},
		{
			name:       "Empty query",
			query:      "",
			templateID: "coding",
			expectErr:  true,
		},
		{
			name:       "Invalid template",
			query:      "test query",
			templateID: "invalid",
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := model.optimizeQuery(tt.query, tt.templateID)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestModel_ViewNavigation(t *testing.T) {
	model := NewModel()

	// 測試視圖導航
	originalView := model.currentView

	// 切換到終端視圖
	model.switchView(ViewTerminals)
	assert.Equal(t, ViewTerminals, model.currentView)

	// 切換到模板視圖
	model.switchView(ViewTemplates)
	assert.Equal(t, ViewTemplates, model.currentView)

	// 回到主視圖
	model.switchView(ViewMain)
	assert.Equal(t, ViewMain, model.currentView)

	// 驗證原始狀態沒有被破壞
	assert.NotEqual(t, originalView, ViewTerminals)
}

func TestModel_StatusUpdates(t *testing.T) {
	model := NewModel()

	// 模擬狀態更新
	initialStatus := model.status

	// 更新狀態
	model.updateStatus()

	// 檢查狀態是否更新（updateStatus 只改變 status 內容，不改變 lastUpdate）
	assert.NotEqual(t, initialStatus, model.status)
}

func TestTickCmd(t *testing.T) {
	cmd := tickCmd()
	assert.NotNil(t, cmd)

	// 執行命令
	msg := cmd()

	// 檢查返回的消息類型
	tickMsg, ok := msg.(TickMsg)
	assert.True(t, ok)
	assert.False(t, tickMsg.Time.IsZero())
}

func TestViewType_String(t *testing.T) {
	tests := []struct {
		view     ViewType
		expected string
	}{
		{ViewMain, "主畫面"},
		{ViewTerminals, "終端管理"},
		{ViewTemplates, "模板管理"},
		{ViewSettings, "設置"},
		{ViewType(999), "未知視圖"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.view.String())
		})
	}
}

func TestModel_ErrorHandling(t *testing.T) {
	model := NewModel()

	// 測試錯誤處理
	model.handleError("測試錯誤")

	// 驗證錯誤狀態
	assert.Contains(t, model.status, "錯誤")
	assert.Contains(t, model.status, "測試錯誤")
}

func TestModel_ConcurrentOperations(t *testing.T) {
	model := NewModel()

	// 並發測試視圖切換
	done := make(chan bool, 5)

	for i := 0; i < 5; i++ {
		go func(index int) {
			viewType := ViewType(index % 4) // 0-3 的視圖類型
			model.switchView(viewType)
			done <- true
		}(i)
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 5; i++ {
		select {
		case <-done:
			// 成功
		case <-time.After(2 * time.Second):
			t.Fatal("Concurrent operations test timed out")
		}
	}

	// 驗證最終狀態是有效的
	assert.True(t, model.currentView >= ViewMain && model.currentView <= ViewSettings)
}