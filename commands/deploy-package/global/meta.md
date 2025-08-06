我想定制项目的协作规范，让我们开始元工作流。请帮我分析项目特征并生成适合的 CLAUDE.md 配置。

## 生成规范

生成的 CLAUDE.md 将：
1. **包含版本信息**
   ```markdown
   <!-- 
   Claude Constitution Version: 2.0.0
   Generated: [日期]
   Template: claude/templates/CLAUDE_MD_TEMPLATE.md
   -->
   ```

2. **标记可定制区域**
   ```markdown
   <!-- LOCAL:BEGIN -->
   项目特定配置...
   <!-- LOCAL:END -->
   ```

3. **包含核心部分**
   - 项目特定的技术栈和开发规范
   - 工作流程定义
   - 文档管理规范（基于标准文档目录结构）
   - 代码质量标准
   - 团队协作约定

## 文档结构规范
所有文档统一存放在 docs/ 目录下：
- API 文档 → docs/api/
- 架构文档 → docs/architecture/
- 使用指南 → docs/guides/
- 开发文档 → docs/development/
- 发布相关 → docs/releases/

## 后续维护
生成后可使用 `/update-constitution` 更新到最新版本，同时保留您的定制内容。