---
arguments: optional
format: "[模块名|all|tree|check|clean]"
examples:
  - "/deps auth - 分析 auth 模块的依赖"
  - "/deps all - 分析所有模块的依赖"
  - "/deps tree - 显示完整依赖树"
  - "/deps check - 检查依赖问题（循环、冲突等）"
  - "/deps clean - 清理未使用的依赖"
---
模块依赖分析工具：

1. **依赖扫描**
   - 直接依赖 vs 间接依赖
   - 生产依赖 vs 开发依赖
   - 版本冲突检测
   - 循环依赖识别

2. **依赖分析**
   - 依赖树可视化（Mermaid 图）
   - 依赖深度分析
   - 包大小和加载时间影响
   - 安全漏洞扫描（已知 CVE）

3. **模块关系**
   - 内部模块依赖图
   - 模块耦合度分析
   - 接口依赖追踪
   - 共享代码识别

4. **优化建议**
   - 重复依赖合并
   - 过时依赖更新
   - 未使用依赖清理
   - 依赖替代方案

5. **生成报告**
   - 依赖健康度评分
   - 风险等级评估
   - 优化路线图
   - 更新到 docs/architecture/dependencies/