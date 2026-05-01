param(
    [string]$Version = "dev",
    [string]$OutputDir = "dist"
)

$ErrorActionPreference = "Stop"

.\scripts\build.ps1 -Version $Version -Output "$OutputDir\mani.exe"

$packageRoot = Join-Path $OutputDir "mani-$Version-windows"
if (Test-Path $packageRoot) {
    Remove-Item -Recurse -Force $packageRoot
}

New-Item -ItemType Directory -Force -Path $packageRoot | Out-Null
Copy-Item "$OutputDir\mani.exe" $packageRoot
Copy-Item "README.md" $packageRoot
Copy-Item "LICENSE" $packageRoot
Copy-Item -Recurse "docs" $packageRoot

$zipPath = Join-Path $OutputDir "mani-$Version-windows-amd64.zip"
if (Test-Path $zipPath) {
    Remove-Item -Force $zipPath
}

Compress-Archive -Path (Join-Path $packageRoot "*") -DestinationPath $zipPath
Write-Host $zipPath
