package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var processes_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(processes_helpCmd).Standalone()

	processesCmd.AddCommand(processes_helpCmd)

	carapace.Gen(processes_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(processesCmd),
	)
}
