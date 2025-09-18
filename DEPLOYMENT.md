# 🚀 Universal AI Coding Framework 部署指南

## 📋 系统要求

### 基础环境
- **Python**: 3.8+
- **操作系统**: Windows 10+, macOS 10.15+, Linux Ubuntu 18.04+
- **内存**: 建议 8GB+ (Ollama 模型需要)
- **存储**: 至少 10GB 可用空间

### 依赖服务
- **Ollama**: 本地 LLM 服务 ([安装指南](https://ollama.ai))
- **Git**: 版本控制 (可选，用于克隆项目)

## 🏗️ 安装部署

### 方法一：一键快速部署 (推荐)

```bash
# 1. 克隆项目
git clone https://github.com/your-org/universal-ai-coding-framework.git
cd universal-ai-coding-framework

# 2. 运行快速启动脚本
python scripts/quick_start.py
```

快速启动脚本会自动：
- ✅ 检查环境依赖
- ✅ 安装 Python 依赖
- ✅ 检查 Ollama 服务
- ✅ 初始化 .addp 项目结构
- ✅ 生成配置文件
- ✅ 显示后续配置步骤

### 方法二：手动部署

#### 步骤 1: 安装 Ollama
```bash
# macOS/Linux
curl -fsSL https://ollama.ai/install.sh | sh

# Windows
# 下载并安装: https://ollama.ai/download
```

#### 步骤 2: 安装 Python 依赖
```bash
pip install -r requirements.txt
```

#### 步骤 3: 下载推荐模型
```bash
ollama pull qwen2.5:14b
```

#### 步骤 4: 初始化项目结构
```bash
python main.py --init
```

#### 步骤 5: 生成配置文件
```bash
python main.py --save-config
```

## 🔧 AI 工具配置

### Claude Code 配置

1. **添加 MCP 服务器**
```bash
claude config mcp-servers add universal-coding-assistant
```

2. **配置服务器路径**
在 Claude Code 配置中添加：
```json
{
  "mcpServers": {
    "universal-coding-assistant": {
      "command": "python",
      "args": ["main.py"],
      "cwd": "/path/to/universal-ai-coding-framework"
    }
  }
}
```

### Gemini CLI 配置

1. **配置 MCP 服务器**
```bash
gemini config mcp-servers.universal-coding-assistant.command "python main.py"
gemini config mcp-servers.universal-coding-assistant.cwd "/path/to/framework"
```

2. **验证配置**
```bash
gemini config list
```

### Cursor 配置

1. **创建 MCP 配置文件**
在项目根目录创建 `.cursor/mcp.json`：
```json
{
  "mcpServers": {
    "universal-coding-assistant": {
      "command": "python",
      "args": ["main.py"],
      "cwd": "./universal-ai-coding-framework"
    }
  }
}
```

2. **重启 Cursor** 以加载配置

## 🧪 功能测试

### 基础功能测试

#### 1. 项目初始化测试
```bash
# Claude Code
claude "初始化 ADDP 项目结构"

# Gemini CLI
gemini "设置统一编程环境"

# 预期结果: 创建完整的 .addp/ 目录结构
```

#### 2. 查询优化测试
```bash
# Claude Code
claude "优化这个查询: 实现用户登录功能"

# 预期结果: 返回详细的技术规格和实施建议
```

#### 3. ADDP 工作流测试
```bash
# Claude Code
claude "启动 ADDP 分析阶段"

# 预期结果: 生成分析阶段的模板和检查清单
```

#### 4. 跨工具同步测试
```bash
# 在 Claude Code 中
claude "保存当前项目状态"

# 在 Gemini CLI 中
gemini "加载项目状态"

# 预期结果: 状态信息成功同步
```

### 高级功能测试

#### 规格驱动开发流程
```bash
# 1. 创建需求规格
claude "/specify 实现用户注册功能，包含邮箱验证"

# 2. 生成技术方案
claude "/plan"

# 3. 分解开发任务
claude "/tasks"

# 4. 启动 ADDP 工作流
claude "/workflow analysis"
```

## 🔍 故障排除

### 常见问题

#### 1. Ollama 连接失败
**症状**: 查询优化功能报错 "Ollama API 错误"

**解决方案**:
```bash
# 检查 Ollama 服务状态
ollama serve

# 验证模型安装
ollama list

# 测试 API 连接
curl http://localhost:11434/api/generate -d '{"model":"qwen2.5:14b","prompt":"test"}'
```

#### 2. MCP 工具不可用
**症状**: AI 工具无法识别 MCP 命令

**解决方案**:
```bash
# 检查配置文件
python main.py --dev

# 验证 MCP 服务器状态
python -c "from src.mcp_server.server import create_mcp_server; print('MCP OK')"
```

#### 3. .addp 目录结构异常
**症状**: 目录结构不完整或缺失

**解决方案**:
```bash
# 重新初始化
python main.py --init

# 检查权限
ls -la .addp/
```

#### 4. 依赖安装失败
**症状**: pip install 报错

**解决方案**:
```bash
# 升级 pip
python -m pip install --upgrade pip

# 使用国内镜像
pip install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple/
```

### 日志调试

#### 启用详细日志
```bash
# 开发模式 (详细日志)
python main.py --dev

# 设置日志级别
export PYTHONPATH=. LOGLEVEL=DEBUG python main.py
```

#### 检查日志文件
```bash
# 查看 MCP 服务器日志
tail -f .addp/analytics/logs/mcp_server.log

# 查看查询优化日志
tail -f .addp/analytics/logs/query_optimization.log
```

## 📊 性能优化

### Ollama 优化
```bash
# 设置环境变量优化性能
export OLLAMA_HOST=0.0.0.0:11434
export OLLAMA_MODELS=/path/to/models
export OLLAMA_NUM_PARALLEL=4
```

### 缓存配置
编辑 `.addp/configs/mcp/server_config.json`:
```json
{
  "project": {
    "cache_enabled": true,
    "cache_ttl": 86400,
    "max_cache_size": "500MB"
  }
}
```

### 内存优化
```bash
# 对于 8GB 内存系统，使用较小模型
ollama pull qwen2.5:7b

# 配置文件中修改模型
{
  "ollama": {
    "model": "qwen2.5:7b"
  }
}
```

## 🔄 更新升级

### 更新框架
```bash
# 拉取最新代码
git pull origin main

# 更新依赖
pip install -r requirements.txt --upgrade

# 重新初始化 (保留现有数据)
python main.py --init
```

### 备份数据
```bash
# 备份项目记忆和配置
tar -czf addp_backup_$(date +%Y%m%d).tar.gz .addp/
```

## 🚀 生产部署

### Docker 部署 (推荐)
```dockerfile
# 创建 Dockerfile
FROM python:3.11-slim

WORKDIR /app
COPY . .
RUN pip install -r requirements.txt

EXPOSE 8000
CMD ["python", "main.py"]
```

### 系统服务
```bash
# 创建 systemd 服务
sudo tee /etc/systemd/system/universal-coding.service << EOF
[Unit]
Description=Universal AI Coding Framework
After=network.target

[Service]
Type=simple
User=your-user
WorkingDirectory=/path/to/framework
ExecStart=/usr/bin/python main.py
Restart=always

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl enable universal-coding
sudo systemctl start universal-coding
```

## 📞 支持与反馈

- **GitHub Issues**: [提交问题](https://github.com/your-org/universal-ai-coding-framework/issues)
- **文档**: [在线文档](https://docs.universal-ai-coding.org)
- **社区**: [GitHub Discussions](https://github.com/your-org/universal-ai-coding-framework/discussions)

---

🎉 **部署完成！开始享受统一的 AI 编程协作体验！**