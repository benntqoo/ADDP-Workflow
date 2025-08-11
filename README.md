# Claude Code Collaboration Framework v3.3

🌐 **Language / 語言**: [English](README.md) | [中文](README_CN.md)

A comprehensive collaboration framework for Claude Code that enhances AI-assisted development through structured commands, memory management, and best practices.

## 📚 Table of Contents

1. [Quick Start](#-quick-start)
2. [Command System v3.0](#-command-system-v30)
3. [Command Reference Guide](#-command-reference-guide)
4. [Project Structure](#-project-structure)
5. [Collaboration Constitution](#-collaboration-constitution)
6. [Usage Guide](#-usage-guide)
7. [Workflows](#-workflows)
8. [Best Practices](#-best-practices)
9. [Version History](#-version-history)

### 🔍 Command Quick Index

| Category | Command | Description | Details |
|----------|---------|-------------|---------|
| **Project Understanding** | `/start` | Quick project understanding | [Details](#1-start---quick-project-understanding) |
| | `/sync` | Restore work state | [Details](#2-sync---state-synchronizer) |
| | `/context` | Confirm understanding | [Details](#3-context---context-sync-checkpoint) |
| **Development Aid** | `/plan` | Task planning & decomposition | [Details](#4-plan---task-planning--design) |
| | `/check` | Complete quality check | [Details](#5-check---complete-quality-check) |
| | `/watch` | Guardian mode | [Details](#6-watch---guardian-mode) |
| | `/test` | Test generation & execution | [Details](#7-test---test-generation--execution) |
| **Problem Solving** | `/debug` | Quick error location | [Details](#8-debug---intelligent-debugging-assistant) |
| | `/analyze` | Deep risk analysis | [Details](#13-analyze---deep-analysis--validation) |
| **Knowledge Management** | `/learn` | Record technical decisions | [Details](#9-learn---learn-and-record-decisions) |
| | `/doc` | Smart documentation maintenance | [Details](#10-doc---intelligent-documentation-maintenance) |
| **Workflow** | `/review` | PR preparation assistant | [Details](#11-review---pr-preparation-assistant) |
| | `/meta` | Project specification customization | [Details](#12-meta---project-specification-customization) |
| | `/update-spec` | Specification update management | [Details](#14-update-spec---claudemd-update-specialist) |
| **SDK Development** | `/sdk-design` | API design guidance | [Details](#1-sdk-design---api-design-assistant) |
| | `/sdk-example` | Example code generation | [Details](#2-sdk-example---example-code-generation) |
| | `/sdk-test` | Test suite generation | [Details](#3-sdk-test---sdk-test-suite) |
| | `/sdk-doc` | SDK documentation generation | [Details](#4-sdk-doc---sdk-documentation-generation) |
| | `/sdk-release` | Release preparation check | [Details](#5-sdk-release---release-preparation-assistant) |

---

## 🚀 Quick Start

### Get Started with Claude Code in 5 Minutes

#### A. New Project
```bash
# Use meta workflow command
/meta

# Claude will:
# 1. Analyze project characteristics
# 2. Ask for key information
# 3. Auto-generate CLAUDE.md
# 4. Set up documentation structure
```

#### B. Existing Project
```bash
# Directly understand the project
/start

# Or restore previous work state
/sync
```

#### C. Install Command System
```bash
# Windows
cd claude\commands\deploy-package
.\deploy.ps1

# macOS/Linux
cd claude/commands/deploy-package
./deploy.sh
```

---

## 🎯 Command System v3.3 + SDK Extension

### Core Philosophy
- **Less is More**: 14 universal commands + 5 SDK-specific commands
- **Intelligent Integration**: Each command completes multiple related tasks
- **Context Aware**: Automatically manages memory and state
- **Scenario Adapted**: Dual-track support for app and SDK development

### 14 Universal Commands (Globally Available)

#### Project Understanding & Management (3)
| Command | Function | When to Use | Parameters |
|---------|----------|-------------|------------|
| `/start` | Quick project understanding | First contact with project | None |
| `/context` | Context sync checkpoint | Ensure consistent understanding | None |
| `/sync` | State synchronizer | Start of new session | None |

#### Development Aid (4)
| Command | Function | When to Use | Parameters |
|---------|----------|-------------|------------|
| `/plan` | Task planning & design | Before new feature | [task description] |
| `/check` | Complete quality check | Before code commit | None |
| `/watch` | Guardian mode | During coding | [on\|off\|status\|report] |
| `/test` | Test generation & execution | Ensure code quality | [file\|feature] |

#### Knowledge Management (2)
| Command | Function | When to Use | Parameters |
|---------|----------|-------------|------------|
| `/learn` | Learn and record decisions | After important decisions | [decision content] |
| `/doc` | Intelligent documentation maintenance | Update project docs | [api\|readme\|changelog\|arch] |

#### Workflow Optimization (3)
| Command | Function | When to Use | Parameters |
|---------|----------|-------------|------------|
| `/review` | PR preparation assistant | Before creating PR | None |
| `/debug` | Intelligent debugging assistant | When encountering issues | [error message] |
| `/meta` | Project specification customization | New project or major changes | None |

#### Quality Assurance (2)
| Command | Function | When to Use | Parameters |
|---------|----------|-------------|------------|
| `/analyze` | Deep analysis & validation | Experience-based risk analysis | [feature/module] [concerns or "deep"] |
| `/update-spec` | CLAUDE.md update specialist | Solidify decisions into specs | [review\|section "content"] |

### 🆕 SDK Development Commands (5)

Designed specifically for SDK/Library development:

#### SDK Specific Commands
| Command | Function | When to Use | Parameters |
|---------|----------|-------------|------------|
| `/sdk-design` | API design assistant | Designing new APIs | [feature description] |
| `/sdk-example` | Example code generation | Creating usage examples | [basic\|advanced\|integration\|all] |
| `/sdk-test` | SDK test suite | Generate professional tests | [unit\|integration\|compat\|performance\|all] |
| `/sdk-doc` | SDK documentation generation | Writing documentation | [api\|guide\|migration\|all] |
| `/sdk-release` | Release preparation assistant | Preparing new version | [major\|minor\|patch\|check] |

---

## 📁 Project Structure

```
claude/
├── README.md                    # Main documentation (English)
├── README_CN.md                 # Chinese documentation / 中文文檔
│
├── constitution/                # Constitution system (reference)
│   └── CLAUDE_CONSTITUTION.md  # Complete Claude collaboration constitution
│
├── commands/                    # Command system
│   └── deploy-package/          # Command deployment package v3.0
│       ├── DEPLOY_GUIDE_EN.md   # Deployment guide
│       ├── CHANGELOG_EN.md      # Version history
│       ├── SIMPLE_COMMANDS_SUMMARY_EN.md # Command system description
│       ├── deploy.ps1           # Windows deployment script
│       ├── deploy.sh            # macOS/Linux deployment script
│       ├── global/              # 14 universal commands
│       │   ├── analyze.md
│       │   ├── check.md
│       │   ├── context.md
│       │   ├── debug.md
│       │   ├── doc.md
│       │   ├── watch.md
│       │   ├── learn.md
│       │   ├── meta.md
│       │   ├── plan.md
│       │   ├── review.md
│       │   ├── start.md
│       │   ├── sync.md
│       │   ├── test.md
│       │   └── update-spec.md
│       └── sdk/                 # 5 SDK-specific commands
│           ├── sdk-design.md
│           ├── sdk-doc.md
│           ├── sdk-example.md
│           ├── sdk-release.md
│           └── sdk-test.md
│
├── guides/                      # In-depth guides (advanced reference)
│   ├── AI_ASSISTANT_COMPARISON.md     # AI assistant comparison
│   ├── COMMAND_WRITING_GUIDE.md       # Command writing guide
│   ├── CONSTITUTION_SYNC_GUIDE.md     # Constitution sync guide
│   ├── CONSTITUTION_USAGE_GUIDE.md    # Constitution usage guide
│   ├── DOCUMENT_STRUCTURE_STANDARD.md # Document structure standard
│   ├── LEGACY_PROJECT_ONBOARDING.md   # Legacy project onboarding
│   ├── MARKET_ANALYSIS.md             # Market analysis
│   ├── NEW_VS_LEGACY_PROJECT.md       # New vs legacy projects
│   └── SDK_DEVELOPMENT_WORKFLOW.md    # SDK development workflow
│
└── templates/                   # Template files
    ├── CLAUDE_MD_TEMPLATE.md    # CLAUDE.md general template
    └── SDK_PROJECT_TEMPLATE.md  # SDK project template
```

### 📝 Documentation Notes
- **README.md**: This file - Main documentation in English
- **README_CN.md**: Chinese version / 中文版本
- **commands/deploy-package/**:
  - English: [Deployment Guide](commands/deploy-package/DEPLOY_GUIDE_EN.md) | [Commands Summary](commands/deploy-package/SIMPLE_COMMANDS_SUMMARY_EN.md) | [Changelog](commands/deploy-package/CHANGELOG_EN.md)
- **guides/**: In-depth topical guides for reference
- **templates/**: Templates for project initialization

---

## 🏛️ Collaboration Constitution

### Core Principles
Claude Code collaboration is based on:

1. **Context First**: Maintain continuity of understanding
2. **Knowledge Accumulation**: Record decisions, avoid repetition
3. **Progressive Improvement**: Small steps, continuous optimization
4. **Human-AI Collaboration**: Clear division of labor, leverage respective strengths

### Working Modes

#### Development Workflow
1. **Understanding Phase**: `/start` or `/sync`
2. **Planning Phase**: `/plan` task decomposition
3. **Implementation Phase**: Coding implementation
4. **Validation Phase**: `/check` and `/test`
5. **Knowledge Consolidation**: `/learn` record decisions

#### Meta Workflow
For establishing and updating project specifications:
1. **Evaluate Project**: Tech stack, team, complexity
2. **Customize Specifications**: Generate project-specific CLAUDE.md
3. **Continuous Optimization**: Adjust specifications based on practice

---

## 📖 Usage Guide

### 📘 Command Reference Guide

#### 1. `/start` - Quick Project Understanding

**Use Cases**:
- First contact with a new project
- Need to quickly understand project structure
- Taking over someone else's codebase

**Usage**:
```bash
/start
```

**Expected Output**:
- Project type identification (Web/API/SDK/Tool etc.)
- Tech stack analysis (language, framework, dependencies)
- Directory structure parsing
- Key file location
- Auto-create `.claude/PROJECT_CONTEXT.md`

**Real Example**:
```bash
# Taking over a React project
/start
> Identified as: React Web Application
> Tech Stack: React 18, TypeScript, Vite
> Main Modules: components/, pages/, services/
> Entry File: src/main.tsx
> Project context file created
```

---

#### 2. `/sync` - State Synchronizer

**Use Cases**:
- Starting a new work session
- After switching branches
- Resuming work after interruption

**Usage**:
```bash
/sync
```

**Expected Output**:
- Load project context
- Restore decision records
- Show current work state
- Remind pending tasks

**Real Example**:
```bash
/sync
> Restoring project: E-commerce Platform v2.0
> Current branch: feature/payment
> Recent decision: Chose Stripe for payments
> Pending: Payment callback handling
```

---

#### 3. `/context` - Context Sync Checkpoint

**Use Cases**:
- Confirm Claude's understanding is correct
- Sync cognition after major changes
- Team member handover

**Usage**:
```bash
/context
```

**Expected Output**:
- Current understanding of project state
- Recent change summary
- Assumptions to confirm

---

#### 4. `/plan` - Task Planning & Design

**Use Cases**:
- Starting new feature development
- Refactoring existing code
- Solving complex problems

**Usage**:
```bash
/plan "task description"
```

**Expected Output**:
- Decomposed subtask list
- Implementation order suggestions
- Potential risk alerts
- Time estimates

**Real Example**:
```bash
/plan "implement shopping cart feature"
> Task breakdown:
> 1. Design cart data model (2h)
> 2. Implement add item API (1h)
> 3. Implement quantity update API (1h)
> 4. Implement remove item API (0.5h)
> 5. Add inventory check logic (1h)
> 6. Implement price calculation (1.5h)
> 7. Write unit tests (2h)
> Risk: Concurrent modifications may cause overselling
```

---

#### 5. `/check` - Complete Quality Check

**Use Cases**:
- Comprehensive check before code commit
- Self-check before Code Review
- Regular quality audit
- Use with `/watch`

**Usage**:
```bash
/check
```

**Expected Output**:
- Code style issues
- Potential bugs
- Performance optimization suggestions
- Security vulnerability warnings
- Quality score report

---

#### 6. `/watch` - Guardian Mode (Collaborative Quality Guard)

**Use Cases**:
- Continuous attention during coding
- Need to actively submit code to trigger checks
- Build good checking habits
- Form complete quality assurance with `/check`

**Usage**:
```bash
/watch on      # Enable guardian mode
/watch off     # Disable guardian mode
/watch status  # Check current status
/watch report  # Generate guardian report
```

**Expected Output**:
- Real-time security warnings
- Code quality reminders
- Performance risk alerts
- Best practice suggestions

**Important Note**:
ℹ️ `/watch` is not truly real-time monitoring - you need to actively submit code snippets to trigger checks. This is a collaborative working mode that helps maintain quality awareness during coding.

---

#### 7. `/test` - Test Generation & Execution

**Use Cases**:
- Writing tests for new features
- Supplementing test coverage
- Verifying bug fixes

**Usage**:
```bash
/test [file|feature]
```

**Expected Output**:
- Generated test code
- Test execution results
- Coverage report
- Boundary condition tests

---

#### 8. `/debug` - Intelligent Debugging Assistant

**Use Cases**:
- Need to locate errors
- Performance issue investigation
- Abnormal behavior analysis

**Usage**:
```bash
/debug "error message or problem description"
```

**Expected Output**:
- Problem cause analysis
- Possible solutions
- Debugging step suggestions
- Related code location

---

#### 9. `/learn` - Learn and Record Decisions

**Use Cases**:
- After important technology selection
- After solving key problems
- When discovering best practices

**Usage**:
```bash
/learn "decision content or experience"
```

**Expected Output**:
- Update DECISIONS.md
- Intelligent category tags
- Related impact analysis

---

#### 10. `/doc` - Intelligent Documentation Maintenance

**Use Cases**:
- Update API documentation
- Maintain README
- Generate changelog

**Usage**:
```bash
/doc [api|readme|changelog|arch]
```

**Expected Output**:
- Auto-update specified documentation
- Preserve manually edited content
- Generate missing parts

---

#### 11. `/review` - PR Preparation Assistant

**Use Cases**:
- Before creating Pull Request
- Need self-review
- Preparing for code review

**Usage**:
```bash
/review
```

**Expected Output**:
- Change summary
- PR description template
- Checklist
- Potential issue reminders

---

#### 12. `/meta` - Project Specification Customization

**Use Cases**:
- New project initialization
- Team specification establishment
- Tech stack changes

**Usage**:
```bash
/meta
```

**Expected Output**:
- Generate CLAUDE.md
- Project-specific specifications
- Workflow definitions

---

#### 13. `/analyze` - Deep Analysis & Validation

**Use Cases**:
- When feature is complete but have concerns
- Need risk assessment
- Performance bottleneck analysis

**Usage**:
```bash
/analyze "feature/module" ["specific concerns" or "deep"]
```

**Expected Output**:
- Risk level assessment
- Boundary condition analysis
- Improvement suggestions
- Test scenarios

---

#### 14. `/update-spec` - CLAUDE.md Update Specialist

**Use Cases**:
- Solidify important decisions into specifications
- Update project rules
- Periodic specification review

**Usage**:
```bash
/update-spec [review|section "content"]
```

**Expected Output**:
- Specification update suggestions
- Version change log
- Conflict detection

---

#### 📦 SDK Command Details

#### 1. `/sdk-design` - API Design Assistant

**Use Cases**:
- Designing new SDK interfaces
- Refactoring existing APIs
- Establishing design standards

**Usage**:
```bash
/sdk-design "feature description"
```

**Expected Output**:
- API structure suggestions
- Naming conventions
- Parameter design
- Error handling strategy

---

#### 2. `/sdk-example` - Example Code Generation

**Use Cases**:
- Creating usage examples for SDK
- Writing quickstart guides
- Demonstrating best practices

**Usage**:
```bash
/sdk-example [basic|advanced|integration|all]
```

---

#### 3. `/sdk-test` - SDK Test Suite

**Use Cases**:
- Generate complete test suite
- Compatibility testing
- Performance benchmark testing

**Usage**:
```bash
/sdk-test [unit|integration|compat|performance|all]
```

---

#### 4. `/sdk-doc` - SDK Documentation Generation

**Use Cases**:
- Generate API reference documentation
- Write usage guides
- Create migration documentation

**Usage**:
```bash
/sdk-doc [api|guide|migration|all]
```

---

#### 5. `/sdk-release` - Release Preparation Assistant

**Use Cases**:
- Prepare new version release
- Check release checklist
- Generate release notes

**Usage**:
```bash
/sdk-release [major|minor|patch|check]
```

---

## 🔄 Workflows

### Development Flow Comparison

#### Application Development Flow
| Phase | Traditional | Claude Code v3.0 |
|-------|-------------|------------------|
| Start | Manual context | `/sync` auto-restore |
| Understanding | Repeat explanation | `/context` sync confirm |
| Planning | Free discussion | `/plan` structured design |
| Development | Solo coding | AI-assisted implementation |
| Testing | Manual writing | `/test` intelligent generation |
| Validation | Experience-based | `/analyze` deep analysis |
| Review | Manual check | `/check` auto review |
| Documentation | Post-hoc supplement | `/doc` sync update |
| Knowledge | Easy to forget | `/learn` persistent record |

#### SDK Development Flow
| Phase | Traditional | Claude Code + SDK Commands |
|-------|-------------|----------------------------|
| API Design | Experience-based | `/sdk-design` professional guidance |
| Example Writing | Manual creation | `/sdk-example` auto generation |
| Test Strategy | Basic testing | `/sdk-test` comprehensive coverage |
| Documentation | Time-consuming | `/sdk-doc` structured generation |
| Version Release | Easy to miss | `/sdk-release` complete check |

### Typical Workflows

#### 1. New Project Initialization

**Application Development**:
```bash
/meta               # Establish project specifications
/start              # Understand project structure
/plan "core feature"  # Plan first task
```

**SDK Development**:
```bash
/meta               # Establish SDK specifications
/sdk-design "core API"  # Design interfaces
/plan "implement core"  # Plan implementation
```

#### 2. Daily Development Cycle
```bash
/sync               # Restore work state
/context            # Confirm understanding
/plan "new feature"   # Plan implementation
# ... coding ...
/check              # Quality check
/test               # Generate and run tests
/learn "tech decision"  # Record important decisions
```

#### 3. Commit & Release
```bash
/check              # Final quality check
/doc                # Update documentation
/review             # Prepare PR
```

#### 4. Problem Solving
```bash
/debug "error message"  # Quick problem location
/test feature       # Verify fix
```

---

## 💡 Best Practices

### 1. Communication Tips
- **Clear Boundaries**: Tell Claude what not to modify
- **Provide Examples**: Give expected code style
- **Segmented Confirmation**: Break complex tasks into checkpoints
- **Record Decisions**: Write important choices into documentation

### 2. Project Organization
```
your-project/
├── .claude/
│   ├── commands/           # Project-specific commands
│   ├── PROJECT_CONTEXT.md  # Project context
│   ├── DECISIONS.md        # Decision records
│   └── state/              # State files
├── CLAUDE.md               # Project specifications
└── ... project files
```

### 3. Efficiency Boost
- **Start with Sync**: Always start with `/sync`
- **Timely Recording**: Use `/learn` to avoid knowledge loss
- **Structured Planning**: Use `/plan` instead of free discussion
- **Automated Checking**: Use `/check` to ensure quality
- **Deep Validation**: Use `/analyze` to verify intuition
- **Specification Management**: Use `/update-spec` to solidify decisions

### 4. Team Collaboration
- Share `.claude/` directory
- Unified use of command system
- Regularly update PROJECT_CONTEXT.md
- Record important decisions in DECISIONS.md

---

## 📊 Version History

### v3.3.0 (2025-08-10) - Current Version
- **Major Adjustment**:
  - Removed all project-level commands, focus on global universal commands
  - Let developers create their own project-specific commands
  - Unified use of 14 global universal commands + 5 SDK-specific commands

### v3.2.2 (2025-08-10)
- **Architecture Fix**:
  - Correctly categorized `/analyze` and `/update-spec` as global commands
  - Corrected universal command count to 14
  - Cleaned up duplicate command files

### v3.2.1 (2025-08-10)
- **Documentation Enhancement**:
  - Added complete command usage manual (18 command details)
  - Each command includes: use cases, usage, expected output, real examples
  - Added 6 typical command combination scenarios

### v3.2.0 (2025-08-10)
- **New Feature**: Deep analysis validation command
  - Created `/analyze` command to fill validation gap
  - Support experience-based risk analysis
  - Provide quantified risk assessment

### v3.1.0 (2024-01-20)
- **New Feature**: Dedicated CLAUDE.md update command
  - Created `/update-spec` command for updating project specifications
  - Support two modes: review mode and targeted update mode

### v3.0.0 (2024-01-15)
- **Major Refactor**: Command system streamlining
- **Core Improvements**:
  - Commands reduced from 31 to 11
  - Intelligent command integration
  - Automated memory management

---

## 🤝 Contributing

Contributions are welcome! Please:
1. Fork this project
2. Create a feature branch
3. Commit your changes
4. Submit a Pull Request

### Reporting Issues
- Use GitHub Issues
- Provide detailed reproduction steps
- Describe expected behavior

---

## 📚 Related Resources

- [Claude Code Official Documentation](https://docs.anthropic.com/en/docs/claude-code)
- [Command System Documentation](https://docs.anthropic.com/en/docs/claude-code/slash-commands)
- [MCP Protocol](https://docs.anthropic.com/en/docs/claude-code/mcp)
- [Project Hook Template](PROJECT_HOOKS_TEMPLATE.md)

---

## 🎯 Core Values

1. **Simple & Efficient**: Fewer commands, higher efficiency
2. **Intelligent Collaboration**: AI understands you, you guide AI
3. **Knowledge Accumulation**: Every decision is wealth
4. **Continuous Evolution**: Constantly optimize based on usage

---

*This collaboration framework was jointly created by Claude and human developers, continuously evolving.*

*Simple, Intelligent, Efficient - Let Claude Code become your best development partner!*