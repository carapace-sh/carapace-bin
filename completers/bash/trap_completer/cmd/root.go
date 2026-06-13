package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trap",
	Short: "Trap signals and other events",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-trap",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("P", "P", false, "display the trap commands associated with each signal_spec")
	rootCmd.Flags().BoolS("l", "l", false, "print a list of signal names and their corresponding numbers")
	rootCmd.Flags().BoolS("p", "p", false, "display the trap commands associated with each signal_spec in a reusable form")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.Batch(
			ps.ActionKillSignals(),
			carapace.ActionValuesDescribed(
				"DEBUG", "executed before every simple command",
				"ERR", "executed when a command's failure would cause the shell to exit",
				"EXIT", "executed on exit from the shell",
				"RETURN", "executed each time a shell function or script finishes",
				"SIGDEBUG", "executed before every simple command",
				"SIGERR", "executed when a command's failure would cause the shell to exit",
				"SIGEXIT", "executed on exit from the shell",
				"SIGRETURN", "executed each time a shell function or script finishes",
			),
		).ToA().FilterArgs(),
	)
}
