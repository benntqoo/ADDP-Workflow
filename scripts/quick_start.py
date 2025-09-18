#!/usr/bin/env python3
"""
Universal AI Coding Framework - 快速启动脚本
===========================================

这个脚本帮助用户快速部署和配置整个 Universal AI Coding Framework。

功能:
1. 环境检查 (Python, Ollama)
2. 依赖安装
3. 项目初始化
4. 配置生成
5. 服务启动
"""

import asyncio
import subprocess
import sys
import shutil
import json
from pathlib import Path
from typing import Dict, List, Tuple

class QuickStart:
    """快速启动助手"""

    def __init__(self):
        self.project_root = Path(__file__).parent.parent
        self.requirements_file = self.project_root / "requirements.txt"
        self.main_script = self.project_root / "main.py"

    async def run(self):
        """运行快速启动流程"""
        print("🚀 Universal AI Coding Framework 快速启动")
        print("=" * 50)

        try:
            # 1. 环境检查
            await self.check_environment()

            # 2. 安装依赖
            await self.install_dependencies()

            # 3. 检查 Ollama
            await self.check_ollama()

            # 4. 初始化项目
            await self.initialize_project()

            # 5. 生成配置
            await self.generate_config()

            # 6. 显示后续步骤
            self.show_next_steps()

        except Exception as e:
            print(f"❌ 快速启动失败: {e}")
            return False

        return True

    async def check_environment(self):
        """检查环境依赖"""
        print("\n🔍 检查环境依赖...")

        # 检查 Python 版本
        python_version = sys.version_info
        if python_version.major < 3 or (python_version.major == 3 and python_version.minor < 8):
            raise Exception(f"需要 Python 3.8+，当前版本: {python_version.major}.{python_version.minor}")

        print(f"✅ Python {python_version.major}.{python_version.minor}.{python_version.micro}")

        # 检查 pip
        try:
            subprocess.run([sys.executable, "-m", "pip", "--version"],
                         check=True, capture_output=True)
            print("✅ pip 可用")
        except subprocess.CalledProcessError:
            raise Exception("pip 不可用，请安装 pip")

    async def install_dependencies(self):
        """安装 Python 依赖"""
        print("\n📦 安装 Python 依赖...")

        if not self.requirements_file.exists():
            print("⚠️  requirements.txt 不存在，跳过依赖安装")
            return

        try:
            # 安装依赖
            result = subprocess.run([
                sys.executable, "-m", "pip", "install", "-r", str(self.requirements_file)
            ], capture_output=True, text=True)

            if result.returncode == 0:
                print("✅ Python 依赖安装完成")
            else:
                print(f"⚠️  依赖安装警告: {result.stderr}")

        except Exception as e:
            print(f"❌ 依赖安装失败: {e}")

    async def check_ollama(self):
        """检查 Ollama 安装和配置"""
        print("\n🤖 检查 Ollama 环境...")

        # 检查 Ollama 是否安装
        if not shutil.which("ollama"):
            print("❌ Ollama 未安装")
            print("💡 请访问 https://ollama.ai 安装 Ollama")
            print("💡 或运行: curl -fsSL https://ollama.ai/install.sh | sh")
            return False

        print("✅ Ollama 已安装")

        # 检查 Ollama 服务状态
        try:
            result = subprocess.run(["ollama", "list"],
                                  capture_output=True, text=True, timeout=10)
            if result.returncode == 0:
                print("✅ Ollama 服务运行正常")

                # 检查推荐模型
                if "qwen2.5:14b" in result.stdout:
                    print("✅ 推荐模型 qwen2.5:14b 已安装")
                else:
                    print("⚠️  推荐模型 qwen2.5:14b 未安装")
                    print("💡 建议运行: ollama pull qwen2.5:14b")

            else:
                print("⚠️  Ollama 服务可能未启动")
                print("💡 请运行: ollama serve")

        except subprocess.TimeoutExpired:
            print("⚠️  Ollama 服务检查超时")
            print("💡 请确保 Ollama 服务正在运行")
        except Exception as e:
            print(f"⚠️  Ollama 检查出错: {e}")

        return True

    async def initialize_project(self):
        """初始化项目结构"""
        print("\n🏗️  初始化项目结构...")

        if not self.main_script.exists():
            raise Exception(f"主脚本不存在: {self.main_script}")

        try:
            # 运行项目初始化
            result = subprocess.run([
                sys.executable, str(self.main_script), "--init"
            ], capture_output=True, text=True, cwd=self.project_root)

            if result.returncode == 0:
                print("✅ 项目结构初始化完成")
                print(result.stdout)
            else:
                print(f"❌ 项目初始化失败: {result.stderr}")

        except Exception as e:
            print(f"❌ 项目初始化异常: {e}")

    async def generate_config(self):
        """生成配置文件"""
        print("\n⚙️  生成配置文件...")

        try:
            # 生成默认配置
            result = subprocess.run([
                sys.executable, str(self.main_script), "--save-config"
            ], capture_output=True, text=True, cwd=self.project_root)

            if result.returncode == 0:
                print("✅ 配置文件生成完成")
            else:
                print(f"⚠️  配置生成警告: {result.stderr}")

        except Exception as e:
            print(f"❌ 配置生成失败: {e}")

    def show_next_steps(self):
        """显示后续步骤"""
        print("\n🎯 快速启动完成! 后续步骤:")
        print("=" * 50)

        print("\n1. 🚀 启动服务:")
        print(f"   cd {self.project_root}")
        print("   python main.py --dev")

        print("\n2. 🔧 配置 AI 工具:")

        print("\n   📌 Claude Code 配置:")
        print("   # 在 Claude Code 中添加 MCP 服务器")
        print("   claude config mcp-servers add universal-coding-assistant")

        print("\n   📌 Gemini CLI 配置:")
        print("   # 配置 Gemini CLI 连接 MCP 服务器")
        print("   gemini config mcp-servers.universal-coding-assistant.command \"python main.py\"")

        print("\n   📌 Cursor 配置:")
        print("   # 在 .cursor/mcp.json 中添加配置")

        print("\n3. 🧪 测试功能:")
        print("   claude \"初始化 ADDP 项目结构\"")
        print("   gemini \"优化这个查询: 实现用户登录\"")

        print("\n4. 📚 查看文档:")
        print("   📖 README.md - 项目概览")
        print("   📋 TARGET.md - 详细规格")
        print("   📁 .addp/ - 项目产出目录")

        print("\n🎉 开始使用 Universal AI Coding Framework!")

def main():
    """主函数"""
    try:
        quick_start = QuickStart()
        success = asyncio.run(quick_start.run())
        sys.exit(0 if success else 1)
    except KeyboardInterrupt:
        print("\n👋 用户取消")
        sys.exit(0)
    except Exception as e:
        print(f"❌ 快速启动异常: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()