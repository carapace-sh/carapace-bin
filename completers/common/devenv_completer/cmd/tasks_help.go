package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tasks_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_helpCmd).Standalone()

	tasksCmd.AddCommand(tasks_helpCmd)

	carapace.Gen(tasks_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(tasksCmd),
	)
}
