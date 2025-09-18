# Contributing to Claude Code Framework

*[中文版](#中文版-贡献指南) | English*

Thank you for your interest in contributing to the Claude Code Framework! This project thrives on community collaboration and we welcome contributions from developers of all skill levels.

## 🌟 Ways to Contribute

### 1. **Agent Development**
- Create new specialized agents for emerging technologies
- Improve existing agent capabilities and knowledge
- Develop context detectors for new programming languages

### 2. **Command Enhancement**
- Optimize existing commands for better performance
- Create project-specific command templates
- Improve command documentation and examples

### 3. **Documentation**
- Translate documentation to new languages
- Create tutorials and guides
- Improve code examples and use cases

### 4. **Bug Reports & Feature Requests**
- Report bugs with detailed reproduction steps
- Suggest new features and improvements
- Provide feedback on user experience

### 5. **Testing & Quality Assurance**
- Write tests for agents and commands
- Test compatibility across different environments
- Validate documentation accuracy

## 🚀 Getting Started

### Prerequisites
- Git for version control
- Basic understanding of Claude Code framework
- Familiarity with Markdown for documentation

### Setup Development Environment
```bash
# 1. Fork the repository on GitHub
# 2. Clone your fork
git clone https://github.com/YOUR_USERNAME/claude-code-framework.git
cd claude-code-framework

# 3. Create a feature branch
git checkout -b feature/your-feature-name

# 4. Make your changes
# 5. Test your changes thoroughly
# 6. Commit with clear messages
git commit -m "feat: add new agent for framework X"

# 7. Push to your fork
git push origin feature/your-feature-name

# 8. Create a Pull Request
```

## 📋 Contribution Guidelines

### Code Standards

#### Agent Development
```markdown
---
name: framework-expert
model: sonnet  # or haiku for simple agents
description: "Brief description of agent capabilities"
trigger: "*.ext, specific patterns"
tools: Read, Write, Edit, Bash, Grep
---

# Agent Name - Brief Description

I am an expert in [Technology/Framework]...

## Key Capabilities
- Capability 1
- Capability 2
- Capability 3

## Usage Examples
[Provide clear examples]

## Best Practices
[Include relevant best practices]
```

#### Command Development
```markdown
---
arguments: required|optional
format: "[parameter format]"
examples:
  - "/command - basic usage"
  - "/command param - with parameter"
---

# Command Title

Brief description of what this command does.

## 🎯 Purpose
Detailed explanation...

## 📋 Usage
Step-by-step usage instructions...

## 💡 Examples
Practical examples with expected outputs...
```

#### Context Detector Pattern
```python
class FrameworkContextDetector:
    def analyze_project(self, file_path: str) -> ProjectContext:
        detections = [
            self.detect_by_imports(file_path),
            self.detect_by_file_structure(file_path),
            self.detect_by_config_files(file_path)
        ]
        
        primary_context = max(detections, key=lambda d: d.confidence)
        
        return ProjectContext(
            primary_type=primary_context.type,
            confidence=primary_context.confidence,
            agent_recommendation=self.get_recommended_agent()
        )
```

### Documentation Standards

#### Bilingual Documentation
- **Primary Language**: English for global accessibility
- **Secondary Language**: Chinese for local community
- **File Naming**: `filename.en.md` for English, `filename.md` for Chinese
- **Cross-References**: Always include language version links

#### Documentation Structure
```markdown
# Title

*[中文版](filename.md) | English*

## Overview
Brief introduction...

## Detailed Sections
...

## Examples
Practical examples with code...

## Related Resources
- Links to related documentation
- External resources
```

### Quality Standards

#### Testing Requirements
- **Agent Testing**: Verify agent triggers and responses
- **Command Testing**: Test all command variations
- **Integration Testing**: Ensure compatibility with existing system
- **Documentation Testing**: Verify all links and examples work

#### Review Criteria
- **Functionality**: Does it work as intended?
- **Performance**: Is it efficient and responsive?
- **Usability**: Is it easy to understand and use?
- **Compatibility**: Works with existing framework?
- **Documentation**: Is it well-documented?

## 🎯 Contribution Areas

### High Priority
1. **New Language Support**: Add context detectors for languages like:
   - PHP (Laravel, Symfony)
   - Ruby (Rails, Sinatra)
   - Swift (iOS, macOS)
   - Dart (Flutter)

2. **Cloud Platform Agents**: Specialized agents for:
   - AWS services
   - Google Cloud Platform
   - Azure services
   - Docker/Kubernetes

3. **Framework Specialists**: New agents for:
   - SvelteKit
   - Remix
   - Astro
   - Fresh (Deno)

### Medium Priority
1. **Developer Tools**: Agents for:
   - CI/CD pipelines
   - Monitoring and logging
   - Database optimization
   - Security scanning

2. **Quality Improvements**:
   - Better error handling
   - Performance optimizations
   - Enhanced user feedback

### Ongoing Needs
1. **Documentation Translation**: Help translate to more languages
2. **Tutorial Creation**: Step-by-step guides for complex workflows
3. **Community Examples**: Real-world usage examples
4. **Bug Fixes**: Address reported issues

## 📝 Pull Request Process

### Before Submitting
1. **Test Thoroughly**: Verify your changes work as expected
2. **Update Documentation**: Include relevant documentation updates
3. **Follow Standards**: Adhere to coding and documentation standards
4. **Check Compatibility**: Ensure backward compatibility

### PR Template
```markdown
## Description
Brief description of changes...

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Documentation update
- [ ] Performance improvement
- [ ] Breaking change

## Testing
- [ ] Local testing completed
- [ ] Documentation updated
- [ ] Examples provided
- [ ] Backward compatibility verified

## Screenshots/Examples
[If applicable]

## Checklist
- [ ] Code follows project standards
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] Tests added/updated
```

### Review Process
1. **Automated Checks**: CI/CD pipeline validation
2. **Code Review**: Maintainer review for quality and standards
3. **Testing**: Community testing when applicable
4. **Documentation Review**: Verify documentation accuracy
5. **Final Approval**: Maintainer approval and merge

## 🏆 Recognition

### Contributor Recognition
- **Hall of Fame**: Featured contributors in documentation
- **GitHub Profile**: Contribution statistics and achievements
- **Community Shoutouts**: Recognition in release notes
- **Beta Access**: Early access to new features

### Becoming a Maintainer
Active contributors may be invited to become maintainers with:
- **Review Privileges**: Help review community contributions
- **Feature Planning**: Input on project direction
- **Release Management**: Help with version releases
- **Community Support**: Guide new contributors

## 💬 Community & Support

### Communication Channels
- **GitHub Issues**: Bug reports and feature requests
- **GitHub Discussions**: Community discussions and Q&A
- **Pull Requests**: Code contributions and reviews

### Getting Help
1. **Check Documentation**: Comprehensive guides available
2. **Search Issues**: Your question might already be answered
3. **Ask the Community**: Use GitHub Discussions
4. **Direct Contact**: Maintainers for complex questions

### Code of Conduct
We follow a simple code of conduct:
- **Be Respectful**: Treat all community members with respect
- **Be Constructive**: Provide helpful, actionable feedback
- **Be Inclusive**: Welcome developers of all backgrounds and skill levels
- **Be Patient**: Remember that everyone is learning

## 🎉 Thank You!

Every contribution, no matter how small, helps make Claude Code Framework better for everyone. Whether you're fixing a typo, adding a new agent, or helping with translations - your efforts are valued and appreciated!

---

# 中文版 - 贡献指南

感谢您对 Claude Code Framework 的贡献兴趣！本项目依靠社区协作蓬勃发展，我们欢迎各个技能水平的开发者参与贡献。

## 🌟 贡献方式

### 1. **代理开发**
- 为新兴技术创建专业代理
- 改进现有代理的能力和知识
- 为新编程语言开发上下文检测器

### 2. **命令增强**
- 优化现有命令以提高性能
- 创建项目特定的命令模板
- 改进命令文档和示例

### 3. **文档完善**
- 将文档翻译为新语言
- 创建教程和指南
- 改进代码示例和用例

### 4. **错误报告和功能请求**
- 报告错误并提供详细的重现步骤
- 建议新功能和改进
- 提供用户体验反馈

### 5. **测试和质量保证**
- 为代理和命令编写测试
- 测试不同环境的兼容性
- 验证文档准确性

## 📝 贡献流程

详细的贡献步骤请参考上方英文版内容，或查看项目的 GitHub 仓库了解最新的贡献指南。

## 💬 社区支持

- **GitHub Issues**: 错误报告和功能请求
- **GitHub Discussions**: 社区讨论和问答
- **Pull Requests**: 代码贡献和审查

感谢您的每一个贡献！无论是修复错字、添加新代理还是帮助翻译，您的努力都很宝贵！

---

**Happy Contributing! 🚀**