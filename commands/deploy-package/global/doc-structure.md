---
arguments: optional
---
管理项目文档目录结构 - $ARGUMENTS：

## 支持的操作

### init - 初始化文档结构
创建标准的文档目录结构：
```
docs/
├── api/          # API 文档
├── architecture/ # 架构文档
├── guides/       # 使用指南
├── development/  # 开发文档
├── references/   # 参考文档
├── releases/     # 发布相关
└── internal/     # 内部文档
```

### check - 检查文档规范性
1. 验证目录结构是否符合规范
2. 检查文档命名是否正确
3. 识别散落在项目中的文档
4. 生成整改建议

### organize - 整理现有文档
1. 扫描项目中的所有文档
2. 根据类型自动归类
3. 修正不规范的命名
4. 生成文档索引

### 默认行为（无参数）
显示当前文档结构并提供改进建议

## 文档归类规则
- README.md → 项目根目录
- API 文档 → docs/api/
- 架构图/设计文档 → docs/architecture/
- 安装/配置指南 → docs/guides/
- 开发规范 → docs/development/
- CHANGELOG → docs/releases/