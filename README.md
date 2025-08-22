# Claude Code Collaboration Framework

*English | [中文](README_zh.md)*

Transform AI into your intelligent development partner through command systems, intelligent agents, and personalized styles.

## 🎯 Overview

Claude Code Collaboration Framework is a comprehensive AI-assisted development system offering three core capabilities:

1. **Command System** - Project management and workflow control
2. **Intelligent Agents** - Professional technical support and quality assurance  
3. **Output Styles** - Personalized communication customization

## 📚 Feature List

### 🎮 Command System

**Core Commands (8)**

| Command | Function | Use Case |
|---------|----------|----------|
| `/start` | Quick project understanding | First time working with project |
| `/sync` | Restore work state | Beginning new session |
| `/context` | Sync context understanding | Ensure understanding consistency |
| `/plan` | Task planning & design | Before starting new features |
| `/learn` | Record decision knowledge | After important decisions |
| `/meta` | Customize project specs | New project or major adjustments |
| `/doc` | Smart document maintenance | Update project docs |
| `/update-spec` | Update CLAUDE.md | Solidify decisions into specs |

**SDK Development Commands (5)**

| Command | Function | Parameters |
|---------|----------|------------|
| `/sdk-design` | API design assistant | [feature description] |
| `/sdk-example` | Example code generation | basic/advanced/all |
| `/sdk-test` | Test suite generation | unit/integration/all |
| `/sdk-doc` | Documentation generation | api/guide/migration/all |
| `/sdk-release` | Release preparation | major/minor/patch/check |

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

### 🎨 Output Styles

| Style Name | Best For | Key Characteristics |
|------------|----------|---------------------|
| `senior-architect` | System design | Comprehensive analysis, risk assessment, strategic thinking |
| `concise-developer` | Quick coding | Minimal explanations, direct solutions, code-first |
| `educational-mentor` | Learning & teaching | Detailed explanations, progressive learning, rich examples |
| `devops-engineer` | Infrastructure | Automation-first, reliability, IaC mindset |
| `security-analyst` | Security review | Threat modeling, vulnerability assessment, compliance |
| `product-expert` | Product requirements | PRD documents, user stories, roadmap planning |
| `sdk-design-expert` | SDK design | API design, cross-platform, developer experience |

📖 Detailed docs: [output-styles/README.md](output-styles/README.md)

## 🚀 Manual Installation

### 1. Command System

**Windows:**
```powershell
# Create commands directory
mkdir "%USERPROFILE%\.claude\commands"

# Copy core commands
xcopy /Y "claude\commands\deploy-package\global\*.md" "%USERPROFILE%\.claude\commands\"

# Copy SDK commands (optional)
xcopy /Y "claude\commands\deploy-package\sdk\*.md" "%USERPROFILE%\.claude\commands\"
```

**macOS/Linux:**
```bash
# Create commands directory
mkdir -p ~/.claude/commands

# Copy core commands
cp claude/commands/deploy-package/global/*.md ~/.claude/commands/

# Copy SDK commands (optional)
cp claude/commands/deploy-package/sdk/*.md ~/.claude/commands/
```

### 2. Intelligent Agents

**Windows:**
```powershell
# Create agents directory
mkdir "%USERPROFILE%\.claude\agents"

# Copy all agents
xcopy /E /Y "claude\agents\*.md" "%USERPROFILE%\.claude\agents\"
```

**macOS/Linux:**
```bash
# Create agents directory
mkdir -p ~/.claude/agents

# Copy all agents
cp -r claude/agents/*.md ~/.claude/agents/
```

### 3. Output Styles

**Windows:**
```powershell
# Create styles directory
mkdir "%USERPROFILE%\.claude\output-styles"

# Copy all styles
xcopy /Y "claude\output-styles\*.md" "%USERPROFILE%\.claude\output-styles\"
```

**macOS/Linux:**
```bash
# Create styles directory
mkdir -p ~/.claude/output-styles

# Copy all styles
cp claude/output-styles/*.md ~/.claude/output-styles/
```

### 4. Verify Installation

```bash
# Check installed files
ls ~/.claude/commands/       # Should see command files
ls ~/.claude/agents/         # Should see agent files
ls ~/.claude/output-styles/  # Should see style files
```

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