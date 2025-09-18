#!/usr/bin/env python3
"""
Universal AI Coding Framework - MCP Server
==========================================

主入口文件，启动 MCP 服务器并处理标准 I/O 通信。

用法:
    python main.py                    # 标准 MCP 模式 (stdio)
    python main.py --config config.json  # 指定配置文件
    python main.py --init            # 初始化项目结构
    python main.py --dev             # 开发模式
"""

import asyncio
import argparse
import sys
import logging
from pathlib import Path

# 添加源代码路径
sys.path.insert(0, str(Path(__file__).parent))

from src.mcp_server.server import create_mcp_server
from src.mcp_server.config import load_config, save_default_config, UniversalConfig
from src.mcp_server.tools.project_tools import ProjectInitializer

# 配置日志
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

async def init_project():
    """初始化项目结构"""
    try:
        print("🚀 开始初始化 Universal AI Coding Framework 项目结构...")

        initializer = ProjectInitializer()
        result = await initializer.initialize_structure()

        if result["success"]:
            print(f"✅ 项目初始化完成!")
            print(f"📁 创建了 {result['directories_created']} 个目录")
            print(f"📄 生成了 {result['files_created']} 个文件")
            print(f"🔧 配置了 {result['configs_created']} 个配置")
            print("\n🎯 下一步:")
            print("  1. 启动 Ollama: ollama serve")
            print("  2. 拉取模型: ollama pull qwen2.5:14b")
            print("  3. 启动 MCP 服务: python main.py")
            print("  4. 配置 AI 工具连接到 MCP 服务器")
            return True
        else:
            print("❌ 项目初始化失败")
            return False

    except Exception as e:
        print(f"❌ 初始化过程中出错: {e}")
        return False

async def run_dev_mode(config: UniversalConfig):
    """开发模式"""
    print("🔧 启动开发模式...")
    print(f"📡 Ollama 端点: {config.ollama.endpoint}")
    print(f"🤖 使用模型: {config.ollama.model}")
    print(f"📁 ADDP 目录: {config.project.addp_directory}")

    # 检查 .addp 目录是否存在
    addp_path = Path(config.project.addp_directory)
    if not addp_path.exists():
        print("⚠️  .addp 目录不存在，是否需要初始化? (y/n)")
        response = input().strip().lower()
        if response == 'y':
            await init_project()
        else:
            print("❌ 开发模式需要 .addp 目录结构")
            return

    # 创建并启动服务器
    server = await create_mcp_server(config)
    print("✅ MCP 服务器已启动 (开发模式)")
    print("💡 在其他终端中测试 MCP 工具调用")
    print("🛑 按 Ctrl+C 停止服务器")

    try:
        # 保持服务运行
        await asyncio.Event().wait()
    except KeyboardInterrupt:
        print("\n🛑 服务器关闭")

async def run_mcp_server(config: UniversalConfig):
    """标准 MCP 服务器模式"""
    logger.info("启动 Universal AI Coding Framework MCP Server")

    # 创建服务器实例
    server = await create_mcp_server(config)

    # 在标准 MCP 模式下，服务器通过 stdio 与客户端通信
    # 这里应该实现标准的 MCP 协议通信
    # 由于这需要与具体的 MCP 客户端实现配合，这里提供框架结构

    try:
        # MCP 标准通信循环
        # 实际实现需要根据 MCP 协议规范来处理 JSON-RPC 消息
        logger.info("MCP 服务器正在监听 stdio...")

        # 模拟服务器运行状态
        await asyncio.Event().wait()

    except KeyboardInterrupt:
        logger.info("服务器关闭")
    except Exception as e:
        logger.error(f"服务器运行错误: {e}")
        raise

def main():
    """主函数"""
    parser = argparse.ArgumentParser(
        description="Universal AI Coding Framework MCP Server"
    )
    parser.add_argument(
        "--config",
        type=str,
        help="配置文件路径"
    )
    parser.add_argument(
        "--init",
        action="store_true",
        help="初始化项目结构"
    )
    parser.add_argument(
        "--dev",
        action="store_true",
        help="开发模式"
    )
    parser.add_argument(
        "--save-config",
        action="store_true",
        help="保存默认配置文件"
    )

    args = parser.parse_args()

    # 保存默认配置
    if args.save_config:
        config_path = args.config or ".addp/configs/mcp/server_config.json"
        save_default_config(config_path)
        print(f"✅ 默认配置已保存到: {config_path}")
        return

    # 初始化项目
    if args.init:
        success = asyncio.run(init_project())
        sys.exit(0 if success else 1)

    # 加载配置
    try:
        config = load_config(args.config)
    except Exception as e:
        print(f"❌ 配置加载失败: {e}")
        print("💡 尝试使用 --save-config 生成默认配置")
        sys.exit(1)

    # 运行服务器
    try:
        if args.dev:
            asyncio.run(run_dev_mode(config))
        else:
            asyncio.run(run_mcp_server(config))
    except KeyboardInterrupt:
        print("\n👋 再见!")
    except Exception as e:
        logger.error(f"程序异常退出: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()