# Agent Capability Matrix

## 📊 Agent System Overview (After Optimization)
- **Total Agents**: 35 (down from 45)
- **Categories**: 9 (ai, architecture, core, frameworks, languages, quality, roles, workflow)
- **Last Updated**: 2025-09-02

## 🎯 Capability Matrix

### Programming Languages

| Language | Primary Agent | Use Case | Token Budget | Trigger Keywords |
|----------|--------------|----------|--------------|------------------|
| **TypeScript** | typescript-expert | All TypeScript development | ~150k | typescript, ts, tsx, type-safe |
| **Python** | python-fullstack-expert | General Python development | ~150k | python, py, django, flask |
| **Python ML** | python-ml-specialist | Machine Learning/AI | ~170k | ml, ai, tensorflow, pytorch, scikit |
| **Kotlin Android** | android-kotlin-architect | Android native apps | ~150k | android, kotlin, compose, mobile |
| **Kotlin Backend** | kotlin-backend-expert | Server-side Kotlin | ~150k | ktor, spring, kotlin + server |
| **Java** | java-enterprise-architect | Enterprise Java | ~160k | java, spring boot, enterprise |
| **Go** | golang-systems-engineer | Systems programming | ~140k | go, golang, microservices |
| **Rust** | rust-zero-cost | Systems programming | ~130k | rust, cargo, memory-safe |
| **C++** | cpp-modern-master | Modern C++ | ~120k | c++, cpp, stl |
| **C#** | csharp-dotnet-master | .NET development | ~160k | c#, csharp, dotnet, asp.net |
| **C** | c-systems-architect | Low-level systems | ~140k | c, embedded, kernel |

### Frontend & UI Development

| Domain | Agent | Responsibility | Token Budget | NOT Responsible For |
|--------|-------|---------------|--------------|---------------------|
| **Web Frontend** | frontend-developer | React, Vue, Angular, Web | ~150k | Native mobile (except React Native) |
| **React Native** | frontend-developer | React Native & Expo | ~150k | Native iOS/Android |
| **iOS Native** | mobile-developer | Swift, SwiftUI, UIKit | ~170k | Android, React Native |
| **Flutter** | mobile-developer | Flutter/Dart cross-platform | ~170k | React Native, Native Android |
| **Android Native** | android-kotlin-architect | Kotlin, Compose | ~150k | Flutter, React Native |

### Backend & Architecture

| Domain | Agent | Focus Area | Token Budget | Key Technologies |
|--------|-------|------------|--------------|------------------|
| **API Design** | api-architect | REST, GraphQL, gRPC | ~120k | OpenAPI, Swagger |
| **Full-Stack** | fullstack-architect | End-to-end architecture | ~200k | Complete systems |
| **Microservices** | golang-systems-engineer | Go microservices | ~140k | Docker, K8s |
| **Backend Kotlin** | kotlin-backend-expert | Ktor, Spring Boot | ~150k | Server-side Kotlin |

### AI/ML & Data Science

| Specialization | Agent | Use Case | Token Budget | Frameworks |
|---------------|-------|----------|--------------|------------|
| **ML Development** | python-ml-specialist | Model development | ~170k | TensorFlow, PyTorch |
| **LLM/RAG** | llm-engineer | LLM applications | ~190k | LangChain, OpenAI |
| **MLOps** | mlops-specialist | ML deployment | ~200k | Kubeflow, MLflow |

### Quality & Testing

| Function | Agent | Responsibility | Token Budget | Approach |
|----------|-------|---------------|--------------|----------|
| **Code Review** | code-reviewer | Basic review | ~80k | Single pass |
| **Validation** | jenny-validator | Correctness check | ~120k | Detailed validation |
| **Reality Check** | karen-realist | Feasibility assessment | ~120k | Practical evaluation |
| **Senior Review** | senior-developer | Best practices | ~120k | Architecture review |
| **Testing** | test-automator | Test generation | ~100k | Unit, integration tests |
| **Bug Fixing** | bug-hunter | Debug & fix | ~110k | Root cause analysis |
| **Performance** | performance-optimizer | Optimization | ~100k | Profiling, tuning |

### Specialized Roles

| Role | Agent | Domain | Token Budget | Key Skills |
|------|-------|--------|--------------|------------|
| **Product** | product-manager | Requirements, roadmap | ~90k | PRD, user stories |
| **UX Design** | ux-designer | User experience | ~100k | Wireframes, flows |
| **DevOps** | devops-engineer | Infrastructure | ~110k | CI/CD, containers |
| **Security** | security-analyst | Security audit | ~100k | Threat modeling |
| **Documentation** | technical-writer | Technical docs | ~90k | API docs, guides |
| **SDK Design** | sdk-product-owner | Developer tools | ~120k | DX, API design |

## 🚫 Removed Agents (No Longer Available)

| Removed Agent | Reason | Replacement |
|--------------|--------|-------------|
| kotlin-expert | Too simplified | kotlin-backend-expert or android-kotlin-architect |
| kotlin-polyglot-master | Too complex, overlapping | Specialized Kotlin agents |
| typescript-expert-core | Fragmented | typescript-expert (unified) |
| typescript-expert-examples | Fragmented | typescript-expert (unified) |
| typescript-fullstack-expert | Fragmented | typescript-expert (unified) |
| *-context-detector (5 agents) | Redundant | Logic embedded in orchestrator |
| token-efficient-loader | Not implementable | N/A |
| work-coordinator | Duplicates orchestrator | orchestrator handles coordination |

## 📈 Selection Guidelines

### Single Agent Preferred
- Simple tasks with clear technology stack
- Performance optimization (always single expert)
- Bug fixing (focused debugging)
- Language-specific tasks

### Multiple Agents Acceptable
- Full-stack projects (max 2 agents)
- Security-critical APIs (api-architect + security-analyst)
- Code quality review (trilogy: jenny + karen + senior)

### Conflict Resolution
1. **React Native**: Always → frontend-developer (NOT mobile-developer)
2. **Kotlin Context**: Check if Android → android-kotlin-architect, if Backend → kotlin-backend-expert
3. **TypeScript**: Unified typescript-expert for all TS needs
4. **ML vs MLOps**: Development → python-ml-specialist, Deployment → mlops-specialist

## 🎯 Token Budget Guidelines

| Task Complexity | Target Budget | Max Agents |
|----------------|---------------|------------|
| Simple | < 150k | 1 |
| Medium | 150-300k | 1-2 |
| Complex | 300-400k | 2-3 |
| Review Only | ~360k | 3 (quality trilogy) |

## 📊 Usage Metrics Template

```yaml
agent_metrics:
  agent_name:
    total_calls: 0
    avg_tokens: 0
    success_rate: 0%
    common_tasks: []
    issues_found: []
```

---

*This matrix should be updated whenever agents are added, removed, or modified.*