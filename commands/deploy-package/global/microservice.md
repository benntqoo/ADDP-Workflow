---
arguments: optional
format: "[analyze|contract|sync|test|monitor|服务名]"
examples:
  - "/microservice analyze - 分析微服务架构全貌"
  - "/microservice contract - 管理 API 契约"
  - "/microservice sync - 协调服务部署"
  - "/microservice test - 集成测试协调"
  - "/microservice monitor - 配置监控体系"
  - "/microservice auth-service - 分析特定服务"
---
微服务架构管理工具：

## 服务发现与分析

1. **服务清单**
   - 扫描所有微服务
   - 识别服务版本和状态
   - 检测服务间通信模式
   - API 契约验证

2. **依赖拓扑**
   - 服务调用链路图
   - 同步 vs 异步通信
   - 消息队列依赖
   - 数据库共享情况

3. **健康检查**
   - 服务可用性监控
   - 性能瓶颈识别
   - 单点故障分析
   - 级联失败风险

## 协调任务

### contract - API 契约管理
- 生成服务间契约文档
- 契约版本控制
- 破坏性变更检测
- 向后兼容性验证

### sync - 服务同步部署
- 依赖顺序分析
- 部署编排计划
- 回滚策略制定
- 配置同步检查

### test - 集成测试协调
- 端到端测试场景
- 服务模拟（Mock）
- 混沌工程测试
- 性能基准测试

### monitor - 监控配置
- 分布式追踪设置
- 日志聚合配置
- 指标收集规范
- 告警规则定义

## 输出位置
- 服务目录 → docs/architecture/services/
- API 契约 → docs/api/contracts/
- 部署文档 → docs/guides/deployment/