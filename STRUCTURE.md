# Claude Code 文档结构

```
claude/
├── README.md                           # 主要说明文档（命令体系总览）
├── QUICK_START.md                      # 快速开始指南（5分钟上手）
├── STRUCTURE.md                        # 本文档
│
├── constitution/                       # 宪法体系
│   └── CLAUDE_CONSTITUTION.md         # Claude 协作宪法（核心文档）
│
├── commands/                          # 命令系统文档
│   ├── GLOBAL_COMMANDS.md             # 全局命令示例集
│   ├── COMMANDS_SUMMARY.md            # 命令总览（需更新）
│   ├── CONSTITUTION_INJECT_COMMANDS.md # 宪法注入命令（已过时）
│   └── SYNC_COMMANDS.md               # 同步命令（已过时）
│
├── guides/                            # 使用指南
│   ├── CONSTITUTION_USAGE_GUIDE.md    # 宪法使用详细指南
│   ├── COMMAND_WRITING_GUIDE.md       # 命令编写指南
│   ├── DOCUMENT_STRUCTURE_STANDARD.md  # 文档目录结构规范（新）
│   ├── AI_ASSISTANT_COMPARISON.md     # Claude vs Gemini 对比
│   ├── LEGACY_PROJECT_ONBOARDING.md   # 遗留项目接入指南
│   ├── NEW_VS_LEGACY_PROJECT.md       # 新旧项目使用策略
│   ├── CONSTITUTION_SYNC_GUIDE.md     # 宪法同步指南（已过时）
│   └── MARKET_ANALYSIS.md             # 市场方案对比分析
│
├── templates/                         # 模板文件（新）
│   └── CLAUDE_MD_TEMPLATE.md          # 标准 CLAUDE.md 模板
│
└── examples/                          # 示例文档
    └── META_WORKFLOW_EXAMPLE.md       # 元工作流完整对话示例
```

## 📖 阅读顺序建议

### 初学者路径
1. `QUICK_START.md` - 5分钟快速上手（使用 /meta 或 /onboard）
2. `README.md` - 了解完整命令体系
3. `constitution/CLAUDE_CONSTITUTION.md` - 理解核心理念（H部分：命令系统）
4. `guides/DOCUMENT_STRUCTURE_STANDARD.md` - 了解文档管理规范

### 进阶用户路径
1. `templates/CLAUDE_MD_TEMPLATE.md` - 查看标准模板结构
2. `guides/COMMAND_WRITING_GUIDE.md` - 创建自定义命令
3. `examples/META_WORKFLOW_EXAMPLE.md` - 查看实际案例
4. `guides/MARKET_ANALYSIS.md` - 了解市场现状

### 遗留项目用户
1. 直接使用 `/onboard` 命令
2. `guides/LEGACY_PROJECT_ONBOARDING.md` - 深入理解接入策略
3. `guides/NEW_VS_LEGACY_PROJECT.md` - 理解新旧项目差异

## 🎯 文档用途说明

| 文档类别 | 主要用途 | 目标读者 | 当前状态 |
|----------|----------|----------|----------|
| **README.md** | 命令体系总览和工作流 | 所有用户 | ✅ 最新 |
| **QUICK_START.md** | 快速上手指南 | 新用户 | ✅ 最新 |
| **宪法文档** | 定义协作模式和原则 | 所有用户 | ✅ 最新 |
| **命令文档** | 命令参考和示例 | 日常使用者 | ⚠️ 部分过时 |
| **指南文档** | 深入理解和最佳实践 | 想深入学习者 | ⚠️ 部分过时 |
| **模板文档** | 标准化配置模板 | 项目管理者 | ✅ 最新 |
| **示例文档** | 实际案例参考 | 实践者 | ⚠️ 需更新 |

## 🔄 文档维护状态

### ✅ 已更新到最新
- `README.md` - 包含完整命令列表和工作流
- `QUICK_START.md` - 使用 /meta 和 /onboard 的新流程
- `constitution/CLAUDE_CONSTITUTION.md` - 包含 H 部分命令系统
- `guides/DOCUMENT_STRUCTURE_STANDARD.md` - 新增的文档规范
- `templates/CLAUDE_MD_TEMPLATE.md` - 包含文档管理的标准模板

### ⚠️ 需要更新或废弃
- `commands/COMMANDS_SUMMARY.md` - 未包含新增的企业级命令
- `commands/CONSTITUTION_INJECT_COMMANDS.md` - 已过时的概念
- `commands/SYNC_COMMANDS.md` - 已过时的概念
- `guides/CONSTITUTION_SYNC_GUIDE.md` - 已过时的流程
- `examples/META_WORKFLOW_EXAMPLE.md` - 需要更新到新的命令体系

## 🔄 文档维护

- **更新频率**: 根据用户反馈每月更新
- **版本控制**: 所有文档纳入 Git 管理
- **贡献方式**: 欢迎提交 PR 改进文档

---

*这套文档体系旨在帮助您快速掌握 Claude Code 的协作规范，从理念到实践，从新手到专家。*