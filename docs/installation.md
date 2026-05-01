# Installation

`mani` is available as a Windows-first CLI and TUI for PowerShell users.

* Windows ZIP archives are available on the [release](https://github.com/alajmo/mani/releases) page. Extract the archive and place `mani.exe` on your `PATH`.

* Via Go
  ```powershell
  go install github.com/alajmo/mani@latest
  ```

## Building From Source

1. Clone the repo
2. Build and run the executable

    ```powershell
    .\scripts\build.ps1
    .\dist\mani.exe --help
    ```

## Packaging

```powershell
.\scripts\package.ps1 -Version dev
```
