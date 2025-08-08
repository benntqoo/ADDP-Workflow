# SDK 開發工作流指南

使用 Claude Code v3.0 + SDK 專用命令開發高品質的 SDK/Library。

## 🎯 SDK 開發生命週期

### 1️⃣ 設計階段
```bash
# 設計 API 接口
/sdk-design "緩存系統"

# 記錄設計決策
/learn "選擇 LRU 算法因為平衡了性能和內存使用"
```

### 2️⃣ 實現階段
```bash
# 規劃實現步驟
/plan "實現 LRU 緩存核心邏輯"

# 開發過程
# ... 編碼 ...

# 代碼審查
/check
```

### 3️⃣ 示例開發
```bash
# 生成示例代碼
/sdk-example basic      # 基礎示例
/sdk-example advanced   # 進階示例
/sdk-example integration # 集成示例
```

### 4️⃣ 測試階段
```bash
# 生成測試套件
/sdk-test unit         # 單元測試
/sdk-test compat       # 兼容性測試
/sdk-test performance  # 性能測試

# 執行測試
/test
```

### 5️⃣ 文檔階段
```bash
# 生成文檔
/sdk-doc api          # API 參考文檔
/sdk-doc guide        # 使用指南
/sdk-doc migration    # 遷移指南
```

### 6️⃣ 發布階段
```bash
# 發布準備
/sdk-release check    # 發布前檢查
/sdk-release minor    # 準備次版本發布
```

## 📋 典型場景

### 場景 1：新增功能
```bash
# 1. 設計 API
/sdk-design "添加批量操作功能"
> 輸出：API 設計方案、向後兼容策略

# 2. 實現功能
/plan "實現批量 get/set/delete 操作"
> 輸出：實施步驟、測試計劃

# 3. 更新示例
/sdk-example advanced
> 輸出：批量操作示例代碼

# 4. 完整測試
/sdk-test all
> 輸出：包含新功能的測試套件

# 5. 更新文檔
/sdk-doc api
> 輸出：更新的 API 文檔
```

### 場景 2：重大重構
```bash
# 1. 評估影響
/sdk-release check
> 輸出：破壞性變更列表

# 2. 設計新架構
/sdk-design "重構為插件架構"
> 輸出：新架構設計、遷移策略

# 3. 創建遷移指南
/sdk-doc migration
> 輸出：詳細的遷移步驟

# 4. 兼容性測試
/sdk-test compat
> 輸出：版本兼容性測試
```

### 場景 3：性能優化
```bash
# 1. 性能基準
/sdk-test performance
> 輸出：當前性能指標

# 2. 優化實現
/plan "優化緩存查找算法"

# 3. 對比測試
/sdk-test performance
> 輸出：優化前後對比

# 4. 記錄優化
/learn "使用哈希表替代數組查找，性能提升 10x"
```

## 🏗️ 項目結構建議

```
my-sdk/
├── src/                    # 源代碼
│   ├── core/              # 核心功能
│   ├── utils/             # 工具函數
│   └── index.ts           # 入口文件
├── tests/                  # 測試文件
│   ├── unit/              # 單元測試
│   ├── integration/       # 集成測試
│   └── compatibility/     # 兼容性測試
├── examples/              # 示例代碼
│   ├── basic/            # 基礎示例
│   ├── advanced/         # 進階示例
│   └── integration/      # 集成示例
├── .claude/              # Claude Code 配置
│   ├── PROJECT_CONTEXT.md # SDK 項目上下文
│   ├── DECISIONS.md       # 技術決策記錄
│   ├── SDK_PRINCIPLES.md  # SDK 設計原則
│   └── docs/              # SDK 文檔 (統一位置)
│       ├── api/           # API 文檔
│       ├── guides/        # 使用指南
│       └── migration/     # 遷移指南
├── CHANGELOG.md          # 變更日誌
├── README.md             # 項目說明
└── package.json          # 項目配置
```

## 💡 SDK 專用配置

### PROJECT_CONTEXT.md 模板
```markdown
# SDK 項目上下文

## 🎯 項目定位
- **SDK 名稱**：[名稱]
- **解決問題**：[核心價值]
- **目標用戶**：[開發者群體]
- **競品分析**：[對比其他方案]

## 🏗️ 設計原則
- **API 優先**：接口穩定性高於實現
- **漸進增強**：基礎功能簡單，高級功能可選
- **零依賴**：最小化外部依賴
- **類型安全**：完整的類型定義

## 📊 版本策略
- **當前版本**：[版本號]
- **發布週期**：[頻率]
- **兼容承諾**：[LTS 策略]
- **廢棄策略**：[提前通知期]

## 🔄 API 狀態
### 穩定 API
- [列表]

### 實驗性 API
- [列表] (標記為 @experimental)

### 廢棄 API
- [列表] (標記為 @deprecated)
```

## 🚀 最佳實踐

### 1. API 設計原則
- **一致性**：相似功能使用相似 API
- **可預測**：行為符合直覺
- **可擴展**：預留擴展點
- **可測試**：易於模擬和測試

### 2. 版本管理
- 遵循語義化版本
- 維護詳細的變更日誌
- 提供清晰的遷移路徑
- 設置合理的廢棄期

### 3. 文檔優先
- 先寫文檔，後寫代碼
- 文檔即規範
- 保持示例代碼可運行
- 及時更新文檔

### 4. 測試策略
- 公共 API 100% 覆蓋
- 注重邊界條件測試
- 保持向後兼容測試
- 定期性能基準測試

## 🎨 命令組合示例

### 快速原型
```bash
/sdk-design "功能名" && /plan "實現計劃" && /sdk-example basic
```

### 完整開發
```bash
/start && /sdk-design "API" && /plan && /check && /sdk-test all && /sdk-doc all
```

### 發布流程
```bash
/sdk-test all && /sdk-release check && /sdk-doc migration && /review
```

---

SDK 開發是一門藝術，需要在易用性、功能性和穩定性之間找到平衡。使用這些專用命令，讓 Claude Code 成為你的 SDK 開發夥伴！