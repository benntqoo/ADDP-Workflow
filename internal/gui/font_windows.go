//go:build windows
// +build windows

package gui

import (
    "log"
    "os"
)

// EnsureCJKFont tries to set a system font that supports CJK so that
// Fyne can render Chinese/Japanese/Korean characters correctly on Windows.
// It sets the FYNE_FONT environment variable if not already set.
func EnsureCJKFont() {
    if os.Getenv("FYNE_FONT") != "" {
        return
    }

    // 优先使用 TTF（Fyne 对 TTC 支持有限），再回退 TTC/OTF。
    candidates := []string{
        `C:\\Windows\\Fonts\\simhei.ttf`,                // SimHei 黑体（常见且为 ttf）
        `C:\\Windows\\Fonts\\msyh.ttf`,                 // Microsoft YaHei
        `C:\\Windows\\Fonts\\msyhbd.ttf`,              // Microsoft YaHei Bold
        `C:\\Windows\\Fonts\\Microsoft YaHei UI.ttf`,   // YaHei UI
        `C:\\Windows\\Fonts\\Deng.ttf`,                 // 等线
        `C:\\Windows\\Fonts\\Dengb.ttf`,                // 等线 Bold
        `C:\\Windows\\Fonts\\simsun.ttc`,               // 宋体（TTC）
        `C:\\Windows\\Fonts\\msyh.ttc`,                 //雅黑（TTC）
        `C:\\Windows\\Fonts\\SourceHanSansCN-Regular.otf`, // 思源黑体
    }

    for _, p := range candidates {
        if _, err := os.Stat(p); err == nil {
            _ = os.Setenv("FYNE_FONT", p)
            log.Printf("已设置 FYNE_FONT = %s", p)
            return
        }
    }

    // 若未找到，保持默认（可能会显示方框）。用户可手动设置 FYNE_FONT。
}

// SelectCJKFont 返回选中的 CJK 字体路径（若找到），同时会设置 FYNE_FONT。
func SelectCJKFont() string {
    if v := os.Getenv("FYNE_FONT"); v != "" {
        return v
    }
    // 复用 EnsureCJKFont 的逻辑
    EnsureCJKFont()
    return os.Getenv("FYNE_FONT")
}
