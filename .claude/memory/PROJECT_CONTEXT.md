# 項目上下文 - AI啟動器GUI系統

## 🎯 項目願景
- **項目名稱**：AI啟動器 (AI Launcher)
- **項目類型**：跨平台GUI桌面應用程序
- **核心目標**：提供統一的AI工具啟動和管理界面，支持多種AI CLI工具集成
- **當前階段**：GUI系統完成，Windows調試中
- **版本**：v2.0 (Desktop GUI版本)

## 🏗️ 技術架構

### 核心技術
- **主要語言**：Go (Golang)
- **GUI框架**：Fyne v2.4.5 (跨平台原生GUI)
- **系統類型**：桌面應用程序 + 終端管理器
- **架構模式**：主窗口 → 標籤頁管理 → 終端整合 → AI工具執行

### 關鍵組件
- **主窗口系統**：基於Fyne的現代GUI界面
- **終端標籤頁**：多標籤頁終端管理，支援AI工具切換
- **項目管理**：項目配置、切換和環境管理
- **AI工具整合**：Claude Code, Gemini CLI, Codex, Aider等
- **Ollama集成**：本地模型配置和優化
- **設置管理**：主題、語言、AI工具配置

## 📁 項目結構

```
ai-launcher/
├── 📋 README.md                    # 項目概覽
├── 📋 ARCHITECTURE.md              # 系統架構和GUI設計規範
├── 📋 WINDOWS_BUILD.md             # Windows編譯指南
├── 📋 DEBUG_EXE.md                 # Windows調試指南
├── 🔧 go.mod                       # Go模組定義
├── 🔧 Dockerfile                   # Linux編譯環境
├── 🔧 Dockerfile.windows           # Windows交叉編譯
├── 📁 cmd/                         # 應用程序入口
│   └── gui/main.go                 # GUI主程序
├── 📁 internal/                    # 內部套件
│   ├── gui/                        # GUI模組
│   │   ├── main_window.go          # 主窗口實現
│   │   ├── terminal_tabs.go        # 終端標籤頁管理
│   │   ├── project_dialog.go       # 項目配置對話框
│   │   ├── settings_dialog.go      # 設置對話框
│   │   ├── new_terminal_dialog.go  # 新終端對話框
│   │   └── status_bar.go           # 狀態欄組件
│   ├── project/                    # 項目管理
│   │   └── manager.go              # 項目配置管理器
│   └── terminal/                   # 終端管理
│       └── manager.go              # 終端管理器
├── 📁 scripts/                     # 構建和部署腳本
│   └── compile-debug.ps1           # PowerShell編譯腳本
└── 📦 Binaries/                    # 編譯輸出
    ├── ai-launcher.exe             # Windows可執行文件 (28.9MB)
    └── ai-launcher-linux           # Linux可執行文件
```

## 📄 重要文件

- **入口文件**：cmd/gui/main.go (GUI應用程序主入口)
- **核心文檔**：
  - ARCHITECTURE.md - 完整的GUI設計規範和系統架構
  - WINDOWS_BUILD.md - Windows編譯完整指南
  - DEBUG_EXE.md - Windows故障排除和調試
  - BUILD_COMMANDS.md - Docker編譯命令手冊
  - QUICK_FIX.md - 快速修復方案
- **配置文件**：
  - go.mod - Go模組依賴管理
  - Dockerfile / Dockerfile.windows - 跨平台編譯環境
- **構建腳本**：
  - compile-debug.ps1 - PowerShell編譯工具
  - BUILD_COMMANDS.md中的Docker命令集

## 🚀 開發環境

### 編譯需求
- **Go語言**：1.23+
- **CGO**：必需 (Fyne GUI需要)
- **Docker**：推薦用於跨平台編譯

### 快速編譯
```bash
# Linux版本 (Docker)
docker build -t ai-launcher:linux .

# Windows版本 (PowerShell)
docker run --rm -v ${PWD}:/workspace -w /workspace golang:1.23-bullseye bash -c "apt-get update -qq && apt-get install -y gcc-mingw-w64 pkg-config && go mod download && CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o ai-launcher.exe ./cmd/gui"
```

### 支持的AI工具
- **Claude Code** - Claude官方CLI工具
- **Gemini CLI** - Google Gemini命令行界面
- **GitHub Codex** - GitHub Copilot CLI
- **Aider** - AI輔助編程工具
- **Ollama** - 本地LLM模型運行

## 📊 當前狀態

- **Git 分支**：master
- **最新版本**：AI Launcher GUI v2.0 (2025-09-20)
- **系統狀態**：
  - Linux編譯：✅ 完成並驗證
  - Windows編譯：✅ 完成
  - Windows運行：⚠️ 調試中 (啟動問題)
- **最近提交**：
  - 99edfaa feat(gui): redesign GUI layout with improved project management
  - d40f7e3 fix: resolve dependency issues with Web GUI fallback solution
  - df8f3b7 refactor: simplify to GUI-only launcher

### 🚀 v2.0 GUI系統突破 (2025-09-20)
- **完整GUI實現**：基於Fyne v2.4.5的現代桌面界面
- **代碼規模**：45個文件變更，2847行新增代碼
- **跨平台編譯**：Docker工具鏈，支援Linux/Windows
- **終端管理**：多標籤頁終端系統，AI工具整合
- **項目管理**：完整的項目配置和切換系統
- **調試工具**：完善的Windows調試和故障排除體系

### 📈 技術指標
- **GUI框架**：Fyne v2.4.5 (原生跨平台)
- **二進制大小**：~29MB (包含所有依賴)
- **啟動時間**：< 3秒 (預期)
- **內存使用**：~50MB (GUI + 終端管理)
- **支援平台**：Linux, Windows, macOS (理論)

## 🎯 功能特性

### v2.0 核心功能
1. **多標籤頁終端**：同時管理多個AI工具session
2. **項目管理系統**：項目配置、切換、環境管理
3. **AI工具整合**：統一界面啟動和管理AI CLI工具
4. **Ollama配置**：本地模型管理和優化設置
5. **主題和設置**：暗色/亮色主題，多語言支援

### GUI系統優勢
- **統一體驗**：一個界面管理所有AI工具
- **可視化操作**：圖形化項目和終端管理
- **跨平台**：Windows/Linux/macOS原生支援
- **高效切換**：快速在不同AI工具間切換
- **配置管理**：可視化的設置和配置界面

## 🤖 使用場景

### 典型工作流
1. **啟動應用**：雙擊 ai-launcher.exe
2. **選擇項目**：從項目列表選擇或添加新項目
3. **啟動AI工具**：點擊工具欄選擇AI工具 (Claude/Gemini/Codex/Aider)
4. **多任務處理**：在不同標籤頁間切換，同時使用多個AI工具
5. **配置優化**：通過設置對話框調整Ollama和AI工具配置

### 適用用戶
- **AI開發者**：需要經常切換不同AI工具的開發者
- **項目管理者**：管理多個項目和AI工具配置的用戶
- **效率追求者**：希望統一界面管理所有AI工具的用戶
- **跨平台用戶**：需要在不同操作系統間保持一致體驗

## 🔧 當前問題與解決方案

### 已解決問題
- ✅ Fyne API兼容性 (v2.4.5)
- ✅ Docker跨平台編譯環境
- ✅ Linux平台編譯和運行
- ✅ 完整的錯誤處理和日志系統

### 待解決問題
- ⚠️ Windows exe啟動立即關閉
- ⚠️ 可能的OpenGL/顯卡驅動兼容性
- ⚠️ Docker Windows路徑映射限制

### 解決策略
1. **錯誤診斷**：通過命令行運行查看具體錯誤
2. **依賴檢查**：驗證Visual C++ Redistributable等運行時
3. **兼容性測試**：測試不同Windows版本和配置
4. **逐步調試**：使用調試版本定位具體問題

## 💡 下一步計劃

### 短期目標 (1週內)
1. **解決Windows啟動問題**：修復exe無法啟動的問題
2. **功能驗證**：確保所有GUI功能正常工作
3. **用戶體驗優化**：改進界面響應和操作流暢度
4. **部署流程完善**：建立完整的安裝和部署指南

### 中期目標 (1個月內)
1. **性能優化**：提升應用啟動速度和運行效率
2. **功能擴展**：添加更多AI工具支援和高級功能
3. **文檔完善**：創建完整的用戶手冊和開發文檔
4. **社區推廣**：開源發布和用戶反饋收集

### 長期願景
- 成為AI開發者的標準工具啟動器
- 建立AI工具生態系統的統一界面標準
- 提供企業級的AI開發工具管理解決方案