package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "claude-proxy",
	Short: "AI Terminal Proxy - 統一所有 AI CLI 工具的終端界面",
	Long: `AI Terminal Proxy 是一個智能終端代理，提供：
• 統一的 GUI 界面管理多個 AI 終端
• Ollama 本地查詢優化，節省 30-50% token
• 1-4 數字鍵快速在不同 AI 工具間切換
• 支援 Claude Code、Gemini CLI、Cursor、Aider`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🤖 AI Terminal Proxy v1.0")
		fmt.Println("啟動 TUI 界面...")
		// TODO: 啟動 Bubble Tea TUI 界面
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "顯示版本信息",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("claude-proxy v1.0.0")
		fmt.Println("基於雙軌架構：Golang 終端代理 + Python MCP 工具")
	},
}

var listCmd = &cobra.Command{
	Use:   "list-terminals",
	Short: "列出可用的終端類型",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("可用終端類型：")
		fmt.Println("  1. claude     - Claude Code CLI")
		fmt.Println("  2. gemini     - Gemini CLI")
		fmt.Println("  3. cursor     - Cursor CLI")
		fmt.Println("  4. aider      - Aider CLI")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(listCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}