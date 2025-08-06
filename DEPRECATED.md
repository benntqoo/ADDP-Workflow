# 已废弃文档说明

本文档列出了已经过时或废弃的文档，这些文档保留仅供历史参考。

## 🚫 已废弃文档列表

### commands/ 目录
1. **CONSTITUTION_INJECT_COMMANDS.md** (**已删除**)
   - 废弃原因：宪法注入的概念已被更简单的方案替代
   - 替代方案：
     - 新项目：使用 `/meta` 一键生成完整 CLAUDE.md
     - 更新版本：使用 `/update-constitution` 智能更新
   - 优势：新方案更简单，同时保留了版本管理功能

2. **SYNC_COMMANDS.md**
   - 废弃原因：同步命令的概念已整合到标准工作流中
   - 替代方案：`/sync-team` 用于团队知识同步

### guides/ 目录
1. **CONSTITUTION_SYNC_GUIDE.md**
   - 废弃原因：宪法同步流程已过时
   - 替代方案：CLAUDE.md 自动包含所有必要配置

## ✅ 当前推荐做法

### 新项目启动
```bash
/meta  # 一键生成完整的 CLAUDE.md，包含文档规范
```

### 现有项目接入
```bash
/onboard  # 智能扫描并生成适配的 CLAUDE.md
```

### 团队知识同步
```bash
/sync-team  # 使用 Graphiti MCP 同步团队知识
```

### examples/ 目录
1. **META_WORKFLOW_EXAMPLE.md** (**已删除**)
   - 废弃原因：示例内容已整合到主 README
   - 替代方案：README.md 中的"元工作流使用示例"章节
   - 优势：避免内容重复，集中文档管理

### commands/ 目录
3. **旧版命令描述**（**已更新**）
   - 更新内容：v2.1 版本优化了所有命令
   - 主要变化：
     - 添加了参数格式说明（format 和 examples）
     - 明确了命令职责边界
     - 增强了项目命令功能
   - 查看最新：`COMMANDS_SUMMARY.md` v2.1 版

4. **OPTIMIZATION_PLAN.md**（**已删除**）
   - 废弃原因：优化任务已全部完成
   - 成果记录：`deploy-package/CHANGELOG.md`
   - 部署包版本：v2.1.0

## 📝 文档维护建议

1. **保留但不更新**：这些过时文档暂时保留，但不再维护
2. **新用户引导**：确保新用户不会误用过时文档
3. **定期清理**：在下个主要版本时考虑移除这些文档

---

*更新日期：2024-01-15*