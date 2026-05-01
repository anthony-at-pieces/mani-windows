package integration

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	maniBin  string
	repoRoot string
)

func TestMain(m *testing.M) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(1)
	}

	repoRoot = filepath.Clean(filepath.Join(filepath.Dir(filename), "..", ".."))
	tmpDir, err := os.MkdirTemp("", "mani-integration-*")
	if err != nil {
		os.Exit(1)
	}

	maniBin = filepath.Join(tmpDir, "mani.exe")
	cmd := exec.Command("go", "build", "-o", maniBin, "main.go")
	cmd.Dir = repoRoot
	cmd.Env = append(os.Environ(), "CGO_ENABLED=0")
	if out, err := cmd.CombinedOutput(); err != nil {
		_, _ = os.Stderr.Write(out)
		os.Exit(1)
	}

	code := m.Run()
	_ = os.RemoveAll(tmpDir)
	os.Exit(code)
}

func runMani(t *testing.T, dir string, args ...string) (string, error) {
	t.Helper()

	cmd := exec.Command(maniBin, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "NO_COLOR=1")
	out, err := cmd.CombinedOutput()

	return string(out), err
}

func runGit(t *testing.T, dir string, args ...string) {
	t.Helper()

	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git %s failed: %v\n%s", strings.Join(args, " "), err, out)
	}
}

func writeFile(t *testing.T, path string, content string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		t.Fatalf("could not create %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("could not write %s: %v", path, err)
	}
}

func TestPowerShellTaskExecutionAndEnvInterpolation(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.Mkdir(filepath.Join(tmpDir, "app"), os.ModePerm); err != nil {
		t.Fatal(err)
	}

	writeFile(t, filepath.Join(tmpDir, "mani.yaml"), `
env:
  GREETING: $(Write-Output "Hello Windows")
projects:
  app:
    path: app
tasks:
  say:
    cmd: Write-Output $env:GREETING
`)

	out, err := runMani(t, tmpDir, "--color=false", "run", "say", "--all")
	if err != nil {
		t.Fatalf("mani run failed: %v\n%s", err, out)
	}
	if !strings.Contains(out, "Hello Windows") {
		t.Fatalf("expected PowerShell env output, got:\n%s", out)
	}
}

func TestConfigDiscoveryAndCwdFiltering(t *testing.T) {
	tmpDir := t.TempDir()
	appDir := filepath.Join(tmpDir, "app")
	if err := os.Mkdir(appDir, os.ModePerm); err != nil {
		t.Fatal(err)
	}

	writeFile(t, filepath.Join(tmpDir, "mani.yaml"), `
projects:
  app:
    path: app
tasks:
  cwd-check:
    cmd: Write-Output "cwd-ok"
`)

	out, err := runMani(t, appDir, "--color=false", "run", "cwd-check", "--cwd")
	if err != nil {
		t.Fatalf("mani run --cwd failed: %v\n%s", err, out)
	}
	if !strings.Contains(out, "cwd-ok") {
		t.Fatalf("expected cwd task output, got:\n%s", out)
	}
}

func TestPowerShellCompletionOnly(t *testing.T) {
	tmpDir := t.TempDir()

	out, err := runMani(t, tmpDir, "completion", "powershell")
	if err != nil {
		t.Fatalf("powershell completion failed: %v\n%s", err, out)
	}
	if !strings.Contains(out, "Register-ArgumentCompleter") {
		t.Fatalf("expected PowerShell completion script, got:\n%s", out)
	}

	out, err = runMani(t, tmpDir, "completion", "cmd")
	if err == nil {
		t.Fatalf("expected non-PowerShell completion to fail, got:\n%s", out)
	}
}

func TestLocalGitSyncAndWorktree(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git is not available")
	}

	tmpDir := t.TempDir()
	remoteDir := filepath.Join(tmpDir, "remote")
	if err := os.Mkdir(remoteDir, os.ModePerm); err != nil {
		t.Fatal(err)
	}

	runGit(t, remoteDir, "init", "-b", "main")
	runGit(t, remoteDir, "config", "user.email", "mani@example.local")
	runGit(t, remoteDir, "config", "user.name", "mani")
	writeFile(t, filepath.Join(remoteDir, "README.md"), "main\n")
	runGit(t, remoteDir, "add", "README.md")
	runGit(t, remoteDir, "commit", "-m", "initial")

	writeFile(t, filepath.Join(tmpDir, "mani.yaml"), strings.ReplaceAll(`
projects:
  app:
    path: app
    url: REMOTE_URL
    worktrees:
      - path: ../app-feature
        branch: feature
tasks:
  status:
    cmd: git status --short
`, "REMOTE_URL", filepath.ToSlash(remoteDir)))

	out, err := runMani(t, tmpDir, "--color=false", "sync")
	if err != nil {
		t.Fatalf("mani sync failed: %v\n%s", err, out)
	}

	if _, err := os.Stat(filepath.Join(tmpDir, "app", ".git")); err != nil {
		t.Fatalf("expected cloned project .git: %v", err)
	}
	if _, err := os.Stat(filepath.Join(tmpDir, "app-feature", ".git")); err != nil {
		t.Fatalf("expected worktree .git: %v", err)
	}
}
