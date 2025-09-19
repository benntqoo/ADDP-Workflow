package main

import (
	"github.com/spf13/cobra"
	"ai-launcher/internal/gui"
)

var guiCmd = &cobra.Command{
	Use:   "gui",
	Short: "启动图形用户界面",
	Long: `启动AI启动器的图形用户界面版本，提供：
• 直观的项目选择和配置
• 拖拽式操作体验
• 可视化的AI模型选择
• 一键启动功能`,
	Run: func(cmd *cobra.Command, args []string) {
		// 启动GUI应用
		app := gui.NewApp()
		app.Run()
	},
}