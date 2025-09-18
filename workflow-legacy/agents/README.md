# 🤖 Claude Code Pro Agents System

*English | [中文](README_cn.md)*

A comprehensive collection of specialized AI agents for Claude Code, designed to provide professional-grade assistance across all aspects of software development.

## 📋 Table of Contents

- [Overview](#overview)
- [Quick Start](#quick-start)
- [Agent Categories](#agent-categories)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Available Agents](#available-agents)
- [Workflows](#workflows)
- [Best Practices](#best-practices)

## 🎯 Overview

This agent system integrates the best practices from leading Claude Code agent projects:
- **wshobson/agents** - Model-tiered agent selection
- **claude_code_agent_farm** - Parallel execution and coordination
- **ClaudeCodeAgents** - Quality assurance specialists

### Key Features

- 🚀 **40+ Specialized Agents** covering all major languages and frameworks
- ⚡ **Smart Selection via Orchestrator** - intelligent single-expert prioritization
- 🔄 **Token-Optimized Workflows** - 60%+ efficiency improvement (800k → 300k avg)
- 🌐 **Language Preference Memory** - persistent cross-session localization
- 📁 **Standardized Memory System** - unified `.claude/memory/` directory structure
- 📊 **Intelligent Coordination** with parallel execution support
- ✅ **Quality Assurance** with dedicated validation agents
- 🎨 **Model Optimization** using Haiku/Sonnet/Opus based on task complexity
- 🎯 **Production-Ready Code** with built-in best practices and security
- 🐛 **Advanced Debugging** with root cause analysis
- 📡 **API Design** with OpenAPI specs and GraphQL schemas

## 🚀 Quick Start

### For Global Installation (Recommended)

1. Copy the entire `claude/agents` directory to your home directory:
```bash
# Windows
xcopy /E /I "D:\Code\ai\claude\agents" "%USERPROFILE%\.claude\agents"

# macOS/Linux
cp -r /path/to/claude/agents ~/.claude/agents
```

2. Copy configuration files:
```bash
# Windows
xcopy /E /I "D:\Code\ai\claude\config" "%USERPROFILE%\.claude\config"

# macOS/Linux
cp -r /path/to/claude/config ~/.claude/config
```

### For Project-Specific Installation

1. Copy to your project root:
```bash
cp -r /path/to/claude/agents .claude/agents
cp -r /path/to/claude/config .claude/config
```

## 📁 Agent Categories

### 1. Core Agents (`agents/core/`)
Essential agents for code quality and optimization:
- `code-reviewer` - Comprehensive code review
- `test-automator` - Test generation and coverage
- `performance-optimizer` - Performance analysis and optimization

### 2. Language Specialists (`agents/languages/`)
Expert agents for specific programming languages:
- `rust-zero-cost` - Rust systems programming
- `cpp-modern-master` - Modern C++20/23
- `golang-systems-engineer` - Go microservices
- `python-ml-specialist` - Python ML/AI development
- `android-kotlin-architect` - Android development
- **`typescript-fullstack-expert`** - TypeScript with advanced type systems, React/Vue/Angular, Node.js
- **`csharp-dotnet-master`** - C# .NET 8+, ASP.NET Core, Entity Framework, Clean Architecture
- **`c-systems-architect`** - C systems programming, memory management, embedded systems, SIMD
- **`python-fullstack-expert`** - Python web frameworks, data science, ML/AI, async programming
- **`java-enterprise-architect`** - Java Spring Boot, microservices, reactive programming, Kafka

### 3. Framework Experts (`agents/frameworks/`)
Specialized agents for frameworks:
- `ktor-backend-architect` - Ktor server development
- `spring-boot-enterprise` - Spring Boot applications
- `react-nextjs-expert` - React/Next.js development
- `vue-nuxt-specialist` - Vue/Nuxt.js development

### 4. Quality Assurance (`agents/quality/`)
Validation and reality-check agents:
- `jenny-validator` - Specification compliance
- `karen-realist` - Timeline and scope reality checks

### 5. Workflow Coordination (`agents/workflow/`)
Orchestration and management agents:
- `work-coordinator` - Multi-agent task coordination
- `context-manager` - Long conversation context management
- **`production-ready-coder`** - Automatically writes production-quality code with security, testing, and documentation
- **`bug-hunter`** - Specializes in finding and fixing bugs with root cause analysis
- **`api-architect`** - Designs REST/GraphQL APIs with OpenAPI specs and best practices
- **`senior-developer`** - Applies 10+ years of experience to every line of code

## 💻 Usage

### Automatic Agent Activation

Agents are automatically triggered based on file types:

```kotlin
// Editing a .kt file automatically activates android-kotlin-architect
class MainActivity : AppCompatActivity() {
    // Agent provides Kotlin and Android best practices
}
```

```rust
// Editing a .rs file automatically activates rust-zero-cost
fn main() {
    // Agent ensures zero-cost abstractions and memory safety
}
```

### Manual Agent Invocation

You can explicitly request specific agents:

```
"Use the security-auditor to check this code"
"Have karen-realist assess this timeline"
"Get performance-optimizer to analyze this function"
```

### Command-Triggered Workflows

Commands automatically trigger agent workflows:

- `/check` → Activates code-reviewer, jenny-validator, security-auditor
- `/test` → Activates test-automator
- `/optimize` → Activates performance-optimizer
- `/review` → Activates comprehensive review workflow

## ⚙️ Configuration

### File Type Triggers (`config/triggers.yaml`)

```yaml
file_triggers:
  "*.rs":
    primary: rust-zero-cost
    support: [test-automator, performance-optimizer]
    
  "*.kt":
    context_dependent:
      android:
        detect: ["AndroidManifest.xml"]
        primary: android-kotlin-architect
      server:
        detect: ["ktor", "spring"]
        primary: ktor-backend-architect
```

### Workflow Definitions (`config/workflows.yaml`)

```yaml
workflows:
  feature_development:
    phases:
      - name: "Architecture Design"
        agents: [system-architect, database-architect]
      - name: "Implementation"
        agents: [$auto]  # Auto-select based on language
      - name: "Testing"
        agents: [test-automator, jenny-validator]
```

## 🤖 Available Agents

### High-Performance (Opus Model)
Best for complex analysis and architecture:
- `rust-zero-cost` - Rust systems programming
- `cpp-modern-master` - Modern C++ expertise
- `performance-optimizer` - Deep performance analysis
- `work-coordinator` - Complex task orchestration

### Balanced (Sonnet Model)
Ideal for implementation and review:
- `android-kotlin-architect` - Android development
- `golang-systems-engineer` - Go services
- `python-ml-specialist` - ML/AI development
- `code-reviewer` - Code quality review
- `karen-realist` - Project reality checks
- `typescript-fullstack-expert` - TypeScript expertise
- `csharp-dotnet-master` - C# and .NET development
- `c-systems-architect` - C systems programming
- `python-fullstack-expert` - Python full-stack development
- `java-enterprise-architect` - Java enterprise development
- `production-ready-coder` - Production-quality code generation
- `bug-hunter` - Bug finding and fixing
- `api-architect` - API design and implementation

### Fast (Haiku Model)
Perfect for quick validations:
- `jenny-validator` - Specification checking
- `doc-maintainer` - Documentation updates
- `style-formatter` - Code formatting

## 🔄 Workflows

### Feature Development
Complete feature implementation workflow:
1. Requirements analysis (jenny-validator, karen-realist)
2. Architecture design (system-architect)
3. Parallel implementation (language specialists)
4. Testing (test-automator)
5. Review (code-reviewer, security-auditor)

### Bug Fix
Rapid bug resolution:
1. Root cause analysis (debug-specialist)
2. Fix implementation (language expert)
3. Regression testing (test-automator)
4. Code review (code-reviewer)

### Performance Optimization
Systematic performance improvement:
1. Profiling (performance-optimizer)
2. Bottleneck analysis (system-architect)
3. Parallel optimization (multiple specialists)
4. Validation (performance-optimizer)

## 🎯 Best Practices

### 1. Let Agents Auto-Activate
The system automatically selects appropriate agents based on context. Trust the automatic selection for optimal results.

### 2. Use Workflows for Complex Tasks
For multi-step tasks, use predefined workflows rather than manual agent coordination.

### 3. Listen to Karen
When `karen-realist` provides timeline estimates, they're based on real-world experience. Ignore at your own risk.

### 4. Validate with Jenny
Always let `jenny-validator` check implementation against specifications to catch missing requirements early.

### 5. Parallel Execution
The system automatically parallelizes independent tasks. Multiple agents can work simultaneously for faster results.

## 🔧 Customization

### Adding New Agents

Create a new agent file in the appropriate category:

```yaml
---
name: your-custom-agent
model: sonnet  # haiku|sonnet|opus
description: "Your agent's expertise"
trigger: "*.xyz"
tools: all
---

# Agent Name

Your agent's prompt and expertise description...
```

### Modifying Triggers

Edit `config/triggers.yaml` to customize when agents activate:

```yaml
file_triggers:
  "*.custom":
    primary: your-custom-agent
    support: [test-automator]
```

### Creating Workflows

Add new workflows to `config/workflows.yaml`:

```yaml
workflows:
  custom_workflow:
    description: "Your workflow description"
    phases:
      - name: "Phase 1"
        agents: [agent1, agent2]
        mode: parallel
```

## 📊 Performance Tips

1. **Model Selection**: Agents use appropriate models (Haiku/Sonnet/Opus) based on task complexity
2. **Parallel Execution**: Up to 20 agents can work simultaneously
3. **Caching**: Results are cached to avoid redundant work
4. **Context Management**: Long conversations automatically trigger context-manager

## 🤝 Contributing

To contribute new agents or improvements:

1. Follow the existing agent format
2. Place agents in appropriate categories
3. Update triggers.yaml for automatic activation
4. Add workflows for complex multi-agent tasks
5. Test thoroughly before submission

## 📄 License

This agent system is designed for use with Claude Code and follows Anthropic's usage guidelines.

## 🙏 Acknowledgments

This system integrates concepts and best practices from:
- [wshobson/agents](https://github.com/wshobson/agents)
- [claude_code_agent_farm](https://github.com/Dicklesworthstone/claude_code_agent_farm)
- [ClaudeCodeAgents](https://github.com/darcyegb/ClaudeCodeAgents)

---

**Ready to supercharge your development with Pro-level AI assistance! 🚀**