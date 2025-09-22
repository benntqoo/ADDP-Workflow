# 技术栈选型（TECH_STACK）

## 总览
- 目标：构建“终端版 AI IDE 代理”+ MCP 服务，先本地优化/执行，再按需转交 AI CLI，降低 token 消耗并标准化流程（ADDP）。
- 语言：Python 3.8+（复用现有 src/ 模块，跨平台友好）。

## 终端代理（CLI/TUI）
- CLI 框架：Typer（快速命令定义、类型提示好）。
- TUI（可选）：Textual；Windows 无 tmux 时提供纯 TUI 回退。
- 进程桥接：`subprocess` 调用 `claude`/`codex`/`gemini`/`aider` 等；统一 handoff 适配器封装 I/O。
- Git 集成：优先直接调用 `git`（worktree/branch/commit/push），必要时引入 GitPython。

## MCP 服务与工具
- 协议：`mcp` Python 包；传输首选 stdio。
- 模块复用：`src/mcp_server/*`（Initializer/Query/Workflow/Sync）。
- 最小工具集：`query.optimize`、`plan.update`、`fs.apply_patch`（白名单）、`test.run`、`guard.validate_flow`、`mem.save_phase`、`prompt.render`、`context.pack`。

## 本地智能与提示词
- 本地 LLM：Ollama（默认 `qwen2.5:14b`，可切换 llama/mistral）。
- HTTP 客户端：`aiohttp`；加强 JSON 容错（code-fence/松散 JSON）。
- 模板渲染：将 `workflow-legacy/commands/deploy-package/global/update-spec.md` 等转为可参数化提示词；本地优化后再 handoff。

## 持久化与状态
- 目录：`.addp/{specifications,workflows,memory,queries,gates,sync,configs}`。
- 存储：JSON 文件（轻量可审计），记录 `MCP_RUN_ID`、阶段快照、门禁报告；后续可选 SQLite。

## 测试与质量
- 测试：`pytest`、`pytest-asyncio`；脚本级 `scripts/test_mcp_tools.py`。
- 格式/类型：`black`、`mypy`；CI 中融合 `guard.validate_flow`（白名单/阶段顺序/测试结果/DoD）。

## 分发与运行
- 包管理：`setuptools` + `pyproject.toml`；提供 `console_scripts`（代理 CLI）与 MCP 入口。
- 运行模式：`cli`（本地执行/优化）与 `stdio`（MCP）。
- 依赖：`ollama`、`git`；可选 `tmux`（并行会话）。

## 选型理由与备选
- 理由：复用现有 Python 模块；跨平台简单；与 MCP/Ollama 生态贴合；便于快速迭代。
- 备选：Go/Node 重写 CLI（更快、易单 binary 发布）；功能稳定后评估迁移。
