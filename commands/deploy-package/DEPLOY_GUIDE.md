# Claude Code 命令部署指南

本包包含了完整的 Claude Code 命令集，可以快速部署到新的开发环境。

## 📦 包内容

```
deploy-package/
├── global/          # 全局命令（25个）
├── project/         # 项目命令（3个）
├── DEPLOY_GUIDE.md  # 本部署指南
└── deploy.sh        # 自动部署脚本
```

## 🚀 快速部署

### Windows (PowerShell)
```powershell
# 1. 创建全局命令目录
New-Item -ItemType Directory -Force -Path "$env:USERPROFILE\.claude\commands"

# 2. 复制全局命令
Copy-Item -Path ".\global\*" -Destination "$env:USERPROFILE\.claude\commands\" -Force

# 3. 在项目中创建命令目录
New-Item -ItemType Directory -Force -Path ".\.claude\commands"

# 4. 复制项目命令
Copy-Item -Path ".\project\*" -Destination ".\.claude\commands\" -Force
```

### macOS/Linux (Bash)
```bash
# 1. 创建全局命令目录
mkdir -p ~/.claude/commands

# 2. 复制全局命令
cp ./global/* ~/.claude/commands/

# 3. 在项目中创建命令目录
mkdir -p ./.claude/commands

# 4. 复制项目命令
cp ./project/* ./.claude/commands/
```

## 📋 命令清单

### 全局命令（25个）

#### 核心工作流（7个）
- `meta.md` - 元工作流，定制项目规范
- `onboard.md` - 遗留项目接入向导
- `update-constitution.md` - 更新 CLAUDE.md 版本
- `constitution.md` - 检查并应用协作宪法
- `deep.md` - 深度开发工作流
- `explore.md` - 探索式工作流
- `batch.md` - 批量操作模式

#### 代码分析（3个）
- `analyze.md` - 深度代码分析
- `audit.md` - 项目健康度审计
- `discover.md` - 理解特定功能模块

#### 文档管理（6个）
- `doc-api.md` - API 文档生成
- `doc-arch.md` - 架构文档更新
- `doc-structure.md` - 文档目录管理
- `doc-sync.md` - 文档一致性检查
- `changelog.md` - 变更日志管理
- `readme.md` - README 生成

#### 企业架构（3个）
- `deps.md` - 模块依赖分析
- `microservice.md` - 微服务协调
- `migrate.md` - 数据库迁移管理

#### 测试质量（3个）
- `test.md` - 智能测试生成
- `coverage.md` - 测试覆盖率分析
- `perf.md` - 性能分析优化

#### 部署运维（3个）
- `deploy-check.md` - 部署前检查
- `rollback.md` - 回滚方案生成
- `config.md` - 环境配置管理

### 项目命令（3个）
- `ai-rules.md` - 检查 AI 协作规范
- `guardian.md` - 启用守护者模式
- `sync.md` - 执行宪法同步

## ⚠️ 已知问题和优化建议

### 1. 命令参数规范化
以下命令使用了 `$ARGUMENTS` 但未明确说明参数格式：
- `config`
- `deps`
- `perf`
- `readme`
- `doc-api`
- `discover`
- `microservice`
- `migrate`
- `rollback`

**建议**：使用时根据实际需求补充参数说明。

### 2. 推荐工作流

#### 新项目启动
```
/meta → /doc-structure init → /test
```

#### 遗留项目接入
```
/audit → /onboard → /discover → /retrofit
```

#### 日常开发
```
/discover → /test → /coverage → /deploy-check
```

#### 故障处理
```
/audit → /rollback → /config → /sync-team
```

### 3. 命令间协调
- `meta`、`update-constitution`、`constitution` 都涉及 CLAUDE.md，使用时注意顺序
- 文档命令建议统一使用 `/doc-structure` 管理输出目录

## 🔧 环境要求

- Claude Code CLI 已安装
- Graphiti MCP（用于 `/sync-team` 命令）
- 项目根目录有 `.claude/` 目录

## 📝 注意事项

1. **重启生效**：部署后需要重启 Claude Code 才能识别新命令
2. **权限检查**：确保命令文件有读取权限
3. **路径规范**：项目命令必须在项目根目录的 `.claude/commands/` 下
4. **版本兼容**：这些命令基于 Claude Code 最新版本设计

---

*部署包版本：2.0.0*
*生成日期：2024-01-15*