# TARGET.md — MCP + Ollama 統一架構實施計劃

**核心使命**：通過 MCP + Ollama 雙層架構解決 AI coding 工具的三大痛點：多工具割裂、提問低效、缺乏紀律性。

## 🎯 核心範圍（Scope）

### 雙層架構設計
```
用戶輸入 → Ollama優化層 → MCP統一服務層 → AI工具執行層 → 統一回饋
```

### 最小可行工具集 (MVP) - 集成 Spec-Kit 理念
```yaml
規格驅動層 (借鑒 GitHub Spec-Kit):
  - spec.create_prd          # 創建產品需求文檔 (/specify)
  - spec.generate_plan       # 生成技術實施方案 (/plan)
  - spec.decompose_tasks     # 分解為可執行任務 (/tasks)
  - spec.validate_gates      # 驗證規格門禁條件

Ollama優化層:
  - query.optimize           # 將模糊輸入轉換為精準技術指令
  - query.analyze_intent     # 深度意圖分析和需求澄清
  - query.enhance_context    # 基於規格文檔增強上下文

MCP統一服務層 (增強 ADDP工作流程):
  - addp.parse_specification # 解析規格文檔為執行計劃
  - addp.start_analysis      # 開始需求分析階段
  - addp.define_scope        # 定義項目範圍和邊界
  - addp.enforce_tdd         # 強制 TDD 開發流程
  - addp.validate_incremental # 增量開發驗證
  - addp.update_memory       # 項目記憶更新

質量門禁層 (Spec-Kit Constitution):
  - gate.check_tdd_first     # TDD 先行門禁檢查
  - gate.check_anti_abstract # 反抽象門禁檢查
  - gate.check_simplify_first # 簡化優先門禁檢查
  - gate.validate_constraints # 企業約束門禁檢查

跨工具同步層:
  - memory.sync_context      # 跨工具上下文同步
  - memory.save_session      # 會話狀態保存
  - memory.save_specification # 保存規格文檔和計劃
  - quality.validate_workflow # 工作流程合規檢查
```

### 支援工具矩陣
- **完全支援**：Claude Code, Gemini CLI (原生 MCP 支援)
- **部分支援**：Cursor (透過 MCP 配置)
- **計劃支援**：Codex, Aider

## 📂 完整的 .addp 目錄規劃

### 🎯 設計原則
- **MCP 自動初始化** - 通過 `project.initialize` 命令一鍵創建完整結構
- **功能導向分類** - 8大功能模組，每個對應特定 MCP 工具組
- **全生命週期管理** - active → archive 動態歸檔，自動化清理
- **跨工具統一** - 所有 AI 工具的產出都統一到 .addp 目錄

### 🗂️ 完整目錄結構

```
claude/
├── .addp/                         # MCP 統一產出目錄 (自動初始化)
│   ├── 📋 specifications/         # 規格驅動產出 (借鑒 Spec-Kit)
│   │   ├── templates/            # 規格模板
│   │   │   ├── prd.template.md    # 產品需求文檔模板
│   │   │   ├── plan.template.md   # 技術方案模板
│   │   │   ├── tasks.template.md  # 任務分解模板
│   │   │   └── adr.template.md    # 架構決策記錄模板
│   │   ├── active/              # 當前活跃規格
│   │   │   ├── current-prd.md    # 當前產品需求文檔
│   │   │   ├── current-plan.md   # 當前技術實施方案
│   │   │   ├── current-tasks.md  # 當前任務清單
│   │   │   └── session-context.json  # 當前會話上下文
│   │   ├── archive/             # 歷史規格存檔
│   │   │   ├── 2025-01-15/      # 按日期歸檔
│   │   │   │   ├── prd-v1.0.md
│   │   │   │   ├── plan-v1.0.md
│   │   │   │   └── completion-report.md
│   │   │   └── 2025-01-20/
│   │   └── constitution/        # 開發憲章和門禁
│   │       ├── CONSTITUTION.md   # 核心開發原則
│   │       ├── gates.yaml       # 質量門禁配置
│   │       ├── constraints.yaml # 企業約束配置
│   │       └── validation-rules.yaml # 驗證規則
│   ├── 🔄 workflows/             # ADDP 四階段工作流程產出
│   │   ├── analysis/            # Analysis 階段產出
│   │   │   ├── requirements-analysis.md  # 需求分析結果
│   │   │   ├── scope-definition.md      # 範圍界定
│   │   │   ├── risk-assessment.md       # 風險評估
│   │   │   └── context-analysis.json   # 上下文分析數據
│   │   ├── design/              # Design 階段產出
│   │   │   ├── architecture-analysis.md    # 架構分析
│   │   │   ├── solution-alternatives.md   # 方案對比
│   │   │   ├── implementation-plan.md     # 實施計劃
│   │   │   └── design-decisions.json     # 設計決策記錄
│   │   ├── development/         # Development 階段產出
│   │   │   ├── tdd-checklist.md         # TDD 檢查清單
│   │   │   ├── test-plans/              # 測試計劃
│   │   │   │   ├── unit-tests.md
│   │   │   │   ├── integration-tests.md
│   │   │   │   └── e2e-tests.md
│   │   │   ├── code-changes/            # 代碼變更記錄
│   │   │   │   ├── change-log.md
│   │   │   │   └── diff-summary.json
│   │   │   └── validation-reports/      # 增量驗證報告
│   │   └── persistence/         # Persistence 階段產出
│   │       ├── completion-report.md     # 完成報告
│   │       ├── quality-metrics.json    # 質量指標
│   │       ├── test-results/           # 測試結果
│   │       │   ├── unit-test-report.xml
│   │       │   ├── coverage-report.html
│   │       │   └── performance-report.json
│   │       └── lessons-learned.md      # 經驗總結
│   ├── 🧠 memory/                # 統一記憶系統
│   │   ├── project-context/      # 項目上下文
│   │   │   ├── PROJECT_CONTEXT.md     # 項目總體上下文
│   │   │   ├── DECISIONS.md           # 重要決策記錄
│   │   │   ├── ARCHITECTURE.md        # 架構演進史
│   │   │   └── CONSTRAINTS.md         # 項目約束記錄
│   │   ├── sessions/            # 會話記錄
│   │   │   ├── 2025-01-15-session.yml    # 具體會話狀態
│   │   │   ├── 2025-01-20-session.yml
│   │   │   └── last-session.yml          # 最近會話快照
│   │   ├── cross-tool/          # 跨工具同步
│   │   │   ├── claude-code-state.json    # Claude Code 狀態
│   │   │   ├── gemini-cli-state.json     # Gemini CLI 狀態
│   │   │   ├── cursor-state.json         # Cursor 狀態
│   │   │   └── sync-manifest.json        # 同步清單
│   │   └── knowledge-base/      # 知識庫
│   │       ├── patterns/              # 常用模式
│   │       ├── solutions/             # 解決方案庫
│   │       └── best-practices/        # 最佳實踐
│   ├── 🔍 queries/               # 查詢優化產出
│   │   ├── raw-queries/         # 原始查詢記錄
│   │   │   ├── 2025-01-15.log
│   │   │   └── 2025-01-20.log
│   │   ├── optimized-queries/   # 優化後查詢
│   │   │   ├── performance-optimization.md
│   │   │   ├── feature-implementation.md
│   │   │   └── bug-fix-queries.md
│   │   ├── context-enhancements/ # 上下文增強結果
│   │   │   ├── tech-stack-analysis.json
│   │   │   ├── dependency-analysis.json
│   │   │   └── constraint-analysis.json
│   │   └── confidence-scores/   # 置信度評分
│   │       ├── daily-scores.json
│   │       └── accuracy-metrics.json
│   ├── ⚡ gates/                 # 質量門禁產出
│   │   ├── validations/         # 門禁驗證結果
│   │   │   ├── tdd-first-check.md        # TDD 先行檢查
│   │   │   ├── anti-abstract-check.md    # 反抽象檢查
│   │   │   ├── simplify-first-check.md   # 簡化優先檢查
│   │   │   └── integration-check.md      # 集成優先檢查
│   │   ├── violations/          # 違規記錄
│   │   │   ├── 2025-01-15-violations.json
│   │   │   └── resolution-actions.md
│   │   ├── metrics/             # 質量指標
│   │   │   ├── gate-pass-rates.json     # 門禁通過率
│   │   │   ├── quality-trends.json      # 質量趨勢
│   │   │   └── improvement-suggestions.md
│   │   └── reports/             # 質量報告
│   │       ├── weekly-quality-report.md
│   │       └── monthly-analysis.md
│   ├── 🔄 sync/                  # 跨工具同步產出
│   │   ├── state-snapshots/     # 狀態快照
│   │   │   ├── pre-sync-snapshot.json
│   │   │   └── post-sync-snapshot.json
│   │   ├── sync-logs/           # 同步日誌
│   │   │   ├── 2025-01-15-sync.log
│   │   │   └── error-logs/
│   │   ├── conflict-resolution/ # 衝突解決
│   │   │   ├── conflicts-detected.json
│   │   │   └── resolution-strategy.md
│   │   └── manifests/           # 同步清單
│   │       ├── tools-manifest.json      # 工具清單
│   │       └── content-manifest.json    # 內容清單
│   ├── 📊 analytics/             # 分析和監控產出
│   │   ├── performance/         # 性能分析
│   │   │   ├── response-times.json      # 響應時間統計
│   │   │   ├── token-usage.json         # Token 使用統計
│   │   │   └── optimization-impact.json # 優化效果
│   │   ├── usage/               # 使用情況分析
│   │   │   ├── tool-usage-stats.json    # 工具使用統計
│   │   │   ├── feature-adoption.json    # 功能採用率
│   │   │   └── user-behavior.json       # 用戶行為分析
│   │   └── reports/             # 綜合報告
│   │       ├── weekly-performance.md
│   │       ├── monthly-usage.md
│   │       └── quarterly-review.md
│   ├── 🧪 experiments/           # 實驗和測試產出
│   │   ├── ab-tests/            # A/B 測試
│   │   │   ├── query-optimization-test.md
│   │   │   └── workflow-comparison.md
│   │   ├── prototypes/          # 原型驗證
│   │   │   ├── new-mcp-tools/
│   │   │   └── alternative-workflows/
│   │   ├── research/            # 研究成果
│   │   │   ├── tool-comparison.md
│   │   │   └── industry-benchmarks.md
│   │   └── legacy/              # 歷史系統存檔
│   │       └── workflow-legacy/ # 舊版 workflow 系統
│   ├── 🔧 configs/               # 配置文件
│   │   ├── mcp-server-config.yaml     # MCP 服務器配置
│   │   ├── ollama-config.yaml         # Ollama 配置
│   │   ├── gate-rules.yaml            # 門禁規則配置
│   │   ├── tool-adapters.yaml         # 工具適配器配置
│   │   └── sync-policies.yaml         # 同步策略配置
│   └── 🗃️ cache/                 # 緩存文件
│       ├── ollama-cache/              # Ollama 查詢緩存
│       ├── spec-cache/                # 規格文檔緩存
│       ├── validation-cache/          # 驗證結果緩存
│       └── temp/                      # 臨時文件
├── .claude/                      # Claude Code 原生配置 (保持兼容)
├── src/                         # 核心實現代碼
│   ├── memory/                 # 統一記憶系統
│   │   ├── PROJECT_CONTEXT.md
│   │   ├── DECISIONS.md
│   │   └── last-session.yml
│   ├── state/                  # 運行狀態
│   │   └── mcp-state.yml
│   └── config/                 # MCP配置
│       └── mcp-servers.json
├── specs/                      # 規格文檔 (借鑒 Spec-Kit)
│   ├── templates/             # 規格模板
│   │   ├── PRD.template.md    # 產品需求文檔模板
│   │   ├── PLAN.template.md   # 技術方案模板
│   │   └── TASKS.template.md  # 任務分解模板
│   ├── constitution/          # 開發憲章和門禁
│   │   ├── CONSTITUTION.md    # 核心開發原則
│   │   ├── gates.yaml         # 質量門禁配置
│   │   └── constraints.yaml   # 企業約束配置
│   └── active/               # 當前活動規格
│       ├── current-prd.md    # 當前產品需求
│       ├── current-plan.md   # 當前技術方案
│       └── current-tasks.md  # 當前任務清單
├── src/                        # 核心實現
│   ├── spec_engine/           # 規格驅動引擎 (新增)
│   │   ├── __init__.py
│   │   ├── prd_generator.py   # PRD 生成器
│   │   ├── plan_generator.py  # 技術方案生成器
│   │   ├── task_decomposer.py # 任務分解器
│   │   └── gate_validator.py  # 門禁驗證器
│   ├── ollama_optimizer/       # Ollama查詢優化器
│   │   ├── __init__.py
│   │   ├── query_optimizer.py
│   │   ├── context_enhancer.py
│   │   └── spec_enhancer.py   # 基於規格的上下文增強 (新增)
│   ├── mcp_server/            # MCP統一服務
│   │   ├── __init__.py
│   │   ├── server.py
│   │   ├── tools/
│   │   │   ├── spec_tools.py      # 規格工具 (新增)
│   │   │   ├── query_tools.py
│   │   │   ├── addp_tools.py
│   │   │   ├── gate_tools.py      # 門禁工具 (新增)
│   │   │   ├── memory_tools.py
│   │   │   └── quality_tools.py
│   │   └── config.yaml
│   └── adapters/              # 平台適配器
│       ├── claude_code.py
│       ├── gemini_cli.py
│       ├── cursor.py
│       └── base.py
├── scripts/                   # 工具腳本
│   ├── start_mcp_server.py
│   ├── setup_ollama.py
│   ├── validate_setup.py
│   └── validate_gates.py      # 門禁驗證腳本 (新增)
├── tests/                     # 測試套件
│   ├── test_spec_engine.py    # 規格引擎測試 (新增)
│   ├── test_gate_validation.py # 門禁驗證測試 (新增)
│   ├── test_ollama_optimizer.py
│   ├── test_mcp_tools.py
│   └── test_cross_platform.py
└── docs/                      # 文檔
    ├── setup.md
    ├── usage.md
    ├── architecture.md
    └── spec-driven-workflow.md # 規格驅動工作流程文檔 (新增)
```

## 🛠️ MCP 自動初始化實現方案

### ✅ **完全可實現 - 無技術瓶頸**

**核心原因**：MCP 工具本質上就是可執行任意操作的函數，包括完整的文件系統操作權限。

### 🎯 **核心 MCP 初始化工具**

#### **1. 一鍵初始化工具**
```python
# src/mcp_server/tools/project_tools.py

@mcp_tool
async def initialize_addp_structure(project_type: str = "universal-coding") -> dict:
    """
    自動初始化完整的 .addp 目錄結構
    用法: claude "初始化 ADDP 項目結構" → 自動調用此工具
    """
    try:
        # 1. 創建完整目錄結構
        directories = [
            ".addp/specifications/templates",
            ".addp/specifications/active",
            ".addp/specifications/archive",
            ".addp/specifications/constitution",
            ".addp/workflows/analysis",
            ".addp/workflows/design",
            ".addp/workflows/development",
            ".addp/workflows/persistence",
            ".addp/memory/project-context",
            ".addp/memory/sessions",
            ".addp/memory/cross-tool",
            ".addp/memory/knowledge-base",
            ".addp/queries/raw-queries",
            ".addp/queries/optimized-queries",
            ".addp/queries/context-enhancements",
            ".addp/queries/confidence-scores",
            ".addp/gates/validations",
            ".addp/gates/violations",
            ".addp/gates/metrics",
            ".addp/gates/reports",
            ".addp/sync/state-snapshots",
            ".addp/sync/sync-logs",
            ".addp/sync/conflict-resolution",
            ".addp/sync/manifests",
            ".addp/analytics/performance",
            ".addp/analytics/usage",
            ".addp/analytics/reports",
            ".addp/experiments/ab-tests",
            ".addp/experiments/prototypes",
            ".addp/experiments/research",
            ".addp/experiments/legacy",
            ".addp/configs",
            ".addp/cache/ollama-cache",
            ".addp/cache/spec-cache",
            ".addp/cache/validation-cache",
            ".addp/cache/temp"
        ]

        # 原子性創建所有目錄
        for directory in directories:
            Path(directory).mkdir(parents=True, exist_ok=True)

        # 2. 創建初始模板文件
        await create_specification_templates()

        # 3. 創建配置文件
        await create_mcp_configs(project_type)

        # 4. 初始化 Constitution 憲章
        await create_constitution_files()

        # 5. 智能檢測並遷移現有內容
        await migrate_existing_content()

        return {
            "success": True,
            "directories_created": len(directories),
            "templates_created": 5,
            "configs_created": 6,
            "structure_ready": True,
            "next_actions": [
                "使用 spec.create_prd 開始創建產品需求",
                "使用 addp.start_analysis 開始分析階段",
                "所有產出自動保存到 .addp/ 對應目錄"
            ]
        }

    except Exception as e:
        return {
            "success": False,
            "error": str(e),
            "recovery_suggestion": "檢查目錄權限或手動創建基礎結構"
        }

@mcp_tool
async def smart_initialize_project() -> dict:
    """
    智能檢測項目類型並初始化對應結構
    自動檢測: React/Vue/Python/Kotlin 等不同項目類型
    """
    project_info = await detect_project_type()

    # 基礎結構初始化
    base_result = await initialize_addp_structure(project_info["type"])

    # 創建項目特定的模板和配置
    if project_info["type"] == "react":
        await create_react_specific_templates()
        await create_react_gates_config()
    elif project_info["type"] == "python":
        await create_python_specific_templates()
        await create_python_gates_config()
    elif project_info["type"] == "kotlin":
        await create_kotlin_specific_templates()
        await create_kotlin_gates_config()

    return {
        **base_result,
        "project_detected": project_info,
        "custom_templates_created": True,
        "project_specific_config": True
    }

async def detect_project_type():
    """智能檢測當前項目類型"""
    if Path("package.json").exists():
        with open("package.json") as f:
            pkg = json.load(f)
            if "react" in pkg.get("dependencies", {}):
                return {"type": "react", "framework": "react", "package_manager": "npm"}
            elif "vue" in pkg.get("dependencies", {}):
                return {"type": "vue", "framework": "vue", "package_manager": "npm"}

    if Path("pyproject.toml").exists() or Path("requirements.txt").exists():
        return {"type": "python", "framework": "python", "package_manager": "pip"}

    if Path("build.gradle.kts").exists() or Path("build.gradle").exists():
        return {"type": "kotlin", "framework": "gradle", "package_manager": "gradle"}

    return {"type": "universal-coding", "framework": "generic", "package_manager": "none"}
```

#### **2. 模板文件自動生成**
```python
async def create_specification_templates():
    """自動創建所有規格模板文件"""

    # PRD 模板 (借鑒 Spec-Kit)
    prd_template = """# Product Requirements Document

## 📋 Summary
{{ prd_summary }}

## 👤 User Stories
{{ user_stories }}

## 🔧 Technical Requirements
{{ technical_requirements }}

## ✅ Acceptance Criteria
{{ acceptance_criteria }}

## ⚡ Quality Gates (借鑒 Spec-Kit Constitution)
- [ ] **TDD First**: 測試必須在實現前編寫
- [ ] **Anti-Abstract**: 避免過度抽象，優先具體解決方案
- [ ] **Simplify First**: 選擇最簡單的可行方案
- [ ] **Integration First**: 優先編寫集成測試

## 🏗️ Architecture Constraints
{{ architecture_constraints }}

## 📊 Success Metrics
{{ success_metrics }}

## 🤖 Generated by
- **Tool**: {{ generating_tool }}
- **Timestamp**: {{ timestamp }}
- **Confidence**: {{ confidence_score }}%
- **Project Type**: {{ project_type }}
"""

    # 技術方案模板
    plan_template = """# Technical Implementation Plan

## 🎯 Overview
{{ plan_overview }}

## 🏗️ Architecture Analysis
{{ architecture_analysis }}

## 🛠️ Implementation Strategy
{{ implementation_strategy }}

## 📝 Task Breakdown
{{ task_breakdown }}

## 🧪 Testing Strategy
{{ testing_strategy }}

## ⚠️ Risk Assessment
{{ risk_assessment }}

## 📈 Success Metrics
{{ success_metrics }}

## 🤖 Generated by
- **Tool**: {{ generating_tool }}
- **Timestamp**: {{ timestamp }}
- **Based on PRD**: {{ source_prd }}
"""

    # 任務分解模板
    tasks_template = """# Task Breakdown

## 📋 Task Overview
{{ task_overview }}

## ✅ Development Tasks
{{ development_tasks }}

## 🧪 Testing Tasks
{{ testing_tasks }}

## 📚 Documentation Tasks
{{ documentation_tasks }}

## ⚡ Quality Gates Checklist
{{ quality_gates_checklist }}

## 🤖 Generated by
- **Tool**: {{ generating_tool }}
- **Timestamp**: {{ timestamp }}
- **Based on Plan**: {{ source_plan }}
"""

    # ADR 模板
    adr_template = """# ADR: {{ decision_title }}

## Status
{{ status }}

## Context
{{ context }}

## Decision
{{ decision }}

## Consequences
{{ consequences }}

## Quality Gates Impact
{{ quality_gates_impact }}

---
**Generated**: {{ timestamp }}
**Tool**: {{ generating_tool }}
"""

    # 保存所有模板
    templates = {
        ".addp/specifications/templates/prd.template.md": prd_template,
        ".addp/specifications/templates/plan.template.md": plan_template,
        ".addp/specifications/templates/tasks.template.md": tasks_template,
        ".addp/specifications/templates/adr.template.md": adr_template
    }

    for path, content in templates.items():
        Path(path).parent.mkdir(parents=True, exist_ok=True)
        with open(path, "w", encoding="utf-8") as f:
            f.write(content)
```

#### **3. Constitution 憲章自動創建**
```python
async def create_constitution_files():
    """創建開發憲章和門禁配置 (借鑒 Spec-Kit)"""

    # 核心開發憲章
    constitution = """# Development Constitution
> 借鑒 [GitHub Spec-Kit](https://github.com/github/spec-kit) 的 Constitution 理念

## 🎯 Core Principles

### 1. TDD First (測試驅動開發)
- **規則**: 所有代碼必須在實現前先寫測試
- **無例外**: 即使是"簡單"函數也必須有測試
- **覆蓋率**: 測試覆蓋率必須 >80%
- **驗證**: `gate.check_tdd_first` 自動檢查

### 2. Anti-Abstract (反過度抽象)
- **規則**: 避免過早抽象
- **優先**: 先選擇具體解決方案
- **重構**: 只在模式明確時才抽象
- **驗證**: `gate.check_anti_abstract` 自動檢查

### 3. Simplify First (簡化優先)
- **規則**: 總是選擇最簡單的可行方案
- **明確**: 優先明確勝過巧妙
- **優化**: 初期優化可讀性而非性能
- **驗證**: `gate.check_simplify_first` 自動檢查

### 4. Integration First (集成優先)
- **規則**: 先寫集成測試再寫單元測試
- **真實**: 測試真實用戶場景
- **Mock**: 只 mock 外部依賴
- **驗證**: `gate.check_integration` 自動檢查

## ⚡ Quality Gates
所有 PRD/Plan/Tasks 必須通過這些門禁才能進入實施階段。

## 🚫 Gate Violations
違反門禁的代碼將被自動阻止，必須修正後才能繼續。
"""

    # 門禁規則配置
    gates_config = {
        "tdd_first": {
            "enabled": True,
            "severity": "error",
            "rules": [
                "每個函數必須有對應測試",
                "測試必須在代碼實現前編寫",
                "測試覆蓋率必須 >= 80%"
            ]
        },
        "anti_abstract": {
            "enabled": True,
            "severity": "warning",
            "rules": [
                "避免創建不必要的抽象層",
                "具體實現優先於通用方案",
                "重構門檻: 至少3個相似模式"
            ]
        },
        "simplify_first": {
            "enabled": True,
            "severity": "error",
            "rules": [
                "選擇最簡單的可行方案",
                "代碼行數限制: 函數 < 20行",
                "循環複雜度限制: < 10"
            ]
        },
        "integration_first": {
            "enabled": True,
            "severity": "warning",
            "rules": [
                "優先編寫 E2E 測試",
                "集成測試覆蓋核心流程",
                "單元測試補充細節邏輯"
            ]
        }
    }

    # 保存憲章文件
    with open(".addp/specifications/constitution/CONSTITUTION.md", "w", encoding="utf-8") as f:
        f.write(constitution)

    # 保存門禁配置
    with open(".addp/specifications/constitution/gates.yaml", "w", encoding="utf-8") as f:
        yaml.dump(gates_config, f, allow_unicode=True)
```

#### **4. 跨工具同步設置**
```python
@mcp_tool
async def setup_cross_tool_sync() -> dict:
    """
    設置跨工具同步配置
    確保 Claude Code、Gemini CLI、Cursor 等工具狀態一致
    """
    tools = ["claude-code", "gemini-cli", "cursor"]
    sync_results = {}

    for tool in tools:
        try:
            # 檢查工具連接狀態
            if await check_tool_connection(tool):
                # 創建工具特定的狀態文件
                tool_state = {
                    "tool": tool,
                    "addp_version": "1.0.0",
                    "directory_structure": ".addp",
                    "last_sync": datetime.now().isoformat(),
                    "sync_enabled": True
                }

                state_file = f".addp/memory/cross-tool/{tool}-state.json"
                with open(state_file, "w", encoding="utf-8") as f:
                    json.dump(tool_state, f, indent=2, ensure_ascii=False)

                sync_results[tool] = "configured"
            else:
                sync_results[tool] = "not_available"

        except Exception as e:
            sync_results[tool] = f"error: {e}"

    return {
        "sync_setup_completed": True,
        "tool_results": sync_results,
        "ready_for_cross_tool_usage": True
    }
```

### 🔧 **使用方式**

#### **在任何 AI 工具中一鍵初始化**
```bash
# 方式 1: 直接調用 (推薦)
claude "初始化 ADDP 項目結構"          # 自動調用 project.initialize
gemini "設置統一編程環境"              # 自動調用 project.setup
cursor "創建 MCP 工作目錄"            # 自動調用 project.initialize

# 方式 2: 通過 CLI 命令
mcp-init --project-type=universal-coding

# 方式 3: 智能檢測初始化
claude "智能初始化項目"                # 自動檢測項目類型並初始化
```

#### **初始化完成後的自動化流程**
```bash
# 1. 結構初始化完成，可以開始使用
claude "優化 React 組件性能"
# → 自動調用 query.optimize (Ollama 優化)
# → 自動調用 spec.create_prd (生成 PRD)
# → 產出保存到 .addp/specifications/active/current-prd.md

# 2. 繼續工作流程
claude "生成技術實施方案"
# → 自動調用 spec.generate_plan
# → 產出保存到 .addp/specifications/active/current-plan.md

# 3. 開始開發
claude "開始 ADDP 分析階段"
# → 自動調用 addp.start_analysis
# → 產出保存到 .addp/workflows/analysis/
```

## 🚀 實施計劃（2週MVP）

### Week 1: 規格驅動層 + Ollama優化層 + 基礎架構
**目標**：實現規格驅動引擎和本地查詢優化

#### TDD 開發順序 (集成 Spec-Kit 理念)：
```bash
# Phase 1: 規格驅動基礎架構
specs/templates/PRD.template.md            # PRD 模板 (借鑒 Spec-Kit)
specs/templates/PLAN.template.md           # 技術方案模板
specs/constitution/CONSTITUTION.md         # 開發憲章
specs/constitution/gates.yaml              # 質量門禁配置

# Phase 2: 測試先行 - 創建測試框架
tests/test_spec_engine.py                  # 規格引擎測試
tests/test_gate_validation.py              # 門禁驗證測試
tests/test_ollama_optimizer.py             # 查詢優化測試
tests/test_mcp_tools.py                    # MCP工具測試

# Phase 3: 核心實現
src/spec_engine/prd_generator.py           # PRD 生成器
src/spec_engine/gate_validator.py          # 門禁驗證器
src/ollama_optimizer/query_optimizer.py    # 核心優化邏輯
src/ollama_optimizer/spec_enhancer.py      # 基於規格的增強
scripts/setup_ollama.py                    # Ollama環境設置
scripts/validate_gates.py                  # 門禁驗證腳本

# Phase 4: 驗收標準 (規格驅動流程)
# 1. 規格生成測試
python -m src.spec_engine.prd_generator "優化React性能"
# 輸出：完整的 PRD 文檔 + 技術方案 + 任務分解

# 2. 門禁驗證測試
python scripts/validate_gates.py specs/active/current-prd.md
# 輸出：TDD先行、反抽象、簡化優先等門禁檢查結果

# 3. 集成優化測試
python -m src.ollama_optimizer.query_optimizer "優化React性能" --with-spec
# 輸出：基於規格文檔增強的精準技術指令
```

### Week 2: MCP統一服務層
**目標**：實現基礎MCP服務和跨工具整合

#### TDD 開發順序：
```bash
# 1. 測試驅動MCP服務開發
src/mcp_server/server.py                   # MCP服務主體
src/mcp_server/tools/query_tools.py        # 查詢優化MCP工具
src/mcp_server/tools/addp_tools.py         # ADDP工作流程MCP工具
src/adapters/claude_code.py                # Claude Code適配器
src/adapters/gemini_cli.py                 # Gemini CLI適配器

# 2. 集成測試
python scripts/start_mcp_server.py         # MCP服務啟動測試
python -m pytest tests/test_cross_platform.py  # 跨工具集成測試
```

## 🎯 TDD 驅動的驗收標準（Definition of Done）

### 測試金字塔
```yaml
單元測試 (70%):
  - test_query_optimizer.py      # 查詢優化核心邏輯
  - test_addp_workflow.py        # ADDP工作流程
  - test_memory_sync.py          # 記憶同步機制

集成測試 (20%):
  - test_mcp_server_integration.py  # MCP服務集成
  - test_ollama_integration.py      # Ollama集成

端到端測試 (10%):
  - test_claude_code_e2e.py      # Claude Code完整流程
  - test_gemini_cli_e2e.py       # Gemini CLI完整流程
```

### 功能驗收 (所有測試必須通過)
```bash
# 1. Ollama優化器測試
python -m pytest tests/test_ollama_optimizer.py -v
echo "優化性能" | python -m src.ollama_optimizer.query_optimizer
# 期望：精準的技術指令 + 建議的agent + 驗收條件

# 2. MCP服務測試
python -m pytest tests/test_mcp_tools.py -v
python scripts/start_mcp_server.py --port 8000 --test-mode
# 期望：所有MCP工具正常註冊和響應

# 3. 跨工具調用測試
python -m pytest tests/test_cross_platform.py -v
# 期望：Claude Code 和 Gemini CLI 產生一致的優化指令
```

### 效果驗收基準
```yaml
精準度提升:
  測試案例: 10個典型模糊查詢
  基準: 直接使用AI工具的回應質量
  目標: 40%+的精準度提升
  測試方法: A/B 對比測試

響應時間:
  Ollama優化: < 3秒 (本地模型)
  MCP調用: < 1秒 (本地服務)
  總體延遲: < 5秒 (可接受範圍)

記憶同步:
  測試場景: Claude Code → Gemini CLI 切換
  驗證標準: 項目上下文100%保留
  測試方法: 狀態差異檢查
```

## 🛠️ 關鍵實現細節 (TDD 模式)

### Ollama查詢優化實現
```python
# tests/test_ollama_optimizer.py
class TestQueryOptimizer:
    def test_optimize_vague_query(self):
        optimizer = QueryOptimizer()
        result = optimizer.optimize_query("優化性能")

        assert result.confidence > 0.7
        assert "React.memo" in result.optimized
        assert result.suggested_agent is not None

# src/ollama_optimizer/query_optimizer.py
class QueryOptimizer:
    def __init__(self, model="qwen2.5:14b"):
        self.ollama = ollama.Client()
        self.model = model

    async def optimize_query(self, raw_input, project_context=None):
        """
        核心優化功能：模糊輸入 → 精準技術指令
        """
        # 實現通過測試驅動開發
        intent = await self._analyze_intent(raw_input)
        enhanced_context = await self._enhance_context(raw_input, intent, project_context)
        optimized = await self._generate_optimized_prompt(raw_input, intent, enhanced_context)

        return OptimizedQuery(
            original=raw_input,
            optimized=optimized.instruction,
            confidence=optimized.confidence,
            suggested_agent=optimized.agent,
            acceptance_criteria=optimized.criteria
        )
```

### MCP工具實現 (測試驅動)
```python
# tests/test_mcp_tools.py
class TestMCPTools:
    def test_optimize_user_query_tool(self):
        result = optimize_user_query("Add login feature", "claude-code")

        assert result["confidence_score"] > 0.7
        assert "suggested_workflow" in result
        assert result["optimized_query"] != "Add login feature"

# src/mcp_server/tools/query_tools.py
@mcp_tool
async def optimize_user_query(raw_query: str, target_platform: str) -> dict:
    """
    MCP工具：查詢優化 - 通過測試驅動開發
    """
    optimizer = QueryOptimizer()
    project_context = await load_project_context()
    result = await optimizer.optimize_query(raw_query, project_context)
    platform_specific = await adapt_for_platform(result, target_platform)

    return {
        "optimized_query": platform_specific,
        "confidence_score": result.confidence,
        "suggested_workflow": "addp_standard"
    }
```

## 🔧 開發環境設置 (TDD 準備)

### 1. 測試環境設置
```bash
# 創建虛擬環境
python -m venv venv
source venv/bin/activate  # Windows: venv\Scripts\activate

# 安裝依賴 (包含測試工具)
pip install pytest pytest-asyncio pytest-cov
pip install ollama mcp anthropic google-generativeai

# 設置測試配置
echo "[tool.pytest.ini_options]
testpaths = ['tests']
python_files = ['test_*.py']
python_classes = ['Test*']
python_functions = ['test_*']
asyncio_mode = 'auto'" > pyproject.toml
```

### 2. Ollama環境 (開發用)
```bash
# 安裝Ollama
curl -fsSL https://ollama.ai/install.sh | sh

# 下載模型
ollama pull qwen2.5:14b

# 測試連接
python -c "import ollama; print('Ollama ready')"
```

### 3. 持續集成設置
```bash
# 創建 Makefile 用於標準化操作
echo "test:
	python -m pytest tests/ -v --cov=src

lint:
	python -m flake8 src/ tests/

typecheck:
	python -m mypy src/

ci: lint typecheck test
	@echo 'All checks passed'

setup:
	pip install -r requirements.txt
	python scripts/setup_ollama.py

.PHONY: test lint typecheck ci setup" > Makefile
```

## ⚠️ 風險與緩解策略

### 技術風險
- **Ollama性能**：使用緩存機制，重複查詢<1秒響應
- **MCP兼容性**：優先支援原生MCP的工具(Claude Code, Gemini CLI)
- **測試覆蓋率**：要求 >80% 測試覆蓋率，CI 自動檢查

### 實施風險
- **TDD 學習曲線**：提供詳細的測試模板和示例
- **開發複雜度**：採用MVP方式，優先核心功能
- **維護成本**：建立完整的自動化測試和CI/CD

## 📈 成功指標 (可測量)

### 短期指標（2週）
- [ ] 所有單元測試通過 (>95% 覆蓋率)
- [ ] Ollama優化器精準度提升40%+ (A/B測試驗證)
- [ ] Claude Code + Gemini CLI基礎整合完成 (E2E測試)

### 中期指標（2個月）
- [ ] 支援4+個AI coding工具 (集成測試覆蓋)
- [ ] 跨工具記憶同步100%可用 (狀態一致性測試)
- [ ] ADDP工作流程完整實現 (工作流程測試)

### 長期指標（6個月）
- [ ] 開源項目：100+ GitHub stars，10+ contributors
- [ ] 性能基準：90%用戶報告開發效率提升
- [ ] 行業影響：推動MCP在AI coding領域標準化

---

**🎯 下一步行動**：開始TDD驅動的Week 1開發

```bash
# TDD 第一步：創建測試框架
mkdir -p tests src/ollama_optimizer src/mcp_server/tools
touch tests/__init__.py
echo "# 查詢優化器測試框架" > tests/test_ollama_optimizer.py
echo "# MCP工具測試框架" > tests/test_mcp_tools.py

# 運行初始測試 (應該失敗，這是TDD的起點)
python -m pytest tests/ -v
```
