---
name: token-efficient-loader
description: Strategy for efficient token usage in agent loading
model: haiku
tools: [read]
---

# Token-Efficient Agent Loading Strategy

## Loading Patterns for Different User Tiers

### 1. Dynamic Loading Based on Context

```yaml
# For Token-Conscious Users (Free/Basic Plans)
typescript_tasks:
  simple_function:
    load: typescript-expert-core  # ~400 tokens
  
  complex_component:
    load: 
      - typescript-expert-core     # Base (~400 tokens)
      - typescript-expert-examples  # Only if needed (+600 tokens)

# For Power Users (Max Plan)  
typescript_tasks:
  any:
    load: typescript-fullstack-expert  # Full ~2500 tokens
```

### 2. Conditional Loading Rules

```typescript
interface AgentLoadingStrategy {
  // Detect user's token budget
  getUserTier(): 'free' | 'basic' | 'pro' | 'max'
  
  // Detect task complexity
  analyzeTaskComplexity(prompt: string): 'simple' | 'medium' | 'complex'
  
  // Smart loading decision
  selectAgent(tier: string, complexity: string): string[] {
    if (tier === 'max') return ['full-expert']
    
    if (complexity === 'simple') return ['core']
    if (complexity === 'medium') return ['core', 'patterns']
    if (complexity === 'complex') return ['core', 'patterns', 'examples']
  }
}
```

### 3. Progressive Enhancement

```markdown
## Level 1: Core (400 tokens)
- Essential rules and patterns
- Basic error handling
- Type safety fundamentals

## Level 2: + Patterns (800 tokens total)
- Common design patterns
- Framework basics
- Testing essentials

## Level 3: + Examples (1200 tokens total)
- Detailed code examples
- Edge cases
- Advanced patterns

## Level 4: Full Expert (2500 tokens)
- Everything above
- Integration examples
- Performance optimizations
- Security best practices
```

### 4. Token Usage Comparison

| User Scenario | Old Approach | Optimized Approach | Savings |
|--------------|--------------|-------------------|---------|
| Simple function | 2500 tokens | 400 tokens (core only) | 84% |
| React component | 2500 tokens | 800 tokens (core+patterns) | 68% |
| Complex system | 2500 tokens | 2500 tokens (full) | 0% |
| Average usage | 2500 tokens | ~1000 tokens | 60% |

### 5. Implementation in CLAUDE.md

```markdown
# Token-Efficient Mode Configuration

## Enable Token Optimization
Set in your project's CLAUDE.md:

\`\`\`
TOKEN_MODE: efficient  # Options: efficient | balanced | full
\`\`\`

## Per-Language Settings
\`\`\`
typescript:
  mode: efficient
  load_examples: on_demand
  
python:
  mode: balanced
  
rust:
  mode: full  # Always load full for systems programming
\`\`\`
```

### 6. Automatic Detection Rules

```typescript
// Auto-detect when to load more context
const needsExamples = (prompt: string): boolean => {
  const indicators = [
    'production', 'enterprise', 'scalable',
    'best practice', 'real world', 'complex',
    'advanced', 'optimize', 'performance'
  ]
  return indicators.some(word => 
    prompt.toLowerCase().includes(word)
  )
}

const needsFullExpert = (context: FileContext): boolean => {
  return (
    context.fileCount > 10 ||
    context.linesOfCode > 1000 ||
    context.hasTests ||
    context.hasCI ||
    context.isMonorepo
  )
}
```