# Claude Code 快速开始指南

5 分钟内开始使用 Claude Code 协作规范体系。

## 🚀 第 1 步：选择您的场景

### A. 我有一个新项目
```bash
# 直接使用元工作流命令
/meta

# Claude 会：
# 1. 分析项目特征
# 2. 询问关键信息
# 3. 自动生成包含文档规范的 CLAUDE.md
# 4. 设置标准文档目录结构
```

### B. 我有一个现有项目
```bash
# 使用遗留项目接入命令
/onboard

# Claude 会：
# 1. 扫描项目结构和技术栈
# 2. 分析现有文档分布
# 3. 生成适配性的 CLAUDE.md
# 4. 创建文档迁移计划
```

## 🛠️ 第 2 步：安装全局命令（可选但推荐）

如果 `/meta` 或 `/onboard` 命令不可用，需要先安装全局命令：

### Windows (PowerShell)
```powershell
# 创建命令目录
New-Item -ItemType Directory -Force -Path "$env:USERPROFILE\.claude\commands"

# 下载并安装命令集（假设有安装脚本）
# 或参考 commands/GLOBAL_COMMANDS.md 手动创建
```

### macOS/Linux (Bash)
```bash
# 创建命令目录
mkdir -p ~/.claude/commands

# 下载并安装命令集
# 或参考 commands/GLOBAL_COMMANDS.md 手动创建
```

**提示**：完整的命令列表和安装说明请查看 `commands/GLOBAL_COMMANDS.md`

## 📋 第 3 步：核心命令速查

### 🎯 起步命令
| 命令 | 用途 | 使用场景 |
|------|------|----------|
| `/meta` | 新项目规范定制 | 从零开始的项目 |
| `/onboard` | 遗留项目接入 | 现有项目引入 Claude |

### 🛠️ 日常开发
| 命令 | 用途 | 使用场景 |
|------|------|----------|
| `/discover` | 理解功能模块 | 探索陌生代码 |
| `/test` | 智能生成测试 | 编写单元/集成测试 |
| `/coverage` | 覆盖率分析 | 识别测试盲点 |
| `/perf` | 性能优化 | 找出性能瓶颈 |

### 📚 文档管理
| 命令 | 用途 | 使用场景 |
|------|------|----------|
| `/doc-api` | API 文档生成 | 自动化文档 |
| `/changelog` | 变更日志 | 版本发布 |
| `/doc-sync` | 文档一致性 | 定期检查 |

## 💬 第 4 步：与 Claude 协作

### 深度模式（复杂任务）
```
用户: "实现用户认证系统"
Claude: [深度分析] → [方案设计] → [任务规划] → [系统实施]
```

### 快速模式（简单任务）
```
用户: "修复这个 null 检查的 bug"
Claude: [立即修复] → [简要说明]
```

## 🎯 推荐工作流

### 1. 日常开发流程
```
/discover → /test → /coverage → /perf → /deploy-check
```
理解代码 → 编写测试 → 检查覆盖率 → 性能优化 → 部署准备

### 2. 故障处理流程
```
/audit → /rollback → /config → /sync-team
```
问题诊断 → 快速回滚 → 配置修复 → 知识同步

### 3. 新功能开发示例
```bash
/discover user-service   # 先理解现有模块
"实现购物车功能"          # Claude 创建完整计划
/test                   # 生成测试代码
/coverage              # 检查测试覆盖率
```

## 📚 深入学习路径

1. **理解协作模式**
   → 阅读 `guides/CONSTITUTION_USAGE_GUIDE.md`

2. **掌握命令系统**
   → 阅读 `guides/COMMAND_WRITING_GUIDE.md`

3. **查看实际案例**
   → 阅读 `examples/META_WORKFLOW_EXAMPLE.md`

4. **了解最佳实践**
   → 阅读 `guides/AI_ASSISTANT_COMPARISON.md`

## ❓ 常见问题

- **命令不可用？**
  → 重启 Claude Code 或检查命令是否在 `~/.claude/commands/`

- **Claude 不理解我的项目？**
  → 使用 `/onboard` 让 Claude 先扫描和理解项目

- **文档散落各处？**
  → 使用 `/doc-structure init` 初始化标准文档结构

- **需要团队协作？**
  → 将 CLAUDE.md 提交到版本控制，使用 `/sync-team` 同步知识

## 🎉 恭喜！

您已经掌握了 Claude Code 的基础使用。核心要点：

1. **一键启动**：`/meta`（新项目）或 `/onboard`（现有项目）
2. **自动规范**：CLAUDE.md 自动包含文档管理等最佳实践
3. **完整工具链**：从开发到部署的全流程命令支持
4. **团队协作**：通过 Graphiti MCP 实现知识共享

---

*下一步：查看 README.md 了解完整的命令列表和高级用法。*