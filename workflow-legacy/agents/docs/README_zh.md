# 代理系统文档

*[English](README.md) | 中文*

## 概述

Claude Code 代理系统为各种开发任务提供专业的 AI 助手。这些代理在特定领域、语言和框架方面拥有深厚的专业知识，可根据上下文自动激活。

## 代理工作原理

1. **自动激活**：代理根据任务描述关键词自动激活
2. **上下文感知**：代理理解文件类型和代码模式
3. **协同工作**：多个代理可以共同处理复杂任务
4. **专业深度**：每个代理在其领域都有专门知识

## 代理分类

### 🔍 质量保证代理

#### `code-reviewer` 代码审查专家
- **专长**：代码质量、最佳实践、安全问题
- **激活场景**：代码审查请求、PR审查
- **能力**：
  - 识别bug和潜在问题
  - 提供改进建议
  - 检查编码标准
  - 安全漏洞检测

#### `test-automator` 测试自动化专家
- **专长**：测试生成、覆盖率分析
- **激活场景**：测试相关任务
- **能力**：
  - 生成单元测试
  - 创建集成测试
  - 设计测试场景
  - 覆盖率建议

#### `performance-optimizer` 性能优化专家
- **专长**：性能分析、优化
- **激活场景**：性能问题、优化请求
- **能力**：
  - 识别瓶颈
  - 提供优化建议
  - 内存使用分析
  - 算法改进

#### `bug-hunter` Bug猎手
- **专长**：调试、根因分析
- **激活场景**：Bug报告、错误信息
- **能力**：
  - 分析堆栈跟踪
  - 查找根本原因
  - 建议修复方案
  - 预防类似问题

### 💻 技术专家代理

#### 语言专家

**`python-ml-specialist` Python机器学习专家**
- **领域**：机器学习、数据科学
- **框架**：PyTorch、TensorFlow、scikit-learn
- **用例**：模型开发、数据分析、AI应用

**`typescript-fullstack-expert` TypeScript全栈专家**
- **领域**：全栈TypeScript开发
- **框架**：React、Node.js、Next.js
- **用例**：Web应用、API、类型安全开发

**`kotlin-expert` Kotlin专家**
- **领域**：跨平台Kotlin开发
- **框架**：Android、Spring Boot、Ktor
- **用例**：移动应用、后端服务

**`golang-systems-engineer` Go系统工程师**
- **领域**：系统编程、微服务
- **框架**：标准库、流行Go包
- **用例**：高性能服务、云原生应用

**`rust-zero-cost` Rust零成本专家**
- **领域**：系统编程、性能关键代码
- **框架**：Tokio、Actix、标准库
- **用例**：系统工具、Web服务、嵌入式

#### 框架专家

**`android-kotlin-architect` Android架构师**
- **领域**：Android应用开发
- **技术**：Jetpack Compose、协程、架构组件
- **用例**：移动应用、UI开发

**`ktor-backend-architect` Ktor后端架构师**
- **领域**：Kotlin后端服务
- **技术**：Ktor框架、协程
- **用例**：REST API、微服务

### 🎭 工作流代理

#### `work-coordinator` 工作协调员
- **用途**：协调多个代理处理复杂任务
- **能力**：
  - 任务分配
  - 结果聚合
  - 依赖管理
  - 跨领域协调

### 🛡️ 验证代理

#### `jenny-validator` Jenny验证器
- **用途**：验证规范和需求
- **能力**：
  - 需求验证
  - 规范合规性
  - 一致性检查

#### `karen-realist` Karen现实主义者
- **用途**：现实检查和可行性评估
- **能力**：
  - 时间线验证
  - 资源评估
  - 风险识别

## 使用代理

### 自动激活

代理根据你的任务自动激活：

```bash
# 这会自动激活 python-ml-specialist
"帮我构建一个文本分类模型"

# 这会自动激活 security-analyst
"审查这个认证代码的安全漏洞"

# 这会自动激活 android-kotlin-architect
"为Android设计购物车功能"
```

### 手动激活

你可以明确请求特定代理：

```bash
"使用 rust-zero-cost 代理优化这段代码"
"让 test-automator 生成全面的测试"
```

### 代理协作

对于复杂任务，多个代理会协同工作：

```bash
"构建一个安全的支付处理系统"
# 激活：
# - security-analyst（安全审查）
# - code-reviewer（代码质量）
# - test-automator（测试生成）
# - performance-optimizer（优化）
```

## 最佳实践

1. **信任代理专业知识**：代理拥有深厚的领域知识
2. **让代理自动激活**：不要过度管理代理选择
3. **使用清晰的任务描述**：更好的描述 = 更好的代理匹配
4. **结合代理与命令**：使用命令控制工作流，使用代理提供专业知识
5. **审查代理建议**：代理提供建议，你做决定

## 代理能力矩阵

| 代理 | 语言 | 框架 | 最适合 |
|------|------|------|--------|
| python-ml-specialist | Python | PyTorch、TensorFlow | ML/AI开发 |
| typescript-fullstack-expert | TypeScript/JavaScript | React、Node.js | Web开发 |
| kotlin-expert | Kotlin | Android、Spring、Ktor | 移动和后端 |
| golang-systems-engineer | Go | 标准库 | 微服务 |
| rust-zero-cost | Rust | Tokio、Actix | 系统编程 |
| android-kotlin-architect | Kotlin | Android SDK | Android应用 |
| code-reviewer | 所有 | 所有 | 代码质量 |
| test-automator | 所有 | 测试框架 | 测试生成 |
| performance-optimizer | 所有 | 所有 | 性能调优 |

## 创建自定义代理

要创建自定义代理，在 `~/.claude/agents/` 中创建文件：

```markdown
---
description: 用于自动激活的简短描述
---

# 代理名称

你是[领域]的专家。你的专长包括...

## 能力
- [能力1]
- [能力2]

## 激活关键词
- [关键词1]
- [关键词2]

## 响应风格
[代理应该如何沟通]
```

## 故障排除

**代理没有激活？**
- 在请求中使用更具体的关键词
- 手动指定代理名称
- 检查代理文件是否存在于 `~/.claude/agents/`

**错误的代理被激活？**
- 在任务描述中更具体
- 通过明确的代理选择手动覆盖

**代理给出通用响应？**
- 确保代理文件格式正确
- 检查代理是否有明确的专业知识定义

---

更多示例请参见[代理使用指南](../../guides/AGENT_GUIDE.md)