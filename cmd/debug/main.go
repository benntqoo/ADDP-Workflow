package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	// 創建日志文件
	logFile, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("無法創建日志文件: %v\n", err)
	} else {
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprintf("程序錯誤: %v", r)
			log.Println(msg)
			fmt.Println(msg)
			fmt.Println("按回車鍵退出...")
			fmt.Scanln()
		}
	}()

	log.Printf("=== 調試程序啟動 ===")
	log.Printf("Go版本: %s", runtime.Version())
	log.Printf("系統: %s/%s", runtime.GOOS, runtime.GOARCH)

	fmt.Println("調試程序啟動...")
	fmt.Printf("Go版本: %s\n", runtime.Version())
	fmt.Printf("系統: %s/%s\n", runtime.GOOS, runtime.GOARCH)

	// 測試Fyne基本初始化
	log.Printf("嘗試導入Fyne...")

	// 動態測試導入
	testFyne()

	log.Printf("=== 調試程序完成 ===")
	fmt.Println("調試完成，檢查debug.log文件")
	fmt.Println("按回車鍵退出...")
	fmt.Scanln()
}

func testFyne() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Fyne初始化失敗: %v", r)
			fmt.Printf("Fyne初始化失敗: %v\n", r)
		}
	}()

	log.Printf("測試基本功能...")
	fmt.Println("測試基本功能...")
}