"""
ADDP 工作流管理器
================

实现 Analysis → Design → Development → Persistence 四阶段工作流的自动化管理。
每个阶段都有明确的输入、处理和输出规范，确保开发过程的纪律性和质量。
"""

import json
import logging
from pathlib import Path
from typing import Dict, List, Any, Optional
from datetime import datetime
from enum import Enum

logger = logging.getLogger(__name__)

class WorkflowPhase(Enum):
    """工作流阶段枚举"""
    ANALYSIS = "analysis"
    DESIGN = "design"
    DEVELOPMENT = "development"
    PERSISTENCE = "persistence"

class WorkflowManager:
    """ADDP 工作流管理器"""

    def __init__(self):
        self.workflows_path = Path(".addp/workflows")
        self.workflows_path.mkdir(parents=True, exist_ok=True)

        # 确保各阶段目录存在
        for phase in WorkflowPhase:
            (self.workflows_path / phase.value).mkdir(exist_ok=True)

    async def execute_phase(
        self,
        phase: str,
        specification: str = "",
        previous_phase_output: str = ""
    ) -> Dict[str, Any]:
        """
        执行指定的工作流阶段

        Args:
            phase: 工作流阶段名称
            specification: 规格文档内容
            previous_phase_output: 前一阶段的输出

        Returns:
            阶段执行结果
        """
        try:
            # 验证阶段名称
            try:
                workflow_phase = WorkflowPhase(phase)
            except ValueError:
                raise Exception(f"无效的工作流阶段: {phase}")

            # 创建阶段上下文
            phase_context = await self._create_phase_context(
                workflow_phase, specification, previous_phase_output
            )

            # 执行阶段处理
            result = await self._execute_phase_logic(workflow_phase, phase_context)

            # 保存阶段输出
            await self._save_phase_output(workflow_phase, result)

            # 更新工作流状态
            await self._update_workflow_status(workflow_phase, result)

            return result

        except Exception as e:
            logger.error(f"工作流阶段 {phase} 执行失败: {e}")
            raise

    async def _create_phase_context(
        self,
        phase: WorkflowPhase,
        specification: str,
        previous_output: str
    ) -> Dict[str, Any]:
        """创建阶段执行上下文"""

        # 加载项目记忆
        project_memory = await self._load_project_memory()

        # 加载相关规格文档
        specifications = await self._load_specifications()

        # 构建阶段上下文
        context = {
            "phase": phase.value,
            "timestamp": datetime.now().isoformat(),
            "specification": specification,
            "previous_output": previous_output,
            "project_memory": project_memory,
            "specifications": specifications,
            "phase_requirements": self._get_phase_requirements(phase),
            "quality_gates": await self._load_quality_gates(phase)
        }

        return context

    async def _execute_phase_logic(
        self,
        phase: WorkflowPhase,
        context: Dict[str, Any]
    ) -> Dict[str, Any]:
        """执行具体的阶段逻辑"""

        if phase == WorkflowPhase.ANALYSIS:
            return await self._execute_analysis_phase(context)
        elif phase == WorkflowPhase.DESIGN:
            return await self._execute_design_phase(context)
        elif phase == WorkflowPhase.DEVELOPMENT:
            return await self._execute_development_phase(context)
        elif phase == WorkflowPhase.PERSISTENCE:
            return await self._execute_persistence_phase(context)
        else:
            raise Exception(f"未实现的阶段逻辑: {phase}")

    async def _execute_analysis_phase(self, context: Dict[str, Any]) -> Dict[str, Any]:
        """执行分析阶段"""

        analysis_tasks = [
            "需求澄清和验证",
            "技术约束识别",
            "风险评估分析",
            "影响范围评估",
            "资源需求评估"
        ]

        # 分析模板
        analysis_template = {
            "phase": "analysis",
            "input": {
                "specification": context["specification"],
                "previous_output": context["previous_output"]
            },
            "tasks": analysis_tasks,
            "outputs": {
                "requirement_clarification": "需求澄清结果",
                "technical_constraints": "技术约束清单",
                "risk_assessment": "风险评估报告",
                "impact_analysis": "影响分析报告",
                "resource_requirements": "资源需求评估"
            },
            "quality_checks": [
                "需求是否清晰明确",
                "技术约束是否完整",
                "风险是否充分识别",
                "影响范围是否准确"
            ],
            "next_phase_input": "为设计阶段准备的详细需求和约束文档"
        }

        # 执行质量门禁检查
        gate_results = await self._check_quality_gates(
            WorkflowPhase.ANALYSIS, analysis_template, context
        )

        result = {
            "phase": "analysis",
            "status": "completed",
            "output": analysis_template,
            "quality_gates": gate_results,
            "timestamp": datetime.now().isoformat(),
            "duration": "自动执行",
            "next_phase": "design"
        }

        return result

    async def _execute_design_phase(self, context: Dict[str, Any]) -> Dict[str, Any]:
        """执行设计阶段"""

        design_tasks = [
            "架构方案设计",
            "技术方案选择",
            "接口设计定义",
            "数据模型设计",
            "实施计划制定"
        ]

        design_template = {
            "phase": "design",
            "input": {
                "analysis_output": context["previous_output"],
                "requirements": context["specification"]
            },
            "tasks": design_tasks,
            "outputs": {
                "architecture_design": "系统架构设计",
                "technical_decisions": "技术方案决策",
                "interface_specifications": "接口规格说明",
                "data_models": "数据模型设计",
                "implementation_plan": "详细实施计划"
            },
            "quality_checks": [
                "架构是否合理可行",
                "技术选择是否恰当",
                "接口设计是否完整",
                "实施计划是否可执行"
            ],
            "design_principles": [
                "遵循 TDD 先行原则",
                "避免过度抽象设计",
                "优先选择简单方案",
                "确保集成测试优先"
            ],
            "next_phase_input": "可直接执行的开发计划和技术规格"
        }

        # 验证设计原则合规性
        design_compliance = await self._check_design_compliance(design_template, context)

        gate_results = await self._check_quality_gates(
            WorkflowPhase.DESIGN, design_template, context
        )

        result = {
            "phase": "design",
            "status": "completed",
            "output": design_template,
            "design_compliance": design_compliance,
            "quality_gates": gate_results,
            "timestamp": datetime.now().isoformat(),
            "duration": "自动执行",
            "next_phase": "development"
        }

        return result

    async def _execute_development_phase(self, context: Dict[str, Any]) -> Dict[str, Any]:
        """执行开发阶段"""

        development_tasks = [
            "TDD 测试用例编写",
            "核心功能实现",
            "单元测试执行",
            "集成测试验证",
            "代码质量检查"
        ]

        development_template = {
            "phase": "development",
            "input": {
                "design_output": context["previous_output"],
                "implementation_plan": "设计阶段的实施计划"
            },
            "tasks": development_tasks,
            "tdd_workflow": [
                "1. 编写失败测试",
                "2. 实现最小代码使测试通过",
                "3. 重构优化代码",
                "4. 重复循环直到完成"
            ],
            "development_rules": [
                "测试先行，代码后行",
                "最小化文件修改（≤3个文件）",
                "每次提交都要通过所有测试",
                "代码审查必须通过"
            ],
            "outputs": {
                "test_cases": "TDD 测试用例",
                "implementation_code": "功能实现代码",
                "test_results": "测试执行结果",
                "code_quality_report": "代码质量报告"
            },
            "quality_checks": [
                "所有测试是否通过",
                "代码覆盖率是否达标",
                "代码质量是否符合标准",
                "是否遵循 TDD 流程"
            ],
            "next_phase_input": "完整的功能实现和测试验证结果"
        }

        # TDD 合规性检查
        tdd_compliance = await self._check_tdd_compliance(development_template, context)

        gate_results = await self._check_quality_gates(
            WorkflowPhase.DEVELOPMENT, development_template, context
        )

        result = {
            "phase": "development",
            "status": "completed",
            "output": development_template,
            "tdd_compliance": tdd_compliance,
            "quality_gates": gate_results,
            "timestamp": datetime.now().isoformat(),
            "duration": "自动执行",
            "next_phase": "persistence"
        }

        return result

    async def _execute_persistence_phase(self, context: Dict[str, Any]) -> Dict[str, Any]:
        """执行持久化阶段"""

        persistence_tasks = [
            "功能验证确认",
            "性能指标检查",
            "项目记忆更新",
            "经验教训记录",
            "状态同步执行"
        ]

        persistence_template = {
            "phase": "persistence",
            "input": {
                "development_output": context["previous_output"],
                "complete_implementation": "开发阶段的完整成果"
            },
            "tasks": persistence_tasks,
            "validation_checks": [
                "功能是否完全实现",
                "性能是否达到预期",
                "质量标准是否满足",
                "文档是否完整更新"
            ],
            "outputs": {
                "validation_results": "功能验证结果",
                "performance_metrics": "性能指标报告",
                "updated_memory": "更新的项目记忆",
                "lessons_learned": "经验教训记录",
                "sync_status": "跨工具同步状态"
            },
            "memory_updates": [
                "记录技术决策和理由",
                "保存成功的实践模式",
                "记录遇到的问题和解决方案",
                "更新项目上下文信息"
            ],
            "quality_checks": [
                "验证结果是否符合预期",
                "性能指标是否达标",
                "记忆更新是否完整",
                "同步状态是否正常"
            ],
            "workflow_completion": "ADDP 工作流循环完成，可开始下一轮迭代"
        }

        # 执行记忆更新
        memory_update_result = await self._update_project_memory(persistence_template, context)

        # 执行跨工具同步
        sync_result = await self._sync_workflow_state(persistence_template, context)

        gate_results = await self._check_quality_gates(
            WorkflowPhase.PERSISTENCE, persistence_template, context
        )

        result = {
            "phase": "persistence",
            "status": "completed",
            "output": persistence_template,
            "memory_update": memory_update_result,
            "sync_result": sync_result,
            "quality_gates": gate_results,
            "timestamp": datetime.now().isoformat(),
            "duration": "自动执行",
            "workflow_status": "完成，可开始新循环"
        }

        return result

    async def _check_quality_gates(
        self,
        phase: WorkflowPhase,
        output: Dict[str, Any],
        context: Dict[str, Any]
    ) -> Dict[str, Any]:
        """执行质量门禁检查"""

        gates = context.get("quality_gates", {})
        gate_results = {
            "phase": phase.value,
            "total_checks": 0,
            "passed_checks": 0,
            "failed_checks": [],
            "warnings": [],
            "overall_status": "unknown"
        }

        # 基础质量检查
        basic_checks = output.get("quality_checks", [])
        gate_results["total_checks"] = len(basic_checks)

        for check in basic_checks:
            # 这里实现具体的检查逻辑
            # 在实际实现中，这些检查会调用具体的验证函数
            check_passed = True  # 占位符，实际需要实现检查逻辑

            if check_passed:
                gate_results["passed_checks"] += 1
            else:
                gate_results["failed_checks"].append(check)

        # 计算通过率
        pass_rate = gate_results["passed_checks"] / max(gate_results["total_checks"], 1)

        if pass_rate >= 0.9:
            gate_results["overall_status"] = "passed"
        elif pass_rate >= 0.7:
            gate_results["overall_status"] = "warning"
        else:
            gate_results["overall_status"] = "failed"

        return gate_results

    async def _check_design_compliance(
        self,
        design_output: Dict[str, Any],
        context: Dict[str, Any]
    ) -> Dict[str, Any]:
        """检查设计原则合规性"""

        principles = design_output.get("design_principles", [])
        compliance_result = {
            "total_principles": len(principles),
            "compliant_principles": 0,
            "violations": [],
            "recommendations": []
        }

        # 检查各项设计原则
        for principle in principles:
            is_compliant = True  # 占位符，实际需要实现检查逻辑

            if is_compliant:
                compliance_result["compliant_principles"] += 1
            else:
                compliance_result["violations"].append(principle)

        # 生成建议
        if compliance_result["violations"]:
            compliance_result["recommendations"] = [
                "请重新评估设计方案，确保符合所有设计原则",
                "考虑简化设计，避免过度复杂的抽象",
                "确保 TDD 流程可以有效执行"
            ]

        return compliance_result

    async def _check_tdd_compliance(
        self,
        development_output: Dict[str, Any],
        context: Dict[str, Any]
    ) -> Dict[str, Any]:
        """检查 TDD 合规性"""

        tdd_workflow = development_output.get("tdd_workflow", [])
        rules = development_output.get("development_rules", [])

        compliance_result = {
            "tdd_workflow_followed": True,  # 占位符
            "rules_compliance": {},
            "test_coverage": 0.0,  # 需要实际测试覆盖率
            "violations": [],
            "recommendations": []
        }

        # 检查开发规则合规性
        for rule in rules:
            # 这里实现具体的规则检查
            compliance_result["rules_compliance"][rule] = True  # 占位符

        return compliance_result

    async def _load_project_memory(self) -> Dict[str, Any]:
        """加载项目记忆"""
        try:
            memory_path = Path(".addp/memory/context/project_context.json")
            if memory_path.exists():
                with open(memory_path, 'r', encoding='utf-8') as f:
                    return json.load(f)
        except Exception as e:
            logger.warning(f"加载项目记忆失败: {e}")

        return {}

    async def _load_specifications(self) -> Dict[str, Any]:
        """加载规格文档"""
        try:
            specs_path = Path(".addp/specifications/active")
            specifications = {}

            if specs_path.exists():
                for spec_file in specs_path.glob("*.md"):
                    with open(spec_file, 'r', encoding='utf-8') as f:
                        specifications[spec_file.stem] = f.read()

            return specifications
        except Exception as e:
            logger.warning(f"加载规格文档失败: {e}")
            return {}

    async def _load_quality_gates(self, phase: WorkflowPhase) -> Dict[str, Any]:
        """加载质量门禁配置"""
        try:
            gates_path = Path(".addp/gates/constitution/validation_rules.json")
            if gates_path.exists():
                with open(gates_path, 'r', encoding='utf-8') as f:
                    return json.load(f)
        except Exception as e:
            logger.warning(f"加载质量门禁失败: {e}")

        return {}

    def _get_phase_requirements(self, phase: WorkflowPhase) -> Dict[str, Any]:
        """获取阶段要求"""
        requirements = {
            WorkflowPhase.ANALYSIS: {
                "required_inputs": ["用户需求", "项目上下文"],
                "expected_outputs": ["需求澄清", "技术约束", "风险评估"],
                "quality_criteria": ["需求明确性", "约束完整性", "风险覆盖度"]
            },
            WorkflowPhase.DESIGN: {
                "required_inputs": ["分析结果", "技术约束"],
                "expected_outputs": ["架构设计", "技术方案", "实施计划"],
                "quality_criteria": ["架构合理性", "方案可行性", "计划可执行性"]
            },
            WorkflowPhase.DEVELOPMENT: {
                "required_inputs": ["设计方案", "实施计划"],
                "expected_outputs": ["测试用例", "功能实现", "质量报告"],
                "quality_criteria": ["TDD 合规性", "测试覆盖率", "代码质量"]
            },
            WorkflowPhase.PERSISTENCE: {
                "required_inputs": ["开发结果", "测试报告"],
                "expected_outputs": ["验证结果", "记忆更新", "同步状态"],
                "quality_criteria": ["功能完整性", "性能达标", "文档完整性"]
            }
        }

        return requirements.get(phase, {})

    async def _save_phase_output(self, phase: WorkflowPhase, result: Dict[str, Any]):
        """保存阶段输出"""
        try:
            phase_path = self.workflows_path / phase.value
            timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
            output_file = phase_path / f"{phase.value}_output_{timestamp}.json"

            with open(output_file, 'w', encoding='utf-8') as f:
                json.dump(result, f, indent=2, ensure_ascii=False)

        except Exception as e:
            logger.error(f"保存阶段输出失败: {e}")

    async def _update_workflow_status(self, phase: WorkflowPhase, result: Dict[str, Any]):
        """更新工作流状态"""
        try:
            status_file = self.workflows_path / "workflow_status.json"

            # 加载现有状态
            status = {}
            if status_file.exists():
                with open(status_file, 'r', encoding='utf-8') as f:
                    status = json.load(f)

            # 更新状态
            status.update({
                "current_phase": phase.value,
                "last_update": datetime.now().isoformat(),
                "phase_status": result.get("status", "unknown"),
                "next_phase": result.get("next_phase", ""),
                "workflow_history": status.get("workflow_history", [])
            })

            # 添加历史记录
            status["workflow_history"].append({
                "phase": phase.value,
                "timestamp": result.get("timestamp"),
                "status": result.get("status"),
                "duration": result.get("duration")
            })

            # 保存更新的状态
            with open(status_file, 'w', encoding='utf-8') as f:
                json.dump(status, f, indent=2, ensure_ascii=False)

        except Exception as e:
            logger.error(f"更新工作流状态失败: {e}")

    async def _update_project_memory(
        self,
        persistence_output: Dict[str, Any],
        context: Dict[str, Any]
    ) -> Dict[str, Any]:
        """更新项目记忆"""
        try:
            memory_path = Path(".addp/memory/context/project_context.json")

            # 加载现有记忆
            memory = {}
            if memory_path.exists():
                with open(memory_path, 'r', encoding='utf-8') as f:
                    memory = json.load(f)

            # 更新记忆
            memory.update({
                "last_workflow_completion": datetime.now().isoformat(),
                "completed_phases": memory.get("completed_phases", []) + ["full_addp_cycle"],
                "lessons_learned": memory.get("lessons_learned", []) + [
                    "完成了完整的 ADDP 工作流循环"
                ]
            })

            # 保存更新的记忆
            with open(memory_path, 'w', encoding='utf-8') as f:
                json.dump(memory, f, indent=2, ensure_ascii=False)

            return {"status": "success", "message": "项目记忆更新完成"}

        except Exception as e:
            logger.error(f"项目记忆更新失败: {e}")
            return {"status": "failed", "error": str(e)}

    async def _sync_workflow_state(
        self,
        persistence_output: Dict[str, Any],
        context: Dict[str, Any]
    ) -> Dict[str, Any]:
        """同步工作流状态到跨工具目录"""
        try:
            sync_path = Path(".addp/sync/universal")
            sync_path.mkdir(parents=True, exist_ok=True)

            sync_data = {
                "workflow_completed": True,
                "completion_time": datetime.now().isoformat(),
                "completed_phases": ["analysis", "design", "development", "persistence"],
                "ready_for_next_iteration": True,
                "sync_targets": ["claude", "gemini", "cursor"]
            }

            # 保存到各工具同步目录
            for tool in sync_data["sync_targets"]:
                tool_sync_path = Path(f".addp/sync/{tool}")
                tool_sync_path.mkdir(parents=True, exist_ok=True)

                with open(tool_sync_path / "workflow_state.json", 'w', encoding='utf-8') as f:
                    json.dump(sync_data, f, indent=2, ensure_ascii=False)

            return {"status": "success", "message": "工作流状态同步完成"}

        except Exception as e:
            logger.error(f"工作流状态同步失败: {e}")
            return {"status": "failed", "error": str(e)}

    async def get_workflow_status(self) -> Dict[str, Any]:
        """获取当前工作流状态"""
        try:
            status_file = self.workflows_path / "workflow_status.json"

            if status_file.exists():
                with open(status_file, 'r', encoding='utf-8') as f:
                    return json.load(f)
            else:
                return {
                    "status": "not_initialized",
                    "message": "工作流尚未开始"
                }

        except Exception as e:
            logger.error(f"获取工作流状态失败: {e}")
            return {"error": str(e)}