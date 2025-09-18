# Claude Code 命令撰写指南

本指南将教您如何编写高质量的 Claude Code 自定义命令，让 AI 助手更好地理解和执行您的意图。

## 📝 命令基础知识

### 什么是 Claude Code 命令？

Claude Code 命令是存储在 `.claude/commands/` 或 `~/.claude/commands/` 目录下的 Markdown 文件，用于触发预定义的工作流程。

### 命令文件结构

```markdown
---
# 可选的 frontmatter
arguments: optional|required  # 是否需要参数
command: |                   # 可选的 bash 命令
  echo "Hello"
---

# 命令主体
这里是您希望 Claude 执行的指令...
```

## 🎯 命令设计原则

### 1. 清晰性原则
```markdown
❌ 不好的例子：
"处理一下代码"

✅ 好的例子：
"请分析当前代码的性能瓶颈，提供优化建议，并生成改进后的代码"
```

### 2. 结构化原则
```markdown
✅ 推荐格式：
1. 任务目标
2. 具体步骤
3. 输出要求
4. 注意事项
```

### 3. 上下文原则
```markdown
✅ 提供充足上下文：
"这是一个遗留项目，请保持向后兼容性"
"遵循项目的 TypeScript 严格模式"
```

## 📊 命令类型与模板

### 1. 分析型命令

**示例：代码质量分析**
```markdown
# /quality-check.md
请对当前代码进行全面的质量检查：

1. **代码规范**
   - 检查命名规范
   - 验证代码格式
   - 识别代码异味

2. **潜在问题**
   - 性能瓶颈
   - 安全隐患
   - 逻辑错误

3. **改进建议**
   - 按优先级排序
   - 提供具体示例
   - 估算改进成本

输出格式：Markdown 报告
```

### 2. 生成型命令

**示例：测试生成**
```markdown
# /generate-tests.md
---
arguments: optional
---
为 $ARGUMENTS 生成完整的测试套件：

要求：
1. 单元测试覆盖所有公共方法
2. 包含正常和边界情况
3. 测试异常处理
4. 使用项目现有的测试框架

测试应该：
- 可读性强
- 独立运行
- 有意义的测试名称
- 包含断言说明
```

### 3. 工作流型命令

**示例：功能开发流程**
```markdown
# /feature.md
---
arguments: required
---
实现新功能：$ARGUMENTS

工作流程：
1. 创建功能分支
2. 分析需求，设计方案
3. 使用 TodoWrite 创建任务列表
4. 实现核心功能
5. 编写测试
6. 更新文档
7. 运行质量检查
8. 准备 PR

确保遵循项目的开发规范。
```

### 4. 集成型命令

**示例：Git 工作流**
```markdown
# /smart-commit.md
---
command: |
  git diff --staged
  git status
---
基于上述 Git 状态：

1. 分析暂存的更改
2. 生成符合规范的提交信息
3. 检查是否有遗漏的文件
4. 提醒潜在的问题

提交信息格式：
type(scope): subject

其中 type 包括：
- feat: 新功能
- fix: 修复
- docs: 文档
- refactor: 重构
- test: 测试
- chore: 构建/工具
```

## 💡 高级技巧

### 1. 条件逻辑
```markdown
请检查项目类型：
- 如果是 Node.js 项目，运行 npm test
- 如果是 Go 项目，运行 go test
- 如果是 Python 项目，运行 pytest
```

### 2. 多步骤协调
```markdown
# /release.md
发布新版本：

阶段 1 - 准备
□ 运行所有测试
□ 更新版本号
□ 生成 CHANGELOG

阶段 2 - 构建
□ 构建生产版本
□ 运行集成测试
□ 生成文档

阶段 3 - 发布
□ 创建 Git 标签
□ 推送到仓库
□ 触发 CI/CD
```

### 3. 错误处理
```markdown
执行以下操作，如果遇到错误：
1. 记录错误详情
2. 尝试自动修复
3. 如果无法修复，提供解决方案
4. 询问用户如何处理
```

## 📁 命令组织最佳实践

### 1. 命名规范
```
✅ 好的命名：
- analyze-performance.md
- generate-api-docs.md
- setup-dev-env.md

❌ 避免：
- cmd1.md
- test.md
- misc.md
```

### 2. 目录结构
```
~/.claude/commands/
├── dev/              # 开发相关
│   ├── analyze.md
│   └── refactor.md
├── test/             # 测试相关
│   ├── unit.md
│   └── e2e.md
├── ops/              # 运维相关
│   ├── deploy.md
│   └── monitor.md
└── workflow/         # 工作流
    ├── feature.md
    └── hotfix.md
```

### 3. 文档化
每个命令都应该包含：
- 用途说明
- 参数说明（如果有）
- 示例用法
- 预期输出

## 🔧 命令测试与优化

### 1. 测试命令
```bash
# 创建测试命令
echo "测试命令内容" > .claude/commands/test-cmd.md

# 使用命令
/test-cmd

# 根据结果优化
```

### 2. 迭代优化
- 记录命令使用情况
- 收集执行结果
- 根据反馈调整
- 版本化管理

### 3. 性能考虑
- 避免过于复杂的命令
- 合理使用 bash 集成
- 考虑执行时间
- 优化输出格式

## 📚 实用命令示例集

### 日常开发
```markdown
# /daily-standup.md
生成今日站会报告：
1. 昨天完成了什么
2. 今天计划做什么
3. 遇到的阻碍
基于 git log 和 todo list
```

### 代码审查
```markdown
# /review-pr.md
---
arguments: required
---
审查 PR #$ARGUMENTS：
- 代码质量
- 测试覆盖
- 文档更新
- 性能影响
- 安全考虑
```

### 问题诊断
```markdown
# /diagnose.md
---
arguments: optional
---
诊断问题：${ARGUMENTS:-最近的错误}
1. 收集相关日志
2. 分析错误模式
3. 追踪根本原因
4. 提供解决方案
```

## ❓ 常见问题

**Q: 命令应该多详细？**
A: 足够详细以确保一致的执行，但不要过度规定。

**Q: 如何处理复杂逻辑？**
A: 拆分成多个命令，或使用 Task 工具处理复杂流程。

**Q: 命令可以调用其他命令吗？**
A: 不直接支持，但可以在命令中提示使用其他命令。

## 🚀 下一步

1. 从简单命令开始
2. 逐步增加复杂度
3. 建立个人命令库
4. 与团队分享优秀命令

---

*编写好的命令是一门艺术，需要不断实践和优化。记住：命令是为了提高效率，而不是增加复杂性。*