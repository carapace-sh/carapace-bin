package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "kill",
	Short:              "Send a signal to a process",
	Long:               "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-kill",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		ps.ActionKillSignals().Prefix("-"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		ps.ActionProcessIds().FilterArgs().Shift(1),
	)
}
