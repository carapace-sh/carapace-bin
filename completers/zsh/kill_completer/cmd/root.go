package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kill",
	Short: "Send a signal to a process",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-kill",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"-TERM", "terminate signal (default)",
			"-HUP", "hangup signal",
			"-INT", "interrupt signal",
			"-KILL", "kill signal (cannot be caught)",
			"-STOP", "stop process",
			"-CONT", "continue stopped process",
		),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValuesDescribed(
			"HUP", "hangup signal",
			"INT", "interrupt signal",
			"KILL", "kill signal (cannot be caught)",
			"QUIT", "quit signal",
			"TERM", "terminate signal",
			"STOP", "stop process",
			"CONT", "continue stopped process",
			"USR1", "user-defined signal 1",
			"USR2", "user-defined signal 2",
			"PIPE", "broken pipe signal",
		),
	)
}
