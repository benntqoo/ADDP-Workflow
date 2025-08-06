---
arguments: optional
format: "[模块名|all|openapi|graphql|postman]"
examples:
  - "/doc-api - 生成所有 API 文档"
  - "/doc-api auth - 生成认证模块 API 文档"
  - "/doc-api openapi - 生成 OpenAPI/Swagger 规范"
  - "/doc-api graphql - 生成 GraphQL Schema 文档"
  - "/doc-api postman - 导出 Postman Collection"
---
API 文档生成器：

1. **扫描 API 端点**
   - REST API 路由
   - GraphQL Schema
   - RPC 接口

2. **提取文档信息**
   - 请求/响应格式
   - 参数说明
   - 状态码定义
   - 示例数据

3. **生成文档格式**
   - OpenAPI/Swagger 规范
   - Markdown 文档
   - Postman Collection

4. **包含内容**
   - 认证方式
   - 错误处理
   - 版本信息
   - 使用示例