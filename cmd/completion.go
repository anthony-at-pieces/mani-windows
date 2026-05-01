package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/alajmo/mani/core"
)

func completionCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "completion powershell",
		Short: "Generate PowerShell completion script",
		Long: `To load completions:
PowerShell:

  PS> mani completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> mani completion powershell > mani.ps1
  # and source this file from your PowerShell profile.
		`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"powershell"},
		Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run:                   generateCompletion,
		DisableAutoGenTag:     true,
	}

	return &cmd
}

func generateCompletion(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "powershell":
		err := cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		core.CheckIfError(err)
	}
}
