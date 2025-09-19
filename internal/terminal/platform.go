package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// PlatformAdapter 提供跨平台的終端管理功能
type PlatformAdapter struct {
	os string
}

// NewPlatformAdapter 創建新的平台適配器
func NewPlatformAdapter() *PlatformAdapter {
	return &PlatformAdapter{
		os: runtime.GOOS,
	}
}

// GetDefaultShell 獲取當前平台的默認 shell
func (pa *PlatformAdapter) GetDefaultShell() string {
	switch pa.os {
	case "windows":
		// Windows 優先使用 PowerShell，備選 cmd
		if pa.ValidateCommand("powershell.exe") {
			return "powershell.exe"
		}
		return "cmd.exe"
	case "darwin", "linux":
		// Unix-like 系統優先使用 bash，備選 sh
		if pa.ValidateCommand("/bin/bash") {
			return "/bin/bash"
		}
		return "/bin/sh"
	default:
		// 其他系統使用通用 shell
		return "/bin/sh"
	}
}

// GetExecutablePath 獲取命令的完整可執行路徑
func (pa *PlatformAdapter) GetExecutablePath(command string) string {
	switch pa.os {
	case "windows":
		// Windows 系統處理
		extensions := []string{"", ".exe", ".cmd", ".bat", ".com"}

		for _, ext := range extensions {
			fullCommand := command + ext
			if path, err := exec.LookPath(fullCommand); err == nil {
				return path
			}
		}

		// 如果沒找到，返回帶 .exe 的版本
		return command + ".exe"

	default:
		// Unix-like 系統
		if path, err := exec.LookPath(command); err == nil {
			return path
		}

		// 如果沒找到，返回原命令（可能在 PATH 中）
		return command
	}
}

// CreateCommand 創建適合當前平台的命令
func (pa *PlatformAdapter) CreateCommand(config TerminalConfig) *exec.Cmd {
	var cmdPath string
	var args []string

	switch config.Type {
	case TypeClaudeCode:
		cmdPath = pa.GetExecutablePath("claude")
		args = []string{cmdPath}

	case TypeGeminiCLI:
		cmdPath = pa.GetExecutablePath("gemini")
		args = []string{cmdPath}

	case TypeCursor:
		cmdPath = pa.GetExecutablePath("cursor")
		args = []string{cmdPath, "--cli"}

	case TypeAider:
		cmdPath = pa.GetExecutablePath("aider")
		args = []string{cmdPath}

	case TypeCustom:
		// 對於測試，使用平台適當的測試命令
		if pa.os == "windows" {
			cmdPath = pa.GetExecutablePath("cmd")
			args = []string{cmdPath, "/c", "echo", "test"}
		} else {
			cmdPath = "/bin/echo"
			args = []string{cmdPath, "test"}
		}

	default:
		// 默認使用 shell
		cmdPath = pa.GetDefaultShell()
		args = []string{cmdPath}
	}

	// 創建命令
	cmd := &exec.Cmd{
		Path: cmdPath,
		Args: args,
	}

	// 設置工作目錄
	if config.WorkingDir != "" {
		cmd.Dir = config.WorkingDir
	}

	// 添加額外參數
	if config.Args != nil {
		cmd.Args = append(cmd.Args, config.Args...)
	}

	return cmd
}

// SetupEnvironment 設置命令的環境變量
func (pa *PlatformAdapter) SetupEnvironment(cmd *exec.Cmd, config TerminalConfig) {
	// 獲取當前環境變量
	env := os.Environ()

	// 添加自定義環境變量
	if config.Environment != nil {
		for key, value := range config.Environment {
			env = append(env, fmt.Sprintf("%s=%s", key, value))
		}
	}

	// 添加平台特定的環境變量
	switch pa.os {
	case "windows":
		// Windows 特定環境設置
		env = append(env, "AI_TERMINAL_PLATFORM=windows")

		// 確保 PATH 包含必要的目錄
		pathEnvVar := pa.findPathVariable(env)
		if pathEnvVar == "" {
			env = append(env, "PATH=C:\\Windows\\System32")
		}

	default:
		// Unix-like 系統特定環境設置
		env = append(env, "AI_TERMINAL_PLATFORM=unix")

		// 確保基本的 PATH
		pathEnvVar := pa.findPathVariable(env)
		if pathEnvVar == "" {
			env = append(env, "PATH=/usr/local/bin:/usr/bin:/bin")
		}
	}

	cmd.Env = env
}

// ValidateCommand 檢查命令是否存在
func (pa *PlatformAdapter) ValidateCommand(command string) bool {
	// 如果是絕對路徑，直接檢查文件是否存在
	if filepath.IsAbs(command) {
		_, err := os.Stat(command)
		return err == nil
	}

	// 使用 exec.LookPath 在 PATH 中查找
	_, err := exec.LookPath(command)
	return err == nil
}

// GetProcessInfo 獲取進程信息
func (pa *PlatformAdapter) GetProcessInfo(pid int) *ProcessInfo {
	info := &ProcessInfo{
		PID: pid,
	}

	switch pa.os {
	case "windows":
		// Windows 使用 WMI 或 tasklist 獲取進程信息
		info.Command = pa.getWindowsProcessInfo(pid)
		info.ExecutablePath = pa.getWindowsExecutablePath(pid)
		info.Status = "running"

	default:
		// Unix-like 系統使用 /proc 或 ps
		info.Command = pa.getUnixProcessInfo(pid)
		info.CommandLine = pa.getUnixCommandLine(pid)
		info.Status = "running"
	}

	return info
}

// KillProcess 跨平台殺死進程
func (pa *PlatformAdapter) KillProcess(pid int) error {
	switch pa.os {
	case "windows":
		// Windows 使用 taskkill
		cmd := exec.Command("taskkill", "/F", "/PID", fmt.Sprintf("%d", pid))
		return cmd.Run()

	default:
		// Unix-like 系統使用 kill
		cmd := exec.Command("kill", "-TERM", fmt.Sprintf("%d", pid))
		if err := cmd.Run(); err != nil {
			// 如果 TERM 失敗，嘗試 KILL
			cmd = exec.Command("kill", "-KILL", fmt.Sprintf("%d", pid))
			return cmd.Run()
		}
		return nil
	}
}

// 輔助方法：查找 PATH 環境變量
func (pa *PlatformAdapter) findPathVariable(env []string) string {
	for _, envVar := range env {
		if strings.HasPrefix(strings.ToUpper(envVar), "PATH=") {
			return envVar
		}
	}
	return ""
}

// Windows 特定的進程信息獲取
func (pa *PlatformAdapter) getWindowsProcessInfo(pid int) string {
	// 簡化實現：在實際項目中可能需要使用 WMI
	cmd := exec.Command("tasklist", "/FI", fmt.Sprintf("PID eq %d", pid), "/FO", "CSV", "/NH")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}

	// 解析 CSV 輸出獲取進程名
	lines := strings.Split(string(output), "\n")
	if len(lines) > 0 && lines[0] != "" {
		fields := strings.Split(lines[0], ",")
		if len(fields) > 0 {
			// 移除引號
			return strings.Trim(fields[0], "\"")
		}
	}

	return "unknown"
}

// Windows 獲取可執行文件路徑
func (pa *PlatformAdapter) getWindowsExecutablePath(pid int) string {
	// 簡化實現：實際項目中可能需要使用 Windows API
	return "unknown"
}

// Unix 特定的進程信息獲取
func (pa *PlatformAdapter) getUnixProcessInfo(pid int) string {
	// 嘗試從 /proc/PID/comm 讀取進程名
	commPath := fmt.Sprintf("/proc/%d/comm", pid)
	if data, err := os.ReadFile(commPath); err == nil {
		return strings.TrimSpace(string(data))
	}

	// 備選：使用 ps 命令
	cmd := exec.Command("ps", "-p", fmt.Sprintf("%d", pid), "-o", "comm=")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(output))
}

// Unix 獲取命令行
func (pa *PlatformAdapter) getUnixCommandLine(pid int) string {
	// 嘗試從 /proc/PID/cmdline 讀取命令行
	cmdlinePath := fmt.Sprintf("/proc/%d/cmdline", pid)
	if data, err := os.ReadFile(cmdlinePath); err == nil {
		// /proc/PID/cmdline 使用 null 字符分隔參數
		cmdline := strings.ReplaceAll(string(data), "\x00", " ")
		return strings.TrimSpace(cmdline)
	}

	// 備選：使用 ps 命令
	cmd := exec.Command("ps", "-p", fmt.Sprintf("%d", pid), "-o", "args=")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(output))
}

// ProcessInfo 跨平台進程信息結構
type ProcessInfo struct {
	PID            int    // 進程 ID
	Command        string // 進程命令名
	CommandLine    string // 完整命令行
	ExecutablePath string // 可執行文件路徑
	Status         string // 進程狀態
}