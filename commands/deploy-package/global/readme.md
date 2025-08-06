---
arguments: optional
format: "[项目路径|update|template|check]"
examples:
  - "/readme - 为当前项目生成 README"
  - "/readme ./subproject - 为子项目生成 README"
  - "/readme update - 更新现有 README"
  - "/readme template - 生成 README 模板"
  - "/readme check - 检查 README 完整性"
---
README 文档生成器：

1. **项目信息提取**
   - 项目名称和描述
   - 主要功能特性
   - 技术栈和依赖
   - 许可证信息

2. **标准章节生成**
   - 项目简介
   - 快速开始指南
   - 安装步骤
   - 使用示例
   - API 参考（如适用）
   - 配置说明
   - 贡献指南
   - 更新日志链接

3. **智能内容**
   - 徽章生成（CI/CD、覆盖率等）
   - 自动检测构建命令
   - 环境变量说明
   - 常见问题解答

4. **本地化支持**
   - 多语言版本
   - 区域特定内容
   - 文档链接管理