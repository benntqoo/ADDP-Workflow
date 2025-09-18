# Repository Guidelines

## 项目结构与模块组织
- `agents/`：按角色/语言/框架划分的代理规范与能力档案，例如 `agents/languages/python-ml-specialist.md`。
- `commands/`：部署与文档相关命令（Markdown 指南），`deploy-package/` 提供常用工作流说明。
- `scripts/`：维护脚本，如 `scripts/check-file-consistency.sh`。
- `guides/`、`docs/`：使用指南与标准文档。
- `config/`、`constitution/`：配置与宪章约束。
- `.claude/`：模块内记忆与状态（`memory/`、`reports/`、`state/`）。
- `.github/`：Issue/PR 模板。

## 构建、测试与开发命令
- 核心引用校验：从仓库根目录运行 `bash claude/scripts/check-file-consistency.sh`（Windows 建议 Git Bash/WSL）。
- 全局检索：`rg "<关键词>" -n` 快速定位规范与示例。
- 文档预览：使用 VS Code/GitHub 预览，检查目录与锚点渲染。

## 编码风格与命名
- 文件名：`lowercase-hyphen.md`；代理档案按类别置于对应子目录。
- 标题：H1 用 Title Case；小节简洁直述，不冗长。
- 链接：使用相对路径；引用全局上下文用 `../.claude/...`，引用子模块上下文用 `./.claude/...`。
- 代码块：使用围栏并标注语言（如 ```bash ）。

## 测试指南
- 提交前运行一致性脚本，修正根目录直连的核心文件引用问题。
- 校验链接有效、目录/代码块渲染正常。
- 变更成对更新：代理规范与其使用说明同步修改。

## 提交与 PR 规范
- 提交信息遵循 `type(scope): summary`，示例：`docs(agents): add rust profile`、`fix(docs): correct .claude path`。
- 使用 `claude/.github/PULL_REQUEST_TEMPLATE.md`；PR 包含：变更摘要、影响范围、截图/片段与关联 Issue。

## 安全与配置
- 禁止提交密钥；敏感配置使用本地环境变量/私密存储，勿写入文档。
- 当工作流/约定更新时，同步维护 `claude/.claude/memory/DECISIONS.md` 与相关指南。
