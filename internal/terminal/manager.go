package terminal

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
	"sync"
	"time"
)

// TerminalManager 實現 Manager 接口
type TerminalManager struct {
	terminals map[string]*Terminal
	mu        sync.RWMutex
	healthy   bool
}

// NewTerminalManager 創建一個新的終端管理器
func NewTerminalManager() *TerminalManager {
	return &TerminalManager{
		terminals: make(map[string]*Terminal),
		healthy:   true,
	}
}

// StartTerminal 啟動指定的終端
func (tm *TerminalManager) StartTerminal(config TerminalConfig) error {
	return tm.StartTerminalWithContext(context.Background(), config)
}

// StartTerminalWithContext 使用上下文啟動指定的終端
func (tm *TerminalManager) StartTerminalWithContext(ctx context.Context, config TerminalConfig) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	// 檢查終端是否已存在
	if _, exists := tm.terminals[config.Name]; exists {
		return fmt.Errorf("terminal '%s' already exists", config.Name)
	}

	// 創建新的終端實例
	terminal := &Terminal{
		Name:     config.Name,
		Type:     config.Type,
		Status:   StatusStarting,
		LastUsed: time.Now().Unix(),
	}

	// 創建命令
	cmd := tm.createCommand(config)
	if cmd == nil {
		return fmt.Errorf("failed to create command for terminal type %s", config.Type.String())
	}

	terminal.Process = cmd

	// 設置輸入輸出管道
	if err := tm.setupPipes(terminal); err != nil {
		return fmt.Errorf("failed to setup pipes: %w", err)
	}

	// 在後台啟動進程
	if err := tm.startProcess(ctx, terminal); err != nil {
		return fmt.Errorf("failed to start process: %w", err)
	}

	// 添加到管理器
	tm.terminals[config.Name] = terminal
	terminal.SetStatus(StatusRunning)

	return nil
}

// StopTerminal 停止指定名稱的終端
func (tm *TerminalManager) StopTerminal(name string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	terminal, exists := tm.terminals[name]
	if !exists {
		return fmt.Errorf("terminal '%s' not found", name)
	}

	// 設置停止狀態
	terminal.SetStatus(StatusStopping)

	// 停止進程
	if terminal.Process != nil && terminal.Process.Process != nil {
		if err := terminal.Process.Process.Kill(); err != nil {
			terminal.SetStatus(StatusError)
			return fmt.Errorf("failed to kill process: %w", err)
		}
	}

	// 等待進程結束
	if terminal.Process != nil {
		_ = terminal.Process.Wait()
	}

	terminal.SetStatus(StatusStopped)
	return nil
}

// SendCommand 向指定終端發送命令
func (tm *TerminalManager) SendCommand(name string, command string) error {
	tm.mu.RLock()
	terminal, exists := tm.terminals[name]
	tm.mu.RUnlock()

	if !exists {
		return fmt.Errorf("terminal '%s' not found", name)
	}

	if !terminal.IsRunning() {
		return fmt.Errorf("terminal '%s' is not running", name)
	}

	// 更新最後使用時間
	terminal.mu.Lock()
	terminal.LastUsed = time.Now().Unix()
	terminal.mu.Unlock()

	// 發送命令
	if terminal.Stdin == nil {
		return fmt.Errorf("terminal '%s' stdin not available", name)
	}

	if _, err := terminal.Stdin.WriteString(command + "\n"); err != nil {
		return fmt.Errorf("failed to write command: %w", err)
	}

	if err := terminal.Stdin.Flush(); err != nil {
		return fmt.Errorf("failed to flush command: %w", err)
	}

	return nil
}

// GetTerminal 獲取指定名稱的終端
func (tm *TerminalManager) GetTerminal(name string) (*Terminal, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	terminal, exists := tm.terminals[name]
	return terminal, exists
}

// ListTerminals 列出所有終端
func (tm *TerminalManager) ListTerminals() []*Terminal {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	terminals := make([]*Terminal, 0, len(tm.terminals))
	for _, terminal := range tm.terminals {
		terminals = append(terminals, terminal)
	}

	return terminals
}

// IsHealthy 檢查終端管理器是否健康
func (tm *TerminalManager) IsHealthy() bool {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return tm.healthy
}

// createCommand 創建對應類型的命令
func (tm *TerminalManager) createCommand(config TerminalConfig) *exec.Cmd {
	var cmd *exec.Cmd

	switch config.Type {
	case TypeClaudeCode:
		cmd = exec.Command("claude")
	case TypeGeminiCLI:
		cmd = exec.Command("gemini")
	case TypeCursor:
		cmd = exec.Command("cursor", "--cli")
	case TypeAider:
		cmd = exec.Command("aider")
	case TypeCustom:
		// 對於測試，使用一個簡單的命令
		cmd = exec.Command("echo", "test")
	default:
		return nil
	}

	// 設置工作目錄
	if config.WorkingDir != "" {
		cmd.Dir = config.WorkingDir
	}

	// 設置環境變量
	if config.Environment != nil {
		for key, value := range config.Environment {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
		}
	}

	// 添加額外參數
	if config.Args != nil {
		cmd.Args = append(cmd.Args, config.Args...)
	}

	return cmd
}

// setupPipes 設置進程的輸入輸出管道
func (tm *TerminalManager) setupPipes(terminal *Terminal) error {
	// 設置標準輸入管道
	stdin, err := terminal.Process.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdin pipe: %w", err)
	}
	terminal.Stdin = bufio.NewWriter(stdin)

	// 設置標準輸出管道
	stdout, err := terminal.Process.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %w", err)
	}
	terminal.Stdout = bufio.NewScanner(stdout)

	return nil
}

// startProcess 在後台啟動進程
func (tm *TerminalManager) startProcess(ctx context.Context, terminal *Terminal) error {
	// 創建一個帶取消功能的上下文
	processCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 設置超時
	done := make(chan error, 1)

	go func() {
		done <- terminal.Process.Start()
	}()

	select {
	case err := <-done:
		if err != nil {
			terminal.SetStatus(StatusError)
			return err
		}
		return nil
	case <-processCtx.Done():
		terminal.SetStatus(StatusError)
		return processCtx.Err()
	case <-time.After(5 * time.Second): // 5秒超時
		terminal.SetStatus(StatusError)
		return fmt.Errorf("timeout starting process")
	}
}