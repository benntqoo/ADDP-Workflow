# Config Directory

## ⚠️ This directory is deprecated

All configuration files have been removed because Claude Code doesn't support custom configuration parsing.

## Removed Files (2025-08-16)

The following files have been removed because they don't work with Claude Code:

### ❌ triggers.yaml (REMOVED)
- **Why removed**: Claude Code doesn't parse custom YAML trigger configurations
- **What it claimed to do**: Auto-trigger agents based on file types
- **Reality**: This never actually worked

### ❌ workflows.yaml (REMOVED)  
- **Why removed**: Workflow orchestration is handled internally by Claude Code
- **What it claimed to do**: Define multi-agent workflows
- **Reality**: Claude Code manages agent coordination automatically

### ❌ token-settings.yaml (REMOVED)
- **Why removed**: Claude Code doesn't parse custom token optimization settings
- **What it claimed to do**: Optimize token usage based on user tier and complexity
- **Reality**: Claude Code manages token usage internally

## How Claude Code Actually Works

Claude Code uses **subagents** which are triggered by:
1. Task description matching agent's `description` field
2. Explicit invocation: "Use the code-reviewer agent"
3. NOT by file types or custom configurations

## Migration Note

If you had custom configurations, focus on:
- Writing clear agent `description` fields with trigger keywords
- Using explicit agent invocation when needed
- Letting Claude Code handle orchestration automatically