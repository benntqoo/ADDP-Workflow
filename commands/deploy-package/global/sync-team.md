同步团队知识到 Graphiti 知识图谱：

## 知识收集与分类

1. **项目知识提取**
   - 架构决策记录（ADRs）
   - API 设计规范
   - 编码最佳实践
   - 故障处理经验

2. **自动识别知识类型**
   - **Preference**：团队偏好和规范
   - **Procedure**：标准操作流程
   - **Fact**：项目事实和关系

3. **知识同步到 Graphiti**
   ```yaml
   使用 mcp__graphiti-memory__add_memory:
     - 架构决策 → source: "text"
     - API 规范 → source: "json"
     - 操作流程 → source: "message"
   ```

## 同步内容

### 技术规范
- 编码标准和约定
- 技术栈选择理由
- 性能优化策略
- 安全实践指南

### 项目知识
- 业务领域术语
- 核心功能说明
- 集成点文档
- 部署架构

### 团队实践
- 代码审查标准
- 测试策略
- 发布流程
- 故障响应程序

## 知识查询
同步后可通过以下方式查询：
- `mcp__graphiti-memory__search_memory_nodes`
- `mcp__graphiti-memory__search_memory_facts`

## 使用场景
- 新成员入职培训
- 技术决策参考
- 问题排查指南
- 知识传承保障