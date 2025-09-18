# 跨 CLI 适配（Adapters）

本层用于将统一命令（plan.update、fs.apply_patch、run.test、guard.validate_flow）映射到各类 AI Coding CLI 使用的调用方式，优先通过本地 MCP 路径实现一致行为。

- 入口：`python -m tools.adapters.dispatch --platform <cli> --command <name> --args '{"k":"v"}'`
- 支持平台（初始）：`claude-code`、`codex`、`cursor`、`aider`
- 共享清单：`tools/commands.yaml`

示例
- 更新计划（Cursor）：
  `python -m tools.adapters.dispatch --platform cursor --command plan.update --args '{"task":"Add rule"}'`
- 应用补丁（Codex）：
  `python -m tools.adapters.dispatch --platform codex --command fs.apply_patch --args '{"patch":"*** Begin Patch..."}'`

合规与校验
- 所有适配器最终路由到 `scripts/mcp_run.py` 并在 `.claude/state/command-state.yml` 记录一次运行。
- 提交前可执行：`python scripts/mcp_validate.py`，用于最小化合规检查。

扩展
- 新增平台：在 `tools/adapters/` 新建 `<platform>.py` 并在 `dispatch.py` 注册。
- 如需平台别名或参数变换，在适配器 `map_command` 中实现。
