# OLLAMA 模型选型与参数策略

## 目标与范围
- 场景：需求理解→要素梳理→命令化重写→打包后交给 AI CLI；优先本地执行，减少 token。
- 要求：结构化输出稳定（JSON/键值对）、中文/双语良好、低延迟、可复现（低温+固定 seed）。

## 任务 → 模型映射（建议）
- 需求理解/要点提取（optimize.user）
  - 首选：`qwen2.5:7b`
  - 回退：`qwen2.5:14b`
- 命令化重写/系统提示词优化（optimize.system, command rewrite）
  - 首选：`qwen2.5:14b`
  - 备选：`gemma2:9b-instruct` 或 `llama3.1:8b-instruct`
- 英文优先场景（快速/紧凑）
  - 首选：`llama3.1:8b-instruct`，回退 `qwen2.5:14b`
- 长上下文（大规格/模板汇总）
  - 提高 `num_ctx`；仍不足则“分块→摘要→汇总”。

## 推理参数（TDD 友好默认）
- `temperature`: 0.1–0.2
- `top_p`: 0.9
- `max_tokens`: 1024–2048（按任务）
- `num_ctx`: 8192（或模型上限）
- `seed`: 42（便于复现）
- `format`: "json"（若模型/版本支持；否则使用 code-fence+JSON 并容错解析）
- `stop`: ["```"]（避免附加说明）

## 路由与回退策略
- 两段式：7B/8B 先“分类+要素抽取+不确定性评估（置信度/冲突数）”；低置信或存在冲突时，升级 14B 做“消歧+命令化重写”。
- 失败/超时/非 JSON：指数退避→同家族上一级模型→记录原因到 `.addp/analytics`。

## 评估指标（纳入测试）
- JSON 合规率 ≥ 99%（N≥200 常见模糊输入集）
- 信息完整性：覆盖 6 类要素（目标/范围/约束/接口/验收/风险）
- P95 时延：结构化优化 ≤3s（7B/8B），深度重写 ≤6s（14B）
- 复现性：同 seed 字段级差异低（结构一致）

## 配置示例
- 环境变量
```
OLLAMA_ENDPOINT=http://localhost:11434
OLLAMA_MODEL=qwen2.5:7b
```
- 任务→模型映射（建议放置：`.addp/configs/ollama/models.json`）
```json
{
  "optimize": {
    "user":    {"primary": "qwen2.5:7b",   "fallback": ["qwen2.5:14b"]},
    "system":  {"primary": "qwen2.5:14b",  "fallback": ["gemma2:9b-instruct", "llama3.1:8b-instruct"]},
    "iterate": {"primary": "llama3.1:8b-instruct", "fallback": ["qwen2.5:7b"]}
  },
  "options": {
    "temperature": 0.2,
    "top_p": 0.9,
    "max_tokens": 2048,
    "num_ctx": 8192,
    "seed": 42,
    "format": "json",
    "stop": ["```"]
  }
}
```

## 模型拉取（示例）
```
ollama pull qwen2.5:7b
ollama pull qwen2.5:14b
ollama pull llama3.1:8b-instruct
ollama pull mistral:7b-instruct
ollama pull gemma2:9b-instruct
```

## 资源与注意事项
- 7–9B（Q4 量化）：CPU 可用；GPU 4–6GB VRAM；P95≈1–3s。
- 14B（Q4 量化）：建议 ≥8–10GB VRAM；P95≈3–6s；不足则降级 7–9B。
- 不同模型对 `format: json` 支持程度差异较大：务必保留“去 code-fence + 容错解析”兜底逻辑。

