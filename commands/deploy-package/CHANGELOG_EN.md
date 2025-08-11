# Changelog

## v3.3.0 (2025-08-10) - Command System Refinement

### 🎯 Major Changes
- Removed all project-level commands, focusing on global universal commands
- Let developers create their own project-specific commands
- Unified to 14 global universal commands + 5 SDK-specific commands

### 📦 Command Distribution
- **Universal Commands**: 14 (available globally)
- **SDK Commands**: 5 (for SDK/library development)
- **Project Commands**: Developer-defined (customizable per project)

### ⚡ Improvements
- Clearer command boundaries and responsibilities
- More focused and specialized functionality
- Better separation between general and specific needs
- Enhanced developer autonomy

---

## v3.2.2 (2025-08-10) - Architecture Correction

### 🔧 Fixes
- Correctly categorized `/analyze` and `/update-spec` as global commands
- Corrected universal command count to 14
- Cleaned up duplicate command files
- Fixed project structure documentation

---

## v3.2.1 (2025-08-10) - Documentation Enhancement

### 📚 Documentation
- Added complete command usage manual (18 command details)
- Each command includes: use cases, usage, expected output, real examples
- Added 6 typical command combination scenarios
- Provided advanced usage tips and efficiency comparison

### 💡 User Experience
- Clearer and more intuitive command descriptions
- Rich practical usage examples
- Added command chaining guide

---

## v3.2.0 (2025-08-10) - Deep Analysis Feature

### 🚀 New Features
- Created `/analyze` command for deep risk analysis
- Support for experience-based risk assessment
- Quantified risk evaluation and priority suggestions

### ⚡ Quality Assurance
- Automatic boundary condition analysis
- Special scenario deduction
- Architecture-level review
- Business logic validation

---

## v3.1.0 (2024-01-20) - Specification Management

### 🚀 New Features
- Created `/update-spec` command specifically for CLAUDE.md updates
- Support for two modes: review mode and targeted update mode
- Implemented single responsibility design for commands

### 🏗️ Architecture Optimization
- Clear command responsibility boundaries
- `/learn` only updates DECISIONS.md and PROJECT_CONTEXT.md
- `/update-spec` only updates CLAUDE.md
- Established clear command responsibility matrix

---

## v3.0.0 (2024-01-15) - Major Refactor

### 🚀 New Features
- Streamlined command system from 31 to 11 core commands
- Intelligent command integration, one command completes multiple related tasks
- Automated memory management system
- Structured project context (PROJECT_CONTEXT.md)
- Decision recording system (DECISIONS.md)

### 💡 Core Commands
1. `/start` - Quick project understanding
2. `/context` - Context sync checkpoint
3. `/sync` - State synchronizer
4. `/plan` - Task planning & design
5. `/check` - Intelligent code review
6. `/test` - Test generation & execution
7. `/learn` - Learn and record decisions
8. `/doc` - Intelligent documentation maintenance
9. `/review` - PR preparation assistant
10. `/debug` - Intelligent debugging assistant
11. `/meta` - Project specification customization

### ⚡ Improvements
- Clearer command responsibilities, no functional overlap
- Most commands require no parameters, intelligent inference
- Better context retention capability
- Simplified learning curve
- Improved human-AI collaboration efficiency

### 🔄 Integrated Commands
- `analyze`, `discover`, `explore` → `/start`
- `audit`, `coverage`, `perf` → `/check`
- `doc-api`, `doc-arch`, `changelog`, `readme` → `/doc`
- `deploy-check`, `rollback`, `config` → `/review`
- Other specialized commands integrated into relevant core commands

### 📝 Documentation Updates
- Added SIMPLE_COMMANDS.md detailed description
- Updated COMMANDS_SUMMARY.md to v3.0
- Rewrote DEPLOY_GUIDE.md for new system
- Created PROJECT_CONTEXT.md template

---

## v2.1.0 (2024-01-14)

### New Features
- Parameter standardization: All commands have clear format and examples
- Command coordination mechanism: Support state sharing and intelligent recommendations
- Project command enhancement: Upgraded from simple functions to complete systems

### Optimizations
- Resolved functional overlap issues
- Clarified command responsibility boundaries
- Optimized command parameter definitions

---

## v2.0.0 (2024-01-13)

### Initial Version
- 31 commands covering complete development lifecycle
- Including global and project commands
- Auto-deployment scripts supporting Windows/macOS/Linux
