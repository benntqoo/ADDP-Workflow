---
arguments: optional
format: "[status|create|plan|validate|rollback|迁移名]"
examples:
  - "/migrate status - 检查迁移状态"
  - "/migrate create add_user_table - 创建新迁移"
  - "/migrate plan - 生成迁移执行计划"
  - "/migrate validate - 验证待执行迁移"
  - "/migrate rollback 3 - 回滚最近3个迁移"
command: |
  echo "检查数据库迁移状态..."
  # 检测常见的迁移工具
  if [ -f "migrations" ] || [ -f "db/migrate" ] || [ -f "database/migrations" ]; then
    echo "发现迁移文件目录"
  fi
---
数据库迁移管理工具：

## 迁移分析

1. **当前状态检查**
   - 已应用的迁移列表
   - 待执行的迁移文件
   - 数据库架构版本
   - 迁移历史追踪

2. **迁移文件分析**
   - 结构变更识别（DDL）
   - 数据变更检测（DML）
   - 破坏性变更警告
   - 依赖关系验证

3. **风险评估**
   - 数据丢失风险
   - 性能影响预测
   - 锁表时间估算
   - 回滚复杂度

## 迁移操作

### create - 创建新迁移
```sql
-- 生成迁移模板
-- 包含 up() 和 down() 方法
-- 自动命名和时间戳
```

### plan - 迁移执行计划
- 执行顺序优化
- 批量迁移策略
- 零停机方案
- 备份检查点

### validate - 迁移验证
- SQL 语法检查
- 约束冲突检测
- 数据完整性验证
- 性能影响分析

### rollback - 回滚方案
- 生成回滚脚本
- 数据恢复策略
- 状态还原验证
- 应急处理流程

## 最佳实践检查

- ✅ 每个迁移都有对应的回滚
- ✅ 避免在迁移中修改大量数据
- ✅ 使用事务保证原子性
- ✅ 添加适当的索引
- ✅ 保留历史数据备份

## 文档生成
- 迁移日志 → docs/releases/migrations/
- 数据库架构 → docs/architecture/database/
- 操作手册 → docs/guides/database/