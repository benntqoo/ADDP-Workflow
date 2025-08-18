# 項目上下文

## 🎯 項目願景
- **項目名稱**：Claude Code Collaboration Framework
- **項目類型**：AI 協作框架/開發工具
- **核心目標**：結合精確命令控制與智能代理自動化，提升開發效率
- **當前階段**：v4.0 穩定版本（生產環境就緒）
- **版本**：v4.0.1

## 🏗️ 技術架構

### 核心技術
- **主要語言**：Markdown 配置文件
- **系統類型**：Claude Code 擴展系統
- **架構模式**：混合命令 + 智能代理雙軌系統

### 關鍵組件
- **命令系統 v3.3**：8個核心命令 + 5個SDK專用命令
- **代理系統 v4.0**：40+ 專業代理覆蓋全技術棧
- **智能路由**：基於文件類型和代碼內容自動選擇代理
- **配置系統**：triggers.yaml 和 workflows.yaml

## 📁 項目結構

```
claude/
├── agents/                     # 智能代理系統（40+ 專業代理）
│   ├── core/                   # 核心代理（代碼審查、性能優化、測試自動化）
│   ├── languages/              # 語言專家代理（支持15+編程語言）
│   ├── frameworks/             # 框架專家代理
│   ├── quality/                # 質量保證代理
│   └── workflow/               # 工作流代理
├── commands/                   # 命令系統 v3.3
│   └── deploy-package/         # 部署包
│       ├── global/             # 8個核心命令
│       └── sdk/                # 5個SDK專用命令
├── commands-legacy/            # 歷史命令存檔
├── config/                     # 配置文件
├── constitution/               # 協作憲法系統
├── guides/                     # 深度指南文檔
└── templates/                  # 模板文件
```

## 📄 重要文件

- **入口文件**：無（配置型框架）
- **核心文檔**：
  - README.md - 中文主文檔
  - README.en.md - 英文文檔
  - agents/README.md - 代理系統文檔
  - commands/deploy-package/SIMPLE_COMMANDS_SUMMARY.md - 命令系統文檔
- **配置文件**：
  - config/triggers.yaml - 智能觸發配置
  - config/workflows.yaml - 工作流定義
  - config/token-settings.yaml - Token 優化設置
- **部署腳本**：
  - commands/deploy-package/deploy.ps1 (Windows)
  - commands/deploy-package/deploy.sh (macOS/Linux)

## 🚀 開發環境

### 安裝方式

#### 自動安裝（推薦）
```bash
# Windows
cd claude\commands\deploy-package
.\deploy.ps1

# macOS/Linux
cd claude/commands/deploy-package
./deploy.sh
```

#### 手動安裝
```bash
# 複製代理到 Claude 主目錄
cp -r claude/agents ~/.claude/agents
cp -r claude/config ~/.claude/config

# 複製命令（可選）
cp -r claude/commands/deploy-package/global ~/.claude/commands
```

### 核心命令
- `/start` - 快速理解項目
- `/sync` - 恢復工作狀態
- `/plan` - 任務規劃
- `/meta` - 創建項目規範
- `/learn` - 記錄決策
- `/doc` - 維護文檔
- `/context` - 同步上下文
- `/update-spec` - 更新規範

## 📊 當前狀態

- **Git 分支**：master
- **最近提交**：
  - bb1db50 fix: 修復 agent 檔案 YAML frontmatter 格式
  - 638d87a feat: 擴展專業代理系統與優化 token 使用
  - 7145541 docs: 重組文檔結構 - 英文為主，中文為輔
- **項目狀態**：Clean（無未提交更改）

## 🎯 開發重點

### v4.0 核心特性
1. **混合架構**：命令系統（人類主導）+ 代理系統（AI驅動）
2. **智能檢測**：5個上下文檢測器解決多用途語言場景衝突
3. **專業深度**：覆蓋 Android、Kotlin、Java、C#、Python、Go、Rust 等全技術棧
4. **零配置**：自動檢測、智能路由、漸進式遷移

### 效率提升
- 項目理解：3倍提升
- 代碼開發：5倍提升  
- 質量檢查：8倍提升
- 測試生成：10倍提升
- 性能優化：15倍提升

## 🤖 AI 協作建議

### 下一步建議
1. **測試安裝**：運行部署腳本安裝到本地環境
2. **體驗代理**：創建測試項目體驗智能代理自動觸發
3. **了解命令**：熟悉8個核心命令的使用場景
4. **自定義配置**：根據團隊需求調整 triggers.yaml

### 注意事項
- v4.0 完全向後兼容 v3.3，可漸進式遷移
- 代理系統會自動啟用，無需手動配置
- 保留的8個核心命令是基礎設施，不可替代
- 建議先觀察一週，再主動協作

### 推薦工作流
1. 使用 `/start` 理解新項目
2. 使用 `/meta` 建立項目規範
3. 讓代理系統自動提供專業支持
4. 使用 `/learn` 記錄重要決策
5. 使用 `/sync` 在新會話恢復狀態