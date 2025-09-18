# Repository Guidelines

## Project Structure & Module Organization
- Root: `README.md` (overview), `TARGET.md`/`TARGET_IMPROVED.md`/`CLAUDE.md` (specs).
- `workflow-legacy/`: legacy-but-authoritative docs and assets:
  - `agents/`: role/language/framework profiles, e.g. `workflow-legacy/agents/languages/python-ml-specialist.md`.
  - `commands/`: `deploy-package/` specs and global command docs.
  - `guides/`, `docs/`, `templates/`, `config/`, `constitution/`, `output-styles/`, `monitoring/`.
  - `scripts/`: maintenance utilities (e.g., `workflow-legacy/scripts/check-file-consistency.sh`).

## Build, Test, and Development Commands
- No build step; this repo is documentation-first.
- Validate references: `bash workflow-legacy/scripts/check-file-consistency.sh`
  - Windows: run via Git Bash or WSL.
- Fast search: `rg "<term>" -n` to locate examples/specs.
- Preview Markdown with VS Code or GitHub to verify anchors and tables.

## Coding Style & Naming Conventions
- Filenames: `lowercase-hyphen.md` inside subfolders; top-level meta docs may be `UPPERCASE.md`.
- Headings: H1 in Title Case; concise H2/H3; keep sections focused.
- Links: use relative paths; update anchors when files or titles change.
- Code blocks: fenced with language tags (`bash`, `json`, `yaml`); keep commands copy‑pastable.
- Agent profiles: include YAML frontmatter fields (`name`, `model`, `description`, `trigger`, `tools`).

## Testing Guidelines
- Run the consistency script before committing.
- Check internal links, images, and Mermaid diagrams render correctly.
- Keep paired docs in sync (e.g., `agents/*` profiles and related `guides/` or `commands/` summaries).

## Commit & Pull Request Guidelines
- Prefer Conventional Commits (seen in history): `feat|fix|docs|refactor|chore(scope): summary`.
- Examples: `docs(agents): add rust profile`; `fix(docs): correct link in commands/deploy-package/global/check.md`.
- PRs include: clear summary, rationale, impacted paths, linked issues, and updated cross‑references; add screenshots for visual changes.

## Security & Configuration Tips
- Do not commit secrets or private data; sanitize examples.
- Avoid large binaries; store heavy media externally if needed.
- Use LF line endings and UTF‑8 encoding.

