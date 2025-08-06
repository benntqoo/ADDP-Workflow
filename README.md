# Claude Code 协作规范体系

本目录包含了完整的 Claude Code 协作规范、命令系统和使用指南。

## 📁 目录结构

```
claude/
├── README.md                    # 本文档
├── constitution/               # 宪法体系
│   └── CLAUDE_CONSTITUTION.md  # Claude 协作宪法
├── commands/                   # 命令系统
│   ├── GLOBAL_COMMANDS.md      # 全局命令集
│   ├── COMMANDS_SUMMARY.md     # 命令总览
│   ├── COMMAND_COORDINATION.md # 命令协调机制
│   └── deploy-package/         # 命令部署包 v2.1.0
├── guides/                     # 使用指南
│   ├── AI_ASSISTANT_COMPARISON.md    # AI 助手对比
│   ├── LEGACY_PROJECT_ONBOARDING.md   # 遗留项目接入
│   ├── NEW_VS_LEGACY_PROJECT.md       # 新旧项目对比
│   └── MARKET_ANALYSIS.md             # 市场分析
└── templates/                  # 模板文件
    └── CLAUDE_MD_TEMPLATE.md   # CLAUDE.md 模板
```

## 🚀 快速开始

### 1. 新项目
```bash
# 复制宪法模板到项目
cp claude/constitution/CLAUDE_CONSTITUTION.md /your-project/CLAUDE.md

# 使用元工作流定制
/meta
```

### 2. 遗留项目
```bash
# 使用专门的接入流程
/onboard
```

### 3. 安装全局命令
参考 `commands/GLOBAL_COMMANDS.md` 安装推荐的命令集。

## 🎮 已配置的自定义命令

### 全局命令（~/.claude/commands/）
这些命令在所有项目中都可以使用：

#### 核心工作流命令
| 命令 | 功能描述 | 使用场景 |
|------|----------|----------|
| `/meta` | 启动元工作流，定制项目规范 | **新项目**初始化或规范定制 |
| `/onboard` | 遗留项目接入向导 | **现有项目**引入 Claude Code |
| `/update-constitution` | 更新 CLAUDE.md 版本 | 保留定制内容，更新标准部分 |
| `/constitution` | 检查并应用 Claude 协作宪法 | 确认协作模式和规范 |
| `/deep` | 启动深度开发工作流 | 复杂任务的系统化实施 |
| `/explore` | 快速浏览项目结构 | 生成项目地图（5-10分钟） |
| `/batch` | 启用批量操作模式 | 需要并行处理多个操作时 |

#### 遗留项目专用命令
| 命令 | 功能描述 | 使用场景 |
|------|----------|----------|
| `/audit` | 遗留项目健康度评估 | 技术债务识别、改进建议 |
| `/discover [模块]` | 业务逻辑深度理解 | 数据流分析（30-60分钟） |
| `/retrofit` | 渐进式改造 | 小步改进，保持兼容性 |

#### 文档管理命令
| 命令 | 功能描述 | 使用场景 |
|------|----------|----------|
| `/doc-api [模块]` | API 文档生成 | 自动生成 OpenAPI/Swagger 文档 |
| `/doc-arch` | 架构文档更新 | 生成架构图和设计文档 |
| `/changelog` | 变更日志管理 | 基于提交历史生成 CHANGELOG |
| `/doc-sync` | 文档一致性检查 | 验证文档与代码的同步性 |
| `/readme [项目]` | README 生成 | 创建或更新项目 README |
| `/doc-structure` | 文档目录管理 | 初始化或检查文档结构规范 |

#### 企业级架构命令
| 命令 | 功能描述 | 使用场景 |
|------|----------|----------|
| `/deps [模块]` | 模块依赖分析 | 分析依赖关系、检测循环依赖、优化建议 |
| `/microservice [操作]` | 微服务协调 | 服务发现、契约管理、部署编排 |
| `/migrate [操作]` | 数据库迁移管理 | 创建迁移、验证脚本、回滚方案 |

#### 团队协作命令
| 命令 | 功能描述 | 使用场景 |
|------|----------|----------|
| `/review` | 代码审查 | **使用 Claude Code 内置命令** |
| `/pr_comments` | PR 评论查看 | **使用 Claude Code 内置命令** |
| `/sync-team` | 团队知识同步 | 同步到 Graphiti 知识图谱，知识传承 |

#### 测试质量命令
| 命令 | 功能描述 | 使用场景 |
|------|----------|----------|
| `/test [代码]` | 智能测试生成 | 自动生成单元测试、集成测试、边界测试 |
| `/coverage` | 测试覆盖率分析 | 分析覆盖率、识别盲点、改进建议 |
| `/perf [模块]` | 性能分析优化 | 识别瓶颈、优化建议、基准测试 |

#### 部署运维命令
| 命令 | 功能描述 | 使用场景 |
|------|----------|----------|
| `/deploy-check` | 部署前检查 | 代码、测试、安全、文档全面检查 |
| `/rollback [类型]` | 回滚方案生成 | 代码回滚、数据库回滚、配置回滚 |
| `/config [环境]` | 环境配置管理 | 配置验证、安全检查、环境同步 |

### 项目命令（.claude/commands/）
这些命令仅在本项目中可用：

| 命令 | 功能描述 | 使用场景 |
|------|----------|----------|
| `/sync [操作]` | 双向智能同步系统 | 上传/下载/冲突解决/版本管理 |
| `/ai-rules [操作]` | 完整规范管理器 | check/apply/validate/report |
| `/guardian [操作]` | 主动监控系统 | 实时质量监控/安全检查/性能分析 |

### 使用方法
1. **命令触发**：直接输入 `/` 后跟命令名，如 `/meta`
2. **重启生效**：新创建的命令需要重启 Claude Code 后才能使用
3. **命令位置**：
   - 全局命令：`C:\Users\[用户名]\.claude\commands\`
   - 项目命令：`[项目根目录]\.claude\commands\`

## 🆕 命令系统 v2.1 更新亮点

### 参数规范化
所有命令现在都有明确的参数格式说明：
- `format`: 参数格式说明
- `examples`: 具体使用示例

例如：`/deps [模块名|all|tree|check|clean]`

### 职责边界明确
- **项目分析三剑客**：
  - `/explore`: 快速浏览（5-10分钟）
  - `/analyze`: 技术分析（15-30分钟）  
  - `/discover`: 业务理解（30-60分钟）

### 项目命令增强
- `/ai-rules`: 完整规范管理（check/apply/validate/report）
- `/guardian`: 主动监控系统（实时质量/安全/性能监控）
- `/sync`: 双向同步系统（版本管理/冲突解决/团队共享）

### 命令协调机制
参考 `commands/COMMAND_COORDINATION.md` 了解详细机制：
- **状态共享**：`.claude/state/command-state.yml` 保存命令执行状态
- **智能推荐**：基于上下文推荐下一个命令
- **链式调用**：自动触发相关命令
- **依赖管理**：确保命令按正确顺序执行

## 🚀 工作流建议

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

### 3. 知识管理流程
```
/doc-arch → /changelog → /sync-team
```
更新架构文档 → 记录变更 → 团队知识共享

### 4. 新项目启动流程
```
/meta → /doc-structure init → /setup → /test
```
定制规范 → 初始化文档结构 → 环境配置 → 测试框架

### 5. 遗留项目接入流程
```
/audit → /onboard → /discover → /retrofit → /doc-sync
```
健康检查 → 正式接入 → 理解现状 → 渐进改造 → 文档规范化

## 📚 核心文档说明

### 宪法体系
- **CLAUDE_CONSTITUTION.md**: 定义了 Claude 的协作模式、工作流程和行为准则

### 命令系统
- **GLOBAL_COMMANDS.md**: 27+ 个精心设计的全局命令
- **COMMANDS_SUMMARY.md**: 所有命令的分类和使用场景
- **COMMAND_COORDINATION.md**: 命令间协调机制和智能推荐
- **deploy-package/**: 优化后的命令部署包 v2.1.0

### 使用指南
- **AI_ASSISTANT_COMPARISON.md**: Claude vs Gemini 特性对比
- **LEGACY_PROJECT_ONBOARDING.md**: 遗留项目的详细接入指南
- **NEW_VS_LEGACY_PROJECT.md**: 不同项目类型的使用策略
- **MARKET_ANALYSIS.md**: 市场上其他方案的对比分析

## 🎯 核心理念

1. **宪法优先**: 不只是工具集，而是完整的协作体系
2. **元工作流**: AI 主导的动态规范生成
3. **双轨支持**: 同时支持新项目和遗留项目
4. **渐进改进**: 小步快跑，持续优化

## 💡 最佳实践

1. 先阅读宪法，理解协作模式
2. 根据项目类型选择合适的接入方式
3. 使用命令系统提升效率
4. 定期更新和优化规范

## 📖 元工作流使用示例

### 什么是元工作流？
元工作流是 Claude 通过对话了解项目特征，自动生成定制化 CLAUDE.md 配置的过程。

### 实际使用示例

#### 场景：为 Go 微服务项目定制规范

```
用户：/meta

Claude：我来帮您定制项目规范！请告诉我：
1. 项目类型和技术栈？
2. 开发阶段？（初期/成长期/维护期）
3. 团队规模和开发工具？

用户：这是一个 Go 微服务项目，使用 Gin 框架，PostgreSQL + GORM，
RabbitMQ 做消息队列，Docker + K8s 部署。团队 5 人，处于成长期。

Claude：了解了！还需要了解：
- 测试策略？（单元测试/集成测试/E2E）
- CI/CD 工具？
- 最常做的任务类型？
- 容易出错的地方？

用户：主要写单元测试和集成测试，用 GitLab CI/CD。
常做：1)添加新 API 2)修改数据库模型 3)调试消息队列
容易出错：忘记数据库迁移，消息队列错误处理

Claude：[生成定制的 CLAUDE.md，包含：]
- Go 微服务专用词汇表
- 项目常用命令（go test, migrate, docker-compose等）
- 架构约定（MUST/AVOID/PREFER规则）
- Claude 行为配置（自动运行测试、检查迁移等）
```

### 元工作流的优势

1. **个性化定制**：根据实际项目特征生成，而非通用模板
2. **智能对话**：通过渐进式提问收集信息
3. **立即可用**：生成的配置包含项目特定的命令和规范
4. **持续优化**：可以随时运行 `/update-constitution` 更新

### 其他项目类型的关注点

- **前端项目**：组件规范、状态管理、构建优化
- **Python 数据分析**：Notebook 管理、数据管道、实验跟踪
- **移动应用**：跨平台兼容、性能优化、发布流程

## 🔗 相关资源

- [Claude Code 官方文档](https://docs.anthropic.com/en/docs/claude-code)
- [命令系统文档](https://docs.anthropic.com/en/docs/claude-code/slash-commands)
- 本项目 GitHub: [待添加]

---

*本规范体系由 Claude 与人类开发者共同创建，持续演进中。*