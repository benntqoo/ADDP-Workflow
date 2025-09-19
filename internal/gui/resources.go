package gui

import "fyne.io/fyne/v2"

// 这里需要嵌入应用图标
// 可以使用 fyne bundle 工具生成资源文件
// 例如: fyne bundle -o resources.go icon.png

// 临时使用默认图标
var resourceAppIconPng fyne.Resource = nil

// SimpleAppIcon 简单的应用图标（内嵌式）
// 如果没有自定义图标，可以创建一个简单的图标
func getAppIcon() fyne.Resource {
	// 这里可以嵌入base64编码的图标
	// 目前返回nil，使用系统默认图标
	return nil
}