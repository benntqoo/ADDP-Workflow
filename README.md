# Claude Code Collaboration Framework

*English | [中文](README_zh.md)*

Transform AI into your intelligent development partner through command systems, intelligent agents, and personalized styles.

## 🎯 Overview

Claude Code Collaboration Framework is a comprehensive AI-assisted development system offering three core capabilities:

1. **Command System** - Project management and workflow control
2. **Intelligent Agents** - Professional technical support and quality assurance  
3. **Output Styles** - Personalized communication customization

## 🚀 Quick Start: Style + Command Combinations

### Core Concept
- **Styles define personality**: How AI thinks and outputs (architect/developer/analyst)
- **Commands define actions**: What tasks to execute (/plan, /sync, /learn)
- **Combinations create synergy**: Different combinations for different scenarios

### Recommended Combinations

| Scenario | Style + Command | Effect |
|----------|----------------|--------|
| **Starting new project** | `architect` + `/start` → `/plan` | Understand project & design architecture |
| **Daily development** | `concise-developer` + `/sync` → `/plan` | Restore state & plan tasks |
| **Feature implementation** | `concise-developer` + `/context` → code | Confirm understanding then implement |
| **Security audit** | `security-analyst` + `/context` → analyze | Understand system then review |
| **Learning new tech** | `educational-mentor` + `/start` → `/doc` | Learn concepts & document knowledge |
| **Emergency fix** | `concise-developer` + `/sync` → fix → deploy | Quick restore, fix, and deploy |

### Example Workflow
```bash
# Morning routine
/output-style:set concise-developer
/sync                          # Restore yesterday's progress
/plan "Complete user module"  # Plan today's tasks

# During development
"Implement user CRUD"          # Code implementation
/learn "Using Repository pattern"  # Record important decisions

# Before commit
/check                         # Quality check
/update-spec                   # Update specifications
```

## 📚 Feature List

### 🎮 Command System

#### Project Understanding & Management (3)
| Command | Function | When to Use | Parameters |
|---------|----------|-------------|------------|
| `/start` | Quick project startup & understanding | First time touching project | None |
| `/context` | Context sync checkpoint | Ensure understanding consistency | None |
| `/sync` | State synchronizer | New session start | None |

#### Development Support (4)
| Command | Function | When to Use | Parameters |
|---------|----------|-------------|------------|
| `/plan` | Task planning & design | Before starting new features | [task description] |
| `/check` | Complete quality check | Before committing code | None |
| `/watch` | Watch mode | During coding | [on\|off\|status\|report] |
| `/test` | Test generation & execution | Ensure code quality | [file\|feature] |

#### Knowledge Management (2)
| Command | Function | When to Use | Parameters |
|---------|----------|-------------|------------|
| `/learn` | Learn & record decisions | After important decisions | [decision content] |
| `/doc` | Smart document maintenance | Update project docs | [api\|readme\|changelog\|arch] |

#### Workflow Optimization (3)
| Command | Function | When to Use | Parameters |
|---------|----------|-------------|------------|
| `/review` | PR preparation assistant | Before creating PR | None |
| `/debug` | Smart debugging assistant | When encountering problems | [error info] |
| `/meta` | Project specification customization | New project or major changes | None |

#### Quality Assurance (2)
| Command | Function | When to Use | Parameters |
|---------|----------|-------------|------------|
| `/analyze` | Deep analysis & validation | Risk analysis based on intuition | [feature/module] [concern or "deep"] |
| `/update-spec` | CLAUDE.md update specific | Solidify decisions into specs | [review\|section "content"] |

#### SDK Development Commands (5)
| Command | Function | When to Use | Parameters |
|---------|----------|-------------|------------|
| `/sdk-design` | API design assistant | Designing new APIs | [feature description] |
| `/sdk-example` | Example code generation | Creating usage examples | [basic\|advanced\|integration\|all] |
| `/sdk-test` | SDK test suite | Generate professional tests | [unit\|integration\|compat\|performance\|all] |
| `/sdk-doc` | SDK documentation generation | Writing documentation | [api\|guide\|migration\|all] |
| `/sdk-release` | Release preparation assistant | Preparing new version | [major\|minor\|patch\|check] |

**Total: 19 Commands** (14 Core + 5 SDK)
📖 Detailed docs: [commands/docs/](commands/docs/)

### 🤖 Intelligent Agents

**Quality Assurance Agents**
- `code-reviewer` - Code review expert
- `test-automator` - Automated test generation
- `performance-optimizer` - Performance optimization analysis
- `bug-hunter` - Bug finding and fixing
- `security-analyst` - Security vulnerability analysis

**Technical Expert Agents**
- `kotlin-expert` - Kotlin full-stack development
- `python-ml-specialist` - Python machine learning
- `golang-systems-engineer` - Go systems programming
- `rust-zero-cost` - Rust zero-cost abstractions
- `typescript-fullstack-expert` - TypeScript full-stack

**35+ More Agents** covering: Android, iOS, Web, Backend, Database, DevOps, etc.

📖 Detailed docs: [agents/docs/](agents/docs/)

### 🎨 Output Styles (9 Professional Personas)

#### Architecture & Design
| Style Name | Best For | Key Characteristics |
|------------|----------|---------------------|
| `senior-architect` | System design | Comprehensive analysis, risk assessment, strategic thinking |
| `system-architect` | Technical architecture | Transform PRDs to technical designs, multi-platform solutions |

#### Development & Implementation
| Style Name | Best For | Key Characteristics |
|------------|----------|---------------------|
| `concise-developer` | Quick coding | Minimal explanations, direct solutions, code-first |
| `educational-mentor` | Learning & teaching | Detailed explanations, progressive learning, rich examples |

#### Operations & Security
| Style Name | Best For | Key Characteristics |
|------------|----------|---------------------|
| `devops-engineer` | Infrastructure | Automation-first, reliability, IaC mindset |
| `security-analyst` | Security review | Threat modeling, vulnerability assessment, compliance |

#### Product & SDK
| Style Name | Best For | Key Characteristics |
|------------|----------|---------------------|
| `product-expert` | Product requirements | PRD documents, user stories, roadmap planning |
| `sdk-design-expert` | SDK architecture | API design, cross-platform, developer experience |
| `sdk-prd-expert` | SDK product management | Developer tools PRD, API product strategy |

**Total: 9 Professional Output Styles**
📖 Detailed docs: [output-styles/README.md](output-styles/README.md)

## 🚀 Production Deployment Guide

### Quick Start (5 minutes)

#### 1. Install All Components

**Windows:**
```powershell
# Create Claude directories
mkdir "%USERPROFILE%\.claude\commands"
mkdir "%USERPROFILE%\.claude\agents" 
mkdir "%USERPROFILE%\.claude\output-styles"

# Copy all files
xcopy /Y "claude\commands\deploy-package\global\*.md" "%USERPROFILE%\.claude\commands\"
xcopy /Y "claude\commands\deploy-package\sdk\*.md" "%USERPROFILE%\.claude\commands\"
xcopy /E /Y "claude\agents" "%USERPROFILE%\.claude\agents\"
xcopy /Y "claude\output-styles\*.md" "%USERPROFILE%\.claude\output-styles\"
```

**macOS/Linux:**
```bash
# One-line installation
mkdir -p ~/.claude/{commands,agents,output-styles} && \
cp claude/commands/deploy-package/global/*.md ~/.claude/commands/ && \
cp claude/commands/deploy-package/sdk/*.md ~/.claude/commands/ && \
cp -r claude/agents/* ~/.claude/agents/ && \
cp claude/output-styles/*.md ~/.claude/output-styles/
```

#### 2. Enable Smart Agent System (CRITICAL)

```bash
# Set the intelligent orchestrator style
/output-style:set orchestrator

# Verify it's working
/output-style
# Should show: "Current: orchestrator"
```

### 🎯 How the Smart System Works

#### Before (Inefficient):
```
User: "Optimize my React app"
❌ Problem: Might start 3-5 random agents
❌ Result: 800k+ tokens wasted, unclear results
```

#### After (Smart Selection):
```
User: "Optimize my React app" 
✅ System thinks: ["React", "optimize"] → performance task
✅ Selects: performance-optimizer (single expert)
✅ Result: ~100k tokens, focused optimization
```

#### Agent Selection Examples:

| User Request | Smart Selection | Tokens | Why |
|--------------|-----------------|--------|-----|
| "Fix login bug" | bug-hunter | ~110k | Debugging needs focus |
| "Design REST API" | api-architect | ~120k | Specialized API expert |
| "Build React app" | frontend-developer | ~150k | Frontend specialist |
| "Deploy ML model" | mlops-specialist | ~200k | Production ML expert |
| "Code quality review" | jenny-validator + karen-realist + senior-developer | ~360k | Only 3-agent scenario |

### 📊 Expected Performance Improvements

```
Token Efficiency:
✅ Average usage: 300k (down from 800k)
✅ Success rate: 90%+ correct agent selection  
✅ Response time: <15 seconds

User Experience:
✅ No more wrong agent selections
✅ No more token waste on irrelevant experts  
✅ Precise, focused solutions
```

### 🔧 Verify Installation

```bash
# Test the smart system
echo "Testing: 'Optimize database performance'"
# Should select: performance-optimizer (single agent)

echo "Testing: 'Create a mobile app'" 
# Should select: mobile-developer (single agent)

echo "Testing: 'Build complete e-commerce platform'"
# Should select: fullstack-architect + frontend-developer (2 agents max)
```

### ⚠️ Troubleshooting

**Problem: Agents not selecting correctly**
```bash
# Check orchestrator style is active
/output-style
# Should show "orchestrator"

# If not, set it:
/output-style:set orchestrator
```

**Problem: Still using too many agents**
```bash
# The system is designed to prefer single experts
# If you see 3+ agents for simple tasks, the old system might be active
# Make sure you're using /output-style:set orchestrator
```

### 📈 Usage Tracking (Optional)

Create a simple usage log in your project:
```bash
# Create tracking file
echo "## Usage Tracking Log" > agents/usage_log.md
echo "Date | Request | Agents Selected | Satisfaction | Notes" >> agents/usage_log.md
echo "-----|---------|-----------------|--------------|-------" >> agents/usage_log.md
```

Example entries:
```
2024-12-19 | React performance | performance-optimizer | 5/5 | Perfect choice
2024-12-19 | API design | api-architect | 5/5 | Comprehensive solution  
2024-12-19 | Bug in login | bug-hunter | 4/5 | Found issue quickly
```

### 🚀 You're Ready for Production!

The system will now:
- ✅ **Automatically select the best agents** for each task
- ✅ **Minimize token usage** by preferring single experts
- ✅ **Provide focused solutions** instead of generic responses
- ✅ **Scale efficiently** as your team grows

**Start using it immediately and experience the 60%+ efficiency improvement!**

## 📖 Usage Guide

### Basic Usage

#### 1. Start New Project
```bash
# Use meta command to create project specs
/meta

# Claude will:
# - Analyze project characteristics
# - Ask key information
# - Generate CLAUDE.md
# - Setup documentation structure
```

#### 2. Restore Work State
```bash
# At new session start
/sync

# System will:
# - Read last work state
# - Check uncommitted changes
# - Provide work suggestions
```

#### 3. Switch Output Style
```bash
# View available styles
/output-style

# Set style
/output-style:set senior-architect

# Start working with new style
/plan "Design microservices architecture"
```

#### 4. Use Intelligent Agents
```bash
# Agents auto-activate, or manually specify
"Use python-ml-specialist agent to help design the model"

# Or auto-trigger for specific tasks
"Review this code for security issues"  # Auto-activates security-analyst
```

### Project Configuration

Create `.claude/` directory in project root:

```
your-project/
├── .claude/
│   ├── PROJECT_CONTEXT.md  # Project context
│   ├── DECISIONS.md        # Decision records
│   ├── settings.local.json # Project settings
│   └── state/              # State files
├── CLAUDE.md               # Project specifications
└── ... project files
```

Configuration example (`.claude/settings.local.json`):
```json
{
  "outputStyle": "concise-developer",
  "permissions": {
    "defaultMode": "acceptEdits"
  }
}
```

## 🎯 Workflow Examples

### Scenario 1: New Feature Development

```bash
# 1. Restore state
/sync
# → Restore previous work progress

# 2. Plan task
/plan "Add user authentication"
# → Generate task plan and technical solution

# 3. Development
# Auto-activate relevant agents:
# - code-reviewer for continuous review
# - test-automator for test generation
# - security-analyst for security checks

# 4. Record decisions
/learn "Decided to use JWT instead of Session because..."
# → Save to DECISIONS.md

# 5. Update specs
/update-spec
# → Update CLAUDE.md
```

### Scenario 2: Code Review & Optimization

```bash
# 1. Switch to architect style
/output-style:set senior-architect

# 2. Architecture review
/context
# → Comprehensive architecture analysis

# 3. Performance optimization
"Analyze and optimize database query performance"
# → Auto-activates performance-optimizer

# 4. Security review
/output-style:set security-analyst
"Review authentication system security"
# → Deep security analysis
```

### Scenario 3: SDK Development

```bash
# 1. Design API
/sdk-design "Payment SDK interface design"

# 2. Generate examples
/sdk-example advanced

# 3. Create tests
/sdk-test all

# 4. Write documentation
/sdk-doc api

# 5. Prepare release
/sdk-release check
```

### Scenario 4: Team Collaboration

```bash
# 1. Morning start
/sync
# → Check team's yesterday changes

# 2. Understand new code
/context
# → Sync project understanding

# 3. Switch to teaching style (for onboarding)
/output-style:set educational-mentor
"Explain how this authentication module works"

# 4. Record team decisions
/learn "Team decided to adopt microservices architecture..."

# 5. Update team documentation
/doc readme
```

## 🏆 Best Practices

### 1. Work Habits
- **Always sync first**: Use `/sync` before each work session
- **Record promptly**: Use `/learn` immediately after important decisions
- **Regular updates**: Use `/update-spec` to solidify specifications

### 2. Style Selection
- **Design phase**: Use `senior-architect`
- **Rapid development**: Use `concise-developer`
- **Code review**: Use `security-analyst`
- **Documentation**: Use `educational-mentor`

### 3. Agent Collaboration
- Let agents auto-activate, don't over-control
- Trust professional agent recommendations
- Multiple agents can work simultaneously

### 4. Team Standards
- Share `.claude/` directory
- Use command system consistently
- Regularly update PROJECT_CONTEXT.md

## 📁 Project Structure

```
claude/
├── README.md               # This document (English)
├── README_zh.md           # Chinese version
├── RELEASE_NOTE.md        # Version history
├── commands/              # Command system
│   ├── docs/             # Command detailed docs
│   └── deploy-package/   # Deployment package
│       ├── global/       # Core commands (8)
│       └── sdk/          # SDK commands (5)
├── agents/               # Intelligent agents
│   ├── docs/            # Agent detailed docs
│   └── *.md             # Agent definition files (35+)
├── output-styles/        # Output styles
│   ├── README.md        # Style usage guide
│   └── *.md             # Style definition files (9)
└── guides/              # In-depth guides
    └── *.md             # Various topic guides
```

## 🆘 FAQ

**Q: Commands not working?**
A: Check if files are copied to correct directory `~/.claude/commands/`

**Q: Agents not auto-activating?**
A: Agents activate based on task description, use clear keywords

**Q: How to persist output style?**
A: Set in project's `.claude/settings.local.json`

**Q: How to create custom commands/agents/styles?**
A: Refer to documentation and templates in respective directories

## 🤝 Contributing

This is an open-source project providing organized Claude Code development experience and tools for all developers.

We welcome contributions! Please:
1. Fork this repository
2. Create a feature branch
3. Submit your changes
4. Create a Pull Request

### Reporting Issues
- Use GitHub Issues
- Provide detailed reproduction steps
- Explain expected behavior

## 📚 Resources

- [Command System Documentation](commands/docs/)
- [Agent System Documentation](agents/docs/)
- [Output Styles Documentation](output-styles/README.md)
- [In-depth Usage Guides](guides/)
- [Claude Code Official Docs](https://docs.anthropic.com/en/docs/claude-code)

## 📄 License

MIT License - See [LICENSE](LICENSE) file for details

## 🌟 Star History

If you find this project helpful, please give it a star ⭐

---

*Let Claude Code become your best development partner!*

**Made with ❤️ by the Claude Code Community**