"""
查询优化工具
============

使用 Ollama 本地模型优化用户查询，提高 AI 工具协作的精确度和效率。
实现智能查询理解、上下文增强和多级优化策略。
"""

import aiohttp
import json
import logging
from typing import Dict, List, Any, Optional
from pathlib import Path
from datetime import datetime

logger = logging.getLogger(__name__)

class QueryOptimizer:
    """查询优化器，使用 Ollama 本地模型"""

    def __init__(self, endpoint: str = "http://localhost:11434", model: str = "qwen2.5:14b"):
        self.endpoint = endpoint.rstrip('/')
        self.model = model
        self.cache_path = Path(".addp/queries/cache")
        self.cache_path.mkdir(parents=True, exist_ok=True)

    async def optimize(
        self,
        user_input: str,
        context: str = "",
        optimization_level: str = "smart"
    ) -> Dict[str, Any]:
        """
        优化用户查询

        Args:
            user_input: 原始用户输入
            context: 项目上下文信息
            optimization_level: 优化级别 (basic, smart, detailed)

        Returns:
            优化结果包含优化后的查询、改进点和建议
        """
        try:
            # 1. 检查缓存
            cached_result = await self._check_cache(user_input, context, optimization_level)
            if cached_result:
                logger.info("使用缓存的查询优化结果")
                return cached_result

            # 2. 构建优化提示
            optimization_prompt = self._build_optimization_prompt(
                user_input, context, optimization_level
            )

            # 3. 调用 Ollama 进行优化
            ollama_response = await self._call_ollama(optimization_prompt)

            # 4. 解析和结构化结果
            optimization_result = self._parse_optimization_result(
                user_input, ollama_response, optimization_level
            )

            # 5. 保存到缓存
            await self._save_to_cache(user_input, context, optimization_level, optimization_result)

            # 6. 记录分析数据
            await self._record_analytics(user_input, optimization_result)

            return optimization_result

        except Exception as e:
            logger.error(f"查询优化失败: {e}")
            # 返回基础优化结果
            return {
                "optimized_query": user_input,
                "improvements": "查询优化服务暂时不可用，使用原始查询",
                "suggestions": "请检查 Ollama 服务是否正常运行",
                "confidence": 0.0,
                "optimization_level": optimization_level,
                "error": str(e)
            }

    def _build_optimization_prompt(
        self,
        user_input: str,
        context: str,
        optimization_level: str
    ) -> str:
        """构建优化提示"""

        base_prompt = f"""你是一个专业的 AI 编程助手查询优化器。你的任务是将模糊或不精确的用户请求转化为清晰、可执行的技术指令。

**用户原始输入**: {user_input}

**项目上下文**: {context or "无特定项目上下文"}

**优化目标**:
1. 明确技术意图和具体需求
2. 添加必要的技术细节和约束
3. 提供可执行的具体步骤
4. 确保符合最佳实践

**输出格式** (严格遵循 JSON 格式):
{{
    "optimized_query": "优化后的精确查询",
    "improvements": "具体改进的方面",
    "suggestions": "额外的建议和注意事项",
    "confidence": 0.85,
    "technical_details": {{
        "framework": "相关技术栈",
        "complexity": "任务复杂度评估",
        "estimated_time": "预估完成时间"
    }}
}}"""

        # 根据优化级别添加特定指导
        if optimization_level == "basic":
            base_prompt += """

**基础优化指导**:
- 澄清基本的技术术语
- 确保查询明确可理解
- 添加基本的技术约束"""

        elif optimization_level == "smart":
            base_prompt += """

**智能优化指导**:
- 基于上下文推理用户真实意图
- 添加相关的技术最佳实践
- 预测可能的技术难点和解决方案
- 建议具体的实施步骤"""

        elif optimization_level == "detailed":
            base_prompt += """

**详细优化指导**:
- 提供全面的技术分析
- 给出多种实现方案的对比
- 包含详细的验收标准
- 预测潜在风险和缓解策略
- 提供完整的开发路径"""

        base_prompt += """

**特别注意**:
- 如果用户请求涉及性能优化，要具体说明优化目标和指标
- 如果涉及新功能开发，要明确功能边界和验收标准
- 如果涉及重构，要说明重构范围和风险控制
- 始终考虑 TDD、代码质量和维护性

请严格按照 JSON 格式输出，不要添加额外的解释文字。"""

        return base_prompt

    async def _call_ollama(self, prompt: str) -> str:
        """调用 Ollama API"""
        try:
            async with aiohttp.ClientSession() as session:
                payload = {
                    "model": self.model,
                    "prompt": prompt,
                    "stream": False,
                    "options": {
                        "temperature": 0.7,
                        "top_p": 0.9,
                        "max_tokens": 2048
                    }
                }

                async with session.post(
                    f"{self.endpoint}/api/generate",
                    json=payload,
                    timeout=aiohttp.ClientTimeout(total=30)
                ) as response:
                    if response.status == 200:
                        result = await response.json()
                        return result.get("response", "")
                    else:
                        raise Exception(f"Ollama API 错误: {response.status}")

        except aiohttp.ClientError as e:
            raise Exception(f"网络请求失败: {e}")
        except Exception as e:
            raise Exception(f"Ollama 调用失败: {e}")

    def _parse_optimization_result(
        self,
        original_input: str,
        ollama_response: str,
        optimization_level: str
    ) -> Dict[str, Any]:
        """解析 Ollama 响应"""
        try:
            # 尝试解析 JSON 响应
            if ollama_response.strip().startswith('{'):
                parsed = json.loads(ollama_response.strip())

                # 验证必需字段
                required_fields = ["optimized_query", "improvements", "suggestions", "confidence"]
                if all(field in parsed for field in required_fields):
                    parsed["original_query"] = original_input
                    parsed["optimization_level"] = optimization_level
                    parsed["timestamp"] = datetime.now().isoformat()
                    return parsed

        except json.JSONDecodeError:
            pass

        # 如果 JSON 解析失败，创建基础优化结果
        lines = ollama_response.split('\n')
        optimized_query = original_input

        # 简单启发式解析
        for line in lines:
            if "优化" in line or "建议" in line:
                potential_query = line.strip().replace("优化后的查询：", "").replace("建议：", "")
                if len(potential_query) > len(original_input) * 0.8:
                    optimized_query = potential_query
                    break

        return {
            "optimized_query": optimized_query,
            "improvements": "基于模型响应的基础优化",
            "suggestions": "请检查 Ollama 模型输出格式",
            "confidence": 0.6,
            "optimization_level": optimization_level,
            "original_query": original_input,
            "timestamp": datetime.now().isoformat(),
            "raw_response": ollama_response[:500]  # 保留部分原始响应用于调试
        }

    async def _check_cache(
        self,
        user_input: str,
        context: str,
        optimization_level: str
    ) -> Optional[Dict[str, Any]]:
        """检查缓存中是否有相同的查询优化结果"""
        try:
            cache_key = self._generate_cache_key(user_input, context, optimization_level)
            cache_file = self.cache_path / f"{cache_key}.json"

            if cache_file.exists():
                with open(cache_file, 'r', encoding='utf-8') as f:
                    cached_data = json.load(f)

                # 检查缓存时效性（24小时）
                cached_time = datetime.fromisoformat(cached_data.get("timestamp", ""))
                current_time = datetime.now()
                if (current_time - cached_time).total_seconds() < 86400:  # 24小时
                    cached_data["from_cache"] = True
                    return cached_data

        except Exception as e:
            logger.warning(f"缓存检查失败: {e}")

        return None

    async def _save_to_cache(
        self,
        user_input: str,
        context: str,
        optimization_level: str,
        result: Dict[str, Any]
    ):
        """保存优化结果到缓存"""
        try:
            cache_key = self._generate_cache_key(user_input, context, optimization_level)
            cache_file = self.cache_path / f"{cache_key}.json"

            with open(cache_file, 'w', encoding='utf-8') as f:
                json.dump(result, f, indent=2, ensure_ascii=False)

        except Exception as e:
            logger.warning(f"缓存保存失败: {e}")

    def _generate_cache_key(self, user_input: str, context: str, optimization_level: str) -> str:
        """生成缓存键"""
        import hashlib

        content = f"{user_input}|{context}|{optimization_level}"
        return hashlib.md5(content.encode('utf-8')).hexdigest()[:16]

    async def _record_analytics(self, original_query: str, optimization_result: Dict[str, Any]):
        """记录分析数据"""
        try:
            analytics_path = Path(".addp/analytics/metrics")
            analytics_path.mkdir(parents=True, exist_ok=True)

            analytics_data = {
                "timestamp": datetime.now().isoformat(),
                "original_query_length": len(original_query),
                "optimized_query_length": len(optimization_result.get("optimized_query", "")),
                "confidence": optimization_result.get("confidence", 0.0),
                "optimization_level": optimization_result.get("optimization_level", "unknown"),
                "improvement_ratio": len(optimization_result.get("optimized_query", "")) / max(len(original_query), 1),
                "from_cache": optimization_result.get("from_cache", False)
            }

            # 添加到每日分析文件
            today = datetime.now().strftime("%Y-%m-%d")
            analytics_file = analytics_path / f"query_optimization_{today}.jsonl"

            with open(analytics_file, 'a', encoding='utf-8') as f:
                f.write(json.dumps(analytics_data, ensure_ascii=False) + '\n')

        except Exception as e:
            logger.warning(f"分析数据记录失败: {e}")

    async def get_optimization_stats(self) -> Dict[str, Any]:
        """获取优化统计数据"""
        try:
            analytics_path = Path(".addp/analytics/metrics")

            if not analytics_path.exists():
                return {"message": "暂无优化统计数据"}

            stats = {
                "total_optimizations": 0,
                "cache_hit_rate": 0.0,
                "average_confidence": 0.0,
                "optimization_levels": {},
                "daily_stats": {}
            }

            # 读取所有分析文件
            for analytics_file in analytics_path.glob("query_optimization_*.jsonl"):
                date = analytics_file.stem.replace("query_optimization_", "")
                daily_count = 0
                daily_confidence = []
                cache_hits = 0

                with open(analytics_file, 'r', encoding='utf-8') as f:
                    for line in f:
                        try:
                            data = json.loads(line.strip())
                            daily_count += 1
                            daily_confidence.append(data.get("confidence", 0.0))

                            if data.get("from_cache", False):
                                cache_hits += 1

                            level = data.get("optimization_level", "unknown")
                            stats["optimization_levels"][level] = stats["optimization_levels"].get(level, 0) + 1

                        except json.JSONDecodeError:
                            continue

                stats["daily_stats"][date] = {
                    "count": daily_count,
                    "average_confidence": sum(daily_confidence) / max(len(daily_confidence), 1),
                    "cache_hits": cache_hits
                }

                stats["total_optimizations"] += daily_count

            # 计算总体统计
            if stats["total_optimizations"] > 0:
                total_cache_hits = sum(day["cache_hits"] for day in stats["daily_stats"].values())
                stats["cache_hit_rate"] = total_cache_hits / stats["total_optimizations"]

                all_confidences = []
                for day in stats["daily_stats"].values():
                    all_confidences.extend([day["average_confidence"]] * day["count"])

                stats["average_confidence"] = sum(all_confidences) / len(all_confidences)

            return stats

        except Exception as e:
            logger.error(f"获取统计数据失败: {e}")
            return {"error": str(e)}

    async def clear_cache(self, older_than_days: int = 7) -> Dict[str, Any]:
        """清理过期缓存"""
        try:
            cache_path = Path(".addp/queries/cache")
            current_time = datetime.now()
            removed_count = 0

            if cache_path.exists():
                for cache_file in cache_path.glob("*.json"):
                    try:
                        # 检查文件修改时间
                        file_time = datetime.fromtimestamp(cache_file.stat().st_mtime)
                        if (current_time - file_time).days > older_than_days:
                            cache_file.unlink()
                            removed_count += 1
                    except Exception:
                        continue

            return {
                "success": True,
                "removed_files": removed_count,
                "message": f"已清理 {removed_count} 个超过 {older_than_days} 天的缓存文件"
            }

        except Exception as e:
            logger.error(f"缓存清理失败: {e}")
            return {"error": str(e)}