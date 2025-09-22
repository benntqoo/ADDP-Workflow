# 🎯 Universal AI Coding Framework 使用示例

本文档提供了完整的使用示例，展示如何在不同场景下使用 Universal AI Coding Framework。

## 📋 快速开始示例

### 场景1: 新项目开发

```bash
# 1. 初始化项目结构
claude "初始化 ADDP 项目结构"

# 2. 创建需求规格
claude "/specify 开发一个用户管理系统，包含注册、登录、个人资料管理功能"

# 3. 生成技术方案
claude "/plan"

# 4. 分解开发任务
claude "/tasks"

# 5. 启动 ADDP 工作流
claude "/workflow analysis"
```

### 场景2: 性能优化

```bash
# 1. 优化查询
claude "优化React应用性能，目标是减少首屏加载时间50%"

# 2. 分析现有架构
claude "/workflow analysis 当前React应用加载缓慢，需要性能优化"

# 3. 设计优化方案
claude "/workflow design"

# 4. 实施优化
claude "/workflow development"
```

### 场景3: 代码重构

```bash
# 1. 分析重构需求
gemini "分析这个组件的重构需求：降低复杂度，提高可维护性"

# 2. 生成重构方案
gemini "/plan 重构用户注册组件"

# 3. 启动重构工作流
gemini "/workflow design 组件重构方案"
```

## 🔄 跨工具协作示例

### Claude Code → Gemini CLI 切换

```bash
# 在 Claude Code 中开始项目
claude "开始开发博客系统的用户认证模块"
claude "/specify 用户认证功能需求"

# 保存状态
claude "保存当前项目状态到 gemini"

# 切换到 Gemini CLI 继续
gemini "加载项目状态"
gemini "继续开发用户认证模块"
```

### Cursor → Claude Code 协作

```bash
# 在 Cursor 中进行 UI 开发
# 通过 MCP 同步状态

# 在 Claude Code 中继续后端开发
claude "同步 Cursor 的前端开发进度"
claude "开发对应的后端 API"
```

## 🏗️ 项目类型示例

### React 前端项目

```bash
# 自动检测 React 项目
claude "初始化 ADDP 项目结构"  # 自动检测 package.json 中的 react 依赖

# React 特定的规格开发
claude "/specify 实现 React 购物车组件，支持添加、删除、修改数量功能"

# 优化 React 性能
claude "优化 React 组件渲染性能，减少不必要的重渲染"
```

### Python 后端项目

```bash
# Django/Flask 项目初始化
claude "为 Python Flask 项目初始化 ADDP 结构"

# Python 特定开发流程
claude "/specify 开发 Flask RESTful API，包含用户认证和数据管理"

# Python 代码质量优化
claude "优化 Python 代码：提高性能，遵循 PEP8，增强类型提示"
```

### 全栈项目

```bash
# 全栈项目初始化
claude "初始化全栈项目：React前端 + Node.js后端 + MongoDB数据库"

# 前后端协调开发
claude "/workflow analysis 全栈项目架构设计"
claude "/workflow design 前后端 API 接口设计"
```

## ⚡ ADDP 工作流详细示例

### Analysis 阶段

```bash
# 启动分析阶段
claude "/workflow analysis"

# 输入需求
"""
需求：开发在线教育平台的视频播放功能
约束：支持多种视频格式，需要进度记录，支持倍速播放
目标：提供流畅的视频学习体验
"""

# 分析输出会包含：
# - 需求澄清结果
# - 技术约束清单
# - 风险评估报告
# - 影响分析报告
# - 资源需求评估
```

### Design 阶段

```bash
# 启动设计阶段
claude "/workflow design"

# 基于分析结果输入：
"""
基于分析阶段的输出，设计视频播放系统：
- 前端：React视频播放器组件
- 后端：视频流处理服务
- 存储：视频文件和进度数据
"""

# 设计输出会包含：
# - 系统架构设计
# - 技术方案决策
# - 接口规格说明
# - 数据模型设计
# - 详细实施计划
```

### Development 阶段

```bash
# 启动开发阶段 (TDD 驱动)
claude "/workflow development"

# 开发流程：
# 1. 编写测试用例
# 2. 实现最小功能
# 3. 运行测试
# 4. 重构代码
# 5. 重复循环

# 开发输出会包含：
# - TDD 测试用例
# - 功能实现代码
# - 测试执行结果
# - 代码质量报告
```

### Persistence 阶段

```bash
# 启动持久化阶段
claude "/workflow persistence"

# 持久化流程：
# 1. 功能验证确认
# 2. 性能指标检查
# 3. 项目记忆更新
# 4. 经验教训记录
# 5. 状态同步执行

# 持久化输出会包含：
# - 验证结果报告
# - 性能指标数据
# - 更新的项目记忆
# - 同步状态信息
```

## 🔍 查询优化示例

### 基础查询 → 优化查询

```bash
# 原始模糊查询
claude "优化性能"

# 优化后的精确查询 (系统自动生成)
"""
使用 React.memo、useMemo 和 useCallback 优化 React 组件性能：
1. 识别渲染性能瓶颈组件
2. 实施 React.memo 包装纯组件
3. 使用 useMemo 缓存计算结果
4. 使用 useCallback 稳定函数引用
5. 设置性能监控和基准测试
目标：减少 50% 不必要的组件重渲染
"""
```

### 上下文感知优化

```bash
# 带上下文的查询
claude "在 TypeScript React 项目中实现状态管理"

# 系统会基于项目上下文优化：
"""
在 TypeScript React 项目中实现 Redux Toolkit 状态管理：
1. 安装 @reduxjs/toolkit 和 react-redux
2. 创建 store.ts 配置文件
3. 定义 TypeScript 类型的 slice
4. 实现 Provider 包装应用
5. 在组件中使用 typed hooks (useAppSelector, useAppDispatch)
考虑因素：TypeScript 类型安全、DevTools 集成、中间件配置
"""
```

## 📊 分析和报告示例

### 查询优化分析

```bash
# 查看优化统计
claude "显示查询优化分析报告"

# 输出示例：
"""
📊 查询优化统计报告
- 总优化次数: 156
- 缓存命中率: 23.1%
- 平均置信度: 0.87
- 优化级别分布:
  * 基础优化: 45%
  * 智能优化: 35%
  * 详细优化: 20%
"""
```

### 工作流分析

```bash
# 查看工作流状态
claude "显示 ADDP 工作流状态"

# 输出示例：
"""
⚡ ADDP 工作流状态
- 当前阶段: development
- 已完成阶段: [analysis, design]
- 下一阶段: persistence
- 工作流健康度: 良好
- 最后更新: 2024-01-26 14:30:25
"""
```

## 🔧 高级配置示例

### 自定义模型配置

```json
{
  "ollama": {
    "endpoint": "http://localhost:11434",
    "model": "qwen2.5:14b",
    "temperature": 0.3,
    "max_tokens": 4096,
    "timeout": 60
  }
}
```

### 项目特定配置

```json
{
  "project": {
    "addp_directory": ".addp",
    "auto_initialize": true,
    "quality_gates_enabled": true,
    "cross_tool_sync": true,
    "analytics_enabled": true,
    "custom_templates": {
      "prd_template": "custom_prd.md",
      "adr_template": "custom_adr.md"
    }
  }
}
```

## 🚨 错误处理示例

### 常见错误和解决方案

```bash
# Ollama 连接失败
claude "优化查询: 实现登录功能"
# 错误: "Ollama API 错误: 连接拒绝"
# 解决: ollama serve

# MCP 工具不可用
claude "初始化项目结构"
# 错误: "MCP 工具调用失败"
# 解决: 检查 MCP 服务器配置

# 权限不足
claude "初始化 ADDP 结构"
# 错误: "无法创建目录"
# 解决: 检查目录写权限
```

## 🎯 最佳实践示例

### 规格驱动开发

```bash
# 1. 需求规格化
claude "/specify 电商网站的商品搜索功能：支持分类筛选、价格排序、关键词搜索"

# 2. 技术方案规划
claude "/plan"

# 3. 任务分解执行
claude "/tasks"

# 4. 工作流驱动开发
claude "/workflow analysis"
```

### 质量门禁实践

```bash
# 确保质量门禁
claude "启动开发阶段，严格执行 TDD 和代码质量检查"

# 检查质量门禁状态
claude "显示当前项目质量门禁检查结果"
```

### 跨工具协作

```bash
# Claude Code: 架构设计
claude "/workflow design 微服务架构设计"

# 同步到 Gemini: 详细实现
claude "同步设计结果到 gemini"
gemini "基于架构设计开始详细实现"

# 同步到 Cursor: UI 开发
gemini "同步后端进度到 cursor"
```

---

💡 **提示**: 这些示例展示了框架的强大功能，根据你的具体项目需求选择合适的工作流程！