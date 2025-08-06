# Claude 宪法同步命令集

这些命令帮助您在项目间同步和更新 Claude 协作规范。

## 📋 核心同步命令

### 1. `/sync-constitution` - 同步最新宪法
**文件**: `~/.claude/commands/sync-constitution.md`
```markdown
---
command: |
  # 检查本地是否有 CLAUDE.md
  if [ -f "CLAUDE.md" ]; then
    echo "发现现有 CLAUDE.md，准备更新..."
  else
    echo "未发现 CLAUDE.md，准备创建..."
  fi
---
请执行以下操作：

1. **获取最新宪法模板**
   - 从 ~/claude-constitution/CLAUDE_CONSTITUTION.md 读取最新版本
   - 或从指定的 Git 仓库获取

2. **保留本地定制内容**
   - 如果存在 CLAUDE.md，提取 E 部分（项目定制内容）
   - 保存其他本地修改的标记

3. **合并更新**
   - 使用最新的宪法框架
   - 保留本地的 E 部分内容
   - 标记版本和更新时间

4. **生成报告**
   - 列出更新的内容
   - 提示需要手动处理的冲突
   - 建议下一步操作
```

### 2. `/init-claude` - 初始化项目宪法
**文件**: `~/.claude/commands/init-claude.md`
```markdown
---
command: |
  # 检查项目类型
  if [ -f "package.json" ]; then
    echo "检测到 Node.js 项目"
  elif [ -f "go.mod" ]; then
    echo "检测到 Go 项目"
  elif [ -f "requirements.txt" ]; then
    echo "检测到 Python 项目"
  fi
---
为当前项目初始化 Claude 协作规范：

1. **检查现有配置**
   - 是否已有 CLAUDE.md
   - 是否有 .claude/ 目录

2. **复制宪法模板**
   - 从标准位置复制 CLAUDE_CONSTITUTION.md
   - 重命名为 CLAUDE.md

3. **启动元工作流**
   - 自动开始项目定制流程
   - 引导用户完成 E 部分配置

4. **创建项目命令**
   - 创建 .claude/commands/ 目录
   - 添加项目特定的命令
```

### 3. `/update-constitution` - 更新宪法框架
**文件**: `~/.claude/commands/update-constitution.md`
```markdown
更新项目的 Claude 宪法到最新版本：

1. **备份当前版本**
   - 保存当前 CLAUDE.md 为 CLAUDE.md.backup
   - 记录备份时间

2. **智能合并**
   - 保留 E 部分（项目定制）
   - 保留自定义的 H 部分（如果有）
   - 更新其他所有部分到最新版

3. **差异分析**
   - 显示主要更新内容
   - 标记不兼容的更改
   - 提供迁移建议

4. **验证结果**
   - 检查合并后的文件完整性
   - 确保所有部分都存在
   - 提示用户检查和确认
```

## 🔧 高级同步命令

### 4. `/export-config` - 导出项目配置
**文件**: `~/.claude/commands/export-config.md`
```markdown
导出当前项目的 Claude 配置：

1. 提取 CLAUDE.md 的 E 部分
2. 收集 .claude/commands/ 下的项目命令
3. 生成可分享的配置包
4. 保存到 claude-config-export.json

用途：团队共享、项目模板、配置备份
```

### 5. `/import-config` - 导入项目配置
**文件**: `~/.claude/commands/import-config.md`
```markdown
---
arguments: optional
---
从配置包导入 Claude 设置：

1. 读取 ${ARGUMENTS:-claude-config-export.json}
2. 验证配置格式和版本兼容性
3. 合并到当前项目
4. 处理冲突和覆盖策略

选项：
- 完全覆盖
- 智能合并
- 仅添加新内容
```

## 📦 批量管理命令

### 6. `/sync-all-projects` - 批量同步
**文件**: `~/.claude/commands/sync-all-projects.md`
```markdown
---
command: |
  find ~/projects -name "CLAUDE.md" -type f | head -5
---
更新所有项目的 Claude 宪法：

1. 扫描包含 CLAUDE.md 的项目
2. 为每个项目执行更新
3. 生成批量更新报告
4. 标记需要手动处理的项目

安全模式：
- 先预览将要更新的项目
- 确认后再执行
- 保留所有备份
```

## 🎯 使用场景

### 场景1：团队统一更新
```bash
# 团队负责人更新了宪法模板
/update-constitution  # 在一个项目中测试
/sync-all-projects   # 批量更新所有项目
```

### 场景2：新成员加入
```bash
# 新成员克隆项目后
/init-claude         # 初始化配置
/import-config team-config.json  # 导入团队配置
```

### 场景3：最佳实践分享
```bash
# 在成功的项目中
/export-config       # 导出配置
# 分享给其他项目使用
```

## 💡 实现建议

### 1. 版本管理
```yaml
# 在 CLAUDE.md 中添加元数据
---
version: 1.0.0
updated: 2024-01-15
base: CLAUDE_CONSTITUTION.md@v1.0
---
```

### 2. 配置中心
```bash
# 设置环境变量指向宪法仓库
export CLAUDE_CONSTITUTION_REPO="https://github.com/org/claude-constitution"

# 或本地路径
export CLAUDE_CONSTITUTION_PATH="~/claude-constitution"
```

### 3. Git 集成
```bash
# 自动提交更新
git add CLAUDE.md
git commit -m "chore: update Claude constitution to v1.0.1"
```

## 🚀 快速设置

```bash
# 创建同步命令
mkdir -p ~/.claude/commands

# 基础同步命令
cat > ~/.claude/commands/sync-constitution.md << 'EOF'
同步最新的 Claude 宪法到当前项目，保留本地定制内容。
EOF

# 初始化命令
cat > ~/.claude/commands/init-claude.md << 'EOF'
为新项目初始化 Claude 协作规范，并启动定制流程。
EOF
```

---

*通过这些同步命令，您可以轻松维护和更新所有项目的 Claude 协作规范，确保团队始终使用最新最佳实践。*