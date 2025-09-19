package terminal

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTerminalManager(t *testing.T) {
	manager := NewTerminalManager()

	assert.NotNil(t, manager)
	assert.True(t, manager.IsHealthy())
	assert.Empty(t, manager.ListTerminals())
}

func TestTerminalManager_StartTerminal(t *testing.T) {
	manager := NewTerminalManager()

	config := TerminalConfig{
		Type:       TypeClaudeCode,
		Name:       "test-claude",
		WorkingDir: "/tmp",
	}

	// 測試啟動終端
	err := manager.StartTerminal(config)
	assert.NoError(t, err)

	// 檢查終端是否被添加
	terminals := manager.ListTerminals()
	assert.Len(t, terminals, 1)
	assert.Equal(t, "test-claude", terminals[0].Name)
	assert.Equal(t, TypeClaudeCode, terminals[0].Type)

	// 測試重複啟動同名終端應該失敗
	err = manager.StartTerminal(config)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")
}

func TestTerminalManager_GetTerminal(t *testing.T) {
	manager := NewTerminalManager()

	// 測試獲取不存在的終端
	terminal, exists := manager.GetTerminal("non-existent")
	assert.Nil(t, terminal)
	assert.False(t, exists)

	// 啟動一個終端
	config := TerminalConfig{
		Type: TypeCustom, // 使用 Custom 類型避免依賴外部命令
		Name: "test-gemini",
	}
	err := manager.StartTerminal(config)
	require.NoError(t, err)

	// 測試獲取存在的終端
	terminal, exists = manager.GetTerminal("test-gemini")
	assert.NotNil(t, terminal)
	assert.True(t, exists)
	assert.Equal(t, "test-gemini", terminal.Name)
	assert.Equal(t, TypeCustom, terminal.Type)
}

func TestTerminalManager_StopTerminal(t *testing.T) {
	manager := NewTerminalManager()

	// 測試停止不存在的終端
	err := manager.StopTerminal("non-existent")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	// 啟動一個終端
	config := TerminalConfig{
		Type: TypeCustom, // 使用 Custom 類型避免依賴外部命令
		Name: "test-cursor",
	}
	err = manager.StartTerminal(config)
	require.NoError(t, err)

	// 驗證終端正在運行
	terminal, exists := manager.GetTerminal("test-cursor")
	require.True(t, exists)
	assert.Equal(t, StatusRunning, terminal.GetStatus())

	// 停止終端
	err = manager.StopTerminal("test-cursor")
	assert.NoError(t, err)

	// 驗證終端已停止
	terminal, exists = manager.GetTerminal("test-cursor")
	assert.True(t, exists) // 終端對象還存在
	assert.Equal(t, StatusStopped, terminal.GetStatus())
}

func TestTerminalManager_SendCommand(t *testing.T) {
	manager := NewTerminalManager()

	// 測試向不存在的終端發送命令
	err := manager.SendCommand("non-existent", "test command")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	// 啟動一個終端
	config := TerminalConfig{
		Type: TypeCustom, // 使用 Custom 類型避免依賴外部命令
		Name: "test-aider",
	}
	err = manager.StartTerminal(config)
	require.NoError(t, err)

	// 測試發送命令
	err = manager.SendCommand("test-aider", "/help")
	assert.NoError(t, err)

	// 測試向已停止的終端發送命令
	err = manager.StopTerminal("test-aider")
	require.NoError(t, err)

	err = manager.SendCommand("test-aider", "another command")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not running")
}

func TestTerminalManager_ListTerminals(t *testing.T) {
	manager := NewTerminalManager()

	// 初始狀態應該為空
	terminals := manager.ListTerminals()
	assert.Empty(t, terminals)

	// 啟動多個終端
	configs := []TerminalConfig{
		{Type: TypeCustom, Name: "claude-1"},
		{Type: TypeCustom, Name: "gemini-1"},
		{Type: TypeCustom, Name: "cursor-1"},
	}

	for _, config := range configs {
		err := manager.StartTerminal(config)
		require.NoError(t, err)
	}

	// 檢查終端列表
	terminals = manager.ListTerminals()
	assert.Len(t, terminals, 3)

	// 檢查每個終端的信息
	names := make([]string, len(terminals))
	for i, terminal := range terminals {
		names[i] = terminal.Name
	}
	assert.Contains(t, names, "claude-1")
	assert.Contains(t, names, "gemini-1")
	assert.Contains(t, names, "cursor-1")
}

func TestTerminalManager_ConcurrentOperations(t *testing.T) {
	manager := NewTerminalManager()

	// 並發啟動多個終端
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(index int) {
			config := TerminalConfig{
				Type: TypeCustom, // 使用 Custom 類型避免依賴外部命令
				Name: fmt.Sprintf("concurrent-terminal-%d", index),
			}
			err := manager.StartTerminal(config)
			assert.NoError(t, err)
			done <- true
		}(i)
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 10; i++ {
		<-done
	}

	// 檢查所有終端都被創建
	terminals := manager.ListTerminals()
	assert.Len(t, terminals, 10)
}

func TestTerminalManager_Timeout(t *testing.T) {
	manager := NewTerminalManager()

	config := TerminalConfig{
		Type: TypeCustom,
		Name: "timeout-test",
	}

	// 創建一個帶超時的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// 測試超時啟動（這個測試可能需要模擬慢啟動的場景）
	err := manager.StartTerminalWithContext(ctx, config)
	// 根據實現，這可能成功或超時，主要是測試上下文傳遞
	if err != nil {
		assert.Contains(t, err.Error(), "timeout")
	}
}

func TestTerminalManager_IsHealthy(t *testing.T) {
	manager := NewTerminalManager()

	// 新創建的管理器應該是健康的
	assert.True(t, manager.IsHealthy())

	// 啟動一些終端
	config := TerminalConfig{
		Type: TypeCustom, // 使用 Custom 類型避免依賴外部命令
		Name: "health-test",
	}
	err := manager.StartTerminal(config)
	require.NoError(t, err)

	// 管理器仍然應該是健康的
	assert.True(t, manager.IsHealthy())
}

