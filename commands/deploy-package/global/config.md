---
arguments: optional
format: "[环境名|check|init|sync]"
examples:
  - "/config dev - 管理开发环境配置"
  - "/config prod - 管理生产环境配置"
  - "/config check - 检查当前配置完整性"
  - "/config init - 初始化配置文件"
  - "/config sync - 同步配置到团队"
---
环境配置管理工具：

## 🔧 配置扫描

### 1. 配置文件检测
```yaml
常见配置文件:
  - .env / .env.* (环境变量)
  - config/*.json (应用配置)
  - application.yml (Spring)
  - settings.py (Django)
  - package.json (Node.js)
```

### 2. 环境识别
- **开发环境** (development)
- **测试环境** (testing/staging)
- **生产环境** (production)
- **本地环境** (local)

### 3. 配置分析
- 必需配置项检查
- 默认值识别
- 敏感信息标记
- 环境差异对比

## 📋 配置管理

### 配置模板生成
```bash
# 生成 .env.example
DATABASE_URL=postgres://user:pass@host:5432/db
REDIS_URL=redis://localhost:6379
API_KEY=your-api-key-here
SECRET_KEY=generate-strong-secret
```

### 配置验证
- [ ] 所有必需变量已设置
- [ ] 数据库连接可用
- [ ] 外部服务可访问
- [ ] 密钥强度足够
- [ ] 无硬编码值

### 安全最佳实践
1. **密钥管理**
   - 使用密钥管理服务
   - 定期轮换密钥
   - 加密敏感配置
   - 审计访问日志

2. **环境隔离**
   - 不同环境不同密钥
   - 最小权限原则
   - 网络访问控制
   - 配置版本控制

## 🔄 配置同步

### 团队同步策略
```yaml
开发团队:
  - 共享非敏感配置
  - 密钥通过安全渠道
  - 配置变更通知
  - 文档保持更新
```

### 配置迁移
- 新环境配置初始化
- 配置项重命名处理
- 废弃配置清理
- 向后兼容维护

## 📊 配置报告

### 健康检查
- 配置完整性 ✓
- 安全合规性 ✓
- 性能影响评估
- 依赖服务状态

### 文档生成
1. **配置说明** → `docs/guides/configuration.md`
2. **环境差异** → `docs/references/env-comparison.md`
3. **迁移指南** → `docs/guides/config-migration.md`

## 🚨 常见问题

### 配置错误诊断
- 连接超时 → 检查网络和防火墙
- 认证失败 → 验证密钥和权限
- 找不到配置 → 确认文件路径
- 类型错误 → 检查值格式

### 故障排查步骤
1. 验证配置文件存在
2. 检查环境变量加载
3. 确认配置优先级
4. 查看错误日志
5. 使用调试模式