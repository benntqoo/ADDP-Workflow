# Command System Documentation

*English | [中文](README_zh.md)*

## Overview

The Claude Code Command System provides structured workflows for AI-assisted development. These commands help manage projects, maintain context, and ensure consistent collaboration between humans and AI.

## Core Commands (8)

### Project Management

#### `/start` - Project Quick Start
- **Purpose**: Quickly understand project structure and purpose
- **When to use**: First time working with a project
- **What it does**:
  - Scans project structure
  - Reads key files (README, package.json, etc.)
  - Creates initial understanding baseline
  - Generates PROJECT_CONTEXT.md

#### `/sync` - State Synchronizer
- **Purpose**: Restore previous work state
- **When to use**: Beginning of each work session
- **What it does**:
  - Reads last session state
  - Checks git status and recent commits
  - Analyzes current work focus
  - Provides work suggestions

#### `/context` - Context Checkpoint
- **Purpose**: Verify understanding consistency
- **When to use**: After complex changes or before critical tasks
- **What it does**:
  - Compares current understanding with baseline
  - Identifies knowledge gaps
  - Ensures alignment with project goals

### Development Support

#### `/plan` - Task Planning
- **Purpose**: Design and plan new features
- **When to use**: Before starting any new development
- **Parameters**: `[task description]`
- **What it does**:
  - Analyzes requirements
  - Creates technical design
  - Breaks down into subtasks
  - Identifies risks and dependencies

#### `/doc` - Document Maintenance
- **Purpose**: Keep documentation up-to-date
- **When to use**: After significant changes
- **Parameters**: `api|readme|changelog|arch`
- **What it does**:
  - Updates specified documentation
  - Maintains consistency
  - Generates missing sections

### Knowledge Management

#### `/learn` - Decision Recording
- **Purpose**: Capture important decisions and learnings
- **When to use**: After making technical decisions
- **Parameters**: `[decision content]`
- **What it does**:
  - Records to DECISIONS.md
  - Maintains decision history
  - Prevents knowledge loss

#### `/meta` - Project Specification
- **Purpose**: Define project-specific rules and conventions
- **When to use**: New project or major changes
- **What it does**:
  - Creates/updates CLAUDE.md
  - Establishes coding standards
  - Defines project vocabulary

#### `/update-spec` - Specification Update
- **Purpose**: Solidify decisions into project specifications
- **When to use**: After validating new patterns or approaches
- **Parameters**: `review|section "content"`
- **What it does**:
  - Updates CLAUDE.md
  - Ensures specifications reflect current practices

## SDK Development Commands (5)

### `/sdk-design` - API Design Assistant
- **Purpose**: Design consistent, user-friendly APIs
- **Parameters**: `[feature description]`
- **Output**: API specifications, method signatures, design patterns

### `/sdk-example` - Example Generation
- **Purpose**: Create comprehensive usage examples
- **Parameters**: `basic|advanced|integration|all`
- **Output**: Working code examples for different scenarios

### `/sdk-test` - Test Suite Generation
- **Purpose**: Generate comprehensive test coverage
- **Parameters**: `unit|integration|compat|performance|all`
- **Output**: Test files with various test scenarios

### `/sdk-doc` - Documentation Generation
- **Purpose**: Create professional SDK documentation
- **Parameters**: `api|guide|migration|all`
- **Output**: API references, guides, migration docs

### `/sdk-release` - Release Preparation
- **Purpose**: Prepare SDK for release
- **Parameters**: `major|minor|patch|check`
- **Output**: Version updates, changelog, release checklist

## Best Practices

1. **Start every session with `/sync`** - Never lose context
2. **Use `/plan` before coding** - Think before you build
3. **Record decisions with `/learn`** - Build institutional knowledge
4. **Update specs with `/update-spec`** - Keep documentation current
5. **Check understanding with `/context`** - Ensure alignment

## Command Workflow

```mermaid
graph LR
    A[New Session] --> B[/sync]
    B --> C{Task Type}
    C -->|New Feature| D[/plan]
    C -->|Bug Fix| E[/context]
    C -->|Documentation| F[/doc]
    D --> G[Development]
    E --> G
    G --> H[/learn]
    H --> I[/update-spec]
```

## File Structure Created

```
.claude/
├── PROJECT_CONTEXT.md    # Project understanding
├── DECISIONS.md          # Decision history
├── settings.local.json   # Project settings
└── state/
    ├── last-session.yml  # Session state
    └── initial-scan.json # Baseline understanding
```

## Tips

- Commands are stateless - each execution is independent
- All commands preserve existing content when updating files
- Use parameters to control command behavior
- Commands can be chained for complex workflows

---

For detailed usage examples, see the [Workflow Guide](../../guides/WORKFLOW_GUIDE.md)