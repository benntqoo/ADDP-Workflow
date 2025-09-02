#!/bin/bash

# 檢查文件路徑一致性腳本
# 確保所有核心文件都指向 .claude/ 目錄

echo "檢查核心文件路徑一致性..."
echo "================================"

# 核心文件列表
CORE_FILES=(
    "PROJECT_CONTEXT.md"
    "DECISIONS.md"
    "last-session.yml"
    "last-session.yaml"
    "CLAUDE.md"
)

# 顏色定義
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 檢查根目錄是否有這些文件（不應該有）
echo -e "\n${YELLOW}檢查根目錄...${NC}"
for file in "${CORE_FILES[@]}"; do
    if [ -f "$file" ]; then
        echo -e "${RED}✗ 發現根目錄文件: $file （應該在 .claude/ 目錄）${NC}"
    fi
done

# 檢查 .claude 目錄是否有這些文件（應該有）
echo -e "\n${YELLOW}檢查 .claude/ 目錄...${NC}"
for file in "${CORE_FILES[@]}"; do
    if [ -f ".claude/$file" ]; then
        echo -e "${GREEN}✓ 正確位置: .claude/$file${NC}"
    else
        if [ "$file" != "last-session.yaml" ]; then  # yaml 是舊格式，可以不存在
            echo -e "${YELLOW}○ 缺少文件: .claude/$file${NC}"
        fi
    fi
done

# 檢查所有代碼中的引用
echo -e "\n${YELLOW}檢查代碼中的引用...${NC}"

# 檢查錯誤引用（直接引用根目錄的文件）
echo -e "\n檢查錯誤引用（應該使用 .claude/ 前綴）："
wrong_refs=$(grep -r --include="*.md" --include="*.yml" --include="*.yaml" \
    -E "(^|[^.])(PROJECT_CONTEXT\.md|DECISIONS\.md|last-session\.(yml|yaml))" \
    . 2>/dev/null | \
    grep -v ".claude/" | \
    grep -v "^Binary" | \
    grep -v "check-file-consistency")

if [ -n "$wrong_refs" ]; then
    echo -e "${RED}發現以下文件可能有錯誤引用：${NC}"
    echo "$wrong_refs" | cut -d: -f1 | sort -u
else
    echo -e "${GREEN}✓ 未發現錯誤引用${NC}"
fi

# 檢查正確引用
echo -e "\n檢查正確引用（使用 .claude/ 前綴）："
correct_refs=$(grep -r --include="*.md" --include="*.yml" --include="*.yaml" \
    -E "\.claude/(PROJECT_CONTEXT\.md|DECISIONS\.md|last-session\.(yml|yaml)|CLAUDE\.md)" \
    . 2>/dev/null | \
    grep -v "^Binary" | \
    grep -v "check-file-consistency")

if [ -n "$correct_refs" ]; then
    echo -e "${GREEN}✓ 以下文件使用了正確的引用：${NC}"
    echo "$correct_refs" | cut -d: -f1 | sort -u | head -10
else
    echo -e "${YELLOW}○ 未找到引用這些文件的代碼${NC}"
fi

echo -e "\n================================"
echo "檢查完成！"

# 建議
echo -e "\n${YELLOW}建議：${NC}"
echo "1. 所有核心文件應該放在 .claude/ 目錄"
echo "2. 代碼中引用時應使用 '.claude/' 前綴"
echo "3. 定期運行此腳本確保一致性"