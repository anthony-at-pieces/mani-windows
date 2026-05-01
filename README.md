<h1 align="center"><code>mani-windows</code></h1>

<div align="center">
  <a href="https://github.com/anthony-at-pieces/mani-windows/releases">
    <img src="https://img.shields.io/github/v/release/anthony-at-pieces/mani-windows?include_prereleases" alt="release">
  </a>

  <a href="LICENSE">
    <img src="https://img.shields.io/badge/license-MIT-green" alt="license">
  </a>

  <img src="https://img.shields.io/badge/platform-windows--amd64-blue" alt="platform">
</div>

<br>

> **Windows-only fork of [`alajmo/mani`](https://github.com/alajmo/mani).** Same `mani` you know — manage many git repositories from one config, run commands across them in parallel — repackaged with native Windows tooling: PowerShell build/test/package scripts, a one-shot install script, and Windows-native integration tests. No WSL, MSYS, or Cygwin required.

`mani` lets you pull, organise, and run commands across multiple git repositories from a single `mani.yaml`. Useful for microservices, polyrepo setups, or any time you find yourself running the same `git` or build command in five places at once.

![demo](res/demo.gif)

## What this fork changes

Compared to upstream [`alajmo/mani`](https://github.com/alajmo/mani):

**Default shell is PowerShell, not bash.** `DEFAULT_SHELL` is now `powershell -NoProfile -Command` and `FormatShell` recognises `powershell`, `pwsh`, and `cmd` (`/C`) in addition to the upstream `bash`/`zsh`/`sh`/`node`/`python` set. Tasks in `mani.yaml` run via PowerShell unless you set `shell:` explicitly. The default editor (when `$EDITOR` is unset) falls back to `notepad`.

**Real `ExecTTY` on Windows.** Upstream's [`core/exec/windows.go`](core/exec/windows.go) was a stub that returned `nil`. This fork implements it via `exec.LookPath` + `os.Stdin/Stdout/Stderr` so `mani exec --tty`-style flows actually attach to the user's terminal on Windows.

**PowerShell build pipeline replaces the Makefile.**

- `Makefile`, `install.sh`, `scripts/release.sh`, `.goreleaser.yaml`, `.dockerignore` — **removed**.
- `scripts/build.ps1`, `scripts/test.ps1`, `scripts/package.ps1`, `scripts/install.ps1` — **added**.
- Man-page generation (`core/man.go`, `core/man_gen.go`, `core/mani.1`, `core/config.man`, `cmd/gen.go`, `cmd/gen_docs.go`) — **removed**. Use `mani --help` and the [`docs/`](docs/) folder instead.

**CI runs on `windows-latest` only.** Both `.github/workflows/test.yml` and `.github/workflows/release.yml` drop the Ubuntu matrix. Releases are produced by `scripts/package.ps1` and uploaded with `softprops/action-gh-release` instead of GoReleaser/Homebrew tap.

**Integration tests are Windows-native.** Upstream's Docker-based suite (`test/integration/{init,list,run,sync,version,exec,describe}_test.go` plus the Alpine `Dockerfile`s and shell-script harnesses under `test/scripts/`) is replaced by a single [`test/integration/windows_test.go`](test/integration/windows_test.go) that runs under PowerShell with no containers.

**Install target is `%LOCALAPPDATA%\Programs\mani`.** [`scripts/install.ps1`](scripts/install.ps1) builds, copies, and (idempotently) adds that directory to the user PATH.

Everything else — projects, tasks, specs, themes, the TUI, `sync`, `exec`, `run`, the YAML config schema — tracks upstream. See [`docs/`](docs/).

## Install

### Recommended: install script

```powershell
.\scripts\install.ps1
```

Builds `dist\mani.exe`, copies it to `%LOCALAPPDATA%\Programs\mani`, and adds that directory to your user PATH. Open a new shell afterwards to pick up the PATH change.

Flags:

- `-InstallDir <path>` — install somewhere else
- `-SkipBuild` — reuse an existing `dist\mani.exe` instead of rebuilding
- `-NoPath` — skip the PATH modification
- `-Version <ver>` — value baked into the binary's `--version`

If PowerShell blocks the script with an execution-policy error:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\install.ps1
```

### Pre-built binary

Download `mani-<version>-windows-amd64.zip` from the [releases page](https://github.com/anthony-at-pieces/mani-windows/releases), unzip, and put `mani.exe` somewhere on your `PATH`.

### Build from source

```powershell
.\scripts\build.ps1
.\dist\mani.exe --help
```

Tab completion: `mani completion powershell`.

## Usage

### Initialize

Run inside a directory containing your git repositories:

```powershell
mani init
```

This generates:

- `mani.yaml` — projects and tasks. Any subdirectory with a `.git` folder is added automatically (turn off with `--auto-discovery=false`).
- `.gitignore` — when run inside a git repo, the listed projects are added (opt out with `--sync-gitignore=false`).

### Examples

```powershell
# List all projects
mani list projects

# Run git status across every project
mani exec --all git status

# Run in parallel with table output
mani exec --all --parallel --output table git status
```

## Documentation

- [Examples](examples)
- [Config reference](docs/config.md)
- [Commands](docs/commands.md)
- [Filtering projects](docs/filtering-projects.md)
- [Variables](docs/variables.md)
- [Output formats](docs/output.md)
- [Changelog](docs/changelog.md)
- [Roadmap](docs/roadmap.md)
- [Project background](docs/project-background.md)
- [Contributing](docs/contributing.md)

## Credits

Original project: [`alajmo/mani`](https://github.com/alajmo/mani) by Samir Alajmovic. This fork ports the build, packaging, and install flow to Windows / PowerShell while tracking upstream behaviour.

## [License](LICENSE)

MIT — see [LICENSE](LICENSE). Copyright (c) 2020-2021 Samir Alajmovic.
