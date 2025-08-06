# 📚 企业级项目文档目录结构规范

## 🎯 核心原则

1. **集中管理**：所有文档统一存放在 `docs/` 目录下
2. **分类清晰**：按文档类型和用途分类存放
3. **版本控制**：重要文档保留历史版本
4. **易于导航**：目录结构直观，命名规范

## 📁 标准目录结构

```
project-root/
├── docs/                       # 所有文档的根目录
│   ├── README.md              # 文档目录索引
│   ├── api/                   # API 文档
│   │   ├── README.md          # API 文档概览
│   │   ├── openapi.yaml       # OpenAPI/Swagger 规范
│   │   ├── rest/              # REST API 文档
│   │   │   ├── v1/            # 版本化 API 文档
│   │   │   └── v2/
│   │   ├── graphql/           # GraphQL Schema
│   │   └── examples/          # API 使用示例
│   │
│   ├── architecture/          # 架构文档
│   │   ├── README.md          # 架构概览
│   │   ├── decisions/         # 架构决策记录 (ADRs)
│   │   │   ├── 0001-use-microservices.md
│   │   │   └── 0002-database-selection.md
│   │   ├── diagrams/          # 架构图
│   │   │   ├── system-architecture.mmd
│   │   │   ├── data-flow.puml
│   │   │   └── deployment.svg
│   │   └── components/        # 组件详细设计
│   │
│   ├── guides/                # 使用指南
│   │   ├── getting-started.md # 快速开始
│   │   ├── installation.md    # 安装指南
│   │   ├── configuration.md   # 配置说明
│   │   ├── deployment.md      # 部署指南
│   │   └── troubleshooting.md # 故障排查
│   │
│   ├── development/           # 开发文档
│   │   ├── setup.md           # 开发环境设置
│   │   ├── coding-standards.md # 编码规范
│   │   ├── testing.md         # 测试指南
│   │   ├── contributing.md    # 贡献指南
│   │   └── workflow.md        # 开发流程
│   │
│   ├── references/            # 参考文档
│   │   ├── glossary.md        # 术语表
│   │   ├── faq.md            # 常见问题
│   │   ├── dependencies.md    # 依赖说明
│   │   └── tools.md          # 工具使用
│   │
│   ├── releases/              # 发布相关
│   │   ├── CHANGELOG.md       # 变更日志
│   │   ├── ROADMAP.md         # 产品路线图
│   │   ├── migration/         # 迁移指南
│   │   │   ├── v1-to-v2.md
│   │   │   └── v2-to-v3.md
│   │   └── notes/             # 发布说明
│   │
│   └── internal/              # 内部文档（不对外公开）
│       ├── team/              # 团队文档
│       ├── meetings/          # 会议记录
│       └── planning/          # 规划文档
│
├── .claude/                   # Claude Code 配置
│   └── commands/              # 项目命令
│
└── README.md                  # 项目主 README

```

## 📝 文档命名规范

### 文件命名
- 使用小写字母和连字符：`user-guide.md`
- 版本化文档加版本号：`api-v2.0.md`
- 日期相关文档使用 ISO 格式：`2024-01-15-meeting-notes.md`

### 目录命名
- 使用小写字母
- 多个单词用连字符连接
- 保持简短明了

## 🔄 文档生成映射

| 命令 | 生成位置 | 文件名规范 |
|------|----------|------------|
| `/doc-api` | `docs/api/` | `{module}-api.yaml` |
| `/doc-arch` | `docs/architecture/` | 相应子目录 |
| `/changelog` | `docs/releases/` | `CHANGELOG.md` |
| `/readme` | 项目根目录 | `README.md` |
| 其他文档 | `docs/` 相应类别 | 按类别规范 |

## 🏷️ 文档模板

### API 文档模板
```markdown
# {API 名称}

## 概述
简要描述 API 的用途和功能

## 认证
说明认证方式

## 端点列表
列出所有可用端点

## 详细说明
每个端点的详细文档
```

### 架构决策记录 (ADR) 模板
```markdown
# ADR-{编号}: {标题}

## 状态
{提议 | 接受 | 弃用 | 替代}

## 背景
问题描述和决策背景

## 决策
做出的具体决策

## 后果
决策带来的影响
```

## 🚀 使用指南

1. **初始化文档结构**
   ```bash
   /doc-structure init
   ```

2. **生成文档时自动归类**
   - 各文档生成命令会自动将文档放到正确位置
   - 遵循命名规范

3. **定期整理**
   ```bash
   /doc-sync
   ```

## 📌 最佳实践

1. **版本控制**
   - 所有文档纳入 Git 管理
   - 重要变更要有提交说明

2. **定期审查**
   - 每个版本发布前审查文档
   - 删除过时内容
   - 更新不准确信息

3. **团队协作**
   - 文档变更需要 review
   - 保持风格一致
   - 及时更新索引

## 🔗 相关命令

- `/doc-structure init` - 初始化文档目录结构
- `/doc-structure check` - 检查文档结构规范性
- `/doc-sync` - 同步和整理文档

---

*本规范是项目文档管理的基础，确保文档有序、易查、易维护。*