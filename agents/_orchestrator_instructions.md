# Intelligent Agent Selection Instructions

## Overview
This document provides Claude Code's orchestrator with intelligent agent selection logic to prevent token waste and improve precision.

## Selection Process

### Step 1: Load Selection Rules
Always read `agents/_selection_engine.json` before making agent selection decisions.

### Step 2: Analyze User Request
Extract key information from the user request:
1. **Primary Domain**: Frontend, Backend, AI/ML, Infrastructure, etc.
2. **Task Type**: Implement, Optimize, Review, Debug, Design
3. **Technology Stack**: Programming languages, frameworks mentioned
4. **Complexity Indicators**: Simple single-task vs complex multi-component
5. **Performance Sensitivity**: Are tokens a constraint?

### Step 3: Apply Selection Logic

#### For Frontend Tasks:
```
IF request contains ["react", "vue", "css", "ui", "component"]
THEN use: frontend-developer (single agent)
AVOID: typescript-fullstack-expert (capability overlap)
```

#### For Performance Optimization:
```  
IF request contains ["performance", "optimize", "slow", "bottleneck"]
THEN use: performance-optimizer (single agent, high efficiency)
UNLESS: specific language expertise needed (check language_specific rules)
```

#### For AI/ML Tasks:
```
IF request contains ["ml", "ai", "model"]:
  IF contains ["deploy", "production", "serving"]
  THEN use: mlops-specialist
  
  ELIF contains ["rag", "llm", "prompt"]  
  THEN use: llm-engineer
  
  ELSE use: python-ml-specialist
```

#### For Quality Review:
```
IF request contains ["review", "check", "quality", "audit"]
THEN use: quality trilogy [jenny-validator, karen-realist, senior-developer]
NOTE: Requires coordination mechanism
```

### Step 4: Conflict Avoidance
Before finalizing agent selection, check for:
1. **Capability Overlaps**: Don't select multiple agents with similar expertise
2. **Token Budget**: Estimate total tokens needed (see token_usage_estimates)
3. **Coordination Complexity**: Avoid parallel agents that need tight coordination

### Step 5: Optimize for Efficiency

#### High-Efficiency Patterns:
- **Single Expert**: For focused tasks, prefer 1 highly relevant agent
- **Parallel Only When Beneficial**: For truly independent components
- **Sequential for Dependencies**: When agents need each other's output

#### Token-Saving Rules:
- Maximum 2 agents for most tasks
- Use trilogy (3 agents) only for quality review
- Prefer specialized agent over general-purpose

## Decision Templates

### Template: Simple Task
```
User Request: "Fix the React component styling"
Analysis: Frontend + Simple + Single technology
Decision: frontend-developer (1 agent, ~150k tokens)
Reasoning: Focused expertise, no overlaps needed
```

### Template: Complex Task
```
User Request: "Build a complete e-commerce platform"
Analysis: Fullstack + Complex + Multiple components
Decision: fullstack-architect + frontend-developer (2 agents, ~350k tokens)
Reasoning: Architecture planning + UI implementation
```

### Template: Optimization Task
```
User Request: "Optimize our Python API performance" 
Analysis: Performance + Backend + Python-specific
Decision: performance-optimizer (1 agent, ~100k tokens)
Reasoning: Specialized performance expertise, language-agnostic
```

## Error Recovery

### If Agent Selection Goes Wrong:
1. **Too Many Agents**: Stop and restart with stricter rules
2. **Wrong Expertise**: Check selection_engine.json for better match
3. **High Token Usage**: Switch to more efficient single-agent approach

### Validation Checklist:
- [ ] Selected agents have clear, non-overlapping roles
- [ ] Total estimated tokens < 400k for most tasks
- [ ] Each agent adds unique value
- [ ] Dependencies between agents are minimal

## Learning and Adaptation

### Track Success Patterns:
- Which agent selections led to best user satisfaction?
- What token usage patterns are most efficient?
- Which combinations create conflicts?

### Update Selection Rules:
Periodically review and update `_selection_engine.json` based on:
- User feedback patterns
- Token efficiency data  
- Agent performance metrics

## Emergency Fallbacks

### If Selection Engine Fails:
1. **Default Safe Choice**: production-ready-coder (general-purpose)
2. **Simple Heuristic**: Match primary keyword to agent name
3. **Ask User**: Present 2-3 options for user selection

Remember: **Precision over power**. Better to use 1 perfect agent than 3 good ones.