# Agent System Documentation

*English | [中文](README_zh.md)*

## Overview

The Claude Code Agent System provides specialized AI assistants for various development tasks. These agents offer deep expertise in specific domains, languages, and frameworks, automatically activating based on context.

## How Agents Work

1. **Automatic Activation**: Agents activate based on task description keywords
2. **Context Awareness**: Agents understand file types and code patterns
3. **Collaborative Work**: Multiple agents can work together on complex tasks
4. **Professional Depth**: Each agent has specialized knowledge in its domain

## Agent Categories

### 🔍 Quality Assurance Agents

#### `code-reviewer`
- **Expertise**: Code quality, best practices, security issues
- **Activates on**: Code review requests, PR reviews
- **Capabilities**:
  - Identify bugs and potential issues
  - Suggest improvements
  - Check coding standards
  - Security vulnerability detection

#### `test-automator`
- **Expertise**: Test generation, coverage analysis
- **Activates on**: Test-related tasks
- **Capabilities**:
  - Generate unit tests
  - Create integration tests
  - Design test scenarios
  - Coverage recommendations

#### `performance-optimizer`
- **Expertise**: Performance analysis, optimization
- **Activates on**: Performance issues, optimization requests
- **Capabilities**:
  - Identify bottlenecks
  - Suggest optimizations
  - Memory usage analysis
  - Algorithm improvements

#### `bug-hunter`
- **Expertise**: Debugging, root cause analysis
- **Activates on**: Bug reports, error messages
- **Capabilities**:
  - Analyze stack traces
  - Find root causes
  - Suggest fixes
  - Prevent similar issues

### 💻 Technical Expert Agents

#### Language Specialists

**`python-ml-specialist`**
- **Domain**: Machine learning, data science
- **Frameworks**: PyTorch, TensorFlow, scikit-learn
- **Use cases**: Model development, data analysis, AI applications

**`typescript-fullstack-expert`**
- **Domain**: Full-stack TypeScript development
- **Frameworks**: React, Node.js, Next.js
- **Use cases**: Web applications, APIs, type-safe development

**`kotlin-expert`**
- **Domain**: Kotlin development across platforms
- **Frameworks**: Android, Spring Boot, Ktor
- **Use cases**: Mobile apps, backend services

**`golang-systems-engineer`**
- **Domain**: Systems programming, microservices
- **Frameworks**: Standard library, popular Go packages
- **Use cases**: High-performance services, cloud-native apps

**`rust-zero-cost`**
- **Domain**: Systems programming, performance-critical code
- **Frameworks**: Tokio, Actix, standard library
- **Use cases**: Systems tools, web services, embedded

#### Framework Specialists

**`android-kotlin-architect`**
- **Domain**: Android application development
- **Technologies**: Jetpack Compose, Coroutines, Architecture Components
- **Use cases**: Mobile apps, UI development

**`ktor-backend-architect`**
- **Domain**: Kotlin backend services
- **Technologies**: Ktor framework, coroutines
- **Use cases**: REST APIs, microservices

### 🎭 Workflow Agents

#### `work-coordinator`
- **Purpose**: Coordinate multiple agents for complex tasks
- **Capabilities**:
  - Task distribution
  - Result aggregation
  - Dependency management
  - Cross-domain coordination

### 🛡️ Validation Agents

#### `jenny-validator`
- **Purpose**: Validate specifications and requirements
- **Capabilities**:
  - Requirement verification
  - Specification compliance
  - Consistency checking

#### `karen-realist`
- **Purpose**: Reality checks and feasibility assessment
- **Capabilities**:
  - Timeline validation
  - Resource assessment
  - Risk identification

## Using Agents

### Automatic Activation

Agents activate automatically based on your task:

```bash
# This automatically activates python-ml-specialist
"Help me build a text classification model"

# This automatically activates security-analyst
"Review this authentication code for vulnerabilities"

# This automatically activates android-kotlin-architect
"Design a shopping cart feature for Android"
```

### Manual Activation

You can explicitly request specific agents:

```bash
"Use the rust-zero-cost agent to optimize this code"
"Have the test-automator generate comprehensive tests"
```

### Agent Collaboration

For complex tasks, multiple agents work together:

```bash
"Build a secure payment processing system"
# Activates:
# - security-analyst (security review)
# - code-reviewer (code quality)
# - test-automator (test generation)
# - performance-optimizer (optimization)
```

## Best Practices

1. **Trust agent expertise**: Agents have deep domain knowledge
2. **Let agents auto-activate**: Don't micromanage agent selection
3. **Use clear task descriptions**: Better descriptions = better agent matching
4. **Combine agents with commands**: Use commands for workflow, agents for expertise
5. **Review agent suggestions**: Agents provide recommendations, you make decisions

## Agent Capabilities Matrix

| Agent | Languages | Frameworks | Best For |
|-------|-----------|------------|----------|
| python-ml-specialist | Python | PyTorch, TensorFlow | ML/AI development |
| typescript-fullstack-expert | TypeScript/JavaScript | React, Node.js | Web development |
| kotlin-expert | Kotlin | Android, Spring, Ktor | Mobile & backend |
| golang-systems-engineer | Go | Standard library | Microservices |
| rust-zero-cost | Rust | Tokio, Actix | Systems programming |
| android-kotlin-architect | Kotlin | Android SDK | Android apps |
| code-reviewer | All | All | Code quality |
| test-automator | All | Testing frameworks | Test generation |
| performance-optimizer | All | All | Performance tuning |

## Creating Custom Agents

To create a custom agent, create a file in `~/.claude/agents/`:

```markdown
---
description: Brief description for auto-activation
---

# Agent Name

You are an expert in [domain]. Your expertise includes...

## Capabilities
- [Capability 1]
- [Capability 2]

## Activation Keywords
- [Keyword 1]
- [Keyword 2]

## Response Style
[How the agent should communicate]
```

## Troubleshooting

**Agent not activating?**
- Use more specific keywords in your request
- Manually specify the agent name
- Check if the agent file exists in `~/.claude/agents/`

**Wrong agent activated?**
- Be more specific in your task description
- Manually override with explicit agent selection

**Agent giving generic responses?**
- Ensure agent file is properly formatted
- Check that agent has clear expertise definition

---

For more examples, see the [Agent Usage Guide](../../guides/AGENT_GUIDE.md)