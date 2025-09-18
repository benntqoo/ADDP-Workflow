"""
MCP Tools 工具包
===============

这个包包含了所有 MCP 工具的实现，提供完整的项目初始化、
查询优化、工作流管理和状态同步功能。

主要模块：
- project_tools: 项目初始化和结构管理
- query_optimizer: Ollama 驱动的查询优化
- workflow_manager: ADDP 四阶段工作流管理
- sync_manager: 跨工具状态同步
"""

from .project_tools import ProjectInitializer
from .query_optimizer import QueryOptimizer
from .workflow_manager import WorkflowManager
from .sync_manager import SyncManager

__all__ = [
    "ProjectInitializer",
    "QueryOptimizer",
    "WorkflowManager",
    "SyncManager"
]