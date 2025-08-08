# SDK 項目 CLAUDE.md 模板

<!-- 
這是專為 SDK/Library 開發設計的 CLAUDE.md 模板
根據你的 SDK 特性調整內容
-->

## 🎯 SDK 概述

**SDK 名稱**：[Your SDK Name]  
**版本**：[Current Version]  
**目標**：[一句話說明 SDK 解決什麼問題]

## 🏗️ 項目結構

```
[sdk-name]/
├── src/           # 源代碼
├── tests/         # 測試文件
├── examples/      # 示例代碼
├── docs/          # 文檔
└── dist/          # 構建輸出
```

## 📋 開發規範

### API 設計原則
1. **一致性優先**：相似功能使用相似的 API 設計
2. **最小驚訝**：API 行為要符合直覺
3. **漸進式複雜度**：簡單任務簡單做，複雜任務也可能
4. **錯誤友好**：提供清晰、可操作的錯誤信息

### 命名規範
- **類名**：PascalCase（如 `CacheManager`）
- **方法名**：camelCase（如 `getValue`）
- **常量**：UPPER_SNAKE_CASE（如 `MAX_SIZE`）
- **私有成員**：下劃線前綴（如 `_internalMethod`）

### 版本管理
- 遵循語義化版本 (SemVer)
- 破壞性變更需要主版本號增加
- 新功能需要次版本號增加
- Bug 修復需要補丁版本號增加

## 🔄 開發工作流

### 新功能開發
1. 使用 `/sdk-design` 設計 API
2. 使用 `/plan` 規劃實現
3. 編寫代碼和測試
4. 使用 `/sdk-example` 創建示例
5. 使用 `/sdk-doc` 更新文檔

### 發布流程
1. 使用 `/sdk-test all` 運行完整測試
2. 使用 `/sdk-release check` 檢查發布就緒
3. 更新 CHANGELOG.md
4. 使用 `/sdk-release [version]` 準備發布

## 💡 Claude 行為配置

### 代碼生成時
- 總是包含完整的 JSDoc/DocString
- 為公共 API 生成使用示例
- 考慮向後兼容性
- 包含錯誤處理

### 測試生成時
- 優先測試公共 API
- 包含邊界條件測試
- 測試錯誤場景
- 生成性能基準測試

### 文檔生成時
- 從用戶角度編寫
- 包含實際使用案例
- 解釋設計決策
- 提供故障排除指南

## 🎨 代碼示例模板

### API 方法模板
```typescript
/**
 * 方法簡短描述
 * 
 * @param {Type} paramName - 參數描述
 * @returns {ReturnType} 返回值描述
 * @throws {ErrorType} 錯誤情況描述
 * 
 * @example
 * ```typescript
 * const result = sdk.methodName(param);
 * ```
 * 
 * @since 1.0.0
 */
public methodName(paramName: Type): ReturnType {
    // 參數驗證
    // 核心邏輯
    // 錯誤處理
    // 返回結果
}
```

### 測試模板
```typescript
describe('ClassName', () => {
    describe('methodName', () => {
        it('should handle normal case', () => {
            // Arrange
            // Act
            // Assert
        });

        it('should handle edge case', () => {
            // 邊界條件測試
        });

        it('should throw error when...', () => {
            // 錯誤場景測試
        });
    });
});
```

## 📚 文檔結構

**SDK 文檔統一存放在 `.claude/docs/` 目錄下**：

- **README.md**：快速開始和基本信息
- **.claude/docs/api/**：API 參考文檔
- **.claude/docs/guides/**：使用指南和教程
- **.claude/docs/examples/**：更多示例代碼
- **CHANGELOG.md**：版本變更歷史
- **CONTRIBUTING.md**：貢獻指南

## ⚠️ 注意事項

### 破壞性變更處理
1. 先標記為 `@deprecated`
2. 提供遷移指南
3. 至少保留一個主版本
4. 在 CHANGELOG 中明確說明

### 性能考慮
- 避免同步阻塞操作
- 提供批量操作 API
- 實現合理的緩存策略
- 定期進行性能測試

### 安全考慮
- 驗證所有輸入
- 避免敏感信息洩露
- 使用安全的默認配置
- 及時更新依賴

## 🔧 開發工具

### 必需工具
- 構建工具：[Webpack/Rollup/etc]
- 測試框架：[Jest/Mocha/etc]
- 文檔工具：[TypeDoc/JSDoc/etc]
- 代碼檢查：[ESLint/TSLint/etc]

### 有用的命令
```bash
npm run build      # 構建 SDK
npm run test       # 運行測試
npm run docs       # 生成文檔
npm run release    # 發布新版本
```

---

<!-- SDK 特定配置區域 -->
## 🎯 SDK 特定配置

[根據你的 SDK 添加特定配置]

---

*使用 Claude Code SDK 命令讓開發更高效！*