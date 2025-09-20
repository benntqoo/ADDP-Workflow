package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"ai-launcher/internal/gui"
)

func main() {
	// 强制将输出重定向到文件以便调试
	logFile, err := os.OpenFile("ai-launcher-debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	// 添加错误恢复处理
	defer func() {
		if r := recover(); r != nil {
			errorMsg := fmt.Sprintf("程序发生错误: %v\n运行环境: %s/%s\n", r, runtime.GOOS, runtime.GOARCH)

			// 同时写入日志文件和控制台
			log.Printf("PANIC: %s", errorMsg)
			fmt.Print(errorMsg)
			fmt.Println("\n按回车键退出...")
			fmt.Scanln()
			os.Exit(1)
		}
	}()

	// 设置日志输出
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Windows环境检查
	log.Printf("程序启动 - Go版本: %s, 系统: %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	if runtime.GOOS == "windows" {
		log.Println("检测到Windows环境")

		// 检查是否在控制台模式下运行
		if len(os.Args) > 1 && os.Args[1] == "--console" {
			log.Println("控制台模式启动")
		}
	}

	log.Printf("开始创建主窗口...")

	// 创建并启动主窗口
	mainWindow := gui.NewMainWindow()
	if mainWindow == nil {
		errMsg := "错误: 无法创建主窗口\n可能的原因:\n1. OpenGL驱动问题\n2. 显卡驱动过旧\n3. 缺少运行时库\n"
		log.Printf("ERROR: %s", errMsg)
		fmt.Print(errMsg)
		fmt.Println("\n按回车键退出...")
		fmt.Scanln()
		return
	}

	log.Printf("主窗口创建成功，开始运行GUI...")

	// 运行GUI
	mainWindow.Run()

	log.Printf("GUI退出")
}