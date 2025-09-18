"""
跨工具状态同步管理器
==================

实现 Claude Code、Gemini CLI、Cursor 等 AI 工具之间的状态同步，
确保所有工具都能共享项目上下文、开发进度和决策记录。
"""

import json
import logging
from pathlib import Path
from typing import Dict, List, Any, Optional
from datetime import datetime

logger = logging.getLogger(__name__)

class SyncManager:
    """跨工具状态同步管理器"""

    def __init__(self):
        self.sync_path = Path(".addp/sync")
        self.sync_path.mkdir(parents=True, exist_ok=True)

        # 支持的工具列表
        self.supported_tools = ["claude", "gemini", "cursor", "universal"]

        # 确保各工具同步目录存在
        for tool in self.supported_tools:
            (self.sync_path / tool).mkdir(exist_ok=True)

    async def sync_state(
        self,
        action: str = "save",
        tool_name: str = "claude",
        state_data: str = ""
    ) -> Dict[str, Any]:
        """
        同步项目状态

        Args:
            action: 操作类型 (save, load, sync_all)
            tool_name: 工具名称
            state_data: 状态数据 (JSON 格式)

        Returns:
            同步操作结果
        """
        try:
            if action == "save":
                return await self._save_tool_state(tool_name, state_data)
            elif action == "load":
                return await self._load_tool_state(tool_name)
            elif action == "sync_all":
                return await self._sync_all_tools()
            else:
                raise Exception(f"不支持的同步操作: {action}")

        except Exception as e:
            logger.error(f"状态同步失败: {e}")
            raise

    async def _save_tool_state(self, tool_name: str, state_data: str) -> Dict[str, Any]:
        """保存工具状态"""
        try:
            if tool_name not in self.supported_tools:
                raise Exception(f"不支持的工具: {tool_name}")

            # 解析状态数据
            if state_data:
                try:
                    parsed_state = json.loads(state_data)
                except json.JSONDecodeError:
                    # 如果不是 JSON，作为文本处理
                    parsed_state = {"raw_data": state_data}
            else:
                # 生成默认状态
                parsed_state = await self._generate_default_state(tool_name)

            # 添加元数据
            full_state = {
                "tool": tool_name,
                "timestamp": datetime.now().isoformat(),
                "version": "1.0.0",
                "state": parsed_state,
                "sync_metadata": {
                    "last_sync": datetime.now().isoformat(),
                    "sync_source": tool_name,
                    "sync_version": 1
                }
            }

            # 保存到工具特定目录
            tool_state_file = self.sync_path / tool_name / "current_state.json"
            with open(tool_state_file, 'w', encoding='utf-8') as f:
                json.dump(full_state, f, indent=2, ensure_ascii=False)

            # 保存到通用同步目录
            await self._update_universal_state(tool_name, full_state)

            # 记录同步历史
            await self._record_sync_history(tool_name, "save", full_state)

            return {
                "success": True,
                "message": f"工具 {tool_name} 状态保存成功",
                "tool": tool_name,
                "timestamp": full_state["timestamp"]
            }

        except Exception as e:
            logger.error(f"保存工具状态失败: {e}")
            raise

    async def _load_tool_state(self, tool_name: str) -> Dict[str, Any]:
        """加载工具状态"""
        try:
            if tool_name not in self.supported_tools:
                raise Exception(f"不支持的工具: {tool_name}")

            tool_state_file = self.sync_path / tool_name / "current_state.json"

            if tool_state_file.exists():
                with open(tool_state_file, 'r', encoding='utf-8') as f:
                    state_data = json.load(f)

                return {
                    "success": True,
                    "message": f"工具 {tool_name} 状态加载成功",
                    "state": state_data,
                    "last_update": state_data.get("timestamp", "unknown")
                }
            else:
                # 如果没有保存的状态，尝试从通用状态创建
                universal_state = await self._load_universal_state()
                if universal_state:
                    adapted_state = await self._adapt_state_for_tool(tool_name, universal_state)
                    return {
                        "success": True,
                        "message": f"为 {tool_name} 创建了适配状态",
                        "state": adapted_state,
                        "source": "universal"
                    }
                else:
                    return {
                        "success": False,
                        "message": f"工具 {tool_name} 暂无保存的状态",
                        "state": {}
                    }

        except Exception as e:
            logger.error(f"加载工具状态失败: {e}")
            raise

    async def _sync_all_tools(self) -> Dict[str, Any]:
        """同步所有工具状态"""
        try:
            sync_results = {}

            # 加载通用状态作为基准
            universal_state = await self._load_universal_state()

            # 为每个工具创建适配状态
            for tool in self.supported_tools:
                if tool == "universal":
                    continue

                try:
                    # 为工具创建适配状态
                    adapted_state = await self._adapt_state_for_tool(tool, universal_state)

                    # 保存适配状态
                    await self._save_tool_state(tool, json.dumps(adapted_state))

                    sync_results[tool] = {
                        "status": "success",
                        "message": f"{tool} 状态同步完成"
                    }

                except Exception as e:
                    sync_results[tool] = {
                        "status": "failed",
                        "error": str(e)
                    }

            # 统计同步结果
            success_count = sum(1 for result in sync_results.values() if result["status"] == "success")
            total_count = len(sync_results)

            return {
                "success": True,
                "message": f"跨工具同步完成: {success_count}/{total_count} 成功",
                "results": sync_results,
                "sync_timestamp": datetime.now().isoformat()
            }

        except Exception as e:
            logger.error(f"跨工具同步失败: {e}")
            raise

    async def _generate_default_state(self, tool_name: str) -> Dict[str, Any]:
        """生成工具的默认状态"""

        # 加载项目基础信息
        project_info = await self._load_project_info()

        base_state = {
            "project_info": project_info,
            "current_session": {
                "start_time": datetime.now().isoformat(),
                "tool": tool_name,
                "mode": "development",
                "context": "universal_ai_coding_framework"
            },
            "workflow_state": {
                "current_phase": "initialized",
                "completed_phases": [],
                "next_actions": []
            },
            "memory": {
                "decisions": [],
                "lessons": [],
                "context": {}
            }
        }

        # 为不同工具添加特定配置
        if tool_name == "claude":
            base_state["claude_specific"] = {
                "subagents_enabled": True,
                "mcp_tools_available": True,
                "output_style": "orchestrator"
            }
        elif tool_name == "gemini":
            base_state["gemini_specific"] = {
                "context_window": "1M_tokens",
                "mcp_integration": "native",
                "workflow_mode": "addp"
            }
        elif tool_name == "cursor":
            base_state["cursor_specific"] = {
                "agent_mode": True,
                "mcp_config": ".cursor/mcp.json",
                "rules_system": "enabled"
            }

        return base_state

    async def _update_universal_state(self, source_tool: str, tool_state: Dict[str, Any]):
        """更新通用状态"""
        try:
            universal_file = self.sync_path / "universal" / "shared_state.json"

            # 加载现有通用状态
            universal_state = {}
            if universal_file.exists():
                with open(universal_file, 'r', encoding='utf-8') as f:
                    universal_state = json.load(f)

            # 合并状态信息
            universal_state.update({
                "last_update": datetime.now().isoformat(),
                "last_update_source": source_tool,
                "project_info": tool_state["state"].get("project_info", {}),
                "workflow_state": tool_state["state"].get("workflow_state", {}),
                "shared_memory": tool_state["state"].get("memory", {}),
                "tool_states": universal_state.get("tool_states", {})
            })

            # 记录各工具状态摘要
            universal_state["tool_states"][source_tool] = {
                "last_sync": tool_state["timestamp"],
                "version": tool_state.get("version", "unknown"),
                "status": "active"
            }

            # 保存更新的通用状态
            with open(universal_file, 'w', encoding='utf-8') as f:
                json.dump(universal_state, f, indent=2, ensure_ascii=False)

        except Exception as e:
            logger.warning(f"更新通用状态失败: {e}")

    async def _load_universal_state(self) -> Dict[str, Any]:
        """加载通用状态"""
        try:
            universal_file = self.sync_path / "universal" / "shared_state.json"

            if universal_file.exists():
                with open(universal_file, 'r', encoding='utf-8') as f:
                    return json.load(f)
            else:
                # 创建初始通用状态
                initial_state = {
                    "created": datetime.now().isoformat(),
                    "version": "1.0.0",
                    "project_info": await self._load_project_info(),
                    "workflow_state": {"current_phase": "initialized"},
                    "shared_memory": {},
                    "tool_states": {}
                }

                with open(universal_file, 'w', encoding='utf-8') as f:
                    json.dump(initial_state, f, indent=2, ensure_ascii=False)

                return initial_state

        except Exception as e:
            logger.warning(f"加载通用状态失败: {e}")
            return {}

    async def _adapt_state_for_tool(self, tool_name: str, universal_state: Dict[str, Any]) -> Dict[str, Any]:
        """为特定工具适配通用状态"""

        base_adapted = {
            "project_info": universal_state.get("project_info", {}),
            "workflow_state": universal_state.get("workflow_state", {}),
            "shared_memory": universal_state.get("shared_memory", {}),
            "adapted_for": tool_name,
            "adaptation_time": datetime.now().isoformat()
        }

        # 工具特定适配
        if tool_name == "claude":
            base_adapted["claude_commands"] = [
                "/sync - 同步项目状态",
                "/workflow [phase] - 启动 ADDP 工作流",
                "/optimize [query] - 优化查询",
                "/specify [requirement] - 规格驱动开发"
            ]
            base_adapted["mcp_tools"] = [
                "initialize_addp_structure",
                "optimize_query",
                "start_addp_workflow",
                "sync_project_state"
            ]

        elif tool_name == "gemini":
            base_adapted["gemini_context"] = {
                "workspace": universal_state.get("project_info", {}).get("path", ""),
                "framework": universal_state.get("project_info", {}).get("framework", ""),
                "mcp_integration": "enabled"
            }

        elif tool_name == "cursor":
            base_adapted["cursor_config"] = {
                "rules_active": True,
                "mcp_enabled": True,
                "agent_mode": "development"
            }

        return base_adapted

    async def _load_project_info(self) -> Dict[str, Any]:
        """加载项目信息"""
        try:
            # 尝试从元数据文件加载
            metadata_file = Path(".addp/metadata.json")
            if metadata_file.exists():
                with open(metadata_file, 'r', encoding='utf-8') as f:
                    metadata = json.load(f)
                    return metadata.get("project_info", {})

            # 尝试从项目记忆加载
            memory_file = Path(".addp/memory/context/project_context.json")
            if memory_file.exists():
                with open(memory_file, 'r', encoding='utf-8') as f:
                    memory = json.load(f)
                    return memory.get("project_info", {})

            # 返回基础信息
            return {
                "name": Path.cwd().name,
                "path": str(Path.cwd()),
                "framework": "unknown",
                "initialized": False
            }

        except Exception as e:
            logger.warning(f"加载项目信息失败: {e}")
            return {}

    async def _record_sync_history(
        self,
        tool_name: str,
        action: str,
        state_data: Dict[str, Any]
    ):
        """记录同步历史"""
        try:
            history_file = self.sync_path / "sync_history.jsonl"

            history_entry = {
                "timestamp": datetime.now().isoformat(),
                "tool": tool_name,
                "action": action,
                "data_size": len(json.dumps(state_data)),
                "success": True
            }

            # 追加到历史文件
            with open(history_file, 'a', encoding='utf-8') as f:
                f.write(json.dumps(history_entry, ensure_ascii=False) + '\n')

        except Exception as e:
            logger.warning(f"记录同步历史失败: {e}")

    async def get_sync_status(self) -> Dict[str, Any]:
        """获取同步状态"""
        try:
            status = {
                "sync_enabled": True,
                "supported_tools": self.supported_tools,
                "tool_status": {},
                "last_sync_times": {},
                "universal_state_available": False
            }

            # 检查各工具状态
            for tool in self.supported_tools:
                tool_state_file = self.sync_path / tool / "current_state.json"

                if tool_state_file.exists():
                    try:
                        with open(tool_state_file, 'r', encoding='utf-8') as f:
                            tool_data = json.load(f)
                            status["tool_status"][tool] = "active"
                            status["last_sync_times"][tool] = tool_data.get("timestamp", "unknown")
                    except:
                        status["tool_status"][tool] = "error"
                else:
                    status["tool_status"][tool] = "not_synced"

            # 检查通用状态
            universal_file = self.sync_path / "universal" / "shared_state.json"
            status["universal_state_available"] = universal_file.exists()

            # 计算同步健康度
            active_tools = sum(1 for s in status["tool_status"].values() if s == "active")
            total_tools = len(self.supported_tools)
            status["sync_health"] = f"{active_tools}/{total_tools} 工具已同步"

            return status

        except Exception as e:
            logger.error(f"获取同步状态失败: {e}")
            return {"error": str(e)}

    async def cleanup_old_states(self, days_to_keep: int = 7) -> Dict[str, Any]:
        """清理旧的状态文件"""
        try:
            current_time = datetime.now()
            cleanup_results = {
                "removed_files": 0,
                "total_size_freed": 0,
                "errors": []
            }

            # 清理历史文件
            history_file = self.sync_path / "sync_history.jsonl"
            if history_file.exists():
                try:
                    # 读取并过滤历史记录
                    with open(history_file, 'r', encoding='utf-8') as f:
                        lines = f.readlines()

                    filtered_lines = []
                    for line in lines:
                        try:
                            entry = json.loads(line.strip())
                            entry_time = datetime.fromisoformat(entry["timestamp"])
                            if (current_time - entry_time).days <= days_to_keep:
                                filtered_lines.append(line)
                        except:
                            continue

                    # 重写历史文件
                    with open(history_file, 'w', encoding='utf-8') as f:
                        f.writelines(filtered_lines)

                    cleanup_results["cleaned_history"] = True

                except Exception as e:
                    cleanup_results["errors"].append(f"清理历史文件失败: {e}")

            return cleanup_results

        except Exception as e:
            logger.error(f"清理状态文件失败: {e}")
            return {"error": str(e)}