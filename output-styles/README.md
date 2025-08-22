# Output Styles System

## 🌍 国际化支持 (Internationalization Support)

所有 Output Style 文件都支持多语言描述：

```yaml
---
description: 
  en: English description for international users
  zh: 中文描述为中文用户
---
```

## 📚 可用风格 (Available Styles)

### 🏛️ 架构设计类 (Architecture & Design)

1. **senior-architect** - 资深架构师
   - 系统设计和架构讨论
   - 全面分析和战略思考
   - 最佳实践和风险评估

2. **system-architect** - 系统架构师
   - PRD转换为技术架构
   - 多平台解决方案设计
   - 任务分解和团队协作

### 💻 开发实施类 (Development & Implementation)

3. **concise-developer** - 简洁开发者
   - 快速直接的编码协助
   - 最少解释，最多代码
   - 专注可执行解决方案

4. **educational-mentor** - 教育导师
   - 详细解释和教学
   - 循序渐进的学习路径
   - 示例丰富，注重理解

### 🔧 运维与安全类 (Operations & Security)

5. **devops-engineer** - DevOps工程师
   - 基础设施和自动化
   - CI/CD和部署管理
   - 监控和运维卓越

6. **security-analyst** - 安全分析师
   - 威胁建模和漏洞评估
   - 安全开发实践
   - 合规性和风险管理

### 📈 产品与SDK类 (Product & SDK)

7. **product-expert** - 产品需求专家
   - 高质量PRD文档创建
   - 用户故事和需求分析
   - 产品路线图规划

8. **sdk-design-expert** - SDK设计专家
   - 开发者友好API设计
   - 跨平台SDK架构
   - 性能和扩展性优化

9. **sdk-prd-expert** - SDK产品需求专家
   - SDK/Library PRD文档
   - 开发者体验设计
   - 技术文档规划

## 🚀 使用方法 (Usage)

### 查看可用风格
```bash
/output-style
```

### 切换风格
```bash
# 切换到架构师风格
/output-style:set senior-architect

# 切换到简洁开发风格
/output-style:set concise-developer

# 切换到安全分析风格
/output-style:set security-analyst
```

### 查看当前风格
```bash
/output-style:current
```

## 🎯 适用场景 (Use Cases)

### 系统设计阶段
```bash
# 使用架构师风格进行系统设计
/output-style:set senior-architect
/plan "设计微服务架构"
```

### 快速编码
```bash
# 使用简洁开发风格快速实现
/output-style:set concise-developer
# 直接开始编码，最少解释
```

### 学习新技术
```bash
# 使用教育导师风格学习
/output-style:set educational-mentor
# 获得详细解释和示例
```

### 安全审查
```bash
# 使用安全分析师风格
/output-style:set security-analyst
/review
```

## 📁 文件结构 (File Structure)

每个 Output Style 文件都遵循统一的结构：

```markdown
---
description: 
  en: English description
  zh: 中文描述
---

# Style Name

## Role Definition
[Define the role and expertise]

## Communication Style
[Define how to communicate]

## Response Structure
[Define response format]

## Code Generation Preferences
[Define coding standards]

## Working Principles
[Core principles to follow]
```

## ✅ 质量标准 (Quality Standards)

所有 Output Styles 必须：

1. **结构完整**：包含所有必要章节
2. **描述清晰**：明确定义角色和职责
3. **国际化**：支持中英文描述
4. **实用性强**：提供具体的工作模式
5. **专业性高**：符合行业最佳实践

## 🔄 更新记录 (Update History)

- **2025-08-22**: 添加国际化支持，所有文件支持中英文描述
- **2025-08-22**: 新增 4 个专业风格：product-expert, system-architect, sdk-design-expert, sdk-prd-expert
- **2025-08-21**: 初始版本，包含 5 个核心风格

## 💡 自定义风格 (Custom Styles)

您可以创建自己的 Output Style：

1. 在 `~/.claude/output-styles/` 目录创建新文件
2. 使用上述标准结构
3. 添加中英文描述
4. 保存为 `.md` 文件

示例：
```markdown
---
description: 
  en: My Custom Style - Specialized for specific domain
  zh: 我的自定义风格 - 专门用于特定领域
---

# My Custom Style

[Your style definition here]
```

## 安装方法

### 自动安装（推荐）

使用部署脚本自动安装：

```bash
# Windows
cd claude\commands\deploy-package
.\deploy.ps1

# macOS/Linux
cd claude/commands/deploy-package
./deploy.sh
```

### 手动安装

#### Windows
```powershell
# 创建目录
mkdir "%USERPROFILE%\.claude\output-styles"

# 复制文件
xcopy /Y "claude\output-styles\*.md" "%USERPROFILE%\.claude\output-styles\"
```

#### macOS/Linux
```bash
# 创建目录
mkdir -p ~/.claude/output-styles

# 复制文件
cp claude/output-styles/*.md ~/.claude/output-styles/

# 设置权限
chmod 644 ~/.claude/output-styles/*.md
```

## 配置方法

### 项目级别配置
编辑 `.claude/settings.local.json`：
```json
{
  "outputStyle": "senior-architect"
}
```

### 全局配置
编辑 `~/.claude/settings.json`：
```json
{
  "defaultOutputStyle": "concise-developer"
}
```

## 团队协作

将自定义风格放入项目的 `.claude/output-styles/` 目录：
```
your-project/
├── .claude/
│   ├── output-styles/
│   │   ├── team-style.md      # 团队自定义风格
│   │   └── project-style.md   # 项目特定风格
│   └── settings.local.json
```

## 最佳实践

### 开发阶段与风格匹配

| 开发阶段 | 推荐风格 | 原因 |
|---------|---------|------|
| 需求分析 | product-expert | 专业PRD文档 |
| 架构设计 | senior-architect | 全面架构思考 |
| 快速原型 | concise-developer | 高效实现 |
| 代码审查 | security-analyst | 安全性检查 |
| 文档编写 | educational-mentor | 清晰解释 |
| 部署配置 | devops-engineer | 自动化最佳实践 |
| SDK开发 | sdk-design-expert | API设计专业性 |

## 常见问题

### Q: Output Style 没有生效？
A: 检查以下几点：
1. 文件是否有正确的 YAML frontmatter
2. `description` 字段格式是否正确
3. 文件是否在正确的目录
4. 尝试重新设置：`/output-style:set <name>`

### Q: 可以同时使用多个风格吗？
A: 不可以。但你可以：
- 快速切换风格
- 创建组合风格（融合多个风格特点）
- 在不同项目使用不同风格

### Q: 如何恢复默认风格？
A: 使用命令：
```bash
/output-style:set default
```

---

*本系统持续更新中，欢迎贡献更多专业风格配置！*