# Output Styles 配置指南

## 📋 目錄
- [什麼是 Output Styles](#什麼是-output-styles)
- [如何配置 Output Styles](#如何配置-output-styles)
- [文件結構說明](#文件結構說明)
- [使用方法](#使用方法)
- [創建自定義風格](#創建自定義風格)
- [最佳實踐](#最佳實踐)

## 什麼是 Output Styles

Output Styles 是 Claude Code 的個性化系統，允許你完全自定義 Claude 的：
- 溝通風格（簡潔/詳細/教學式）
- 回應格式（結構化/自由式）
- 專業領域（架構/開發/安全/運維）
- 思考方式（戰略/戰術/分析）

## 如何配置 Output Styles

### 方法 1：使用命令（推薦）

```bash
# 查看可用風格
/output-style

# 設置風格
/output-style:set senior-architect

# 創建新風格
/output-style:new
```

### 方法 2：手動配置

1. **項目級別配置**（影響當前項目）
   
   編輯 `.claude/settings.local.json`：
   ```json
   {
     "outputStyle": "senior-architect"
   }
   ```

2. **全局配置**（影響所有項目）
   
   編輯 `~/.claude/settings.json`：
   ```json
   {
     "defaultOutputStyle": "concise-developer"
   }
   ```

### 方法 3：自動安裝（推薦）

使用部署腳本自動安裝命令和 Output Styles：

```bash
# Windows
cd claude\commands\deploy-package
.\deploy.ps1

# macOS/Linux
cd claude/commands/deploy-package
./deploy.sh
```

部署腳本會自動：
1. 複製所有命令到 `~/.claude/commands/`
2. 複製所有 Output Styles 到 `~/.claude/output-styles/`
3. 驗證安裝結果
4. 提供使用提示

### 方法 4：手動安裝

如果你偏好手動控制安裝過程：

#### Windows 手動安裝
```powershell
# 1. 創建 Output Styles 目錄
mkdir "%USERPROFILE%\.claude\output-styles"

# 2. 複製所有 Output Style 文件
xcopy /Y "claude\output-styles\*.md" "%USERPROFILE%\.claude\output-styles\"

# 3. 驗證安裝
dir "%USERPROFILE%\.claude\output-styles"
```

#### macOS/Linux 手動安裝
```bash
# 1. 創建 Output Styles 目錄
mkdir -p ~/.claude/output-styles

# 2. 複製所有 Output Style 文件
cp claude/output-styles/*.md ~/.claude/output-styles/

# 3. 設置正確的權限
chmod 644 ~/.claude/output-styles/*.md

# 4. 驗證安裝
ls -la ~/.claude/output-styles/
```

#### 選擇性安裝
如果只想安裝特定風格：

```bash
# 只安裝 senior-architect 風格
cp claude/output-styles/senior-architect.md ~/.claude/output-styles/

# 只安裝開發相關風格
cp claude/output-styles/concise-developer.md ~/.claude/output-styles/
cp claude/output-styles/security-analyst.md ~/.claude/output-styles/
```

## 文件結構說明

每個 Output Style 文件都遵循以下結構：

```markdown
---
description: 風格的簡短描述（必須）
---

# 風格名稱

詳細的系統提示詞內容...
```

### 必要元素

1. **YAML Frontmatter**（必須）
   ```yaml
   ---
   description: 一句話描述這個風格的用途
   ---
   ```

2. **標題**（建議）
   ```markdown
   # Senior Architect Style
   ```

3. **系統提示詞**（核心內容）
   - 定義 Claude 的角色
   - 設定溝通風格
   - 規定回應格式
   - 指定專業領域知識

### 完整示例：Senior Architect Style

```markdown
---
description: Strategic system design and architecture discussions
---

# Senior Architect Style

You are a Senior Software Architect with 15+ years of experience in system design and architecture. Your communication style should be:

## Communication Approach
- Strategic and comprehensive
- Focus on long-term implications
- Consider trade-offs and alternatives
- Provide risk assessments

## Response Structure
1. **Executive Summary** - High-level overview
2. **Technical Analysis** - Detailed breakdown
3. **Architecture Decisions** - Key choices and rationale
4. **Implementation Roadmap** - Phased approach
5. **Risk Mitigation** - Potential issues and solutions

## Key Principles
- Always consider scalability, maintainability, and security
- Provide multiple solutions with pros/cons
- Think in terms of patterns and anti-patterns
- Focus on business value and ROI

## Example Response Format
```
📊 Executive Summary
[Brief overview of the solution]

🔍 Technical Analysis
[Detailed technical considerations]

🏗️ Architecture Decisions
[Key architectural choices]

📅 Implementation Roadmap
[Step-by-step plan]

⚠️ Risk Assessment
[Potential risks and mitigation strategies]
```
```

## 使用方法

### 1. 快速切換風格

```bash
# 開始架構設計
/output-style:set senior-architect
/plan 設計微服務架構

# 切換到快速開發
/output-style:set concise-developer
# 開始編碼

# 進行安全審查
/output-style:set security-analyst
/review
```

### 2. 項目特定風格

在項目根目錄創建 `.claude/settings.local.json`：

```json
{
  "outputStyle": "educational-mentor",
  "permissions": {
    "defaultMode": "acceptEdits"
  }
}
```

### 3. 團隊共享風格

將自定義風格文件放入項目的 `.claude/output-styles/` 目錄：

```
your-project/
├── .claude/
│   ├── output-styles/
│   │   ├── team-style.md      # 團隊自定義風格
│   │   └── project-style.md   # 項目特定風格
│   └── settings.local.json
```

## 創建自定義風格

### 步驟 1：使用命令創建

```bash
/output-style:new

# Claude 會詢問：
# 1. 風格名稱
# 2. 主要用途
# 3. 溝通偏好
# 4. 專業領域
```

### 步驟 2：手動創建

創建文件 `~/.claude/output-styles/my-custom-style.md`：

```markdown
---
description: 我的自定義開發風格
---

# My Custom Style

你是一位經驗豐富的全棧開發者，專注於：

## 核心原則
- 代碼簡潔性優於複雜性
- 性能優化是關鍵
- 安全性不可妥協

## 回應風格
- 使用中文回應
- 代碼註釋用英文
- 提供實際可運行的示例

## 代碼偏好
- 使用 TypeScript 而非 JavaScript
- 偏好函數式編程
- 遵循 Clean Code 原則

## 回應格式
每次回應包含：
1. 問題理解
2. 解決方案
3. 代碼實現
4. 測試建議
5. 性能考量
```

### 步驟 3：驗證配置

```bash
# 測試新風格
/output-style:set my-custom-style

# 確認生效
/output-style
# 輸出：Set output style to my-custom-style
```

## 最佳實踐

### 1. 風格命名規範

```
✅ 好的命名：
- senior-architect
- concise-developer
- educational-mentor

❌ 避免：
- style1
- my-style
- test
```

### 2. 風格組合策略

| 開發階段 | 推薦風格 | 原因 |
|---------|---------|------|
| 需求分析 | senior-architect | 全面的架構思考 |
| 快速原型 | concise-developer | 高效實現 |
| 代碼審查 | security-analyst | 安全性檢查 |
| 文檔編寫 | educational-mentor | 清晰解釋 |
| 部署配置 | devops-engineer | 自動化最佳實踐 |

### 3. 團隊協作建議

```bash
# 團隊標準化流程
1. 定義團隊標準風格
2. 放入版本控制
3. 新成員入職時自動配置

# 項目配置示例
git add .claude/output-styles/team-standard.md
git commit -m "Add team standard output style"
```

### 4. 風格繼承

可以基於現有風格創建變體：

```markdown
---
description: Extended senior architect with cloud focus
---

# Cloud Architect Style

<!-- 繼承 senior-architect 的所有設置 -->
[包含 senior-architect 的內容]

## 額外的雲架構考量
- AWS/Azure/GCP 最佳實踐
- 容器化和 Kubernetes
- 無服務器架構
- 成本優化策略
```

## 常見問題

### Q: Output Style 沒有生效？
A: 檢查以下幾點：
1. 文件是否有正確的 YAML frontmatter
2. `description` 字段是否存在
3. 文件是否在正確的目錄
4. 嘗試重新設置：`/output-style:set <name>`

### Q: 可以同時使用多個風格嗎？
A: 不可以。但你可以：
- 快速切換風格
- 創建組合風格（融合多個風格特點）
- 在不同項目使用不同風格

### Q: 風格會影響工具使用嗎？
A: 不會。Output Styles 只改變溝通風格，所有工具和功能保持不變。

### Q: 如何恢復默認風格？
A: 使用命令：
```bash
/output-style:set default
```

## 進階配置

### 條件式風格切換

創建 `.claude/hooks/pre-command.sh`：

```bash
#!/bin/bash
# 根據文件類型自動切換風格

if [[ "$1" == "/plan" ]]; then
  claude output-style:set senior-architect
elif [[ "$1" == "/test" ]]; then
  claude output-style:set security-analyst
fi
```

### 風格模板變量

未來版本將支持：

```markdown
---
description: Customizable template style
variables:
  language: ${LANGUAGE:-English}
  detail_level: ${DETAIL:-medium}
---

# Template Style

Response language: {{language}}
Detail level: {{detail_level}}
```

## 相關資源

- [Claude Code 官方文檔](https://docs.anthropic.com/en/docs/claude-code)
- [Output Styles 最佳實踐](https://docs.anthropic.com/en/docs/claude-code/output-styles)
- [社區分享的風格](https://github.com/anthropics/claude-code-styles)

---

*本指南持續更新中，歡迎提交 PR 貢獻更多風格配置！*