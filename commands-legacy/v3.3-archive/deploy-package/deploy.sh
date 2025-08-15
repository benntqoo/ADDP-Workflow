#!/bin/bash

# Claude Code 命令自动部署脚本
# 支持 macOS 和 Linux

echo "🚀 Claude Code 命令部署脚本"
echo "=========================="

# 检测操作系统
OS=$(uname -s)
if [[ "$OS" == "Darwin" ]]; then
    echo "✓ 检测到 macOS 系统"
elif [[ "$OS" == "Linux" ]]; then
    echo "✓ 检测到 Linux 系统"
else
    echo "❌ 不支持的操作系统: $OS"
    exit 1
fi

# 设置命令目录路径
GLOBAL_CMD_DIR="$HOME/.claude/commands"
PROJECT_CMD_DIR="./.claude/commands"

# 创建全局命令目录
echo ""
echo "📁 创建全局命令目录..."
mkdir -p "$GLOBAL_CMD_DIR"
if [[ $? -eq 0 ]]; then
    echo "✓ 全局命令目录创建成功: $GLOBAL_CMD_DIR"
else
    echo "❌ 创建全局命令目录失败"
    exit 1
fi

# 复制全局命令
echo ""
echo "📋 复制全局命令..."
cp ./global/*.md "$GLOBAL_CMD_DIR/"
if [[ $? -eq 0 ]]; then
    echo "✓ 成功复制 $(ls -1 ./global/*.md | wc -l) 个全局命令"
else
    echo "❌ 复制全局命令失败"
    exit 1
fi

# 询问是否部署项目命令
echo ""
read -p "是否要部署项目命令到当前目录? (y/n) " -n 1 -r
echo ""
if [[ $REPLY =~ ^[Yy]$ ]]; then
    # 创建项目命令目录
    echo "📁 创建项目命令目录..."
    mkdir -p "$PROJECT_CMD_DIR"
    if [[ $? -eq 0 ]]; then
        echo "✓ 项目命令目录创建成功: $PROJECT_CMD_DIR"
    else
        echo "❌ 创建项目命令目录失败"
        exit 1
    fi
    
    # 复制项目命令
    echo "📋 复制项目命令..."
    cp ./project/*.md "$PROJECT_CMD_DIR/"
    if [[ $? -eq 0 ]]; then
        echo "✓ 成功复制 $(ls -1 ./project/*.md | wc -l) 个项目命令"
    else
        echo "❌ 复制项目命令失败"
        exit 1
    fi
fi

# 验证部署
echo ""
echo "🔍 验证部署结果..."
echo "全局命令数量: $(ls -1 "$GLOBAL_CMD_DIR"/*.md 2>/dev/null | wc -l)"
if [[ -d "$PROJECT_CMD_DIR" ]]; then
    echo "项目命令数量: $(ls -1 "$PROJECT_CMD_DIR"/*.md 2>/dev/null | wc -l)"
fi

echo ""
echo "✨ 部署完成！"
echo ""
echo "提示："
echo "1. 请重启 Claude Code 以加载新命令"
echo "2. 使用 /meta 命令开始定制项目规范"
echo "3. 查看 DEPLOY_GUIDE.md 了解更多信息"
echo ""