//go:build windows
// +build windows

package exec

import (
	"os"
	"os/exec"
)

func ExecTTY(shell string, args []string, envs []string) error {
	shellPath, err := exec.LookPath(shell)
	if err != nil {
		return err
	}

	userEnv := append(os.Environ(), envs...)
	command := exec.Command(shellPath, args...)
	command.Env = userEnv
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command.Run()
}
