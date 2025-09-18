"""
MCP Server 主服务器文件
实现完整的 Model Context Protocol 服务器
"""

import asyncio
import json
import logging
from pathlib import Path
from typing import Dict, List, Any, Optional
from dataclasses import dataclass

# MCP 核心导入 (需要安装 mcp package)
try:
    from mcp import types
    from mcp.server import NotificationOptions, Server
    from mcp.server.models import InitializationOptions
    from mcp.types import Resource, Tool, TextContent, ImageContent, EmbeddedResource
except ImportError:
    print("请安装 MCP 依赖: pip install mcp")
    raise

from .tools.project_tools import ProjectInitializer
from .tools.query_optimizer import QueryOptimizer
from .tools.workflow_manager import WorkflowManager
from .tools.sync_manager import SyncManager

# 配置日志
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

@dataclass
class ServerConfig:
    """服务器配置"""
    name: str = "universal-coding-assistant"
    version: str = "1.0.0"
    description: str = "Universal AI Coding Framework MCP Server"
    ollama_endpoint: str = "http://localhost:11434"
    ollama_model: str = "qwen2.5:14b"

class UniversalCodingServer:
    """统一编程协作 MCP 服务器"""

    def __init__(self, config: ServerConfig = None):
        self.config = config or ServerConfig()
        self.server = Server(self.config.name)
        self.project_initializer = ProjectInitializer()
        self.query_optimizer = QueryOptimizer(
            endpoint=self.config.ollama_endpoint,
            model=self.config.ollama_model
        )
        self.workflow_manager = WorkflowManager()
        self.sync_manager = SyncManager()

        # 注册所有工具和资源
        self._register_tools()
        self._register_resources()
        self._register_prompts()

    def _register_tools(self):
        """注册所有 MCP 工具"""

        # 1. 项目初始化工具
        @self.server.call_tool()
        async def initialize_addp_structure(
            project_type: str = "universal-coding",
            project_name: str = "",
            framework: str = "auto-detect"
        ) -> List[types.TextContent]:
            """
            自动初始化完整的 .addp 项目结构

            Args:
                project_type: 项目类型 (universal-coding, react, vue, python, kotlin, etc.)
                project_name: 项目名称 (可选，自动检测)
                framework: 框架类型 (auto-detect, react, vue, angular, django, etc.)

            用法示例:
                claude "初始化 ADDP 项目结构"
                gemini "设置统一编程环境"
            """
            try:
                result = await self.project_initializer.initialize_structure(
                    project_type=project_type,
                    project_name=project_name,
                    framework=framework
                )

                return [types.TextContent(
                    type="text",
                    text=f"✅ ADDP 项目结构初始化完成!\n\n{result['summary']}\n\n"
                         f"📁 已创建 {result['directories_created']} 个目录\n"
                         f"📄 已生成 {result['files_created']} 个模板文件\n"
                         f"🔧 已配置 {result['configs_created']} 个配置文件\n\n"
                         f"🎯 下一步: 使用 'claude \"/specify 你的需求\"' 开始规格驱动开发"
                )]
            except Exception as e:
                logger.error(f"初始化失败: {e}")
                return [types.TextContent(
                    type="text",
                    text=f"❌ 初始化失败: {str(e)}"
                )]

        # 2. 查询优化工具
        @self.server.call_tool()
        async def optimize_query(
            user_input: str,
            context: str = "",
            optimization_level: str = "smart"
        ) -> List[types.TextContent]:
            """
            使用 Ollama 本地优化用户查询

            Args:
                user_input: 用户原始输入
                context: 项目上下文 (可选)
                optimization_level: 优化级别 (basic, smart, detailed)
            """
            try:
                result = await self.query_optimizer.optimize(
                    user_input, context, optimization_level
                )

                return [types.TextContent(
                    type="text",
                    text=f"🧠 查询优化结果:\n\n"
                         f"**原始查询**: {user_input}\n\n"
                         f"**优化查询**: {result['optimized_query']}\n\n"
                         f"**改进点**: {result['improvements']}\n\n"
                         f"**建议**: {result['suggestions']}"
                )]
            except Exception as e:
                return [types.TextContent(
                    type="text",
                    text=f"❌ 查询优化失败: {str(e)}"
                )]

        # 3. ADDP 工作流工具
        @self.server.call_tool()
        async def start_addp_workflow(
            phase: str,
            specification: str = "",
            previous_phase_output: str = ""
        ) -> List[types.TextContent]:
            """
            启动 ADDP 工作流的特定阶段

            Args:
                phase: 工作流阶段 (analysis, design, development, persistence)
                specification: 规格文档内容
                previous_phase_output: 前一阶段的输出
            """
            try:
                result = await self.workflow_manager.execute_phase(
                    phase, specification, previous_phase_output
                )

                return [types.TextContent(
                    type="text",
                    text=f"⚡ ADDP {phase.title()} 阶段执行完成:\n\n{result['output']}"
                )]
            except Exception as e:
                return [types.TextContent(
                    type="text",
                    text=f"❌ ADDP 工作流执行失败: {str(e)}"
                )]

        # 4. 跨工具同步工具
        @self.server.call_tool()
        async def sync_project_state(
            action: str = "save",
            tool_name: str = "claude",
            state_data: str = ""
        ) -> List[types.TextContent]:
            """
            同步项目状态到所有 AI 工具

            Args:
                action: 操作类型 (save, load, sync_all)
                tool_name: 工具名称 (claude, gemini, cursor, codex)
                state_data: 状态数据 (JSON 格式)
            """
            try:
                result = await self.sync_manager.sync_state(
                    action, tool_name, state_data
                )

                return [types.TextContent(
                    type="text",
                    text=f"🔄 项目状态同步完成:\n\n{result['message']}"
                )]
            except Exception as e:
                return [types.TextContent(
                    type="text",
                    text=f"❌ 状态同步失败: {str(e)}"
                )]

    def _register_resources(self):
        """注册 MCP 资源"""

        @self.server.list_resources()
        async def list_resources() -> List[types.Resource]:
            """列出所有可用资源"""
            return [
                types.Resource(
                    uri="addp://specifications",
                    name="ADDP 规格文档",
                    description="项目规格文档和模板",
                    mimeType="text/markdown"
                ),
                types.Resource(
                    uri="addp://workflows",
                    name="ADDP 工作流",
                    description="四阶段工作流状态和输出",
                    mimeType="application/json"
                ),
                types.Resource(
                    uri="addp://memory",
                    name="项目记忆",
                    description="跨工具共享的项目记忆",
                    mimeType="application/json"
                ),
                types.Resource(
                    uri="addp://configs",
                    name="配置文件",
                    description="MCP 工具配置和模板",
                    mimeType="application/json"
                )
            ]

        @self.server.read_resource()
        async def read_resource(uri: str) -> str:
            """读取特定资源"""
            if uri.startswith("addp://"):
                resource_type = uri.replace("addp://", "")
                resource_path = Path(f".addp/{resource_type}")

                if resource_path.exists():
                    # 返回目录内容摘要
                    files = list(resource_path.rglob("*"))
                    content = {
                        "resource_type": resource_type,
                        "total_files": len([f for f in files if f.is_file()]),
                        "files": [str(f.relative_to(resource_path)) for f in files if f.is_file()]
                    }
                    return json.dumps(content, indent=2, ensure_ascii=False)

            return json.dumps({"error": f"Resource not found: {uri}"})

    def _register_prompts(self):
        """注册提示模板"""

        @self.server.list_prompts()
        async def list_prompts() -> List[types.Prompt]:
            """列出所有可用的提示模板"""
            return [
                types.Prompt(
                    name="spec-driven-analysis",
                    description="规格驱动的需求分析模板",
                    arguments=[
                        types.PromptArgument(
                            name="requirement",
                            description="用户需求描述",
                            required=True
                        )
                    ]
                ),
                types.Prompt(
                    name="addp-workflow",
                    description="ADDP 四阶段工作流模板",
                    arguments=[
                        types.PromptArgument(
                            name="phase",
                            description="工作流阶段",
                            required=True
                        ),
                        types.PromptArgument(
                            name="input",
                            description="阶段输入",
                            required=True
                        )
                    ]
                )
            ]

        @self.server.get_prompt()
        async def get_prompt(name: str, arguments: Dict[str, str]) -> types.GetPromptResult:
            """获取特定的提示模板"""

            if name == "spec-driven-analysis":
                requirement = arguments.get("requirement", "")

                prompt_text = f"""# 规格驱动需求分析

## 用户需求
{requirement}

## 分析流程

### 1. 需求澄清
- 核心功能是什么？
- 用户场景有哪些？
- 成功标准是什么？

### 2. 技术约束
- 现有技术栈
- 性能要求
- 兼容性需求

### 3. 实施规划
- 技术方案选择
- 开发阶段划分
- 验收标准定义

请基于以上分析生成详细的 PRD (产品需求文档)。
"""

                return types.GetPromptResult(
                    description=f"规格驱动分析: {requirement}",
                    messages=[
                        types.PromptMessage(
                            role="user",
                            content=types.TextContent(type="text", text=prompt_text)
                        )
                    ]
                )

            elif name == "addp-workflow":
                phase = arguments.get("phase", "")
                input_data = arguments.get("input", "")

                phase_templates = {
                    "analysis": "深入分析需求，识别技术约束和风险",
                    "design": "设计架构方案，权衡技术选择",
                    "development": "TDD 驱动开发，最小化修改",
                    "persistence": "验证结果，更新项目记忆"
                }

                template = phase_templates.get(phase, "执行 ADDP 工作流")

                prompt_text = f"""# ADDP {phase.title()} 阶段

## 输入
{input_data}

## 执行指导
{template}

## 输出要求
- 具体可执行的步骤
- 明确的验收标准
- 下一阶段的输入准备
"""

                return types.GetPromptResult(
                    description=f"ADDP {phase} 阶段执行",
                    messages=[
                        types.PromptMessage(
                            role="user",
                            content=types.TextContent(type="text", text=prompt_text)
                        )
                    ]
                )

            raise ValueError(f"未知的提示模板: {name}")

async def create_mcp_server(config: ServerConfig = None) -> UniversalCodingServer:
    """创建并配置 MCP 服务器实例"""
    server = UniversalCodingServer(config)

    # 服务器初始化
    @server.server.initialize()
    async def initialize(
        initialization_options: InitializationOptions
    ) -> InitializationOptions:
        logger.info(f"Universal Coding MCP Server 初始化完成")
        logger.info(f"服务器名称: {server.config.name}")
        logger.info(f"版本: {server.config.version}")
        logger.info(f"Ollama 端点: {server.config.ollama_endpoint}")
        logger.info(f"使用模型: {server.config.ollama_model}")

        return initialization_options

    return server

if __name__ == "__main__":
    # 开发模式启动
    async def main():
        server = await create_mcp_server()

        # 这里应该启动 MCP 协议的标准 stdio 通信
        # 实际部署时会通过 MCP 客户端调用
        logger.info("MCP Server 已启动，等待客户端连接...")

        # 保持服务运行
        try:
            await asyncio.Event().wait()
        except KeyboardInterrupt:
            logger.info("服务器关闭")

    asyncio.run(main())