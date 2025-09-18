"""
Universal AI Coding Framework - MCP Server
==========================================

这是基于 MCP (Model Context Protocol) 的统一 AI 编程协作服务器，
提供跨工具的项目初始化、状态同步和工作流管理功能。

核心功能：
- 自动初始化 .addp 项目结构
- Ollama 查询优化
- 跨工具状态同步
- ADDP 工作流管理
- 规格驱动开发支持
"""

__version__ = "1.0.0"
__author__ = "Universal AI Coding Framework Team"

from .server import create_mcp_server
from .tools import *

__all__ = ["create_mcp_server"]