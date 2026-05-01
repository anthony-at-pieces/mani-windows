package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/alajmo/mani/core/dao"
)

const (
	appName      = "mani"
	shortAppDesc = "repositories manager and task runner"
)

var (
	config         dao.Config
	configErr      error
	configFilepath string
	userConfigPath string
	color          bool
	version        = "dev"
	commit         = "none"
	date           = "n/a"
	rootCmd        = &cobra.Command{
		Use:     appName,
		Short:   shortAppDesc,
		Version: version,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// When user input's wrong command or flag
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&configFilepath, "config", "c", "", "specify config")
	rootCmd.PersistentFlags().StringVarP(&userConfigPath, "user-config", "u", "", "specify user config")
	rootCmd.PersistentFlags().BoolVar(&color, "color", true, "enable color")

	rootCmd.AddCommand(
		completionCmd(),
		initCmd(),
		execCmd(&config, &configErr),
		runCmd(&config, &configErr),
		listCmd(&config, &configErr),
		describeCmd(&config, &configErr),
		syncCmd(&config, &configErr),
		editCmd(&config, &configErr),
		checkCmd(&configErr),
		tuiCmd(&config, &configErr),
	)

	rootCmd.SetVersionTemplate(fmt.Sprintf("Version: %-10s\nCommit: %-10s\nDate: %-10s\n", version, commit, date))

	// Add custom help template with footer
	defaultHelpTemplate := rootCmd.HelpTemplate()
	rootCmd.SetHelpTemplate(defaultHelpTemplate + `
Documentation: https://manicli.com
Issues:        https://github.com/alajmo/mani/issues
`)

	rootCmd.DisableAutoGenTag = true
}

func initConfig() {
	config, configErr = dao.ReadConfig(configFilepath, userConfigPath, color)
}
