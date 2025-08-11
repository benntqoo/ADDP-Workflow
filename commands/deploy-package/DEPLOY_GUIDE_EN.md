# Claude Code Streamlined Command System Deployment Guide v3.0

## 🚀 Quick Deploy

### Windows
```powershell
cd claude\commands\deploy-package
.\deploy.ps1
```

### macOS/Linux
```bash
cd claude/commands/deploy-package
chmod +x deploy.sh
./deploy.sh
```

## 📦 Package Contents

### Universal Commands (14)
1. **start** - Quick project understanding
2. **context** - Context sync checkpoint
3. **sync** - State synchronizer
4. **plan** - Task planning & design
5. **check** - Intelligent code review
6. **test** - Test generation & execution
7. **learn** - Learn and record decisions
8. **doc** - Intelligent documentation maintenance
9. **review** - PR preparation assistant
10. **debug** - Intelligent debugging assistant
11. **meta** - Project specification customization
12. **analyze** - Deep risk analysis
13. **update-spec** - Update project specifications
14. **watch** - Guardian mode

### Deployment Location
- **Windows**: `%USERPROFILE%\.claude\commands\`
- **macOS/Linux**: `~/.claude/commands/`

## 🔧 Manual Deployment

If automatic scripts fail, you can manually copy:

```bash
# 1. Create target directory
mkdir -p ~/.claude/commands

# 2. Copy global commands
cp global/*.md ~/.claude/commands/

# 3. Create project-specific commands (developer customization)
# Developers can create their own project commands in .claude/commands/
mkdir -p YOUR_PROJECT/.claude/commands
# Refer to global command format to create custom commands
```

## 📋 Post-Deployment Verification

1. Type `/` in Claude Code to see available commands
2. Test core commands:
   ```bash
   /start    # Should start analyzing project
   /context  # Should display current understanding
   ```

## ⚙️ Custom Configuration

### Selective Deployment
If you only need certain commands, copy only the required `.md` files.

### Project-Specific Commands
Create custom commands in your project's `.claude/commands/` directory.

### Command Priority
1. Project commands (`.claude/commands/`)
2. Global commands (`~/.claude/commands/`)
3. Built-in commands

## 🔄 Update Instructions

Upgrading from old version (v2.x):
1. Backup existing custom commands
2. Delete old global commands:
   ```bash
   rm ~/.claude/commands/*.md
   ```
3. Run new deployment script
4. Restore custom commands (if any)

## 💡 Usage Suggestions

### New Project Startup
```bash
/meta      # Establish project specifications
/start     # Understand project structure
/plan      # Plan first feature
```

### Daily Development Flow
```bash
/sync      # Restore work state
/context   # Confirm understanding
/plan      # Plan new task
/check     # Code quality check
/test      # Execute tests
/learn     # Record important decisions
```

### Code Submission
```bash
/check     # Final check
/doc       # Update documentation
/review    # Prepare PR
```

## ❓ FAQ

### Q: Commands not showing?
A: 
- Ensure file extension is `.md`
- Check if files are in correct directory
- Restart Claude Code

### Q: Command conflicts?
A: Project commands override global commands with the same name.

### Q: How to completely uninstall?
A: Delete corresponding files in `~/.claude/commands/`.

### Q: Difference from old commands?
A: v3.0 consolidates 31 commands into 14 universal commands, smarter functionality, simpler usage.

## 📝 Version Notes

### v3.3 (2025-08-10)
- **Streamlined**: Reduced from 31 to 14 universal commands
- **Focused**: Removed project-level commands, let developers define their own
- **Intelligent**: Each command integrates multiple related functions
- **Automated**: Memory management and state sync automation
- **Simplified**: Most commands require no parameters

### Major Improvements
- Command count reduced by 55%
- Feature coverage maintained at 100%
- Learning cost significantly reduced
- Human-AI collaboration efficiency improved

### Removed Commands
The following functions have been integrated into new commands:
- analyze, discover, explore → `/start`
- audit, coverage, perf → `/check`
- doc-api, doc-arch, changelog, readme → `/doc`
- deploy-check, rollback, config → `/review`
- Other specialized commands integrated into relevant core commands

---

*Simple and efficient, let Claude Code become your best development partner!*

*Deployment Package Version: 3.0.0*  
*Release Date: 2024-01-15*
