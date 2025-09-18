"""
项目初始化工具
==============

实现完整的 .addp 项目结构自动初始化功能，包括：
- 智能项目类型检测
- 目录结构创建
- 模板文件生成
- 配置文件设置
"""

import json
import os
from pathlib import Path
from typing import Dict, List, Any, Optional
from datetime import datetime
import logging

logger = logging.getLogger(__name__)

class ProjectInitializer:
    """项目初始化器"""

    def __init__(self):
        self.addp_structure = {
            "specifications": {
                "desc": "规格驱动文档",
                "subdirs": ["templates", "active", "archive", "reviews"]
            },
            "workflows": {
                "desc": "ADDP 四阶段工作流",
                "subdirs": ["analysis", "design", "development", "persistence"]
            },
            "memory": {
                "desc": "跨工具项目记忆",
                "subdirs": ["context", "decisions", "lessons", "sessions"]
            },
            "queries": {
                "desc": "Ollama 查询优化",
                "subdirs": ["optimized", "cache", "analytics", "feedback"]
            },
            "gates": {
                "desc": "质量门禁检查",
                "subdirs": ["constitution", "rules", "validations", "reports"]
            },
            "sync": {
                "desc": "工具状态同步",
                "subdirs": ["claude", "gemini", "cursor", "universal"]
            },
            "analytics": {
                "desc": "性能使用分析",
                "subdirs": ["metrics", "reports", "trends", "benchmarks"]
            },
            "experiments": {
                "desc": "A/B 测试研究",
                "subdirs": ["configs", "results", "comparisons", "insights"]
            },
            "configs": {
                "desc": "配置管理",
                "subdirs": ["mcp", "ollama", "tools", "templates"]
            },
            "cache": {
                "desc": "缓存优化",
                "subdirs": ["queries", "results", "models", "states"]
            }
        }

    async def initialize_structure(
        self,
        project_type: str = "universal-coding",
        project_name: str = "",
        framework: str = "auto-detect"
    ) -> Dict[str, Any]:
        """
        初始化完整的 .addp 项目结构

        Args:
            project_type: 项目类型
            project_name: 项目名称
            framework: 框架类型

        Returns:
            初始化结果统计
        """
        try:
            # 1. 检测项目信息
            project_info = await self._detect_project_info(project_name, framework)

            # 2. 创建目录结构
            directories_created = await self._create_directory_structure()

            # 3. 生成模板文件
            files_created = await self._create_template_files(project_info)

            # 4. 创建配置文件
            configs_created = await self._create_configuration_files(project_info)

            # 5. 生成项目元数据
            await self._create_project_metadata(project_info, project_type)

            # 6. 创建 Constitution 文件
            await self._create_constitution_files()

            # 7. 初始化 README
            await self._create_addp_readme()

            result = {
                "success": True,
                "project_info": project_info,
                "directories_created": directories_created,
                "files_created": files_created,
                "configs_created": configs_created,
                "summary": self._generate_summary(project_info, directories_created, files_created)
            }

            logger.info(f"项目结构初始化完成: {result['summary']}")
            return result

        except Exception as e:
            logger.error(f"项目初始化失败: {e}")
            raise

    async def _detect_project_info(self, project_name: str, framework: str) -> Dict[str, str]:
        """检测项目信息"""
        current_dir = Path.cwd()

        # 检测项目名称
        if not project_name:
            project_name = current_dir.name

        # 自动检测框架
        detected_framework = framework
        if framework == "auto-detect":
            if (current_dir / "package.json").exists():
                try:
                    with open(current_dir / "package.json", "r", encoding="utf-8") as f:
                        package_data = json.load(f)
                        dependencies = {**package_data.get("dependencies", {}), **package_data.get("devDependencies", {})}

                        if "react" in dependencies:
                            detected_framework = "react"
                        elif "vue" in dependencies:
                            detected_framework = "vue"
                        elif "angular" in dependencies:
                            detected_framework = "angular"
                        elif "next" in dependencies:
                            detected_framework = "nextjs"
                        else:
                            detected_framework = "nodejs"
                except:
                    detected_framework = "nodejs"

            elif (current_dir / "requirements.txt").exists() or (current_dir / "pyproject.toml").exists():
                detected_framework = "python"
            elif (current_dir / "build.gradle.kts").exists() or (current_dir / "build.gradle").exists():
                detected_framework = "kotlin"
            elif (current_dir / "go.mod").exists():
                detected_framework = "golang"
            elif (current_dir / "Cargo.toml").exists():
                detected_framework = "rust"
            else:
                detected_framework = "universal"

        return {
            "name": project_name,
            "framework": detected_framework,
            "path": str(current_dir),
            "initialized_at": datetime.now().isoformat()
        }

    async def _create_directory_structure(self) -> int:
        """创建目录结构"""
        directories_created = 0
        base_path = Path(".addp")

        # 创建根目录
        base_path.mkdir(exist_ok=True)
        directories_created += 1

        # 创建所有子目录
        for main_dir, config in self.addp_structure.items():
            main_path = base_path / main_dir
            main_path.mkdir(exist_ok=True)
            directories_created += 1

            # 创建子目录
            for subdir in config["subdirs"]:
                sub_path = main_path / subdir
                sub_path.mkdir(exist_ok=True)
                directories_created += 1

        return directories_created

    async def _create_template_files(self, project_info: Dict[str, str]) -> int:
        """创建模板文件"""
        files_created = 0
        base_path = Path(".addp")

        # 1. 规格模板
        await self._create_specification_templates(base_path, project_info)
        files_created += 4  # PRD, Plan, Tasks, ADR

        # 2. 工作流模板
        await self._create_workflow_templates(base_path, project_info)
        files_created += 4  # 四个阶段

        # 3. 记忆模板
        await self._create_memory_templates(base_path, project_info)
        files_created += 2  # context, decisions

        # 4. 查询模板
        await self._create_query_templates(base_path, project_info)
        files_created += 2  # optimization, cache

        return files_created

    async def _create_specification_templates(self, base_path: Path, project_info: Dict[str, str]):
        """创建规格文档模板"""
        templates_path = base_path / "specifications" / "templates"

        # PRD 模板
        prd_template = f"""# 产品需求文档 (PRD) 模板

## 项目信息
- **项目名称**: {project_info['name']}
- **技术栈**: {project_info['framework']}
- **创建时间**: {project_info['initialized_at']}

## 1. 项目概述
### 背景
[描述项目背景和动机]

### 目标
[明确项目要解决的问题]

### 成功标准
[定义项目成功的具体指标]

## 2. 功能需求
### 核心功能
- [ ] 功能1: [详细描述]
- [ ] 功能2: [详细描述]

### 约束条件
- **性能要求**: [具体指标]
- **兼容性要求**: [支持的平台/浏览器]
- **安全要求**: [安全标准]

## 3. 技术需求
### 技术栈
- **前端**: {project_info['framework']}
- **后端**: [选择的后端技术]
- **数据库**: [数据库选择]

### 架构要求
[系统架构描述]

## 4. 验收标准
### 功能验收
- [ ] 所有核心功能正常工作
- [ ] 性能指标达标
- [ ] 安全测试通过

### 质量门禁
- [ ] 代码覆盖率 > 80%
- [ ] 所有 lint 检查通过
- [ ] 所有单元测试通过

---
*此文档遵循规格驱动开发 (SDD) 原则*
"""

        with open(templates_path / "prd_template.md", "w", encoding="utf-8") as f:
            f.write(prd_template)

        # 技术方案模板
        plan_template = f"""# 技术实施方案模板

## 项目信息
- **项目名称**: {project_info['name']}
- **基于 PRD**: [关联的 PRD 文档]
- **方案版本**: v1.0

## 1. 技术方案概览
### 整体架构
[架构图和描述]

### 技术选择
| 层级 | 技术选择 | 理由 |
|------|----------|------|
| 前端 | {project_info['framework']} | [选择理由] |
| 后端 | [选择] | [理由] |
| 数据库 | [选择] | [理由] |

## 2. 实施计划
### 阶段1: 基础设施 (Week 1-2)
- [ ] 项目脚手架搭建
- [ ] 开发环境配置
- [ ] CI/CD 流水线

### 阶段2: 核心功能 (Week 3-6)
- [ ] 功能模块1
- [ ] 功能模块2
- [ ] 单元测试

### 阶段3: 集成测试 (Week 7-8)
- [ ] 集成测试
- [ ] 性能优化
- [ ] 文档完善

## 3. 风险评估
### 技术风险
- **风险1**: [描述] - 缓解策略: [策略]
- **风险2**: [描述] - 缓解策略: [策略]

### 进度风险
- **依赖风险**: [外部依赖分析]
- **资源风险**: [人力/时间分析]

## 4. 质量保证
### 开发标准
- TDD 先行开发
- 代码审查机制
- 自动化测试

### 部署策略
- 灰度发布
- 回滚机制
- 监控告警

---
*此方案基于 ADDP 框架设计*
"""

        with open(templates_path / "plan_template.md", "w", encoding="utf-8") as f:
            f.write(plan_template)

        # 任务分解模板
        tasks_template = """# 开发任务分解模板

## 任务概览
- **来源**: [对应的技术方案]
- **总工期**: [预估时间]
- **负责人**: [开发者]

## 任务分解

### Epic 1: 基础设施建设
#### Task 1.1: 项目初始化
- **描述**: 搭建项目基础结构
- **验收标准**:
  - [ ] 项目脚手架创建完成
  - [ ] 开发环境可正常启动
  - [ ] 基础依赖安装完成
- **预估工时**: 0.5 天
- **优先级**: P0 (阻塞)

#### Task 1.2: CI/CD 配置
- **描述**: 配置自动化构建和部署
- **验收标准**:
  - [ ] GitHub Actions 配置完成
  - [ ] 自动化测试流水线运行
  - [ ] 代码质量检查集成
- **预估工时**: 1 天
- **优先级**: P1 (重要)

### Epic 2: 核心功能开发
#### Task 2.1: [具体功能]
- **描述**: [详细功能描述]
- **验收标准**:
  - [ ] 功能实现完成
  - [ ] 单元测试覆盖率 > 80%
  - [ ] 集成测试通过
- **预估工时**: [时间]
- **优先级**: [P0/P1/P2]
- **依赖**: [前置任务]

## 进度跟踪

| 任务 | 状态 | 开始时间 | 完成时间 | 实际工时 | 备注 |
|------|------|----------|----------|----------|------|
| Task 1.1 | 🟡 进行中 | [日期] | - | - | - |
| Task 1.2 | ⚪ 待开始 | - | - | - | - |

## 风险跟踪
- **阻塞问题**: [当前阻塞点]
- **技术难点**: [需要攻克的技术点]
- **依赖等待**: [等待的外部依赖]

---
*任务遵循 TDD 和最小化修改原则*
"""

        with open(templates_path / "tasks_template.md", "w", encoding="utf-8") as f:
            f.write(tasks_template)

        # ADR 模板
        adr_template = """# 架构决策记录 (ADR) 模板

## ADR-001: [决策标题]

**状态**: 提议中 | 已接受 | 已废弃 | 已替代
**决策者**: [决策人员]
**决策日期**: [日期]
**技术故事**: [关联的需求或问题]

### 上下文
[描述促使此决策的情况和问题]

### 决策
[描述我们的反应，即我们选择的决策]

### 结果
[描述应用决策后的结果上下文]

### 合规性
此决策需要遵循以下约束：
- [ ] TDD 先行原则
- [ ] 反抽象原则 (避免过度抽象)
- [ ] 简化优先原则
- [ ] 集成优先测试

### 后果
**正面影响**:
- [积极后果]

**负面影响**:
- [消极后果]

**风险缓解**:
- [如何处理负面影响]

### 相关决策
- ADR-XXX: [相关决策]
- 替代方案: [被拒绝的其他选项]

---
*此 ADR 遵循 [MADR](https://adr.github.io/madr/) 格式*
"""

        with open(templates_path / "adr_template.md", "w", encoding="utf-8") as f:
            f.write(adr_template)

    async def _create_workflow_templates(self, base_path: Path, project_info: Dict[str, str]):
        """创建工作流模板"""
        for phase in ["analysis", "design", "development", "persistence"]:
            phase_path = base_path / "workflows" / phase
            template_content = f"""# ADDP {phase.title()} 阶段模板

## 阶段目标
{self._get_phase_objective(phase)}

## 输入要求
- 前一阶段的输出
- 相关上下文信息
- 约束条件

## 执行检查清单
{self._get_phase_checklist(phase)}

## 输出格式
- 阶段结果文档
- 下一阶段输入
- 风险和建议

## 质量门禁
{self._get_phase_gates(phase)}

---
*ADDP Framework v1.0*
"""

            with open(phase_path / f"{phase}_template.md", "w", encoding="utf-8") as f:
                f.write(template_content)

    async def _create_memory_templates(self, base_path: Path, project_info: Dict[str, str]):
        """创建记忆模板"""
        # 项目上下文模板
        context_template = f"""# 项目上下文记忆

## 项目基本信息
- **名称**: {project_info['name']}
- **类型**: {project_info['framework']}
- **路径**: {project_info['path']}
- **初始化**: {project_info['initialized_at']}

## 技术栈信息
- **前端**: {project_info['framework']}
- **后端**: [待确定]
- **数据库**: [待确定]
- **部署**: [待确定]

## 项目约束
- **性能要求**: [待定义]
- **安全要求**: [待定义]
- **兼容性**: [待定义]

## 当前状态
- **开发阶段**: 初始化
- **当前分支**: main
- **最后更新**: {project_info['initialized_at']}

## 关键决策
- [记录重要的技术决策]

## 学习记录
- [记录开发过程中的经验教训]

---
*此文件由 MCP 自动维护，记录项目演进历史*
"""

        context_path = base_path / "memory" / "context"
        with open(context_path / "project_context.json", "w", encoding="utf-8") as f:
            json.dump({
                "project_info": project_info,
                "technical_stack": {"frontend": project_info['framework']},
                "constraints": {},
                "current_state": {"phase": "initialized"},
                "decisions": [],
                "lessons": []
            }, f, indent=2, ensure_ascii=False)

    async def _create_query_templates(self, base_path: Path, project_info: Dict[str, str]):
        """创建查询优化模板"""
        optimization_config = {
            "ollama_endpoint": "http://localhost:11434",
            "model": "qwen2.5:14b",
            "optimization_levels": {
                "basic": "基础语法和术语优化",
                "smart": "上下文感知的智能优化",
                "detailed": "深度分析和多方案生成"
            },
            "prompt_templates": {
                "optimization": "请优化以下查询使其更加精确和可执行: {query}",
                "context_enhancement": "基于项目上下文 {context}，优化查询: {query}",
                "specification_generation": "将需求 '{query}' 转化为详细的技术规格"
            }
        }

        queries_path = base_path / "queries" / "optimized"
        with open(queries_path / "optimization_config.json", "w", encoding="utf-8") as f:
            json.dump(optimization_config, f, indent=2, ensure_ascii=False)

    async def _create_configuration_files(self, project_info: Dict[str, str]) -> int:
        """创建配置文件"""
        configs_created = 0
        configs_path = Path(".addp") / "configs"

        # MCP 配置
        mcp_config = {
            "server": {
                "name": "universal-coding-assistant",
                "version": "1.0.0",
                "description": "Universal AI Coding Framework MCP Server"
            },
            "tools": [
                {
                    "name": "initialize_addp_structure",
                    "description": "自动初始化 ADDP 项目结构"
                },
                {
                    "name": "optimize_query",
                    "description": "使用 Ollama 优化用户查询"
                },
                {
                    "name": "start_addp_workflow",
                    "description": "启动 ADDP 工作流阶段"
                },
                {
                    "name": "sync_project_state",
                    "description": "同步项目状态到所有工具"
                }
            ],
            "project_info": project_info
        }

        with open(configs_path / "mcp" / "server_config.json", "w", encoding="utf-8") as f:
            json.dump(mcp_config, f, indent=2, ensure_ascii=False)
        configs_created += 1

        # Ollama 配置
        ollama_config = {
            "endpoint": "http://localhost:11434",
            "model": "qwen2.5:14b",
            "temperature": 0.7,
            "max_tokens": 2048,
            "timeout": 30
        }

        with open(configs_path / "ollama" / "model_config.json", "w", encoding="utf-8") as f:
            json.dump(ollama_config, f, indent=2, ensure_ascii=False)
        configs_created += 1

        # 工具配置
        tools_config = {
            "claude_code": {
                "mcp_server": "universal-coding-assistant",
                "commands": ["initialize", "optimize", "workflow", "sync"]
            },
            "gemini_cli": {
                "mcp_server": "universal-coding-assistant",
                "integration": "mcp-protocol"
            },
            "cursor": {
                "mcp_config": ".cursor/mcp.json",
                "integration": "plugin"
            }
        }

        with open(configs_path / "tools" / "integration_config.json", "w", encoding="utf-8") as f:
            json.dump(tools_config, f, indent=2, ensure_ascii=False)
        configs_created += 1

        return configs_created

    async def _create_constitution_files(self):
        """创建 Constitution 质量门禁文件"""
        gates_path = Path(".addp") / "gates" / "constitution"

        constitution_content = """# Universal Coding Framework Constitution

## 核心原则 (借鉴 Spec-Kit)

### 1. TDD 先行原则
**规则**: 任何代码修改必须先写测试
**检查点**:
- [ ] 测试用例已编写
- [ ] 测试用例验证需求
- [ ] 测试用例可重现失败

**执行**:
```bash
# 每次开发前必须先有测试
npm test -- --watch
# 或
pytest --watch
```

### 2. 反抽象原则
**规则**: 避免过度抽象，优先具体实现
**检查点**:
- [ ] 代码直接解决问题，无过度抽象
- [ ] 重复代码在3次以上才考虑抽象
- [ ] 抽象层级不超过2层

### 3. 简化优先原则
**规则**: 选择最简单可行的解决方案
**检查点**:
- [ ] 方案易于理解和维护
- [ ] 依赖最少
- [ ] 认知复杂度最低

### 4. 集成优先原则
**规则**: 集成测试优先于单元测试
**检查点**:
- [ ] 端到端测试覆盖主要流程
- [ ] API 集成测试完整
- [ ] 用户场景测试覆盖

## 强制门禁

### 代码质量门禁
```bash
# 必须通过的检查
npm run lint          # 代码风格检查
npm run type-check    # 类型检查
npm run test          # 单元测试
npm run test:e2e      # 集成测试
npm run build         # 构建检查
```

### 性能门禁
- 构建时间 < 30秒
- 测试执行时间 < 5分钟
- 包大小增长 < 10%

### 安全门禁
- 无高危漏洞
- 无敏感信息泄露
- 依赖安全扫描通过

## 违反处理
- **违反 TDD**: 拒绝合并，要求补充测试
- **过度抽象**: 要求重构为具体实现
- **复杂方案**: 要求简化或提供简化理由
- **缺少集成测试**: 补充端到端测试

---
*此 Constitution 确保代码质量和开发纪律*
"""

        with open(gates_path / "constitution.md", "w", encoding="utf-8") as f:
            f.write(constitution_content)

        # 验证规则配置
        validation_rules = {
            "tdd_enforcement": {
                "enabled": True,
                "pre_commit_check": True,
                "test_coverage_threshold": 80
            },
            "anti_abstraction": {
                "enabled": True,
                "max_abstraction_layers": 2,
                "duplication_threshold": 3
            },
            "simplify_first": {
                "enabled": True,
                "complexity_threshold": 10,
                "dependency_limit": 20
            },
            "integration_priority": {
                "enabled": True,
                "e2e_coverage_threshold": 70,
                "api_test_required": True
            }
        }

        with open(gates_path / "validation_rules.json", "w", encoding="utf-8") as f:
            json.dump(validation_rules, f, indent=2, ensure_ascii=False)

    async def _create_project_metadata(self, project_info: Dict[str, str], project_type: str):
        """创建项目元数据"""
        metadata = {
            "version": "1.0.0",
            "framework_version": "1.0.0",
            "project_type": project_type,
            "project_info": project_info,
            "addp_structure_version": "1.0.0",
            "initialized_by": "Universal Coding Framework MCP",
            "initialization_date": datetime.now().isoformat(),
            "features": {
                "spec_driven_development": True,
                "addp_workflow": True,
                "cross_tool_sync": True,
                "query_optimization": True,
                "quality_gates": True
            }
        }

        with open(Path(".addp") / "metadata.json", "w", encoding="utf-8") as f:
            json.dump(metadata, f, indent=2, ensure_ascii=False)

    async def _create_addp_readme(self):
        """创建 .addp 目录说明文件"""
        readme_content = """# .addp 目录说明

这个目录包含了所有 Universal AI Coding Framework 的产出文件和配置。

## 📁 目录结构

```
.addp/
├── 📋 specifications/     # 规格驱动文档 (PRD/Plan/Tasks)
├── 🔄 workflows/         # ADDP 四阶段工作流产出
├── 🧠 memory/            # 跨工具项目记忆同步
├── 🔍 queries/           # Ollama 查询优化缓存
├── ⚡ gates/             # 质量门禁检查规则
├── 🔄 sync/              # 工具状态同步数据
├── 📊 analytics/         # 性能使用分析报告
├── 🧪 experiments/       # A/B 测试配置结果
├── 🔧 configs/           # MCP/Ollama 配置文件
└── 🗃️ cache/             # 缓存优化数据
```

## 🚀 使用指南

### 开始新需求
```bash
claude "/specify 你的需求描述"
# 自动生成 specifications/active/ 下的 PRD
```

### 执行开发计划
```bash
claude "/plan"
# 基于 PRD 生成技术方案到 specifications/active/
```

### 启动 ADDP 工作流
```bash
claude "/workflow analysis"
# 启动四阶段工作流，产出保存到 workflows/
```

### 跨工具同步
```bash
claude "同步项目状态"
# 将当前状态同步到 sync/ 目录，供其他工具使用
```

## 🔄 跨工具支持

此目录结构被以下工具共享：
- **Claude Code**: 原生 MCP 支援
- **Gemini CLI**: 完整 MCP 支援
- **Cursor**: 通过 MCP 配置
- **其他工具**: 通过 MCP 协议

## 📜 文件管理

- **不要手动修改** `cache/` 和 `sync/` 目录
- **可以编辑** `specifications/` 和 `configs/` 文件
- **建议备份** 重要的规格文档和配置
- **定期清理** 过期的缓存和实验数据

## 🆘 故障排除

如果遇到问题：
1. 检查 `configs/mcp/server_config.json` 配置
2. 验证 Ollama 服务是否运行 (`ollama serve`)
3. 查看 `analytics/` 目录的错误日志
4. 重新初始化: `claude "重新初始化 ADDP 结构"`

---
*此目录由 Universal AI Coding Framework 自动管理*
"""

        with open(Path(".addp") / "README.md", "w", encoding="utf-8") as f:
            f.write(readme_content)

    def _get_phase_objective(self, phase: str) -> str:
        """获取阶段目标"""
        objectives = {
            "analysis": "深入分析需求，识别技术约束和风险，为设计阶段提供清晰的输入",
            "design": "基于分析结果设计技术方案，权衡各种选择，输出详细的实施计划",
            "development": "按照 TDD 原则执行开发，保持最小化修改，确保质量",
            "persistence": "验证开发结果，更新项目记忆，为下一轮迭代做准备"
        }
        return objectives.get(phase, "执行指定的工作流阶段")

    def _get_phase_checklist(self, phase: str) -> str:
        """获取阶段检查清单"""
        checklists = {
            "analysis": """- [ ] 需求已澄清和确认
- [ ] 技术约束已识别
- [ ] 风险评估已完成
- [ ] 影响分析已进行
- [ ] 输入数据已验证""",
            "design": """- [ ] 技术方案已选择
- [ ] 架构设计已完成
- [ ] 实施计划已制定
- [ ] 风险缓解策略已定义
- [ ] 设计评审已通过""",
            "development": """- [ ] 测试用例已编写 (TDD)
- [ ] 代码实现已完成
- [ ] 单元测试已通过
- [ ] 代码审查已完成
- [ ] 集成测试已执行""",
            "persistence": """- [ ] 功能验证已完成
- [ ] 性能指标已确认
- [ ] 文档已更新
- [ ] 项目记忆已同步
- [ ] 经验教训已记录"""
        }
        return checklists.get(phase, "- [ ] 阶段任务已完成")

    def _get_phase_gates(self, phase: str) -> str:
        """获取阶段质量门禁"""
        gates = {
            "analysis": """- 需求清晰度检查
- 技术可行性验证
- 风险可控性评估""",
            "design": """- 方案完整性检查
- 架构合理性验证
- 实施可行性评估""",
            "development": """- TDD 流程合规检查
- 代码质量标准验证
- 测试覆盖率要求""",
            "persistence": """- 功能完整性验证
- 性能指标达标检查
- 文档完整性评估"""
        }
        return gates.get(phase, "基础质量门禁检查")

    def _generate_summary(self, project_info: Dict[str, str], directories: int, files: int) -> str:
        """生成初始化摘要"""
        return f"""✅ 项目 '{project_info['name']}' 初始化完成
🔧 检测到技术栈: {project_info['framework']}
📁 创建了 {directories} 个目录
📄 生成了 {files} 个模板文件
🎯 已配置规格驱动开发和 ADDP 工作流
🔄 已启用跨工具状态同步功能"""