# Repository Guidelines

## Project Structure & Modules
- Root docs: `README.md` (overview), `ARCHITECTURE.md` (design & interfaces), `DEVELOPMENT_PLAN.md`, `CLAUDE.md`, `TECH_STACK.md`.
- Source: `src/mcp_server/` (MCP server, tools, config), `main.py` (entry), `scripts/` (quick start, tests).
- State: `.addp/` for specs, workflows, memory, queries, gates, sync, configs.
- Legacy references: `workflow-legacy/` holds historical templates/commands for prompt rendering.

## Build, Test, Dev Commands
- Initialize `.addp`: `python main.py --init`
- Save default config: `python main.py --save-config`
- Run dev server (stdio wiring WIP): `python main.py --dev`
- Tool tests: `python scripts/test_mcp_tools.py`
- Search docs/specs: `rg "<term>" -n`
- Ollama: ensure local service and model (e.g., `qwen2.5:14b`) are available.

## Architecture & MCP Tools
- Local‑first: terminal代理先做模板渲染与本地优化，再按需 handoff 到 AI CLI。
- Implemented tools: `initialize_addp_structure`, `optimize_query`, `start_addp_workflow`, `sync_project_state`。
- Planned minimal set (aliases incoming): `query.optimize`, `plan.update`, `fs.apply_patch`(allowlist+dry‑run), `test.run`, `guard.validate_flow`, `mem.save_phase`, `prompt.render`, `context.pack`。

## Coding Style & Naming
- Python: format with Black (line length 100); add types where practical; prefer descriptive names.
- Markdown: `lowercase-hyphen.md`; H1 Title Case; relative links; fenced code blocks with language tags.
- Keep changes minimal and focused; avoid unrelated refactors in a single PR.

## Testing Guidelines
- TDD 优先：新增/变更 MCP 工具前先补测试；目标覆盖率≥80%（新代码）。
- 运行 `scripts/test_mcp_tools.py` 做最小端到端验证；后续以 `test.run` + `guard.validate_flow` 作为门禁。

## Commit & PR Guidelines
- Conventional Commits：`feat|fix|docs|refactor|chore(scope): summary`。
- PR 应包含：变更摘要、影响范围、运行命令与关键输出（DoD 证据）、相关文档更新链接。

## Security & Configuration
- 禁止提交密钥/敏感数据；审慎处理日志与示例。
- `fs.apply_patch` 仅允许白名单目录；避免路径遍历与越权写入。
- Windows 建议用 Python 命令替代 bash；必要时用 Git Bash/WSL 运行脚本。
