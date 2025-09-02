# Agent系統優化計劃

## 📍 理解確認
我理解你需要：基於剛才的深度分析，制定一個系統化的Agent優化計劃，解決職責重疊、相互污染和目的不明確的問題。

## 🎯 優化目標
- **減少Agent數量**：從45個減少到約30個（33%精簡）
- **提升Token效率**：減少40-50%的token浪費
- **提高選擇準確性**：從60%提升到90%+
- **清晰職責邊界**：每個agent單一職責，無重疊

## 📊 現狀問題總結

### 高風險（P0）
1. **Kotlin生態混亂**：3個agents嚴重重疊
2. **TypeScript碎片化**：3個agents功能分散
3. **移動開發衝突**：邊界不清，多重觸發

### 中風險（P1）
1. **Context Detectors冗余**：與orchestrator功能重疊
2. **Token-Loader不可用**：理論設計無法實現
3. **Work-Coordinator重複**：與orchestrator職責衝突

## 🚀 實施計劃

### Phase 1：立即執行（Day 1-2）🔥

#### 1.1 Kotlin生態整合
```bash
# 執行步驟
1. 備份現有agents
   git add agents/languages/kotlin-*.md
   git commit -m "backup: Kotlin agents before optimization"

2. 刪除冗余agents
   rm agents/languages/kotlin-expert.md           # 過於簡化
   rm agents/languages/kotlin-polyglot-master.md  # 過度複雜

3. 優化保留的agent
   - 增強 android-kotlin-architect.md (Android專門)
   - 創建 kotlin-backend-expert.md (Ktor/Spring專門)

4. 更新orchestrator選擇邏輯
   - Kotlin + Android → android-kotlin-architect
   - Kotlin + (Ktor|Spring) → kotlin-backend-expert
```

#### 1.2 TypeScript統一整合
```bash
# 執行步驟
1. 合併三個TypeScript agents
   - 提取 typescript-expert-core 的簡潔規則
   - 整合 typescript-expert-examples 的範例
   - 保留 typescript-fullstack-expert 的完整功能

2. 創建統一的 typescript-expert.md
   - 包含核心規則（必需）
   - 包含常用範例（精選）
   - 包含框架支援（React/Vue/Node）

3. 刪除碎片文件
   rm agents/languages/typescript-expert-*.md
```

#### 1.3 移除Context Detectors
```bash
# 執行步驟
1. 提取有用的檢測邏輯
   - 從每個 *-context-detector.md 提取規則
   - 整合到 orchestrator.md 的選擇邏輯中

2. 刪除所有context detector agents
   rm agents/languages/*-context-detector.md

3. 更新orchestrator v2.1
   - 添加語言特定的檢測規則
   - 增強關鍵詞匹配邏輯
```

### Phase 2：短期優化（Day 3-7）📈

#### 2.1 明確移動開發邊界
```yaml
mobile-developer:
  focus: "Native iOS (Swift) and Flutter only"
  exclude: ["React Native", "Android"]
  
android-kotlin-architect:
  focus: "Native Android with Kotlin/Compose"
  exclude: ["React Native", "Flutter"]
  
frontend-developer:
  focus: "Web and React Native"
  exclude: ["Native iOS", "Native Android", "Flutter"]
```

#### 2.2 清理無用agents
```bash
# 需要刪除或合併的agents
- token-efficient-loader.md  # 無法實現
- work-coordinator.md        # 與orchestrator重複
- kotlin-context-detector.md # 已整合到orchestrator

# 需要增強的agents
- code-reviewer.md          # 增加更多審查規則
- performance-optimizer.md  # 添加具體優化策略
- test-automator.md         # 支援更多測試框架
```

#### 2.3 優化Orchestrator選擇邏輯
```typescript
// 新的選擇邏輯結構
interface SmartSelection {
  // 1. 語言檢測（基於文件擴展名）
  detectLanguage(files: string[]): Language
  
  // 2. 框架檢測（基於import和關鍵詞）
  detectFramework(content: string): Framework
  
  // 3. 任務類型（基於用戶描述）
  detectTaskType(request: string): TaskType
  
  // 4. 精確選擇（單一agent優先）
  selectAgent(lang: Language, framework: Framework, task: TaskType): Agent
}
```

### Phase 3：中期改進（Week 2）🔧

#### 3.1 建立Agent能力矩陣
```markdown
| 能力領域 | 負責Agent | 觸發條件 | Token預算 |
|---------|-----------|---------|-----------|
| Android開發 | android-kotlin-architect | *.kt + Android | ~150k |
| iOS開發 | mobile-developer | *.swift | ~150k |
| React前端 | frontend-developer | *.tsx + React | ~120k |
| Node後端 | api-architect | *.js + Express | ~100k |
| Python ML | python-ml-specialist | *.py + ML庫 | ~170k |
```

#### 3.2 實施測試驗證
```bash
# 創建測試場景
tests/
├── kotlin-android-test.kt    # 應觸發 android-kotlin-architect
├── kotlin-backend-test.kt    # 應觸發 kotlin-backend-expert  
├── react-native-test.tsx     # 應觸發 frontend-developer
├── swift-ios-test.swift      # 應觸發 mobile-developer
└── typescript-web-test.ts    # 應觸發 typescript-expert
```

#### 3.3 建立監控機制
```yaml
# .claude/memory/agent-metrics.yml
metrics:
  agent_usage:
    - agent: "android-kotlin-architect"
      calls: 0
      token_avg: 0
      success_rate: 0
      
  selection_accuracy:
    correct: 0
    incorrect: 0
    accuracy: 0%
    
  token_efficiency:
    before_optimization: 800k
    after_optimization: 0
    savings: 0%
```

### Phase 4：長期維護（Month 1）🚀

#### 4.1 持續優化
- 每週review agent使用統計
- 識別未使用的agents並考慮移除
- 根據實際使用優化選擇邏輯

#### 4.2 文檔更新
- 更新 agents/README.md 反映新結構
- 創建 agents/ARCHITECTURE.md 說明設計原則
- 維護 agents/CHANGELOG.md 記錄變更

#### 4.3 社區反饋
- 收集用戶對agent選擇的反饋
- 調整不準確的選擇規則
- 添加缺失的專業agents

## ⚠️ 注意事項

### 風險控制
1. **備份優先**：每次刪除前先commit
2. **逐步執行**：一次優化一個語言生態
3. **測試驗證**：每個改動後測試選擇準確性
4. **回滾準備**：保留原始agents 30天

### 兼容性考慮
1. **Orchestrator更新**：確保v2.1正確加載
2. **緩存清理**：可能需要重啟Claude Code
3. **用戶教育**：更新文檔說明變更

## ✅ 檢查清單

### Phase 1 (立即)
- [ ] 備份所有agents到Git
- [ ] 整合Kotlin生態系統
- [ ] 合併TypeScript agents
- [ ] 移除context detectors
- [ ] 更新orchestrator選擇邏輯
- [ ] 測試基本場景

### Phase 2 (本週)
- [ ] 明確移動開發邊界
- [ ] 清理無用agents
- [ ] 建立能力矩陣
- [ ] 創建測試套件
- [ ] 實施監控機制

### Phase 3 (本月)
- [ ] 優化選擇準確性到90%+
- [ ] Token使用減少40%+
- [ ] 完成所有文檔更新
- [ ] 建立維護流程

## 🚀 建議開始

**第一步：執行Kotlin生態整合**
```bash
# 1. 先備份
git add -A && git commit -m "backup: before agent optimization"

# 2. 開始Kotlin整合
cd agents/languages
# 按計劃執行...
```

**預期效果**：
- 立即減少3-5個冗余agents
- Token使用減少30%+
- 選擇準確性提升到80%+

## 📈 成功指標

### 短期（1週）
- ✅ Agent數量 < 35個
- ✅ 無明顯職責重疊
- ✅ 選擇準確性 > 80%

### 中期（2週）
- ✅ Agent數量 ≈ 30個
- ✅ Token平均使用 < 400k
- ✅ 選擇準確性 > 85%

### 長期（1月）
- ✅ Agent數量穩定在30個
- ✅ Token平均使用 < 300k
- ✅ 選擇準確性 > 90%
- ✅ 用戶滿意度提升

---

這個優化計劃將徹底解決Agent系統的混亂問題，實現高效、準確的智能選擇！