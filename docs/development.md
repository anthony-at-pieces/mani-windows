# Development

## Build instructions

### Prerequisites

- [go 1.25 or above](https://golang.org/doc/install)
- Git

### Building

```powershell
# Build mani for your platform target
.\scripts\build.ps1

# Run unit, integration, and smoke tests
.\scripts\test.ps1

# Create a Windows ZIP package
.\scripts\package.ps1 -Version dev
```

## Developing

```powershell
# Format code
gofmt -w .\cmd .\core .\test\integration

# Manage dependencies (download/remove unused)
go mod tidy

# Lint code
golangci-lint run ./...

# Debug completion
.\dist\mani.exe __complete list tags --projects ""
```

## Releasing

The following workflow is used for releasing a new `mani` version:

1. Create pull request with changes
2. Verify the Windows build works
   - `.\scripts\build.ps1`
3. Pass all integration and unit tests locally
   - `.\scripts\test.ps1`
4. Update docs if any config or command behavior changes
5. Update `CHANGELOG.md` with the release notes
6. Tag the release with `vX.Y.Z`; GitHub Actions packages `mani.exe` into a Windows ZIP archive
