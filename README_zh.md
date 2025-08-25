# Claude Code 协作框架

*[English](README.md) | 中文*

通过命令系统、智能代理和个性化风格，让 AI 成为你的智能开发伙伴。

## 🎯 项目简介

Claude Code 协作框架是一套完整的 AI 辅助开发系统，提供三大核心功能：

1. **命令系统** - 项目管理和工作流控制
2. **智能代理** - 专业技术支持和质量保证  
3. **输出风格** - 个性化沟通方式定制

## 🚀 快速开始：Style + Command 组合使用

### 核心概念
- **Style定义角色性格**：AI如何思考和输出（架构师/开发者/分析师）
- **Command定义具体动作**：执行什么任务（/plan、/sync、/learn）
- **组合产生协同效应**：不同组合适用不同场景

### 推荐组合

| 场景 | Style + Command | 效果 |
|------|----------------|------|
| **启动新项目** | `architect` + `/start` → `/plan` | 理解项目并设计架构 |
| **日常开发** | `concise-developer` + `/sync` → `/plan` | 恢复状态并规划任务 |
| **功能实现** | `concise-developer` + `/context` → 编码 | 确认理解后实现 |
| **安全审计** | `security-analyst` + `/context` → 分析 | 理解系统后审查 |
| **学习新技术** | `educational-mentor` + `/start` → `/doc` | 学习概念并记录知识 |
| **紧急修复** | `concise-developer` + `/sync` → 修复 → 部署 | 快速恢复、修复、部署 |

### 工作流示例
```bash
# 早晨例行
/output-style:set concise-developer
/sync                          # 恢复昨天的进度
/plan "完成用户模块"            # 规划今天的任务

# 开发过程中
"实现用户CRUD"                  # 代码实现
/learn "使用Repository模式"     # 记录重要决策

# 提交前
/check                         # 质量检查
/update-spec                   # 更新规范
```

## 📚 功能清单

### 🎮 命令系统

#### 项目理解与管理（3个）
| 命令 | 功能 | 使用时机 | 参数 |
|------|------|----------|------|
| `/start` | 项目快速启动与理解 | 初次接触项目 | 无 |
| `/context` | 上下文同步检查点 | 确保理解一致 | 无 |
| `/sync` | 状态同步器 | 新会话开始 | 无 |

#### 开发辅助（4个）
| 命令 | 功能 | 使用时机 | 参数 |
|------|------|----------|------|
| `/plan` | 任务规划与设计 | 开始新功能前 | [任务描述] |
| `/check` | 完整质量检查 | 提交代码前 | 无 |
| `/watch` | 监察模式 | 编码过程中 | [on\|off\|status\|report] |
| `/test` | 测试生成与执行 | 确保代码质量 | [文件\|功能] |

#### 知识管理（2个）
| 命令 | 功能 | 使用时机 | 参数 |
|------|------|----------|------|
| `/learn` | 学习并记录决策 | 重要决定后 | [决策内容] |
| `/doc` | 智能文档维护 | 更新项目文档 | [api\|readme\|changelog\|arch] |

#### 工作流优化（3个）
| 命令 | 功能 | 使用时机 | 参数 |
|------|------|----------|------|
| `/review` | PR 准备助手 | 创建 PR 前 | 无 |
| `/debug` | 智能调试助手 | 遇到问题时 | [错误信息] |
| `/meta` | 项目规范定制 | 新项目或重大变更 | 无 |

#### 质量保证（2个）
| 命令 | 功能 | 使用时机 | 参数 |
|------|------|----------|------|
| `/analyze` | 深度分析与验证 | 基于经验直觉的风险分析 | [功能/模块] [疑虑或"deep"] |
| `/update-spec` | CLAUDE.md 更新专用 | 固化决策为规范 | [review\|section "content"] |

#### SDK 开发专用命令（5个）
| 命令 | 功能 | 使用时机 | 参数 |
|------|------|----------|------|
| `/sdk-design` | API 设计助手 | 设计新 API 时 | [功能描述] |
| `/sdk-example` | 示例代码生成 | 创建使用示例 | [basic\|advanced\|integration\|all] |
| `/sdk-test` | SDK 测试套件 | 生成专业测试 | [unit\|integration\|compat\|performance\|all] |
| `/sdk-doc` | SDK 文档生成 | 编写文档时 | [api\|guide\|migration\|all] |
| `/sdk-release` | 发布准备助手 | 准备新版本 | [major\|minor\|patch\|check] |

📖 详细文档：[commands/docs/](commands/docs/)

### 🤖 智能代理

**质量保证代理**
- `code-reviewer` - 代码审查专家
- `test-automator` - 自动化测试生成
- `performance-optimizer` - 性能优化分析
- `bug-hunter` - Bug查找和修复
- `security-analyst` - 安全漏洞分析

**技术专家代理**
- `kotlin-expert` - Kotlin全栈开发
- `python-ml-specialist` - Python机器学习
- `golang-systems-engineer` - Go系统编程
- `rust-zero-cost` - Rust零成本抽象
- `typescript-fullstack-expert` - TypeScript全栈

**更多代理（35+）**涵盖：Android、iOS、Web、后端、数据库、DevOps等

📖 详细文档：[agents/docs/](agents/docs/)

### 🎨 输出风格（9个专业人格）

#### 架构与设计
| 风格名称 | 适用场景 | 特点 |
|---------|---------|------|
| `senior-architect` | 系统设计 | 全面分析、风险评估、战略思考 |
| `system-architect` | 技术架构 | PRD转换为技术设计、多平台解决方案 |

#### 开发与实施
| 风格名称 | 适用场景 | 特点 |
|---------|---------|------|
| `concise-developer` | 快速编码 | 最少解释、直接方案、代码优先 |
| `educational-mentor` | 学习教学 | 详细解释、循序渐进、示例丰富 |

#### 运维与安全
| 风格名称 | 适用场景 | 特点 |
|---------|---------|------|
| `devops-engineer` | 基础设施 | 自动化优先、可靠性、IaC思维 |
| `security-analyst` | 安全审查 | 威胁建模、漏洞评估、合规检查 |

#### 产品与SDK
| 风格名称 | 适用场景 | 特点 |
|---------|---------|------|
| `product-expert` | 产品需求 | PRD文档、用户故事、路线规划 |
| `sdk-design-expert` | SDK架构 | API设计、跨平台、开发者体验 |
| `sdk-prd-expert` | SDK产品管理 | 开发者工具PRD、API产品策略 |

**共计：9个专业输出风格**
📖 详细文档：[output-styles/README.md](output-styles/README.md)

## 🚀 手动部署

### 1. 命令系统安装

**Windows:**
```powershell
# 创建命令目录
mkdir "%USERPROFILE%\.claude\commands"

# 复制核心命令
xcopy /Y "claude\commands\deploy-package\global\*.md" "%USERPROFILE%\.claude\commands\"

# 复制SDK命令（可选）
xcopy /Y "claude\commands\deploy-package\sdk\*.md" "%USERPROFILE%\.claude\commands\"
```

**macOS/Linux:**
```bash
# 创建命令目录
mkdir -p ~/.claude/commands

# 复制核心命令
cp claude/commands/deploy-package/global/*.md ~/.claude/commands/

# 复制SDK命令（可选）
cp claude/commands/deploy-package/sdk/*.md ~/.claude/commands/
```

### 2. 智能代理安装

**Windows:**
```powershell
# 创建代理目录
mkdir "%USERPROFILE%\.claude\agents"

# 复制所有代理
xcopy /E /Y "claude\agents\*.md" "%USERPROFILE%\.claude\agents\"
```

**macOS/Linux:**
```bash
# 创建代理目录
mkdir -p ~/.claude/agents

# 复制所有代理
cp -r claude/agents/*.md ~/.claude/agents/
```

### 3. 输出风格安装

**Windows:**
```powershell
# 创建风格目录
mkdir "%USERPROFILE%\.claude\output-styles"

# 复制所有风格
xcopy /Y "claude\output-styles\*.md" "%USERPROFILE%\.claude\output-styles\"
```

**macOS/Linux:**
```bash
# 创建风格目录
mkdir -p ~/.claude/output-styles

# 复制所有风格
cp claude/output-styles/*.md ~/.claude/output-styles/
```

### 4. 验证安装

```bash
# 检查安装的文件
ls ~/.claude/commands/       # 应看到命令文件
ls ~/.claude/agents/         # 应看到代理文件
ls ~/.claude/output-styles/  # 应看到风格文件
```

## 📖 使用方法

### 基础使用

#### 1. 开始新项目
```bash
# 使用meta命令创建项目规范
/meta

# Claude会：
# - 分析项目特征
# - 询问关键信息
# - 生成CLAUDE.md
# - 设置文档结构
```

#### 2. 恢复工作状态
```bash
# 新会话开始时
/sync

# 系统会：
# - 读取上次工作状态
# - 检查未提交更改
# - 提供工作建议
```

#### 3. 切换输出风格
```bash
# 查看可用风格
/output-style

# 设置风格
/output-style:set senior-architect

# 开始使用新风格工作
/plan "设计微服务架构"
```

#### 4. 使用智能代理
```bash
# 代理会自动激活，也可手动指定
"使用 python-ml-specialist 代理帮我设计模型"

# 或在特定任务时自动触发
"帮我审查这段代码的安全性"  # 自动激活 security-analyst
```

### 项目配置

在项目根目录创建 `.claude/` 目录：

```
your-project/
├── .claude/
│   ├── PROJECT_CONTEXT.md  # 项目上下文
│   ├── DECISIONS.md        # 决策记录
│   ├── settings.local.json # 项目设置
│   └── state/              # 状态文件
├── CLAUDE.md               # 项目规范
└── ... 项目文件
```

配置示例 (`.claude/settings.local.json`):
```json
{
  "outputStyle": "concise-developer",
  "permissions": {
    "defaultMode": "acceptEdits"
  }
}
```

## 🎯 工作流程示例

### 场景1：开始新功能开发

```bash
# 1. 恢复状态
/sync
# → 恢复上次工作进度

# 2. 规划任务
/plan "添加用户认证功能"
# → 生成任务计划和技术方案

# 3. 开发实现
# 自动激活相关代理提供支持
# - code-reviewer 持续审查
# - test-automator 生成测试
# - security-analyst 检查安全

# 4. 记录决策
/learn "决定使用JWT而非Session，因为..."
# → 保存到DECISIONS.md

# 5. 更新规范
/update-spec
# → 更新CLAUDE.md
```

### 场景2：代码审查与优化

```bash
# 1. 切换到架构师风格
/output-style:set senior-architect

# 2. 进行架构审查
/context
# → 全面分析当前架构

# 3. 性能优化
"分析并优化数据库查询性能"
# → 自动激活 performance-optimizer

# 4. 安全审查
/output-style:set security-analyst
"审查认证系统的安全性"
# → 深度安全分析
```

### 场景3：SDK开发流程

```bash
# 1. 设计API
/sdk-design "支付SDK接口设计"

# 2. 生成示例
/sdk-example advanced

# 3. 创建测试
/sdk-test all

# 4. 编写文档
/sdk-doc api

# 5. 准备发布
/sdk-release check
```

### 场景4：团队协作

```bash
# 1. 早晨开始
/sync
# → 查看团队昨天的更改

# 2. 理解新代码
/context
# → 同步对项目的理解

# 3. 切换到教学风格（给新人讲解）
/output-style:set educational-mentor
"解释这个认证模块的工作原理"

# 4. 记录团队决策
/learn "团队决定采用微服务架构..."

# 5. 更新团队文档
/doc readme
```

## 🏆 最佳实践

### 1. 工作习惯
- **开始必sync**：每次工作前使用 `/sync` 恢复状态
- **及时记录**：重要决策立即用 `/learn` 记录
- **定期更新**：用 `/update-spec` 固化规范

### 2. 风格选择
- **设计阶段**：使用 `senior-architect`
- **快速开发**：使用 `concise-developer`
- **代码审查**：使用 `security-analyst`
- **文档编写**：使用 `educational-mentor`

### 3. 代理协作
- 让代理自动激活，不要过度干预
- 信任专业代理的建议
- 多个代理可以同时工作

### 4. 团队规范
- 共享 `.claude/` 目录
- 统一使用命令系统
- 定期更新 PROJECT_CONTEXT.md

## 📁 项目结构

```
claude/
├── README.md               # 英文文档
├── README_zh.md           # 本文档（中文）
├── RELEASE_NOTE.md        # 版本历史
├── commands/              # 命令系统
│   ├── docs/             # 命令详细文档
│   └── deploy-package/   # 部署包
│       ├── global/       # 核心命令(8个)
│       └── sdk/          # SDK命令(5个)
├── agents/               # 智能代理
│   ├── docs/            # 代理详细文档
│   └── *.md             # 代理定义文件(35+)
├── output-styles/        # 输出风格
│   ├── README.md        # 风格使用指南
│   └── *.md             # 风格定义文件(9个)
└── guides/              # 深度指南
    └── *.md             # 各种专题指南
```

## 🆘 常见问题

**Q: 命令没有生效？**
A: 检查文件是否复制到正确目录 `~/.claude/commands/`

**Q: 代理没有自动激活？**
A: 代理基于任务描述自动激活，使用明确的关键词

**Q: 输出风格如何持久化？**
A: 在项目的 `.claude/settings.local.json` 中设置

**Q: 如何创建自定义命令/代理/风格？**
A: 参考相应目录下的文档和模板

## 🤝 贡献

这是一个开源项目，为所有开发者提供整理好的 Claude Code 开发经验和工具。

欢迎贡献！请：
1. Fork 本仓库
2. 创建功能分支
3. 提交更改
4. 创建 Pull Request

### 报告问题
- 使用 GitHub Issues
- 提供详细的复现步骤
- 说明预期行为

## 📚 更多资源

- [命令系统详细文档](commands/docs/)
- [代理系统详细文档](agents/docs/)
- [输出风格详细文档](output-styles/README.md)
- [深度使用指南](guides/)
- [Claude Code 官方文档](https://docs.anthropic.com/en/docs/claude-code)

## 📄 许可证

MIT License - 详见 [LICENSE](LICENSE) 文件

## 🌟 Star 历史

如果这个项目对你有帮助，请给个 star ⭐

---

*让 Claude Code 成为你最好的开发伙伴！*

**由 Claude Code 社区用 ❤️ 制作**