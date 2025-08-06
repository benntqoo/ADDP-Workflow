# Claude Code 命令总览 v2.1

本文档包含 Claude Code 协作体系中所有可用的命令，包括全局命令和项目命令。

## 📊 命令统计

- **核心工作流**: 7 个命令
- **项目分析**: 3 个命令（职责明确）
- **文档管理**: 6 个命令
- **企业架构**: 3 个命令
- **团队协作**: 3 个命令
- **测试质量**: 3 个命令
- **部署运维**: 3 个命令
- **项目专用**: 3 个命令（功能增强）

**总计**: 31 个命令，覆盖从项目初始化到部署运维的完整生命周期。

## 🆕 v2.1 版本更新

1. **参数规范化**：所有命令都有明确的 `format` 和 `examples`
2. **职责边界明确**：解决了功能重叠问题
3. **项目命令增强**：从简单功能升级为完整系统
4. **命令协调机制**：支持状态共享和智能推荐

## 🎯 核心工作流命令

| 命令 | 功能描述 | 参数格式 | 使用场景 |
|------|----------|----------|----------|
| `/meta` | 启动元工作流，定制项目规范 | 无参数 | **新项目**初始化或规范定制 |
| `/onboard` | 遗留项目接入向导 | 无参数 | **现有项目**引入 Claude Code |
| `/update-constitution` | 更新 CLAUDE.md 版本 | 无参数 | 保留定制内容，更新标准部分 |
| `/constitution` | 检查并应用协作宪法 | 无参数 | 确认协作模式和规范 |
| `/deep` | 启动深度开发工作流 | 无参数 | 复杂任务的系统化实施 |
| `/explore` | 快速浏览项目结构 | 无参数 | 生成项目地图（5-10分钟） |
| `/batch` | 启用批量操作模式 | 无参数 | 需要并行处理多个操作时 |

## 🔍 项目分析命令（职责明确）

| 命令 | 功能描述 | 参数格式 | 执行时间 |
|------|----------|----------|----------|
| `/explore` | 快速浏览，生成项目地图 | 无参数 | 5-10分钟 |
| `/analyze` | 技术栈分析，依赖检查 | 无参数 | 15-30分钟 |
| `/discover` | 业务逻辑理解，数据流分析 | `[模块名\|功能名\|flow\|data\|risk]` | 30-60分钟 |

### 使用建议
- 初次接触项目：先 `/explore` 获得全局认识
- 技术调研：使用 `/analyze` 深入技术细节
- 功能开发：用 `/discover` 理解业务逻辑

## 📚 文档管理命令

| 命令 | 功能描述 | 参数格式 | 输出位置 |
|------|----------|----------|----------|
| `/doc-api` | API 文档生成 | `[模块名\|all\|openapi\|graphql\|postman]` | `docs/api/` |
| `/doc-arch` | 架构文档更新 | 无参数 | `docs/architecture/` |
| `/changelog` | 变更日志管理 | 无参数 | `docs/releases/CHANGELOG.md` |
| `/doc-sync` | 文档一致性检查 | 无参数 | 控制台报告 |
| `/readme` | README 生成 | `[项目路径\|update\|template\|check]` | `./README.md` |
| `/doc-structure` | 文档目录管理 | `[init\|check]` | `docs/` 目录结构 |

## 🏗️ 企业级架构命令

| 命令 | 功能描述 | 参数格式 | 特色功能 |
|------|----------|----------|----------|
| `/deps` | 模块依赖分析 | `[模块名\|all\|tree\|check\|clean]` | 循环依赖检测、依赖树可视化 |
| `/microservice` | 微服务协调 | `[analyze\|contract\|sync\|test\|monitor\|服务名]` | 服务拓扑、契约管理 |
| `/migrate` | 数据库迁移管理 | `[status\|create\|plan\|validate\|rollback\|迁移名]` | 风险评估、回滚方案 |

## 👥 团队协作命令

| 命令 | 功能描述 | 参数格式 | 集成说明 |
|------|----------|----------|----------|
| `/review` | 代码审查 | - | 使用 Claude Code 内置命令 |
| `/pr_comments` | PR 评论查看 | - | 使用 Claude Code 内置命令 |
| `/sync-team` | 团队知识同步 | 无参数 | 集成 Graphiti 知识图谱 |

## 🧪 测试质量命令

| 命令 | 功能描述 | 参数格式 | 覆盖范围 |
|------|----------|----------|----------|
| `/test` | 智能测试生成 | `[代码路径\|--setup\|--generate-missing]` | 单元/集成/边界测试 |
| `/coverage` | 测试覆盖率分析 | `[--detail]` | 覆盖率报告、盲点识别 |
| `/perf` | 性能分析优化 | `[模块名\|api\|db\|frontend\|all]` | 瓶颈识别、优化建议 |

## 🚀 部署运维命令

| 命令 | 功能描述 | 参数格式 | 检查项目 |
|------|----------|----------|----------|
| `/deploy-check` | 部署前检查 | `[--quick]` | 代码/测试/安全/文档 |
| `/rollback` | 回滚方案生成 | `[code\|db\|config\|all\|emergency]` | 快速/标准/紧急回滚 |
| `/config` | 环境配置管理 | `[环境名\|check\|init\|sync]` | 配置验证、安全检查 |

## 🔒 项目专用命令（增强版）

| 命令 | 功能描述 | 参数格式 | v2.1 新增功能 |
|------|----------|----------|--------------|
| `/sync` | 双向智能同步系统 | `[up\|down\|status\|force\|rollback]` | 版本管理、冲突解决、团队共享 |
| `/ai-rules` | 完整规范管理器 | `[check\|apply\|validate\|report]` | 合规验证、自动修复、CI/CD集成 |
| `/guardian` | 主动监控系统 | `[on\|off\|status\|config\|report]` | 实时监控、安全检查、性能分析 |

## 🔄 推荐工作流

### 1. 项目初识流程
```
/explore → /analyze → /discover main
```
快速浏览 → 技术分析 → 业务理解

### 2. 日常开发流程
```
/discover feature → /test → /coverage → /perf → /deploy-check
```
理解功能 → 编写测试 → 检查覆盖率 → 性能优化 → 部署准备

### 3. 故障处理流程
```
/audit → /rollback prepare → /config check → /sync-team
```
问题诊断 → 准备回滚 → 配置检查 → 知识同步

### 4. 知识管理流程
```
/doc-arch → /changelog → /sync-team
```
更新架构文档 → 记录变更 → 团队知识共享

### 5. 新项目启动流程
```
/meta → /doc-structure init → /ai-rules apply → /guardian on
```
定制规范 → 初始化文档 → 应用规则 → 开启监控

### 6. 遗留项目接入流程
```
/audit → /onboard → /discover → /retrofit → /doc-sync
```
健康检查 → 正式接入 → 理解现状 → 渐进改造 → 文档规范化

## 💡 命令使用技巧

### 参数说明
- `[参数]` 表示可选参数
- `|` 表示参数选项分隔符
- 无参数的命令直接使用

### 命令协调
- 命令执行状态保存在 `.claude/state/command-state.yml`
- 支持基于上下文的智能推荐
- 部分命令支持自动触发相关命令

### 批量操作
- 使用 `/batch` 启用并行执行模式
- 适合多文件操作和批量处理

## 📦 快速部署

使用优化后的部署包 v2.1.0：
```bash
# 获取部署包
cd claude/commands/deploy-package

# Windows
./deploy.ps1

# macOS/Linux  
./deploy.sh
```

## 🔧 自定义命令

创建自定义命令的基本结构：
```markdown
---
arguments: optional
format: "[参数格式说明]"
examples:
  - "/命令 参数1 - 说明"
  - "/命令 参数2 - 说明"
---

命令功能描述...
```

## 🎯 选择建议

### 根据项目阶段
- **新项目**：从 `/meta` 开始
- **遗留项目**：从 `/audit` 开始
- **日常开发**：常用 `/discover`、`/test`、`/deploy-check`

### 根据任务类型
- **理解代码**：`/explore` → `/analyze` → `/discover`
- **质量保障**：`/test` → `/coverage` → `/guardian`
- **文档维护**：`/doc-sync` → `/doc-api` → `/readme`

---

*本文档是 Claude Code 命令体系的完整参考。v2.1 版本优化了所有命令的参数定义、职责边界和协作机制。*