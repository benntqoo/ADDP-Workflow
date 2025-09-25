# Repository Guidelines

> AGENTS Contract v1.1 · Last-Updated: 2025-09-24 · Owner: Core Maintainers

This AGENTS.md defines how agents and contributors operate here. It covers scope, precedence, encoding, Windows specifics, and safe-edit rules. Nested AGENTS.md files override this one in their subtree; explicit user/developer instructions override AGENTS.md.

Scope & Precedence
- Scope: Applies to the entire repo rooted at this file’s directory.
- Overrides: Nested AGENTS.md > this file; Prompt/Task > AGENTS.md.
- Enforcement: Collect all applicable AGENTS.md for files you modify and honor the highest-precedence rules.

Change Control
- Propose changes via PR with Motivation, Scope, Examples, Do/Don’t, Impact, Migration.
- Mark breaking changes and include transition notes if needed.
- Keep a short “What changed for agents” summary in the PR body.

## Project Structure & Modules
- Root docs: `README.md`, `ARCHITECTURE.md`, `DEVELOPMENT_PLAN.md`, `TECH_STACK.md`, `OLLAMA_MODEL_GUIDE.md`, `CLAUDE.md`.
- Code: `src/mcp_server/` (server/tools/config), `main.py` (entry), `scripts/` (quick start/tests).
- State: `.addp/` holds `specifications/`, `workflows/`, `memory, `queries/`, `gates/`, `sync/`, `configs/`.
- Templates: `workflow-legacy/` (prompt/command docs for rendering).

## Build, Test, Dev Commands
- Initialize `.addp`: `python main.py --init`
- Save default config: `python main.py --save-config`
- Dev server (MCP stdio WIP): `python main.py --dev`
- Tool tests: `python scripts/test_mcp_tools.py`
- Quick start: `python scripts/quick_start.py`
- Search docs/specs: `rg "<term>" -n`
- Ollama: run `ollama serve` and pull models (see `OLLAMA_MODEL_GUIDE.md`).

## Architecture & MCP Tools
- Local-first: terminal agent (planned) renders prompts locally and optimizes via Ollama; then hand off to AI CLI.
- Implemented: `initialize_addp_structure`, `optimize_query`, `start_addp_workflow`, `sync_project_state`.
- Planned minimal set: `query.optimize`, `plan.update`, `fs.apply_patch` (allowlist+dry-run), `test.run`, `guard.validate_flow`, `mem.save_phase`, `prompt.render`, `context.pack`.

## Coding Style & Naming
- Python: Black (line length 100); add types where practical; descriptive names.
- Markdown: `lowercase-hyphen.md`; H1 Title Case; relative links; fenced code blocks with language tags.
- Keep changes minimal and focused; avoid unrelated refactors in a single PR.

## i18n & Encoding (UTF-8, no BOM)
- All source/text files use UTF-8 (no BOM). Do not commit ANSI/GBK/UTF-16.
- Eliminate “mojibake” (乱码/锟斤拷等); replace with proper Simplified Chinese.
- Fyne GUI: prefer CJK TTF (e.g., `simhei.ttf`, `msyh.ttf`). Set `FYNE_FONT` if needed. See `docs/ui/CJK_FONT.md`.

## PowerShell Writing (Windows)
- Always write files as UTF-8 (no BOM):
  - `Set-Content -Encoding UTF8`
  - `Out-File -Encoding utf8`
  - `[IO.File]::WriteAllText($p,$text,[Text.UTF8Encoding]::new($false))`
- Console output for Chinese:
  - `chcp 65001`
  - `$OutputEncoding = [Console]::OutputEncoding = [Text.UTF8Encoding]::new($false)`
- Compose/YAML: avoid fragile multiline command lists; prefer single-line `bash -lc` chains; mind CRLF.

## Testing Guidelines
- TDD 优先：先编写失败用例（Red）再实现（Green）；新代码覆盖率≥80%。
- 运行 `scripts/test_mcp_tools.py` 做端到端校验；门禁以 `test.run` + `guard.validate_flow` 为准（CI/pre‑push）。

## Commit & PR Guidelines
- Conventional Commits：`feat|fix|docs|refactor|chore(scope): summary`。
- PR 包含：变更摘要、影响范围、运行命令与关键输出（DoD 证据）、相关文档更新链接。

## Security & Configuration
- 禁止提交密钥/敏感数据；审慎处理日志与示例。
- `fs.apply_patch`（规划）仅允许白名单目录；避免路径遍历与越权写入。
- Windows：优先使用 Python 命令；bash 脚本可在 Git Bash/WSL 运行。

## CI Hooks & Checks（建议）
- Lint encoding: ensure UTF-8 (no BOM) and LF endings where applicable.
- Detect mojibake: fail if suspicious bytes introduced.
- Disallow writes outside allowed directories from automation.
- Optional: pre-commit hook to normalize EOL and reject mixed encodings.

## Docker/Compose Notes（Windows）
- 在 `build/` 目录执行 compose 时，不要再传 `-f .\build\docker-compose.yml` 以免路径变成 `build\build\...`。
- 在容器内导入 PATH 时用 `$$PATH`（防止宿主 PATH 注入）。
- 组合长命令尽量写成单行 `bash -lc "... && ..."`，避免 CRLF 断行。

## 发现并清理“中文乱码”的建议流程
- 用 `rg -n "[\x80-\xFF]" --glob "**/*.go"` 初筛非 ASCII 文本；结合上下文判断是否乱码。
- 将不可读的“伪中文”替换为正常简体中文；避免多行/未闭合字符串导致 `newline in string`。
- 保存为 UTF-8（无 BOM），并编译自测，确保消除转义/换行类错误。
- GUI 文案优先简洁直白；必要时提供图标/状态提示。

## References
- 中文字体与显示：`docs/ui/CJK_FONT.md`

