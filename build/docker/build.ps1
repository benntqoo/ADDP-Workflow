param(
    [ValidateSet("windows", "linux")]
    [string]$Target = "windows",

    [string]$OutputDir = "dist",

    [switch]$KeepImage
)

$ErrorActionPreference = "Stop"

function Invoke-Docker {
    param(
        [Parameter(Mandatory = $true, Position = 0)]
        [string[]]$Args
    )

    $output = & docker @Args
    $exitCode = $LASTEXITCODE
    if ($exitCode -ne 0) {
        throw "docker $($Args -join ' ') failed with exit code $exitCode."
    }
    return $output
}

if (-not (Get-Command docker -ErrorAction SilentlyContinue)) {
    throw "docker command not found. Please install Docker Desktop or CLI."
}

$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$repoRoot  = [System.IO.Path]::GetFullPath([System.IO.Path]::Combine($scriptDir, "..", ".."))
$outputPath = [System.IO.Path]::GetFullPath((New-Item -ItemType Directory -Path $OutputDir -Force).FullName)

switch ($Target) {
    "windows" {
        $dockerfile   = Join-Path $scriptDir "Dockerfile.windows"
        $imageTag     = "ai-launcher:windows"
        $artifactPath = "/ai-launcher.exe"
        $artifactName = "ai-launcher.exe"
    }
    "linux" {
        $dockerfile   = Join-Path $scriptDir "Dockerfile"
        $imageTag     = "ai-launcher:linux"
        $artifactPath = "/app/ai-launcher"
        $artifactName = "ai-launcher"
    }
    default {
        throw "Unsupported target '$Target'."
    }
}

Write-Host "Building Docker image '$imageTag' using $dockerfile" -ForegroundColor Cyan
Invoke-Docker @("build", "--file", $dockerfile, "--tag", $imageTag, $repoRoot) | Out-Null

$containerName = "ai-launcher-build-" + ([Guid]::NewGuid().ToString("N").Substring(0, 12))
$containerCreated = $false
try {
    Write-Host "Creating temporary container $containerName" -ForegroundColor Cyan
    Invoke-Docker @("create", "--name", $containerName, $imageTag) | Out-Null
    $containerCreated = $true

    $destination = Join-Path $outputPath $artifactName
    Write-Host "Copying artifact to $destination" -ForegroundColor Cyan
    Invoke-Docker @("cp", ("{0}:{1}" -f $containerName, $artifactPath), $destination) | Out-Null
}
finally {
    if ($containerCreated -and (Get-Command docker -ErrorAction SilentlyContinue)) {
        try { Invoke-Docker @("rm", $containerName) | Out-Null } catch { }
    }
}

if (-not $KeepImage.IsPresent) {
    Write-Host "Removing intermediate image $imageTag" -ForegroundColor DarkGray
    try { Invoke-Docker @("rmi", $imageTag) | Out-Null } catch { }
}

Write-Host "Build complete. Artifact located at $(Join-Path $outputPath $artifactName)" -ForegroundColor Green
