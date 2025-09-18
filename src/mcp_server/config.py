"""
MCP Server 配置管理
==================

管理服务器配置、工具注册和环境设置。
"""

import json
import os
from pathlib import Path
from typing import Dict, List, Any, Optional
from dataclasses import dataclass, asdict

@dataclass
class OllamaConfig:
    """Ollama 配置"""
    endpoint: str = "http://localhost:11434"
    model: str = "qwen2.5:14b"
    temperature: float = 0.7
    max_tokens: int = 2048
    timeout: int = 30

@dataclass
class MCPServerConfig:
    """MCP 服务器配置"""
    name: str = "universal-coding-assistant"
    version: str = "1.0.0"
    description: str = "Universal AI Coding Framework MCP Server"
    host: str = "localhost"
    port: int = 8000
    debug: bool = False

@dataclass
class ProjectConfig:
    """项目配置"""
    addp_directory: str = ".addp"
    auto_initialize: bool = True
    quality_gates_enabled: bool = True
    cross_tool_sync: bool = True
    analytics_enabled: bool = True

@dataclass
class UniversalConfig:
    """统一配置"""
    server: MCPServerConfig
    ollama: OllamaConfig
    project: ProjectConfig

    @classmethod
    def from_file(cls, config_path: str) -> "UniversalConfig":
        """从配置文件加载"""
        try:
            with open(config_path, 'r', encoding='utf-8') as f:
                config_data = json.load(f)

            return cls(
                server=MCPServerConfig(**config_data.get("server", {})),
                ollama=OllamaConfig(**config_data.get("ollama", {})),
                project=ProjectConfig(**config_data.get("project", {}))
            )
        except FileNotFoundError:
            # 返回默认配置
            return cls(
                server=MCPServerConfig(),
                ollama=OllamaConfig(),
                project=ProjectConfig()
            )
        except Exception as e:
            raise Exception(f"配置文件加载失败: {e}")

    def to_file(self, config_path: str):
        """保存到配置文件"""
        config_data = {
            "server": asdict(self.server),
            "ollama": asdict(self.ollama),
            "project": asdict(self.project)
        }

        # 确保配置目录存在
        Path(config_path).parent.mkdir(parents=True, exist_ok=True)

        with open(config_path, 'w', encoding='utf-8') as f:
            json.dump(config_data, f, indent=2, ensure_ascii=False)

    def merge_env_vars(self):
        """合并环境变量"""
        # Ollama 配置
        if os.getenv("OLLAMA_ENDPOINT"):
            self.ollama.endpoint = os.getenv("OLLAMA_ENDPOINT")
        if os.getenv("OLLAMA_MODEL"):
            self.ollama.model = os.getenv("OLLAMA_MODEL")

        # 服务器配置
        if os.getenv("MCP_SERVER_HOST"):
            self.server.host = os.getenv("MCP_SERVER_HOST")
        if os.getenv("MCP_SERVER_PORT"):
            self.server.port = int(os.getenv("MCP_SERVER_PORT"))
        if os.getenv("MCP_SERVER_DEBUG"):
            self.server.debug = os.getenv("MCP_SERVER_DEBUG").lower() == "true"

        # 项目配置
        if os.getenv("ADDP_DIRECTORY"):
            self.project.addp_directory = os.getenv("ADDP_DIRECTORY")

# 工具注册表
TOOL_REGISTRY = {
    "initialize_addp_structure": {
        "name": "initialize_addp_structure",
        "description": "自动初始化完整的 .addp 项目结构",
        "parameters": [
            {
                "name": "project_type",
                "type": "string",
                "description": "项目类型",
                "default": "universal-coding"
            },
            {
                "name": "project_name",
                "type": "string",
                "description": "项目名称",
                "default": ""
            },
            {
                "name": "framework",
                "type": "string",
                "description": "框架类型",
                "default": "auto-detect"
            }
        ]
    },
    "optimize_query": {
        "name": "optimize_query",
        "description": "使用 Ollama 本地优化用户查询",
        "parameters": [
            {
                "name": "user_input",
                "type": "string",
                "description": "用户原始输入",
                "required": True
            },
            {
                "name": "context",
                "type": "string",
                "description": "项目上下文",
                "default": ""
            },
            {
                "name": "optimization_level",
                "type": "string",
                "description": "优化级别",
                "default": "smart",
                "enum": ["basic", "smart", "detailed"]
            }
        ]
    },
    "start_addp_workflow": {
        "name": "start_addp_workflow",
        "description": "启动 ADDP 工作流的特定阶段",
        "parameters": [
            {
                "name": "phase",
                "type": "string",
                "description": "工作流阶段",
                "required": True,
                "enum": ["analysis", "design", "development", "persistence"]
            },
            {
                "name": "specification",
                "type": "string",
                "description": "规格文档内容",
                "default": ""
            },
            {
                "name": "previous_phase_output",
                "type": "string",
                "description": "前一阶段的输出",
                "default": ""
            }
        ]
    },
    "sync_project_state": {
        "name": "sync_project_state",
        "description": "同步项目状态到所有 AI 工具",
        "parameters": [
            {
                "name": "action",
                "type": "string",
                "description": "操作类型",
                "default": "save",
                "enum": ["save", "load", "sync_all"]
            },
            {
                "name": "tool_name",
                "type": "string",
                "description": "工具名称",
                "default": "claude",
                "enum": ["claude", "gemini", "cursor", "codex"]
            },
            {
                "name": "state_data",
                "type": "string",
                "description": "状态数据 (JSON 格式)",
                "default": ""
            }
        ]
    }
}

# 资源注册表
RESOURCE_REGISTRY = {
    "addp://specifications": {
        "uri": "addp://specifications",
        "name": "ADDP 规格文档",
        "description": "项目规格文档和模板",
        "mimeType": "text/markdown"
    },
    "addp://workflows": {
        "uri": "addp://workflows",
        "name": "ADDP 工作流",
        "description": "四阶段工作流状态和输出",
        "mimeType": "application/json"
    },
    "addp://memory": {
        "uri": "addp://memory",
        "name": "项目记忆",
        "description": "跨工具共享的项目记忆",
        "mimeType": "application/json"
    },
    "addp://configs": {
        "uri": "addp://configs",
        "name": "配置文件",
        "description": "MCP 工具配置和模板",
        "mimeType": "application/json"
    }
}

# 提示模板注册表
PROMPT_REGISTRY = {
    "spec-driven-analysis": {
        "name": "spec-driven-analysis",
        "description": "规格驱动的需求分析模板",
        "arguments": [
            {
                "name": "requirement",
                "description": "用户需求描述",
                "required": True
            }
        ]
    },
    "addp-workflow": {
        "name": "addp-workflow",
        "description": "ADDP 四阶段工作流模板",
        "arguments": [
            {
                "name": "phase",
                "description": "工作流阶段",
                "required": True
            },
            {
                "name": "input",
                "description": "阶段输入",
                "required": True
            }
        ]
    }
}

def load_config(config_path: Optional[str] = None) -> UniversalConfig:
    """加载配置"""
    if not config_path:
        # 默认配置路径
        config_path = Path(".addp/configs/mcp/server_config.json")
        if not config_path.exists():
            config_path = "config.json"

    config = UniversalConfig.from_file(str(config_path))
    config.merge_env_vars()
    return config

def save_default_config(config_path: str = ".addp/configs/mcp/server_config.json"):
    """保存默认配置"""
    config = UniversalConfig(
        server=MCPServerConfig(),
        ollama=OllamaConfig(),
        project=ProjectConfig()
    )
    config.to_file(config_path)