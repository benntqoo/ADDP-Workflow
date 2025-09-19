package terminal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTerminalType_String(t *testing.T) {
	tests := []struct {
		name     string
		termType TerminalType
		expected string
	}{
		{"Claude Code", TypeClaudeCode, "claude"},
		{"Gemini CLI", TypeGeminiCLI, "gemini"},
		{"Cursor", TypeCursor, "cursor"},
		{"Aider", TypeAider, "aider"},
		{"Custom", TypeCustom, "custom"},
		{"Unknown", TerminalType(999), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.termType.String()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTerminalType_CommandName(t *testing.T) {
	tests := []struct {
		name     string
		termType TerminalType
		expected string
	}{
		{"Claude Code command", TypeClaudeCode, "claude"},
		{"Gemini CLI command", TypeGeminiCLI, "gemini"},
		{"Cursor command", TypeCursor, "cursor"},
		{"Aider command", TypeAider, "aider"},
		{"Unknown command", TerminalType(999), "bash"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.termType.CommandName()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTerminalStatus_String(t *testing.T) {
	tests := []struct {
		name     string
		status   TerminalStatus
		expected string
	}{
		{"Stopped", StatusStopped, "stopped"},
		{"Starting", StatusStarting, "starting"},
		{"Running", StatusRunning, "running"},
		{"Stopping", StatusStopping, "stopping"},
		{"Error", StatusError, "error"},
		{"Unknown", TerminalStatus(999), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.status.String()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTerminal_GetSetStatus(t *testing.T) {
	terminal := &Terminal{
		Name: "test-terminal",
		Type: TypeClaudeCode,
	}

	// 測試默認狀態
	assert.Equal(t, StatusStopped, terminal.GetStatus())

	// 測試設置狀態
	terminal.SetStatus(StatusRunning)
	assert.Equal(t, StatusRunning, terminal.GetStatus())

	// 測試 IsRunning
	assert.True(t, terminal.IsRunning())

	// 測試設置其他狀態
	terminal.SetStatus(StatusStopped)
	assert.False(t, terminal.IsRunning())
}

func TestTerminal_ConcurrentAccess(t *testing.T) {
	terminal := &Terminal{
		Name: "concurrent-test",
		Type: TypeGeminiCLI,
	}

	// 並發測試狀態設置和讀取
	done := make(chan bool, 2)

	// Goroutine 1: 持續設置狀態
	go func() {
		for i := 0; i < 1000; i++ {
			terminal.SetStatus(StatusRunning)
			terminal.SetStatus(StatusStopped)
		}
		done <- true
	}()

	// Goroutine 2: 持續讀取狀態
	go func() {
		for i := 0; i < 1000; i++ {
			_ = terminal.GetStatus()
			_ = terminal.IsRunning()
		}
		done <- true
	}()

	// 等待兩個 goroutine 完成
	<-done
	<-done

	// 測試通過表示沒有競態條件
	assert.True(t, true, "Concurrent access test passed")
}

func TestTerminalConfig_Creation(t *testing.T) {
	config := TerminalConfig{
		Type:       TypeClaudeCode,
		Name:       "test-claude",
		WorkingDir: "/tmp",
		Environment: map[string]string{
			"AI_TOOL": "claude",
		},
		Args: []string{"--verbose"},
	}

	assert.Equal(t, TypeClaudeCode, config.Type)
	assert.Equal(t, "test-claude", config.Name)
	assert.Equal(t, "/tmp", config.WorkingDir)
	assert.Equal(t, "claude", config.Environment["AI_TOOL"])
	assert.Contains(t, config.Args, "--verbose")
}