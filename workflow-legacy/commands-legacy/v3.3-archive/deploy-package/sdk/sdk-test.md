---
arguments: optional
format: "[unit|integration|compat|performance|all]"
examples:
  - "/sdk-test - 生成完整測試套件"
  - "/sdk-test unit - 單元測試"
  - "/sdk-test compat - 兼容性測試"
  - "/sdk-test performance - 性能測試"
---

# SDK 測試專家

為 SDK 生成全面的測試套件，確保穩定性和兼容性：

## 🧪 測試策略

### 1. 單元測試 (unit)
- **核心功能測試**：每個公共 API
- **邊界條件**：極限值、空值、異常輸入
- **錯誤場景**：異常處理、錯誤恢復
- **內部邏輯**：關鍵算法和數據結構

### 2. 集成測試 (integration)
- **模塊協作**：多個組件的交互
- **真實環境**：實際使用場景
- **外部依賴**：第三方服務集成
- **並發測試**：多線程/異步場景

### 3. 兼容性測試 (compat)
- **版本兼容**：不同版本間的兼容性
- **平台兼容**：不同運行環境
- **依賴兼容**：不同依賴版本
- **向後兼容**：API 變更影響

### 4. 性能測試 (performance)
- **基準測試**：關鍵操作的性能
- **壓力測試**：高負載下的表現
- **內存分析**：內存使用和洩漏
- **性能回歸**：版本間性能對比

## 📋 測試模板

### SDK 測試結構
```
tests/
├── unit/              # 單元測試
│   ├── core/         # 核心功能
│   ├── utils/        # 工具函數
│   └── edge-cases/   # 邊界情況
├── integration/       # 集成測試
│   ├── scenarios/    # 使用場景
│   └── external/     # 外部集成
├── compatibility/     # 兼容性測試
│   ├── versions/     # 版本測試
│   └── platforms/    # 平台測試
├── performance/       # 性能測試
│   ├── benchmarks/   # 基準測試
│   └── stress/       # 壓力測試
└── fixtures/         # 測試數據
```

### 測試用例示例

#### 單元測試
```[語言]
describe('SDK Core API', () => {
  // 正常情況
  test('should handle valid input', () => {
    // 準備
    // 執行
    // 斷言
  });

  // 邊界條件
  test('should handle edge cases', () => {
    // 空值、極值、特殊字符
  });

  // 錯誤處理
  test('should throw meaningful errors', () => {
    // 錯誤類型、錯誤信息
  });
});
```

#### 兼容性測試
```[語言]
// 測試不同版本的行為
describe('Version Compatibility', () => {
  test('v1.x API still works', () => {
    // 舊 API 調用方式
    // 確保向後兼容
  });

  test('deprecated features show warnings', () => {
    // 廢棄功能提示
  });
});
```

## 🎯 測試覆蓋要求

### SDK 特定要求
- **公共 API 100% 覆蓋**：所有導出的接口
- **錯誤路徑 90% 覆蓋**：異常處理完整
- **文檔示例可運行**：所有示例都有測試
- **破壞性變更檢測**：自動發現不兼容

### 質量指標
```yaml
coverage:
  statements: 90%
  branches: 85%
  functions: 95%
  lines: 90%

performance:
  regression_threshold: 10%
  memory_limit: 100MB
  response_time: <100ms
```

## 🔄 持續測試

### 測試自動化
- **提交時**：運行快速測試
- **PR 時**：完整測試套件
- **發布前**：兼容性和性能測試
- **定期**：依賴更新測試

### 測試報告
```markdown
## 測試報告
- ✅ 單元測試：245/245 通過
- ✅ 集成測試：48/48 通過
- ✅ 兼容性：所有版本通過
- ⚠️ 性能：比 v1.2 慢 5%

## 覆蓋率
- 語句：94.5%
- 分支：88.2%
- 函數：96.8%
```

## 💡 SDK 測試最佳實踐

1. **測試公共 API，不測試內部實現**
2. **模擬外部依賴，確保測試獨立**
3. **使用真實數據，覆蓋實際場景**
4. **關注錯誤信息的可讀性**
5. **性能測試要有基準對比**

全面的測試是 SDK 品質的保證！