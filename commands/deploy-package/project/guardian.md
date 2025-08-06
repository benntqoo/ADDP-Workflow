---
arguments: optional
format: "[on|off|status|config|report]"
examples:
  - "/guardian on - 启用守护者模式"
  - "/guardian off - 关闭守护者模式"
  - "/guardian status - 查看当前状态"
  - "/guardian config - 配置监控规则"
  - "/guardian report - 生成质量报告"
---

代码质量守护者：

## 🛡️ 守护者模式

### 实时监控功能
当守护者模式启用时，Claude 会主动：

1. **安全检查**
   - 🔴 硬编码密钥和敏感信息
   - 🔴 SQL 注入风险
   - 🔴 XSS 和 CSRF 漏洞
   - 🟡 不安全的依赖版本
   - 🟡 权限配置问题

2. **代码质量**
   - 🔍 复杂度过高的函数
   - 🔍 重复代码检测
   - 🔍 未使用的变量和导入
   - 🔍 不一致的命名风格
   - 🔍 缺失的错误处理

3. **性能优化**
   - ⚡ N+1 查询问题
   - ⚡ 不必要的重渲染
   - ⚡ 内存泄漏风险
   - ⚡ 阻塞操作识别
   - ⚡ 大文件和包体积

4. **最佳实践**
   - ✅ 测试覆盖率不足
   - ✅ 缺失的文档
   - ✅ 不符合项目规范
   - ✅ 缺少类型定义
   - ✅ 违反 SOLID 原则

## 🎯 主动建议

### 即时反馈
```yaml
示例提醒:
  - "检测到硬编码的 API 密钥，建议使用环境变量"
  - "这个函数复杂度为 15，建议拆分为更小的函数"
  - "发现潜在的 N+1 查询，考虑使用 includes/join"
  - "缺少对 null/undefined 的检查"
```

### 改进建议
- 提供具体的代码示例
- 推荐相关的最佳实践
- 链接到项目规范文档
- 生成重构任务列表

## ⚙️ 配置选项

### 监控级别
```yaml
guardian_config:
  level: "strict"  # strict | normal | relaxed
  
  rules:
    security: 
      enabled: true
      severity: "error"
    
    performance:
      enabled: true
      severity: "warning"
    
    code_quality:
      enabled: true
      complexity_threshold: 10
      duplication_threshold: 50
    
    testing:
      enabled: true
      coverage_threshold: 80
```

### 自定义规则
```yaml
custom_rules:
  - name: "项目特定命名"
    pattern: "^[A-Z][a-zA-Z0-9]*$"
    apply_to: ["class", "interface"]
    
  - name: "必需的文件头"
    require: "copyright"
    file_types: [".ts", ".js"]
```

## 📊 质量报告

### 报告内容
1. **问题汇总**
   - 按严重程度分类
   - 按类型统计
   - 趋势分析

2. **具体问题列表**
   - 文件位置
   - 问题描述
   - 修复建议
   - 优先级

3. **改进建议**
   - 快速修复项
   - 长期改进项
   - 架构优化建议

### 输出格式
- 控制台摘要
- 详细报告 → `docs/internal/guardian-report.md`
- 任务列表 → TodoWrite
- CI/CD 集成 → JSON 格式

## 🔄 与其他命令协作

- **触发测试**: 发现问题后自动运行 `/test`
- **性能分析**: 检测到性能问题时调用 `/perf`
- **规范检查**: 配合 `/ai-rules` 验证合规性
- **文档更新**: 提醒使用 `/doc-sync` 更新文档

## 💡 使用建议

### 开发阶段
```bash
/guardian on
# 开启实时监控，获得即时反馈
```

### 代码审查
```bash
/guardian report
# 生成完整的质量报告
```

### CI/CD 集成
```bash
/guardian validate --exit-on-error
# 作为质量门禁使用
```

### 定制监控
```bash
/guardian config --rule security:strict
# 针对特定需求调整规则
```

## 🚨 智能提醒

守护者不会：
- 打断正常的开发流程
- 对每个小问题都发出警告
- 重复提醒已知问题

守护者会：
- 优先提醒严重问题
- 批量汇总相似问题
- 提供可操作的建议
- 学习项目偏好