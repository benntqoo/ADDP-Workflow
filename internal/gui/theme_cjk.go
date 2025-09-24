package gui

import (
    "image/color"
    "io/ioutil"
    "log"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/theme"
)

type cjkTheme struct {
    base    fyne.Theme
    regular []byte
}

func (t *cjkTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
    return t.base.Color(n, v)
}

func (t *cjkTheme) Font(style fyne.TextStyle) fyne.Resource {
    // 使用同一字体覆盖常规/粗体/斜体，保证 CJK 显示。
    return fyne.NewStaticResource("cjk.ttf", t.regular)
}

func (t *cjkTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
    return t.base.Icon(n)
}

func (t *cjkTheme) Size(n fyne.ThemeSizeName) float32 {
    return t.base.Size(n)
}

// ApplyCJKTheme 尝试用给定字体路径应用自定义主题，以确保 CJK 字体覆盖。
func ApplyCJKTheme(app fyne.App, fontPath string) bool {
    data, err := ioutil.ReadFile(fontPath)
    if err != nil || len(data) == 0 {
        return false
    }
    base := app.Settings().Theme()
    if base == nil {
        base = theme.DarkTheme()
    }
    app.Settings().SetTheme(&cjkTheme{base: base, regular: data})
    log.Printf("CJK 主题应用成功: %s", fontPath)
    return true
}
