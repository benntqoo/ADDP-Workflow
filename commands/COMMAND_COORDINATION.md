# 命令间协调机制

本文档定义了 Claude Code 命令系统的协调规则，确保命令间的顺畅配合。

## 🔗 命令依赖关系

### 项目接入流程
```mermaid
graph LR
    audit[/audit] --> onboard[/onboard]
    onboard --> constitution[/constitution]
    constitution --> doc-structure[/doc-structure init]
```

### 代码分析流程
```mermaid
graph LR
    explore[/explore] --> analyze[/analyze]
    analyze --> discover[/discover]
    discover --> test[/test]
```

### 文档管理流程
```mermaid
graph LR
    doc-structure[/doc-structure] --> doc-api[/doc-api]
    doc-api --> doc-sync[/doc-sync]
    doc-sync --> readme[/readme]
```

## 📋 命令协调规则

### 1. 状态共享机制

#### 共享状态文件
```yaml
# .claude/state/command-state.yml
last_commands:
  - command: "/audit"
    timestamp: "2024-01-15T10:00:00Z"
    output: "health-report.json"
    
  - command: "/analyze"
    timestamp: "2024-01-15T10:15:00Z"
    output: "tech-stack.json"

shared_data:
  project_type: "nodejs"
  framework: "express"
  test_framework: "jest"
  documentation_standard: "docs/"
```

#### 状态读写规则
- 每个命令执行后更新状态
- 后续命令可读取前置命令结果
- 状态文件使用锁机制防止冲突

### 2. 命令链式调用

#### 自动触发机制
```yaml
command_chains:
  onboard_complete:
    triggers:
      - "/doc-structure init"
      - "/ai-rules apply"
    
  test_failure:
    triggers:
      - "/coverage --detail"
      - "/guardian report"
    
  deploy_check_pass:
    triggers:
      - "/changelog --update"
      - "/sync-team"
```

#### 条件触发
```yaml
conditional_triggers:
  - when: "audit.score < 'C'"
    trigger: "/retrofit --plan"
    
  - when: "coverage.percentage < 80"
    trigger: "/test --generate-missing"
    
  - when: "perf.issues > 0"
    trigger: "/perf --deep-analysis"
```

### 3. 数据传递协议

#### 标准输出格式
```json
{
  "command": "/analyze",
  "status": "success",
  "data": {
    "tech_stack": ["nodejs", "typescript", "react"],
    "dependencies": ["express", "jest", "eslint"],
    "architecture": "mvc"
  },
  "next_commands": ["/test", "/deps"],
  "warnings": []
}
```

#### 数据使用示例
- `/test` 读取 `/analyze` 的 test_framework
- `/doc-api` 使用 `/discover` 的 API 端点信息
- `/migrate` 参考 `/analyze` 的数据库类型

### 4. 冲突避免机制

#### 文件锁定
```yaml
file_locks:
  CLAUDE.md:
    locked_by: ["/meta", "/onboard", "/update-constitution"]
    lock_type: "exclusive"
    
  "docs/**":
    locked_by: ["/doc-*"]
    lock_type: "shared"
```

#### 操作优先级
```yaml
priority_rules:
  high:
    - "/rollback"  # 紧急回滚最高优先级
    - "/deploy-check"  # 部署检查优先
    
  medium:
    - "/test"
    - "/coverage"
    
  low:
    - "/doc-sync"
    - "/changelog"
```

## 🎯 智能推荐系统

### 基于上下文的命令推荐
```yaml
context_recommendations:
  after_file_edit:
    - "/test [edited_file]"
    - "/guardian status"
    
  after_new_feature:
    - "/test --generate"
    - "/doc-api [feature]"
    - "/changelog --add-feature"
    
  after_bug_fix:
    - "/test --regression"
    - "/coverage"
    - "/changelog --add-fix"
```

### 工作流模板
```yaml
workflow_templates:
  feature_development:
    steps:
      1: "/discover [feature_area]"
      2: "implement_feature"
      3: "/test"
      4: "/coverage"
      5: "/doc-api"
      6: "/deploy-check"
    
  hotfix:
    steps:
      1: "/discover [bug_area]"
      2: "fix_bug"
      3: "/test --regression"
      4: "/rollback --prepare"
      5: "/deploy-check --quick"
```

## 🔧 实现细节

### 1. 状态管理器
```typescript
interface CommandState {
  lastCommand: string;
  timestamp: Date;
  output: any;
  nextCommands: string[];
}

class StateManager {
  static save(state: CommandState): void
  static load(command: string): CommandState
  static clean(): void
}
```

### 2. 协调器接口
```typescript
interface CommandCoordinator {
  canExecute(command: string): boolean
  getPrerequisites(command: string): string[]
  getRecommendations(context: Context): string[]
  resolveConflicts(commands: string[]): string[]
}
```

### 3. 钩子系统
```typescript
interface CommandHooks {
  beforeExecute?: (command: string) => void
  afterExecute?: (command: string, result: any) => void
  onError?: (command: string, error: Error) => void
  onConflict?: (commands: string[]) => string
}
```

## 📊 监控和报告

### 命令使用统计
- 最常用的命令组合
- 平均执行时间
- 成功/失败率
- 用户偏好模式

### 优化建议
- 基于使用模式推荐工作流
- 识别低效的命令序列
- 建议更好的命令组合

## 🚀 未来增强

1. **机器学习优化**
   - 学习用户使用模式
   - 预测下一个命令
   - 自动优化工作流

2. **并行执行**
   - 识别可并行的命令
   - 自动并行化执行
   - 结果自动合并

3. **智能回滚**
   - 命令操作历史
   - 批量撤销功能
   - 状态快照恢复

---

*此协调机制确保命令系统的高效运作，随着使用不断优化。*