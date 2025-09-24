package main

import (
    "fmt"
    "log"
    "os"
    "runtime"
    "runtime/debug"

    "ai-launcher/internal/gui"
)

func main() {
    // log to file for diagnostics
    logFile, err := os.OpenFile("ai-launcher-debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err == nil {
        defer logFile.Close()
        log.SetOutput(logFile)
    }

    // panic guard with stack
    defer func() {
        if r := recover(); r != nil {
            errorMsg := fmt.Sprintf("程序发生错误: %v\n运行环境: %s/%s\n", r, runtime.GOOS, runtime.GOARCH)
            log.Printf("PANIC: %s\nSTACK:\n%s", errorMsg, string(debug.Stack()))
            _ = os.WriteFile("ai-launcher-crash.log", append([]byte(errorMsg+"\n\n"), debug.Stack()...), 0644)
            fmt.Print(errorMsg)
            fmt.Println("\n按回车退出...")
            fmt.Scanln()
            os.Exit(1)
        }
    }()

    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Printf("程序启动 - Go版本: %s, 系统: %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)

    if runtime.GOOS == "windows" {
        log.Println("检测到Windows环境")
        if len(os.Args) > 1 && os.Args[1] == "--console" {
            log.Println("console mode")
        }
    }

    log.Printf("开始创建主窗口...")

    // Windows 下优先设置 CJK 字体，避免中文显示为方框/乱码
    if runtime.GOOS == "windows" {
        gui.EnsureCJKFont()
    }

    // create and run GUI
    mainWindow := gui.NewMainWindow()
    if mainWindow == nil {
        errMsg := "错误: 无法创建主窗口\n可能原因:\n1. OpenGL/显卡驱动问题\n2. 显卡驱动过旧\n3. 缺少运行库\n"
        log.Printf("ERROR: %s", errMsg)
        fmt.Print(errMsg)
        fmt.Println("\n按回车退出...")
        fmt.Scanln()
        return
    }

    log.Printf("主窗口创建成功，启动GUI...")
    mainWindow.Run()
    log.Printf("GUI退出")
}
