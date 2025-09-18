#!/usr/bin/env python3
"""
MCP 工具测试脚本
================

测试所有 MCP 工具的功能，验证 Universal AI Coding Framework 的完整性。

使用方法:
    python scripts/test_mcp_tools.py              # 运行所有测试
    python scripts/test_mcp_tools.py --tool init  # 测试特定工具
    python scripts/test_mcp_tools.py --verbose    # 详细输出
"""

import asyncio
import argparse
import json
import sys
import tempfile
import shutil
from pathlib import Path
from typing import Dict, List, Any, Optional

# 添加源代码路径
sys.path.insert(0, str(Path(__file__).parent.parent))

from src.mcp_server.tools.project_tools import ProjectInitializer
from src.mcp_server.tools.query_optimizer import QueryOptimizer
from src.mcp_server.tools.workflow_manager import WorkflowManager
from src.mcp_server.tools.sync_manager import SyncManager

class MCPToolsTester:
    """MCP 工具测试器"""

    def __init__(self, verbose: bool = False):
        self.verbose = verbose
        self.test_results = {}
        self.temp_dir = None

    async def run_all_tests(self) -> Dict[str, Any]:
        """运行所有测试"""
        print("🧪 开始 MCP 工具测试")
        print("=" * 50)

        # 创建临时测试环境
        self.temp_dir = Path(tempfile.mkdtemp(prefix="mcp_test_"))
        original_cwd = Path.cwd()

        try:
            # 切换到测试目录
            import os
            os.chdir(self.temp_dir)

            # 运行各项测试
            await self.test_project_initialization()
            await self.test_query_optimization()
            await self.test_workflow_management()
            await self.test_sync_management()

            # 生成测试报告
            self.generate_test_report()

        finally:
            # 恢复原目录
            os.chdir(original_cwd)
            # 清理测试环境
            if self.temp_dir and self.temp_dir.exists():
                shutil.rmtree(self.temp_dir)

        return self.test_results

    async def test_project_initialization(self):
        """测试项目初始化功能"""
        print("\n🏗️  测试项目初始化...")

        test_name = "project_initialization"
        try:
            initializer = ProjectInitializer()

            # 测试自动检测
            result = await initializer.initialize_structure(
                project_type="test",
                project_name="mcp_test_project",
                framework="auto-detect"
            )

            # 验证结果
            checks = []
            checks.append(("初始化成功", result.get("success", False)))
            checks.append(("目录创建", result.get("directories_created", 0) > 0))
            checks.append(("文件生成", result.get("files_created", 0) > 0))
            checks.append(("配置创建", result.get("configs_created", 0) > 0))

            # 验证目录结构
            addp_path = Path(".addp")
            checks.append(("ADDP目录存在", addp_path.exists()))

            if addp_path.exists():
                expected_dirs = ["specifications", "workflows", "memory", "queries", "gates", "sync"]
                for dir_name in expected_dirs:
                    checks.append((f"{dir_name}目录", (addp_path / dir_name).exists()))

            # 验证关键文件
            key_files = [
                ".addp/metadata.json",
                ".addp/README.md",
                ".addp/specifications/templates/prd_template.md",
                ".addp/configs/mcp/server_config.json"
            ]

            for file_path in key_files:
                checks.append((f"文件: {Path(file_path).name}", Path(file_path).exists()))

            self.test_results[test_name] = {
                "status": "passed" if all(check[1] for check in checks) else "failed",
                "checks": checks,
                "details": result
            }

            if self.verbose:
                for check_name, passed in checks:
                    status = "✅" if passed else "❌"
                    print(f"  {status} {check_name}")

        except Exception as e:
            self.test_results[test_name] = {
                "status": "error",
                "error": str(e),
                "checks": []
            }
            print(f"❌ 项目初始化测试失败: {e}")

    async def test_query_optimization(self):
        """测试查询优化功能"""
        print("\n🧠 测试查询优化...")

        test_name = "query_optimization"
        try:
            optimizer = QueryOptimizer()

            # 测试用例
            test_cases = [
                {
                    "name": "基础优化",
                    "input": "优化性能",
                    "level": "basic",
                    "context": ""
                },
                {
                    "name": "智能优化",
                    "input": "实现用户登录",
                    "level": "smart",
                    "context": "React 前端项目"
                },
                {
                    "name": "详细优化",
                    "input": "重构代码",
                    "level": "detailed",
                    "context": "Python Flask 后端"
                }
            ]

            checks = []
            for test_case in test_cases:
                try:
                    result = await optimizer.optimize(
                        test_case["input"],
                        test_case["context"],
                        test_case["level"]
                    )

                    # 验证优化结果
                    required_fields = ["optimized_query", "improvements", "suggestions", "confidence"]
                    case_passed = all(field in result for field in required_fields)

                    # 验证优化质量
                    if case_passed:
                        optimized_length = len(result.get("optimized_query", ""))
                        original_length = len(test_case["input"])
                        improvement_ratio = optimized_length / max(original_length, 1)

                        # 优化后的查询应该更详细 (通常更长)
                        case_passed = improvement_ratio > 0.8

                    checks.append((test_case["name"], case_passed))

                except Exception as e:
                    checks.append((test_case["name"], False))
                    if self.verbose:
                        print(f"  ❌ {test_case['name']} 失败: {e}")

            # 测试缓存功能
            try:
                # 第二次调用相同查询，应该使用缓存
                cached_result = await optimizer.optimize("优化性能", "", "basic")
                cache_working = cached_result.get("from_cache", False)
                checks.append(("缓存功能", cache_working))
            except:
                checks.append(("缓存功能", False))

            # 测试统计功能
            try:
                stats = await optimizer.get_optimization_stats()
                stats_working = "total_optimizations" in stats
                checks.append(("统计功能", stats_working))
            except:
                checks.append(("统计功能", False))

            self.test_results[test_name] = {
                "status": "passed" if all(check[1] for check in checks) else "failed",
                "checks": checks,
                "total_tests": len(checks)
            }

            if self.verbose:
                for check_name, passed in checks:
                    status = "✅" if passed else "❌"
                    print(f"  {status} {check_name}")

        except Exception as e:
            self.test_results[test_name] = {
                "status": "error",
                "error": str(e),
                "checks": []
            }
            print(f"❌ 查询优化测试失败: {e}")

    async def test_workflow_management(self):
        """测试工作流管理功能"""
        print("\n⚡ 测试 ADDP 工作流...")

        test_name = "workflow_management"
        try:
            workflow_manager = WorkflowManager()

            # 测试各个阶段
            phases = ["analysis", "design", "development", "persistence"]
            checks = []

            # 模拟完整的工作流循环
            previous_output = ""
            for phase in phases:
                try:
                    result = await workflow_manager.execute_phase(
                        phase=phase,
                        specification="测试需求：实现用户注册功能",
                        previous_phase_output=previous_output
                    )

                    # 验证阶段执行结果
                    phase_passed = all([
                        result.get("phase") == phase,
                        result.get("status") == "completed",
                        "output" in result,
                        "quality_gates" in result
                    ])

                    checks.append((f"{phase}阶段", phase_passed))

                    # 准备下一阶段的输入
                    if phase_passed:
                        previous_output = json.dumps(result.get("output", {}))

                except Exception as e:
                    checks.append((f"{phase}阶段", False))
                    if self.verbose:
                        print(f"  ❌ {phase} 阶段失败: {e}")

            # 测试工作流状态
            try:
                status = await workflow_manager.get_workflow_status()
                status_working = "current_phase" in status or "status" in status
                checks.append(("状态查询", status_working))
            except:
                checks.append(("状态查询", False))

            self.test_results[test_name] = {
                "status": "passed" if all(check[1] for check in checks) else "failed",
                "checks": checks,
                "phases_tested": len(phases)
            }

            if self.verbose:
                for check_name, passed in checks:
                    status = "✅" if passed else "❌"
                    print(f"  {status} {check_name}")

        except Exception as e:
            self.test_results[test_name] = {
                "status": "error",
                "error": str(e),
                "checks": []
            }
            print(f"❌ 工作流管理测试失败: {e}")

    async def test_sync_management(self):
        """测试跨工具同步功能"""
        print("\n🔄 测试跨工具同步...")

        test_name = "sync_management"
        try:
            sync_manager = SyncManager()

            # 测试用例
            test_tools = ["claude", "gemini", "cursor"]
            checks = []

            # 测试状态保存和加载
            for tool in test_tools:
                try:
                    # 保存测试状态
                    test_state = {
                        "session_id": f"test_{tool}",
                        "timestamp": "2024-01-01T00:00:00",
                        "data": {"test": True}
                    }

                    save_result = await sync_manager.sync_state(
                        action="save",
                        tool_name=tool,
                        state_data=json.dumps(test_state)
                    )

                    save_success = save_result.get("success", False)
                    checks.append((f"{tool}保存", save_success))

                    if save_success:
                        # 测试状态加载
                        load_result = await sync_manager.sync_state(
                            action="load",
                            tool_name=tool
                        )

                        load_success = load_result.get("success", False)
                        checks.append((f"{tool}加载", load_success))

                except Exception as e:
                    checks.append((f"{tool}保存", False))
                    checks.append((f"{tool}加载", False))
                    if self.verbose:
                        print(f"  ❌ {tool} 同步失败: {e}")

            # 测试全工具同步
            try:
                sync_all_result = await sync_manager.sync_state(action="sync_all")
                sync_all_success = sync_all_result.get("success", False)
                checks.append(("全工具同步", sync_all_success))
            except:
                checks.append(("全工具同步", False))

            # 测试同步状态查询
            try:
                status = await sync_manager.get_sync_status()
                status_working = "supported_tools" in status and "tool_status" in status
                checks.append(("状态查询", status_working))
            except:
                checks.append(("状态查询", False))

            self.test_results[test_name] = {
                "status": "passed" if all(check[1] for check in checks) else "failed",
                "checks": checks,
                "tools_tested": len(test_tools)
            }

            if self.verbose:
                for check_name, passed in checks:
                    status = "✅" if passed else "❌"
                    print(f"  {status} {check_name}")

        except Exception as e:
            self.test_results[test_name] = {
                "status": "error",
                "error": str(e),
                "checks": []
            }
            print(f"❌ 跨工具同步测试失败: {e}")

    def generate_test_report(self):
        """生成测试报告"""
        print("\n📊 测试报告")
        print("=" * 50)

        total_tests = len(self.test_results)
        passed_tests = sum(1 for result in self.test_results.values()
                          if result["status"] == "passed")
        failed_tests = sum(1 for result in self.test_results.values()
                          if result["status"] == "failed")
        error_tests = sum(1 for result in self.test_results.values()
                         if result["status"] == "error")

        print(f"总测试数: {total_tests}")
        print(f"✅ 通过: {passed_tests}")
        print(f"❌ 失败: {failed_tests}")
        print(f"💥 错误: {error_tests}")
        print(f"成功率: {passed_tests/total_tests*100:.1f}%")

        # 详细结果
        for test_name, result in self.test_results.items():
            status_icon = {"passed": "✅", "failed": "❌", "error": "💥"}[result["status"]]
            print(f"\n{status_icon} {test_name}:")

            if result["status"] == "error":
                print(f"   错误: {result.get('error', 'Unknown error')}")
            else:
                checks = result.get("checks", [])
                for check_name, passed in checks:
                    check_icon = "✅" if passed else "❌"
                    print(f"   {check_icon} {check_name}")

        # 保存报告到文件
        if hasattr(self, 'temp_dir') and self.temp_dir:
            report_file = Path.cwd() / f"mcp_test_report_{asyncio.get_event_loop().time():.0f}.json"
            with open(report_file, 'w', encoding='utf-8') as f:
                json.dump(self.test_results, f, indent=2, ensure_ascii=False)
            print(f"\n📄 详细报告已保存到: {report_file}")

    async def test_specific_tool(self, tool_name: str):
        """测试特定工具"""
        print(f"🔧 测试特定工具: {tool_name}")

        if tool_name == "init":
            await self.test_project_initialization()
        elif tool_name == "optimize":
            await self.test_query_optimization()
        elif tool_name == "workflow":
            await self.test_workflow_management()
        elif tool_name == "sync":
            await self.test_sync_management()
        else:
            print(f"❌ 未知工具: {tool_name}")
            return

        self.generate_test_report()

def main():
    """主函数"""
    parser = argparse.ArgumentParser(description="MCP 工具测试脚本")
    parser.add_argument("--tool", type=str,
                       choices=["init", "optimize", "workflow", "sync"],
                       help="测试特定工具")
    parser.add_argument("--verbose", action="store_true",
                       help="详细输出")

    args = parser.parse_args()

    tester = MCPToolsTester(verbose=args.verbose)

    try:
        if args.tool:
            asyncio.run(tester.test_specific_tool(args.tool))
        else:
            results = asyncio.run(tester.run_all_tests())

            # 根据测试结果设置退出码
            if all(result["status"] == "passed" for result in results.values()):
                print("\n🎉 所有测试通过!")
                sys.exit(0)
            else:
                print("\n⚠️  部分测试失败，请检查报告")
                sys.exit(1)

    except KeyboardInterrupt:
        print("\n👋 测试被用户取消")
        sys.exit(0)
    except Exception as e:
        print(f"❌ 测试运行异常: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()