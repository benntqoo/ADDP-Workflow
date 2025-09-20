# Windows啟動問題解決方案

## 立即測試步驟

### 1. 檢查依賴項
請運行以下命令檢查系統依賴：

```cmd
# 檢查程序信息
cd "D:\Code\fos\AI\claude"
dir ai-launcher.exe

# 嘗試運行並捕獲錯誤
ai-launcher.exe > output.txt 2>&1
type output.txt
```

### 2. 安裝必要的運行時
下載並安裝以下組件：

1. **Visual C++ Redistributable** (最重要)
   - 下載地址: https://aka.ms/vs/17/release/vc_redist.x64.exe
   - 安裝完成後重試

2. **更新顯卡驅動**
   - NVIDIA: https://www.nvidia.com/Download/index.aspx
   - AMD: https://www.amd.com/support
   - Intel: https://www.intel.com/content/www/us/en/support/

### 3. 重新編譯（如果上述方法無效）
使用正確的編譯標誌：

```powershell
# 使用PowerShell執行
docker run --rm -v "${PWD}:/workspace" -w /workspace golang:1.23-bullseye bash -c "
  apt-get update -qq &&
  apt-get install -y gcc-mingw-w64 pkg-config libxxf86vm-dev &&
  go mod download &&
  CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -ldflags='-H windowsgui' -o ai-launcher-fixed.exe ./cmd/gui
"
```

### 4. 診斷模式
如果仍然有問題，創建診斷版本：

```powershell
# 編譯控制台版本用於調試
docker run --rm -v "${PWD}:/workspace" -w /workspace golang:1.23-bullseye bash -c "
  apt-get update -qq &&
  apt-get install -y gcc-mingw-w64 pkg-config libxxf86vm-dev &&
  go mod download &&
  CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o ai-launcher-console.exe ./cmd/gui
"
```

### 5. 系統兼容性測試
```cmd
# 以管理員身份運行
右鍵點擊程序 -> "以管理員身份運行"

# 兼容模式
右鍵程序 -> 屬性 -> 兼容性 -> 選擇 "Windows 10"
```

## 最可能的解決方案

根據症狀分析，最有可能的問題是：

1. **缺少Visual C++ Redistributable** (70%概率)
2. **OpenGL驅動過舊** (20%概率)
3. **編譯標誌問題** (10%概率)

## 快速修復
請按順序嘗試：

1. 安裝VC++ Redistributable
2. 更新顯卡驅動
3. 以管理員身份運行
4. 重新編譯程序

如果以上方法都無效，請運行診斷程序並提供輸出結果。