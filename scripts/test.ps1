param(
    [switch]$SkipSmoke
)

$ErrorActionPreference = "Stop"

go test ./...

.\scripts\build.ps1

if (-not $SkipSmoke) {
    .\dist\mani.exe --help | Out-Null
    .\dist\mani.exe --version | Out-Null
    .\dist\mani.exe completion powershell | Out-Null
}
