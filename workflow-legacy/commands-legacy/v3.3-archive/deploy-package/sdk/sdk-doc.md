---
arguments: optional
format: "[api|guide|migration|all]"
examples:
  - "/sdk-doc - 生成完整文檔"
  - "/sdk-doc api - 只生成 API 參考"
  - "/sdk-doc guide - 生成使用指南"
  - "/sdk-doc migration - 生成遷移文檔"
---

# SDK 文檔專家

為 SDK 生成專業、完整、易懂的文檔：

## 📚 文檔類型

### 1. API 參考文檔 (api)
- **完整性**：所有公共 API
- **結構化**：按模塊/功能分組
- **詳細度**：參數、返回值、異常
- **示例化**：每個 API 都有示例

### 2. 使用指南 (guide)
- **快速開始**：5 分鐘上手
- **核心概念**：理解設計理念
- **最佳實踐**：推薦用法
- **常見問題**：FAQ

### 3. 遷移指南 (migration)
- **版本對比**：變更詳情
- **遷移步驟**：漸進式升級
- **破壞性變更**：詳細說明
- **自動化工具**：遷移腳本

## 📋 SDK 文檔結構標準

**所有 SDK 文檔統一存放在 `.claude/docs/` 目錄下**：

```
.claude/docs/
├── README.md          # SDK 文檔首頁
├── getting-started/   # 快速開始
│   ├── installation.md
│   ├── quick-start.md
│   └── first-app.md
├── api/              # API 參考
│   ├── core/         # 核心 API
│   ├── utilities/    # 工具函數
│   └── advanced/     # 高級功能
├── guides/           # 使用指南
│   ├── concepts.md   # 核心概念
│   ├── patterns.md   # 設計模式
│   └── best-practices.md
├── examples/         # 示例代碼
├── migration/        # 遷移指南
│   └── v1-to-v2.md
└── reference/        # 參考資料
    ├── config.md     # 配置選項
    ├── errors.md     # 錯誤代碼
    └── glossary.md   # 術語表
```

## 🎨 文檔模板

### API 文檔模板
```markdown
# 類/模塊名稱

簡短描述功能用途。

## 導入方式
\`\`\`[語言]
import { ClassName } from 'sdk-name';
\`\`\`

## 構造函數/初始化
\`\`\`[語言]
new ClassName(options)
\`\`\`

### 參數
| 參數 | 類型 | 必需 | 默認值 | 描述 |
|------|------|------|--------|------|
| options | Object | 是 | - | 配置選項 |
| options.key | string | 是 | - | API 密鑰 |
| options.timeout | number | 否 | 5000 | 超時時間(ms) |

### 示例
\`\`\`[語言]
const instance = new ClassName({
  key: 'your-api-key',
  timeout: 10000
});
\`\`\`

## 方法

### methodName(param1, param2)
方法描述。

#### 參數
- \`param1\` (Type) - 參數描述
- \`param2\` (Type, 可選) - 參數描述

#### 返回值
- \`Promise<ReturnType>\` - 返回值描述

#### 異常
- \`ErrorType\` - 錯誤情況說明

#### 示例
\`\`\`[語言]
const result = await instance.methodName('value1', {
  option: true
});
\`\`\`
```

### 使用指南模板
```markdown
# 使用指南

## 核心概念

### 概念名稱
解釋核心概念，使用類比和圖表。

## 常見用例

### 用例 1：基本使用
步驟說明和代碼示例。

### 用例 2：高級功能
複雜場景的解決方案。

## 最佳實踐

### 性能優化
- 使用連接池
- 批量操作
- 緩存策略

### 錯誤處理
- 重試機制
- 優雅降級
- 日誌記錄

## 故障排除

### 常見錯誤
問題描述和解決方案。
```

## 💡 文檔最佳實踐

### 1. 寫作原則
- **用戶視角**：從使用者角度寫
- **循序漸進**：從簡單到複雜
- **示例豐富**：代碼勝於描述
- **保持更新**：與代碼同步

### 2. 格式規範
- **一致性**：統一的格式和術語
- **可讀性**：短句、段落、列表
- **可搜索**：良好的標題結構
- **可複製**：代碼塊可直接使用

### 3. 內容組織
- **快速導航**：目錄和錨點
- **交叉引用**：相關內容鏈接
- **版本標註**：API 可用版本
- **平台說明**：兼容性信息

## 🔄 文檔維護

### 自動化生成
- API 文檔從代碼註釋生成
- 示例代碼自動測試
- 版本變更自動對比
- 破壞性變更高亮

### 審查流程
- [ ] 技術準確性
- [ ] 示例可運行
- [ ] 格式一致性
- [ ] 完整性檢查

優秀的文檔是 SDK 成功的關鍵！