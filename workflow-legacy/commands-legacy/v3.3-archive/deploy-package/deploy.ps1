# Claude Code 命令自动部署脚本
# 适用于 Windows PowerShell

Write-Host "🚀 Claude Code 命令部署脚本" -ForegroundColor Cyan
Write-Host "==========================" -ForegroundColor Cyan

# 设置命令目录路径
$GlobalCmdDir = "$env:USERPROFILE\.claude\commands"
$ProjectCmdDir = ".\.claude\commands"

# 创建全局命令目录
Write-Host ""
Write-Host "📁 创建全局命令目录..." -ForegroundColor Yellow
if (!(Test-Path $GlobalCmdDir)) {
    New-Item -ItemType Directory -Force -Path $GlobalCmdDir | Out-Null
}
if (Test-Path $GlobalCmdDir) {
    Write-Host "✓ 全局命令目录创建成功: $GlobalCmdDir" -ForegroundColor Green
} else {
    Write-Host "❌ 创建全局命令目录失败" -ForegroundColor Red
    exit 1
}

# 复制全局命令
Write-Host ""
Write-Host "📋 复制全局命令..." -ForegroundColor Yellow
try {
    Copy-Item -Path ".\global\*.md" -Destination $GlobalCmdDir -Force
    $globalCount = (Get-ChildItem -Path ".\global\*.md").Count
    Write-Host "✓ 成功复制 $globalCount 个全局命令" -ForegroundColor Green
} catch {
    Write-Host "❌ 复制全局命令失败: $_" -ForegroundColor Red
    exit 1
}

# 询问是否部署项目命令
Write-Host ""
$response = Read-Host "是否要部署项目命令到当前目录? (y/n)"
if ($response -eq 'y' -or $response -eq 'Y') {
    # 创建项目命令目录
    Write-Host "📁 创建项目命令目录..." -ForegroundColor Yellow
    if (!(Test-Path $ProjectCmdDir)) {
        New-Item -ItemType Directory -Force -Path $ProjectCmdDir | Out-Null
    }
    if (Test-Path $ProjectCmdDir) {
        Write-Host "✓ 项目命令目录创建成功: $ProjectCmdDir" -ForegroundColor Green
    } else {
        Write-Host "❌ 创建项目命令目录失败" -ForegroundColor Red
        exit 1
    }
    
    # 复制项目命令
    Write-Host "📋 复制项目命令..." -ForegroundColor Yellow
    try {
        Copy-Item -Path ".\project\*.md" -Destination $ProjectCmdDir -Force
        $projectCount = (Get-ChildItem -Path ".\project\*.md").Count
        Write-Host "✓ 成功复制 $projectCount 个项目命令" -ForegroundColor Green
    } catch {
        Write-Host "❌ 复制项目命令失败: $_" -ForegroundColor Red
        exit 1
    }
}

# 验证部署
Write-Host ""
Write-Host "🔍 验证部署结果..." -ForegroundColor Yellow
$globalDeployed = (Get-ChildItem -Path "$GlobalCmdDir\*.md" -ErrorAction SilentlyContinue).Count
Write-Host "全局命令数量: $globalDeployed" -ForegroundColor Cyan
if (Test-Path $ProjectCmdDir) {
    $projectDeployed = (Get-ChildItem -Path "$ProjectCmdDir\*.md" -ErrorAction SilentlyContinue).Count
    Write-Host "项目命令数量: $projectDeployed" -ForegroundColor Cyan
}

Write-Host ""
Write-Host "✨ 部署完成！" -ForegroundColor Green
Write-Host ""
Write-Host "提示：" -ForegroundColor Yellow
Write-Host "1. 请重启 Claude Code 以加载新命令"
Write-Host "2. 使用 /meta 命令开始定制项目规范"
Write-Host "3. 查看 DEPLOY_GUIDE.md 了解更多信息"
Write-Host ""