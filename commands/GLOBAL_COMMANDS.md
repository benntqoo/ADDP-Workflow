# Claude Code 全局自定义命令集

本文档包含可以放置在 `~/.claude/commands/` 目录下的全局命令，这些命令在所有项目中都可以使用。

## 🚀 安装方法

### Windows
```bash
# 创建全局命令目录
mkdir -p %USERPROFILE%\.claude\commands

# 复制命令文件到该目录
```

### macOS/Linux
```bash
# 创建全局命令目录
mkdir -p ~/.claude/commands

# 复制命令文件到该目录
```

## 📝 全局命令示例

### 1. `/meta` - 启动元工作流
**文件**: `~/.claude/commands/meta.md`
```markdown
我想定制项目的协作规范，让我们开始元工作流。请帮我生成适合这个项目的 CLAUDE.md 配置。
```

### 2. `/analyze` - 深度分析代码库
**文件**: `~/.claude/commands/analyze.md`
```markdown
请对当前代码库进行深度分析：
1. 识别项目的技术栈和架构模式
2. 找出核心模块和关键路径
3. 分析代码质量和潜在问题
4. 生成项目概览报告
```

### 3. `/refactor` - 代码重构助手
**文件**: `~/.claude/commands/refactor.md`
```markdown
我需要重构代码。请：
1. 分析当前代码的问题
2. 提出重构方案
3. 评估重构风险
4. 使用 TodoWrite 创建重构任务列表
```

### 4. `/test` - 智能测试生成
**文件**: `~/.claude/commands/test.md`
```markdown
为当前代码生成测试：
1. 分析代码逻辑
2. 识别边界条件
3. 生成单元测试
4. 包含正常和异常场景
```

### 5. `/security` - 安全审查
**文件**: `~/.claude/commands/security.md`
```markdown
执行安全审查：
1. 检查潜在的安全漏洞
2. 验证输入验证
3. 检查敏感信息泄露
4. 提供修复建议
```

## 🏗️ 遗留项目专用命令

### `/onboard` - 遗留项目接入向导
**文件**: `~/.claude/commands/onboard.md`
```markdown
我需要接入一个现有项目。请执行以下步骤：

1. 扫描项目结构和技术栈
2. 分析代码规模和质量
3. 识别现有的编码规范
4. 生成项目健康度报告
5. 创建适配性的 CLAUDE.md

注：这是遗留项目，请保持谦逊和谨慎。
```

### `/audit` - 项目健康度审计
**文件**: `~/.claude/commands/audit.md`
```markdown
对这个遗留项目进行全面健康检查：
- 代码质量评估
- 技术债务识别
- 安全风险扫描
- 性能瓶颈分析
- 依赖过时检查

生成可操作的改进建议。
```

### `/discover` - 功能模块探索
**文件**: `~/.claude/commands/discover.md`
```markdown
---
arguments: optional
---
帮我深入理解：$ARGUMENTS

追踪执行流程，识别：
- 业务逻辑
- 数据流向
- 依赖关系
- 潜在风险
```

### `/retrofit` - 渐进式改造
**文件**: `~/.claude/commands/retrofit.md`
```markdown
选择一个小模块进行改进：
1. 保持向后兼容
2. 添加测试覆盖
3. 逐步现代化
4. 最小化风险
```

## 🎯 带参数的命令示例

### 6. `/explain` - 代码解释器
**文件**: `~/.claude/commands/explain.md`
```markdown
---
arguments: optional
---
请详细解释这段代码：
$ARGUMENTS

包括：
1. 代码的功能和目的
2. 实现原理
3. 使用的设计模式
4. 潜在的改进点
```

### 7. `/optimize` - 性能优化
**文件**: `~/.claude/commands/optimize.md`
```markdown
---
arguments: optional
---
优化以下代码的性能：
$ARGUMENTS

考虑：
1. 时间复杂度
2. 空间复杂度
3. 并发优化
4. 缓存策略
```

## 🔧 带 Bash 命令的示例

### 8. `/setup` - 项目初始化
**文件**: `~/.claude/commands/setup.md`
```markdown
---
command: |
  echo "检查项目类型..."
  if [ -f "package.json" ]; then
    echo "Node.js 项目检测到"
  elif [ -f "go.mod" ]; then
    echo "Go 项目检测到"
  elif [ -f "requirements.txt" ]; then
    echo "Python 项目检测到"
  fi
---
基于检测到的项目类型，请：
1. 安装必要的依赖
2. 设置开发环境
3. 创建基础配置文件
4. 生成 .gitignore
```

### 9. `/commit` - 智能提交
**文件**: `~/.claude/commands/commit.md`
```markdown
---
command: git diff --staged
---
基于上述暂存的更改，请：
1. 生成符合约定的提交信息
2. 解释主要改动
3. 检查是否有遗漏的文件
```

## 📂 命名空间示例

### 工作流命令组
**目录**: `~/.claude/commands/workflow/`

- `/workflow/pr` - 创建 Pull Request
- `/workflow/review` - 代码审查
- `/workflow/deploy` - 部署检查清单

### AI 辅助命令组
**目录**: `~/.claude/commands/ai/`

- `/ai/prompt` - 优化 AI 提示词
- `/ai/context` - 生成上下文摘要
- `/ai/learn` - 学习新技术

## 💡 高级用法

### 10. `/template` - 代码模板生成器
**文件**: `~/.claude/commands/template.md`
```markdown
---
arguments: required
name: "Code Template Generator"
description: "Generate code from templates"
---
生成 $ARGUMENTS 的代码模板。

支持的模板类型：
- component: React/Vue 组件
- api: REST API 端点
- test: 测试文件
- model: 数据模型
- service: 服务层代码
```

## 🎨 自定义命令最佳实践

1. **命名规范**
   - 使用简短、描述性的名称
   - 动词开头（analyze, create, check）
   - 避免与内置命令冲突

2. **内容结构**
   - 清晰的指令
   - 明确的输出期望
   - 分步骤的任务

3. **参数使用**
   - 明确标注是否需要参数
   - 提供参数使用示例
   - 验证参数有效性

4. **命令组织**
   - 相关命令放在同一目录
   - 使用命名空间避免冲突
   - 保持命令数量适中

## 🔄 维护建议

1. **定期更新**: 根据项目需求更新命令
2. **版本控制**: 将常用命令备份到 Git
3. **团队共享**: 导出优秀的命令供团队使用
4. **性能优化**: 避免命令过于复杂

---

*这些全局命令可以显著提升您与 Claude Code 的协作效率。根据您的具体需求定制和扩展这些命令。*