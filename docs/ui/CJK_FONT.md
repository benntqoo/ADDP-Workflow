# Fyne 中文字体显示指南（CJK 显示）

Fyne 在 Windows 上默认字体可能不包含中文（CJK），会出现中文显示为方框/问号。建议按以下方式确保中文正常显示。

## 方式一：环境变量（推荐、最简单）

- 临时运行（PowerShell）：

```
$env:FYNE_FONT="C:\\Windows\\Fonts\\simhei.ttf"
./ai-launcher.exe
```

- 常用字体候选（优先 TTF）：
  - `C:\Windows\Fonts\simhei.ttf`（黑体）
  - `C:\Windows\Fonts\msyh.ttf`（微软雅黑）
  - `C:\Windows\Fonts\Microsoft YaHei UI.ttf`

说明：TTC 在部分 Fyne 版本兼容性不如 TTF，优先选择 TTF。

## 方式二：程序内设置（已内置）

本项目在 Windows 下启动时会自动尝试：

- 通过 `EnsureCJKFont()` 搜索本机常见中文字体并设置环境变量 `FYNE_FONT`；
- 通过 `ApplyCJKTheme(app, fontPath)` 以自定义主题方式覆盖字体；

因此多数机器无需手动设置也能显示中文。如果仍有方框/问号，请按“方式一”强制指定字体。

## 方式三：随包分发字体（可选）

将开源中文字体（如 Noto Sans CJK SC）随程序一起分发，并在启动时优先加载该字体。优点是跨机器可复现，但会增加程序体积。

## 常见问题

- 仍显示乱码（如“鍒/閻/闂”）：这是源码中文文本被错误编码写入的“假中文”，需要在源码中替换为正常简体中文。本项目已逐步清理剩余文案，如发现遗漏欢迎反馈具体界面。
- 仍显示方框：确认 `FYNE_FONT` 指向存在的 TTF；如为 `TTC` 尝试改用同款 `TTF`；或使用 `msyh.ttf/simhei.ttf`。

