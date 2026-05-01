# Test

`mani` uses native Windows unit and integration tests. The integration tests build `mani.exe` into a temporary directory and exercise PowerShell task execution, config discovery, completion generation, and local git sync/worktree behavior.

## Directory Structure

```sh
.
├── fixtures    # files needed for testing purposes
└── integration # native Windows integration tests
```

## Prerequisites

- Go
- Git
- [golangci-lint](https://golangci-lint.run)

## Testing & Development

Run tests from the repository root in PowerShell.

```powershell
# Run tests
.\scripts\test.ps1

# Run Go tests directly
go test ./...

# Debug completion
mani __complete list tags --projects ""

# Build and run a smoke check
.\scripts\build.ps1
.\dist\mani.exe --version
```
