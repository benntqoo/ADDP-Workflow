# 部署包更新日志

## v2.1.0 (2024-01-15)

### 🎯 主要改进

#### 1. 参数规范化
- 为 9 个命令添加了明确的参数格式说明和使用示例
- 受影响命令：config, deps, perf, readme, doc-api, discover, microservice, migrate, rollback

#### 2. 职责边界明确化
- **项目分析类**：
  - explore: 快速浏览，生成项目地图（5-10分钟）
  - analyze: 技术栈分析，依赖检查（15-30分钟）
  - discover: 业务逻辑理解，数据流分析（30-60分钟）

- **项目接入类**：
  - audit: 遗留项目健康度评估
  - onboard: 智能接入向导，生成适配性 CLAUDE.md
  - constitution: 应用已存在的协作规范

#### 3. 项目命令增强
- **ai-rules**: 从简单检查升级为完整的规范管理器
  - 支持 check/apply/validate/report 操作
  - 集成 CLAUDE.md 版本管理
  - 提供合规性验证和报告

- **guardian**: 从被动提醒升级为主动监控系统
  - 实时代码质量监控
  - 安全、性能、质量多维度检查
  - 支持自定义规则和 CI/CD 集成

- **sync**: 从单向同步升级为双向智能同步
  - 支持上传/下载/冲突解决
  - 版本历史追踪和回滚
  - 团队配置共享机制

#### 4. 命令协调机制
- 创建了 `COMMAND_COORDINATION.md` 定义协调规则
- 实现了状态共享机制（`.claude/state/command-state.yml`）
- 定义了命令依赖关系和链式调用
- 添加了智能推荐系统

### 📋 文件变更

#### 新增文件
- `COMMAND_COORDINATION.md` - 命令协调机制文档
- `.claude/state/command-state.yml` - 命令状态示例
- `CHANGELOG.md` - 本更新日志

#### 更新文件
- 14 个全局命令文件（参数和描述优化）
- 3 个项目命令文件（功能增强）

#### 已删除文件
- `OPTIMIZATION_PLAN.md` - 优化完成后删除，避免混淆

### 🔧 技术细节

1. **参数格式标准化**
   ```markdown
   format: "[参数1|参数2|参数3]"
   examples:
     - "/命令 参数1 - 说明"
     - "/命令 参数2 - 说明"
   ```

2. **职责描述标准化**
   ```markdown
   描述: 一句话说明命令用途
   重点: 关注的核心内容
   速度/适用: 执行时间或适用场景
   ```

3. **状态共享格式**
   ```yaml
   last_commands:
     - command: "命令名"
       timestamp: "时间戳"
       output: "输出数据"
   ```

### 🚀 使用建议

1. 部署新版本命令后，建议运行 `/ai-rules apply` 加载最新规范
2. 使用 `/guardian on` 开启实时代码质量监控
3. 定期运行 `/sync status` 检查配置更新
4. 参考 `COMMAND_COORDINATION.md` 了解命令最佳组合

### 🐛 已知问题

- 命令状态文件需要手动创建 `.claude/state` 目录
- 部分命令的 bash 脚本可能需要根据平台调整

### 📅 下一版本计划

- 实现命令并行执行优化
- 添加命令执行历史可视化
- 完善错误处理和回滚机制
- 增加更多预定义工作流模板

---

*此版本由 Claude Code 协作生成*