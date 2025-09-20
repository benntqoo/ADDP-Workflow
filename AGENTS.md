# Repository Guidelines

## Project Structure & Modules
- Root docs: `README.md` (overview), `ARCHITECTURE.md` (interfaces/DoD), `DEVELOPMENT_PLAN.md`, `TECH_STACK.md`, `OLLAMA_MODEL_GUIDE.md`, `CLAUDE.md`.
- Code: `src/mcp_server/` (server, tools, config), `main.py` (entry), `scripts/` (quick start, tests).
- State: `.addp/` holds `specifications/`, `workflows/`, `memory/`, `queries/`, `gates/`, `sync/`, `configs/`.
- Templates: `workflow-legacy/` (prompt/command docs used for rendering).

## Build, Test, Dev Commands
- Initialize `.addp`: `python main.py --init`
- Save default config: `python main.py --save-config`
- Dev server (MCP stdio WIP): `python main.py --dev`
- Tool tests: `python scripts/test_mcp_tools.py`
- Quick start: `python scripts/quick_start.py`
- Search docs/specs: `rg "<term>" -n`
- Ollama: run `ollama serve` and pull models (see `OLLAMA_MODEL_GUIDE.md`).

## Architecture & MCP Tools
- Local‑first: terminal agent (planned) renders prompts locally and optimizes via Ollama; only then handoff to AI CLI.
- Implemented: `initialize_addp_structure`, `optimize_query`, `start_addp_workflow`, `sync_project_state`.
- Planned minimal set: `query.optimize`, `plan.update`, `fs.apply_patch` (allowlist+dry‑run), `test.run`, `guard.validate_flow`, `mem.save_phase`, `prompt.render`, `context.pack`.

## Coding Style & Naming
- Python: format with Black (line length 100); add types where practical; prefer descriptive names.
- Markdown: `lowercase-hyphen.md`; H1 Title Case; relative links; fenced code blocks with language tags.
- Keep changes minimal and focused; avoid unrelated refactors in a single PR.

## Testing Guidelines
- TDD 优先：先编写失败用例（Red）再实现（Green）；目标覆盖率≥80%（新代码）。
- 运行 `scripts/test_mcp_tools.py` 做端到端校验；门禁以 `test.run` + `guard.validate_flow` 为准（CI/pre‑push）。

## Commit & PR Guidelines
- Conventional Commits：`feat|fix|docs|refactor|chore(scope): summary`。
- PR 包含：变更摘要、影响范围、运行命令与关键输出（DoD 证据）、相关文档更新链接。

## Security & Configuration
- 禁止提交密钥/敏感数据；审慎处理日志与示例。
- `fs.apply_patch`（规划）仅允许白名单目录；避免路径遍历与越权写入。
- Windows：优先使用 Python 命令；bash 脚本可在 Git Bash/WSL 运行。
