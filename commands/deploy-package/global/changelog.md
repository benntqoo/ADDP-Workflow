---
command: |
  git log --oneline -10
  echo "---"
  git diff --stat HEAD~1
---
基于最近的提交历史，生成或更新 CHANGELOG：

1. **分析提交类型**
   - feat: 新功能
   - fix: 错误修复
   - docs: 文档更新
   - style: 代码格式
   - refactor: 重构
   - perf: 性能优化
   - test: 测试相关
   - chore: 构建/工具

2. **生成版本记录**
   - 遵循语义化版本（Semantic Versioning）
   - 按类型分组变更
   - 突出破坏性变更
   - 标注贡献者

3. **格式规范**
   - 遵循 Keep a Changelog 格式
   - 支持多语言版本
   - 包含发布日期
   - 链接相关 Issue/PR