param(
    [string]$Version = "0.2.0",
    [string]$InstallDir = (Join-Path $env:LOCALAPPDATA 'Programs\mani'),
    [switch]$SkipBuild,
    [switch]$NoPath
)

$ErrorActionPreference = "Stop"

$repoRoot = Split-Path -Parent $PSScriptRoot
$binary = Join-Path $repoRoot 'dist\mani.exe'

if (-not $SkipBuild) {
    & (Join-Path $PSScriptRoot 'build.ps1') -Version $Version -Output $binary
}

if (-not (Test-Path $binary)) {
    throw "mani.exe not found at $binary. Run without -SkipBuild or build first."
}

New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
$target = Join-Path $InstallDir 'mani.exe'
Copy-Item -Force $binary $target
Write-Host "Installed: $target"

if ($NoPath) { return }

$userPath = [Environment]::GetEnvironmentVariable('Path', 'User')
$entries = if ($userPath) { $userPath.Split(';') | Where-Object { $_ } } else { @() }

if ($entries -contains $InstallDir) {
    Write-Host "PATH: $InstallDir already on user PATH"
} else {
    $newPath = if ($userPath) { "$userPath;$InstallDir" } else { $InstallDir }
    [Environment]::SetEnvironmentVariable('Path', $newPath, 'User')
    Write-Host "PATH: added $InstallDir to user PATH (open a new shell to pick it up)"
}
