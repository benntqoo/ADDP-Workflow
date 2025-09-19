package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"ai-launcher/internal/tui"
)

var rootCmd = &cobra.Command{
	Use:   "ai-launcher",
	Short: "AI启动器 - 快速启动各种AI编程助手",
	Long: `AI启动器是一个智能的多AI工具启动器，提供：
• 项目目录快速选择
• 多种AI模型支持 (Claude Code, Gemini CLI, Codex)
• YOLO模式支持，跳过安全确认
• 配置自动保存和最近项目记录
• 一键启动新终端窗口`,
	Run: func(cmd *cobra.Command, args []string) {
		// 启动新的UI界面
		model := tui.NewNewUIModel()
		p := tea.NewProgram(model, tea.WithAltScreen())

		if _, err := p.Run(); err != nil {
			log.Fatal("启动UI失败:", err)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本信息",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ai-launcher v2.0.0")
		fmt.Println("AI启动器 - 智能多AI工具启动器")
		fmt.Println("支持项目管理、多模型选择、YOLO模式")
	},
}

var listCmd = &cobra.Command{
	Use:   "list-models",
	Short: "列出支持的AI模型",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("支持的AI模型：")
		fmt.Println("  🤖 Claude Code    - claude / claude --dangerously-skip-permissions")
		fmt.Println("  💎 Gemini CLI     - gemini / gemini --yolo")
		fmt.Println("  🔧 Codex          - codex / codex --dangerously-bypass-approvals-and-sandbox")
		fmt.Println("")
		fmt.Println("运行模式：")
		fmt.Println("  🛡️  普通模式 - 需要用户确认操作")
		fmt.Println("  🚀 YOLO模式 - 跳过安全检查，自动执行")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(guiCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}