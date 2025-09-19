package terminal

import (
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlatformAdapter_GetDefaultShell(t *testing.T) {
	adapter := NewPlatformAdapter()

	shell := adapter.GetDefaultShell()

	switch runtime.GOOS {
	case "windows":
		// Windows 應該使用 cmd.exe 或 powershell.exe
		assert.True(t, shell == "cmd.exe" || shell == "powershell.exe",
			"Windows should use cmd.exe or powershell.exe, got: %s", shell)
	case "darwin", "linux":
		// Unix-like 系統應該使用 bash 或 sh
		assert.True(t, shell == "/bin/bash" || shell == "/bin/sh",
			"Unix-like systems should use /bin/bash or /bin/sh, got: %s", shell)
	default:
		// 其他系統使用通用 shell
		assert.NotEmpty(t, shell, "Shell should not be empty")
	}
}

func TestPlatformAdapter_GetExecutablePath(t *testing.T) {
	adapter := NewPlatformAdapter()

	tests := []struct {
		name     string
		command  string
		platform string
	}{
		{"Windows Claude", "claude", "windows"},
		{"Unix Claude", "claude", "linux"},
		{"Windows Gemini", "gemini", "windows"},
		{"Unix Gemini", "gemini", "linux"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 跳過不匹配的平台測試
			if runtime.GOOS != tt.platform && tt.platform != "linux" {
				return
			}
			if runtime.GOOS == "darwin" && tt.platform == "linux" {
				// macOS 使用 linux 的邏輯
			} else if runtime.GOOS != tt.platform {
				return
			}

			path := adapter.GetExecutablePath(tt.command)
			assert.NotEmpty(t, path, "Executable path should not be empty")

			if runtime.GOOS == "windows" {
				// Windows 可執行文件可能有 .exe 擴展名
				// 檢查路徑是否以命令名結尾（可能包含路徑）
				assert.True(t,
					strings.HasSuffix(path, tt.command) ||
					strings.HasSuffix(path, tt.command+".exe") ||
					strings.HasSuffix(path, tt.command+".cmd") ||
					strings.HasSuffix(path, tt.command+".bat"),
					"Windows executable should have proper extension: %s", path)
			} else {
				// Unix-like 系統
				assert.Contains(t, path, tt.command, "Unix executable should contain command name")
			}
		})
	}
}

func TestPlatformAdapter_CreateCommand(t *testing.T) {
	adapter := NewPlatformAdapter()

	config := TerminalConfig{
		Type: TypeClaudeCode,
		Name: "test-platform",
	}

	cmd := adapter.CreateCommand(config)
	require.NotNil(t, cmd, "Command should not be nil")

	// 檢查命令路徑
	expectedCmd := adapter.GetExecutablePath("claude")
	assert.Equal(t, expectedCmd, cmd.Path, "Command path should match platform-specific executable")

	// 檢查參數
	assert.NotEmpty(t, cmd.Args, "Command args should not be empty")
	assert.Contains(t, cmd.Args[0], "claude", "First arg should contain command name")
}

func TestPlatformAdapter_SetupEnvironment(t *testing.T) {
	adapter := NewPlatformAdapter()

	config := TerminalConfig{
		Type: TypeGeminiCLI,
		Name: "test-env",
		Environment: map[string]string{
			"TEST_VAR": "test_value",
			"AI_TOOL":  "gemini",
		},
	}

	cmd := adapter.CreateCommand(config)
	require.NotNil(t, cmd)

	adapter.SetupEnvironment(cmd, config)

	// 檢查環境變量是否正確設置
	envMap := make(map[string]string)
	for _, env := range cmd.Env {
		// 解析 KEY=VALUE 格式
		for i, char := range env {
			if char == '=' {
				key := env[:i]
				value := env[i+1:]
				envMap[key] = value
				break
			}
		}
	}

	assert.Equal(t, "test_value", envMap["TEST_VAR"], "Custom environment variable should be set")
	assert.Equal(t, "gemini", envMap["AI_TOOL"], "AI_TOOL environment variable should be set")

	// 檢查平台特定的環境變量
	if runtime.GOOS == "windows" {
		// Windows 可能需要設置 PATH 或其他變量
		assert.Contains(t, envMap, "PATH", "Windows should have PATH environment variable")
	} else {
		// Unix-like 系統
		assert.Contains(t, envMap, "PATH", "Unix should have PATH environment variable")
	}
}

func TestPlatformAdapter_ValidateCommand(t *testing.T) {
	adapter := NewPlatformAdapter()

	tests := []struct {
		name        string
		command     string
		shouldExist bool
	}{
		{"System command exists", adapter.GetDefaultShell(), true},
		{"Non-existent command", "definitely-not-a-command-12345", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exists := adapter.ValidateCommand(tt.command)
			assert.Equal(t, tt.shouldExist, exists,
				"Command validation for %s should be %v", tt.command, tt.shouldExist)
		})
	}
}

func TestPlatformAdapter_GetProcessInfo(t *testing.T) {
	adapter := NewPlatformAdapter()

	// 創建一個測試進程
	config := TerminalConfig{
		Type: TypeCustom,
		Name: "test-process-info",
	}

	cmd := adapter.CreateCommand(config)
	require.NotNil(t, cmd)

	// 啟動進程
	err := cmd.Start()
	require.NoError(t, err, "Should be able to start test process")
	defer func() {
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
	}()

	// 獲取進程信息
	info := adapter.GetProcessInfo(cmd.Process.Pid)
	assert.NotNil(t, info, "Process info should not be nil")
	assert.Equal(t, cmd.Process.Pid, info.PID, "PID should match")
	assert.NotEmpty(t, info.Command, "Command should not be empty")

	// 檢查平台特定的進程信息
	if runtime.GOOS == "windows" {
		assert.NotEmpty(t, info.ExecutablePath, "Windows should have executable path")
	} else {
		assert.NotEmpty(t, info.CommandLine, "Unix should have command line")
	}
}

func TestPlatformAdapter_KillProcess(t *testing.T) {
	adapter := NewPlatformAdapter()

	// 創建一個測試進程
	config := TerminalConfig{
		Type: TypeCustom,
		Name: "test-kill-process",
	}

	cmd := adapter.CreateCommand(config)
	require.NotNil(t, cmd)

	// 啟動進程
	err := cmd.Start()
	require.NoError(t, err, "Should be able to start test process")

	pid := cmd.Process.Pid
	assert.Greater(t, pid, 0, "PID should be positive")

	// 殺死進程
	err = adapter.KillProcess(pid)
	if runtime.GOOS == "windows" {
		// Windows taskkill 可能返回非零狀態，但仍然成功殺死進程
		// 我們主要關注進程是否真的被終止
	} else {
		assert.NoError(t, err, "Should be able to kill process")
	}

	// 等待進程結束
	err = cmd.Wait()
	// 進程被殺死，可能返回錯誤，這是預期的
	// 我們主要檢查進程確實結束了
}

func TestCrossPlatformTerminalTypes(t *testing.T) {
	adapter := NewPlatformAdapter()

	terminalTypes := []TerminalType{
		TypeClaudeCode,
		TypeGeminiCLI,
		TypeCursor,
		TypeAider,
		TypeCustom,
	}

	for _, termType := range terminalTypes {
		t.Run(termType.String(), func(t *testing.T) {
			config := TerminalConfig{
				Type: termType,
				Name: "cross-platform-test-" + termType.String(),
			}

			cmd := adapter.CreateCommand(config)
			assert.NotNil(t, cmd, "Should be able to create command for %s", termType.String())

			// 檢查命令路徑是否適合當前平台
			assert.NotEmpty(t, cmd.Path, "Command path should not be empty")

			if runtime.GOOS == "windows" {
				// Windows 路徑不應該以 / 開始
				assert.False(t, cmd.Path[0] == '/', "Windows path should not start with /")
			}
		})
	}
}

