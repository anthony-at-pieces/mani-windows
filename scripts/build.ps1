param(
    [string]$Version = "dev",
    [string]$Output = "dist\mani.exe"
)

$ErrorActionPreference = "Stop"

$commit = "none"
if (Test-Path ".git") {
    $commit = (git rev-parse --short HEAD).Trim()
}

$date = (Get-Date).ToUniversalTime().ToString("yyyy-MM-ddTHH:mm:ssZ")
$package = "github.com/alajmo/mani"
$outputDir = Split-Path -Parent $Output

if ($outputDir) {
    New-Item -ItemType Directory -Force -Path $outputDir | Out-Null
}

$env:CGO_ENABLED = "0"
go build `
    -ldflags "-w -X '$package/cmd.version=$Version' -X '$package/core/tui/views.Version=$Version' -X '$package/cmd.commit=$commit' -X '$package/cmd.date=$date'" `
    -a `
    -tags netgo `
    -o $Output `
    main.go
